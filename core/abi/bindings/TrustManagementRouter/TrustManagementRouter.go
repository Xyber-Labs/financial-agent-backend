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

// Deposit is an auto generated low-level Go binding around an user-defined struct.
type Deposit struct {
	Amount      *big.Int
	LockedUntil *big.Int
}

// Permit is an auto generated low-level Go binding around an user-defined struct.
type Permit struct {
	Holder   common.Address
	Amount   *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

// Transaction is an auto generated low-level Go binding around an user-defined struct.
type Transaction struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
}

// TrustManagementRouterMetaData contains all meta data concerning the TrustManagementRouter contract.
var TrustManagementRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletBeacon\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2EmptyBytecode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Create2FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"Create2InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EllipticCurve__InvalidNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeVerifier__InvalidEnclaveSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeVerifier__InvalidIntermediateCert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeVerifier__InvalidLeafCert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeVerifier__InvalidQeReportDataHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeVerifier__InvalidQeSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeVerifier__ReportMustBe384Bytes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeVerifier__SignatureMustBe64Bytes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XyberTrustManagementRouter__DeadlineExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"XyberTrustManagementRouter__ExecuteCallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perfomanceFeeRate\",\"type\":\"uint256\"}],\"name\":\"XyberTrustManagementRouter__FeeRateExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"XyberTrustManagementRouter__InvalidCallTarget\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"XyberTrustManagementRouter__InvalidCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedData\",\"type\":\"bytes32\"}],\"name\":\"XyberTrustManagementRouter__InvalidHashedData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"XyberTrustManagementRouter__InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"XyberTrustManagementRouter__UnallowedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mrEnclave\",\"type\":\"bytes32\"}],\"name\":\"XyberTrustManagementRouter__UnauthorizedMrEnclave\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XyberTrustManagementRouter__UsedSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XyberTrustManagementRouter__WalletCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XyberTrustManagementRouter__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XyberTrustManagementRouter__ZeroSessionKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XyberUpgradeChecker__E0\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AllowedTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"FeeCollectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mrEnclave\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAuthorized\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"MrEnclaveSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPerfomanceFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"PerformanceFeeRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RootCertSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mrEnclave\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hashedData\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"SessionInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"WalletCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SESSION_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WALLET_BEACON\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"allowedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"createWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structPermit\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"depositWithPermit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"ecdsaOnKeccak256r1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction[]\",\"name\":\"txs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeCollectorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDeposits\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"}],\"internalType\":\"structDeposit[]\",\"name\":\"deposits\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getWalletAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDeployed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getWithdrawableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawableAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getWithdrawableDeposits\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"withdrawableDepositsIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"bodyPartOne\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bodyPartTwo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedX509Cert\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"bodyPartOne\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bodyPartTwo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedX509Cert\",\"name\":\"intermediate\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"isvReport\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"isvReportSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestationKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"qeReport\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"qeReportSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"qeAuthenticationData\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedSGXQuote\",\"name\":\"quote\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"initSessionKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"bodyPartOne\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bodyPartTwo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedX509Cert\",\"name\":\"rootCert\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mrEnclave\",\"type\":\"bytes32\"}],\"name\":\"mrEnclaveAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isAuthorized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"perfomanceFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"performanceFeeRateValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rootCert\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"bodyPartOne\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bodyPartTwo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedX509Cert\",\"name\":\"rootCertificate\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"sessions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"setAllowedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mrEnclave\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isAuthorized\",\"type\":\"bool\"}],\"name\":\"setMrEnclave\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPerfomanceFeeRate\",\"type\":\"uint256\"}],\"name\":\"setPerfomanceFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"totalDeposited\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalDepositedTokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signatureHash\",\"type\":\"bytes32\"}],\"name\":\"usedSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isUsed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"bodyPartOne\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bodyPartTwo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structChunkedX509Cert\",\"name\":\"cert\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"verifyCert\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"depositIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithYield\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xyberContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"contractName\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_TrustManagementRouter *TrustManagementRouterCaller) ETHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "ETH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_TrustManagementRouter *TrustManagementRouterSession) ETHADDRESS() (common.Address, error) {
	return _TrustManagementRouter.Contract.ETHADDRESS(&_TrustManagementRouter.CallOpts)
}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) ETHADDRESS() (common.Address, error) {
	return _TrustManagementRouter.Contract.ETHADDRESS(&_TrustManagementRouter.CallOpts)
}

// SESSIONADMINROLE is a free data retrieval call binding the contract method 0x6f46b405.
//
// Solidity: function SESSION_ADMIN_ROLE() view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterCaller) SESSIONADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "SESSION_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SESSIONADMINROLE is a free data retrieval call binding the contract method 0x6f46b405.
//
// Solidity: function SESSION_ADMIN_ROLE() view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterSession) SESSIONADMINROLE() ([32]byte, error) {
	return _TrustManagementRouter.Contract.SESSIONADMINROLE(&_TrustManagementRouter.CallOpts)
}

// SESSIONADMINROLE is a free data retrieval call binding the contract method 0x6f46b405.
//
// Solidity: function SESSION_ADMIN_ROLE() view returns(bytes32)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) SESSIONADMINROLE() ([32]byte, error) {
	return _TrustManagementRouter.Contract.SESSIONADMINROLE(&_TrustManagementRouter.CallOpts)
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

// WALLETBEACON is a free data retrieval call binding the contract method 0x78efbba7.
//
// Solidity: function WALLET_BEACON() view returns(address)
func (_TrustManagementRouter *TrustManagementRouterCaller) WALLETBEACON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "WALLET_BEACON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WALLETBEACON is a free data retrieval call binding the contract method 0x78efbba7.
//
// Solidity: function WALLET_BEACON() view returns(address)
func (_TrustManagementRouter *TrustManagementRouterSession) WALLETBEACON() (common.Address, error) {
	return _TrustManagementRouter.Contract.WALLETBEACON(&_TrustManagementRouter.CallOpts)
}

// WALLETBEACON is a free data retrieval call binding the contract method 0x78efbba7.
//
// Solidity: function WALLET_BEACON() view returns(address)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) WALLETBEACON() (common.Address, error) {
	return _TrustManagementRouter.Contract.WALLETBEACON(&_TrustManagementRouter.CallOpts)
}

// AllowedToken is a free data retrieval call binding the contract method 0x756742f8.
//
// Solidity: function allowedToken(address tokenAddress) view returns(bool isAllowed)
func (_TrustManagementRouter *TrustManagementRouterCaller) AllowedToken(opts *bind.CallOpts, tokenAddress common.Address) (bool, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "allowedToken", tokenAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToken is a free data retrieval call binding the contract method 0x756742f8.
//
// Solidity: function allowedToken(address tokenAddress) view returns(bool isAllowed)
func (_TrustManagementRouter *TrustManagementRouterSession) AllowedToken(tokenAddress common.Address) (bool, error) {
	return _TrustManagementRouter.Contract.AllowedToken(&_TrustManagementRouter.CallOpts, tokenAddress)
}

// AllowedToken is a free data retrieval call binding the contract method 0x756742f8.
//
// Solidity: function allowedToken(address tokenAddress) view returns(bool isAllowed)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) AllowedToken(tokenAddress common.Address) (bool, error) {
	return _TrustManagementRouter.Contract.AllowedToken(&_TrustManagementRouter.CallOpts, tokenAddress)
}

// EcdsaOnKeccak256r1 is a free data retrieval call binding the contract method 0xce27ac17.
//
// Solidity: function ecdsaOnKeccak256r1(bytes message, bytes signature, bytes publicKey) pure returns(bool isValid)
func (_TrustManagementRouter *TrustManagementRouterCaller) EcdsaOnKeccak256r1(opts *bind.CallOpts, message []byte, signature []byte, publicKey []byte) (bool, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "ecdsaOnKeccak256r1", message, signature, publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EcdsaOnKeccak256r1 is a free data retrieval call binding the contract method 0xce27ac17.
//
// Solidity: function ecdsaOnKeccak256r1(bytes message, bytes signature, bytes publicKey) pure returns(bool isValid)
func (_TrustManagementRouter *TrustManagementRouterSession) EcdsaOnKeccak256r1(message []byte, signature []byte, publicKey []byte) (bool, error) {
	return _TrustManagementRouter.Contract.EcdsaOnKeccak256r1(&_TrustManagementRouter.CallOpts, message, signature, publicKey)
}

// EcdsaOnKeccak256r1 is a free data retrieval call binding the contract method 0xce27ac17.
//
// Solidity: function ecdsaOnKeccak256r1(bytes message, bytes signature, bytes publicKey) pure returns(bool isValid)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) EcdsaOnKeccak256r1(message []byte, signature []byte, publicKey []byte) (bool, error) {
	return _TrustManagementRouter.Contract.EcdsaOnKeccak256r1(&_TrustManagementRouter.CallOpts, message, signature, publicKey)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address feeCollectorAddress)
func (_TrustManagementRouter *TrustManagementRouterCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address feeCollectorAddress)
func (_TrustManagementRouter *TrustManagementRouterSession) FeeCollector() (common.Address, error) {
	return _TrustManagementRouter.Contract.FeeCollector(&_TrustManagementRouter.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address feeCollectorAddress)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) FeeCollector() (common.Address, error) {
	return _TrustManagementRouter.Contract.FeeCollector(&_TrustManagementRouter.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(address user, address[] tokens) view returns(uint256[] balances)
func (_TrustManagementRouter *TrustManagementRouterCaller) GetBalances(opts *bind.CallOpts, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "getBalances", user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(address user, address[] tokens) view returns(uint256[] balances)
func (_TrustManagementRouter *TrustManagementRouterSession) GetBalances(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _TrustManagementRouter.Contract.GetBalances(&_TrustManagementRouter.CallOpts, user, tokens)
}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(address user, address[] tokens) view returns(uint256[] balances)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) GetBalances(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _TrustManagementRouter.Contract.GetBalances(&_TrustManagementRouter.CallOpts, user, tokens)
}

// GetDeposits is a free data retrieval call binding the contract method 0x018fa360.
//
// Solidity: function getDeposits(address user, address token) view returns((uint256,uint256)[] deposits)
func (_TrustManagementRouter *TrustManagementRouterCaller) GetDeposits(opts *bind.CallOpts, user common.Address, token common.Address) ([]Deposit, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "getDeposits", user, token)

	if err != nil {
		return *new([]Deposit), err
	}

	out0 := *abi.ConvertType(out[0], new([]Deposit)).(*[]Deposit)

	return out0, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x018fa360.
//
// Solidity: function getDeposits(address user, address token) view returns((uint256,uint256)[] deposits)
func (_TrustManagementRouter *TrustManagementRouterSession) GetDeposits(user common.Address, token common.Address) ([]Deposit, error) {
	return _TrustManagementRouter.Contract.GetDeposits(&_TrustManagementRouter.CallOpts, user, token)
}

// GetDeposits is a free data retrieval call binding the contract method 0x018fa360.
//
// Solidity: function getDeposits(address user, address token) view returns((uint256,uint256)[] deposits)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) GetDeposits(user common.Address, token common.Address) ([]Deposit, error) {
	return _TrustManagementRouter.Contract.GetDeposits(&_TrustManagementRouter.CallOpts, user, token)
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

// GetWalletAddress is a free data retrieval call binding the contract method 0x1fd9afa5.
//
// Solidity: function getWalletAddress(address user) view returns(address walletAddress, bool isDeployed)
func (_TrustManagementRouter *TrustManagementRouterCaller) GetWalletAddress(opts *bind.CallOpts, user common.Address) (struct {
	WalletAddress common.Address
	IsDeployed    bool
}, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "getWalletAddress", user)

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
func (_TrustManagementRouter *TrustManagementRouterSession) GetWalletAddress(user common.Address) (struct {
	WalletAddress common.Address
	IsDeployed    bool
}, error) {
	return _TrustManagementRouter.Contract.GetWalletAddress(&_TrustManagementRouter.CallOpts, user)
}

// GetWalletAddress is a free data retrieval call binding the contract method 0x1fd9afa5.
//
// Solidity: function getWalletAddress(address user) view returns(address walletAddress, bool isDeployed)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) GetWalletAddress(user common.Address) (struct {
	WalletAddress common.Address
	IsDeployed    bool
}, error) {
	return _TrustManagementRouter.Contract.GetWalletAddress(&_TrustManagementRouter.CallOpts, user)
}

// GetWithdrawableAmount is a free data retrieval call binding the contract method 0x46c58227.
//
// Solidity: function getWithdrawableAmount(address user, address token) view returns(uint256 withdrawableAmount)
func (_TrustManagementRouter *TrustManagementRouterCaller) GetWithdrawableAmount(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "getWithdrawableAmount", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableAmount is a free data retrieval call binding the contract method 0x46c58227.
//
// Solidity: function getWithdrawableAmount(address user, address token) view returns(uint256 withdrawableAmount)
func (_TrustManagementRouter *TrustManagementRouterSession) GetWithdrawableAmount(user common.Address, token common.Address) (*big.Int, error) {
	return _TrustManagementRouter.Contract.GetWithdrawableAmount(&_TrustManagementRouter.CallOpts, user, token)
}

// GetWithdrawableAmount is a free data retrieval call binding the contract method 0x46c58227.
//
// Solidity: function getWithdrawableAmount(address user, address token) view returns(uint256 withdrawableAmount)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) GetWithdrawableAmount(user common.Address, token common.Address) (*big.Int, error) {
	return _TrustManagementRouter.Contract.GetWithdrawableAmount(&_TrustManagementRouter.CallOpts, user, token)
}

// GetWithdrawableDeposits is a free data retrieval call binding the contract method 0x7274b0c3.
//
// Solidity: function getWithdrawableDeposits(address user, address token) view returns(uint256[] withdrawableDepositsIds)
func (_TrustManagementRouter *TrustManagementRouterCaller) GetWithdrawableDeposits(opts *bind.CallOpts, user common.Address, token common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "getWithdrawableDeposits", user, token)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetWithdrawableDeposits is a free data retrieval call binding the contract method 0x7274b0c3.
//
// Solidity: function getWithdrawableDeposits(address user, address token) view returns(uint256[] withdrawableDepositsIds)
func (_TrustManagementRouter *TrustManagementRouterSession) GetWithdrawableDeposits(user common.Address, token common.Address) ([]*big.Int, error) {
	return _TrustManagementRouter.Contract.GetWithdrawableDeposits(&_TrustManagementRouter.CallOpts, user, token)
}

// GetWithdrawableDeposits is a free data retrieval call binding the contract method 0x7274b0c3.
//
// Solidity: function getWithdrawableDeposits(address user, address token) view returns(uint256[] withdrawableDepositsIds)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) GetWithdrawableDeposits(user common.Address, token common.Address) ([]*big.Int, error) {
	return _TrustManagementRouter.Contract.GetWithdrawableDeposits(&_TrustManagementRouter.CallOpts, user, token)
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

// MrEnclaveAuthorized is a free data retrieval call binding the contract method 0x9a16def0.
//
// Solidity: function mrEnclaveAuthorized(bytes32 mrEnclave) view returns(bool isAuthorized)
func (_TrustManagementRouter *TrustManagementRouterCaller) MrEnclaveAuthorized(opts *bind.CallOpts, mrEnclave [32]byte) (bool, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "mrEnclaveAuthorized", mrEnclave)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MrEnclaveAuthorized is a free data retrieval call binding the contract method 0x9a16def0.
//
// Solidity: function mrEnclaveAuthorized(bytes32 mrEnclave) view returns(bool isAuthorized)
func (_TrustManagementRouter *TrustManagementRouterSession) MrEnclaveAuthorized(mrEnclave [32]byte) (bool, error) {
	return _TrustManagementRouter.Contract.MrEnclaveAuthorized(&_TrustManagementRouter.CallOpts, mrEnclave)
}

// MrEnclaveAuthorized is a free data retrieval call binding the contract method 0x9a16def0.
//
// Solidity: function mrEnclaveAuthorized(bytes32 mrEnclave) view returns(bool isAuthorized)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) MrEnclaveAuthorized(mrEnclave [32]byte) (bool, error) {
	return _TrustManagementRouter.Contract.MrEnclaveAuthorized(&_TrustManagementRouter.CallOpts, mrEnclave)
}

// PerfomanceFeeRate is a free data retrieval call binding the contract method 0x2571eb15.
//
// Solidity: function perfomanceFeeRate() view returns(uint256 performanceFeeRateValue)
func (_TrustManagementRouter *TrustManagementRouterCaller) PerfomanceFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "perfomanceFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerfomanceFeeRate is a free data retrieval call binding the contract method 0x2571eb15.
//
// Solidity: function perfomanceFeeRate() view returns(uint256 performanceFeeRateValue)
func (_TrustManagementRouter *TrustManagementRouterSession) PerfomanceFeeRate() (*big.Int, error) {
	return _TrustManagementRouter.Contract.PerfomanceFeeRate(&_TrustManagementRouter.CallOpts)
}

// PerfomanceFeeRate is a free data retrieval call binding the contract method 0x2571eb15.
//
// Solidity: function perfomanceFeeRate() view returns(uint256 performanceFeeRateValue)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) PerfomanceFeeRate() (*big.Int, error) {
	return _TrustManagementRouter.Contract.PerfomanceFeeRate(&_TrustManagementRouter.CallOpts)
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

// RootCert is a free data retrieval call binding the contract method 0x7221e6fc.
//
// Solidity: function rootCert() view returns((bytes,bytes,bytes,bytes) rootCertificate)
func (_TrustManagementRouter *TrustManagementRouterCaller) RootCert(opts *bind.CallOpts) (ChunkedX509Cert, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "rootCert")

	if err != nil {
		return *new(ChunkedX509Cert), err
	}

	out0 := *abi.ConvertType(out[0], new(ChunkedX509Cert)).(*ChunkedX509Cert)

	return out0, err

}

// RootCert is a free data retrieval call binding the contract method 0x7221e6fc.
//
// Solidity: function rootCert() view returns((bytes,bytes,bytes,bytes) rootCertificate)
func (_TrustManagementRouter *TrustManagementRouterSession) RootCert() (ChunkedX509Cert, error) {
	return _TrustManagementRouter.Contract.RootCert(&_TrustManagementRouter.CallOpts)
}

// RootCert is a free data retrieval call binding the contract method 0x7221e6fc.
//
// Solidity: function rootCert() view returns((bytes,bytes,bytes,bytes) rootCertificate)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) RootCert() (ChunkedX509Cert, error) {
	return _TrustManagementRouter.Contract.RootCert(&_TrustManagementRouter.CallOpts)
}

// Sessions is a free data retrieval call binding the contract method 0x431a1b97.
//
// Solidity: function sessions(address sessionKey) view returns(bool isActive)
func (_TrustManagementRouter *TrustManagementRouterCaller) Sessions(opts *bind.CallOpts, sessionKey common.Address) (bool, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "sessions", sessionKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Sessions is a free data retrieval call binding the contract method 0x431a1b97.
//
// Solidity: function sessions(address sessionKey) view returns(bool isActive)
func (_TrustManagementRouter *TrustManagementRouterSession) Sessions(sessionKey common.Address) (bool, error) {
	return _TrustManagementRouter.Contract.Sessions(&_TrustManagementRouter.CallOpts, sessionKey)
}

// Sessions is a free data retrieval call binding the contract method 0x431a1b97.
//
// Solidity: function sessions(address sessionKey) view returns(bool isActive)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) Sessions(sessionKey common.Address) (bool, error) {
	return _TrustManagementRouter.Contract.Sessions(&_TrustManagementRouter.CallOpts, sessionKey)
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

// TotalDeposited is a free data retrieval call binding the contract method 0x0cac4d76.
//
// Solidity: function totalDeposited(address user, address token) view returns(uint256 totalDepositedTokens)
func (_TrustManagementRouter *TrustManagementRouterCaller) TotalDeposited(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "totalDeposited", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeposited is a free data retrieval call binding the contract method 0x0cac4d76.
//
// Solidity: function totalDeposited(address user, address token) view returns(uint256 totalDepositedTokens)
func (_TrustManagementRouter *TrustManagementRouterSession) TotalDeposited(user common.Address, token common.Address) (*big.Int, error) {
	return _TrustManagementRouter.Contract.TotalDeposited(&_TrustManagementRouter.CallOpts, user, token)
}

// TotalDeposited is a free data retrieval call binding the contract method 0x0cac4d76.
//
// Solidity: function totalDeposited(address user, address token) view returns(uint256 totalDepositedTokens)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) TotalDeposited(user common.Address, token common.Address) (*big.Int, error) {
	return _TrustManagementRouter.Contract.TotalDeposited(&_TrustManagementRouter.CallOpts, user, token)
}

// UsedSignature is a free data retrieval call binding the contract method 0x0f945e9f.
//
// Solidity: function usedSignature(bytes32 signatureHash) view returns(bool isUsed)
func (_TrustManagementRouter *TrustManagementRouterCaller) UsedSignature(opts *bind.CallOpts, signatureHash [32]byte) (bool, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "usedSignature", signatureHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedSignature is a free data retrieval call binding the contract method 0x0f945e9f.
//
// Solidity: function usedSignature(bytes32 signatureHash) view returns(bool isUsed)
func (_TrustManagementRouter *TrustManagementRouterSession) UsedSignature(signatureHash [32]byte) (bool, error) {
	return _TrustManagementRouter.Contract.UsedSignature(&_TrustManagementRouter.CallOpts, signatureHash)
}

// UsedSignature is a free data retrieval call binding the contract method 0x0f945e9f.
//
// Solidity: function usedSignature(bytes32 signatureHash) view returns(bool isUsed)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) UsedSignature(signatureHash [32]byte) (bool, error) {
	return _TrustManagementRouter.Contract.UsedSignature(&_TrustManagementRouter.CallOpts, signatureHash)
}

// VerifyCert is a free data retrieval call binding the contract method 0xbee94f11.
//
// Solidity: function verifyCert((bytes,bytes,bytes,bytes) cert, bytes publicKey) pure returns(bool isValid)
func (_TrustManagementRouter *TrustManagementRouterCaller) VerifyCert(opts *bind.CallOpts, cert ChunkedX509Cert, publicKey []byte) (bool, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "verifyCert", cert, publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCert is a free data retrieval call binding the contract method 0xbee94f11.
//
// Solidity: function verifyCert((bytes,bytes,bytes,bytes) cert, bytes publicKey) pure returns(bool isValid)
func (_TrustManagementRouter *TrustManagementRouterSession) VerifyCert(cert ChunkedX509Cert, publicKey []byte) (bool, error) {
	return _TrustManagementRouter.Contract.VerifyCert(&_TrustManagementRouter.CallOpts, cert, publicKey)
}

// VerifyCert is a free data retrieval call binding the contract method 0xbee94f11.
//
// Solidity: function verifyCert((bytes,bytes,bytes,bytes) cert, bytes publicKey) pure returns(bool isValid)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) VerifyCert(cert ChunkedX509Cert, publicKey []byte) (bool, error) {
	return _TrustManagementRouter.Contract.VerifyCert(&_TrustManagementRouter.CallOpts, cert, publicKey)
}

// XyberContractName is a free data retrieval call binding the contract method 0xe60e9654.
//
// Solidity: function xyberContractName() pure returns(string contractName)
func (_TrustManagementRouter *TrustManagementRouterCaller) XyberContractName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TrustManagementRouter.contract.Call(opts, &out, "xyberContractName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// XyberContractName is a free data retrieval call binding the contract method 0xe60e9654.
//
// Solidity: function xyberContractName() pure returns(string contractName)
func (_TrustManagementRouter *TrustManagementRouterSession) XyberContractName() (string, error) {
	return _TrustManagementRouter.Contract.XyberContractName(&_TrustManagementRouter.CallOpts)
}

// XyberContractName is a free data retrieval call binding the contract method 0xe60e9654.
//
// Solidity: function xyberContractName() pure returns(string contractName)
func (_TrustManagementRouter *TrustManagementRouterCallerSession) XyberContractName() (string, error) {
	return _TrustManagementRouter.Contract.XyberContractName(&_TrustManagementRouter.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0xfff6f90b.
//
// Solidity: function claim(address user, address receiver, address[] tokens, uint256[] amounts) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) Claim(opts *bind.TransactOpts, user common.Address, receiver common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "claim", user, receiver, tokens, amounts)
}

// Claim is a paid mutator transaction binding the contract method 0xfff6f90b.
//
// Solidity: function claim(address user, address receiver, address[] tokens, uint256[] amounts) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) Claim(user common.Address, receiver common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Claim(&_TrustManagementRouter.TransactOpts, user, receiver, tokens, amounts)
}

// Claim is a paid mutator transaction binding the contract method 0xfff6f90b.
//
// Solidity: function claim(address user, address receiver, address[] tokens, uint256[] amounts) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Claim(user common.Address, receiver common.Address, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Claim(&_TrustManagementRouter.TransactOpts, user, receiver, tokens, amounts)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xb054a9e8.
//
// Solidity: function createWallet(address user) returns(address walletAddress)
func (_TrustManagementRouter *TrustManagementRouterTransactor) CreateWallet(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "createWallet", user)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xb054a9e8.
//
// Solidity: function createWallet(address user) returns(address walletAddress)
func (_TrustManagementRouter *TrustManagementRouterSession) CreateWallet(user common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.CreateWallet(&_TrustManagementRouter.TransactOpts, user)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xb054a9e8.
//
// Solidity: function createWallet(address user) returns(address walletAddress)
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) CreateWallet(user common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.CreateWallet(&_TrustManagementRouter.TransactOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x20e8c565.
//
// Solidity: function deposit(address receiver, address token, uint256 amount, uint256 lockDuration) payable returns(address walletAddress)
func (_TrustManagementRouter *TrustManagementRouterTransactor) Deposit(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, lockDuration *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "deposit", receiver, token, amount, lockDuration)
}

// Deposit is a paid mutator transaction binding the contract method 0x20e8c565.
//
// Solidity: function deposit(address receiver, address token, uint256 amount, uint256 lockDuration) payable returns(address walletAddress)
func (_TrustManagementRouter *TrustManagementRouterSession) Deposit(receiver common.Address, token common.Address, amount *big.Int, lockDuration *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Deposit(&_TrustManagementRouter.TransactOpts, receiver, token, amount, lockDuration)
}

// Deposit is a paid mutator transaction binding the contract method 0x20e8c565.
//
// Solidity: function deposit(address receiver, address token, uint256 amount, uint256 lockDuration) payable returns(address walletAddress)
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Deposit(receiver common.Address, token common.Address, amount *big.Int, lockDuration *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Deposit(&_TrustManagementRouter.TransactOpts, receiver, token, amount, lockDuration)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x96931cf5.
//
// Solidity: function depositWithPermit(address receiver, address token, uint256 lockDuration, (address,uint256,uint256,uint8,bytes32,bytes32) permit) returns(address walletAddress)
func (_TrustManagementRouter *TrustManagementRouterTransactor) DepositWithPermit(opts *bind.TransactOpts, receiver common.Address, token common.Address, lockDuration *big.Int, permit Permit) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "depositWithPermit", receiver, token, lockDuration, permit)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x96931cf5.
//
// Solidity: function depositWithPermit(address receiver, address token, uint256 lockDuration, (address,uint256,uint256,uint8,bytes32,bytes32) permit) returns(address walletAddress)
func (_TrustManagementRouter *TrustManagementRouterSession) DepositWithPermit(receiver common.Address, token common.Address, lockDuration *big.Int, permit Permit) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.DepositWithPermit(&_TrustManagementRouter.TransactOpts, receiver, token, lockDuration, permit)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x96931cf5.
//
// Solidity: function depositWithPermit(address receiver, address token, uint256 lockDuration, (address,uint256,uint256,uint8,bytes32,bytes32) permit) returns(address walletAddress)
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) DepositWithPermit(receiver common.Address, token common.Address, lockDuration *big.Int, permit Permit) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.DepositWithPermit(&_TrustManagementRouter.TransactOpts, receiver, token, lockDuration, permit)
}

// Execute is a paid mutator transaction binding the contract method 0x2f37da74.
//
// Solidity: function execute((address,uint256,bytes)[] txs, bytes signature, uint256 deadline) payable returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) Execute(opts *bind.TransactOpts, txs []Transaction, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "execute", txs, signature, deadline)
}

// Execute is a paid mutator transaction binding the contract method 0x2f37da74.
//
// Solidity: function execute((address,uint256,bytes)[] txs, bytes signature, uint256 deadline) payable returns()
func (_TrustManagementRouter *TrustManagementRouterSession) Execute(txs []Transaction, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Execute(&_TrustManagementRouter.TransactOpts, txs, signature, deadline)
}

// Execute is a paid mutator transaction binding the contract method 0x2f37da74.
//
// Solidity: function execute((address,uint256,bytes)[] txs, bytes signature, uint256 deadline) payable returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Execute(txs []Transaction, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Execute(&_TrustManagementRouter.TransactOpts, txs, signature, deadline)
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

// InitSessionKey is a paid mutator transaction binding the contract method 0x4c7c04c4.
//
// Solidity: function initSessionKey((bytes,bytes,bytes,bytes) leaf, (bytes,bytes,bytes,bytes) intermediate, (bytes,bytes,bytes,bytes,bytes,bytes,bytes) quote, address sessionKey) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) InitSessionKey(opts *bind.TransactOpts, leaf ChunkedX509Cert, intermediate ChunkedX509Cert, quote ChunkedSGXQuote, sessionKey common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "initSessionKey", leaf, intermediate, quote, sessionKey)
}

// InitSessionKey is a paid mutator transaction binding the contract method 0x4c7c04c4.
//
// Solidity: function initSessionKey((bytes,bytes,bytes,bytes) leaf, (bytes,bytes,bytes,bytes) intermediate, (bytes,bytes,bytes,bytes,bytes,bytes,bytes) quote, address sessionKey) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) InitSessionKey(leaf ChunkedX509Cert, intermediate ChunkedX509Cert, quote ChunkedSGXQuote, sessionKey common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.InitSessionKey(&_TrustManagementRouter.TransactOpts, leaf, intermediate, quote, sessionKey)
}

// InitSessionKey is a paid mutator transaction binding the contract method 0x4c7c04c4.
//
// Solidity: function initSessionKey((bytes,bytes,bytes,bytes) leaf, (bytes,bytes,bytes,bytes) intermediate, (bytes,bytes,bytes,bytes,bytes,bytes,bytes) quote, address sessionKey) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) InitSessionKey(leaf ChunkedX509Cert, intermediate ChunkedX509Cert, quote ChunkedSGXQuote, sessionKey common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.InitSessionKey(&_TrustManagementRouter.TransactOpts, leaf, intermediate, quote, sessionKey)
}

// Initialize is a paid mutator transaction binding the contract method 0x776ad9dc.
//
// Solidity: function initialize(address defaultAdmin, (bytes,bytes,bytes,bytes) rootCert) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) Initialize(opts *bind.TransactOpts, defaultAdmin common.Address, rootCert ChunkedX509Cert) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "initialize", defaultAdmin, rootCert)
}

// Initialize is a paid mutator transaction binding the contract method 0x776ad9dc.
//
// Solidity: function initialize(address defaultAdmin, (bytes,bytes,bytes,bytes) rootCert) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) Initialize(defaultAdmin common.Address, rootCert ChunkedX509Cert) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Initialize(&_TrustManagementRouter.TransactOpts, defaultAdmin, rootCert)
}

// Initialize is a paid mutator transaction binding the contract method 0x776ad9dc.
//
// Solidity: function initialize(address defaultAdmin, (bytes,bytes,bytes,bytes) rootCert) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Initialize(defaultAdmin common.Address, rootCert ChunkedX509Cert) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Initialize(&_TrustManagementRouter.TransactOpts, defaultAdmin, rootCert)
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

// SetAllowedToken is a paid mutator transaction binding the contract method 0x8aaa2284.
//
// Solidity: function setAllowedToken(address token, bool isAllowed) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) SetAllowedToken(opts *bind.TransactOpts, token common.Address, isAllowed bool) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "setAllowedToken", token, isAllowed)
}

// SetAllowedToken is a paid mutator transaction binding the contract method 0x8aaa2284.
//
// Solidity: function setAllowedToken(address token, bool isAllowed) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) SetAllowedToken(token common.Address, isAllowed bool) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.SetAllowedToken(&_TrustManagementRouter.TransactOpts, token, isAllowed)
}

// SetAllowedToken is a paid mutator transaction binding the contract method 0x8aaa2284.
//
// Solidity: function setAllowedToken(address token, bool isAllowed) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) SetAllowedToken(token common.Address, isAllowed bool) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.SetAllowedToken(&_TrustManagementRouter.TransactOpts, token, isAllowed)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address newFeeCollector) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) SetFeeCollector(opts *bind.TransactOpts, newFeeCollector common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "setFeeCollector", newFeeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address newFeeCollector) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) SetFeeCollector(newFeeCollector common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.SetFeeCollector(&_TrustManagementRouter.TransactOpts, newFeeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address newFeeCollector) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) SetFeeCollector(newFeeCollector common.Address) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.SetFeeCollector(&_TrustManagementRouter.TransactOpts, newFeeCollector)
}

// SetMrEnclave is a paid mutator transaction binding the contract method 0x3a343014.
//
// Solidity: function setMrEnclave(bytes32 mrEnclave, bool isAuthorized) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) SetMrEnclave(opts *bind.TransactOpts, mrEnclave [32]byte, isAuthorized bool) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "setMrEnclave", mrEnclave, isAuthorized)
}

// SetMrEnclave is a paid mutator transaction binding the contract method 0x3a343014.
//
// Solidity: function setMrEnclave(bytes32 mrEnclave, bool isAuthorized) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) SetMrEnclave(mrEnclave [32]byte, isAuthorized bool) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.SetMrEnclave(&_TrustManagementRouter.TransactOpts, mrEnclave, isAuthorized)
}

// SetMrEnclave is a paid mutator transaction binding the contract method 0x3a343014.
//
// Solidity: function setMrEnclave(bytes32 mrEnclave, bool isAuthorized) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) SetMrEnclave(mrEnclave [32]byte, isAuthorized bool) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.SetMrEnclave(&_TrustManagementRouter.TransactOpts, mrEnclave, isAuthorized)
}

// SetPerfomanceFeeRate is a paid mutator transaction binding the contract method 0x24402bba.
//
// Solidity: function setPerfomanceFeeRate(uint256 newPerfomanceFeeRate) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) SetPerfomanceFeeRate(opts *bind.TransactOpts, newPerfomanceFeeRate *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "setPerfomanceFeeRate", newPerfomanceFeeRate)
}

// SetPerfomanceFeeRate is a paid mutator transaction binding the contract method 0x24402bba.
//
// Solidity: function setPerfomanceFeeRate(uint256 newPerfomanceFeeRate) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) SetPerfomanceFeeRate(newPerfomanceFeeRate *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.SetPerfomanceFeeRate(&_TrustManagementRouter.TransactOpts, newPerfomanceFeeRate)
}

// SetPerfomanceFeeRate is a paid mutator transaction binding the contract method 0x24402bba.
//
// Solidity: function setPerfomanceFeeRate(uint256 newPerfomanceFeeRate) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) SetPerfomanceFeeRate(newPerfomanceFeeRate *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.SetPerfomanceFeeRate(&_TrustManagementRouter.TransactOpts, newPerfomanceFeeRate)
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

// Withdraw is a paid mutator transaction binding the contract method 0x9b40e41e.
//
// Solidity: function withdraw(address user, address token, address receiver, uint256[] depositIds, uint256 baseAmount, uint256 amountWithYield) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactor) Withdraw(opts *bind.TransactOpts, user common.Address, token common.Address, receiver common.Address, depositIds []*big.Int, baseAmount *big.Int, amountWithYield *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.contract.Transact(opts, "withdraw", user, token, receiver, depositIds, baseAmount, amountWithYield)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b40e41e.
//
// Solidity: function withdraw(address user, address token, address receiver, uint256[] depositIds, uint256 baseAmount, uint256 amountWithYield) returns()
func (_TrustManagementRouter *TrustManagementRouterSession) Withdraw(user common.Address, token common.Address, receiver common.Address, depositIds []*big.Int, baseAmount *big.Int, amountWithYield *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Withdraw(&_TrustManagementRouter.TransactOpts, user, token, receiver, depositIds, baseAmount, amountWithYield)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b40e41e.
//
// Solidity: function withdraw(address user, address token, address receiver, uint256[] depositIds, uint256 baseAmount, uint256 amountWithYield) returns()
func (_TrustManagementRouter *TrustManagementRouterTransactorSession) Withdraw(user common.Address, token common.Address, receiver common.Address, depositIds []*big.Int, baseAmount *big.Int, amountWithYield *big.Int) (*types.Transaction, error) {
	return _TrustManagementRouter.Contract.Withdraw(&_TrustManagementRouter.TransactOpts, user, token, receiver, depositIds, baseAmount, amountWithYield)
}

// TrustManagementRouterAllowedTokenSetIterator is returned from FilterAllowedTokenSet and is used to iterate over the raw logs and unpacked data for AllowedTokenSet events raised by the TrustManagementRouter contract.
type TrustManagementRouterAllowedTokenSetIterator struct {
	Event *TrustManagementRouterAllowedTokenSet // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterAllowedTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterAllowedTokenSet)
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
		it.Event = new(TrustManagementRouterAllowedTokenSet)
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
func (it *TrustManagementRouterAllowedTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterAllowedTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterAllowedTokenSet represents a AllowedTokenSet event raised by the TrustManagementRouter contract.
type TrustManagementRouterAllowedTokenSet struct {
	Token     common.Address
	IsAllowed bool
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAllowedTokenSet is a free log retrieval operation binding the contract event 0x08e00fb187b2ed4d0d966e3e6885897b97cd8695c9731410b7d2a189e474c95e.
//
// Solidity: event AllowedTokenSet(address indexed token, bool indexed isAllowed, address caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterAllowedTokenSet(opts *bind.FilterOpts, token []common.Address, isAllowed []bool) (*TrustManagementRouterAllowedTokenSetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var isAllowedRule []interface{}
	for _, isAllowedItem := range isAllowed {
		isAllowedRule = append(isAllowedRule, isAllowedItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "AllowedTokenSet", tokenRule, isAllowedRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterAllowedTokenSetIterator{contract: _TrustManagementRouter.contract, event: "AllowedTokenSet", logs: logs, sub: sub}, nil
}

// WatchAllowedTokenSet is a free log subscription operation binding the contract event 0x08e00fb187b2ed4d0d966e3e6885897b97cd8695c9731410b7d2a189e474c95e.
//
// Solidity: event AllowedTokenSet(address indexed token, bool indexed isAllowed, address caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchAllowedTokenSet(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterAllowedTokenSet, token []common.Address, isAllowed []bool) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var isAllowedRule []interface{}
	for _, isAllowedItem := range isAllowed {
		isAllowedRule = append(isAllowedRule, isAllowedItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "AllowedTokenSet", tokenRule, isAllowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterAllowedTokenSet)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "AllowedTokenSet", log); err != nil {
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

// ParseAllowedTokenSet is a log parse operation binding the contract event 0x08e00fb187b2ed4d0d966e3e6885897b97cd8695c9731410b7d2a189e474c95e.
//
// Solidity: event AllowedTokenSet(address indexed token, bool indexed isAllowed, address caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseAllowedTokenSet(log types.Log) (*TrustManagementRouterAllowedTokenSet, error) {
	event := new(TrustManagementRouterAllowedTokenSet)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "AllowedTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementRouterClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the TrustManagementRouter contract.
type TrustManagementRouterClaimedIterator struct {
	Event *TrustManagementRouterClaimed // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterClaimed)
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
		it.Event = new(TrustManagementRouterClaimed)
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
func (it *TrustManagementRouterClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterClaimed represents a Claimed event raised by the TrustManagementRouter contract.
type TrustManagementRouterClaimed struct {
	Wallet   common.Address
	User     common.Address
	Tokens   []common.Address
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x6341e54c99570fa6967fa5800dd288c6b0c405bacae23a5ad6a94e0f970318f8.
//
// Solidity: event Claimed(address wallet, address indexed user, address[] tokens, address receiver)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterClaimed(opts *bind.FilterOpts, user []common.Address) (*TrustManagementRouterClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "Claimed", userRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterClaimedIterator{contract: _TrustManagementRouter.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x6341e54c99570fa6967fa5800dd288c6b0c405bacae23a5ad6a94e0f970318f8.
//
// Solidity: event Claimed(address wallet, address indexed user, address[] tokens, address receiver)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "Claimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterClaimed)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x6341e54c99570fa6967fa5800dd288c6b0c405bacae23a5ad6a94e0f970318f8.
//
// Solidity: event Claimed(address wallet, address indexed user, address[] tokens, address receiver)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseClaimed(log types.Log) (*TrustManagementRouterClaimed, error) {
	event := new(TrustManagementRouterClaimed)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementRouterDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the TrustManagementRouter contract.
type TrustManagementRouterDepositedIterator struct {
	Event *TrustManagementRouterDeposited // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterDeposited)
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
		it.Event = new(TrustManagementRouterDeposited)
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
func (it *TrustManagementRouterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterDeposited represents a Deposited event raised by the TrustManagementRouter contract.
type TrustManagementRouterDeposited struct {
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
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*TrustManagementRouterDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "Deposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterDepositedIterator{contract: _TrustManagementRouter.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xb045190548dadae679cfe9e337437613ca6dd73efdf984f75e56f152ccee22f0.
//
// Solidity: event Deposited(address wallet, address indexed user, address indexed token, uint256 amount, uint256 lockDuration)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterDeposited, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "Deposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterDeposited)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseDeposited(log types.Log) (*TrustManagementRouterDeposited, error) {
	event := new(TrustManagementRouterDeposited)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementRouterExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the TrustManagementRouter contract.
type TrustManagementRouterExecutedIterator struct {
	Event *TrustManagementRouterExecuted // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterExecuted)
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
		it.Event = new(TrustManagementRouterExecuted)
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
func (it *TrustManagementRouterExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterExecuted represents a Executed event raised by the TrustManagementRouter contract.
type TrustManagementRouterExecuted struct {
	Target common.Address
	Data   []byte
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x40e4545157549759798da7c1865562e1e6473f16e61954ae287208f49962927f.
//
// Solidity: event Executed(address indexed target, bytes data, address caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterExecuted(opts *bind.FilterOpts, target []common.Address) (*TrustManagementRouterExecutedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "Executed", targetRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterExecutedIterator{contract: _TrustManagementRouter.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x40e4545157549759798da7c1865562e1e6473f16e61954ae287208f49962927f.
//
// Solidity: event Executed(address indexed target, bytes data, address caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterExecuted, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "Executed", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterExecuted)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x40e4545157549759798da7c1865562e1e6473f16e61954ae287208f49962927f.
//
// Solidity: event Executed(address indexed target, bytes data, address caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseExecuted(log types.Log) (*TrustManagementRouterExecuted, error) {
	event := new(TrustManagementRouterExecuted)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementRouterFeeCollectorSetIterator is returned from FilterFeeCollectorSet and is used to iterate over the raw logs and unpacked data for FeeCollectorSet events raised by the TrustManagementRouter contract.
type TrustManagementRouterFeeCollectorSetIterator struct {
	Event *TrustManagementRouterFeeCollectorSet // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterFeeCollectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterFeeCollectorSet)
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
		it.Event = new(TrustManagementRouterFeeCollectorSet)
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
func (it *TrustManagementRouterFeeCollectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterFeeCollectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterFeeCollectorSet represents a FeeCollectorSet event raised by the TrustManagementRouter contract.
type TrustManagementRouterFeeCollectorSet struct {
	NewFeeCollector common.Address
	Caller          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorSet is a free log retrieval operation binding the contract event 0x089588e3f10370c99a6f74177eacb5361ba90e1b70a123bfeccb6619c21cd721.
//
// Solidity: event FeeCollectorSet(address newFeeCollector, address caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterFeeCollectorSet(opts *bind.FilterOpts) (*TrustManagementRouterFeeCollectorSetIterator, error) {

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "FeeCollectorSet")
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterFeeCollectorSetIterator{contract: _TrustManagementRouter.contract, event: "FeeCollectorSet", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorSet is a free log subscription operation binding the contract event 0x089588e3f10370c99a6f74177eacb5361ba90e1b70a123bfeccb6619c21cd721.
//
// Solidity: event FeeCollectorSet(address newFeeCollector, address caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchFeeCollectorSet(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterFeeCollectorSet) (event.Subscription, error) {

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "FeeCollectorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterFeeCollectorSet)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
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

// ParseFeeCollectorSet is a log parse operation binding the contract event 0x089588e3f10370c99a6f74177eacb5361ba90e1b70a123bfeccb6619c21cd721.
//
// Solidity: event FeeCollectorSet(address newFeeCollector, address caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseFeeCollectorSet(log types.Log) (*TrustManagementRouterFeeCollectorSet, error) {
	event := new(TrustManagementRouterFeeCollectorSet)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "FeeCollectorSet", log); err != nil {
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

// TrustManagementRouterMrEnclaveSetIterator is returned from FilterMrEnclaveSet and is used to iterate over the raw logs and unpacked data for MrEnclaveSet events raised by the TrustManagementRouter contract.
type TrustManagementRouterMrEnclaveSetIterator struct {
	Event *TrustManagementRouterMrEnclaveSet // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterMrEnclaveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterMrEnclaveSet)
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
		it.Event = new(TrustManagementRouterMrEnclaveSet)
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
func (it *TrustManagementRouterMrEnclaveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterMrEnclaveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterMrEnclaveSet represents a MrEnclaveSet event raised by the TrustManagementRouter contract.
type TrustManagementRouterMrEnclaveSet struct {
	MrEnclave    [32]byte
	IsAuthorized bool
	Caller       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMrEnclaveSet is a free log retrieval operation binding the contract event 0xf2ec1200fd7533c710e00ced1a72001aba7b6d3eec4a7c6e281de2ad16910598.
//
// Solidity: event MrEnclaveSet(bytes32 mrEnclave, bool isAuthorized, address indexed caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterMrEnclaveSet(opts *bind.FilterOpts, caller []common.Address) (*TrustManagementRouterMrEnclaveSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "MrEnclaveSet", callerRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterMrEnclaveSetIterator{contract: _TrustManagementRouter.contract, event: "MrEnclaveSet", logs: logs, sub: sub}, nil
}

// WatchMrEnclaveSet is a free log subscription operation binding the contract event 0xf2ec1200fd7533c710e00ced1a72001aba7b6d3eec4a7c6e281de2ad16910598.
//
// Solidity: event MrEnclaveSet(bytes32 mrEnclave, bool isAuthorized, address indexed caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchMrEnclaveSet(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterMrEnclaveSet, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "MrEnclaveSet", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterMrEnclaveSet)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "MrEnclaveSet", log); err != nil {
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

// ParseMrEnclaveSet is a log parse operation binding the contract event 0xf2ec1200fd7533c710e00ced1a72001aba7b6d3eec4a7c6e281de2ad16910598.
//
// Solidity: event MrEnclaveSet(bytes32 mrEnclave, bool isAuthorized, address indexed caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseMrEnclaveSet(log types.Log) (*TrustManagementRouterMrEnclaveSet, error) {
	event := new(TrustManagementRouterMrEnclaveSet)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "MrEnclaveSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementRouterPerformanceFeeRateSetIterator is returned from FilterPerformanceFeeRateSet and is used to iterate over the raw logs and unpacked data for PerformanceFeeRateSet events raised by the TrustManagementRouter contract.
type TrustManagementRouterPerformanceFeeRateSetIterator struct {
	Event *TrustManagementRouterPerformanceFeeRateSet // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterPerformanceFeeRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterPerformanceFeeRateSet)
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
		it.Event = new(TrustManagementRouterPerformanceFeeRateSet)
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
func (it *TrustManagementRouterPerformanceFeeRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterPerformanceFeeRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterPerformanceFeeRateSet represents a PerformanceFeeRateSet event raised by the TrustManagementRouter contract.
type TrustManagementRouterPerformanceFeeRateSet struct {
	NewPerfomanceFeeRate *big.Int
	Caller               common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPerformanceFeeRateSet is a free log retrieval operation binding the contract event 0xaa7531042271e1c8eee267fa80dead7e054f25a9dc0ea5709de69a1ca378c5ef.
//
// Solidity: event PerformanceFeeRateSet(uint256 newPerfomanceFeeRate, address caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterPerformanceFeeRateSet(opts *bind.FilterOpts) (*TrustManagementRouterPerformanceFeeRateSetIterator, error) {

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "PerformanceFeeRateSet")
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterPerformanceFeeRateSetIterator{contract: _TrustManagementRouter.contract, event: "PerformanceFeeRateSet", logs: logs, sub: sub}, nil
}

// WatchPerformanceFeeRateSet is a free log subscription operation binding the contract event 0xaa7531042271e1c8eee267fa80dead7e054f25a9dc0ea5709de69a1ca378c5ef.
//
// Solidity: event PerformanceFeeRateSet(uint256 newPerfomanceFeeRate, address caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchPerformanceFeeRateSet(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterPerformanceFeeRateSet) (event.Subscription, error) {

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "PerformanceFeeRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterPerformanceFeeRateSet)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "PerformanceFeeRateSet", log); err != nil {
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

// ParsePerformanceFeeRateSet is a log parse operation binding the contract event 0xaa7531042271e1c8eee267fa80dead7e054f25a9dc0ea5709de69a1ca378c5ef.
//
// Solidity: event PerformanceFeeRateSet(uint256 newPerfomanceFeeRate, address caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParsePerformanceFeeRateSet(log types.Log) (*TrustManagementRouterPerformanceFeeRateSet, error) {
	event := new(TrustManagementRouterPerformanceFeeRateSet)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "PerformanceFeeRateSet", log); err != nil {
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

// TrustManagementRouterRootCertSetIterator is returned from FilterRootCertSet and is used to iterate over the raw logs and unpacked data for RootCertSet events raised by the TrustManagementRouter contract.
type TrustManagementRouterRootCertSetIterator struct {
	Event *TrustManagementRouterRootCertSet // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterRootCertSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterRootCertSet)
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
		it.Event = new(TrustManagementRouterRootCertSet)
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
func (it *TrustManagementRouterRootCertSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterRootCertSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterRootCertSet represents a RootCertSet event raised by the TrustManagementRouter contract.
type TrustManagementRouterRootCertSet struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRootCertSet is a free log retrieval operation binding the contract event 0xd941fc17d44329a3280c0dfd2da1fc4a0e7fb2337b63abff947578411d678e9c.
//
// Solidity: event RootCertSet(address indexed caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterRootCertSet(opts *bind.FilterOpts, caller []common.Address) (*TrustManagementRouterRootCertSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "RootCertSet", callerRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterRootCertSetIterator{contract: _TrustManagementRouter.contract, event: "RootCertSet", logs: logs, sub: sub}, nil
}

// WatchRootCertSet is a free log subscription operation binding the contract event 0xd941fc17d44329a3280c0dfd2da1fc4a0e7fb2337b63abff947578411d678e9c.
//
// Solidity: event RootCertSet(address indexed caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchRootCertSet(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterRootCertSet, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "RootCertSet", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterRootCertSet)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "RootCertSet", log); err != nil {
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

// ParseRootCertSet is a log parse operation binding the contract event 0xd941fc17d44329a3280c0dfd2da1fc4a0e7fb2337b63abff947578411d678e9c.
//
// Solidity: event RootCertSet(address indexed caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseRootCertSet(log types.Log) (*TrustManagementRouterRootCertSet, error) {
	event := new(TrustManagementRouterRootCertSet)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "RootCertSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementRouterSessionInitializedIterator is returned from FilterSessionInitialized and is used to iterate over the raw logs and unpacked data for SessionInitialized events raised by the TrustManagementRouter contract.
type TrustManagementRouterSessionInitializedIterator struct {
	Event *TrustManagementRouterSessionInitialized // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterSessionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterSessionInitialized)
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
		it.Event = new(TrustManagementRouterSessionInitialized)
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
func (it *TrustManagementRouterSessionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterSessionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterSessionInitialized represents a SessionInitialized event raised by the TrustManagementRouter contract.
type TrustManagementRouterSessionInitialized struct {
	SessionKey common.Address
	MrEnclave  [32]byte
	HashedData [32]byte
	Caller     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSessionInitialized is a free log retrieval operation binding the contract event 0xec00d59ce40b347d1ecc55021a8dabf207d4b9084d2820fe25b69680eaab0f7d.
//
// Solidity: event SessionInitialized(address sessionKey, bytes32 mrEnclave, bytes32 hashedData, address indexed caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterSessionInitialized(opts *bind.FilterOpts, caller []common.Address) (*TrustManagementRouterSessionInitializedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "SessionInitialized", callerRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterSessionInitializedIterator{contract: _TrustManagementRouter.contract, event: "SessionInitialized", logs: logs, sub: sub}, nil
}

// WatchSessionInitialized is a free log subscription operation binding the contract event 0xec00d59ce40b347d1ecc55021a8dabf207d4b9084d2820fe25b69680eaab0f7d.
//
// Solidity: event SessionInitialized(address sessionKey, bytes32 mrEnclave, bytes32 hashedData, address indexed caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchSessionInitialized(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterSessionInitialized, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "SessionInitialized", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterSessionInitialized)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "SessionInitialized", log); err != nil {
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

// ParseSessionInitialized is a log parse operation binding the contract event 0xec00d59ce40b347d1ecc55021a8dabf207d4b9084d2820fe25b69680eaab0f7d.
//
// Solidity: event SessionInitialized(address sessionKey, bytes32 mrEnclave, bytes32 hashedData, address indexed caller)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseSessionInitialized(log types.Log) (*TrustManagementRouterSessionInitialized, error) {
	event := new(TrustManagementRouterSessionInitialized)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "SessionInitialized", log); err != nil {
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

// TrustManagementRouterWalletCreatedIterator is returned from FilterWalletCreated and is used to iterate over the raw logs and unpacked data for WalletCreated events raised by the TrustManagementRouter contract.
type TrustManagementRouterWalletCreatedIterator struct {
	Event *TrustManagementRouterWalletCreated // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterWalletCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterWalletCreated)
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
		it.Event = new(TrustManagementRouterWalletCreated)
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
func (it *TrustManagementRouterWalletCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterWalletCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterWalletCreated represents a WalletCreated event raised by the TrustManagementRouter contract.
type TrustManagementRouterWalletCreated struct {
	Wallet common.Address
	User   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWalletCreated is a free log retrieval operation binding the contract event 0x5b03bfed1c14a02bdeceb5fa582eb1a5765fc0bc64ca0e6af4c20afc9487f081.
//
// Solidity: event WalletCreated(address wallet, address user)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterWalletCreated(opts *bind.FilterOpts) (*TrustManagementRouterWalletCreatedIterator, error) {

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "WalletCreated")
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterWalletCreatedIterator{contract: _TrustManagementRouter.contract, event: "WalletCreated", logs: logs, sub: sub}, nil
}

// WatchWalletCreated is a free log subscription operation binding the contract event 0x5b03bfed1c14a02bdeceb5fa582eb1a5765fc0bc64ca0e6af4c20afc9487f081.
//
// Solidity: event WalletCreated(address wallet, address user)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchWalletCreated(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterWalletCreated) (event.Subscription, error) {

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "WalletCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterWalletCreated)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "WalletCreated", log); err != nil {
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

// ParseWalletCreated is a log parse operation binding the contract event 0x5b03bfed1c14a02bdeceb5fa582eb1a5765fc0bc64ca0e6af4c20afc9487f081.
//
// Solidity: event WalletCreated(address wallet, address user)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseWalletCreated(log types.Log) (*TrustManagementRouterWalletCreated, error) {
	event := new(TrustManagementRouterWalletCreated)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "WalletCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustManagementRouterWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the TrustManagementRouter contract.
type TrustManagementRouterWithdrawnIterator struct {
	Event *TrustManagementRouterWithdrawn // Event containing the contract specifics and raw log

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
func (it *TrustManagementRouterWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustManagementRouterWithdrawn)
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
		it.Event = new(TrustManagementRouterWithdrawn)
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
func (it *TrustManagementRouterWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustManagementRouterWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustManagementRouterWithdrawn represents a Withdrawn event raised by the TrustManagementRouter contract.
type TrustManagementRouterWithdrawn struct {
	Wallet    common.Address
	User      common.Address
	Token     common.Address
	Receiver  common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x554e88281043242e15d672fb9bd9320d0c660a025e8103b8ed42709235b1796e.
//
// Solidity: event Withdrawn(address wallet, address indexed user, address indexed token, address receiver, uint256 amount, uint256 feeAmount)
func (_TrustManagementRouter *TrustManagementRouterFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*TrustManagementRouterWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.FilterLogs(opts, "Withdrawn", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TrustManagementRouterWithdrawnIterator{contract: _TrustManagementRouter.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x554e88281043242e15d672fb9bd9320d0c660a025e8103b8ed42709235b1796e.
//
// Solidity: event Withdrawn(address wallet, address indexed user, address indexed token, address receiver, uint256 amount, uint256 feeAmount)
func (_TrustManagementRouter *TrustManagementRouterFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *TrustManagementRouterWithdrawn, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TrustManagementRouter.contract.WatchLogs(opts, "Withdrawn", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustManagementRouterWithdrawn)
				if err := _TrustManagementRouter.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x554e88281043242e15d672fb9bd9320d0c660a025e8103b8ed42709235b1796e.
//
// Solidity: event Withdrawn(address wallet, address indexed user, address indexed token, address receiver, uint256 amount, uint256 feeAmount)
func (_TrustManagementRouter *TrustManagementRouterFilterer) ParseWithdrawn(log types.Log) (*TrustManagementRouterWithdrawn, error) {
	event := new(TrustManagementRouterWithdrawn)
	if err := _TrustManagementRouter.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
