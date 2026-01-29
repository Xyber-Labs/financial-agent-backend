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

// MockTEEWalletChunkedSGXQuote is an auto generated low-level Go binding around an user-defined struct.
type MockTEEWalletChunkedSGXQuote struct {
	Header               []byte
	IsvReport            []byte
	IsvReportSignature   []byte
	AttestationKey       []byte
	QeReport             []byte
	QeReportSignature    []byte
	QeAuthenticationData []byte
}

// MockTEEWalletChunkedX509Cert is an auto generated low-level Go binding around an user-defined struct.
type MockTEEWalletChunkedX509Cert struct {
	BodyPartOne []byte
	PublicKey   []byte
	BodyPartTwo []byte
	Signature   []byte
}

// MockTEEWalletMetaData contains all meta data concerning the MockTEEWallet contract.
var MockTEEWalletMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"initSessionKey\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structMockTEEWallet.ChunkedX509Cert\",\"components\":[{\"name\":\"bodyPartOne\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"bodyPartTwo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"intermediate\",\"type\":\"tuple\",\"internalType\":\"structMockTEEWallet.ChunkedX509Cert\",\"components\":[{\"name\":\"bodyPartOne\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"bodyPartTwo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"quote\",\"type\":\"tuple\",\"internalType\":\"structMockTEEWallet.ChunkedSGXQuote\",\"components\":[{\"name\":\"header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"isvReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"isvReportSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"attestationKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeReportSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeAuthenticationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506101df8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c80634c7c04c41461002d575b5f5ffd5b610047600480360381019061004291906100f1565b610049565b005b50505050565b5f5ffd5b5f5ffd5b5f5ffd5b5f608082840312156100705761006f610057565b5b81905092915050565b5f60e0828403121561008e5761008d610057565b5b81905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100c082610097565b9050919050565b6100d0816100b6565b81146100da575f5ffd5b50565b5f813590506100eb816100c7565b92915050565b5f5f5f5f608085870312156101095761010861004f565b5b5f85013567ffffffffffffffff81111561012657610125610053565b5b6101328782880161005b565b945050602085013567ffffffffffffffff81111561015357610152610053565b5b61015f8782880161005b565b935050604085013567ffffffffffffffff8111156101805761017f610053565b5b61018c87828801610079565b925050606061019d878288016100dd565b9150509295919450925056fea2646970667358221220cebf37883c717f7b6937ad3b9aeb7532840807315a1495f5da90a03713385f6364736f6c634300081e0033",
}

// MockTEEWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use MockTEEWalletMetaData.ABI instead.
var MockTEEWalletABI = MockTEEWalletMetaData.ABI

// MockTEEWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockTEEWalletMetaData.Bin instead.
var MockTEEWalletBin = MockTEEWalletMetaData.Bin

// DeployMockTEEWallet deploys a new Ethereum contract, binding an instance of MockTEEWallet to it.
func DeployMockTEEWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockTEEWallet, error) {
	parsed, err := MockTEEWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockTEEWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockTEEWallet{MockTEEWalletCaller: MockTEEWalletCaller{contract: contract}, MockTEEWalletTransactor: MockTEEWalletTransactor{contract: contract}, MockTEEWalletFilterer: MockTEEWalletFilterer{contract: contract}}, nil
}

// MockTEEWallet is an auto generated Go binding around an Ethereum contract.
type MockTEEWallet struct {
	MockTEEWalletCaller     // Read-only binding to the contract
	MockTEEWalletTransactor // Write-only binding to the contract
	MockTEEWalletFilterer   // Log filterer for contract events
}

// MockTEEWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockTEEWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTEEWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockTEEWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTEEWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockTEEWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTEEWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockTEEWalletSession struct {
	Contract     *MockTEEWallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockTEEWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockTEEWalletCallerSession struct {
	Contract *MockTEEWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MockTEEWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockTEEWalletTransactorSession struct {
	Contract     *MockTEEWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MockTEEWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockTEEWalletRaw struct {
	Contract *MockTEEWallet // Generic contract binding to access the raw methods on
}

// MockTEEWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockTEEWalletCallerRaw struct {
	Contract *MockTEEWalletCaller // Generic read-only contract binding to access the raw methods on
}

// MockTEEWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockTEEWalletTransactorRaw struct {
	Contract *MockTEEWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockTEEWallet creates a new instance of MockTEEWallet, bound to a specific deployed contract.
func NewMockTEEWallet(address common.Address, backend bind.ContractBackend) (*MockTEEWallet, error) {
	contract, err := bindMockTEEWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockTEEWallet{MockTEEWalletCaller: MockTEEWalletCaller{contract: contract}, MockTEEWalletTransactor: MockTEEWalletTransactor{contract: contract}, MockTEEWalletFilterer: MockTEEWalletFilterer{contract: contract}}, nil
}

// NewMockTEEWalletCaller creates a new read-only instance of MockTEEWallet, bound to a specific deployed contract.
func NewMockTEEWalletCaller(address common.Address, caller bind.ContractCaller) (*MockTEEWalletCaller, error) {
	contract, err := bindMockTEEWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockTEEWalletCaller{contract: contract}, nil
}

// NewMockTEEWalletTransactor creates a new write-only instance of MockTEEWallet, bound to a specific deployed contract.
func NewMockTEEWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*MockTEEWalletTransactor, error) {
	contract, err := bindMockTEEWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockTEEWalletTransactor{contract: contract}, nil
}

// NewMockTEEWalletFilterer creates a new log filterer instance of MockTEEWallet, bound to a specific deployed contract.
func NewMockTEEWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*MockTEEWalletFilterer, error) {
	contract, err := bindMockTEEWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockTEEWalletFilterer{contract: contract}, nil
}

// bindMockTEEWallet binds a generic wrapper to an already deployed contract.
func bindMockTEEWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockTEEWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTEEWallet *MockTEEWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTEEWallet.Contract.MockTEEWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTEEWallet *MockTEEWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTEEWallet.Contract.MockTEEWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTEEWallet *MockTEEWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTEEWallet.Contract.MockTEEWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTEEWallet *MockTEEWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTEEWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTEEWallet *MockTEEWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTEEWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTEEWallet *MockTEEWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTEEWallet.Contract.contract.Transact(opts, method, params...)
}

// InitSessionKey is a paid mutator transaction binding the contract method 0x4c7c04c4.
//
// Solidity: function initSessionKey((bytes,bytes,bytes,bytes) leaf, (bytes,bytes,bytes,bytes) intermediate, (bytes,bytes,bytes,bytes,bytes,bytes,bytes) quote, address sessionKey) returns()
func (_MockTEEWallet *MockTEEWalletTransactor) InitSessionKey(opts *bind.TransactOpts, leaf MockTEEWalletChunkedX509Cert, intermediate MockTEEWalletChunkedX509Cert, quote MockTEEWalletChunkedSGXQuote, sessionKey common.Address) (*types.Transaction, error) {
	return _MockTEEWallet.contract.Transact(opts, "initSessionKey", leaf, intermediate, quote, sessionKey)
}

// InitSessionKey is a paid mutator transaction binding the contract method 0x4c7c04c4.
//
// Solidity: function initSessionKey((bytes,bytes,bytes,bytes) leaf, (bytes,bytes,bytes,bytes) intermediate, (bytes,bytes,bytes,bytes,bytes,bytes,bytes) quote, address sessionKey) returns()
func (_MockTEEWallet *MockTEEWalletSession) InitSessionKey(leaf MockTEEWalletChunkedX509Cert, intermediate MockTEEWalletChunkedX509Cert, quote MockTEEWalletChunkedSGXQuote, sessionKey common.Address) (*types.Transaction, error) {
	return _MockTEEWallet.Contract.InitSessionKey(&_MockTEEWallet.TransactOpts, leaf, intermediate, quote, sessionKey)
}

// InitSessionKey is a paid mutator transaction binding the contract method 0x4c7c04c4.
//
// Solidity: function initSessionKey((bytes,bytes,bytes,bytes) leaf, (bytes,bytes,bytes,bytes) intermediate, (bytes,bytes,bytes,bytes,bytes,bytes,bytes) quote, address sessionKey) returns()
func (_MockTEEWallet *MockTEEWalletTransactorSession) InitSessionKey(leaf MockTEEWalletChunkedX509Cert, intermediate MockTEEWalletChunkedX509Cert, quote MockTEEWalletChunkedSGXQuote, sessionKey common.Address) (*types.Transaction, error) {
	return _MockTEEWallet.Contract.InitSessionKey(&_MockTEEWallet.TransactOpts, leaf, intermediate, quote, sessionKey)
}
