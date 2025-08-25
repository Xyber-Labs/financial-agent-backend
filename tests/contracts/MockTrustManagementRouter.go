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
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositWithPermit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.Permit\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"txs\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Transaction[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Deposit[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initSessionKey\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.ChunkedX509Cert\",\"components\":[{\"name\":\"bodyPartOne\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"bodyPartTwo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"intermediate\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.ChunkedX509Cert\",\"components\":[{\"name\":\"bodyPartOne\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"bodyPartTwo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"quote\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.ChunkedSGXQuote\",\"components\":[{\"name\":\"header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"isvReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"isvReportSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"attestationKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeReportSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeAuthenticationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockSetDeposits\",\"inputs\":[{\"name\":\"deposits\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Deposit[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"depositIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amountsWithYield\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610b828061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061007b575f3560e01c8063385270b911610059578063385270b9146100fb5780634c7c04c414610117578063c860efa814610133578063d3413cac146101635761007b565b8063018fa3601461007f57806320e8c565146100af5780632f37da74146100df575b5f5ffd5b61009960048036038101906100949190610334565b61017f565b6040516100a6919061045f565b60405180910390f35b6100c960048036038101906100c491906104a9565b6101f0565b6040516100d6919061051c565b60405180910390f35b6100f960048036038101906100f491906105eb565b6101f9565b005b610115600480360381019061011091906106d1565b610200565b005b610131600480360381019061012c919061075c565b6102b7565b005b61014d60048036038101906101489190610832565b6102bd565b60405161015a919061051c565b60405180910390f35b61017d600480360381019061017891906108ec565b6102c6565b005b60605f805480602002602001604051908101604052809291908181526020015f905b828210156101e4578382905f5260205f2090600202016040518060400160405290815f8201548152602001600182015481525050815260200190600101906101a1565b50505050905092915050565b5f949350505050565b5050505050565b5b5f5f805490501115610247575f80548061021e5761021d6109e9565b5b600190038181905f5260205f2090600202015f5f82015f9055600182015f905550509055610201565b5f5f90505b828290508110156102b2575f83838381811061026b5761026a610a16565b5b905060400201908060018154018082558091505060019003905f5260205f2090600202015f9091909190915081816102a39190610b3e565b5050808060010191505061024c565b505050565b50505050565b5f949350505050565b50505050505050505050565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610303826102da565b9050919050565b610313816102f9565b811461031d575f5ffd5b50565b5f8135905061032e8161030a565b92915050565b5f5f6040838503121561034a576103496102d2565b5b5f61035785828601610320565b925050602061036885828601610320565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b6103ad8161039b565b82525050565b604082015f8201516103c75f8501826103a4565b5060208201516103da60208501826103a4565b50505050565b5f6103eb83836103b3565b60408301905092915050565b5f602082019050919050565b5f61040d82610372565b610417818561037c565b93506104228361038c565b805f5b8381101561045257815161043988826103e0565b9750610444836103f7565b925050600181019050610425565b5085935050505092915050565b5f6020820190508181035f8301526104778184610403565b905092915050565b6104888161039b565b8114610492575f5ffd5b50565b5f813590506104a38161047f565b92915050565b5f5f5f5f608085870312156104c1576104c06102d2565b5b5f6104ce87828801610320565b94505060206104df87828801610320565b93505060406104f087828801610495565b925050606061050187828801610495565b91505092959194509250565b610516816102f9565b82525050565b5f60208201905061052f5f83018461050d565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261055657610555610535565b5b8235905067ffffffffffffffff81111561057357610572610539565b5b60208301915083602082028301111561058f5761058e61053d565b5b9250929050565b5f5f83601f8401126105ab576105aa610535565b5b8235905067ffffffffffffffff8111156105c8576105c7610539565b5b6020830191508360018202830111156105e4576105e361053d565b5b9250929050565b5f5f5f5f5f60608688031215610604576106036102d2565b5b5f86013567ffffffffffffffff811115610621576106206102d6565b5b61062d88828901610541565b9550955050602086013567ffffffffffffffff8111156106505761064f6102d6565b5b61065c88828901610596565b9350935050604061066f88828901610495565b9150509295509295909350565b5f5f83601f84011261069157610690610535565b5b8235905067ffffffffffffffff8111156106ae576106ad610539565b5b6020830191508360408202830111156106ca576106c961053d565b5b9250929050565b5f5f602083850312156106e7576106e66102d2565b5b5f83013567ffffffffffffffff811115610704576107036102d6565b5b6107108582860161067c565b92509250509250929050565b5f5ffd5b5f608082840312156107355761073461071c565b5b81905092915050565b5f60e082840312156107535761075261071c565b5b81905092915050565b5f5f5f5f60808587031215610774576107736102d2565b5b5f85013567ffffffffffffffff811115610791576107906102d6565b5b61079d87828801610720565b945050602085013567ffffffffffffffff8111156107be576107bd6102d6565b5b6107ca87828801610720565b935050604085013567ffffffffffffffff8111156107eb576107ea6102d6565b5b6107f78782880161073e565b925050606061080887828801610320565b91505092959194509250565b5f60a082840312156108295761082861071c565b5b81905092915050565b5f5f5f5f610100858703121561084b5761084a6102d2565b5b5f61085887828801610320565b945050602061086987828801610320565b935050604061087a87828801610495565b925050606061088b87828801610814565b91505092959194509250565b5f5f83601f8401126108ac576108ab610535565b5b8235905067ffffffffffffffff8111156108c9576108c8610539565b5b6020830191508360208202830111156108e5576108e461053d565b5b9250929050565b5f5f5f5f5f5f5f5f5f5f60e08b8d03121561090a576109096102d2565b5b5f6109178d828e01610320565b9a505060206109288d828e01610320565b99505060406109398d828e01610320565b98505060608b013567ffffffffffffffff81111561095a576109596102d6565b5b6109668d828e01610897565b975097505060808b013567ffffffffffffffff811115610989576109886102d6565b5b6109958d828e01610897565b955095505060a08b013567ffffffffffffffff8111156109b8576109b76102d6565b5b6109c48d828e01610596565b935093505060c06109d78d828e01610495565b9150509295989b9194979a5092959850565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f8135610a4f8161047f565b80915050919050565b5f815f1b9050919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610a8e84610a58565b9350801983169250808416831791505092915050565b5f819050919050565b5f610ac7610ac2610abd8461039b565b610aa4565b61039b565b9050919050565b5f819050919050565b610ae082610aad565b610af3610aec82610ace565b8354610a63565b8255505050565b5f81015f830180610b0a81610a43565b9050610b168184610ad7565b505050600181016020830180610b2b81610a43565b9050610b378184610ad7565b5050505050565b610b488282610afa565b505056fea26469706673582212208917528be325055fc4008417db2ac5c66f1acb7d2216fad8ff84f12790b1ca4d64736f6c634300081e0033",
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
