// Code generated by counterfeiter. DO NOT EDIT.
package protocolsfakes

import (
	"lucaswilliameufrasio/golang-fiber-api/src/data/protocols"
	"sync"
)

type FakeLoadUserByEmailRepository struct {
	LoadByEmailStub        func(string) (*protocols.LoadUserByIDRepositoryResult, error)
	loadByEmailMutex       sync.RWMutex
	loadByEmailArgsForCall []struct {
		arg1 string
	}
	loadByEmailReturns struct {
		result1 *protocols.LoadUserByIDRepositoryResult
		result2 error
	}
	loadByEmailReturnsOnCall map[int]struct {
		result1 *protocols.LoadUserByIDRepositoryResult
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLoadUserByEmailRepository) LoadByEmail(arg1 string) (*protocols.LoadUserByIDRepositoryResult, error) {
	fake.loadByEmailMutex.Lock()
	ret, specificReturn := fake.loadByEmailReturnsOnCall[len(fake.loadByEmailArgsForCall)]
	fake.loadByEmailArgsForCall = append(fake.loadByEmailArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LoadByEmailStub
	fakeReturns := fake.loadByEmailReturns
	fake.recordInvocation("LoadByEmail", []interface{}{arg1})
	fake.loadByEmailMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLoadUserByEmailRepository) LoadByEmailCallCount() int {
	fake.loadByEmailMutex.RLock()
	defer fake.loadByEmailMutex.RUnlock()
	return len(fake.loadByEmailArgsForCall)
}

func (fake *FakeLoadUserByEmailRepository) LoadByEmailCalls(stub func(string) (*protocols.LoadUserByIDRepositoryResult, error)) {
	fake.loadByEmailMutex.Lock()
	defer fake.loadByEmailMutex.Unlock()
	fake.LoadByEmailStub = stub
}

func (fake *FakeLoadUserByEmailRepository) LoadByEmailArgsForCall(i int) string {
	fake.loadByEmailMutex.RLock()
	defer fake.loadByEmailMutex.RUnlock()
	argsForCall := fake.loadByEmailArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLoadUserByEmailRepository) LoadByEmailReturns(result1 *protocols.LoadUserByIDRepositoryResult, result2 error) {
	fake.loadByEmailMutex.Lock()
	defer fake.loadByEmailMutex.Unlock()
	fake.LoadByEmailStub = nil
	fake.loadByEmailReturns = struct {
		result1 *protocols.LoadUserByIDRepositoryResult
		result2 error
	}{result1, result2}
}

func (fake *FakeLoadUserByEmailRepository) LoadByEmailReturnsOnCall(i int, result1 *protocols.LoadUserByIDRepositoryResult, result2 error) {
	fake.loadByEmailMutex.Lock()
	defer fake.loadByEmailMutex.Unlock()
	fake.LoadByEmailStub = nil
	if fake.loadByEmailReturnsOnCall == nil {
		fake.loadByEmailReturnsOnCall = make(map[int]struct {
			result1 *protocols.LoadUserByIDRepositoryResult
			result2 error
		})
	}
	fake.loadByEmailReturnsOnCall[i] = struct {
		result1 *protocols.LoadUserByIDRepositoryResult
		result2 error
	}{result1, result2}
}

func (fake *FakeLoadUserByEmailRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loadByEmailMutex.RLock()
	defer fake.loadByEmailMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLoadUserByEmailRepository) recordInvocation(key string, args []interface{}) {
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

var _ protocols.LoadUserByEmailRepository = new(FakeLoadUserByEmailRepository)
