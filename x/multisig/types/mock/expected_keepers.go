// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	utils "github.com/axelarnetwork/axelar-core/utils"
	github_com_axelarnetwork_axelar_core_x_multisig_exported "github.com/scalarorg/scalar-core/x/multisig/exported"
	"github.com/scalarorg/scalar-core/x/multisig/types"
	github_com_axelarnetwork_axelar_core_x_nexus_exported "github.com/scalarorg/scalar-core/x/nexus/exported"
	reward "github.com/scalarorg/scalar-core/x/reward/exported"
	exported "github.com/axelarnetwork/axelar-core/x/snapshot/exported"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/libs/log"
	"sync"
)

// Ensure, that KeeperMock does implement types.Keeper.
// If this is not the case, regenerate this file with moq.
var _ types.Keeper = &KeeperMock{}

// KeeperMock is a mock implementation of types.Keeper.
//
//	func TestSomethingThatUsesKeeper(t *testing.T) {
//
//		// make and configure a mocked types.Keeper
//		mockedKeeper := &KeeperMock{
//			DeleteKeygenSessionFunc: func(ctx sdk.Context, id github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID)  {
//				panic("mock out the DeleteKeygenSession method")
//			},
//			DeleteSigningSessionFunc: func(ctx sdk.Context, id uint64)  {
//				panic("mock out the DeleteSigningSession method")
//			},
//			GetCurrentKeyIDFunc: func(ctx sdk.Context, chainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName) (github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID, bool) {
//				panic("mock out the GetCurrentKeyID method")
//			},
//			GetKeyFunc: func(ctx sdk.Context, keyID github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID) (github_com_axelarnetwork_axelar_core_x_multisig_exported.Key, bool) {
//				panic("mock out the GetKey method")
//			},
//			GetKeygenSessionFunc: func(ctx sdk.Context, id github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID) (types.KeygenSession, bool) {
//				panic("mock out the GetKeygenSession method")
//			},
//			GetKeygenSessionsByExpiryFunc: func(ctx sdk.Context, expiry int64) []types.KeygenSession {
//				panic("mock out the GetKeygenSessionsByExpiry method")
//			},
//			GetNextKeyIDFunc: func(ctx sdk.Context, chainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName) (github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID, bool) {
//				panic("mock out the GetNextKeyID method")
//			},
//			GetParamsFunc: func(ctx sdk.Context) types.Params {
//				panic("mock out the GetParams method")
//			},
//			GetSigRouterFunc: func() types.SigRouter {
//				panic("mock out the GetSigRouter method")
//			},
//			GetSigningSessionsByExpiryFunc: func(ctx sdk.Context, expiry int64) []types.SigningSession {
//				panic("mock out the GetSigningSessionsByExpiry method")
//			},
//			LoggerFunc: func(ctx sdk.Context) log.Logger {
//				panic("mock out the Logger method")
//			},
//			SetKeyFunc: func(ctx sdk.Context, key types.Key)  {
//				panic("mock out the SetKey method")
//			},
//		}
//
//		// use mockedKeeper in code that requires types.Keeper
//		// and then make assertions.
//
//	}
type KeeperMock struct {
	// DeleteKeygenSessionFunc mocks the DeleteKeygenSession method.
	DeleteKeygenSessionFunc func(ctx sdk.Context, id github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID)

	// DeleteSigningSessionFunc mocks the DeleteSigningSession method.
	DeleteSigningSessionFunc func(ctx sdk.Context, id uint64)

	// GetCurrentKeyIDFunc mocks the GetCurrentKeyID method.
	GetCurrentKeyIDFunc func(ctx sdk.Context, chainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName) (github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID, bool)

	// GetKeyFunc mocks the GetKey method.
	GetKeyFunc func(ctx sdk.Context, keyID github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID) (github_com_axelarnetwork_axelar_core_x_multisig_exported.Key, bool)

	// GetKeygenSessionFunc mocks the GetKeygenSession method.
	GetKeygenSessionFunc func(ctx sdk.Context, id github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID) (types.KeygenSession, bool)

	// GetKeygenSessionsByExpiryFunc mocks the GetKeygenSessionsByExpiry method.
	GetKeygenSessionsByExpiryFunc func(ctx sdk.Context, expiry int64) []types.KeygenSession

	// GetNextKeyIDFunc mocks the GetNextKeyID method.
	GetNextKeyIDFunc func(ctx sdk.Context, chainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName) (github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID, bool)

	// GetParamsFunc mocks the GetParams method.
	GetParamsFunc func(ctx sdk.Context) types.Params

	// GetSigRouterFunc mocks the GetSigRouter method.
	GetSigRouterFunc func() types.SigRouter

	// GetSigningSessionsByExpiryFunc mocks the GetSigningSessionsByExpiry method.
	GetSigningSessionsByExpiryFunc func(ctx sdk.Context, expiry int64) []types.SigningSession

	// LoggerFunc mocks the Logger method.
	LoggerFunc func(ctx sdk.Context) log.Logger

	// SetKeyFunc mocks the SetKey method.
	SetKeyFunc func(ctx sdk.Context, key types.Key)

	// calls tracks calls to the methods.
	calls struct {
		// DeleteKeygenSession holds details about calls to the DeleteKeygenSession method.
		DeleteKeygenSession []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// ID is the id argument value.
			ID github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID
		}
		// DeleteSigningSession holds details about calls to the DeleteSigningSession method.
		DeleteSigningSession []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// ID is the id argument value.
			ID uint64
		}
		// GetCurrentKeyID holds details about calls to the GetCurrentKeyID method.
		GetCurrentKeyID []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// ChainName is the chainName argument value.
			ChainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
		}
		// GetKey holds details about calls to the GetKey method.
		GetKey []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// KeyID is the keyID argument value.
			KeyID github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID
		}
		// GetKeygenSession holds details about calls to the GetKeygenSession method.
		GetKeygenSession []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// ID is the id argument value.
			ID github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID
		}
		// GetKeygenSessionsByExpiry holds details about calls to the GetKeygenSessionsByExpiry method.
		GetKeygenSessionsByExpiry []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Expiry is the expiry argument value.
			Expiry int64
		}
		// GetNextKeyID holds details about calls to the GetNextKeyID method.
		GetNextKeyID []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// ChainName is the chainName argument value.
			ChainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
		}
		// GetParams holds details about calls to the GetParams method.
		GetParams []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
		}
		// GetSigRouter holds details about calls to the GetSigRouter method.
		GetSigRouter []struct {
		}
		// GetSigningSessionsByExpiry holds details about calls to the GetSigningSessionsByExpiry method.
		GetSigningSessionsByExpiry []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Expiry is the expiry argument value.
			Expiry int64
		}
		// Logger holds details about calls to the Logger method.
		Logger []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
		}
		// SetKey holds details about calls to the SetKey method.
		SetKey []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Key is the key argument value.
			Key types.Key
		}
	}
	lockDeleteKeygenSession        sync.RWMutex
	lockDeleteSigningSession       sync.RWMutex
	lockGetCurrentKeyID            sync.RWMutex
	lockGetKey                     sync.RWMutex
	lockGetKeygenSession           sync.RWMutex
	lockGetKeygenSessionsByExpiry  sync.RWMutex
	lockGetNextKeyID               sync.RWMutex
	lockGetParams                  sync.RWMutex
	lockGetSigRouter               sync.RWMutex
	lockGetSigningSessionsByExpiry sync.RWMutex
	lockLogger                     sync.RWMutex
	lockSetKey                     sync.RWMutex
}

// DeleteKeygenSession calls DeleteKeygenSessionFunc.
func (mock *KeeperMock) DeleteKeygenSession(ctx sdk.Context, id github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID) {
	if mock.DeleteKeygenSessionFunc == nil {
		panic("KeeperMock.DeleteKeygenSessionFunc: method is nil but Keeper.DeleteKeygenSession was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
		ID  github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteKeygenSession.Lock()
	mock.calls.DeleteKeygenSession = append(mock.calls.DeleteKeygenSession, callInfo)
	mock.lockDeleteKeygenSession.Unlock()
	mock.DeleteKeygenSessionFunc(ctx, id)
}

// DeleteKeygenSessionCalls gets all the calls that were made to DeleteKeygenSession.
// Check the length with:
//
//	len(mockedKeeper.DeleteKeygenSessionCalls())
func (mock *KeeperMock) DeleteKeygenSessionCalls() []struct {
	Ctx sdk.Context
	ID  github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID
} {
	var calls []struct {
		Ctx sdk.Context
		ID  github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID
	}
	mock.lockDeleteKeygenSession.RLock()
	calls = mock.calls.DeleteKeygenSession
	mock.lockDeleteKeygenSession.RUnlock()
	return calls
}

// DeleteSigningSession calls DeleteSigningSessionFunc.
func (mock *KeeperMock) DeleteSigningSession(ctx sdk.Context, id uint64) {
	if mock.DeleteSigningSessionFunc == nil {
		panic("KeeperMock.DeleteSigningSessionFunc: method is nil but Keeper.DeleteSigningSession was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
		ID  uint64
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteSigningSession.Lock()
	mock.calls.DeleteSigningSession = append(mock.calls.DeleteSigningSession, callInfo)
	mock.lockDeleteSigningSession.Unlock()
	mock.DeleteSigningSessionFunc(ctx, id)
}

// DeleteSigningSessionCalls gets all the calls that were made to DeleteSigningSession.
// Check the length with:
//
//	len(mockedKeeper.DeleteSigningSessionCalls())
func (mock *KeeperMock) DeleteSigningSessionCalls() []struct {
	Ctx sdk.Context
	ID  uint64
} {
	var calls []struct {
		Ctx sdk.Context
		ID  uint64
	}
	mock.lockDeleteSigningSession.RLock()
	calls = mock.calls.DeleteSigningSession
	mock.lockDeleteSigningSession.RUnlock()
	return calls
}

// GetCurrentKeyID calls GetCurrentKeyIDFunc.
func (mock *KeeperMock) GetCurrentKeyID(ctx sdk.Context, chainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName) (github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID, bool) {
	if mock.GetCurrentKeyIDFunc == nil {
		panic("KeeperMock.GetCurrentKeyIDFunc: method is nil but Keeper.GetCurrentKeyID was just called")
	}
	callInfo := struct {
		Ctx       sdk.Context
		ChainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
	}{
		Ctx:       ctx,
		ChainName: chainName,
	}
	mock.lockGetCurrentKeyID.Lock()
	mock.calls.GetCurrentKeyID = append(mock.calls.GetCurrentKeyID, callInfo)
	mock.lockGetCurrentKeyID.Unlock()
	return mock.GetCurrentKeyIDFunc(ctx, chainName)
}

// GetCurrentKeyIDCalls gets all the calls that were made to GetCurrentKeyID.
// Check the length with:
//
//	len(mockedKeeper.GetCurrentKeyIDCalls())
func (mock *KeeperMock) GetCurrentKeyIDCalls() []struct {
	Ctx       sdk.Context
	ChainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
} {
	var calls []struct {
		Ctx       sdk.Context
		ChainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
	}
	mock.lockGetCurrentKeyID.RLock()
	calls = mock.calls.GetCurrentKeyID
	mock.lockGetCurrentKeyID.RUnlock()
	return calls
}

// GetKey calls GetKeyFunc.
func (mock *KeeperMock) GetKey(ctx sdk.Context, keyID github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID) (github_com_axelarnetwork_axelar_core_x_multisig_exported.Key, bool) {
	if mock.GetKeyFunc == nil {
		panic("KeeperMock.GetKeyFunc: method is nil but Keeper.GetKey was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		KeyID github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID
	}{
		Ctx:   ctx,
		KeyID: keyID,
	}
	mock.lockGetKey.Lock()
	mock.calls.GetKey = append(mock.calls.GetKey, callInfo)
	mock.lockGetKey.Unlock()
	return mock.GetKeyFunc(ctx, keyID)
}

// GetKeyCalls gets all the calls that were made to GetKey.
// Check the length with:
//
//	len(mockedKeeper.GetKeyCalls())
func (mock *KeeperMock) GetKeyCalls() []struct {
	Ctx   sdk.Context
	KeyID github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID
} {
	var calls []struct {
		Ctx   sdk.Context
		KeyID github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID
	}
	mock.lockGetKey.RLock()
	calls = mock.calls.GetKey
	mock.lockGetKey.RUnlock()
	return calls
}

// GetKeygenSession calls GetKeygenSessionFunc.
func (mock *KeeperMock) GetKeygenSession(ctx sdk.Context, id github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID) (types.KeygenSession, bool) {
	if mock.GetKeygenSessionFunc == nil {
		panic("KeeperMock.GetKeygenSessionFunc: method is nil but Keeper.GetKeygenSession was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
		ID  github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetKeygenSession.Lock()
	mock.calls.GetKeygenSession = append(mock.calls.GetKeygenSession, callInfo)
	mock.lockGetKeygenSession.Unlock()
	return mock.GetKeygenSessionFunc(ctx, id)
}

// GetKeygenSessionCalls gets all the calls that were made to GetKeygenSession.
// Check the length with:
//
//	len(mockedKeeper.GetKeygenSessionCalls())
func (mock *KeeperMock) GetKeygenSessionCalls() []struct {
	Ctx sdk.Context
	ID  github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID
} {
	var calls []struct {
		Ctx sdk.Context
		ID  github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID
	}
	mock.lockGetKeygenSession.RLock()
	calls = mock.calls.GetKeygenSession
	mock.lockGetKeygenSession.RUnlock()
	return calls
}

// GetKeygenSessionsByExpiry calls GetKeygenSessionsByExpiryFunc.
func (mock *KeeperMock) GetKeygenSessionsByExpiry(ctx sdk.Context, expiry int64) []types.KeygenSession {
	if mock.GetKeygenSessionsByExpiryFunc == nil {
		panic("KeeperMock.GetKeygenSessionsByExpiryFunc: method is nil but Keeper.GetKeygenSessionsByExpiry was just called")
	}
	callInfo := struct {
		Ctx    sdk.Context
		Expiry int64
	}{
		Ctx:    ctx,
		Expiry: expiry,
	}
	mock.lockGetKeygenSessionsByExpiry.Lock()
	mock.calls.GetKeygenSessionsByExpiry = append(mock.calls.GetKeygenSessionsByExpiry, callInfo)
	mock.lockGetKeygenSessionsByExpiry.Unlock()
	return mock.GetKeygenSessionsByExpiryFunc(ctx, expiry)
}

// GetKeygenSessionsByExpiryCalls gets all the calls that were made to GetKeygenSessionsByExpiry.
// Check the length with:
//
//	len(mockedKeeper.GetKeygenSessionsByExpiryCalls())
func (mock *KeeperMock) GetKeygenSessionsByExpiryCalls() []struct {
	Ctx    sdk.Context
	Expiry int64
} {
	var calls []struct {
		Ctx    sdk.Context
		Expiry int64
	}
	mock.lockGetKeygenSessionsByExpiry.RLock()
	calls = mock.calls.GetKeygenSessionsByExpiry
	mock.lockGetKeygenSessionsByExpiry.RUnlock()
	return calls
}

// GetNextKeyID calls GetNextKeyIDFunc.
func (mock *KeeperMock) GetNextKeyID(ctx sdk.Context, chainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName) (github_com_axelarnetwork_axelar_core_x_multisig_exported.KeyID, bool) {
	if mock.GetNextKeyIDFunc == nil {
		panic("KeeperMock.GetNextKeyIDFunc: method is nil but Keeper.GetNextKeyID was just called")
	}
	callInfo := struct {
		Ctx       sdk.Context
		ChainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
	}{
		Ctx:       ctx,
		ChainName: chainName,
	}
	mock.lockGetNextKeyID.Lock()
	mock.calls.GetNextKeyID = append(mock.calls.GetNextKeyID, callInfo)
	mock.lockGetNextKeyID.Unlock()
	return mock.GetNextKeyIDFunc(ctx, chainName)
}

// GetNextKeyIDCalls gets all the calls that were made to GetNextKeyID.
// Check the length with:
//
//	len(mockedKeeper.GetNextKeyIDCalls())
func (mock *KeeperMock) GetNextKeyIDCalls() []struct {
	Ctx       sdk.Context
	ChainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
} {
	var calls []struct {
		Ctx       sdk.Context
		ChainName github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
	}
	mock.lockGetNextKeyID.RLock()
	calls = mock.calls.GetNextKeyID
	mock.lockGetNextKeyID.RUnlock()
	return calls
}

// GetParams calls GetParamsFunc.
func (mock *KeeperMock) GetParams(ctx sdk.Context) types.Params {
	if mock.GetParamsFunc == nil {
		panic("KeeperMock.GetParamsFunc: method is nil but Keeper.GetParams was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetParams.Lock()
	mock.calls.GetParams = append(mock.calls.GetParams, callInfo)
	mock.lockGetParams.Unlock()
	return mock.GetParamsFunc(ctx)
}

// GetParamsCalls gets all the calls that were made to GetParams.
// Check the length with:
//
//	len(mockedKeeper.GetParamsCalls())
func (mock *KeeperMock) GetParamsCalls() []struct {
	Ctx sdk.Context
} {
	var calls []struct {
		Ctx sdk.Context
	}
	mock.lockGetParams.RLock()
	calls = mock.calls.GetParams
	mock.lockGetParams.RUnlock()
	return calls
}

// GetSigRouter calls GetSigRouterFunc.
func (mock *KeeperMock) GetSigRouter() types.SigRouter {
	if mock.GetSigRouterFunc == nil {
		panic("KeeperMock.GetSigRouterFunc: method is nil but Keeper.GetSigRouter was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetSigRouter.Lock()
	mock.calls.GetSigRouter = append(mock.calls.GetSigRouter, callInfo)
	mock.lockGetSigRouter.Unlock()
	return mock.GetSigRouterFunc()
}

// GetSigRouterCalls gets all the calls that were made to GetSigRouter.
// Check the length with:
//
//	len(mockedKeeper.GetSigRouterCalls())
func (mock *KeeperMock) GetSigRouterCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetSigRouter.RLock()
	calls = mock.calls.GetSigRouter
	mock.lockGetSigRouter.RUnlock()
	return calls
}

// GetSigningSessionsByExpiry calls GetSigningSessionsByExpiryFunc.
func (mock *KeeperMock) GetSigningSessionsByExpiry(ctx sdk.Context, expiry int64) []types.SigningSession {
	if mock.GetSigningSessionsByExpiryFunc == nil {
		panic("KeeperMock.GetSigningSessionsByExpiryFunc: method is nil but Keeper.GetSigningSessionsByExpiry was just called")
	}
	callInfo := struct {
		Ctx    sdk.Context
		Expiry int64
	}{
		Ctx:    ctx,
		Expiry: expiry,
	}
	mock.lockGetSigningSessionsByExpiry.Lock()
	mock.calls.GetSigningSessionsByExpiry = append(mock.calls.GetSigningSessionsByExpiry, callInfo)
	mock.lockGetSigningSessionsByExpiry.Unlock()
	return mock.GetSigningSessionsByExpiryFunc(ctx, expiry)
}

// GetSigningSessionsByExpiryCalls gets all the calls that were made to GetSigningSessionsByExpiry.
// Check the length with:
//
//	len(mockedKeeper.GetSigningSessionsByExpiryCalls())
func (mock *KeeperMock) GetSigningSessionsByExpiryCalls() []struct {
	Ctx    sdk.Context
	Expiry int64
} {
	var calls []struct {
		Ctx    sdk.Context
		Expiry int64
	}
	mock.lockGetSigningSessionsByExpiry.RLock()
	calls = mock.calls.GetSigningSessionsByExpiry
	mock.lockGetSigningSessionsByExpiry.RUnlock()
	return calls
}

// Logger calls LoggerFunc.
func (mock *KeeperMock) Logger(ctx sdk.Context) log.Logger {
	if mock.LoggerFunc == nil {
		panic("KeeperMock.LoggerFunc: method is nil but Keeper.Logger was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
	}{
		Ctx: ctx,
	}
	mock.lockLogger.Lock()
	mock.calls.Logger = append(mock.calls.Logger, callInfo)
	mock.lockLogger.Unlock()
	return mock.LoggerFunc(ctx)
}

// LoggerCalls gets all the calls that were made to Logger.
// Check the length with:
//
//	len(mockedKeeper.LoggerCalls())
func (mock *KeeperMock) LoggerCalls() []struct {
	Ctx sdk.Context
} {
	var calls []struct {
		Ctx sdk.Context
	}
	mock.lockLogger.RLock()
	calls = mock.calls.Logger
	mock.lockLogger.RUnlock()
	return calls
}

// SetKey calls SetKeyFunc.
func (mock *KeeperMock) SetKey(ctx sdk.Context, key types.Key) {
	if mock.SetKeyFunc == nil {
		panic("KeeperMock.SetKeyFunc: method is nil but Keeper.SetKey was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
		Key types.Key
	}{
		Ctx: ctx,
		Key: key,
	}
	mock.lockSetKey.Lock()
	mock.calls.SetKey = append(mock.calls.SetKey, callInfo)
	mock.lockSetKey.Unlock()
	mock.SetKeyFunc(ctx, key)
}

// SetKeyCalls gets all the calls that were made to SetKey.
// Check the length with:
//
//	len(mockedKeeper.SetKeyCalls())
func (mock *KeeperMock) SetKeyCalls() []struct {
	Ctx sdk.Context
	Key types.Key
} {
	var calls []struct {
		Ctx sdk.Context
		Key types.Key
	}
	mock.lockSetKey.RLock()
	calls = mock.calls.SetKey
	mock.lockSetKey.RUnlock()
	return calls
}

// Ensure, that SnapshotterMock does implement types.Snapshotter.
// If this is not the case, regenerate this file with moq.
var _ types.Snapshotter = &SnapshotterMock{}

// SnapshotterMock is a mock implementation of types.Snapshotter.
//
//	func TestSomethingThatUsesSnapshotter(t *testing.T) {
//
//		// make and configure a mocked types.Snapshotter
//		mockedSnapshotter := &SnapshotterMock{
//			CreateSnapshotFunc: func(ctx sdk.Context, candidates []sdk.ValAddress, filterFunc func(exported.ValidatorI) bool, weightFunc func(consensusPower sdk.Uint) sdk.Uint, threshold utils.Threshold) (exported.Snapshot, error) {
//				panic("mock out the CreateSnapshot method")
//			},
//			GetOperatorFunc: func(ctx sdk.Context, proxy sdk.AccAddress) sdk.ValAddress {
//				panic("mock out the GetOperator method")
//			},
//			GetProxyFunc: func(ctx sdk.Context, operator sdk.ValAddress) (sdk.AccAddress, bool) {
//				panic("mock out the GetProxy method")
//			},
//		}
//
//		// use mockedSnapshotter in code that requires types.Snapshotter
//		// and then make assertions.
//
//	}
type SnapshotterMock struct {
	// CreateSnapshotFunc mocks the CreateSnapshot method.
	CreateSnapshotFunc func(ctx sdk.Context, candidates []sdk.ValAddress, filterFunc func(exported.ValidatorI) bool, weightFunc func(consensusPower sdk.Uint) sdk.Uint, threshold utils.Threshold) (exported.Snapshot, error)

	// GetOperatorFunc mocks the GetOperator method.
	GetOperatorFunc func(ctx sdk.Context, proxy sdk.AccAddress) sdk.ValAddress

	// GetProxyFunc mocks the GetProxy method.
	GetProxyFunc func(ctx sdk.Context, operator sdk.ValAddress) (sdk.AccAddress, bool)

	// calls tracks calls to the methods.
	calls struct {
		// CreateSnapshot holds details about calls to the CreateSnapshot method.
		CreateSnapshot []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Candidates is the candidates argument value.
			Candidates []sdk.ValAddress
			// FilterFunc is the filterFunc argument value.
			FilterFunc func(exported.ValidatorI) bool
			// WeightFunc is the weightFunc argument value.
			WeightFunc func(consensusPower sdk.Uint) sdk.Uint
			// Threshold is the threshold argument value.
			Threshold utils.Threshold
		}
		// GetOperator holds details about calls to the GetOperator method.
		GetOperator []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Proxy is the proxy argument value.
			Proxy sdk.AccAddress
		}
		// GetProxy holds details about calls to the GetProxy method.
		GetProxy []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Operator is the operator argument value.
			Operator sdk.ValAddress
		}
	}
	lockCreateSnapshot sync.RWMutex
	lockGetOperator    sync.RWMutex
	lockGetProxy       sync.RWMutex
}

// CreateSnapshot calls CreateSnapshotFunc.
func (mock *SnapshotterMock) CreateSnapshot(ctx sdk.Context, candidates []sdk.ValAddress, filterFunc func(exported.ValidatorI) bool, weightFunc func(consensusPower sdk.Uint) sdk.Uint, threshold utils.Threshold) (exported.Snapshot, error) {
	if mock.CreateSnapshotFunc == nil {
		panic("SnapshotterMock.CreateSnapshotFunc: method is nil but Snapshotter.CreateSnapshot was just called")
	}
	callInfo := struct {
		Ctx        sdk.Context
		Candidates []sdk.ValAddress
		FilterFunc func(exported.ValidatorI) bool
		WeightFunc func(consensusPower sdk.Uint) sdk.Uint
		Threshold  utils.Threshold
	}{
		Ctx:        ctx,
		Candidates: candidates,
		FilterFunc: filterFunc,
		WeightFunc: weightFunc,
		Threshold:  threshold,
	}
	mock.lockCreateSnapshot.Lock()
	mock.calls.CreateSnapshot = append(mock.calls.CreateSnapshot, callInfo)
	mock.lockCreateSnapshot.Unlock()
	return mock.CreateSnapshotFunc(ctx, candidates, filterFunc, weightFunc, threshold)
}

// CreateSnapshotCalls gets all the calls that were made to CreateSnapshot.
// Check the length with:
//
//	len(mockedSnapshotter.CreateSnapshotCalls())
func (mock *SnapshotterMock) CreateSnapshotCalls() []struct {
	Ctx        sdk.Context
	Candidates []sdk.ValAddress
	FilterFunc func(exported.ValidatorI) bool
	WeightFunc func(consensusPower sdk.Uint) sdk.Uint
	Threshold  utils.Threshold
} {
	var calls []struct {
		Ctx        sdk.Context
		Candidates []sdk.ValAddress
		FilterFunc func(exported.ValidatorI) bool
		WeightFunc func(consensusPower sdk.Uint) sdk.Uint
		Threshold  utils.Threshold
	}
	mock.lockCreateSnapshot.RLock()
	calls = mock.calls.CreateSnapshot
	mock.lockCreateSnapshot.RUnlock()
	return calls
}

// GetOperator calls GetOperatorFunc.
func (mock *SnapshotterMock) GetOperator(ctx sdk.Context, proxy sdk.AccAddress) sdk.ValAddress {
	if mock.GetOperatorFunc == nil {
		panic("SnapshotterMock.GetOperatorFunc: method is nil but Snapshotter.GetOperator was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		Proxy sdk.AccAddress
	}{
		Ctx:   ctx,
		Proxy: proxy,
	}
	mock.lockGetOperator.Lock()
	mock.calls.GetOperator = append(mock.calls.GetOperator, callInfo)
	mock.lockGetOperator.Unlock()
	return mock.GetOperatorFunc(ctx, proxy)
}

// GetOperatorCalls gets all the calls that were made to GetOperator.
// Check the length with:
//
//	len(mockedSnapshotter.GetOperatorCalls())
func (mock *SnapshotterMock) GetOperatorCalls() []struct {
	Ctx   sdk.Context
	Proxy sdk.AccAddress
} {
	var calls []struct {
		Ctx   sdk.Context
		Proxy sdk.AccAddress
	}
	mock.lockGetOperator.RLock()
	calls = mock.calls.GetOperator
	mock.lockGetOperator.RUnlock()
	return calls
}

// GetProxy calls GetProxyFunc.
func (mock *SnapshotterMock) GetProxy(ctx sdk.Context, operator sdk.ValAddress) (sdk.AccAddress, bool) {
	if mock.GetProxyFunc == nil {
		panic("SnapshotterMock.GetProxyFunc: method is nil but Snapshotter.GetProxy was just called")
	}
	callInfo := struct {
		Ctx      sdk.Context
		Operator sdk.ValAddress
	}{
		Ctx:      ctx,
		Operator: operator,
	}
	mock.lockGetProxy.Lock()
	mock.calls.GetProxy = append(mock.calls.GetProxy, callInfo)
	mock.lockGetProxy.Unlock()
	return mock.GetProxyFunc(ctx, operator)
}

// GetProxyCalls gets all the calls that were made to GetProxy.
// Check the length with:
//
//	len(mockedSnapshotter.GetProxyCalls())
func (mock *SnapshotterMock) GetProxyCalls() []struct {
	Ctx      sdk.Context
	Operator sdk.ValAddress
} {
	var calls []struct {
		Ctx      sdk.Context
		Operator sdk.ValAddress
	}
	mock.lockGetProxy.RLock()
	calls = mock.calls.GetProxy
	mock.lockGetProxy.RUnlock()
	return calls
}

// Ensure, that StakerMock does implement types.Staker.
// If this is not the case, regenerate this file with moq.
var _ types.Staker = &StakerMock{}

// StakerMock is a mock implementation of types.Staker.
//
//	func TestSomethingThatUsesStaker(t *testing.T) {
//
//		// make and configure a mocked types.Staker
//		mockedStaker := &StakerMock{
//			GetBondedValidatorsByPowerFunc: func(ctx sdk.Context) []stakingTypes.Validator {
//				panic("mock out the GetBondedValidatorsByPower method")
//			},
//		}
//
//		// use mockedStaker in code that requires types.Staker
//		// and then make assertions.
//
//	}
type StakerMock struct {
	// GetBondedValidatorsByPowerFunc mocks the GetBondedValidatorsByPower method.
	GetBondedValidatorsByPowerFunc func(ctx sdk.Context) []stakingTypes.Validator

	// calls tracks calls to the methods.
	calls struct {
		// GetBondedValidatorsByPower holds details about calls to the GetBondedValidatorsByPower method.
		GetBondedValidatorsByPower []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
		}
	}
	lockGetBondedValidatorsByPower sync.RWMutex
}

// GetBondedValidatorsByPower calls GetBondedValidatorsByPowerFunc.
func (mock *StakerMock) GetBondedValidatorsByPower(ctx sdk.Context) []stakingTypes.Validator {
	if mock.GetBondedValidatorsByPowerFunc == nil {
		panic("StakerMock.GetBondedValidatorsByPowerFunc: method is nil but Staker.GetBondedValidatorsByPower was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetBondedValidatorsByPower.Lock()
	mock.calls.GetBondedValidatorsByPower = append(mock.calls.GetBondedValidatorsByPower, callInfo)
	mock.lockGetBondedValidatorsByPower.Unlock()
	return mock.GetBondedValidatorsByPowerFunc(ctx)
}

// GetBondedValidatorsByPowerCalls gets all the calls that were made to GetBondedValidatorsByPower.
// Check the length with:
//
//	len(mockedStaker.GetBondedValidatorsByPowerCalls())
func (mock *StakerMock) GetBondedValidatorsByPowerCalls() []struct {
	Ctx sdk.Context
} {
	var calls []struct {
		Ctx sdk.Context
	}
	mock.lockGetBondedValidatorsByPower.RLock()
	calls = mock.calls.GetBondedValidatorsByPower
	mock.lockGetBondedValidatorsByPower.RUnlock()
	return calls
}

// Ensure, that SlasherMock does implement types.Slasher.
// If this is not the case, regenerate this file with moq.
var _ types.Slasher = &SlasherMock{}

// SlasherMock is a mock implementation of types.Slasher.
//
//	func TestSomethingThatUsesSlasher(t *testing.T) {
//
//		// make and configure a mocked types.Slasher
//		mockedSlasher := &SlasherMock{
//			IsTombstonedFunc: func(ctx sdk.Context, consAddr sdk.ConsAddress) bool {
//				panic("mock out the IsTombstoned method")
//			},
//		}
//
//		// use mockedSlasher in code that requires types.Slasher
//		// and then make assertions.
//
//	}
type SlasherMock struct {
	// IsTombstonedFunc mocks the IsTombstoned method.
	IsTombstonedFunc func(ctx sdk.Context, consAddr sdk.ConsAddress) bool

	// calls tracks calls to the methods.
	calls struct {
		// IsTombstoned holds details about calls to the IsTombstoned method.
		IsTombstoned []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// ConsAddr is the consAddr argument value.
			ConsAddr sdk.ConsAddress
		}
	}
	lockIsTombstoned sync.RWMutex
}

// IsTombstoned calls IsTombstonedFunc.
func (mock *SlasherMock) IsTombstoned(ctx sdk.Context, consAddr sdk.ConsAddress) bool {
	if mock.IsTombstonedFunc == nil {
		panic("SlasherMock.IsTombstonedFunc: method is nil but Slasher.IsTombstoned was just called")
	}
	callInfo := struct {
		Ctx      sdk.Context
		ConsAddr sdk.ConsAddress
	}{
		Ctx:      ctx,
		ConsAddr: consAddr,
	}
	mock.lockIsTombstoned.Lock()
	mock.calls.IsTombstoned = append(mock.calls.IsTombstoned, callInfo)
	mock.lockIsTombstoned.Unlock()
	return mock.IsTombstonedFunc(ctx, consAddr)
}

// IsTombstonedCalls gets all the calls that were made to IsTombstoned.
// Check the length with:
//
//	len(mockedSlasher.IsTombstonedCalls())
func (mock *SlasherMock) IsTombstonedCalls() []struct {
	Ctx      sdk.Context
	ConsAddr sdk.ConsAddress
} {
	var calls []struct {
		Ctx      sdk.Context
		ConsAddr sdk.ConsAddress
	}
	mock.lockIsTombstoned.RLock()
	calls = mock.calls.IsTombstoned
	mock.lockIsTombstoned.RUnlock()
	return calls
}

// Ensure, that RewarderMock does implement types.Rewarder.
// If this is not the case, regenerate this file with moq.
var _ types.Rewarder = &RewarderMock{}

// RewarderMock is a mock implementation of types.Rewarder.
//
//	func TestSomethingThatUsesRewarder(t *testing.T) {
//
//		// make and configure a mocked types.Rewarder
//		mockedRewarder := &RewarderMock{
//			GetPoolFunc: func(ctx sdk.Context, name string) reward.RewardPool {
//				panic("mock out the GetPool method")
//			},
//		}
//
//		// use mockedRewarder in code that requires types.Rewarder
//		// and then make assertions.
//
//	}
type RewarderMock struct {
	// GetPoolFunc mocks the GetPool method.
	GetPoolFunc func(ctx sdk.Context, name string) reward.RewardPool

	// calls tracks calls to the methods.
	calls struct {
		// GetPool holds details about calls to the GetPool method.
		GetPool []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Name is the name argument value.
			Name string
		}
	}
	lockGetPool sync.RWMutex
}

// GetPool calls GetPoolFunc.
func (mock *RewarderMock) GetPool(ctx sdk.Context, name string) reward.RewardPool {
	if mock.GetPoolFunc == nil {
		panic("RewarderMock.GetPoolFunc: method is nil but Rewarder.GetPool was just called")
	}
	callInfo := struct {
		Ctx  sdk.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockGetPool.Lock()
	mock.calls.GetPool = append(mock.calls.GetPool, callInfo)
	mock.lockGetPool.Unlock()
	return mock.GetPoolFunc(ctx, name)
}

// GetPoolCalls gets all the calls that were made to GetPool.
// Check the length with:
//
//	len(mockedRewarder.GetPoolCalls())
func (mock *RewarderMock) GetPoolCalls() []struct {
	Ctx  sdk.Context
	Name string
} {
	var calls []struct {
		Ctx  sdk.Context
		Name string
	}
	mock.lockGetPool.RLock()
	calls = mock.calls.GetPool
	mock.lockGetPool.RUnlock()
	return calls
}

// Ensure, that NexusMock does implement types.Nexus.
// If this is not the case, regenerate this file with moq.
var _ types.Nexus = &NexusMock{}

// NexusMock is a mock implementation of types.Nexus.
//
//	func TestSomethingThatUsesNexus(t *testing.T) {
//
//		// make and configure a mocked types.Nexus
//		mockedNexus := &NexusMock{
//			GetChainFunc: func(ctx sdk.Context, chain github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName) (github_com_axelarnetwork_axelar_core_x_nexus_exported.Chain, bool) {
//				panic("mock out the GetChain method")
//			},
//			GetChainsFunc: func(ctx sdk.Context) []github_com_axelarnetwork_axelar_core_x_nexus_exported.Chain {
//				panic("mock out the GetChains method")
//			},
//		}
//
//		// use mockedNexus in code that requires types.Nexus
//		// and then make assertions.
//
//	}
type NexusMock struct {
	// GetChainFunc mocks the GetChain method.
	GetChainFunc func(ctx sdk.Context, chain github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName) (github_com_axelarnetwork_axelar_core_x_nexus_exported.Chain, bool)

	// GetChainsFunc mocks the GetChains method.
	GetChainsFunc func(ctx sdk.Context) []github_com_axelarnetwork_axelar_core_x_nexus_exported.Chain

	// calls tracks calls to the methods.
	calls struct {
		// GetChain holds details about calls to the GetChain method.
		GetChain []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Chain is the chain argument value.
			Chain github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
		}
		// GetChains holds details about calls to the GetChains method.
		GetChains []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
		}
	}
	lockGetChain  sync.RWMutex
	lockGetChains sync.RWMutex
}

// GetChain calls GetChainFunc.
func (mock *NexusMock) GetChain(ctx sdk.Context, chain github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName) (github_com_axelarnetwork_axelar_core_x_nexus_exported.Chain, bool) {
	if mock.GetChainFunc == nil {
		panic("NexusMock.GetChainFunc: method is nil but Nexus.GetChain was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		Chain github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
	}{
		Ctx:   ctx,
		Chain: chain,
	}
	mock.lockGetChain.Lock()
	mock.calls.GetChain = append(mock.calls.GetChain, callInfo)
	mock.lockGetChain.Unlock()
	return mock.GetChainFunc(ctx, chain)
}

// GetChainCalls gets all the calls that were made to GetChain.
// Check the length with:
//
//	len(mockedNexus.GetChainCalls())
func (mock *NexusMock) GetChainCalls() []struct {
	Ctx   sdk.Context
	Chain github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
} {
	var calls []struct {
		Ctx   sdk.Context
		Chain github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName
	}
	mock.lockGetChain.RLock()
	calls = mock.calls.GetChain
	mock.lockGetChain.RUnlock()
	return calls
}

// GetChains calls GetChainsFunc.
func (mock *NexusMock) GetChains(ctx sdk.Context) []github_com_axelarnetwork_axelar_core_x_nexus_exported.Chain {
	if mock.GetChainsFunc == nil {
		panic("NexusMock.GetChainsFunc: method is nil but Nexus.GetChains was just called")
	}
	callInfo := struct {
		Ctx sdk.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetChains.Lock()
	mock.calls.GetChains = append(mock.calls.GetChains, callInfo)
	mock.lockGetChains.Unlock()
	return mock.GetChainsFunc(ctx)
}

// GetChainsCalls gets all the calls that were made to GetChains.
// Check the length with:
//
//	len(mockedNexus.GetChainsCalls())
func (mock *NexusMock) GetChainsCalls() []struct {
	Ctx sdk.Context
} {
	var calls []struct {
		Ctx sdk.Context
	}
	mock.lockGetChains.RLock()
	calls = mock.calls.GetChains
	mock.lockGetChains.RUnlock()
	return calls
}

// Ensure, that KeygenParticipatorMock does implement types.KeygenParticipator.
// If this is not the case, regenerate this file with moq.
var _ types.KeygenParticipator = &KeygenParticipatorMock{}

// KeygenParticipatorMock is a mock implementation of types.KeygenParticipator.
//
//	func TestSomethingThatUsesKeygenParticipator(t *testing.T) {
//
//		// make and configure a mocked types.KeygenParticipator
//		mockedKeygenParticipator := &KeygenParticipatorMock{
//			HasOptedOutFunc: func(ctx sdk.Context, participant sdk.AccAddress) bool {
//				panic("mock out the HasOptedOut method")
//			},
//		}
//
//		// use mockedKeygenParticipator in code that requires types.KeygenParticipator
//		// and then make assertions.
//
//	}
type KeygenParticipatorMock struct {
	// HasOptedOutFunc mocks the HasOptedOut method.
	HasOptedOutFunc func(ctx sdk.Context, participant sdk.AccAddress) bool

	// calls tracks calls to the methods.
	calls struct {
		// HasOptedOut holds details about calls to the HasOptedOut method.
		HasOptedOut []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Participant is the participant argument value.
			Participant sdk.AccAddress
		}
	}
	lockHasOptedOut sync.RWMutex
}

// HasOptedOut calls HasOptedOutFunc.
func (mock *KeygenParticipatorMock) HasOptedOut(ctx sdk.Context, participant sdk.AccAddress) bool {
	if mock.HasOptedOutFunc == nil {
		panic("KeygenParticipatorMock.HasOptedOutFunc: method is nil but KeygenParticipator.HasOptedOut was just called")
	}
	callInfo := struct {
		Ctx         sdk.Context
		Participant sdk.AccAddress
	}{
		Ctx:         ctx,
		Participant: participant,
	}
	mock.lockHasOptedOut.Lock()
	mock.calls.HasOptedOut = append(mock.calls.HasOptedOut, callInfo)
	mock.lockHasOptedOut.Unlock()
	return mock.HasOptedOutFunc(ctx, participant)
}

// HasOptedOutCalls gets all the calls that were made to HasOptedOut.
// Check the length with:
//
//	len(mockedKeygenParticipator.HasOptedOutCalls())
func (mock *KeygenParticipatorMock) HasOptedOutCalls() []struct {
	Ctx         sdk.Context
	Participant sdk.AccAddress
} {
	var calls []struct {
		Ctx         sdk.Context
		Participant sdk.AccAddress
	}
	mock.lockHasOptedOut.RLock()
	calls = mock.calls.HasOptedOut
	mock.lockHasOptedOut.RUnlock()
	return calls
}
