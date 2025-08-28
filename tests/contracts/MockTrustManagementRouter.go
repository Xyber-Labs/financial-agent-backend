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

// MockTrustManagementRouterChunkedSGXQuote is an auto generated low-level Go binding around an user-defined struct.
type MockTrustManagementRouterChunkedSGXQuote struct {
	Header               []byte
	IsvReport            []byte
	IsvReportSignature   []byte
	AttestationKey       []byte
	QeReport             []byte
	QeReportSignature    []byte
	QeAuthenticationData []byte
}

// MockTrustManagementRouterChunkedX509Cert is an auto generated low-level Go binding around an user-defined struct.
type MockTrustManagementRouterChunkedX509Cert struct {
	BodyPartOne []byte
	PublicKey   []byte
	BodyPartTwo []byte
	Signature   []byte
}

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
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositWithPermit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.Permit\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"txs\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Transaction[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Deposit[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initSessionKey\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.ChunkedX509Cert\",\"components\":[{\"name\":\"bodyPartOne\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"bodyPartTwo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"intermediate\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.ChunkedX509Cert\",\"components\":[{\"name\":\"bodyPartOne\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"bodyPartTwo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"quote\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.ChunkedSGXQuote\",\"components\":[{\"name\":\"header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"isvReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"isvReportSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"attestationKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeReportSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeAuthenticationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockSetDeposits\",\"inputs\":[{\"name\":\"deposits\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Deposit[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"depositIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amountsWithYield\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610c2f8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061007b575f3560e01c8063385270b911610059578063385270b9146100fb5780634c7c04c414610117578063c860efa814610133578063d3413cac146101635761007b565b8063018fa3601461007f57806320e8c565146100af5780632f37da74146100df575b5f5ffd5b6100996004803603810190610094919061039d565b61017f565b6040516100a691906104c8565b60405180910390f35b6100c960048036038101906100c49190610512565b6101f0565b6040516100d69190610585565b60405180910390f35b6100f960048036038101906100f49190610654565b610262565b005b6101156004803603810190610110919061073a565b610269565b005b610131600480360381019061012c91906107c5565b610320565b005b61014d6004803603810190610148919061089b565b610326565b60405161015a9190610585565b60405180910390f35b61017d60048036038101906101789190610955565b61032f565b005b60605f805480602002602001604051908101604052809291908181526020015f905b828210156101e4578382905f5260205f2090600202016040518060400160405290815f8201548152602001600182015481525050815260200190600101906101a1565b50505050905092915050565b5f8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fb045190548dadae679cfe9e337437613ca6dd73efdf984f75e56f152ccee22f087868660405161025293929190610a61565b60405180910390a3949350505050565b5050505050565b5b5f5f8054905011156102b0575f80548061028757610286610a96565b5b600190038181905f5260205f2090600202015f5f82015f9055600182015f90555050905561026a565b5f5f90505b8282905081101561031b575f8383838181106102d4576102d3610ac3565b5b905060400201908060018154018082558091505060019003905f5260205f2090600202015f90919091909150818161030c9190610beb565b505080806001019150506102b5565b505050565b50505050565b5f949350505050565b50505050505050505050565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61036c82610343565b9050919050565b61037c81610362565b8114610386575f5ffd5b50565b5f8135905061039781610373565b92915050565b5f5f604083850312156103b3576103b261033b565b5b5f6103c085828601610389565b92505060206103d185828601610389565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b61041681610404565b82525050565b604082015f8201516104305f85018261040d565b506020820151610443602085018261040d565b50505050565b5f610454838361041c565b60408301905092915050565b5f602082019050919050565b5f610476826103db565b61048081856103e5565b935061048b836103f5565b805f5b838110156104bb5781516104a28882610449565b97506104ad83610460565b92505060018101905061048e565b5085935050505092915050565b5f6020820190508181035f8301526104e0818461046c565b905092915050565b6104f181610404565b81146104fb575f5ffd5b50565b5f8135905061050c816104e8565b92915050565b5f5f5f5f6080858703121561052a5761052961033b565b5b5f61053787828801610389565b945050602061054887828801610389565b9350506040610559878288016104fe565b925050606061056a878288016104fe565b91505092959194509250565b61057f81610362565b82525050565b5f6020820190506105985f830184610576565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126105bf576105be61059e565b5b8235905067ffffffffffffffff8111156105dc576105db6105a2565b5b6020830191508360208202830111156105f8576105f76105a6565b5b9250929050565b5f5f83601f8401126106145761061361059e565b5b8235905067ffffffffffffffff811115610631576106306105a2565b5b60208301915083600182028301111561064d5761064c6105a6565b5b9250929050565b5f5f5f5f5f6060868803121561066d5761066c61033b565b5b5f86013567ffffffffffffffff81111561068a5761068961033f565b5b610696888289016105aa565b9550955050602086013567ffffffffffffffff8111156106b9576106b861033f565b5b6106c5888289016105ff565b935093505060406106d8888289016104fe565b9150509295509295909350565b5f5f83601f8401126106fa576106f961059e565b5b8235905067ffffffffffffffff811115610717576107166105a2565b5b602083019150836040820283011115610733576107326105a6565b5b9250929050565b5f5f602083850312156107505761074f61033b565b5b5f83013567ffffffffffffffff81111561076d5761076c61033f565b5b610779858286016106e5565b92509250509250929050565b5f5ffd5b5f6080828403121561079e5761079d610785565b5b81905092915050565b5f60e082840312156107bc576107bb610785565b5b81905092915050565b5f5f5f5f608085870312156107dd576107dc61033b565b5b5f85013567ffffffffffffffff8111156107fa576107f961033f565b5b61080687828801610789565b945050602085013567ffffffffffffffff8111156108275761082661033f565b5b61083387828801610789565b935050604085013567ffffffffffffffff8111156108545761085361033f565b5b610860878288016107a7565b925050606061087187828801610389565b91505092959194509250565b5f60a0828403121561089257610891610785565b5b81905092915050565b5f5f5f5f61010085870312156108b4576108b361033b565b5b5f6108c187828801610389565b94505060206108d287828801610389565b93505060406108e3878288016104fe565b92505060606108f48782880161087d565b91505092959194509250565b5f5f83601f8401126109155761091461059e565b5b8235905067ffffffffffffffff811115610932576109316105a2565b5b60208301915083602082028301111561094e5761094d6105a6565b5b9250929050565b5f5f5f5f5f5f5f5f5f5f60e08b8d0312156109735761097261033b565b5b5f6109808d828e01610389565b9a505060206109918d828e01610389565b99505060406109a28d828e01610389565b98505060608b013567ffffffffffffffff8111156109c3576109c261033f565b5b6109cf8d828e01610900565b975097505060808b013567ffffffffffffffff8111156109f2576109f161033f565b5b6109fe8d828e01610900565b955095505060a08b013567ffffffffffffffff811115610a2157610a2061033f565b5b610a2d8d828e016105ff565b935093505060c0610a408d828e016104fe565b9150509295989b9194979a5092959850565b610a5b81610404565b82525050565b5f606082019050610a745f830186610576565b610a816020830185610a52565b610a8e6040830184610a52565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f8135610afc816104e8565b80915050919050565b5f815f1b9050919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610b3b84610b05565b9350801983169250808416831791505092915050565b5f819050919050565b5f610b74610b6f610b6a84610404565b610b51565b610404565b9050919050565b5f819050919050565b610b8d82610b5a565b610ba0610b9982610b7b565b8354610b10565b8255505050565b5f81015f830180610bb781610af0565b9050610bc38184610b84565b505050600181016020830180610bd881610af0565b9050610be48184610b84565b5050505050565b610bf58282610ba7565b505056fea264697066735822122023f2f1caa8a0bf1a16f6fb9e51d77e1b78600a255f2551b124c82eeaed378be764736f6c634300081e0033",
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

// InitSessionKey is a paid mutator transaction binding the contract method 0x4c7c04c4.
//
// Solidity: function initSessionKey((bytes,bytes,bytes,bytes) leaf, (bytes,bytes,bytes,bytes) intermediate, (bytes,bytes,bytes,bytes,bytes,bytes,bytes) quote, address sessionKey) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactor) InitSessionKey(opts *bind.TransactOpts, leaf MockTrustManagementRouterChunkedX509Cert, intermediate MockTrustManagementRouterChunkedX509Cert, quote MockTrustManagementRouterChunkedSGXQuote, sessionKey common.Address) (*types.Transaction, error) {
	return _MockTrustManagementRouter.contract.Transact(opts, "initSessionKey", leaf, intermediate, quote, sessionKey)
}

// InitSessionKey is a paid mutator transaction binding the contract method 0x4c7c04c4.
//
// Solidity: function initSessionKey((bytes,bytes,bytes,bytes) leaf, (bytes,bytes,bytes,bytes) intermediate, (bytes,bytes,bytes,bytes,bytes,bytes,bytes) quote, address sessionKey) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterSession) InitSessionKey(leaf MockTrustManagementRouterChunkedX509Cert, intermediate MockTrustManagementRouterChunkedX509Cert, quote MockTrustManagementRouterChunkedSGXQuote, sessionKey common.Address) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.InitSessionKey(&_MockTrustManagementRouter.TransactOpts, leaf, intermediate, quote, sessionKey)
}

// InitSessionKey is a paid mutator transaction binding the contract method 0x4c7c04c4.
//
// Solidity: function initSessionKey((bytes,bytes,bytes,bytes) leaf, (bytes,bytes,bytes,bytes) intermediate, (bytes,bytes,bytes,bytes,bytes,bytes,bytes) quote, address sessionKey) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactorSession) InitSessionKey(leaf MockTrustManagementRouterChunkedX509Cert, intermediate MockTrustManagementRouterChunkedX509Cert, quote MockTrustManagementRouterChunkedSGXQuote, sessionKey common.Address) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.InitSessionKey(&_MockTrustManagementRouter.TransactOpts, leaf, intermediate, quote, sessionKey)
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

// MockTrustManagementRouterDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the MockTrustManagementRouter contract.
type MockTrustManagementRouterDepositedIterator struct {
	Event *MockTrustManagementRouterDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockTrustManagementRouterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockTrustManagementRouterDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockTrustManagementRouterDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockTrustManagementRouterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockTrustManagementRouterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockTrustManagementRouterDeposited represents a Deposited event raised by the MockTrustManagementRouter contract.
type MockTrustManagementRouterDeposited struct {
	Wallet       common.Address
	User         common.Address
	Token        common.Address
	Amount       *big.Int
	LockDuration *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xb045190548dadae679cfe9e337437613ca6dd73efdf984f75e56f152ccee22f0.
//
// Solidity: event Deposited(address wallet, address indexed user, address indexed token, uint256 amount, uint256 lockDuration)
func (_MockTrustManagementRouter *MockTrustManagementRouterFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*MockTrustManagementRouterDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MockTrustManagementRouter.contract.FilterLogs(opts, "Deposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MockTrustManagementRouterDepositedIterator{contract: _MockTrustManagementRouter.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xb045190548dadae679cfe9e337437613ca6dd73efdf984f75e56f152ccee22f0.
//
// Solidity: event Deposited(address wallet, address indexed user, address indexed token, uint256 amount, uint256 lockDuration)
func (_MockTrustManagementRouter *MockTrustManagementRouterFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *MockTrustManagementRouterDeposited, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MockTrustManagementRouter.contract.WatchLogs(opts, "Deposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockTrustManagementRouterDeposited)
				if err := _MockTrustManagementRouter.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0xb045190548dadae679cfe9e337437613ca6dd73efdf984f75e56f152ccee22f0.
//
// Solidity: event Deposited(address wallet, address indexed user, address indexed token, uint256 amount, uint256 lockDuration)
func (_MockTrustManagementRouter *MockTrustManagementRouterFilterer) ParseDeposited(log types.Log) (*MockTrustManagementRouterDeposited, error) {
	event := new(MockTrustManagementRouterDeposited)
	if err := _MockTrustManagementRouter.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
