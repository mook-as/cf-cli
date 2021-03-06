// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeCancelDeploymentActor struct {
	CancelDeploymentStub        func(string) (v7action.Warnings, error)
	cancelDeploymentMutex       sync.RWMutex
	cancelDeploymentArgsForCall []struct {
		arg1 string
	}
	cancelDeploymentReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	cancelDeploymentReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	GetApplicationByNameAndSpaceStub        func(string, string) (v7action.Application, v7action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	GetLatestActiveDeploymentForAppStub        func(string) (v7action.Deployment, v7action.Warnings, error)
	getLatestActiveDeploymentForAppMutex       sync.RWMutex
	getLatestActiveDeploymentForAppArgsForCall []struct {
		arg1 string
	}
	getLatestActiveDeploymentForAppReturns struct {
		result1 v7action.Deployment
		result2 v7action.Warnings
		result3 error
	}
	getLatestActiveDeploymentForAppReturnsOnCall map[int]struct {
		result1 v7action.Deployment
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCancelDeploymentActor) CancelDeployment(arg1 string) (v7action.Warnings, error) {
	fake.cancelDeploymentMutex.Lock()
	ret, specificReturn := fake.cancelDeploymentReturnsOnCall[len(fake.cancelDeploymentArgsForCall)]
	fake.cancelDeploymentArgsForCall = append(fake.cancelDeploymentArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("CancelDeployment", []interface{}{arg1})
	fake.cancelDeploymentMutex.Unlock()
	if fake.CancelDeploymentStub != nil {
		return fake.CancelDeploymentStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.cancelDeploymentReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCancelDeploymentActor) CancelDeploymentCallCount() int {
	fake.cancelDeploymentMutex.RLock()
	defer fake.cancelDeploymentMutex.RUnlock()
	return len(fake.cancelDeploymentArgsForCall)
}

func (fake *FakeCancelDeploymentActor) CancelDeploymentCalls(stub func(string) (v7action.Warnings, error)) {
	fake.cancelDeploymentMutex.Lock()
	defer fake.cancelDeploymentMutex.Unlock()
	fake.CancelDeploymentStub = stub
}

func (fake *FakeCancelDeploymentActor) CancelDeploymentArgsForCall(i int) string {
	fake.cancelDeploymentMutex.RLock()
	defer fake.cancelDeploymentMutex.RUnlock()
	argsForCall := fake.cancelDeploymentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCancelDeploymentActor) CancelDeploymentReturns(result1 v7action.Warnings, result2 error) {
	fake.cancelDeploymentMutex.Lock()
	defer fake.cancelDeploymentMutex.Unlock()
	fake.CancelDeploymentStub = nil
	fake.cancelDeploymentReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCancelDeploymentActor) CancelDeploymentReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.cancelDeploymentMutex.Lock()
	defer fake.cancelDeploymentMutex.Unlock()
	fake.CancelDeploymentStub = nil
	if fake.cancelDeploymentReturnsOnCall == nil {
		fake.cancelDeploymentReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.cancelDeploymentReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCancelDeploymentActor) GetApplicationByNameAndSpace(arg1 string, arg2 string) (v7action.Application, v7action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCancelDeploymentActor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeCancelDeploymentActor) GetApplicationByNameAndSpaceCalls(stub func(string, string) (v7action.Application, v7action.Warnings, error)) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = stub
}

func (fake *FakeCancelDeploymentActor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCancelDeploymentActor) GetApplicationByNameAndSpaceReturns(result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCancelDeploymentActor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.Application
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCancelDeploymentActor) GetLatestActiveDeploymentForApp(arg1 string) (v7action.Deployment, v7action.Warnings, error) {
	fake.getLatestActiveDeploymentForAppMutex.Lock()
	ret, specificReturn := fake.getLatestActiveDeploymentForAppReturnsOnCall[len(fake.getLatestActiveDeploymentForAppArgsForCall)]
	fake.getLatestActiveDeploymentForAppArgsForCall = append(fake.getLatestActiveDeploymentForAppArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetLatestActiveDeploymentForApp", []interface{}{arg1})
	fake.getLatestActiveDeploymentForAppMutex.Unlock()
	if fake.GetLatestActiveDeploymentForAppStub != nil {
		return fake.GetLatestActiveDeploymentForAppStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getLatestActiveDeploymentForAppReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCancelDeploymentActor) GetLatestActiveDeploymentForAppCallCount() int {
	fake.getLatestActiveDeploymentForAppMutex.RLock()
	defer fake.getLatestActiveDeploymentForAppMutex.RUnlock()
	return len(fake.getLatestActiveDeploymentForAppArgsForCall)
}

func (fake *FakeCancelDeploymentActor) GetLatestActiveDeploymentForAppCalls(stub func(string) (v7action.Deployment, v7action.Warnings, error)) {
	fake.getLatestActiveDeploymentForAppMutex.Lock()
	defer fake.getLatestActiveDeploymentForAppMutex.Unlock()
	fake.GetLatestActiveDeploymentForAppStub = stub
}

func (fake *FakeCancelDeploymentActor) GetLatestActiveDeploymentForAppArgsForCall(i int) string {
	fake.getLatestActiveDeploymentForAppMutex.RLock()
	defer fake.getLatestActiveDeploymentForAppMutex.RUnlock()
	argsForCall := fake.getLatestActiveDeploymentForAppArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCancelDeploymentActor) GetLatestActiveDeploymentForAppReturns(result1 v7action.Deployment, result2 v7action.Warnings, result3 error) {
	fake.getLatestActiveDeploymentForAppMutex.Lock()
	defer fake.getLatestActiveDeploymentForAppMutex.Unlock()
	fake.GetLatestActiveDeploymentForAppStub = nil
	fake.getLatestActiveDeploymentForAppReturns = struct {
		result1 v7action.Deployment
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCancelDeploymentActor) GetLatestActiveDeploymentForAppReturnsOnCall(i int, result1 v7action.Deployment, result2 v7action.Warnings, result3 error) {
	fake.getLatestActiveDeploymentForAppMutex.Lock()
	defer fake.getLatestActiveDeploymentForAppMutex.Unlock()
	fake.GetLatestActiveDeploymentForAppStub = nil
	if fake.getLatestActiveDeploymentForAppReturnsOnCall == nil {
		fake.getLatestActiveDeploymentForAppReturnsOnCall = make(map[int]struct {
			result1 v7action.Deployment
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getLatestActiveDeploymentForAppReturnsOnCall[i] = struct {
		result1 v7action.Deployment
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCancelDeploymentActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cancelDeploymentMutex.RLock()
	defer fake.cancelDeploymentMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getLatestActiveDeploymentForAppMutex.RLock()
	defer fake.getLatestActiveDeploymentForAppMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCancelDeploymentActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v7.CancelDeploymentActor = new(FakeCancelDeploymentActor)
