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
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositWithPermit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.Permit\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"txs\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Transaction[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Deposit[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWalletAddress\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeployed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initSessionKey\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.ChunkedX509Cert\",\"components\":[{\"name\":\"bodyPartOne\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"bodyPartTwo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"intermediate\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.ChunkedX509Cert\",\"components\":[{\"name\":\"bodyPartOne\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"bodyPartTwo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"quote\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.ChunkedSGXQuote\",\"components\":[{\"name\":\"header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"isvReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"isvReportSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"attestationKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeReportSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeAuthenticationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockSetDeposits\",\"inputs\":[{\"name\":\"deposits\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Deposit[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockSetWalletAddress\",\"inputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeployed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"depositIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amountWithYield\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610d968061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610091575f3560e01c80632f37da74116100645780632f37da7414610142578063385270b91461015e5780634c7c04c41461017a57806354f3f0b514610196578063c860efa8146101b257610091565b8063018fa360146100955780631fd9afa5146100c557806320e8c565146100f6578063240aabe314610126575b5f5ffd5b6100af60048036038101906100aa9190610497565b6101e2565b6040516100bc91906105c2565b60405180910390f35b6100df60048036038101906100da91906105e2565b610254565b6040516100ed929190610636565b60405180910390f35b610110600480360381019061010b9190610687565b61028f565b60405161011d91906106eb565b60405180910390f35b610140600480360381019061013b919061072e565b610301565b005b61015c60048036038101906101579190610822565b61035d565b005b61017860048036038101906101739190610908565b610364565b005b610194600480360381019061018f9190610993565b61041e565b005b6101b060048036038101906101ab9190610aa0565b610424565b005b6101cc60048036038101906101c79190610b54565b61042c565b6040516101d991906106eb565b60405180910390f35b60606001805480602002602001604051908101604052809291908181526020015f905b82821015610248578382905f5260205f2090600202016040518060400160405290815f820154815260200160018201548152505081526020019060010190610205565b50505050905092915050565b5f5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff165f60149054906101000a900460ff1691509150915091565b5f8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fb045190548dadae679cfe9e337437613ca6dd73efdf984f75e56f152ccee22f08786866040516102f193929190610bc8565b60405180910390a3949350505050565b815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550805f60146101000a81548160ff0219169083151502179055505050565b5050505050565b5b5f60018054905011156103ad57600180548061038457610383610bfd565b5b600190038181905f5260205f2090600202015f5f82015f9055600182015f905550509055610365565b5f5f90505b828290508110156104195760018383838181106103d2576103d1610c2a565b5b905060400201908060018154018082558091505060019003905f5260205f2090600202015f90919091909150818161040a9190610d52565b505080806001019150506103b2565b505050565b50505050565b505050505050565b5f949350505050565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104668261043d565b9050919050565b6104768161045c565b8114610480575f5ffd5b50565b5f813590506104918161046d565b92915050565b5f5f604083850312156104ad576104ac610435565b5b5f6104ba85828601610483565b92505060206104cb85828601610483565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b610510816104fe565b82525050565b604082015f82015161052a5f850182610507565b50602082015161053d6020850182610507565b50505050565b5f61054e8383610516565b60408301905092915050565b5f602082019050919050565b5f610570826104d5565b61057a81856104df565b9350610585836104ef565b805f5b838110156105b557815161059c8882610543565b97506105a78361055a565b925050600181019050610588565b5085935050505092915050565b5f6020820190508181035f8301526105da8184610566565b905092915050565b5f602082840312156105f7576105f6610435565b5b5f61060484828501610483565b91505092915050565b6106168161045c565b82525050565b5f8115159050919050565b6106308161061c565b82525050565b5f6040820190506106495f83018561060d565b6106566020830184610627565b9392505050565b610666816104fe565b8114610670575f5ffd5b50565b5f813590506106818161065d565b92915050565b5f5f5f5f6080858703121561069f5761069e610435565b5b5f6106ac87828801610483565b94505060206106bd87828801610483565b93505060406106ce87828801610673565b92505060606106df87828801610673565b91505092959194509250565b5f6020820190506106fe5f83018461060d565b92915050565b61070d8161061c565b8114610717575f5ffd5b50565b5f8135905061072881610704565b92915050565b5f5f6040838503121561074457610743610435565b5b5f61075185828601610483565b92505060206107628582860161071a565b9150509250929050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261078d5761078c61076c565b5b8235905067ffffffffffffffff8111156107aa576107a9610770565b5b6020830191508360208202830111156107c6576107c5610774565b5b9250929050565b5f5f83601f8401126107e2576107e161076c565b5b8235905067ffffffffffffffff8111156107ff576107fe610770565b5b60208301915083600182028301111561081b5761081a610774565b5b9250929050565b5f5f5f5f5f6060868803121561083b5761083a610435565b5b5f86013567ffffffffffffffff81111561085857610857610439565b5b61086488828901610778565b9550955050602086013567ffffffffffffffff81111561088757610886610439565b5b610893888289016107cd565b935093505060406108a688828901610673565b9150509295509295909350565b5f5f83601f8401126108c8576108c761076c565b5b8235905067ffffffffffffffff8111156108e5576108e4610770565b5b60208301915083604082028301111561090157610900610774565b5b9250929050565b5f5f6020838503121561091e5761091d610435565b5b5f83013567ffffffffffffffff81111561093b5761093a610439565b5b610947858286016108b3565b92509250509250929050565b5f5ffd5b5f6080828403121561096c5761096b610953565b5b81905092915050565b5f60e0828403121561098a57610989610953565b5b81905092915050565b5f5f5f5f608085870312156109ab576109aa610435565b5b5f85013567ffffffffffffffff8111156109c8576109c7610439565b5b6109d487828801610957565b945050602085013567ffffffffffffffff8111156109f5576109f4610439565b5b610a0187828801610957565b935050604085013567ffffffffffffffff811115610a2257610a21610439565b5b610a2e87828801610975565b9250506060610a3f87828801610483565b91505092959194509250565b5f5f83601f840112610a6057610a5f61076c565b5b8235905067ffffffffffffffff811115610a7d57610a7c610770565b5b602083019150836020820283011115610a9957610a98610774565b5b9250929050565b5f5f5f5f5f5f60a08789031215610aba57610ab9610435565b5b5f610ac789828a01610483565b9650506020610ad889828a01610483565b9550506040610ae989828a01610483565b945050606087013567ffffffffffffffff811115610b0a57610b09610439565b5b610b1689828a01610a4b565b93509350506080610b2989828a01610673565b9150509295509295509295565b5f60a08284031215610b4b57610b4a610953565b5b81905092915050565b5f5f5f5f6101008587031215610b6d57610b6c610435565b5b5f610b7a87828801610483565b9450506020610b8b87828801610483565b9350506040610b9c87828801610673565b9250506060610bad87828801610b36565b91505092959194509250565b610bc2816104fe565b82525050565b5f606082019050610bdb5f83018661060d565b610be86020830185610bb9565b610bf56040830184610bb9565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f8135610c638161065d565b80915050919050565b5f815f1b9050919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610ca284610c6c565b9350801983169250808416831791505092915050565b5f819050919050565b5f610cdb610cd6610cd1846104fe565b610cb8565b6104fe565b9050919050565b5f819050919050565b610cf482610cc1565b610d07610d0082610ce2565b8354610c77565b8255505050565b5f81015f830180610d1e81610c57565b9050610d2a8184610ceb565b505050600181016020830180610d3f81610c57565b9050610d4b8184610ceb565b5050505050565b610d5c8282610d0e565b505056fea26469706673582212200581a493fcf45ae44f2b506f92d2258c6774cf3922f3e851f7acf25aeaafb08a64736f6c634300081e0033",
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

// GetWalletAddress is a free data retrieval call binding the contract method 0x1fd9afa5.
//
// Solidity: function getWalletAddress(address user) view returns(address walletAddress, bool isDeployed)
func (_MockTrustManagementRouter *MockTrustManagementRouterCaller) GetWalletAddress(opts *bind.CallOpts, user common.Address) (struct {
	WalletAddress common.Address
	IsDeployed    bool
}, error) {
	var out []interface{}
	err := _MockTrustManagementRouter.contract.Call(opts, &out, "getWalletAddress", user)

	outstruct := new(struct {
		WalletAddress common.Address
		IsDeployed    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WalletAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsDeployed = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetWalletAddress is a free data retrieval call binding the contract method 0x1fd9afa5.
//
// Solidity: function getWalletAddress(address user) view returns(address walletAddress, bool isDeployed)
func (_MockTrustManagementRouter *MockTrustManagementRouterSession) GetWalletAddress(user common.Address) (struct {
	WalletAddress common.Address
	IsDeployed    bool
}, error) {
	return _MockTrustManagementRouter.Contract.GetWalletAddress(&_MockTrustManagementRouter.CallOpts, user)
}

// GetWalletAddress is a free data retrieval call binding the contract method 0x1fd9afa5.
//
// Solidity: function getWalletAddress(address user) view returns(address walletAddress, bool isDeployed)
func (_MockTrustManagementRouter *MockTrustManagementRouterCallerSession) GetWalletAddress(user common.Address) (struct {
	WalletAddress common.Address
	IsDeployed    bool
}, error) {
	return _MockTrustManagementRouter.Contract.GetWalletAddress(&_MockTrustManagementRouter.CallOpts, user)
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

// MockSetWalletAddress is a paid mutator transaction binding the contract method 0x240aabe3.
//
// Solidity: function mockSetWalletAddress(address walletAddress, bool isDeployed) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactor) MockSetWalletAddress(opts *bind.TransactOpts, walletAddress common.Address, isDeployed bool) (*types.Transaction, error) {
	return _MockTrustManagementRouter.contract.Transact(opts, "mockSetWalletAddress", walletAddress, isDeployed)
}

// MockSetWalletAddress is a paid mutator transaction binding the contract method 0x240aabe3.
//
// Solidity: function mockSetWalletAddress(address walletAddress, bool isDeployed) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterSession) MockSetWalletAddress(walletAddress common.Address, isDeployed bool) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.MockSetWalletAddress(&_MockTrustManagementRouter.TransactOpts, walletAddress, isDeployed)
}

// MockSetWalletAddress is a paid mutator transaction binding the contract method 0x240aabe3.
//
// Solidity: function mockSetWalletAddress(address walletAddress, bool isDeployed) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactorSession) MockSetWalletAddress(walletAddress common.Address, isDeployed bool) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.MockSetWalletAddress(&_MockTrustManagementRouter.TransactOpts, walletAddress, isDeployed)
}

// Withdraw is a paid mutator transaction binding the contract method 0x54f3f0b5.
//
// Solidity: function withdraw(address user, address token, address receiver, uint256[] depositIds, uint256 amountWithYield) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactor) Withdraw(opts *bind.TransactOpts, user common.Address, token common.Address, receiver common.Address, depositIds []*big.Int, amountWithYield *big.Int) (*types.Transaction, error) {
	return _MockTrustManagementRouter.contract.Transact(opts, "withdraw", user, token, receiver, depositIds, amountWithYield)
}

// Withdraw is a paid mutator transaction binding the contract method 0x54f3f0b5.
//
// Solidity: function withdraw(address user, address token, address receiver, uint256[] depositIds, uint256 amountWithYield) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterSession) Withdraw(user common.Address, token common.Address, receiver common.Address, depositIds []*big.Int, amountWithYield *big.Int) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.Withdraw(&_MockTrustManagementRouter.TransactOpts, user, token, receiver, depositIds, amountWithYield)
}

// Withdraw is a paid mutator transaction binding the contract method 0x54f3f0b5.
//
// Solidity: function withdraw(address user, address token, address receiver, uint256[] depositIds, uint256 amountWithYield) returns()
func (_MockTrustManagementRouter *MockTrustManagementRouterTransactorSession) Withdraw(user common.Address, token common.Address, receiver common.Address, depositIds []*big.Int, amountWithYield *big.Int) (*types.Transaction, error) {
	return _MockTrustManagementRouter.Contract.Withdraw(&_MockTrustManagementRouter.TransactOpts, user, token, receiver, depositIds, amountWithYield)
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
