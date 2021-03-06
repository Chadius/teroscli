// Code generated by counterfeiter. DO NOT EDIT.
package terosgameserverfakes

import (
	"context"
	"sync"

	"github.com/chadius/terosgameserver/rpc/github.com/chadius/teros_game_server"
)

type FakeTerosGameServer struct {
	ReplayBattleScriptStub        func(context.Context, *teros_game_server.DataStreams) (*teros_game_server.Results, error)
	replayBattleScriptMutex       sync.RWMutex
	replayBattleScriptArgsForCall []struct {
		arg1 context.Context
		arg2 *teros_game_server.DataStreams
	}
	replayBattleScriptReturns struct {
		result1 *teros_game_server.Results
		result2 error
	}
	replayBattleScriptReturnsOnCall map[int]struct {
		result1 *teros_game_server.Results
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTerosGameServer) ReplayBattleScript(arg1 context.Context, arg2 *teros_game_server.DataStreams) (*teros_game_server.Results, error) {
	fake.replayBattleScriptMutex.Lock()
	ret, specificReturn := fake.replayBattleScriptReturnsOnCall[len(fake.replayBattleScriptArgsForCall)]
	fake.replayBattleScriptArgsForCall = append(fake.replayBattleScriptArgsForCall, struct {
		arg1 context.Context
		arg2 *teros_game_server.DataStreams
	}{arg1, arg2})
	stub := fake.ReplayBattleScriptStub
	fakeReturns := fake.replayBattleScriptReturns
	fake.recordInvocation("ReplayBattleScript", []interface{}{arg1, arg2})
	fake.replayBattleScriptMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTerosGameServer) ReplayBattleScriptCallCount() int {
	fake.replayBattleScriptMutex.RLock()
	defer fake.replayBattleScriptMutex.RUnlock()
	return len(fake.replayBattleScriptArgsForCall)
}

func (fake *FakeTerosGameServer) ReplayBattleScriptCalls(stub func(context.Context, *teros_game_server.DataStreams) (*teros_game_server.Results, error)) {
	fake.replayBattleScriptMutex.Lock()
	defer fake.replayBattleScriptMutex.Unlock()
	fake.ReplayBattleScriptStub = stub
}

func (fake *FakeTerosGameServer) ReplayBattleScriptArgsForCall(i int) (context.Context, *teros_game_server.DataStreams) {
	fake.replayBattleScriptMutex.RLock()
	defer fake.replayBattleScriptMutex.RUnlock()
	argsForCall := fake.replayBattleScriptArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTerosGameServer) ReplayBattleScriptReturns(result1 *teros_game_server.Results, result2 error) {
	fake.replayBattleScriptMutex.Lock()
	defer fake.replayBattleScriptMutex.Unlock()
	fake.ReplayBattleScriptStub = nil
	fake.replayBattleScriptReturns = struct {
		result1 *teros_game_server.Results
		result2 error
	}{result1, result2}
}

func (fake *FakeTerosGameServer) ReplayBattleScriptReturnsOnCall(i int, result1 *teros_game_server.Results, result2 error) {
	fake.replayBattleScriptMutex.Lock()
	defer fake.replayBattleScriptMutex.Unlock()
	fake.ReplayBattleScriptStub = nil
	if fake.replayBattleScriptReturnsOnCall == nil {
		fake.replayBattleScriptReturnsOnCall = make(map[int]struct {
			result1 *teros_game_server.Results
			result2 error
		})
	}
	fake.replayBattleScriptReturnsOnCall[i] = struct {
		result1 *teros_game_server.Results
		result2 error
	}{result1, result2}
}

func (fake *FakeTerosGameServer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.replayBattleScriptMutex.RLock()
	defer fake.replayBattleScriptMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTerosGameServer) recordInvocation(key string, args []interface{}) {
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

var _ teros_game_server.TerosGameServer = new(FakeTerosGameServer)
