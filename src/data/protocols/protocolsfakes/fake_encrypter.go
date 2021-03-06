// Code generated by counterfeiter. DO NOT EDIT.
package protocolsfakes

import (
	"lucaswilliameufrasio/golang-fiber-api/src/data/protocols"
	"sync"
)

type FakeEncrypter struct {
	EncryptStub        func(string) (string, error)
	encryptMutex       sync.RWMutex
	encryptArgsForCall []struct {
		arg1 string
	}
	encryptReturns struct {
		result1 string
		result2 error
	}
	encryptReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEncrypter) Encrypt(arg1 string) (string, error) {
	fake.encryptMutex.Lock()
	ret, specificReturn := fake.encryptReturnsOnCall[len(fake.encryptArgsForCall)]
	fake.encryptArgsForCall = append(fake.encryptArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.EncryptStub
	fakeReturns := fake.encryptReturns
	fake.recordInvocation("Encrypt", []interface{}{arg1})
	fake.encryptMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEncrypter) EncryptCallCount() int {
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	return len(fake.encryptArgsForCall)
}

func (fake *FakeEncrypter) EncryptCalls(stub func(string) (string, error)) {
	fake.encryptMutex.Lock()
	defer fake.encryptMutex.Unlock()
	fake.EncryptStub = stub
}

func (fake *FakeEncrypter) EncryptArgsForCall(i int) string {
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	argsForCall := fake.encryptArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEncrypter) EncryptReturns(result1 string, result2 error) {
	fake.encryptMutex.Lock()
	defer fake.encryptMutex.Unlock()
	fake.EncryptStub = nil
	fake.encryptReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeEncrypter) EncryptReturnsOnCall(i int, result1 string, result2 error) {
	fake.encryptMutex.Lock()
	defer fake.encryptMutex.Unlock()
	fake.EncryptStub = nil
	if fake.encryptReturnsOnCall == nil {
		fake.encryptReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.encryptReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeEncrypter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEncrypter) recordInvocation(key string, args []interface{}) {
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

var _ protocols.Encrypter = new(FakeEncrypter)
