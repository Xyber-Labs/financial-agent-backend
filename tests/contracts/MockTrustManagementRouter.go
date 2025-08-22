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

// MockTrustManagementRouterDeposit is an auto generated low-level Go binding around an user-defined struct.
type MockTrustManagementRouterDeposit struct {
	Amount      *big.Int
	LockedUntil *big.Int
}

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
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositWithPermit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.Permit\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"txs\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Transaction[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Deposit[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mockSetDeposits\",\"inputs\":[{\"name\":\"deposits\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Deposit[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"depositIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amountsWithYield\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610a518061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c8063018fa3601461006457806320e8c565146100945780632f37da74146100c4578063385270b9146100e0578063c860efa8146100fc578063d3413cac1461012c575b5f5ffd5b61007e600480360381019061007991906102f7565b610148565b60405161008b9190610422565b60405180910390f35b6100ae60048036038101906100a9919061046c565b6101b9565b6040516100bb91906104df565b60405180910390f35b6100de60048036038101906100d991906105ae565b6101c2565b005b6100fa60048036038101906100f59190610694565b6101c9565b005b61011660048036038101906101119190610701565b610280565b60405161012391906104df565b60405180910390f35b610146600480360381019061014191906107bb565b610289565b005b60605f805480602002602001604051908101604052809291908181526020015f905b828210156101ad578382905f5260205f2090600202016040518060400160405290815f82015481526020016001820154815250508152602001906001019061016a565b50505050905092915050565b5f949350505050565b5050505050565b5b5f5f805490501115610210575f8054806101e7576101e66108b8565b5b600190038181905f5260205f2090600202015f5f82015f9055600182015f9055505090556101ca565b5f5f90505b8282905081101561027b575f838383818110610234576102336108e5565b5b905060400201908060018154018082558091505060019003905f5260205f2090600202015f90919091909150818161026c9190610a0d565b50508080600101915050610215565b505050565b5f949350505050565b50505050505050505050565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102c68261029d565b9050919050565b6102d6816102bc565b81146102e0575f5ffd5b50565b5f813590506102f1816102cd565b92915050565b5f5f6040838503121561030d5761030c610295565b5b5f61031a858286016102e3565b925050602061032b858286016102e3565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b6103708161035e565b82525050565b604082015f82015161038a5f850182610367565b50602082015161039d6020850182610367565b50505050565b5f6103ae8383610376565b60408301905092915050565b5f602082019050919050565b5f6103d082610335565b6103da818561033f565b93506103e58361034f565b805f5b838110156104155781516103fc88826103a3565b9750610407836103ba565b9250506001810190506103e8565b5085935050505092915050565b5f6020820190508181035f83015261043a81846103c6565b905092915050565b61044b8161035e565b8114610455575f5ffd5b50565b5f8135905061046681610442565b92915050565b5f5f5f5f6080858703121561048457610483610295565b5b5f610491878288016102e3565b94505060206104a2878288016102e3565b93505060406104b387828801610458565b92505060606104c487828801610458565b91505092959194509250565b6104d9816102bc565b82525050565b5f6020820190506104f25f8301846104d0565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610519576105186104f8565b5b8235905067ffffffffffffffff811115610536576105356104fc565b5b60208301915083602082028301111561055257610551610500565b5b9250929050565b5f5f83601f84011261056e5761056d6104f8565b5b8235905067ffffffffffffffff81111561058b5761058a6104fc565b5b6020830191508360018202830111156105a7576105a6610500565b5b9250929050565b5f5f5f5f5f606086880312156105c7576105c6610295565b5b5f86013567ffffffffffffffff8111156105e4576105e3610299565b5b6105f088828901610504565b9550955050602086013567ffffffffffffffff81111561061357610612610299565b5b61061f88828901610559565b9350935050604061063288828901610458565b9150509295509295909350565b5f5f83601f840112610654576106536104f8565b5b8235905067ffffffffffffffff811115610671576106706104fc565b5b60208301915083604082028301111561068d5761068c610500565b5b9250929050565b5f5f602083850312156106aa576106a9610295565b5b5f83013567ffffffffffffffff8111156106c7576106c6610299565b5b6106d38582860161063f565b92509250509250929050565b5f5ffd5b5f60a082840312156106f8576106f76106df565b5b81905092915050565b5f5f5f5f610100858703121561071a57610719610295565b5b5f610727878288016102e3565b9450506020610738878288016102e3565b935050604061074987828801610458565b925050606061075a878288016106e3565b91505092959194509250565b5f5f83601f84011261077b5761077a6104f8565b5b8235905067ffffffffffffffff811115610798576107976104fc565b5b6020830191508360208202830111156107b4576107b3610500565b5b9250929050565b5f5f5f5f5f5f5f5f5f5f60e08b8d0312156107d9576107d8610295565b5b5f6107e68d828e016102e3565b9a505060206107f78d828e016102e3565b99505060406108088d828e016102e3565b98505060608b013567ffffffffffffffff81111561082957610828610299565b5b6108358d828e01610766565b975097505060808b013567ffffffffffffffff81111561085857610857610299565b5b6108648d828e01610766565b955095505060a08b013567ffffffffffffffff81111561088757610886610299565b5b6108938d828e01610559565b935093505060c06108a68d828e01610458565b9150509295989b9194979a5092959850565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f813561091e81610442565b80915050919050565b5f815f1b9050919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61095d84610927565b9350801983169250808416831791505092915050565b5f819050919050565b5f61099661099161098c8461035e565b610973565b61035e565b9050919050565b5f819050919050565b6109af8261097c565b6109c26109bb8261099d565b8354610932565b8255505050565b5f81015f8301806109d981610912565b90506109e581846109a6565b5050506001810160208301806109fa81610912565b9050610a0681846109a6565b5050505050565b610a1782826109c9565b505056fea2646970667358221220cc199dc71acef720eb72e9aea24c3c5cf1bd979345ba5ab439369f848ab856fd64736f6c634300081e0033",
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

// GetDeposits is a free data retrieval call binding the contract method 0x018fa360.
//
// Solidity: function getDeposits(address user, address token) view returns((uint256,uint256)[])
func (_MockTrustManagementRouter *MockTrustManagementRouterCaller) GetDeposits(opts *bind.CallOpts, user common.Address, token common.Address) ([]MockTrustManagementRouterDeposit, error) {
	var out []interface{}
	err := _MockTrustManagementRouter.contract.Call(opts, &out, "getDeposits", user, token)

	if err != nil {
		return *new([]MockTrustManagementRouterDeposit), err
	}

	out0 := *abi.ConvertType(out[0], new([]MockTrustManagementRouterDeposit)).(*[]MockTrustManagementRouterDeposit)

	return out0, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x018fa360.
//
// Solidity: function getDeposits(address user, address token) view returns((uint256,uint256)[])
func (_MockTrustManagementRouter *MockTrustManagementRouterSession) GetDeposits(user common.Address, token common.Address) ([]MockTrustManagementRouterDeposit, error) {
	return _MockTrustManagementRouter.Contract.GetDeposits(&_MockTrustManagementRouter.CallOpts, user, token)
}

// GetDeposits is a free data retrieval call binding the contract method 0x018fa360.
//
// Solidity: function getDeposits(address user, address token) view returns((uint256,uint256)[])
func (_MockTrustManagementRouter *MockTrustManagementRouterCallerSession) GetDeposits(user common.Address, token common.Address) ([]MockTrustManagementRouterDeposit, error) {
	return _MockTrustManagementRouter.Contract.GetDeposits(&_MockTrustManagementRouter.CallOpts, user, token)
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

// MockSetDeposits is a paid mutator transaction binding the contract method 0x385270b9.
//
// Solidity: function mockSetDeposits((uint256,uint256)[] deposits) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactor) MockSetDeposits(opts *bind.TransactOpts, deposits []MockTrustManagementRouterDeposit) (*types.Transaction, error) {
	return _MockTrustManagementRouter.contract.Transact(opts, "mockSetDeposits", deposits)
}

// MockSetDeposits is a paid mutator transaction binding the contract method 0x385270b9.
//
// Solidity: function mockSetDeposits((uint256,uint256)[] deposits) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterSession) MockSetDeposits(deposits []MockTrustManagementRouterDeposit) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.MockSetDeposits(&_MockTrustManagementRouter.TransactOpts, deposits)
}

// MockSetDeposits is a paid mutator transaction binding the contract method 0x385270b9.
//
// Solidity: function mockSetDeposits((uint256,uint256)[] deposits) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactorSession) MockSetDeposits(deposits []MockTrustManagementRouterDeposit) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.MockSetDeposits(&_MockTrustManagementRouter.TransactOpts, deposits)
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
