// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MockAavePoolMetaData contains all meta data concerning the MockAavePool contract.
var MockAavePoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"supply\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"referralCode\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506101b18061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063617ba0371461002d575b5f5ffd5b61004760048036038101906100429190610117565b610049565b005b50505050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61007c82610053565b9050919050565b61008c81610072565b8114610096575f5ffd5b50565b5f813590506100a781610083565b92915050565b5f819050919050565b6100bf816100ad565b81146100c9575f5ffd5b50565b5f813590506100da816100b6565b92915050565b5f61ffff82169050919050565b6100f6816100e0565b8114610100575f5ffd5b50565b5f81359050610111816100ed565b92915050565b5f5f5f5f6080858703121561012f5761012e61004f565b5b5f61013c87828801610099565b945050602061014d878288016100cc565b935050604061015e87828801610099565b925050606061016f87828801610103565b9150509295919450925056fea264697066735822122066d3fc3f239f59442357a3c510b4dc8e692dce35b7ab42ee177d5d164c0a212264736f6c634300081e0033",
}

// MockAavePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use MockAavePoolMetaData.ABI instead.
var MockAavePoolABI = MockAavePoolMetaData.ABI

// MockAavePoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockAavePoolMetaData.Bin instead.
var MockAavePoolBin = MockAavePoolMetaData.Bin

// DeployMockAavePool deploys a new Ethereum contract, binding an instance of MockAavePool to it.
func DeployMockAavePool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockAavePool, error) {
	parsed, err := MockAavePoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockAavePoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockAavePool{MockAavePoolCaller: MockAavePoolCaller{contract: contract}, MockAavePoolTransactor: MockAavePoolTransactor{contract: contract}, MockAavePoolFilterer: MockAavePoolFilterer{contract: contract}}, nil
}

// MockAavePool is an auto generated Go binding around an Ethereum contract.
type MockAavePool struct {
	MockAavePoolCaller     // Read-only binding to the contract
	MockAavePoolTransactor // Write-only binding to the contract
	MockAavePoolFilterer   // Log filterer for contract events
}

// MockAavePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockAavePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockAavePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockAavePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockAavePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockAavePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockAavePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockAavePoolSession struct {
	Contract     *MockAavePool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockAavePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockAavePoolCallerSession struct {
	Contract *MockAavePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MockAavePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockAavePoolTransactorSession struct {
	Contract     *MockAavePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MockAavePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockAavePoolRaw struct {
	Contract *MockAavePool // Generic contract binding to access the raw methods on
}

// MockAavePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockAavePoolCallerRaw struct {
	Contract *MockAavePoolCaller // Generic read-only contract binding to access the raw methods on
}

// MockAavePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockAavePoolTransactorRaw struct {
	Contract *MockAavePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockAavePool creates a new instance of MockAavePool, bound to a specific deployed contract.
func NewMockAavePool(address common.Address, backend bind.ContractBackend) (*MockAavePool, error) {
	contract, err := bindMockAavePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockAavePool{MockAavePoolCaller: MockAavePoolCaller{contract: contract}, MockAavePoolTransactor: MockAavePoolTransactor{contract: contract}, MockAavePoolFilterer: MockAavePoolFilterer{contract: contract}}, nil
}

// NewMockAavePoolCaller creates a new read-only instance of MockAavePool, bound to a specific deployed contract.
func NewMockAavePoolCaller(address common.Address, caller bind.ContractCaller) (*MockAavePoolCaller, error) {
	contract, err := bindMockAavePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockAavePoolCaller{contract: contract}, nil
}

// NewMockAavePoolTransactor creates a new write-only instance of MockAavePool, bound to a specific deployed contract.
func NewMockAavePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*MockAavePoolTransactor, error) {
	contract, err := bindMockAavePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockAavePoolTransactor{contract: contract}, nil
}

// NewMockAavePoolFilterer creates a new log filterer instance of MockAavePool, bound to a specific deployed contract.
func NewMockAavePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*MockAavePoolFilterer, error) {
	contract, err := bindMockAavePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockAavePoolFilterer{contract: contract}, nil
}

// bindMockAavePool binds a generic wrapper to an already deployed contract.
func bindMockAavePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockAavePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockAavePool *MockAavePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockAavePool.Contract.MockAavePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockAavePool *MockAavePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockAavePool.Contract.MockAavePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockAavePool *MockAavePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockAavePool.Contract.MockAavePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockAavePool *MockAavePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockAavePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockAavePool *MockAavePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockAavePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockAavePool *MockAavePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockAavePool.Contract.contract.Transact(opts, method, params...)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_MockAavePool *MockAavePoolTransactor) Supply(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _MockAavePool.contract.Transact(opts, "supply", asset, amount, onBehalfOf, referralCode)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_MockAavePool *MockAavePoolSession) Supply(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _MockAavePool.Contract.Supply(&_MockAavePool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_MockAavePool *MockAavePoolTransactorSession) Supply(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _MockAavePool.Contract.Supply(&_MockAavePool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}
