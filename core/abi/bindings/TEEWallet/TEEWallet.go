// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TEEWallet

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

// ChunkedSGXQuote is an auto generated low-level Go binding around an user-defined struct.
type ChunkedSGXQuote struct {
	Header               []byte
	IsvReport            []byte
	IsvReportSignature   []byte
	AttestationKey       []byte
	QeReport             []byte
	QeReportSignature    []byte
	QeAuthenticationData []byte
}

// ChunkedX509Cert is an auto generated low-level Go binding around an user-defined struct.
type ChunkedX509Cert struct {
	BodyPartOne []byte
	PublicKey   []byte
	BodyPartTwo []byte
	Signature   []byte
}

// ITEEWalletuserData is an auto generated low-level Go binding around an user-defined struct.
type ITEEWalletuserData struct {
	UserAddress common.Address
	Amount      *big.Int
}

// Transaction is an auto generated low-level Go binding around an user-defined struct.
type Transaction struct {
	Target     common.Address
	Value      *big.Int
	Data       []byte
	IsDelegate bool
}

// TEEWalletMetaData contains all meta data concerning the TEEWallet contract.
var TEEWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"TEEWallet__CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"}],\"name\":\"TEEWallet__InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TEEWallet__InvalidEnclaveSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TEEWallet__InvalidIntermediateCert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TEEWallet__InvalidLeafCert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"TEEWallet__InvalidNonce_or_RefundAlreadyProcessed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TEEWallet__InvalidQeReportDataHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TEEWallet__InvalidQeSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"TEEWallet__InvalidSessionKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"TEEWallet__InvalidTarget\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"callIndex\",\"type\":\"uint256\"}],\"name\":\"TEEWallet__InvalidValueForDelegateCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"TEEWallet__MessageValueIsTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"TEEWallet__NewImplementationIsInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TEEWallet__UserDataAmountIsTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHashFromSessionKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"reportDataHash\",\"type\":\"bytes32\"}],\"name\":\"TEEWallet__WrongDataHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mrEnclave\",\"type\":\"bytes32\"}],\"name\":\"TEEWallet__WrongMrEnclave\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msgsender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"FundsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"scriptHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gitHubLinkPublic\",\"type\":\"string\"}],\"name\":\"ScriptSetted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_mrEnclave\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reportDataHash\",\"type\":\"bytes32\"}],\"name\":\"SessionInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SESSION_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"ecdsaOnKeccak256r1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isDelegate\",\"type\":\"bool\"}],\"internalType\":\"structTransaction[]\",\"name\":\"txs\",\"type\":\"tuple[]\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gitHubLink\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"bodyPartOne\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bodyPartTwo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedX509Cert\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"bodyPartOne\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bodyPartTwo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedX509Cert\",\"name\":\"intermediate\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"isvReport\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"isvReportSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestationKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"qeReport\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"qeReportSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"qeAuthenticationData\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedSGXQuote\",\"name\":\"quote\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"initSessionKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sessionAdmin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"bodyPartOne\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bodyPartTwo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedX509Cert\",\"name\":\"_rootCert\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"nonceToUserData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structITEEWallet.userData\",\"name\":\"uData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootCert\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"bodyPartOne\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bodyPartTwo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedX509Cert\",\"name\":\"cert\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scriptHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"sessions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"bodyPartOne\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bodyPartTwo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedX509Cert\",\"name\":\"_newRootCert\",\"type\":\"tuple\"}],\"name\":\"setRootCert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"scriptHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"gitHubLinkPublic\",\"type\":\"string\"}],\"name\":\"setScript\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"bodyPartOne\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bodyPartTwo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedX509Cert\",\"name\":\"cert\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"verifyCert\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TEEWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use TEEWalletMetaData.ABI instead.
var TEEWalletABI = TEEWalletMetaData.ABI

// TEEWallet is an auto generated Go binding around an Ethereum contract.
type TEEWallet struct {
	TEEWalletCaller     // Read-only binding to the contract
	TEEWalletTransactor // Write-only binding to the contract
	TEEWalletFilterer   // Log filterer for contract events
}

// TEEWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type TEEWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TEEWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TEEWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TEEWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TEEWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TEEWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TEEWalletSession struct {
	Contract     *TEEWallet        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TEEWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TEEWalletCallerSession struct {
	Contract *TEEWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TEEWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TEEWalletTransactorSession struct {
	Contract     *TEEWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TEEWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type TEEWalletRaw struct {
	Contract *TEEWallet // Generic contract binding to access the raw methods on
}

// TEEWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TEEWalletCallerRaw struct {
	Contract *TEEWalletCaller // Generic read-only contract binding to access the raw methods on
}

// TEEWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TEEWalletTransactorRaw struct {
	Contract *TEEWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTEEWallet creates a new instance of TEEWallet, bound to a specific deployed contract.
func NewTEEWallet(address common.Address, backend bind.ContractBackend) (*TEEWallet, error) {
	contract, err := bindTEEWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TEEWallet{TEEWalletCaller: TEEWalletCaller{contract: contract}, TEEWalletTransactor: TEEWalletTransactor{contract: contract}, TEEWalletFilterer: TEEWalletFilterer{contract: contract}}, nil
}

// NewTEEWalletCaller creates a new read-only instance of TEEWallet, bound to a specific deployed contract.
func NewTEEWalletCaller(address common.Address, caller bind.ContractCaller) (*TEEWalletCaller, error) {
	contract, err := bindTEEWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TEEWalletCaller{contract: contract}, nil
}

// NewTEEWalletTransactor creates a new write-only instance of TEEWallet, bound to a specific deployed contract.
func NewTEEWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*TEEWalletTransactor, error) {
	contract, err := bindTEEWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TEEWalletTransactor{contract: contract}, nil
}

// NewTEEWalletFilterer creates a new log filterer instance of TEEWallet, bound to a specific deployed contract.
func NewTEEWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*TEEWalletFilterer, error) {
	contract, err := bindTEEWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TEEWalletFilterer{contract: contract}, nil
}

// bindTEEWallet binds a generic wrapper to an already deployed contract.
func bindTEEWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TEEWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TEEWallet *TEEWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TEEWallet.Contract.TEEWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TEEWallet *TEEWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TEEWallet.Contract.TEEWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TEEWallet *TEEWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TEEWallet.Contract.TEEWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TEEWallet *TEEWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TEEWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TEEWallet *TEEWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TEEWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TEEWallet *TEEWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TEEWallet.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_TEEWallet *TEEWalletCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_TEEWallet *TEEWalletSession) ADMINROLE() ([32]byte, error) {
	return _TEEWallet.Contract.ADMINROLE(&_TEEWallet.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_TEEWallet *TEEWalletCallerSession) ADMINROLE() ([32]byte, error) {
	return _TEEWallet.Contract.ADMINROLE(&_TEEWallet.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TEEWallet *TEEWalletCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TEEWallet *TEEWalletSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TEEWallet.Contract.DEFAULTADMINROLE(&_TEEWallet.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TEEWallet *TEEWalletCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TEEWallet.Contract.DEFAULTADMINROLE(&_TEEWallet.CallOpts)
}

// SESSIONADMINROLE is a free data retrieval call binding the contract method 0x6f46b405.
//
// Solidity: function SESSION_ADMIN_ROLE() view returns(bytes32)
func (_TEEWallet *TEEWalletCaller) SESSIONADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "SESSION_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SESSIONADMINROLE is a free data retrieval call binding the contract method 0x6f46b405.
//
// Solidity: function SESSION_ADMIN_ROLE() view returns(bytes32)
func (_TEEWallet *TEEWalletSession) SESSIONADMINROLE() ([32]byte, error) {
	return _TEEWallet.Contract.SESSIONADMINROLE(&_TEEWallet.CallOpts)
}

// SESSIONADMINROLE is a free data retrieval call binding the contract method 0x6f46b405.
//
// Solidity: function SESSION_ADMIN_ROLE() view returns(bytes32)
func (_TEEWallet *TEEWalletCallerSession) SESSIONADMINROLE() ([32]byte, error) {
	return _TEEWallet.Contract.SESSIONADMINROLE(&_TEEWallet.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TEEWallet *TEEWalletCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TEEWallet *TEEWalletSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TEEWallet.Contract.UPGRADEINTERFACEVERSION(&_TEEWallet.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TEEWallet *TEEWalletCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TEEWallet.Contract.UPGRADEINTERFACEVERSION(&_TEEWallet.CallOpts)
}

// EcdsaOnKeccak256r1 is a free data retrieval call binding the contract method 0xce27ac17.
//
// Solidity: function ecdsaOnKeccak256r1(bytes message, bytes signature, bytes publicKey) pure returns(bool)
func (_TEEWallet *TEEWalletCaller) EcdsaOnKeccak256r1(opts *bind.CallOpts, message []byte, signature []byte, publicKey []byte) (bool, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "ecdsaOnKeccak256r1", message, signature, publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EcdsaOnKeccak256r1 is a free data retrieval call binding the contract method 0xce27ac17.
//
// Solidity: function ecdsaOnKeccak256r1(bytes message, bytes signature, bytes publicKey) pure returns(bool)
func (_TEEWallet *TEEWalletSession) EcdsaOnKeccak256r1(message []byte, signature []byte, publicKey []byte) (bool, error) {
	return _TEEWallet.Contract.EcdsaOnKeccak256r1(&_TEEWallet.CallOpts, message, signature, publicKey)
}

// EcdsaOnKeccak256r1 is a free data retrieval call binding the contract method 0xce27ac17.
//
// Solidity: function ecdsaOnKeccak256r1(bytes message, bytes signature, bytes publicKey) pure returns(bool)
func (_TEEWallet *TEEWalletCallerSession) EcdsaOnKeccak256r1(message []byte, signature []byte, publicKey []byte) (bool, error) {
	return _TEEWallet.Contract.EcdsaOnKeccak256r1(&_TEEWallet.CallOpts, message, signature, publicKey)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_TEEWallet *TEEWalletCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_TEEWallet *TEEWalletSession) Fee() (*big.Int, error) {
	return _TEEWallet.Contract.Fee(&_TEEWallet.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_TEEWallet *TEEWalletCallerSession) Fee() (*big.Int, error) {
	return _TEEWallet.Contract.Fee(&_TEEWallet.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TEEWallet *TEEWalletCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TEEWallet *TEEWalletSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TEEWallet.Contract.GetRoleAdmin(&_TEEWallet.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TEEWallet *TEEWalletCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TEEWallet.Contract.GetRoleAdmin(&_TEEWallet.CallOpts, role)
}

// GitHubLink is a free data retrieval call binding the contract method 0x7bffa75a.
//
// Solidity: function gitHubLink() view returns(string)
func (_TEEWallet *TEEWalletCaller) GitHubLink(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "gitHubLink")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GitHubLink is a free data retrieval call binding the contract method 0x7bffa75a.
//
// Solidity: function gitHubLink() view returns(string)
func (_TEEWallet *TEEWalletSession) GitHubLink() (string, error) {
	return _TEEWallet.Contract.GitHubLink(&_TEEWallet.CallOpts)
}

// GitHubLink is a free data retrieval call binding the contract method 0x7bffa75a.
//
// Solidity: function gitHubLink() view returns(string)
func (_TEEWallet *TEEWalletCallerSession) GitHubLink() (string, error) {
	return _TEEWallet.Contract.GitHubLink(&_TEEWallet.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TEEWallet *TEEWalletCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TEEWallet *TEEWalletSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TEEWallet.Contract.HasRole(&_TEEWallet.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TEEWallet *TEEWalletCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TEEWallet.Contract.HasRole(&_TEEWallet.CallOpts, role, account)
}

// NonceToUserData is a free data retrieval call binding the contract method 0xfa3dfc5e.
//
// Solidity: function nonceToUserData(uint256 _nonce) view returns((address,uint256) uData)
func (_TEEWallet *TEEWalletCaller) NonceToUserData(opts *bind.CallOpts, _nonce *big.Int) (ITEEWalletuserData, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "nonceToUserData", _nonce)

	if err != nil {
		return *new(ITEEWalletuserData), err
	}

	out0 := *abi.ConvertType(out[0], new(ITEEWalletuserData)).(*ITEEWalletuserData)

	return out0, err

}

// NonceToUserData is a free data retrieval call binding the contract method 0xfa3dfc5e.
//
// Solidity: function nonceToUserData(uint256 _nonce) view returns((address,uint256) uData)
func (_TEEWallet *TEEWalletSession) NonceToUserData(_nonce *big.Int) (ITEEWalletuserData, error) {
	return _TEEWallet.Contract.NonceToUserData(&_TEEWallet.CallOpts, _nonce)
}

// NonceToUserData is a free data retrieval call binding the contract method 0xfa3dfc5e.
//
// Solidity: function nonceToUserData(uint256 _nonce) view returns((address,uint256) uData)
func (_TEEWallet *TEEWalletCallerSession) NonceToUserData(_nonce *big.Int) (ITEEWalletuserData, error) {
	return _TEEWallet.Contract.NonceToUserData(&_TEEWallet.CallOpts, _nonce)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TEEWallet *TEEWalletCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TEEWallet *TEEWalletSession) ProxiableUUID() ([32]byte, error) {
	return _TEEWallet.Contract.ProxiableUUID(&_TEEWallet.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TEEWallet *TEEWalletCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TEEWallet.Contract.ProxiableUUID(&_TEEWallet.CallOpts)
}

// RootCert is a free data retrieval call binding the contract method 0x7221e6fc.
//
// Solidity: function rootCert() view returns((bytes,bytes,bytes,bytes) cert)
func (_TEEWallet *TEEWalletCaller) RootCert(opts *bind.CallOpts) (ChunkedX509Cert, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "rootCert")

	if err != nil {
		return *new(ChunkedX509Cert), err
	}

	out0 := *abi.ConvertType(out[0], new(ChunkedX509Cert)).(*ChunkedX509Cert)

	return out0, err

}

// RootCert is a free data retrieval call binding the contract method 0x7221e6fc.
//
// Solidity: function rootCert() view returns((bytes,bytes,bytes,bytes) cert)
func (_TEEWallet *TEEWalletSession) RootCert() (ChunkedX509Cert, error) {
	return _TEEWallet.Contract.RootCert(&_TEEWallet.CallOpts)
}

// RootCert is a free data retrieval call binding the contract method 0x7221e6fc.
//
// Solidity: function rootCert() view returns((bytes,bytes,bytes,bytes) cert)
func (_TEEWallet *TEEWalletCallerSession) RootCert() (ChunkedX509Cert, error) {
	return _TEEWallet.Contract.RootCert(&_TEEWallet.CallOpts)
}

// ScriptHash is a free data retrieval call binding the contract method 0x86ba212d.
//
// Solidity: function scriptHash() view returns(bytes32)
func (_TEEWallet *TEEWalletCaller) ScriptHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "scriptHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ScriptHash is a free data retrieval call binding the contract method 0x86ba212d.
//
// Solidity: function scriptHash() view returns(bytes32)
func (_TEEWallet *TEEWalletSession) ScriptHash() ([32]byte, error) {
	return _TEEWallet.Contract.ScriptHash(&_TEEWallet.CallOpts)
}

// ScriptHash is a free data retrieval call binding the contract method 0x86ba212d.
//
// Solidity: function scriptHash() view returns(bytes32)
func (_TEEWallet *TEEWalletCallerSession) ScriptHash() ([32]byte, error) {
	return _TEEWallet.Contract.ScriptHash(&_TEEWallet.CallOpts)
}

// Sessions is a free data retrieval call binding the contract method 0x431a1b97.
//
// Solidity: function sessions(address sessionKey) view returns(bool)
func (_TEEWallet *TEEWalletCaller) Sessions(opts *bind.CallOpts, sessionKey common.Address) (bool, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "sessions", sessionKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Sessions is a free data retrieval call binding the contract method 0x431a1b97.
//
// Solidity: function sessions(address sessionKey) view returns(bool)
func (_TEEWallet *TEEWalletSession) Sessions(sessionKey common.Address) (bool, error) {
	return _TEEWallet.Contract.Sessions(&_TEEWallet.CallOpts, sessionKey)
}

// Sessions is a free data retrieval call binding the contract method 0x431a1b97.
//
// Solidity: function sessions(address sessionKey) view returns(bool)
func (_TEEWallet *TEEWalletCallerSession) Sessions(sessionKey common.Address) (bool, error) {
	return _TEEWallet.Contract.Sessions(&_TEEWallet.CallOpts, sessionKey)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TEEWallet *TEEWalletCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TEEWallet *TEEWalletSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TEEWallet.Contract.SupportsInterface(&_TEEWallet.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TEEWallet *TEEWalletCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TEEWallet.Contract.SupportsInterface(&_TEEWallet.CallOpts, interfaceId)
}

// VerifyCert is a free data retrieval call binding the contract method 0xbee94f11.
//
// Solidity: function verifyCert((bytes,bytes,bytes,bytes) cert, bytes publicKey) pure returns(bool)
func (_TEEWallet *TEEWalletCaller) VerifyCert(opts *bind.CallOpts, cert ChunkedX509Cert, publicKey []byte) (bool, error) {
	var out []interface{}
	err := _TEEWallet.contract.Call(opts, &out, "verifyCert", cert, publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCert is a free data retrieval call binding the contract method 0xbee94f11.
//
// Solidity: function verifyCert((bytes,bytes,bytes,bytes) cert, bytes publicKey) pure returns(bool)
func (_TEEWallet *TEEWalletSession) VerifyCert(cert ChunkedX509Cert, publicKey []byte) (bool, error) {
	return _TEEWallet.Contract.VerifyCert(&_TEEWallet.CallOpts, cert, publicKey)
}

// VerifyCert is a free data retrieval call binding the contract method 0xbee94f11.
//
// Solidity: function verifyCert((bytes,bytes,bytes,bytes) cert, bytes publicKey) pure returns(bool)
func (_TEEWallet *TEEWalletCallerSession) VerifyCert(cert ChunkedX509Cert, publicKey []byte) (bool, error) {
	return _TEEWallet.Contract.VerifyCert(&_TEEWallet.CallOpts, cert, publicKey)
}

// Execute is a paid mutator transaction binding the contract method 0x87614bb9.
//
// Solidity: function execute((address,uint256,bytes,bool)[] txs) payable returns(bytes[] results)
func (_TEEWallet *TEEWalletTransactor) Execute(opts *bind.TransactOpts, txs []Transaction) (*types.Transaction, error) {
	return _TEEWallet.contract.Transact(opts, "execute", txs)
}

// Execute is a paid mutator transaction binding the contract method 0x87614bb9.
//
// Solidity: function execute((address,uint256,bytes,bool)[] txs) payable returns(bytes[] results)
func (_TEEWallet *TEEWalletSession) Execute(txs []Transaction) (*types.Transaction, error) {
	return _TEEWallet.Contract.Execute(&_TEEWallet.TransactOpts, txs)
}

// Execute is a paid mutator transaction binding the contract method 0x87614bb9.
//
// Solidity: function execute((address,uint256,bytes,bool)[] txs) payable returns(bytes[] results)
func (_TEEWallet *TEEWalletTransactorSession) Execute(txs []Transaction) (*types.Transaction, error) {
	return _TEEWallet.Contract.Execute(&_TEEWallet.TransactOpts, txs)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TEEWallet *TEEWalletTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TEEWallet.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TEEWallet *TEEWalletSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TEEWallet.Contract.GrantRole(&_TEEWallet.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TEEWallet *TEEWalletTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TEEWallet.Contract.GrantRole(&_TEEWallet.TransactOpts, role, account)
}

// InitSessionKey is a paid mutator transaction binding the contract method 0x4c7c04c4.
//
// Solidity: function initSessionKey((bytes,bytes,bytes,bytes) leaf, (bytes,bytes,bytes,bytes) intermediate, (bytes,bytes,bytes,bytes,bytes,bytes,bytes) quote, address sessionKey) returns()
func (_TEEWallet *TEEWalletTransactor) InitSessionKey(opts *bind.TransactOpts, leaf ChunkedX509Cert, intermediate ChunkedX509Cert, quote ChunkedSGXQuote, sessionKey common.Address) (*types.Transaction, error) {
	return _TEEWallet.contract.Transact(opts, "initSessionKey", leaf, intermediate, quote, sessionKey)
}

// InitSessionKey is a paid mutator transaction binding the contract method 0x4c7c04c4.
//
// Solidity: function initSessionKey((bytes,bytes,bytes,bytes) leaf, (bytes,bytes,bytes,bytes) intermediate, (bytes,bytes,bytes,bytes,bytes,bytes,bytes) quote, address sessionKey) returns()
func (_TEEWallet *TEEWalletSession) InitSessionKey(leaf ChunkedX509Cert, intermediate ChunkedX509Cert, quote ChunkedSGXQuote, sessionKey common.Address) (*types.Transaction, error) {
	return _TEEWallet.Contract.InitSessionKey(&_TEEWallet.TransactOpts, leaf, intermediate, quote, sessionKey)
}

// InitSessionKey is a paid mutator transaction binding the contract method 0x4c7c04c4.
//
// Solidity: function initSessionKey((bytes,bytes,bytes,bytes) leaf, (bytes,bytes,bytes,bytes) intermediate, (bytes,bytes,bytes,bytes,bytes,bytes,bytes) quote, address sessionKey) returns()
func (_TEEWallet *TEEWalletTransactorSession) InitSessionKey(leaf ChunkedX509Cert, intermediate ChunkedX509Cert, quote ChunkedSGXQuote, sessionKey common.Address) (*types.Transaction, error) {
	return _TEEWallet.Contract.InitSessionKey(&_TEEWallet.TransactOpts, leaf, intermediate, quote, sessionKey)
}

// Initialize is a paid mutator transaction binding the contract method 0xec3a0860.
//
// Solidity: function initialize(address admin, address sessionAdmin, (bytes,bytes,bytes,bytes) _rootCert) returns()
func (_TEEWallet *TEEWalletTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, sessionAdmin common.Address, _rootCert ChunkedX509Cert) (*types.Transaction, error) {
	return _TEEWallet.contract.Transact(opts, "initialize", admin, sessionAdmin, _rootCert)
}

// Initialize is a paid mutator transaction binding the contract method 0xec3a0860.
//
// Solidity: function initialize(address admin, address sessionAdmin, (bytes,bytes,bytes,bytes) _rootCert) returns()
func (_TEEWallet *TEEWalletSession) Initialize(admin common.Address, sessionAdmin common.Address, _rootCert ChunkedX509Cert) (*types.Transaction, error) {
	return _TEEWallet.Contract.Initialize(&_TEEWallet.TransactOpts, admin, sessionAdmin, _rootCert)
}

// Initialize is a paid mutator transaction binding the contract method 0xec3a0860.
//
// Solidity: function initialize(address admin, address sessionAdmin, (bytes,bytes,bytes,bytes) _rootCert) returns()
func (_TEEWallet *TEEWalletTransactorSession) Initialize(admin common.Address, sessionAdmin common.Address, _rootCert ChunkedX509Cert) (*types.Transaction, error) {
	return _TEEWallet.Contract.Initialize(&_TEEWallet.TransactOpts, admin, sessionAdmin, _rootCert)
}

// Refund is a paid mutator transaction binding the contract method 0x22ab1334.
//
// Solidity: function refund(uint256 _nonce, bytes _signature) payable returns()
func (_TEEWallet *TEEWalletTransactor) Refund(opts *bind.TransactOpts, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _TEEWallet.contract.Transact(opts, "refund", _nonce, _signature)
}

// Refund is a paid mutator transaction binding the contract method 0x22ab1334.
//
// Solidity: function refund(uint256 _nonce, bytes _signature) payable returns()
func (_TEEWallet *TEEWalletSession) Refund(_nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _TEEWallet.Contract.Refund(&_TEEWallet.TransactOpts, _nonce, _signature)
}

// Refund is a paid mutator transaction binding the contract method 0x22ab1334.
//
// Solidity: function refund(uint256 _nonce, bytes _signature) payable returns()
func (_TEEWallet *TEEWalletTransactorSession) Refund(_nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _TEEWallet.Contract.Refund(&_TEEWallet.TransactOpts, _nonce, _signature)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TEEWallet *TEEWalletTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TEEWallet.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TEEWallet *TEEWalletSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TEEWallet.Contract.RenounceRole(&_TEEWallet.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TEEWallet *TEEWalletTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TEEWallet.Contract.RenounceRole(&_TEEWallet.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TEEWallet *TEEWalletTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TEEWallet.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TEEWallet *TEEWalletSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TEEWallet.Contract.RevokeRole(&_TEEWallet.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TEEWallet *TEEWalletTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TEEWallet.Contract.RevokeRole(&_TEEWallet.TransactOpts, role, account)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _newFee) returns()
func (_TEEWallet *TEEWalletTransactor) SetFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _TEEWallet.contract.Transact(opts, "setFee", _newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _newFee) returns()
func (_TEEWallet *TEEWalletSession) SetFee(_newFee *big.Int) (*types.Transaction, error) {
	return _TEEWallet.Contract.SetFee(&_TEEWallet.TransactOpts, _newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _newFee) returns()
func (_TEEWallet *TEEWalletTransactorSession) SetFee(_newFee *big.Int) (*types.Transaction, error) {
	return _TEEWallet.Contract.SetFee(&_TEEWallet.TransactOpts, _newFee)
}

// SetRootCert is a paid mutator transaction binding the contract method 0x0642e02f.
//
// Solidity: function setRootCert((bytes,bytes,bytes,bytes) _newRootCert) returns()
func (_TEEWallet *TEEWalletTransactor) SetRootCert(opts *bind.TransactOpts, _newRootCert ChunkedX509Cert) (*types.Transaction, error) {
	return _TEEWallet.contract.Transact(opts, "setRootCert", _newRootCert)
}

// SetRootCert is a paid mutator transaction binding the contract method 0x0642e02f.
//
// Solidity: function setRootCert((bytes,bytes,bytes,bytes) _newRootCert) returns()
func (_TEEWallet *TEEWalletSession) SetRootCert(_newRootCert ChunkedX509Cert) (*types.Transaction, error) {
	return _TEEWallet.Contract.SetRootCert(&_TEEWallet.TransactOpts, _newRootCert)
}

// SetRootCert is a paid mutator transaction binding the contract method 0x0642e02f.
//
// Solidity: function setRootCert((bytes,bytes,bytes,bytes) _newRootCert) returns()
func (_TEEWallet *TEEWalletTransactorSession) SetRootCert(_newRootCert ChunkedX509Cert) (*types.Transaction, error) {
	return _TEEWallet.Contract.SetRootCert(&_TEEWallet.TransactOpts, _newRootCert)
}

// SetScript is a paid mutator transaction binding the contract method 0x87a4ddd1.
//
// Solidity: function setScript(bytes32 scriptHash, string gitHubLinkPublic) returns()
func (_TEEWallet *TEEWalletTransactor) SetScript(opts *bind.TransactOpts, scriptHash [32]byte, gitHubLinkPublic string) (*types.Transaction, error) {
	return _TEEWallet.contract.Transact(opts, "setScript", scriptHash, gitHubLinkPublic)
}

// SetScript is a paid mutator transaction binding the contract method 0x87a4ddd1.
//
// Solidity: function setScript(bytes32 scriptHash, string gitHubLinkPublic) returns()
func (_TEEWallet *TEEWalletSession) SetScript(scriptHash [32]byte, gitHubLinkPublic string) (*types.Transaction, error) {
	return _TEEWallet.Contract.SetScript(&_TEEWallet.TransactOpts, scriptHash, gitHubLinkPublic)
}

// SetScript is a paid mutator transaction binding the contract method 0x87a4ddd1.
//
// Solidity: function setScript(bytes32 scriptHash, string gitHubLinkPublic) returns()
func (_TEEWallet *TEEWalletTransactorSession) SetScript(scriptHash [32]byte, gitHubLinkPublic string) (*types.Transaction, error) {
	return _TEEWallet.Contract.SetScript(&_TEEWallet.TransactOpts, scriptHash, gitHubLinkPublic)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TEEWallet *TEEWalletTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TEEWallet.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TEEWallet *TEEWalletSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TEEWallet.Contract.UpgradeToAndCall(&_TEEWallet.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TEEWallet *TEEWalletTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TEEWallet.Contract.UpgradeToAndCall(&_TEEWallet.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TEEWallet *TEEWalletTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TEEWallet.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TEEWallet *TEEWalletSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TEEWallet.Contract.Fallback(&_TEEWallet.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TEEWallet *TEEWalletTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TEEWallet.Contract.Fallback(&_TEEWallet.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TEEWallet *TEEWalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TEEWallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TEEWallet *TEEWalletSession) Receive() (*types.Transaction, error) {
	return _TEEWallet.Contract.Receive(&_TEEWallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TEEWallet *TEEWalletTransactorSession) Receive() (*types.Transaction, error) {
	return _TEEWallet.Contract.Receive(&_TEEWallet.TransactOpts)
}

// TEEWalletExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the TEEWallet contract.
type TEEWalletExecutedIterator struct {
	Event *TEEWalletExecuted // Event containing the contract specifics and raw log

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
func (it *TEEWalletExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TEEWalletExecuted)
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
		it.Event = new(TEEWalletExecuted)
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
func (it *TEEWalletExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TEEWalletExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TEEWalletExecuted represents a Executed event raised by the TEEWallet contract.
type TEEWalletExecuted struct {
	Target   common.Address
	Caller   common.Address
	Data     []byte
	Response []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x6f048d58d8eae4b8db180e5334619d0da212adcbc67ef49bfe1fbae2cb4f33b2.
//
// Solidity: event Executed(address indexed target, address indexed caller, bytes data, bytes response)
func (_TEEWallet *TEEWalletFilterer) FilterExecuted(opts *bind.FilterOpts, target []common.Address, caller []common.Address) (*TEEWalletExecutedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TEEWallet.contract.FilterLogs(opts, "Executed", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &TEEWalletExecutedIterator{contract: _TEEWallet.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x6f048d58d8eae4b8db180e5334619d0da212adcbc67ef49bfe1fbae2cb4f33b2.
//
// Solidity: event Executed(address indexed target, address indexed caller, bytes data, bytes response)
func (_TEEWallet *TEEWalletFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *TEEWalletExecuted, target []common.Address, caller []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TEEWallet.contract.WatchLogs(opts, "Executed", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TEEWalletExecuted)
				if err := _TEEWallet.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x6f048d58d8eae4b8db180e5334619d0da212adcbc67ef49bfe1fbae2cb4f33b2.
//
// Solidity: event Executed(address indexed target, address indexed caller, bytes data, bytes response)
func (_TEEWallet *TEEWalletFilterer) ParseExecuted(log types.Log) (*TEEWalletExecuted, error) {
	event := new(TEEWalletExecuted)
	if err := _TEEWallet.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TEEWalletFundsReceivedIterator is returned from FilterFundsReceived and is used to iterate over the raw logs and unpacked data for FundsReceived events raised by the TEEWallet contract.
type TEEWalletFundsReceivedIterator struct {
	Event *TEEWalletFundsReceived // Event containing the contract specifics and raw log

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
func (it *TEEWalletFundsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TEEWalletFundsReceived)
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
		it.Event = new(TEEWalletFundsReceived)
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
func (it *TEEWalletFundsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TEEWalletFundsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TEEWalletFundsReceived represents a FundsReceived event raised by the TEEWallet contract.
type TEEWalletFundsReceived struct {
	Nonce     *big.Int
	Msgsender common.Address
	MsgValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFundsReceived is a free log retrieval operation binding the contract event 0xfb932f39cc5ea177af75a4f1502d59d13404be49a2edb139943b43e4fb3fe45b.
//
// Solidity: event FundsReceived(uint256 nonce, address msgsender, uint256 msgValue)
func (_TEEWallet *TEEWalletFilterer) FilterFundsReceived(opts *bind.FilterOpts) (*TEEWalletFundsReceivedIterator, error) {

	logs, sub, err := _TEEWallet.contract.FilterLogs(opts, "FundsReceived")
	if err != nil {
		return nil, err
	}
	return &TEEWalletFundsReceivedIterator{contract: _TEEWallet.contract, event: "FundsReceived", logs: logs, sub: sub}, nil
}

// WatchFundsReceived is a free log subscription operation binding the contract event 0xfb932f39cc5ea177af75a4f1502d59d13404be49a2edb139943b43e4fb3fe45b.
//
// Solidity: event FundsReceived(uint256 nonce, address msgsender, uint256 msgValue)
func (_TEEWallet *TEEWalletFilterer) WatchFundsReceived(opts *bind.WatchOpts, sink chan<- *TEEWalletFundsReceived) (event.Subscription, error) {

	logs, sub, err := _TEEWallet.contract.WatchLogs(opts, "FundsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TEEWalletFundsReceived)
				if err := _TEEWallet.contract.UnpackLog(event, "FundsReceived", log); err != nil {
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

// ParseFundsReceived is a log parse operation binding the contract event 0xfb932f39cc5ea177af75a4f1502d59d13404be49a2edb139943b43e4fb3fe45b.
//
// Solidity: event FundsReceived(uint256 nonce, address msgsender, uint256 msgValue)
func (_TEEWallet *TEEWalletFilterer) ParseFundsReceived(log types.Log) (*TEEWalletFundsReceived, error) {
	event := new(TEEWalletFundsReceived)
	if err := _TEEWallet.contract.UnpackLog(event, "FundsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TEEWalletInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TEEWallet contract.
type TEEWalletInitializedIterator struct {
	Event *TEEWalletInitialized // Event containing the contract specifics and raw log

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
func (it *TEEWalletInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TEEWalletInitialized)
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
		it.Event = new(TEEWalletInitialized)
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
func (it *TEEWalletInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TEEWalletInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TEEWalletInitialized represents a Initialized event raised by the TEEWallet contract.
type TEEWalletInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TEEWallet *TEEWalletFilterer) FilterInitialized(opts *bind.FilterOpts) (*TEEWalletInitializedIterator, error) {

	logs, sub, err := _TEEWallet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TEEWalletInitializedIterator{contract: _TEEWallet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TEEWallet *TEEWalletFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TEEWalletInitialized) (event.Subscription, error) {

	logs, sub, err := _TEEWallet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TEEWalletInitialized)
				if err := _TEEWallet.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TEEWallet *TEEWalletFilterer) ParseInitialized(log types.Log) (*TEEWalletInitialized, error) {
	event := new(TEEWalletInitialized)
	if err := _TEEWallet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TEEWalletRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the TEEWallet contract.
type TEEWalletRefundIterator struct {
	Event *TEEWalletRefund // Event containing the contract specifics and raw log

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
func (it *TEEWalletRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TEEWalletRefund)
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
		it.Event = new(TEEWalletRefund)
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
func (it *TEEWalletRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TEEWalletRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TEEWalletRefund represents a Refund event raised by the TEEWallet contract.
type TEEWalletRefund struct {
	Nonce    *big.Int
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x21e12a7cad0da5928167e1084ea4d5fdf8d9af66657a2543a9ac76a0ca081477.
//
// Solidity: event Refund(uint256 nonce, address receiver, uint256 amount)
func (_TEEWallet *TEEWalletFilterer) FilterRefund(opts *bind.FilterOpts) (*TEEWalletRefundIterator, error) {

	logs, sub, err := _TEEWallet.contract.FilterLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return &TEEWalletRefundIterator{contract: _TEEWallet.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x21e12a7cad0da5928167e1084ea4d5fdf8d9af66657a2543a9ac76a0ca081477.
//
// Solidity: event Refund(uint256 nonce, address receiver, uint256 amount)
func (_TEEWallet *TEEWalletFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *TEEWalletRefund) (event.Subscription, error) {

	logs, sub, err := _TEEWallet.contract.WatchLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TEEWalletRefund)
				if err := _TEEWallet.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0x21e12a7cad0da5928167e1084ea4d5fdf8d9af66657a2543a9ac76a0ca081477.
//
// Solidity: event Refund(uint256 nonce, address receiver, uint256 amount)
func (_TEEWallet *TEEWalletFilterer) ParseRefund(log types.Log) (*TEEWalletRefund, error) {
	event := new(TEEWalletRefund)
	if err := _TEEWallet.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TEEWalletRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TEEWallet contract.
type TEEWalletRoleAdminChangedIterator struct {
	Event *TEEWalletRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TEEWalletRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TEEWalletRoleAdminChanged)
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
		it.Event = new(TEEWalletRoleAdminChanged)
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
func (it *TEEWalletRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TEEWalletRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TEEWalletRoleAdminChanged represents a RoleAdminChanged event raised by the TEEWallet contract.
type TEEWalletRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TEEWallet *TEEWalletFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TEEWalletRoleAdminChangedIterator, error) {

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

	logs, sub, err := _TEEWallet.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TEEWalletRoleAdminChangedIterator{contract: _TEEWallet.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TEEWallet *TEEWalletFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TEEWalletRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TEEWallet.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TEEWalletRoleAdminChanged)
				if err := _TEEWallet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TEEWallet *TEEWalletFilterer) ParseRoleAdminChanged(log types.Log) (*TEEWalletRoleAdminChanged, error) {
	event := new(TEEWalletRoleAdminChanged)
	if err := _TEEWallet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TEEWalletRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TEEWallet contract.
type TEEWalletRoleGrantedIterator struct {
	Event *TEEWalletRoleGranted // Event containing the contract specifics and raw log

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
func (it *TEEWalletRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TEEWalletRoleGranted)
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
		it.Event = new(TEEWalletRoleGranted)
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
func (it *TEEWalletRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TEEWalletRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TEEWalletRoleGranted represents a RoleGranted event raised by the TEEWallet contract.
type TEEWalletRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TEEWallet *TEEWalletFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TEEWalletRoleGrantedIterator, error) {

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

	logs, sub, err := _TEEWallet.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TEEWalletRoleGrantedIterator{contract: _TEEWallet.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TEEWallet *TEEWalletFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TEEWalletRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TEEWallet.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TEEWalletRoleGranted)
				if err := _TEEWallet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TEEWallet *TEEWalletFilterer) ParseRoleGranted(log types.Log) (*TEEWalletRoleGranted, error) {
	event := new(TEEWalletRoleGranted)
	if err := _TEEWallet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TEEWalletRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TEEWallet contract.
type TEEWalletRoleRevokedIterator struct {
	Event *TEEWalletRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TEEWalletRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TEEWalletRoleRevoked)
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
		it.Event = new(TEEWalletRoleRevoked)
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
func (it *TEEWalletRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TEEWalletRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TEEWalletRoleRevoked represents a RoleRevoked event raised by the TEEWallet contract.
type TEEWalletRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TEEWallet *TEEWalletFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TEEWalletRoleRevokedIterator, error) {

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

	logs, sub, err := _TEEWallet.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TEEWalletRoleRevokedIterator{contract: _TEEWallet.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TEEWallet *TEEWalletFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TEEWalletRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TEEWallet.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TEEWalletRoleRevoked)
				if err := _TEEWallet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TEEWallet *TEEWalletFilterer) ParseRoleRevoked(log types.Log) (*TEEWalletRoleRevoked, error) {
	event := new(TEEWalletRoleRevoked)
	if err := _TEEWallet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TEEWalletScriptSettedIterator is returned from FilterScriptSetted and is used to iterate over the raw logs and unpacked data for ScriptSetted events raised by the TEEWallet contract.
type TEEWalletScriptSettedIterator struct {
	Event *TEEWalletScriptSetted // Event containing the contract specifics and raw log

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
func (it *TEEWalletScriptSettedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TEEWalletScriptSetted)
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
		it.Event = new(TEEWalletScriptSetted)
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
func (it *TEEWalletScriptSettedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TEEWalletScriptSettedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TEEWalletScriptSetted represents a ScriptSetted event raised by the TEEWallet contract.
type TEEWalletScriptSetted struct {
	ScriptHash       [32]byte
	GitHubLinkPublic string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterScriptSetted is a free log retrieval operation binding the contract event 0x921f2a6ac928f09d07e6d7a36bae6cfb5f3467ca4158136adceba5c11512032c.
//
// Solidity: event ScriptSetted(bytes32 scriptHash, string gitHubLinkPublic)
func (_TEEWallet *TEEWalletFilterer) FilterScriptSetted(opts *bind.FilterOpts) (*TEEWalletScriptSettedIterator, error) {

	logs, sub, err := _TEEWallet.contract.FilterLogs(opts, "ScriptSetted")
	if err != nil {
		return nil, err
	}
	return &TEEWalletScriptSettedIterator{contract: _TEEWallet.contract, event: "ScriptSetted", logs: logs, sub: sub}, nil
}

// WatchScriptSetted is a free log subscription operation binding the contract event 0x921f2a6ac928f09d07e6d7a36bae6cfb5f3467ca4158136adceba5c11512032c.
//
// Solidity: event ScriptSetted(bytes32 scriptHash, string gitHubLinkPublic)
func (_TEEWallet *TEEWalletFilterer) WatchScriptSetted(opts *bind.WatchOpts, sink chan<- *TEEWalletScriptSetted) (event.Subscription, error) {

	logs, sub, err := _TEEWallet.contract.WatchLogs(opts, "ScriptSetted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TEEWalletScriptSetted)
				if err := _TEEWallet.contract.UnpackLog(event, "ScriptSetted", log); err != nil {
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

// ParseScriptSetted is a log parse operation binding the contract event 0x921f2a6ac928f09d07e6d7a36bae6cfb5f3467ca4158136adceba5c11512032c.
//
// Solidity: event ScriptSetted(bytes32 scriptHash, string gitHubLinkPublic)
func (_TEEWallet *TEEWalletFilterer) ParseScriptSetted(log types.Log) (*TEEWalletScriptSetted, error) {
	event := new(TEEWalletScriptSetted)
	if err := _TEEWallet.contract.UnpackLog(event, "ScriptSetted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TEEWalletSessionInitializedIterator is returned from FilterSessionInitialized and is used to iterate over the raw logs and unpacked data for SessionInitialized events raised by the TEEWallet contract.
type TEEWalletSessionInitializedIterator struct {
	Event *TEEWalletSessionInitialized // Event containing the contract specifics and raw log

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
func (it *TEEWalletSessionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TEEWalletSessionInitialized)
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
		it.Event = new(TEEWalletSessionInitialized)
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
func (it *TEEWalletSessionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TEEWalletSessionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TEEWalletSessionInitialized represents a SessionInitialized event raised by the TEEWallet contract.
type TEEWalletSessionInitialized struct {
	SessionKey     common.Address
	MrEnclave      [32]byte
	ReportDataHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSessionInitialized is a free log retrieval operation binding the contract event 0x170ab4f1c291451427dd313e007f2d10e6bc15876e4cbdc09044dc5d73774136.
//
// Solidity: event SessionInitialized(address sessionKey, bytes32 _mrEnclave, bytes32 reportDataHash)
func (_TEEWallet *TEEWalletFilterer) FilterSessionInitialized(opts *bind.FilterOpts) (*TEEWalletSessionInitializedIterator, error) {

	logs, sub, err := _TEEWallet.contract.FilterLogs(opts, "SessionInitialized")
	if err != nil {
		return nil, err
	}
	return &TEEWalletSessionInitializedIterator{contract: _TEEWallet.contract, event: "SessionInitialized", logs: logs, sub: sub}, nil
}

// WatchSessionInitialized is a free log subscription operation binding the contract event 0x170ab4f1c291451427dd313e007f2d10e6bc15876e4cbdc09044dc5d73774136.
//
// Solidity: event SessionInitialized(address sessionKey, bytes32 _mrEnclave, bytes32 reportDataHash)
func (_TEEWallet *TEEWalletFilterer) WatchSessionInitialized(opts *bind.WatchOpts, sink chan<- *TEEWalletSessionInitialized) (event.Subscription, error) {

	logs, sub, err := _TEEWallet.contract.WatchLogs(opts, "SessionInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TEEWalletSessionInitialized)
				if err := _TEEWallet.contract.UnpackLog(event, "SessionInitialized", log); err != nil {
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

// ParseSessionInitialized is a log parse operation binding the contract event 0x170ab4f1c291451427dd313e007f2d10e6bc15876e4cbdc09044dc5d73774136.
//
// Solidity: event SessionInitialized(address sessionKey, bytes32 _mrEnclave, bytes32 reportDataHash)
func (_TEEWallet *TEEWalletFilterer) ParseSessionInitialized(log types.Log) (*TEEWalletSessionInitialized, error) {
	event := new(TEEWalletSessionInitialized)
	if err := _TEEWallet.contract.UnpackLog(event, "SessionInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TEEWalletUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TEEWallet contract.
type TEEWalletUpgradedIterator struct {
	Event *TEEWalletUpgraded // Event containing the contract specifics and raw log

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
func (it *TEEWalletUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TEEWalletUpgraded)
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
		it.Event = new(TEEWalletUpgraded)
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
func (it *TEEWalletUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TEEWalletUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TEEWalletUpgraded represents a Upgraded event raised by the TEEWallet contract.
type TEEWalletUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TEEWallet *TEEWalletFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TEEWalletUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TEEWallet.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TEEWalletUpgradedIterator{contract: _TEEWallet.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TEEWallet *TEEWalletFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TEEWalletUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TEEWallet.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TEEWalletUpgraded)
				if err := _TEEWallet.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TEEWallet *TEEWalletFilterer) ParseUpgraded(log types.Log) (*TEEWalletUpgraded, error) {
	event := new(TEEWalletUpgraded)
	if err := _TEEWallet.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
