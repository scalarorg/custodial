// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	github_com_scalarorg_scalar_core_x_vote_exported "github.com/scalarorg/scalar-core/x/vote/exported"
	"github.com/scalarorg/scalar-core/x/vote/types"
	"sync"
)

// Ensure, that VoteRouterMock does implement types.VoteRouter.
// If this is not the case, regenerate this file with moq.
var _ types.VoteRouter = &VoteRouterMock{}

// VoteRouterMock is a mock implementation of types.VoteRouter.
//
//	func TestSomethingThatUsesVoteRouter(t *testing.T) {
//
//		// make and configure a mocked types.VoteRouter
//		mockedVoteRouter := &VoteRouterMock{
//			AddHandlerFunc: func(module string, handler github_com_scalarorg_scalar_core_x_vote_exported.VoteHandler) types.VoteRouter {
//				panic("mock out the AddHandler method")
//			},
//			GetHandlerFunc: func(module string) github_com_scalarorg_scalar_core_x_vote_exported.VoteHandler {
//				panic("mock out the GetHandler method")
//			},
//			HasHandlerFunc: func(module string) bool {
//				panic("mock out the HasHandler method")
//			},
//			SealFunc: func()  {
//				panic("mock out the Seal method")
//			},
//		}
//
//		// use mockedVoteRouter in code that requires types.VoteRouter
//		// and then make assertions.
//
//	}
type VoteRouterMock struct {
	// AddHandlerFunc mocks the AddHandler method.
	AddHandlerFunc func(module string, handler github_com_scalarorg_scalar_core_x_vote_exported.VoteHandler) types.VoteRouter

	// GetHandlerFunc mocks the GetHandler method.
	GetHandlerFunc func(module string) github_com_scalarorg_scalar_core_x_vote_exported.VoteHandler

	// HasHandlerFunc mocks the HasHandler method.
	HasHandlerFunc func(module string) bool

	// SealFunc mocks the Seal method.
	SealFunc func()

	// calls tracks calls to the methods.
	calls struct {
		// AddHandler holds details about calls to the AddHandler method.
		AddHandler []struct {
			// Module is the module argument value.
			Module string
			// Handler is the handler argument value.
			Handler github_com_scalarorg_scalar_core_x_vote_exported.VoteHandler
		}
		// GetHandler holds details about calls to the GetHandler method.
		GetHandler []struct {
			// Module is the module argument value.
			Module string
		}
		// HasHandler holds details about calls to the HasHandler method.
		HasHandler []struct {
			// Module is the module argument value.
			Module string
		}
		// Seal holds details about calls to the Seal method.
		Seal []struct {
		}
	}
	lockAddHandler sync.RWMutex
	lockGetHandler sync.RWMutex
	lockHasHandler sync.RWMutex
	lockSeal       sync.RWMutex
}

// AddHandler calls AddHandlerFunc.
func (mock *VoteRouterMock) AddHandler(module string, handler github_com_scalarorg_scalar_core_x_vote_exported.VoteHandler) types.VoteRouter {
	if mock.AddHandlerFunc == nil {
		panic("VoteRouterMock.AddHandlerFunc: method is nil but VoteRouter.AddHandler was just called")
	}
	callInfo := struct {
		Module  string
		Handler github_com_scalarorg_scalar_core_x_vote_exported.VoteHandler
	}{
		Module:  module,
		Handler: handler,
	}
	mock.lockAddHandler.Lock()
	mock.calls.AddHandler = append(mock.calls.AddHandler, callInfo)
	mock.lockAddHandler.Unlock()
	return mock.AddHandlerFunc(module, handler)
}

// AddHandlerCalls gets all the calls that were made to AddHandler.
// Check the length with:
//
//	len(mockedVoteRouter.AddHandlerCalls())
func (mock *VoteRouterMock) AddHandlerCalls() []struct {
	Module  string
	Handler github_com_scalarorg_scalar_core_x_vote_exported.VoteHandler
} {
	var calls []struct {
		Module  string
		Handler github_com_scalarorg_scalar_core_x_vote_exported.VoteHandler
	}
	mock.lockAddHandler.RLock()
	calls = mock.calls.AddHandler
	mock.lockAddHandler.RUnlock()
	return calls
}

// GetHandler calls GetHandlerFunc.
func (mock *VoteRouterMock) GetHandler(module string) github_com_scalarorg_scalar_core_x_vote_exported.VoteHandler {
	if mock.GetHandlerFunc == nil {
		panic("VoteRouterMock.GetHandlerFunc: method is nil but VoteRouter.GetHandler was just called")
	}
	callInfo := struct {
		Module string
	}{
		Module: module,
	}
	mock.lockGetHandler.Lock()
	mock.calls.GetHandler = append(mock.calls.GetHandler, callInfo)
	mock.lockGetHandler.Unlock()
	return mock.GetHandlerFunc(module)
}

// GetHandlerCalls gets all the calls that were made to GetHandler.
// Check the length with:
//
//	len(mockedVoteRouter.GetHandlerCalls())
func (mock *VoteRouterMock) GetHandlerCalls() []struct {
	Module string
} {
	var calls []struct {
		Module string
	}
	mock.lockGetHandler.RLock()
	calls = mock.calls.GetHandler
	mock.lockGetHandler.RUnlock()
	return calls
}

// HasHandler calls HasHandlerFunc.
func (mock *VoteRouterMock) HasHandler(module string) bool {
	if mock.HasHandlerFunc == nil {
		panic("VoteRouterMock.HasHandlerFunc: method is nil but VoteRouter.HasHandler was just called")
	}
	callInfo := struct {
		Module string
	}{
		Module: module,
	}
	mock.lockHasHandler.Lock()
	mock.calls.HasHandler = append(mock.calls.HasHandler, callInfo)
	mock.lockHasHandler.Unlock()
	return mock.HasHandlerFunc(module)
}

// HasHandlerCalls gets all the calls that were made to HasHandler.
// Check the length with:
//
//	len(mockedVoteRouter.HasHandlerCalls())
func (mock *VoteRouterMock) HasHandlerCalls() []struct {
	Module string
} {
	var calls []struct {
		Module string
	}
	mock.lockHasHandler.RLock()
	calls = mock.calls.HasHandler
	mock.lockHasHandler.RUnlock()
	return calls
}

// Seal calls SealFunc.
func (mock *VoteRouterMock) Seal() {
	if mock.SealFunc == nil {
		panic("VoteRouterMock.SealFunc: method is nil but VoteRouter.Seal was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSeal.Lock()
	mock.calls.Seal = append(mock.calls.Seal, callInfo)
	mock.lockSeal.Unlock()
	mock.SealFunc()
}

// SealCalls gets all the calls that were made to Seal.
// Check the length with:
//
//	len(mockedVoteRouter.SealCalls())
func (mock *VoteRouterMock) SealCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSeal.RLock()
	calls = mock.calls.Seal
	mock.lockSeal.RUnlock()
	return calls
}