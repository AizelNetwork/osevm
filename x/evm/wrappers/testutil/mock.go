// Code generated by MockGen. DO NOT EDIT.
// Source: ./x/evm/types/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=./x/evm/types/interfaces.go -package testutil -destination=./x/evm/wrappers/testutil/mock.go
//

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	big "math/big"
	reflect "reflect"

	math "cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/types"
	types3 "github.com/AizelNetwork/osevm/x/feemarket/types"
	gomock "go.uber.org/mock/gomock"
)

// MockFeeMarketKeeper is a mock of FeeMarketKeeper interface.
type MockFeeMarketKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockFeeMarketKeeperMockRecorder
	isgomock struct{}
}

// MockFeeMarketKeeperMockRecorder is the mock recorder for MockFeeMarketKeeper.
type MockFeeMarketKeeperMockRecorder struct {
	mock *MockFeeMarketKeeper
}

// NewMockFeeMarketKeeper creates a new mock instance.
func NewMockFeeMarketKeeper(ctrl *gomock.Controller) *MockFeeMarketKeeper {
	mock := &MockFeeMarketKeeper{ctrl: ctrl}
	mock.recorder = &MockFeeMarketKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeeMarketKeeper) EXPECT() *MockFeeMarketKeeperMockRecorder {
	return m.recorder
}

// CalculateBaseFee mocks base method.
func (m *MockFeeMarketKeeper) CalculateBaseFee(ctx types.Context) math.LegacyDec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateBaseFee", ctx)
	ret0, _ := ret[0].(math.LegacyDec)
	return ret0
}

// CalculateBaseFee indicates an expected call of CalculateBaseFee.
func (mr *MockFeeMarketKeeperMockRecorder) CalculateBaseFee(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateBaseFee", reflect.TypeOf((*MockFeeMarketKeeper)(nil).CalculateBaseFee), ctx)
}

// GetBaseFee mocks base method.
func (m *MockFeeMarketKeeper) GetBaseFee(ctx types.Context) math.LegacyDec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBaseFee", ctx)
	ret0, _ := ret[0].(math.LegacyDec)
	return ret0
}

// GetBaseFee indicates an expected call of GetBaseFee.
func (mr *MockFeeMarketKeeperMockRecorder) GetBaseFee(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaseFee", reflect.TypeOf((*MockFeeMarketKeeper)(nil).GetBaseFee), ctx)
}

// GetParams mocks base method.
func (m *MockFeeMarketKeeper) GetParams(ctx types.Context) types3.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types3.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockFeeMarketKeeperMockRecorder) GetParams(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockFeeMarketKeeper)(nil).GetParams), ctx)
}

// MockBankWrapper is a mock of BankWrapper interface.
type MockBankWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockBankWrapperMockRecorder
	isgomock struct{}
}

// MockBankWrapperMockRecorder is the mock recorder for MockBankWrapper.
type MockBankWrapperMockRecorder struct {
	mock *MockBankWrapper
}

// NewMockBankWrapper creates a new mock instance.
func NewMockBankWrapper(ctrl *gomock.Controller) *MockBankWrapper {
	mock := &MockBankWrapper{ctrl: ctrl}
	mock.recorder = &MockBankWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankWrapper) EXPECT() *MockBankWrapperMockRecorder {
	return m.recorder
}

// BurnAmountFromAccount mocks base method.
func (m *MockBankWrapper) BurnAmountFromAccount(ctx context.Context, account types.AccAddress, amt *big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnAmountFromAccount", ctx, account, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnAmountFromAccount indicates an expected call of BurnAmountFromAccount.
func (mr *MockBankWrapperMockRecorder) BurnAmountFromAccount(ctx, account, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnAmountFromAccount", reflect.TypeOf((*MockBankWrapper)(nil).BurnAmountFromAccount), ctx, account, amt)
}

// BurnCoins mocks base method.
func (m *MockBankWrapper) BurnCoins(ctx context.Context, moduleName string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnCoins indicates an expected call of BurnCoins.
func (mr *MockBankWrapperMockRecorder) BurnCoins(ctx, moduleName, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnCoins", reflect.TypeOf((*MockBankWrapper)(nil).BurnCoins), ctx, moduleName, amt)
}

// GetAllBalances mocks base method.
func (m *MockBankWrapper) GetAllBalances(ctx context.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankWrapperMockRecorder) GetAllBalances(ctx, addr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankWrapper)(nil).GetAllBalances), ctx, addr)
}

// GetBalance mocks base method.
func (m *MockBankWrapper) GetBalance(ctx context.Context, addr types.AccAddress, denom string) types.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, addr, denom)
	ret0, _ := ret[0].(types.Coin)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockBankWrapperMockRecorder) GetBalance(ctx, addr, denom any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBankWrapper)(nil).GetBalance), ctx, addr, denom)
}

// IsSendEnabledCoins mocks base method.
func (m *MockBankWrapper) IsSendEnabledCoins(ctx context.Context, coins ...types.Coin) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range coins {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsSendEnabledCoins", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsSendEnabledCoins indicates an expected call of IsSendEnabledCoins.
func (mr *MockBankWrapperMockRecorder) IsSendEnabledCoins(ctx any, coins ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, coins...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSendEnabledCoins", reflect.TypeOf((*MockBankWrapper)(nil).IsSendEnabledCoins), varargs...)
}

// MintAmountToAccount mocks base method.
func (m *MockBankWrapper) MintAmountToAccount(ctx context.Context, recipientAddr types.AccAddress, amt *big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintAmountToAccount", ctx, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintAmountToAccount indicates an expected call of MintAmountToAccount.
func (mr *MockBankWrapperMockRecorder) MintAmountToAccount(ctx, recipientAddr, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintAmountToAccount", reflect.TypeOf((*MockBankWrapper)(nil).MintAmountToAccount), ctx, recipientAddr, amt)
}

// MintCoins mocks base method.
func (m *MockBankWrapper) MintCoins(ctx context.Context, moduleName string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintCoins indicates an expected call of MintCoins.
func (mr *MockBankWrapperMockRecorder) MintCoins(ctx, moduleName, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCoins", reflect.TypeOf((*MockBankWrapper)(nil).MintCoins), ctx, moduleName, amt)
}

// SendCoins mocks base method.
func (m *MockBankWrapper) SendCoins(ctx context.Context, from, to types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", ctx, from, to, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockBankWrapperMockRecorder) SendCoins(ctx, from, to, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockBankWrapper)(nil).SendCoins), ctx, from, to, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankWrapper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankWrapperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankWrapper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankWrapper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankWrapperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankWrapper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}
