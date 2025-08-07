// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TrustManagementRouter

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

// ITrustManagementStructsPermit is an auto generated low-level Go binding around an user-defined struct.
type ITrustManagementStructsPermit struct {
	Amount   *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

// ITrustManagementStructsTransaction is an auto generated low-level Go binding around an user-defined struct.
type ITrustManagementStructsTransaction struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
}

// TrustManagementRouterMetaData contains all meta data concerning the TrustManagementRouter contract.
var TrustManagementRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"TrustManagementRouter__CallerDoesntHaveAgentRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustManagementRouter__CallerIsNotAgent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"TrustManagementRouter__CallerIsNotUser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TrustManagementRouter__CallerIsNotUserOrWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"TrustManagementRouter__InvalidWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callResponse\",\"type\":\"bytes\"}],\"name\":\"TrustManagementRouter__MsgValueSendingFail\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"providedAddress\",\"type\":\"address\"}],\"name\":\"TrustManagementRouter__ProvidedAddressDoesntHaveAgentRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TrustManagementRouter__TokenNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockTimestamp\",\"type\":\"uint256\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGENT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"_depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_amountWithYield\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"_depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountWithYield\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"createWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newWallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_lockDuration\",\"type\":\"uint256\"}],\"name\":\"createWalletAndDeposit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newWallet\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockDuration\",\"type\":\"uint256\"}],\"name\":\"createWalletAndDeposit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockTimestamp\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structITrustManagementStructs.Permit\",\"name\":\"_permit\",\"type\":\"tuple\"}],\"name\":\"createWalletAndDepositWithPermit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlockTimestamp\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlockTimestamp\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockTimestamp\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockTimestamp\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structITrustManagementStructs.Permit\",\"name\":\"_permit\",\"type\":\"tuple\"}],\"name\":\"depositWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"_depositsIds\",\"type\":\"uint256[][]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structITrustManagementStructs.Transaction[]\",\"name\":\"txs\",\"type\":\"tuple[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPerfomanceFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"perfomanceFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initBlueprint\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeReceiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPerfomanceFeeRate\",\"type\":\"uint256\"}],\"name\":\"setPerfomanceFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amountWithYield\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TrustManagementRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use TrustManagementRouterMetaData.ABI instead.
var TrustManagementRouterABI = TrustManagementRouterMetaData.ABI

// TrustManagementRouter is an auto generated Go binding around an Ethereum contract.
type TrustManagementRouter struct {
	TrustManagementRouterCaller     // Read-only binding to the contract
	TrustManagementRouterTransactor // Write-only binding to the contract
	TrustManagementRouterFilterer   // Log filterer for contract events
}

// TrustManagementRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrustManagementRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustManagementRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrustManagementRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustManagementRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrustManagementRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustManagementRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrustManagementRouterSession struct {
	Contract     *TrustManagementRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TrustManagementRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrustManagementRouterCallerSession struct {
	Contract *TrustManagementRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TrustManagementRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrustManagementRouterTransactorSession struct {
	Contract     *TrustManagementRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TrustManagementRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrustManagementRouterRaw struct {
	Contract *TrustManagementRouter // Generic contract binding to access the raw methods on
}

// TrustManagementRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrustManagementRouterCallerRaw struct {
	Contract *TrustManagementRouterCaller // Generic read-only contract binding to access the raw methods on
}

// TrustManagementRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrustManagementRouterTransactorRaw struct {
	Contract *TrustManagementRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrustManagementRouter creates a new instance of TrustManagementRouter, bound to a specific deployed contract.
func NewTrustManagementRouter(address common.Address, backend bind.ContractBackend) (*TrustManagementRouter, error) {
	contract, err := bindTrustManagementRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouter{TrustManagementRouterCaller: TrustManagementRouterCaller{contract: contract}, TrustManagementRouterTransactor: TrustManagementRouterTransactor{contract: contract}, TrustManagementRouterFilterer: TrustManagementRouterFilterer{contract: contract}}, nil
}

// NewTrustManagementRouterCaller creates a new read-only instance of TrustManagementRouter, bound to a specific deployed contract.
func NewTrustManagementRouterCaller(address common.Address, caller bind.ContractCaller) (*TrustManagementRouterCaller, error) {
	contract, err := bindTrustManagementRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterCaller{contract: contract}, nil
}

// NewTrustManagementRouterTransactor creates a new write-only instance of TrustManagementRouter, bound to a specific deployed contract.
func NewTrustManagementRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*TrustManagementRouterTransactor, error) {
	contract, err := bindTrustManagementRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterTransactor{contract: contract}, nil
}

// NewTrustManagementRouterFilterer creates a new log filterer instance of TrustManagementRouter, bound to a specific deployed contract.
func NewTrustManagementRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*TrustManagementRouterFilterer, error) {
	contract, err := bindTrustManagementRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterFilterer{contract: contract}, nil
}

// bindTrustManagementRouter binds a generic wrapper to an already deployed contract.
func bindTrustManagementRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TrustManagementRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustManagementRouter *TrustManagementRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustManagementRouter.Contract.TrustManagementRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustManagementRouter *TrustManagementRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.TrustManagementRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustManagementRouter *TrustManagementRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.TrustManagementRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustManagementRouter *TrustManagementRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustManagementRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustManagementRouter *TrustManagementRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustManagementRouter *TrustManagementRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.contract.Transact(opts, method, params...)
}

// AGENTROLE is a free data retrieval call binding the contract method 0x22459e18.
//
// Solidity: function AGENT_ROLE() view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterCaller) AGENTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "AGENT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AGENTROLE is a free data retrieval call binding the contract method 0x22459e18.
//
// Solidity: function AGENT_ROLE() view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterSession) AGENTROLE() ([32]byte, error) {
	return _TrustManagementRouter.Contract.AGENTROLE(&_TrustManagementRouter.CallOpts)
}

// AGENTROLE is a free data retrieval call binding the contract method 0x22459e18.
//
// Solidity: function AGENT_ROLE() view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) AGENTROLE() ([32]byte, error) {
	return _TrustManagementRouter.Contract.AGENTROLE(&_TrustManagementRouter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TrustManagementRouter.Contract.DEFAULTADMINROLE(&_TrustManagementRouter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TrustManagementRouter.Contract.DEFAULTADMINROLE(&_TrustManagementRouter.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TrustManagementRouter *TrustManagementRouterCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TrustManagementRouter *TrustManagementRouterSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TrustManagementRouter.Contract.UPGRADEINTERFACEVERSION(&_TrustManagementRouter.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TrustManagementRouter.Contract.UPGRADEINTERFACEVERSION(&_TrustManagementRouter.CallOpts)
}

// BeaconAddress is a free data retrieval call binding the contract method 0x7e2ec6d0.
//
// Solidity: function beaconAddress() view returns(address)
func (_TrustManagementRouter *TrustManagementRouterCaller) BeaconAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "beaconAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconAddress is a free data retrieval call binding the contract method 0x7e2ec6d0.
//
// Solidity: function beaconAddress() view returns(address)
func (_TrustManagementRouter *TrustManagementRouterSession) BeaconAddress() (common.Address, error) {
	return _TrustManagementRouter.Contract.BeaconAddress(&_TrustManagementRouter.CallOpts)
}

// BeaconAddress is a free data retrieval call binding the contract method 0x7e2ec6d0.
//
// Solidity: function beaconAddress() view returns(address)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) BeaconAddress() (common.Address, error) {
	return _TrustManagementRouter.Contract.BeaconAddress(&_TrustManagementRouter.CallOpts)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0xe8a35392.
//
// Solidity: function getFeeReceiver() view returns(address feeReceiver)
func (_TrustManagementRouter *TrustManagementRouterCaller) GetFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "getFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeReceiver is a free data retrieval call binding the contract method 0xe8a35392.
//
// Solidity: function getFeeReceiver() view returns(address feeReceiver)
func (_TrustManagementRouter *TrustManagementRouterSession) GetFeeReceiver() (common.Address, error) {
	return _TrustManagementRouter.Contract.GetFeeReceiver(&_TrustManagementRouter.CallOpts)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0xe8a35392.
//
// Solidity: function getFeeReceiver() view returns(address feeReceiver)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) GetFeeReceiver() (common.Address, error) {
	return _TrustManagementRouter.Contract.GetFeeReceiver(&_TrustManagementRouter.CallOpts)
}

// GetPerfomanceFeeRate is a free data retrieval call binding the contract method 0x832675e4.
//
// Solidity: function getPerfomanceFeeRate() view returns(uint256 perfomanceFeeRate)
func (_TrustManagementRouter *TrustManagementRouterCaller) GetPerfomanceFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "getPerfomanceFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPerfomanceFeeRate is a free data retrieval call binding the contract method 0x832675e4.
//
// Solidity: function getPerfomanceFeeRate() view returns(uint256 perfomanceFeeRate)
func (_TrustManagementRouter *TrustManagementRouterSession) GetPerfomanceFeeRate() (*big.Int, error) {
	return _TrustManagementRouter.Contract.GetPerfomanceFeeRate(&_TrustManagementRouter.CallOpts)
}

// GetPerfomanceFeeRate is a free data retrieval call binding the contract method 0x832675e4.
//
// Solidity: function getPerfomanceFeeRate() view returns(uint256 perfomanceFeeRate)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) GetPerfomanceFeeRate() (*big.Int, error) {
	return _TrustManagementRouter.Contract.GetPerfomanceFeeRate(&_TrustManagementRouter.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TrustManagementRouter.Contract.GetRoleAdmin(&_TrustManagementRouter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TrustManagementRouter.Contract.GetRoleAdmin(&_TrustManagementRouter.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TrustManagementRouter *TrustManagementRouterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TrustManagementRouter *TrustManagementRouterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TrustManagementRouter.Contract.HasRole(&_TrustManagementRouter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TrustManagementRouter.Contract.HasRole(&_TrustManagementRouter.CallOpts, role, account)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterSession) ProxiableUUID() ([32]byte, error) {
	return _TrustManagementRouter.Contract.ProxiableUUID(&_TrustManagementRouter.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TrustManagementRouter.Contract.ProxiableUUID(&_TrustManagementRouter.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TrustManagementRouter *TrustManagementRouterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TrustManagementRouter *TrustManagementRouterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TrustManagementRouter.Contract.SupportsInterface(&_TrustManagementRouter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TrustManagementRouter.Contract.SupportsInterface(&_TrustManagementRouter.CallOpts, interfaceId)
}

// Claim is a paid mutator transaction binding the contract method 0x78bffdec.
//
// Solidity: function claim(address _wallet, bytes _signature, uint256[] _depositIds, uint256 _amountWithYield, address _receiver) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) Claim(opts *bind.TransactOpts, _wallet common.Address, _signature []byte, _depositIds []*big.Int, _amountWithYield *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "claim", _wallet, _signature, _depositIds, _amountWithYield, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x78bffdec.
//
// Solidity: function claim(address _wallet, bytes _signature, uint256[] _depositIds, uint256 _amountWithYield, address _receiver) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) Claim(_wallet common.Address, _signature []byte, _depositIds []*big.Int, _amountWithYield *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Claim(&_TrustManagementRouter.TransactOpts, _wallet, _signature, _depositIds, _amountWithYield, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x78bffdec.
//
// Solidity: function claim(address _wallet, bytes _signature, uint256[] _depositIds, uint256 _amountWithYield, address _receiver) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Claim(_wallet common.Address, _signature []byte, _depositIds []*big.Int, _amountWithYield *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Claim(&_TrustManagementRouter.TransactOpts, _wallet, _signature, _depositIds, _amountWithYield, _receiver)
}

// Claim0 is a paid mutator transaction binding the contract method 0x7bd4f668.
//
// Solidity: function claim(address _wallet, bytes _signature, uint256[] _depositIds, address _token, uint256 _amountWithYield, address _receiver) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) Claim0(opts *bind.TransactOpts, _wallet common.Address, _signature []byte, _depositIds []*big.Int, _token common.Address, _amountWithYield *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "claim0", _wallet, _signature, _depositIds, _token, _amountWithYield, _receiver)
}

// Claim0 is a paid mutator transaction binding the contract method 0x7bd4f668.
//
// Solidity: function claim(address _wallet, bytes _signature, uint256[] _depositIds, address _token, uint256 _amountWithYield, address _receiver) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) Claim0(_wallet common.Address, _signature []byte, _depositIds []*big.Int, _token common.Address, _amountWithYield *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Claim0(&_TrustManagementRouter.TransactOpts, _wallet, _signature, _depositIds, _token, _amountWithYield, _receiver)
}

// Claim0 is a paid mutator transaction binding the contract method 0x7bd4f668.
//
// Solidity: function claim(address _wallet, bytes _signature, uint256[] _depositIds, address _token, uint256 _amountWithYield, address _receiver) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Claim0(_wallet common.Address, _signature []byte, _depositIds []*big.Int, _token common.Address, _amountWithYield *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Claim0(&_TrustManagementRouter.TransactOpts, _wallet, _signature, _depositIds, _token, _amountWithYield, _receiver)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xe8eef270.
//
// Solidity: function createWallet(address _user, address _agent, bytes32 _salt) returns(address newWallet)
func (_TrustManagementRouter *TrustManagementRouterTransactor) CreateWallet(opts *bind.TransactOpts, _user common.Address, _agent common.Address, _salt [32]byte) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "createWallet", _user, _agent, _salt)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xe8eef270.
//
// Solidity: function createWallet(address _user, address _agent, bytes32 _salt) returns(address newWallet)
func (_TrustManagementRouter *TrustManagementRouterSession) CreateWallet(_user common.Address, _agent common.Address, _salt [32]byte) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.CreateWallet(&_TrustManagementRouter.TransactOpts, _user, _agent, _salt)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xe8eef270.
//
// Solidity: function createWallet(address _user, address _agent, bytes32 _salt) returns(address newWallet)
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) CreateWallet(_user common.Address, _agent common.Address, _salt [32]byte) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.CreateWallet(&_TrustManagementRouter.TransactOpts, _user, _agent, _salt)
}

// CreateWalletAndDeposit is a paid mutator transaction binding the contract method 0x22a8adac.
//
// Solidity: function createWalletAndDeposit(address _user, address _agent, bytes32 _salt, uint256 _lockDuration) payable returns(address newWallet)
func (_TrustManagementRouter *TrustManagementRouterTransactor) CreateWalletAndDeposit(opts *bind.TransactOpts, _user common.Address, _agent common.Address, _salt [32]byte, _lockDuration *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "createWalletAndDeposit", _user, _agent, _salt, _lockDuration)
}

// CreateWalletAndDeposit is a paid mutator transaction binding the contract method 0x22a8adac.
//
// Solidity: function createWalletAndDeposit(address _user, address _agent, bytes32 _salt, uint256 _lockDuration) payable returns(address newWallet)
func (_TrustManagementRouter *TrustManagementRouterSession) CreateWalletAndDeposit(_user common.Address, _agent common.Address, _salt [32]byte, _lockDuration *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.CreateWalletAndDeposit(&_TrustManagementRouter.TransactOpts, _user, _agent, _salt, _lockDuration)
}

// CreateWalletAndDeposit is a paid mutator transaction binding the contract method 0x22a8adac.
//
// Solidity: function createWalletAndDeposit(address _user, address _agent, bytes32 _salt, uint256 _lockDuration) payable returns(address newWallet)
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) CreateWalletAndDeposit(_user common.Address, _agent common.Address, _salt [32]byte, _lockDuration *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.CreateWalletAndDeposit(&_TrustManagementRouter.TransactOpts, _user, _agent, _salt, _lockDuration)
}

// CreateWalletAndDeposit0 is a paid mutator transaction binding the contract method 0xa41ee59c.
//
// Solidity: function createWalletAndDeposit(address _user, address _agent, bytes32 _salt, address _token, uint256 _amount, uint256 _lockDuration) returns(address)
func (_TrustManagementRouter *TrustManagementRouterTransactor) CreateWalletAndDeposit0(opts *bind.TransactOpts, _user common.Address, _agent common.Address, _salt [32]byte, _token common.Address, _amount *big.Int, _lockDuration *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "createWalletAndDeposit0", _user, _agent, _salt, _token, _amount, _lockDuration)
}

// CreateWalletAndDeposit0 is a paid mutator transaction binding the contract method 0xa41ee59c.
//
// Solidity: function createWalletAndDeposit(address _user, address _agent, bytes32 _salt, address _token, uint256 _amount, uint256 _lockDuration) returns(address)
func (_TrustManagementRouter *TrustManagementRouterSession) CreateWalletAndDeposit0(_user common.Address, _agent common.Address, _salt [32]byte, _token common.Address, _amount *big.Int, _lockDuration *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.CreateWalletAndDeposit0(&_TrustManagementRouter.TransactOpts, _user, _agent, _salt, _token, _amount, _lockDuration)
}

// CreateWalletAndDeposit0 is a paid mutator transaction binding the contract method 0xa41ee59c.
//
// Solidity: function createWalletAndDeposit(address _user, address _agent, bytes32 _salt, address _token, uint256 _amount, uint256 _lockDuration) returns(address)
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) CreateWalletAndDeposit0(_user common.Address, _agent common.Address, _salt [32]byte, _token common.Address, _amount *big.Int, _lockDuration *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.CreateWalletAndDeposit0(&_TrustManagementRouter.TransactOpts, _user, _agent, _salt, _token, _amount, _lockDuration)
}

// CreateWalletAndDepositWithPermit is a paid mutator transaction binding the contract method 0x781ea3be.
//
// Solidity: function createWalletAndDepositWithPermit(address _user, address _agent, bytes32 _salt, address _token, uint256 _unlockTimestamp, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns(address)
func (_TrustManagementRouter *TrustManagementRouterTransactor) CreateWalletAndDepositWithPermit(opts *bind.TransactOpts, _user common.Address, _agent common.Address, _salt [32]byte, _token common.Address, _unlockTimestamp *big.Int, _permit ITrustManagementStructsPermit) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "createWalletAndDepositWithPermit", _user, _agent, _salt, _token, _unlockTimestamp, _permit)
}

// CreateWalletAndDepositWithPermit is a paid mutator transaction binding the contract method 0x781ea3be.
//
// Solidity: function createWalletAndDepositWithPermit(address _user, address _agent, bytes32 _salt, address _token, uint256 _unlockTimestamp, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns(address)
func (_TrustManagementRouter *TrustManagementRouterSession) CreateWalletAndDepositWithPermit(_user common.Address, _agent common.Address, _salt [32]byte, _token common.Address, _unlockTimestamp *big.Int, _permit ITrustManagementStructsPermit) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.CreateWalletAndDepositWithPermit(&_TrustManagementRouter.TransactOpts, _user, _agent, _salt, _token, _unlockTimestamp, _permit)
}

// CreateWalletAndDepositWithPermit is a paid mutator transaction binding the contract method 0x781ea3be.
//
// Solidity: function createWalletAndDepositWithPermit(address _user, address _agent, bytes32 _salt, address _token, uint256 _unlockTimestamp, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns(address)
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) CreateWalletAndDepositWithPermit(_user common.Address, _agent common.Address, _salt [32]byte, _token common.Address, _unlockTimestamp *big.Int, _permit ITrustManagementStructsPermit) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.CreateWalletAndDepositWithPermit(&_TrustManagementRouter.TransactOpts, _user, _agent, _salt, _token, _unlockTimestamp, _permit)
}

// Deposit is a paid mutator transaction binding the contract method 0x02b9446c.
//
// Solidity: function deposit(address _user, address _wallet, address _token, uint256 _amount, uint256 _unlockTimestamp) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) Deposit(opts *bind.TransactOpts, _user common.Address, _wallet common.Address, _token common.Address, _amount *big.Int, _unlockTimestamp *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "deposit", _user, _wallet, _token, _amount, _unlockTimestamp)
}

// Deposit is a paid mutator transaction binding the contract method 0x02b9446c.
//
// Solidity: function deposit(address _user, address _wallet, address _token, uint256 _amount, uint256 _unlockTimestamp) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) Deposit(_user common.Address, _wallet common.Address, _token common.Address, _amount *big.Int, _unlockTimestamp *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Deposit(&_TrustManagementRouter.TransactOpts, _user, _wallet, _token, _amount, _unlockTimestamp)
}

// Deposit is a paid mutator transaction binding the contract method 0x02b9446c.
//
// Solidity: function deposit(address _user, address _wallet, address _token, uint256 _amount, uint256 _unlockTimestamp) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Deposit(_user common.Address, _wallet common.Address, _token common.Address, _amount *big.Int, _unlockTimestamp *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Deposit(&_TrustManagementRouter.TransactOpts, _user, _wallet, _token, _amount, _unlockTimestamp)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x20e8c565.
//
// Solidity: function deposit(address _wallet, address _token, uint256 _amount, uint256 _unlockTimestamp) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) Deposit0(opts *bind.TransactOpts, _wallet common.Address, _token common.Address, _amount *big.Int, _unlockTimestamp *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "deposit0", _wallet, _token, _amount, _unlockTimestamp)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x20e8c565.
//
// Solidity: function deposit(address _wallet, address _token, uint256 _amount, uint256 _unlockTimestamp) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) Deposit0(_wallet common.Address, _token common.Address, _amount *big.Int, _unlockTimestamp *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Deposit0(&_TrustManagementRouter.TransactOpts, _wallet, _token, _amount, _unlockTimestamp)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x20e8c565.
//
// Solidity: function deposit(address _wallet, address _token, uint256 _amount, uint256 _unlockTimestamp) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Deposit0(_wallet common.Address, _token common.Address, _amount *big.Int, _unlockTimestamp *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Deposit0(&_TrustManagementRouter.TransactOpts, _wallet, _token, _amount, _unlockTimestamp)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _wallet, uint256 _unlockTimestamp) payable returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) Deposit1(opts *bind.TransactOpts, _wallet common.Address, _unlockTimestamp *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "deposit1", _wallet, _unlockTimestamp)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _wallet, uint256 _unlockTimestamp) payable returns()
func (_TrustManagementRouter *TrustManagementRouterSession) Deposit1(_wallet common.Address, _unlockTimestamp *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Deposit1(&_TrustManagementRouter.TransactOpts, _wallet, _unlockTimestamp)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _wallet, uint256 _unlockTimestamp) payable returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Deposit1(_wallet common.Address, _unlockTimestamp *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Deposit1(&_TrustManagementRouter.TransactOpts, _wallet, _unlockTimestamp)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0xc860efa8.
//
// Solidity: function depositWithPermit(address _wallet, address _token, uint256 _unlockTimestamp, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) DepositWithPermit(opts *bind.TransactOpts, _wallet common.Address, _token common.Address, _unlockTimestamp *big.Int, _permit ITrustManagementStructsPermit) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "depositWithPermit", _wallet, _token, _unlockTimestamp, _permit)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0xc860efa8.
//
// Solidity: function depositWithPermit(address _wallet, address _token, uint256 _unlockTimestamp, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) DepositWithPermit(_wallet common.Address, _token common.Address, _unlockTimestamp *big.Int, _permit ITrustManagementStructsPermit) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.DepositWithPermit(&_TrustManagementRouter.TransactOpts, _wallet, _token, _unlockTimestamp, _permit)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0xc860efa8.
//
// Solidity: function depositWithPermit(address _wallet, address _token, uint256 _unlockTimestamp, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) DepositWithPermit(_wallet common.Address, _token common.Address, _unlockTimestamp *big.Int, _permit ITrustManagementStructsPermit) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.DepositWithPermit(&_TrustManagementRouter.TransactOpts, _wallet, _token, _unlockTimestamp, _permit)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x62441c23.
//
// Solidity: function emergencyWithdraw(address _wallet, address[] _tokens, uint256[][] _depositsIds, address _receiver) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _wallet common.Address, _tokens []common.Address, _depositsIds [][]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "emergencyWithdraw", _wallet, _tokens, _depositsIds, _receiver)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x62441c23.
//
// Solidity: function emergencyWithdraw(address _wallet, address[] _tokens, uint256[][] _depositsIds, address _receiver) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) EmergencyWithdraw(_wallet common.Address, _tokens []common.Address, _depositsIds [][]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.EmergencyWithdraw(&_TrustManagementRouter.TransactOpts, _wallet, _tokens, _depositsIds, _receiver)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x62441c23.
//
// Solidity: function emergencyWithdraw(address _wallet, address[] _tokens, uint256[][] _depositsIds, address _receiver) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) EmergencyWithdraw(_wallet common.Address, _tokens []common.Address, _depositsIds [][]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.EmergencyWithdraw(&_TrustManagementRouter.TransactOpts, _wallet, _tokens, _depositsIds, _receiver)
}

// EmergencyWithdraw0 is a paid mutator transaction binding the contract method 0x62472394.
//
// Solidity: function emergencyWithdraw(address _wallet, uint256[] _depositIds, address _receiver) payable returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) EmergencyWithdraw0(opts *bind.TransactOpts, _wallet common.Address, _depositIds []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "emergencyWithdraw0", _wallet, _depositIds, _receiver)
}

// EmergencyWithdraw0 is a paid mutator transaction binding the contract method 0x62472394.
//
// Solidity: function emergencyWithdraw(address _wallet, uint256[] _depositIds, address _receiver) payable returns()
func (_TrustManagementRouter *TrustManagementRouterSession) EmergencyWithdraw0(_wallet common.Address, _depositIds []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.EmergencyWithdraw0(&_TrustManagementRouter.TransactOpts, _wallet, _depositIds, _receiver)
}

// EmergencyWithdraw0 is a paid mutator transaction binding the contract method 0x62472394.
//
// Solidity: function emergencyWithdraw(address _wallet, uint256[] _depositIds, address _receiver) payable returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) EmergencyWithdraw0(_wallet common.Address, _depositIds []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.EmergencyWithdraw0(&_TrustManagementRouter.TransactOpts, _wallet, _depositIds, _receiver)
}

// Execute is a paid mutator transaction binding the contract method 0x42cc1b84.
//
// Solidity: function execute(address target, (address,uint256,bytes)[] txs) payable returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) Execute(opts *bind.TransactOpts, target common.Address, txs []ITrustManagementStructsTransaction) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "execute", target, txs)
}

// Execute is a paid mutator transaction binding the contract method 0x42cc1b84.
//
// Solidity: function execute(address target, (address,uint256,bytes)[] txs) payable returns()
func (_TrustManagementRouter *TrustManagementRouterSession) Execute(target common.Address, txs []ITrustManagementStructsTransaction) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Execute(&_TrustManagementRouter.TransactOpts, target, txs)
}

// Execute is a paid mutator transaction binding the contract method 0x42cc1b84.
//
// Solidity: function execute(address target, (address,uint256,bytes)[] txs) payable returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Execute(target common.Address, txs []ITrustManagementStructsTransaction) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Execute(&_TrustManagementRouter.TransactOpts, target, txs)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.GrantRole(&_TrustManagementRouter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.GrantRole(&_TrustManagementRouter.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _initBlueprint) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) Initialize(opts *bind.TransactOpts, _initBlueprint common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "initialize", _initBlueprint)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _initBlueprint) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) Initialize(_initBlueprint common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Initialize(&_TrustManagementRouter.TransactOpts, _initBlueprint)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _initBlueprint) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Initialize(_initBlueprint common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Initialize(&_TrustManagementRouter.TransactOpts, _initBlueprint)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.RenounceRole(&_TrustManagementRouter.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.RenounceRole(&_TrustManagementRouter.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.RevokeRole(&_TrustManagementRouter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.RevokeRole(&_TrustManagementRouter.TransactOpts, role, account)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _newFeeReceiver) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) SetFeeReceiver(opts *bind.TransactOpts, _newFeeReceiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "setFeeReceiver", _newFeeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _newFeeReceiver) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) SetFeeReceiver(_newFeeReceiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.SetFeeReceiver(&_TrustManagementRouter.TransactOpts, _newFeeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _newFeeReceiver) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) SetFeeReceiver(_newFeeReceiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.SetFeeReceiver(&_TrustManagementRouter.TransactOpts, _newFeeReceiver)
}

// SetPerfomanceFeeRate is a paid mutator transaction binding the contract method 0x24402bba.
//
// Solidity: function setPerfomanceFeeRate(uint256 _newPerfomanceFeeRate) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) SetPerfomanceFeeRate(opts *bind.TransactOpts, _newPerfomanceFeeRate *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "setPerfomanceFeeRate", _newPerfomanceFeeRate)
}

// SetPerfomanceFeeRate is a paid mutator transaction binding the contract method 0x24402bba.
//
// Solidity: function setPerfomanceFeeRate(uint256 _newPerfomanceFeeRate) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) SetPerfomanceFeeRate(_newPerfomanceFeeRate *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.SetPerfomanceFeeRate(&_TrustManagementRouter.TransactOpts, _newPerfomanceFeeRate)
}

// SetPerfomanceFeeRate is a paid mutator transaction binding the contract method 0x24402bba.
//
// Solidity: function setPerfomanceFeeRate(uint256 _newPerfomanceFeeRate) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) SetPerfomanceFeeRate(_newPerfomanceFeeRate *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.SetPerfomanceFeeRate(&_TrustManagementRouter.TransactOpts, _newPerfomanceFeeRate)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address _newImplementation) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) UpgradeTo(opts *bind.TransactOpts, _newImplementation common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "upgradeTo", _newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address _newImplementation) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) UpgradeTo(_newImplementation common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.UpgradeTo(&_TrustManagementRouter.TransactOpts, _newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address _newImplementation) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) UpgradeTo(_newImplementation common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.UpgradeTo(&_TrustManagementRouter.TransactOpts, _newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TrustManagementRouter *TrustManagementRouterSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.UpgradeToAndCall(&_TrustManagementRouter.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.UpgradeToAndCall(&_TrustManagementRouter.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x572bcaa7.
//
// Solidity: function withdraw(address _wallet, bytes _signature, address _token, uint256[] _depositIds, uint256[] _amountWithYield, address _receiver) payable returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) Withdraw(opts *bind.TransactOpts, _wallet common.Address, _signature []byte, _token common.Address, _depositIds []*big.Int, _amountWithYield []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "withdraw", _wallet, _signature, _token, _depositIds, _amountWithYield, _receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x572bcaa7.
//
// Solidity: function withdraw(address _wallet, bytes _signature, address _token, uint256[] _depositIds, uint256[] _amountWithYield, address _receiver) payable returns()
func (_TrustManagementRouter *TrustManagementRouterSession) Withdraw(_wallet common.Address, _signature []byte, _token common.Address, _depositIds []*big.Int, _amountWithYield []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Withdraw(&_TrustManagementRouter.TransactOpts, _wallet, _signature, _token, _depositIds, _amountWithYield, _receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x572bcaa7.
//
// Solidity: function withdraw(address _wallet, bytes _signature, address _token, uint256[] _depositIds, uint256[] _amountWithYield, address _receiver) payable returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Withdraw(_wallet common.Address, _signature []byte, _token common.Address, _depositIds []*big.Int, _amountWithYield []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Withdraw(&_TrustManagementRouter.TransactOpts, _wallet, _signature, _token, _depositIds, _amountWithYield, _receiver)
}

// TrustManagementRouterDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the TrustManagementRouter contract.
type TrustManagementRouterDepositEventIterator struct {
	Event *TrustManagementRouterDepositEvent // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterDepositEvent)
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
		it.Event = new(TrustManagementRouterDepositEvent)
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
func (it *TrustManagementRouterDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterDepositEvent represents a DepositEvent event raised by the TrustManagementRouter contract.
type TrustManagementRouterDepositEvent struct {
	Wallet          common.Address
	User            common.Address
	Token           common.Address
	Amount          *big.Int
	UnlockTimestamp *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0x06e15f44f14902aa990d1ee87818dbf5fefcd7cc78920c3833ef6aca41672704.
//
// Solidity: event DepositEvent(address wallet, address user, address token, uint256 amount, uint256 unlockTimestamp)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterDepositEvent(opts *bind.FilterOpts) (*TrustManagementRouterDepositEventIterator, error) {

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterDepositEventIterator{contract: _TrustManagementRouter.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0x06e15f44f14902aa990d1ee87818dbf5fefcd7cc78920c3833ef6aca41672704.
//
// Solidity: event DepositEvent(address wallet, address user, address token, uint256 amount, uint256 unlockTimestamp)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterDepositEvent) (event.Subscription, error) {

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterDepositEvent)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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

// ParseDepositEvent is a log parse operation binding the contract event 0x06e15f44f14902aa990d1ee87818dbf5fefcd7cc78920c3833ef6aca41672704.
//
// Solidity: event DepositEvent(address wallet, address user, address token, uint256 amount, uint256 unlockTimestamp)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseDepositEvent(log types.Log) (*TrustManagementRouterDepositEvent, error) {
	event := new(TrustManagementRouterDepositEvent)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "DepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementRouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TrustManagementRouter contract.
type TrustManagementRouterInitializedIterator struct {
	Event *TrustManagementRouterInitialized // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterInitialized)
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
		it.Event = new(TrustManagementRouterInitialized)
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
func (it *TrustManagementRouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterInitialized represents a Initialized event raised by the TrustManagementRouter contract.
type TrustManagementRouterInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*TrustManagementRouterInitializedIterator, error) {

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterInitializedIterator{contract: _TrustManagementRouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterInitialized) (event.Subscription, error) {

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterInitialized)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseInitialized(log types.Log) (*TrustManagementRouterInitialized, error) {
	event := new(TrustManagementRouterInitialized)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementRouterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TrustManagementRouter contract.
type TrustManagementRouterRoleAdminChangedIterator struct {
	Event *TrustManagementRouterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterRoleAdminChanged)
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
		it.Event = new(TrustManagementRouterRoleAdminChanged)
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
func (it *TrustManagementRouterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterRoleAdminChanged represents a RoleAdminChanged event raised by the TrustManagementRouter contract.
type TrustManagementRouterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TrustManagementRouterRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterRoleAdminChangedIterator{contract: _TrustManagementRouter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterRoleAdminChanged)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseRoleAdminChanged(log types.Log) (*TrustManagementRouterRoleAdminChanged, error) {
	event := new(TrustManagementRouterRoleAdminChanged)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementRouterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TrustManagementRouter contract.
type TrustManagementRouterRoleGrantedIterator struct {
	Event *TrustManagementRouterRoleGranted // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterRoleGranted)
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
		it.Event = new(TrustManagementRouterRoleGranted)
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
func (it *TrustManagementRouterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterRoleGranted represents a RoleGranted event raised by the TrustManagementRouter contract.
type TrustManagementRouterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TrustManagementRouterRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterRoleGrantedIterator{contract: _TrustManagementRouter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterRoleGranted)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseRoleGranted(log types.Log) (*TrustManagementRouterRoleGranted, error) {
	event := new(TrustManagementRouterRoleGranted)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementRouterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TrustManagementRouter contract.
type TrustManagementRouterRoleRevokedIterator struct {
	Event *TrustManagementRouterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterRoleRevoked)
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
		it.Event = new(TrustManagementRouterRoleRevoked)
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
func (it *TrustManagementRouterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterRoleRevoked represents a RoleRevoked event raised by the TrustManagementRouter contract.
type TrustManagementRouterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TrustManagementRouterRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterRoleRevokedIterator{contract: _TrustManagementRouter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterRoleRevoked)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseRoleRevoked(log types.Log) (*TrustManagementRouterRoleRevoked, error) {
	event := new(TrustManagementRouterRoleRevoked)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementRouterUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TrustManagementRouter contract.
type TrustManagementRouterUpgradedIterator struct {
	Event *TrustManagementRouterUpgraded // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterUpgraded)
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
		it.Event = new(TrustManagementRouterUpgraded)
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
func (it *TrustManagementRouterUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterUpgraded represents a Upgraded event raised by the TrustManagementRouter contract.
type TrustManagementRouterUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TrustManagementRouterUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterUpgradedIterator{contract: _TrustManagementRouter.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterUpgraded)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseUpgraded(log types.Log) (*TrustManagementRouterUpgraded, error) {
	event := new(TrustManagementRouterUpgraded)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
