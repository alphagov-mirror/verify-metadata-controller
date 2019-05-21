// Code generated by counterfeiter. DO NOT EDIT.
package hsmfakes

import (
	"sync"

	"github.com/alphagov/verify-metadata-controller/pkg/apis/verify/v1beta1"
	"github.com/alphagov/verify-metadata-controller/pkg/hsm"
)

type FakeClient struct {
	CreateRSAKeyPairStub        func(string, hsm.Credentials) ([]byte, error)
	createRSAKeyPairMutex       sync.RWMutex
	createRSAKeyPairArgsForCall []struct {
		arg1 string
		arg2 hsm.Credentials
	}
	createRSAKeyPairReturns struct {
		result1 []byte
		result2 error
	}
	createRSAKeyPairReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	FindOrCreateRSAKeyPairStub        func(string, hsm.Credentials) ([]byte, error)
	findOrCreateRSAKeyPairMutex       sync.RWMutex
	findOrCreateRSAKeyPairArgsForCall []struct {
		arg1 string
		arg2 hsm.Credentials
	}
	findOrCreateRSAKeyPairReturns struct {
		result1 []byte
		result2 error
	}
	findOrCreateRSAKeyPairReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GenerateAndSignMetadataStub        func([]byte, string, v1beta1.MetadataSpec, hsm.Credentials) ([]byte, error)
	generateAndSignMetadataMutex       sync.RWMutex
	generateAndSignMetadataArgsForCall []struct {
		arg1 []byte
		arg2 string
		arg3 v1beta1.MetadataSpec
		arg4 hsm.Credentials
	}
	generateAndSignMetadataReturns struct {
		result1 []byte
		result2 error
	}
	generateAndSignMetadataReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) CreateRSAKeyPair(arg1 string, arg2 hsm.Credentials) ([]byte, error) {
	fake.createRSAKeyPairMutex.Lock()
	ret, specificReturn := fake.createRSAKeyPairReturnsOnCall[len(fake.createRSAKeyPairArgsForCall)]
	fake.createRSAKeyPairArgsForCall = append(fake.createRSAKeyPairArgsForCall, struct {
		arg1 string
		arg2 hsm.Credentials
	}{arg1, arg2})
	fake.recordInvocation("CreateRSAKeyPair", []interface{}{arg1, arg2})
	fake.createRSAKeyPairMutex.Unlock()
	if fake.CreateRSAKeyPairStub != nil {
		return fake.CreateRSAKeyPairStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createRSAKeyPairReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) CreateRSAKeyPairCallCount() int {
	fake.createRSAKeyPairMutex.RLock()
	defer fake.createRSAKeyPairMutex.RUnlock()
	return len(fake.createRSAKeyPairArgsForCall)
}

func (fake *FakeClient) CreateRSAKeyPairCalls(stub func(string, hsm.Credentials) ([]byte, error)) {
	fake.createRSAKeyPairMutex.Lock()
	defer fake.createRSAKeyPairMutex.Unlock()
	fake.CreateRSAKeyPairStub = stub
}

func (fake *FakeClient) CreateRSAKeyPairArgsForCall(i int) (string, hsm.Credentials) {
	fake.createRSAKeyPairMutex.RLock()
	defer fake.createRSAKeyPairMutex.RUnlock()
	argsForCall := fake.createRSAKeyPairArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) CreateRSAKeyPairReturns(result1 []byte, result2 error) {
	fake.createRSAKeyPairMutex.Lock()
	defer fake.createRSAKeyPairMutex.Unlock()
	fake.CreateRSAKeyPairStub = nil
	fake.createRSAKeyPairReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateRSAKeyPairReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.createRSAKeyPairMutex.Lock()
	defer fake.createRSAKeyPairMutex.Unlock()
	fake.CreateRSAKeyPairStub = nil
	if fake.createRSAKeyPairReturnsOnCall == nil {
		fake.createRSAKeyPairReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.createRSAKeyPairReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindOrCreateRSAKeyPair(arg1 string, arg2 hsm.Credentials) ([]byte, error) {
	fake.findOrCreateRSAKeyPairMutex.Lock()
	ret, specificReturn := fake.findOrCreateRSAKeyPairReturnsOnCall[len(fake.findOrCreateRSAKeyPairArgsForCall)]
	fake.findOrCreateRSAKeyPairArgsForCall = append(fake.findOrCreateRSAKeyPairArgsForCall, struct {
		arg1 string
		arg2 hsm.Credentials
	}{arg1, arg2})
	fake.recordInvocation("FindOrCreateRSAKeyPair", []interface{}{arg1, arg2})
	fake.findOrCreateRSAKeyPairMutex.Unlock()
	if fake.FindOrCreateRSAKeyPairStub != nil {
		return fake.FindOrCreateRSAKeyPairStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findOrCreateRSAKeyPairReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) FindOrCreateRSAKeyPairCallCount() int {
	fake.findOrCreateRSAKeyPairMutex.RLock()
	defer fake.findOrCreateRSAKeyPairMutex.RUnlock()
	return len(fake.findOrCreateRSAKeyPairArgsForCall)
}

func (fake *FakeClient) FindOrCreateRSAKeyPairCalls(stub func(string, hsm.Credentials) ([]byte, error)) {
	fake.findOrCreateRSAKeyPairMutex.Lock()
	defer fake.findOrCreateRSAKeyPairMutex.Unlock()
	fake.FindOrCreateRSAKeyPairStub = stub
}

func (fake *FakeClient) FindOrCreateRSAKeyPairArgsForCall(i int) (string, hsm.Credentials) {
	fake.findOrCreateRSAKeyPairMutex.RLock()
	defer fake.findOrCreateRSAKeyPairMutex.RUnlock()
	argsForCall := fake.findOrCreateRSAKeyPairArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) FindOrCreateRSAKeyPairReturns(result1 []byte, result2 error) {
	fake.findOrCreateRSAKeyPairMutex.Lock()
	defer fake.findOrCreateRSAKeyPairMutex.Unlock()
	fake.FindOrCreateRSAKeyPairStub = nil
	fake.findOrCreateRSAKeyPairReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindOrCreateRSAKeyPairReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.findOrCreateRSAKeyPairMutex.Lock()
	defer fake.findOrCreateRSAKeyPairMutex.Unlock()
	fake.FindOrCreateRSAKeyPairStub = nil
	if fake.findOrCreateRSAKeyPairReturnsOnCall == nil {
		fake.findOrCreateRSAKeyPairReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.findOrCreateRSAKeyPairReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GenerateAndSignMetadata(arg1 []byte, arg2 string, arg3 v1beta1.MetadataSpec, arg4 hsm.Credentials) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.generateAndSignMetadataMutex.Lock()
	ret, specificReturn := fake.generateAndSignMetadataReturnsOnCall[len(fake.generateAndSignMetadataArgsForCall)]
	fake.generateAndSignMetadataArgsForCall = append(fake.generateAndSignMetadataArgsForCall, struct {
		arg1 []byte
		arg2 string
		arg3 v1beta1.MetadataSpec
		arg4 hsm.Credentials
	}{arg1Copy, arg2, arg3, arg4})
	fake.recordInvocation("GenerateAndSignMetadata", []interface{}{arg1Copy, arg2, arg3, arg4})
	fake.generateAndSignMetadataMutex.Unlock()
	if fake.GenerateAndSignMetadataStub != nil {
		return fake.GenerateAndSignMetadataStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.generateAndSignMetadataReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) GenerateAndSignMetadataCallCount() int {
	fake.generateAndSignMetadataMutex.RLock()
	defer fake.generateAndSignMetadataMutex.RUnlock()
	return len(fake.generateAndSignMetadataArgsForCall)
}

func (fake *FakeClient) GenerateAndSignMetadataCalls(stub func([]byte, string, v1beta1.MetadataSpec, hsm.Credentials) ([]byte, error)) {
	fake.generateAndSignMetadataMutex.Lock()
	defer fake.generateAndSignMetadataMutex.Unlock()
	fake.GenerateAndSignMetadataStub = stub
}

func (fake *FakeClient) GenerateAndSignMetadataArgsForCall(i int) ([]byte, string, v1beta1.MetadataSpec, hsm.Credentials) {
	fake.generateAndSignMetadataMutex.RLock()
	defer fake.generateAndSignMetadataMutex.RUnlock()
	argsForCall := fake.generateAndSignMetadataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClient) GenerateAndSignMetadataReturns(result1 []byte, result2 error) {
	fake.generateAndSignMetadataMutex.Lock()
	defer fake.generateAndSignMetadataMutex.Unlock()
	fake.GenerateAndSignMetadataStub = nil
	fake.generateAndSignMetadataReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GenerateAndSignMetadataReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.generateAndSignMetadataMutex.Lock()
	defer fake.generateAndSignMetadataMutex.Unlock()
	fake.GenerateAndSignMetadataStub = nil
	if fake.generateAndSignMetadataReturnsOnCall == nil {
		fake.generateAndSignMetadataReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.generateAndSignMetadataReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createRSAKeyPairMutex.RLock()
	defer fake.createRSAKeyPairMutex.RUnlock()
	fake.findOrCreateRSAKeyPairMutex.RLock()
	defer fake.findOrCreateRSAKeyPairMutex.RUnlock()
	fake.generateAndSignMetadataMutex.RLock()
	defer fake.generateAndSignMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ hsm.Client = new(FakeClient)
