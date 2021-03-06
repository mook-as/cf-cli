package command

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"code.cloudfoundry.org/cli/util"
	logcache "code.cloudfoundry.org/log-cache/pkg/client"
)

type RequestLoggerOutput interface {
	Start() error
	Stop() error
	DisplayType(name string, requestDate time.Time) error
	DisplayDump(dump string) error
}

type DebugPrinter struct {
	outputs []RequestLoggerOutput
}

func (p DebugPrinter) Print(title, dump string) {
	for _, output := range p.outputs {
		_ = output.Start()                        //nolint
		_ = output.DisplayType(title, time.Now()) //nolint
		_ = output.DisplayDump(dump)              //nolint
		_ = output.Stop()                         //nolint
	}
}

func (p *DebugPrinter) addOutput(output RequestLoggerOutput) {
	p.outputs = append(p.outputs, output)
}

type tokenHTTPClient struct {
	c           logcache.HTTPClient
	accessToken func() string
}

func (c *tokenHTTPClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", c.accessToken())
	return c.c.Do(req)
}

type httpDebugClient struct {
	printer DebugPrinter
	c       logcache.HTTPClient
}

func (c *httpDebugClient) Do(req *http.Request) (*http.Response, error) {
	c.printer.Print("HTTP REQUEST",
		fmt.Sprintf("GET %s HTTP/1.1\nHOST: %s://%s\nQUERY: %s\n%s",
			req.URL.Path,
			req.URL.Scheme,
			req.URL.Host,
			req.URL.RawQuery,
			headersString(req.Header)),
	)

	resp, err := c.c.Do(req)
	if err != nil {
		c.printer.Print("HTTP ERROR", err.Error())
		return nil, err
	}

	c.printer.Print("HTTP RESPONSE",
		fmt.Sprintf("%s\n%s",
			resp.Status,
			headersString(resp.Header)),
	)

	return resp, err
}

// NewLogCacheClient returns back a configured Log Cache Client.
func NewLogCacheClient(logCacheEndpoint string, config Config, ui UI) *logcache.Client {
	tr := &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: util.NewTLSConfig(nil, config.SkipSSLValidation()),
		DialContext: (&net.Dialer{
			KeepAlive: 30 * time.Second,
			Timeout:   config.DialTimeout(),
		}).DialContext,
	}

	var client logcache.HTTPClient //nolint
	client = &http.Client{Transport: tr}

	verbose, location := config.Verbose()
	if verbose {
		printer := DebugPrinter{}
		printer.addOutput(ui.RequestLoggerTerminalDisplay())
		if location != nil {
			printer.addOutput(ui.RequestLoggerFileWriter(location))
		}

		client = &httpDebugClient{printer: printer, c: client}
	}

	return logcache.NewClient(
		logCacheEndpoint,
		logcache.WithHTTPClient(&tokenHTTPClient{
			c:           client,
			accessToken: config.AccessToken,
		}),
	)
}
func headersString(header http.Header) string {
	var result string
	for name, values := range header {
		result += name + ": " + strings.Join(values, ", ") + "\n"
	}
	return result
}
