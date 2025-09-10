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
	ABI: "[{\"type\":\"function\",\"name\":\"WALLET_BEACON\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositWithPermit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.Permit\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"txs\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Transaction[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Deposit[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWalletAddress\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeployed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initSessionKey\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.ChunkedX509Cert\",\"components\":[{\"name\":\"bodyPartOne\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"bodyPartTwo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"intermediate\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.ChunkedX509Cert\",\"components\":[{\"name\":\"bodyPartOne\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"publicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"bodyPartTwo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"quote\",\"type\":\"tuple\",\"internalType\":\"structMockTrustManagementRouter.ChunkedSGXQuote\",\"components\":[{\"name\":\"header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"isvReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"isvReportSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"attestationKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeReportSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"qeAuthenticationData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockSetDeposits\",\"inputs\":[{\"name\":\"deposits\",\"type\":\"tuple[]\",\"internalType\":\"structMockTrustManagementRouter.Deposit[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockSetWalletAddress\",\"inputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isDeployed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"depositIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amountWithYield\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"lockDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610dc68061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061009c575f3560e01c8063385270b911610064578063385270b9146101695780634c7c04c41461018557806354f3f0b5146101a157806378efbba7146101bd578063c860efa8146101db5761009c565b8063018fa360146100a05780631fd9afa5146100d057806320e8c56514610101578063240aabe3146101315780632f37da741461014d575b5f5ffd5b6100ba60048036038101906100b591906104c7565b61020b565b6040516100c791906105f2565b60405180910390f35b6100ea60048036038101906100e59190610612565b61027d565b6040516100f8929190610666565b60405180910390f35b61011b600480360381019061011691906106b7565b6102b8565b604051610128919061071b565b60405180910390f35b61014b6004803603810190610146919061075e565b61032a565b005b61016760048036038101906101629190610852565b610386565b005b610183600480360381019061017e9190610938565b61038d565b005b61019f600480360381019061019a91906109c3565b610447565b005b6101bb60048036038101906101b69190610ad0565b61044d565b005b6101c5610455565b6040516101d2919061071b565b60405180910390f35b6101f560048036038101906101f09190610b84565b61045c565b604051610202919061071b565b60405180910390f35b60606001805480602002602001604051908101604052809291908181526020015f905b82821015610271578382905f5260205f2090600202016040518060400160405290815f82015481526020016001820154815250508152602001906001019061022e565b50505050905092915050565b5f5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff165f60149054906101000a900460ff1691509150915091565b5f8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fb045190548dadae679cfe9e337437613ca6dd73efdf984f75e56f152ccee22f087868660405161031a93929190610bf8565b60405180910390a3949350505050565b815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550805f60146101000a81548160ff0219169083151502179055505050565b5050505050565b5b5f60018054905011156103d65760018054806103ad576103ac610c2d565b5b600190038181905f5260205f2090600202015f5f82015f9055600182015f90555050905561038e565b5f5f90505b828290508110156104425760018383838181106103fb576103fa610c5a565b5b905060400201908060018154018082558091505060019003905f5260205f2090600202015f9091909190915081816104339190610d82565b505080806001019150506103db565b505050565b50505050565b505050505050565b5f30905090565b5f949350505050565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104968261046d565b9050919050565b6104a68161048c565b81146104b0575f5ffd5b50565b5f813590506104c18161049d565b92915050565b5f5f604083850312156104dd576104dc610465565b5b5f6104ea858286016104b3565b92505060206104fb858286016104b3565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b6105408161052e565b82525050565b604082015f82015161055a5f850182610537565b50602082015161056d6020850182610537565b50505050565b5f61057e8383610546565b60408301905092915050565b5f602082019050919050565b5f6105a082610505565b6105aa818561050f565b93506105b58361051f565b805f5b838110156105e55781516105cc8882610573565b97506105d78361058a565b9250506001810190506105b8565b5085935050505092915050565b5f6020820190508181035f83015261060a8184610596565b905092915050565b5f6020828403121561062757610626610465565b5b5f610634848285016104b3565b91505092915050565b6106468161048c565b82525050565b5f8115159050919050565b6106608161064c565b82525050565b5f6040820190506106795f83018561063d565b6106866020830184610657565b9392505050565b6106968161052e565b81146106a0575f5ffd5b50565b5f813590506106b18161068d565b92915050565b5f5f5f5f608085870312156106cf576106ce610465565b5b5f6106dc878288016104b3565b94505060206106ed878288016104b3565b93505060406106fe878288016106a3565b925050606061070f878288016106a3565b91505092959194509250565b5f60208201905061072e5f83018461063d565b92915050565b61073d8161064c565b8114610747575f5ffd5b50565b5f8135905061075881610734565b92915050565b5f5f6040838503121561077457610773610465565b5b5f610781858286016104b3565b92505060206107928582860161074a565b9150509250929050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126107bd576107bc61079c565b5b8235905067ffffffffffffffff8111156107da576107d96107a0565b5b6020830191508360208202830111156107f6576107f56107a4565b5b9250929050565b5f5f83601f8401126108125761081161079c565b5b8235905067ffffffffffffffff81111561082f5761082e6107a0565b5b60208301915083600182028301111561084b5761084a6107a4565b5b9250929050565b5f5f5f5f5f6060868803121561086b5761086a610465565b5b5f86013567ffffffffffffffff81111561088857610887610469565b5b610894888289016107a8565b9550955050602086013567ffffffffffffffff8111156108b7576108b6610469565b5b6108c3888289016107fd565b935093505060406108d6888289016106a3565b9150509295509295909350565b5f5f83601f8401126108f8576108f761079c565b5b8235905067ffffffffffffffff811115610915576109146107a0565b5b602083019150836040820283011115610931576109306107a4565b5b9250929050565b5f5f6020838503121561094e5761094d610465565b5b5f83013567ffffffffffffffff81111561096b5761096a610469565b5b610977858286016108e3565b92509250509250929050565b5f5ffd5b5f6080828403121561099c5761099b610983565b5b81905092915050565b5f60e082840312156109ba576109b9610983565b5b81905092915050565b5f5f5f5f608085870312156109db576109da610465565b5b5f85013567ffffffffffffffff8111156109f8576109f7610469565b5b610a0487828801610987565b945050602085013567ffffffffffffffff811115610a2557610a24610469565b5b610a3187828801610987565b935050604085013567ffffffffffffffff811115610a5257610a51610469565b5b610a5e878288016109a5565b9250506060610a6f878288016104b3565b91505092959194509250565b5f5f83601f840112610a9057610a8f61079c565b5b8235905067ffffffffffffffff811115610aad57610aac6107a0565b5b602083019150836020820283011115610ac957610ac86107a4565b5b9250929050565b5f5f5f5f5f5f60a08789031215610aea57610ae9610465565b5b5f610af789828a016104b3565b9650506020610b0889828a016104b3565b9550506040610b1989828a016104b3565b945050606087013567ffffffffffffffff811115610b3a57610b39610469565b5b610b4689828a01610a7b565b93509350506080610b5989828a016106a3565b9150509295509295509295565b5f60a08284031215610b7b57610b7a610983565b5b81905092915050565b5f5f5f5f6101008587031215610b9d57610b9c610465565b5b5f610baa878288016104b3565b9450506020610bbb878288016104b3565b9350506040610bcc878288016106a3565b9250506060610bdd87828801610b66565b91505092959194509250565b610bf28161052e565b82525050565b5f606082019050610c0b5f83018661063d565b610c186020830185610be9565b610c256040830184610be9565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f8135610c938161068d565b80915050919050565b5f815f1b9050919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610cd284610c9c565b9350801983169250808416831791505092915050565b5f819050919050565b5f610d0b610d06610d018461052e565b610ce8565b61052e565b9050919050565b5f819050919050565b610d2482610cf1565b610d37610d3082610d12565b8354610ca7565b8255505050565b5f81015f830180610d4e81610c87565b9050610d5a8184610d1b565b505050600181016020830180610d6f81610c87565b9050610d7b8184610d1b565b5050505050565b610d8c8282610d3e565b505056fea26469706673582212208e9835a632635670a8bfa3faceba22b26bde36534b8ee40b605b317c77e65a5564736f6c634300081e0033",
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

// WALLETBEACON is a free data retrieval call binding the contract method 0x78efbba7.
//
// Solidity: function WALLET_BEACON() view returns(address)
func (_MockTrustManagementRouter *MockTrustManagementRouterCaller) WALLETBEACON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockTrustManagementRouter.contract.Call(opts, &out, "WALLET_BEACON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WALLETBEACON is a free data retrieval call binding the contract method 0x78efbba7.
//
// Solidity: function WALLET_BEACON() view returns(address)
func (_MockTrustManagementRouter *MockTrustManagementRouterSession) WALLETBEACON() (common.Address, error) {
	return _MockTrustManagementRouter.Contract.WALLETBEACON(&_MockTrustManagementRouter.CallOpts)
}

// WALLETBEACON is a free data retrieval call binding the contract method 0x78efbba7.
//
// Solidity: function WALLET_BEACON() view returns(address)
func (_MockTrustManagementRouter *MockTrustManagementRouterCallerSession) WALLETBEACON() (common.Address, error) {
	return _MockTrustManagementRouter.Contract.WALLETBEACON(&_MockTrustManagementRouter.CallOpts)
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
