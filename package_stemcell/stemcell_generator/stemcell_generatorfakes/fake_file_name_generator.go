// Code generated by counterfeiter. DO NOT EDIT.
package stemcell_generatorfakes

import (
	sync "sync"

	stemcell_generator "github.com/cloudfoundry-incubator/stembuild/package_stemcell/stemcell_generator"
)

type FakeFileNameGenerator struct {
	FileNameStub        func() string
	fileNameMutex       sync.RWMutex
	fileNameArgsForCall []struct {
	}
	fileNameReturns struct {
		result1 string
	}
	fileNameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFileNameGenerator) FileName() string {
	fake.fileNameMutex.Lock()
	ret, specificReturn := fake.fileNameReturnsOnCall[len(fake.fileNameArgsForCall)]
	fake.fileNameArgsForCall = append(fake.fileNameArgsForCall, struct {
	}{})
	fake.recordInvocation("FileName", []interface{}{})
	fake.fileNameMutex.Unlock()
	if fake.FileNameStub != nil {
		return fake.FileNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.fileNameReturns
	return fakeReturns.result1
}

func (fake *FakeFileNameGenerator) FileNameCallCount() int {
	fake.fileNameMutex.RLock()
	defer fake.fileNameMutex.RUnlock()
	return len(fake.fileNameArgsForCall)
}

func (fake *FakeFileNameGenerator) FileNameCalls(stub func() string) {
	fake.fileNameMutex.Lock()
	defer fake.fileNameMutex.Unlock()
	fake.FileNameStub = stub
}

func (fake *FakeFileNameGenerator) FileNameReturns(result1 string) {
	fake.fileNameMutex.Lock()
	defer fake.fileNameMutex.Unlock()
	fake.FileNameStub = nil
	fake.fileNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFileNameGenerator) FileNameReturnsOnCall(i int, result1 string) {
	fake.fileNameMutex.Lock()
	defer fake.fileNameMutex.Unlock()
	fake.FileNameStub = nil
	if fake.fileNameReturnsOnCall == nil {
		fake.fileNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.fileNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeFileNameGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fileNameMutex.RLock()
	defer fake.fileNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFileNameGenerator) recordInvocation(key string, args []interface{}) {
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

var _ stemcell_generator.FileNameGenerator = new(FakeFileNameGenerator)
