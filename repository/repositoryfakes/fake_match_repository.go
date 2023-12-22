// Code generated by counterfeiter. DO NOT EDIT.
package repositoryfakes

import (
	"context"
	"sync"

	"github.com/SyaibanAhmadRamadhan/go-export-import/model"
	"github.com/SyaibanAhmadRamadhan/go-export-import/repository"
)

type FakeMatchRepository struct {
	InsertManyStub        func(context.Context, []model.Match) error
	insertManyMutex       sync.RWMutex
	insertManyArgsForCall []struct {
		arg1 context.Context
		arg2 []model.Match
	}
	insertManyReturns struct {
		result1 error
	}
	insertManyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMatchRepository) InsertMany(arg1 context.Context, arg2 []model.Match) error {
	var arg2Copy []model.Match
	if arg2 != nil {
		arg2Copy = make([]model.Match, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.insertManyMutex.Lock()
	ret, specificReturn := fake.insertManyReturnsOnCall[len(fake.insertManyArgsForCall)]
	fake.insertManyArgsForCall = append(fake.insertManyArgsForCall, struct {
		arg1 context.Context
		arg2 []model.Match
	}{arg1, arg2Copy})
	stub := fake.InsertManyStub
	fakeReturns := fake.insertManyReturns
	fake.recordInvocation("InsertMany", []interface{}{arg1, arg2Copy})
	fake.insertManyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMatchRepository) InsertManyCallCount() int {
	fake.insertManyMutex.RLock()
	defer fake.insertManyMutex.RUnlock()
	return len(fake.insertManyArgsForCall)
}

func (fake *FakeMatchRepository) InsertManyCalls(stub func(context.Context, []model.Match) error) {
	fake.insertManyMutex.Lock()
	defer fake.insertManyMutex.Unlock()
	fake.InsertManyStub = stub
}

func (fake *FakeMatchRepository) InsertManyArgsForCall(i int) (context.Context, []model.Match) {
	fake.insertManyMutex.RLock()
	defer fake.insertManyMutex.RUnlock()
	argsForCall := fake.insertManyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMatchRepository) InsertManyReturns(result1 error) {
	fake.insertManyMutex.Lock()
	defer fake.insertManyMutex.Unlock()
	fake.InsertManyStub = nil
	fake.insertManyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMatchRepository) InsertManyReturnsOnCall(i int, result1 error) {
	fake.insertManyMutex.Lock()
	defer fake.insertManyMutex.Unlock()
	fake.InsertManyStub = nil
	if fake.insertManyReturnsOnCall == nil {
		fake.insertManyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.insertManyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMatchRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.insertManyMutex.RLock()
	defer fake.insertManyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMatchRepository) recordInvocation(key string, args []interface{}) {
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

var _ repository.MatchRepository = new(FakeMatchRepository)
