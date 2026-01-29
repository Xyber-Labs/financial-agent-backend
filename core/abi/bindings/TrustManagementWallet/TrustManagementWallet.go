// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TrustManagementWallet

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

// Deposit is an auto generated low-level Go binding around an user-defined struct.
type Deposit struct {
	Amount      *big.Int
	LockedUntil *big.Int
}

// Transaction is an auto generated low-level Go binding around an user-defined struct.
type Transaction struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
}

// TrustManagementWalletMetaData contains all meta data concerning the TrustManagementWallet contract.
var TrustManagementWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minLockDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLockDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XyberTrustManagementRouter__EthereumTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"XyberTrustManagementWallet__CallerIsNotRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XyberTrustManagementWallet__DeadlineExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"}],\"name\":\"XyberTrustManagementWallet__DepositLocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"XyberTrustManagementWallet__ExecuteCallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithYield\",\"type\":\"uint256\"}],\"name\":\"XyberTrustManagementWallet__InvalidAmountWithYield\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"XyberTrustManagementWallet__InvalidCallTarget\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"}],\"name\":\"XyberTrustManagementWallet__InvalidDepositId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"XyberTrustManagementWallet__InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XyberTrustManagementWallet__InvalidUnlockTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XyberTrustManagementWallet__ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XyberUpgradeChecker__E0\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposited\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ETH_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LOCK_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_LOCK_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction[]\",\"name\":\"txs\",\"type\":\"tuple[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDeposits\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"}],\"internalType\":\"structDeposit[]\",\"name\":\"deposits\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getWithdrawableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawableAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getWithdrawableDeposits\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"withdrawableDepositsIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"totalDeposited\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalDepositedTokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amountWithYield\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xyberContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"contractName\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TrustManagementWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use TrustManagementWalletMetaData.ABI instead.
var TrustManagementWalletABI = TrustManagementWalletMetaData.ABI

// TrustManagementWallet is an auto generated Go binding around an Ethereum contract.
type TrustManagementWallet struct {
	TrustManagementWalletCaller     // Read-only binding to the contract
	TrustManagementWalletTransactor // Write-only binding to the contract
	TrustManagementWalletFilterer   // Log filterer for contract events
}

// TrustManagementWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrustManagementWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustManagementWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrustManagementWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustManagementWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrustManagementWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustManagementWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrustManagementWalletSession struct {
	Contract     *TrustManagementWallet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TrustManagementWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrustManagementWalletCallerSession struct {
	Contract *TrustManagementWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TrustManagementWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrustManagementWalletTransactorSession struct {
	Contract     *TrustManagementWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TrustManagementWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrustManagementWalletRaw struct {
	Contract *TrustManagementWallet // Generic contract binding to access the raw methods on
}

// TrustManagementWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrustManagementWalletCallerRaw struct {
	Contract *TrustManagementWalletCaller // Generic read-only contract binding to access the raw methods on
}

// TrustManagementWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrustManagementWalletTransactorRaw struct {
	Contract *TrustManagementWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrustManagementWallet creates a new instance of TrustManagementWallet, bound to a specific deployed contract.
func NewTrustManagementWallet(address common.Address, backend bind.ContractBackend) (*TrustManagementWallet, error) {
	contract, err := bindTrustManagementWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrustManagementWallet{TrustManagementWalletCaller: TrustManagementWalletCaller{contract: contract}, TrustManagementWalletTransactor: TrustManagementWalletTransactor{contract: contract}, TrustManagementWalletFilterer: TrustManagementWalletFilterer{contract: contract}}, nil
}

// NewTrustManagementWalletCaller creates a new read-only instance of TrustManagementWallet, bound to a specific deployed contract.
func NewTrustManagementWalletCaller(address common.Address, caller bind.ContractCaller) (*TrustManagementWalletCaller, error) {
	contract, err := bindTrustManagementWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrustManagementWalletCaller{contract: contract}, nil
}

// NewTrustManagementWalletTransactor creates a new write-only instance of TrustManagementWallet, bound to a specific deployed contract.
func NewTrustManagementWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*TrustManagementWalletTransactor, error) {
	contract, err := bindTrustManagementWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrustManagementWalletTransactor{contract: contract}, nil
}

// NewTrustManagementWalletFilterer creates a new log filterer instance of TrustManagementWallet, bound to a specific deployed contract.
func NewTrustManagementWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*TrustManagementWalletFilterer, error) {
	contract, err := bindTrustManagementWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrustManagementWalletFilterer{contract: contract}, nil
}

// bindTrustManagementWallet binds a generic wrapper to an already deployed contract.
func bindTrustManagementWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TrustManagementWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustManagementWallet *TrustManagementWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustManagementWallet.Contract.TrustManagementWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustManagementWallet *TrustManagementWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.TrustManagementWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustManagementWallet *TrustManagementWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.TrustManagementWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustManagementWallet *TrustManagementWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustManagementWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustManagementWallet *TrustManagementWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustManagementWallet *TrustManagementWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.contract.Transact(opts, method, params...)
}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_TrustManagementWallet *TrustManagementWalletCaller) ETHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrustManagementWallet.contract.Call(opts, &out, "ETH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_TrustManagementWallet *TrustManagementWalletSession) ETHADDRESS() (common.Address, error) {
	return _TrustManagementWallet.Contract.ETHADDRESS(&_TrustManagementWallet.CallOpts)
}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_TrustManagementWallet *TrustManagementWalletCallerSession) ETHADDRESS() (common.Address, error) {
	return _TrustManagementWallet.Contract.ETHADDRESS(&_TrustManagementWallet.CallOpts)
}

// MAXLOCKDURATION is a free data retrieval call binding the contract method 0x4f1bfc9e.
//
// Solidity: function MAX_LOCK_DURATION() view returns(uint256)
func (_TrustManagementWallet *TrustManagementWalletCaller) MAXLOCKDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrustManagementWallet.contract.Call(opts, &out, "MAX_LOCK_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLOCKDURATION is a free data retrieval call binding the contract method 0x4f1bfc9e.
//
// Solidity: function MAX_LOCK_DURATION() view returns(uint256)
func (_TrustManagementWallet *TrustManagementWalletSession) MAXLOCKDURATION() (*big.Int, error) {
	return _TrustManagementWallet.Contract.MAXLOCKDURATION(&_TrustManagementWallet.CallOpts)
}

// MAXLOCKDURATION is a free data retrieval call binding the contract method 0x4f1bfc9e.
//
// Solidity: function MAX_LOCK_DURATION() view returns(uint256)
func (_TrustManagementWallet *TrustManagementWalletCallerSession) MAXLOCKDURATION() (*big.Int, error) {
	return _TrustManagementWallet.Contract.MAXLOCKDURATION(&_TrustManagementWallet.CallOpts)
}

// MINLOCKDURATION is a free data retrieval call binding the contract method 0x78b4330f.
//
// Solidity: function MIN_LOCK_DURATION() view returns(uint256)
func (_TrustManagementWallet *TrustManagementWalletCaller) MINLOCKDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrustManagementWallet.contract.Call(opts, &out, "MIN_LOCK_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINLOCKDURATION is a free data retrieval call binding the contract method 0x78b4330f.
//
// Solidity: function MIN_LOCK_DURATION() view returns(uint256)
func (_TrustManagementWallet *TrustManagementWalletSession) MINLOCKDURATION() (*big.Int, error) {
	return _TrustManagementWallet.Contract.MINLOCKDURATION(&_TrustManagementWallet.CallOpts)
}

// MINLOCKDURATION is a free data retrieval call binding the contract method 0x78b4330f.
//
// Solidity: function MIN_LOCK_DURATION() view returns(uint256)
func (_TrustManagementWallet *TrustManagementWalletCallerSession) MINLOCKDURATION() (*big.Int, error) {
	return _TrustManagementWallet.Contract.MINLOCKDURATION(&_TrustManagementWallet.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_TrustManagementWallet *TrustManagementWalletCaller) ROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrustManagementWallet.contract.Call(opts, &out, "ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_TrustManagementWallet *TrustManagementWalletSession) ROUTER() (common.Address, error) {
	return _TrustManagementWallet.Contract.ROUTER(&_TrustManagementWallet.CallOpts)
}

// ROUTER is a free data retrieval call binding the contract method 0x32fe7b26.
//
// Solidity: function ROUTER() view returns(address)
func (_TrustManagementWallet *TrustManagementWalletCallerSession) ROUTER() (common.Address, error) {
	return _TrustManagementWallet.Contract.ROUTER(&_TrustManagementWallet.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x2d2ae1c1.
//
// Solidity: function getBalances(address[] tokens) view returns(uint256[] balances)
func (_TrustManagementWallet *TrustManagementWalletCaller) GetBalances(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TrustManagementWallet.contract.Call(opts, &out, "getBalances", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x2d2ae1c1.
//
// Solidity: function getBalances(address[] tokens) view returns(uint256[] balances)
func (_TrustManagementWallet *TrustManagementWalletSession) GetBalances(tokens []common.Address) ([]*big.Int, error) {
	return _TrustManagementWallet.Contract.GetBalances(&_TrustManagementWallet.CallOpts, tokens)
}

// GetBalances is a free data retrieval call binding the contract method 0x2d2ae1c1.
//
// Solidity: function getBalances(address[] tokens) view returns(uint256[] balances)
func (_TrustManagementWallet *TrustManagementWalletCallerSession) GetBalances(tokens []common.Address) ([]*big.Int, error) {
	return _TrustManagementWallet.Contract.GetBalances(&_TrustManagementWallet.CallOpts, tokens)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address token) view returns((uint256,uint256)[] deposits)
func (_TrustManagementWallet *TrustManagementWalletCaller) GetDeposits(opts *bind.CallOpts, token common.Address) ([]Deposit, error) {
	var out []interface{}
	err := _TrustManagementWallet.contract.Call(opts, &out, "getDeposits", token)

	if err != nil {
		return *new([]Deposit), err
	}

	out0 := *abi.ConvertType(out[0], new([]Deposit)).(*[]Deposit)

	return out0, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address token) view returns((uint256,uint256)[] deposits)
func (_TrustManagementWallet *TrustManagementWalletSession) GetDeposits(token common.Address) ([]Deposit, error) {
	return _TrustManagementWallet.Contract.GetDeposits(&_TrustManagementWallet.CallOpts, token)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address token) view returns((uint256,uint256)[] deposits)
func (_TrustManagementWallet *TrustManagementWalletCallerSession) GetDeposits(token common.Address) ([]Deposit, error) {
	return _TrustManagementWallet.Contract.GetDeposits(&_TrustManagementWallet.CallOpts, token)
}

// GetWithdrawableAmount is a free data retrieval call binding the contract method 0xf6777175.
//
// Solidity: function getWithdrawableAmount(address token) view returns(uint256 withdrawableAmount)
func (_TrustManagementWallet *TrustManagementWalletCaller) GetWithdrawableAmount(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TrustManagementWallet.contract.Call(opts, &out, "getWithdrawableAmount", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableAmount is a free data retrieval call binding the contract method 0xf6777175.
//
// Solidity: function getWithdrawableAmount(address token) view returns(uint256 withdrawableAmount)
func (_TrustManagementWallet *TrustManagementWalletSession) GetWithdrawableAmount(token common.Address) (*big.Int, error) {
	return _TrustManagementWallet.Contract.GetWithdrawableAmount(&_TrustManagementWallet.CallOpts, token)
}

// GetWithdrawableAmount is a free data retrieval call binding the contract method 0xf6777175.
//
// Solidity: function getWithdrawableAmount(address token) view returns(uint256 withdrawableAmount)
func (_TrustManagementWallet *TrustManagementWalletCallerSession) GetWithdrawableAmount(token common.Address) (*big.Int, error) {
	return _TrustManagementWallet.Contract.GetWithdrawableAmount(&_TrustManagementWallet.CallOpts, token)
}

// GetWithdrawableDeposits is a free data retrieval call binding the contract method 0x34e57ad8.
//
// Solidity: function getWithdrawableDeposits(address token) view returns(uint256[] withdrawableDepositsIds)
func (_TrustManagementWallet *TrustManagementWalletCaller) GetWithdrawableDeposits(opts *bind.CallOpts, token common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TrustManagementWallet.contract.Call(opts, &out, "getWithdrawableDeposits", token)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetWithdrawableDeposits is a free data retrieval call binding the contract method 0x34e57ad8.
//
// Solidity: function getWithdrawableDeposits(address token) view returns(uint256[] withdrawableDepositsIds)
func (_TrustManagementWallet *TrustManagementWalletSession) GetWithdrawableDeposits(token common.Address) ([]*big.Int, error) {
	return _TrustManagementWallet.Contract.GetWithdrawableDeposits(&_TrustManagementWallet.CallOpts, token)
}

// GetWithdrawableDeposits is a free data retrieval call binding the contract method 0x34e57ad8.
//
// Solidity: function getWithdrawableDeposits(address token) view returns(uint256[] withdrawableDepositsIds)
func (_TrustManagementWallet *TrustManagementWalletCallerSession) GetWithdrawableDeposits(token common.Address) ([]*big.Int, error) {
	return _TrustManagementWallet.Contract.GetWithdrawableDeposits(&_TrustManagementWallet.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address ownerAddress)
func (_TrustManagementWallet *TrustManagementWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrustManagementWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address ownerAddress)
func (_TrustManagementWallet *TrustManagementWalletSession) Owner() (common.Address, error) {
	return _TrustManagementWallet.Contract.Owner(&_TrustManagementWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address ownerAddress)
func (_TrustManagementWallet *TrustManagementWalletCallerSession) Owner() (common.Address, error) {
	return _TrustManagementWallet.Contract.Owner(&_TrustManagementWallet.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TrustManagementWallet *TrustManagementWalletCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TrustManagementWallet.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TrustManagementWallet *TrustManagementWalletSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TrustManagementWallet.Contract.SupportsInterface(&_TrustManagementWallet.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TrustManagementWallet *TrustManagementWalletCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TrustManagementWallet.Contract.SupportsInterface(&_TrustManagementWallet.CallOpts, interfaceId)
}

// TotalDeposited is a free data retrieval call binding the contract method 0x53055481.
//
// Solidity: function totalDeposited(address token) view returns(uint256 totalDepositedTokens)
func (_TrustManagementWallet *TrustManagementWalletCaller) TotalDeposited(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TrustManagementWallet.contract.Call(opts, &out, "totalDeposited", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeposited is a free data retrieval call binding the contract method 0x53055481.
//
// Solidity: function totalDeposited(address token) view returns(uint256 totalDepositedTokens)
func (_TrustManagementWallet *TrustManagementWalletSession) TotalDeposited(token common.Address) (*big.Int, error) {
	return _TrustManagementWallet.Contract.TotalDeposited(&_TrustManagementWallet.CallOpts, token)
}

// TotalDeposited is a free data retrieval call binding the contract method 0x53055481.
//
// Solidity: function totalDeposited(address token) view returns(uint256 totalDepositedTokens)
func (_TrustManagementWallet *TrustManagementWalletCallerSession) TotalDeposited(token common.Address) (*big.Int, error) {
	return _TrustManagementWallet.Contract.TotalDeposited(&_TrustManagementWallet.CallOpts, token)
}

// XyberContractName is a free data retrieval call binding the contract method 0xe60e9654.
//
// Solidity: function xyberContractName() pure returns(string contractName)
func (_TrustManagementWallet *TrustManagementWalletCaller) XyberContractName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TrustManagementWallet.contract.Call(opts, &out, "xyberContractName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// XyberContractName is a free data retrieval call binding the contract method 0xe60e9654.
//
// Solidity: function xyberContractName() pure returns(string contractName)
func (_TrustManagementWallet *TrustManagementWalletSession) XyberContractName() (string, error) {
	return _TrustManagementWallet.Contract.XyberContractName(&_TrustManagementWallet.CallOpts)
}

// XyberContractName is a free data retrieval call binding the contract method 0xe60e9654.
//
// Solidity: function xyberContractName() pure returns(string contractName)
func (_TrustManagementWallet *TrustManagementWalletCallerSession) XyberContractName() (string, error) {
	return _TrustManagementWallet.Contract.XyberContractName(&_TrustManagementWallet.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x9fae2d52.
//
// Solidity: function claim(address receiver, address[] tokens, uint256[] amounts) payable returns(bool success)
func (_TrustManagementWallet *TrustManagementWalletTransactor) Claim(opts *bind.TransactOpts, receiver common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TrustManagementWallet.contract.Transact(opts, "claim", receiver, tokens, amounts)
}

// Claim is a paid mutator transaction binding the contract method 0x9fae2d52.
//
// Solidity: function claim(address receiver, address[] tokens, uint256[] amounts) payable returns(bool success)
func (_TrustManagementWallet *TrustManagementWalletSession) Claim(receiver common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.Claim(&_TrustManagementWallet.TransactOpts, receiver, tokens, amounts)
}

// Claim is a paid mutator transaction binding the contract method 0x9fae2d52.
//
// Solidity: function claim(address receiver, address[] tokens, uint256[] amounts) payable returns(bool success)
func (_TrustManagementWallet *TrustManagementWalletTransactorSession) Claim(receiver common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.Claim(&_TrustManagementWallet.TransactOpts, receiver, tokens, amounts)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address token, uint256 amount, uint256 lockDuration) returns(bool success)
func (_TrustManagementWallet *TrustManagementWalletTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int, lockDuration *big.Int) (*types.Transaction, error) {
	return _TrustManagementWallet.contract.Transact(opts, "deposit", token, amount, lockDuration)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address token, uint256 amount, uint256 lockDuration) returns(bool success)
func (_TrustManagementWallet *TrustManagementWalletSession) Deposit(token common.Address, amount *big.Int, lockDuration *big.Int) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.Deposit(&_TrustManagementWallet.TransactOpts, token, amount, lockDuration)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address token, uint256 amount, uint256 lockDuration) returns(bool success)
func (_TrustManagementWallet *TrustManagementWalletTransactorSession) Deposit(token common.Address, amount *big.Int, lockDuration *big.Int) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.Deposit(&_TrustManagementWallet.TransactOpts, token, amount, lockDuration)
}

// Execute is a paid mutator transaction binding the contract method 0x3f707e6b.
//
// Solidity: function execute((address,uint256,bytes)[] txs) returns()
func (_TrustManagementWallet *TrustManagementWalletTransactor) Execute(opts *bind.TransactOpts, txs []Transaction) (*types.Transaction, error) {
	return _TrustManagementWallet.contract.Transact(opts, "execute", txs)
}

// Execute is a paid mutator transaction binding the contract method 0x3f707e6b.
//
// Solidity: function execute((address,uint256,bytes)[] txs) returns()
func (_TrustManagementWallet *TrustManagementWalletSession) Execute(txs []Transaction) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.Execute(&_TrustManagementWallet.TransactOpts, txs)
}

// Execute is a paid mutator transaction binding the contract method 0x3f707e6b.
//
// Solidity: function execute((address,uint256,bytes)[] txs) returns()
func (_TrustManagementWallet *TrustManagementWalletTransactorSession) Execute(txs []Transaction) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.Execute(&_TrustManagementWallet.TransactOpts, txs)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newOwner) returns()
func (_TrustManagementWallet *TrustManagementWalletTransactor) Initialize(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TrustManagementWallet.contract.Transact(opts, "initialize", newOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newOwner) returns()
func (_TrustManagementWallet *TrustManagementWalletSession) Initialize(newOwner common.Address) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.Initialize(&_TrustManagementWallet.TransactOpts, newOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newOwner) returns()
func (_TrustManagementWallet *TrustManagementWalletTransactorSession) Initialize(newOwner common.Address) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.Initialize(&_TrustManagementWallet.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb0d5c7ef.
//
// Solidity: function withdraw(address token, address receiver, uint256[] depositIds, uint256 amountWithYield) payable returns(uint256 amount, uint256 feeAmount)
func (_TrustManagementWallet *TrustManagementWalletTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, receiver common.Address, depositIds []*big.Int, amountWithYield *big.Int) (*types.Transaction, error) {
	return _TrustManagementWallet.contract.Transact(opts, "withdraw", token, receiver, depositIds, amountWithYield)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb0d5c7ef.
//
// Solidity: function withdraw(address token, address receiver, uint256[] depositIds, uint256 amountWithYield) payable returns(uint256 amount, uint256 feeAmount)
func (_TrustManagementWallet *TrustManagementWalletSession) Withdraw(token common.Address, receiver common.Address, depositIds []*big.Int, amountWithYield *big.Int) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.Withdraw(&_TrustManagementWallet.TransactOpts, token, receiver, depositIds, amountWithYield)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb0d5c7ef.
//
// Solidity: function withdraw(address token, address receiver, uint256[] depositIds, uint256 amountWithYield) payable returns(uint256 amount, uint256 feeAmount)
func (_TrustManagementWallet *TrustManagementWalletTransactorSession) Withdraw(token common.Address, receiver common.Address, depositIds []*big.Int, amountWithYield *big.Int) (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.Withdraw(&_TrustManagementWallet.TransactOpts, token, receiver, depositIds, amountWithYield)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TrustManagementWallet *TrustManagementWalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustManagementWallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TrustManagementWallet *TrustManagementWalletSession) Receive() (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.Receive(&_TrustManagementWallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TrustManagementWallet *TrustManagementWalletTransactorSession) Receive() (*types.Transaction, error) {
	return _TrustManagementWallet.Contract.Receive(&_TrustManagementWallet.TransactOpts)
}

// TrustManagementWalletClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the TrustManagementWallet contract.
type TrustManagementWalletClaimedIterator struct {
	Event *TrustManagementWalletClaimed // Event containing the contract specifics and raw log

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
func (it *TrustManagementWalletClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementWalletClaimed)
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
		it.Event = new(TrustManagementWalletClaimed)
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
func (it *TrustManagementWalletClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementWalletClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementWalletClaimed represents a Claimed event raised by the TrustManagementWallet contract.
type TrustManagementWalletClaimed struct {
	Token     common.Address
	Receiver  common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x2f6639d24651730c7bf57c95ddbf96d66d11477e4ec626876f92c22e5f365e68.
//
// Solidity: event Claimed(address indexed token, address indexed receiver, uint256 amount, uint256 feeAmount)
func (_TrustManagementWallet *TrustManagementWalletFilterer) FilterClaimed(opts *bind.FilterOpts, token []common.Address, receiver []common.Address) (*TrustManagementWalletClaimedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _TrustManagementWallet.contract.FilterLogs(opts, "Claimed", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementWalletClaimedIterator{contract: _TrustManagementWallet.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x2f6639d24651730c7bf57c95ddbf96d66d11477e4ec626876f92c22e5f365e68.
//
// Solidity: event Claimed(address indexed token, address indexed receiver, uint256 amount, uint256 feeAmount)
func (_TrustManagementWallet *TrustManagementWalletFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *TrustManagementWalletClaimed, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _TrustManagementWallet.contract.WatchLogs(opts, "Claimed", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementWalletClaimed)
				if err := _TrustManagementWallet.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x2f6639d24651730c7bf57c95ddbf96d66d11477e4ec626876f92c22e5f365e68.
//
// Solidity: event Claimed(address indexed token, address indexed receiver, uint256 amount, uint256 feeAmount)
func (_TrustManagementWallet *TrustManagementWalletFilterer) ParseClaimed(log types.Log) (*TrustManagementWalletClaimed, error) {
	event := new(TrustManagementWalletClaimed)
	if err := _TrustManagementWallet.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementWalletDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the TrustManagementWallet contract.
type TrustManagementWalletDepositedIterator struct {
	Event *TrustManagementWalletDeposited // Event containing the contract specifics and raw log

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
func (it *TrustManagementWalletDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementWalletDeposited)
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
		it.Event = new(TrustManagementWalletDeposited)
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
func (it *TrustManagementWalletDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementWalletDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementWalletDeposited represents a Deposited event raised by the TrustManagementWallet contract.
type TrustManagementWalletDeposited struct {
	Token          common.Address
	Amount         *big.Int
	LockDuration   *big.Int
	DepositId      *big.Int
	TotalDeposited *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xe3b53cd1a44fbf11535e145d80b8ef1ed6d57a73bf5daa7e939b6b01657d6549.
//
// Solidity: event Deposited(address indexed token, uint256 amount, uint256 lockDuration, uint256 depositId, uint256 totalDeposited)
func (_TrustManagementWallet *TrustManagementWalletFilterer) FilterDeposited(opts *bind.FilterOpts, token []common.Address) (*TrustManagementWalletDepositedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TrustManagementWallet.contract.FilterLogs(opts, "Deposited", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementWalletDepositedIterator{contract: _TrustManagementWallet.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xe3b53cd1a44fbf11535e145d80b8ef1ed6d57a73bf5daa7e939b6b01657d6549.
//
// Solidity: event Deposited(address indexed token, uint256 amount, uint256 lockDuration, uint256 depositId, uint256 totalDeposited)
func (_TrustManagementWallet *TrustManagementWalletFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *TrustManagementWalletDeposited, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TrustManagementWallet.contract.WatchLogs(opts, "Deposited", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementWalletDeposited)
				if err := _TrustManagementWallet.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xe3b53cd1a44fbf11535e145d80b8ef1ed6d57a73bf5daa7e939b6b01657d6549.
//
// Solidity: event Deposited(address indexed token, uint256 amount, uint256 lockDuration, uint256 depositId, uint256 totalDeposited)
func (_TrustManagementWallet *TrustManagementWalletFilterer) ParseDeposited(log types.Log) (*TrustManagementWalletDeposited, error) {
	event := new(TrustManagementWalletDeposited)
	if err := _TrustManagementWallet.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementWalletInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TrustManagementWallet contract.
type TrustManagementWalletInitializedIterator struct {
	Event *TrustManagementWalletInitialized // Event containing the contract specifics and raw log

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
func (it *TrustManagementWalletInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementWalletInitialized)
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
		it.Event = new(TrustManagementWalletInitialized)
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
func (it *TrustManagementWalletInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementWalletInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementWalletInitialized represents a Initialized event raised by the TrustManagementWallet contract.
type TrustManagementWalletInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TrustManagementWallet *TrustManagementWalletFilterer) FilterInitialized(opts *bind.FilterOpts) (*TrustManagementWalletInitializedIterator, error) {

	logs, sub, err := _TrustManagementWallet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TrustManagementWalletInitializedIterator{contract: _TrustManagementWallet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TrustManagementWallet *TrustManagementWalletFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TrustManagementWalletInitialized) (event.Subscription, error) {

	logs, sub, err := _TrustManagementWallet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementWalletInitialized)
				if err := _TrustManagementWallet.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TrustManagementWallet *TrustManagementWalletFilterer) ParseInitialized(log types.Log) (*TrustManagementWalletInitialized, error) {
	event := new(TrustManagementWalletInitialized)
	if err := _TrustManagementWallet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementWalletWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the TrustManagementWallet contract.
type TrustManagementWalletWithdrawnIterator struct {
	Event *TrustManagementWalletWithdrawn // Event containing the contract specifics and raw log

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
func (it *TrustManagementWalletWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementWalletWithdrawn)
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
		it.Event = new(TrustManagementWalletWithdrawn)
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
func (it *TrustManagementWalletWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementWalletWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementWalletWithdrawn represents a Withdrawn event raised by the TrustManagementWallet contract.
type TrustManagementWalletWithdrawn struct {
	Token     common.Address
	Receiver  common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x91fb9d98b786c57d74c099ccd2beca1739e9f6a81fb49001ca465c4b7591bbe2.
//
// Solidity: event Withdrawn(address indexed token, address indexed receiver, uint256 amount, uint256 feeAmount)
func (_TrustManagementWallet *TrustManagementWalletFilterer) FilterWithdrawn(opts *bind.FilterOpts, token []common.Address, receiver []common.Address) (*TrustManagementWalletWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _TrustManagementWallet.contract.FilterLogs(opts, "Withdrawn", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementWalletWithdrawnIterator{contract: _TrustManagementWallet.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x91fb9d98b786c57d74c099ccd2beca1739e9f6a81fb49001ca465c4b7591bbe2.
//
// Solidity: event Withdrawn(address indexed token, address indexed receiver, uint256 amount, uint256 feeAmount)
func (_TrustManagementWallet *TrustManagementWalletFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *TrustManagementWalletWithdrawn, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _TrustManagementWallet.contract.WatchLogs(opts, "Withdrawn", tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementWalletWithdrawn)
				if err := _TrustManagementWallet.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x91fb9d98b786c57d74c099ccd2beca1739e9f6a81fb49001ca465c4b7591bbe2.
//
// Solidity: event Withdrawn(address indexed token, address indexed receiver, uint256 amount, uint256 feeAmount)
func (_TrustManagementWallet *TrustManagementWalletFilterer) ParseWithdrawn(log types.Log) (*TrustManagementWalletWithdrawn, error) {
	event := new(TrustManagementWalletWithdrawn)
	if err := _TrustManagementWallet.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
