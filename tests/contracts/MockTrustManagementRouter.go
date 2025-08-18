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

// MockTrustManagementRouterPermit is an auto generated low-level Go binding around an user-defined struct.
type MockTrustManagementRouterPermit struct {
	Amount   *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

// MockTrustManagementRouterTransaction is an auto generated low-level Go binding around an user-defined struct.
type MockTrustManagementRouterTransaction struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
}

// MockTrustManagementRouterMetaData contains all meta data concerning the MockTrustManagementRouter contract.
var MockTrustManagementRouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositWithPermit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.Permit\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"txs\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Transaction[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"depositIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amountsWithYield\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506105828061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806320e8c5651461004e5780632f37da741461007e578063c860efa81461009a578063d3413cac146100ca575b5f5ffd5b610068600480360381019061006391906101a0565b6100e6565b6040516100759190610213565b60405180910390f35b610098600480360381019061009391906102e2565b6100ef565b005b6100b460048036038101906100af9190610395565b6100f6565b6040516100c19190610213565b60405180910390f35b6100e460048036038101906100df919061044f565b6100ff565b005b5f949350505050565b5050505050565b5f949350505050565b50505050505050505050565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61013c82610113565b9050919050565b61014c81610132565b8114610156575f5ffd5b50565b5f8135905061016781610143565b92915050565b5f819050919050565b61017f8161016d565b8114610189575f5ffd5b50565b5f8135905061019a81610176565b92915050565b5f5f5f5f608085870312156101b8576101b761010b565b5b5f6101c587828801610159565b94505060206101d687828801610159565b93505060406101e78782880161018c565b92505060606101f88782880161018c565b91505092959194509250565b61020d81610132565b82525050565b5f6020820190506102265f830184610204565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261024d5761024c61022c565b5b8235905067ffffffffffffffff81111561026a57610269610230565b5b60208301915083602082028301111561028657610285610234565b5b9250929050565b5f5f83601f8401126102a2576102a161022c565b5b8235905067ffffffffffffffff8111156102bf576102be610230565b5b6020830191508360018202830111156102db576102da610234565b5b9250929050565b5f5f5f5f5f606086880312156102fb576102fa61010b565b5b5f86013567ffffffffffffffff8111156103185761031761010f565b5b61032488828901610238565b9550955050602086013567ffffffffffffffff8111156103475761034661010f565b5b6103538882890161028d565b935093505060406103668882890161018c565b9150509295509295909350565b5f5ffd5b5f60a0828403121561038c5761038b610373565b5b81905092915050565b5f5f5f5f61010085870312156103ae576103ad61010b565b5b5f6103bb87828801610159565b94505060206103cc87828801610159565b93505060406103dd8782880161018c565b92505060606103ee87828801610377565b91505092959194509250565b5f5f83601f84011261040f5761040e61022c565b5b8235905067ffffffffffffffff81111561042c5761042b610230565b5b60208301915083602082028301111561044857610447610234565b5b9250929050565b5f5f5f5f5f5f5f5f5f5f60e08b8d03121561046d5761046c61010b565b5b5f61047a8d828e01610159565b9a5050602061048b8d828e01610159565b995050604061049c8d828e01610159565b98505060608b013567ffffffffffffffff8111156104bd576104bc61010f565b5b6104c98d828e016103fa565b975097505060808b013567ffffffffffffffff8111156104ec576104eb61010f565b5b6104f88d828e016103fa565b955095505060a08b013567ffffffffffffffff81111561051b5761051a61010f565b5b6105278d828e0161028d565b935093505060c061053a8d828e0161018c565b9150509295989b9194979a509295985056fea264697066735822122031859329241f91d756646ce2f1c9f86b8bb4d66143b075a7d1efc0ae2928788664736f6c634300081e0033",
}

// MockTrustManagementRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use MockTrustManagementRouterMetaData.ABI instead.
var MockTrustManagementRouterABI = MockTrustManagementRouterMetaData.ABI

// MockTrustManagementRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockTrustManagementRouterMetaData.Bin instead.
var MockTrustManagementRouterBin = MockTrustManagementRouterMetaData.Bin

// DeployMockTrustManagementRouter deploys a new Ethereum contract, binding an instance of MockTrustManagementRouter to it.
func DeployMockTrustManagementRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockTrustManagementRouter, error) {
	parsed, err := MockTrustManagementRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockTrustManagementRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockTrustManagementRouter{MockTrustManagementRouterCaller: MockTrustManagementRouterCaller{contract: contract}, MockTrustManagementRouterTransactor: MockTrustManagementRouterTransactor{contract: contract}, MockTrustManagementRouterFilterer: MockTrustManagementRouterFilterer{contract: contract}}, nil
}

// MockTrustManagementRouter is an auto generated Go binding around an Ethereum contract.
type MockTrustManagementRouter struct {
	MockTrustManagementRouterCaller     // Read-only binding to the contract
	MockTrustManagementRouterTransactor // Write-only binding to the contract
	MockTrustManagementRouterFilterer   // Log filterer for contract events
}

// MockTrustManagementRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockTrustManagementRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTrustManagementRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockTrustManagementRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTrustManagementRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockTrustManagementRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTrustManagementRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockTrustManagementRouterSession struct {
	Contract     *MockTrustManagementRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MockTrustManagementRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockTrustManagementRouterCallerSession struct {
	Contract *MockTrustManagementRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// MockTrustManagementRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockTrustManagementRouterTransactorSession struct {
	Contract     *MockTrustManagementRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// MockTrustManagementRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockTrustManagementRouterRaw struct {
	Contract *MockTrustManagementRouter // Generic contract binding to access the raw methods on
}

// MockTrustManagementRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockTrustManagementRouterCallerRaw struct {
	Contract *MockTrustManagementRouterCaller // Generic read-only contract binding to access the raw methods on
}

// MockTrustManagementRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockTrustManagementRouterTransactorRaw struct {
	Contract *MockTrustManagementRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockTrustManagementRouter creates a new instance of MockTrustManagementRouter, bound to a specific deployed contract.
func NewMockTrustManagementRouter(address common.Address, backend bind.ContractBackend) (*MockTrustManagementRouter, error) {
	contract, err := bindMockTrustManagementRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockTrustManagementRouter{MockTrustManagementRouterCaller: MockTrustManagementRouterCaller{contract: contract}, MockTrustManagementRouterTransactor: MockTrustManagementRouterTransactor{contract: contract}, MockTrustManagementRouterFilterer: MockTrustManagementRouterFilterer{contract: contract}}, nil
}

// NewMockTrustManagementRouterCaller creates a new read-only instance of MockTrustManagementRouter, bound to a specific deployed contract.
func NewMockTrustManagementRouterCaller(address common.Address, caller bind.ContractCaller) (*MockTrustManagementRouterCaller, error) {
	contract, err := bindMockTrustManagementRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockTrustManagementRouterCaller{contract: contract}, nil
}

// NewMockTrustManagementRouterTransactor creates a new write-only instance of MockTrustManagementRouter, bound to a specific deployed contract.
func NewMockTrustManagementRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*MockTrustManagementRouterTransactor, error) {
	contract, err := bindMockTrustManagementRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockTrustManagementRouterTransactor{contract: contract}, nil
}

// NewMockTrustManagementRouterFilterer creates a new log filterer instance of MockTrustManagementRouter, bound to a specific deployed contract.
func NewMockTrustManagementRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*MockTrustManagementRouterFilterer, error) {
	contract, err := bindMockTrustManagementRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockTrustManagementRouterFilterer{contract: contract}, nil
}

// bindMockTrustManagementRouter binds a generic wrapper to an already deployed contract.
func bindMockTrustManagementRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockTrustManagementRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTrustManagementRouter *MockTrustManagementRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTrustManagementRouter.Contract.MockTrustManagementRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTrustManagementRouter *MockTrustManagementRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.MockTrustManagementRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTrustManagementRouter *MockTrustManagementRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.MockTrustManagementRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockTrustManagementRouter *MockTrustManagementRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTrustManagementRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x20e8c565.
//
// Solidity: function deposit(address receiver, address token, uint256 amount, uint256 lockDuration) returns(address walletAddress)
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactor) Deposit(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, lockDuration *big.Int) (*types.Transaction, error) {
	return _MockTrustManagementRouter.contract.Transact(opts, "deposit", receiver, token, amount, lockDuration)
}

// Deposit is a paid mutator transaction binding the contract method 0x20e8c565.
//
// Solidity: function deposit(address receiver, address token, uint256 amount, uint256 lockDuration) returns(address walletAddress)
func (_MockTrustManagementRouter *MockTrustManagementRouterSession) Deposit(receiver common.Address, token common.Address, amount *big.Int, lockDuration *big.Int) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.Deposit(&_MockTrustManagementRouter.TransactOpts, receiver, token, amount, lockDuration)
}

// Deposit is a paid mutator transaction binding the contract method 0x20e8c565.
//
// Solidity: function deposit(address receiver, address token, uint256 amount, uint256 lockDuration) returns(address walletAddress)
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactorSession) Deposit(receiver common.Address, token common.Address, amount *big.Int, lockDuration *big.Int) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.Deposit(&_MockTrustManagementRouter.TransactOpts, receiver, token, amount, lockDuration)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0xc860efa8.
//
// Solidity: function depositWithPermit(address receiver, address token, uint256 lockDuration, (uint256,uint256,uint8,bytes32,bytes32) permit) returns(address walletAddress)
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactor) DepositWithPermit(opts *bind.TransactOpts, receiver common.Address, token common.Address, lockDuration *big.Int, permit MockTrustManagementRouterPermit) (*types.Transaction, error) {
	return _MockTrustManagementRouter.contract.Transact(opts, "depositWithPermit", receiver, token, lockDuration, permit)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0xc860efa8.
//
// Solidity: function depositWithPermit(address receiver, address token, uint256 lockDuration, (uint256,uint256,uint8,bytes32,bytes32) permit) returns(address walletAddress)
func (_MockTrustManagementRouter *MockTrustManagementRouterSession) DepositWithPermit(receiver common.Address, token common.Address, lockDuration *big.Int, permit MockTrustManagementRouterPermit) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.DepositWithPermit(&_MockTrustManagementRouter.TransactOpts, receiver, token, lockDuration, permit)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0xc860efa8.
//
// Solidity: function depositWithPermit(address receiver, address token, uint256 lockDuration, (uint256,uint256,uint8,bytes32,bytes32) permit) returns(address walletAddress)
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactorSession) DepositWithPermit(receiver common.Address, token common.Address, lockDuration *big.Int, permit MockTrustManagementRouterPermit) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.DepositWithPermit(&_MockTrustManagementRouter.TransactOpts, receiver, token, lockDuration, permit)
}

// Execute is a paid mutator transaction binding the contract method 0x2f37da74.
//
// Solidity: function execute((address,uint256,bytes)[] txs, bytes signature, uint256 deadline) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactor) Execute(opts *bind.TransactOpts, txs []MockTrustManagementRouterTransaction, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _MockTrustManagementRouter.contract.Transact(opts, "execute", txs, signature, deadline)
}

// Execute is a paid mutator transaction binding the contract method 0x2f37da74.
//
// Solidity: function execute((address,uint256,bytes)[] txs, bytes signature, uint256 deadline) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterSession) Execute(txs []MockTrustManagementRouterTransaction, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.Execute(&_MockTrustManagementRouter.TransactOpts, txs, signature, deadline)
}

// Execute is a paid mutator transaction binding the contract method 0x2f37da74.
//
// Solidity: function execute((address,uint256,bytes)[] txs, bytes signature, uint256 deadline) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactorSession) Execute(txs []MockTrustManagementRouterTransaction, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.Execute(&_MockTrustManagementRouter.TransactOpts, txs, signature, deadline)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd3413cac.
//
// Solidity: function withdraw(address wallet, address token, address receiver, uint256[] depositIds, uint256[] amountsWithYield, bytes signature, uint256 deadline) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactor) Withdraw(opts *bind.TransactOpts, wallet common.Address, token common.Address, receiver common.Address, depositIds []*big.Int, amountsWithYield []*big.Int, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _MockTrustManagementRouter.contract.Transact(opts, "withdraw", wallet, token, receiver, depositIds, amountsWithYield, signature, deadline)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd3413cac.
//
// Solidity: function withdraw(address wallet, address token, address receiver, uint256[] depositIds, uint256[] amountsWithYield, bytes signature, uint256 deadline) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterSession) Withdraw(wallet common.Address, token common.Address, receiver common.Address, depositIds []*big.Int, amountsWithYield []*big.Int, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.Withdraw(&_MockTrustManagementRouter.TransactOpts, wallet, token, receiver, depositIds, amountsWithYield, signature, deadline)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd3413cac.
//
// Solidity: function withdraw(address wallet, address token, address receiver, uint256[] depositIds, uint256[] amountsWithYield, bytes signature, uint256 deadline) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactorSession) Withdraw(wallet common.Address, token common.Address, receiver common.Address, depositIds []*big.Int, amountsWithYield []*big.Int, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.Withdraw(&_MockTrustManagementRouter.TransactOpts, wallet, token, receiver, depositIds, amountsWithYield, signature, deadline)
}
