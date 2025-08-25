// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AavePool

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

// DataTypesCollateralConfig is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCollateralConfig struct {
	Ltv                  uint16
	LiquidationThreshold uint16
	LiquidationBonus     uint16
}

// DataTypesEModeCategoryBaseConfiguration is an auto generated low-level Go binding around an user-defined struct.
type DataTypesEModeCategoryBaseConfiguration struct {
	Ltv                  uint16
	LiquidationThreshold uint16
	LiquidationBonus     uint16
	Label                string
}

// DataTypesEModeCategoryLegacy is an auto generated low-level Go binding around an user-defined struct.
type DataTypesEModeCategoryLegacy struct {
	Ltv                  uint16
	LiquidationThreshold uint16
	LiquidationBonus     uint16
	PriceSource          common.Address
	Label                string
}

// DataTypesReserveConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveConfigurationMap struct {
	Data *big.Int
}

// DataTypesReserveDataLegacy is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveDataLegacy struct {
	Configuration               DataTypesReserveConfigurationMap
	LiquidityIndex              *big.Int
	CurrentLiquidityRate        *big.Int
	VariableBorrowIndex         *big.Int
	CurrentVariableBorrowRate   *big.Int
	CurrentStableBorrowRate     *big.Int
	LastUpdateTimestamp         *big.Int
	Id                          uint16
	ATokenAddress               common.Address
	StableDebtTokenAddress      common.Address
	VariableDebtTokenAddress    common.Address
	InterestRateStrategyAddress common.Address
	AccruedToTreasury           *big.Int
	Unbacked                    *big.Int
	IsolationModeTotalDebt      *big.Int
}

// DataTypesUserConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesUserConfigurationMap struct {
	Data *big.Int
}

// AavePoolMetaData contains all meta data concerning the AavePool contract.
var AavePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"contractIReserveInterestRateStrategy\",\"name\":\"interestRateStrategy_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetNotListed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotAToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPoolAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPoolConfigurator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPositionManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotUmbrella\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EModeCategoryReserved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddressesProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotValid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDataTypes.InterestRateMode\",\"name\":\"interestRateMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountCovered\",\"type\":\"uint256\"}],\"name\":\"DeficitCovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountCreated\",\"type\":\"uint256\"}],\"name\":\"DeficitCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDataTypes.InterestRateMode\",\"name\":\"interestRateMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDebt\",\"type\":\"uint256\"}],\"name\":\"IsolationModeTotalDebtUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidatedCollateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"receiveAToken\",\"type\":\"bool\"}],\"name\":\"LiquidationCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountMinted\",\"type\":\"uint256\"}],\"name\":\"MintedToTreasury\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"positionManager\",\"type\":\"address\"}],\"name\":\"PositionManagerApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"positionManager\",\"type\":\"address\"}],\"name\":\"PositionManagerRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"repayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useATokens\",\"type\":\"bool\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"ReserveDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"categoryId\",\"type\":\"uint8\"}],\"name\":\"UserEModeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLASHLOAN_PREMIUM_TOTAL\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLASHLOAN_PREMIUM_TO_PROTOCOL\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUMBER_RESERVES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_INTEREST_RATE_STRATEGY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UMBRELLA\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"positionManager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approve\",\"type\":\"bool\"}],\"name\":\"approvePositionManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"internalType\":\"structDataTypes.EModeCategoryBaseConfiguration\",\"name\":\"category\",\"type\":\"tuple\"}],\"name\":\"configureEModeCategory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"borrowableBitmap\",\"type\":\"uint128\"}],\"name\":\"configureEModeCategoryBorrowableBitmap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"collateralBitmap\",\"type\":\"uint128\"}],\"name\":\"configureEModeCategoryCollateralBitmap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"dropReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"eliminateReserveDeficit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scaledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledBalanceFromBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledBalanceToBefore\",\"type\":\"uint256\"}],\"name\":\"finalizeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"interestRateModes\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"flashLoanSimple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBorrowLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getEModeCategoryBorrowableBitmap\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getEModeCategoryCollateralBitmap\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getEModeCategoryCollateralConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationBonus\",\"type\":\"uint16\"}],\"internalType\":\"structDataTypes.CollateralConfig\",\"name\":\"res\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getEModeCategoryData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceSource\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"internalType\":\"structDataTypes.EModeCategoryLegacy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getEModeCategoryLabel\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEModeLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFlashLoanLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getLiquidationGracePeriod\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidationLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveAToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"getReserveAddressById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentLiquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentVariableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentStableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accruedToTreasury\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"unbacked\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"isolationModeTotalDebt\",\"type\":\"uint128\"}],\"internalType\":\"structDataTypes.ReserveDataLegacy\",\"name\":\"res\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveDeficit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedVariableDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveVariableDebtToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupplyLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserEMode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getVirtualUnderlyingBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtAddress\",\"type\":\"address\"}],\"name\":\"initReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"positionManager\",\"type\":\"address\"}],\"name\":\"isApprovedPositionManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"receiveAToken\",\"type\":\"bool\"}],\"name\":\"liquidationCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"mintToTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"renouncePositionManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"}],\"name\":\"repayWithATokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"permitV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"permitR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"permitS\",\"type\":\"bytes32\"}],\"name\":\"repayWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"resetIsolationModeTotalDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"}],\"name\":\"setConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"until\",\"type\":\"uint40\"}],\"name\":\"setLiquidationGracePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"categoryId\",\"type\":\"uint8\"}],\"name\":\"setUserEMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"categoryId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"setUserEModeOnBehalfOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useAsCollateral\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"setUserUseReserveAsCollateralOnBehalfOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"supply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"permitV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"permitR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"permitS\",\"type\":\"bytes32\"}],\"name\":\"supplyWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"syncIndexesState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"syncRatesState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"flashLoanPremium\",\"type\":\"uint128\"}],\"name\":\"updateFlashloanPremium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AavePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use AavePoolMetaData.ABI instead.
var AavePoolABI = AavePoolMetaData.ABI

// AavePool is an auto generated Go binding around an Ethereum contract.
type AavePool struct {
	AavePoolCaller     // Read-only binding to the contract
	AavePoolTransactor // Write-only binding to the contract
	AavePoolFilterer   // Log filterer for contract events
}

// AavePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AavePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AavePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AavePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AavePoolSession struct {
	Contract     *AavePool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AavePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AavePoolCallerSession struct {
	Contract *AavePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AavePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AavePoolTransactorSession struct {
	Contract     *AavePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AavePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AavePoolRaw struct {
	Contract *AavePool // Generic contract binding to access the raw methods on
}

// AavePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AavePoolCallerRaw struct {
	Contract *AavePoolCaller // Generic read-only contract binding to access the raw methods on
}

// AavePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AavePoolTransactorRaw struct {
	Contract *AavePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAavePool creates a new instance of AavePool, bound to a specific deployed contract.
func NewAavePool(address common.Address, backend bind.ContractBackend) (*AavePool, error) {
	contract, err := bindAavePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AavePool{AavePoolCaller: AavePoolCaller{contract: contract}, AavePoolTransactor: AavePoolTransactor{contract: contract}, AavePoolFilterer: AavePoolFilterer{contract: contract}}, nil
}

// NewAavePoolCaller creates a new read-only instance of AavePool, bound to a specific deployed contract.
func NewAavePoolCaller(address common.Address, caller bind.ContractCaller) (*AavePoolCaller, error) {
	contract, err := bindAavePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AavePoolCaller{contract: contract}, nil
}

// NewAavePoolTransactor creates a new write-only instance of AavePool, bound to a specific deployed contract.
func NewAavePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AavePoolTransactor, error) {
	contract, err := bindAavePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AavePoolTransactor{contract: contract}, nil
}

// NewAavePoolFilterer creates a new log filterer instance of AavePool, bound to a specific deployed contract.
func NewAavePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AavePoolFilterer, error) {
	contract, err := bindAavePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AavePoolFilterer{contract: contract}, nil
}

// bindAavePool binds a generic wrapper to an already deployed contract.
func bindAavePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AavePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AavePool *AavePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AavePool.Contract.AavePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AavePool *AavePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AavePool.Contract.AavePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AavePool *AavePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AavePool.Contract.AavePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AavePool *AavePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AavePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AavePool *AavePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AavePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AavePool *AavePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AavePool.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AavePool *AavePoolCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AavePool *AavePoolSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AavePool.Contract.ADDRESSESPROVIDER(&_AavePool.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AavePool *AavePoolCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AavePool.Contract.ADDRESSESPROVIDER(&_AavePool.CallOpts)
}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint128)
func (_AavePool *AavePoolCaller) FLASHLOANPREMIUMTOTAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "FLASHLOAN_PREMIUM_TOTAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint128)
func (_AavePool *AavePoolSession) FLASHLOANPREMIUMTOTAL() (*big.Int, error) {
	return _AavePool.Contract.FLASHLOANPREMIUMTOTAL(&_AavePool.CallOpts)
}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint128)
func (_AavePool *AavePoolCallerSession) FLASHLOANPREMIUMTOTAL() (*big.Int, error) {
	return _AavePool.Contract.FLASHLOANPREMIUMTOTAL(&_AavePool.CallOpts)
}

// FLASHLOANPREMIUMTOPROTOCOL is a free data retrieval call binding the contract method 0x6a99c036.
//
// Solidity: function FLASHLOAN_PREMIUM_TO_PROTOCOL() view returns(uint128)
func (_AavePool *AavePoolCaller) FLASHLOANPREMIUMTOPROTOCOL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "FLASHLOAN_PREMIUM_TO_PROTOCOL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLASHLOANPREMIUMTOPROTOCOL is a free data retrieval call binding the contract method 0x6a99c036.
//
// Solidity: function FLASHLOAN_PREMIUM_TO_PROTOCOL() view returns(uint128)
func (_AavePool *AavePoolSession) FLASHLOANPREMIUMTOPROTOCOL() (*big.Int, error) {
	return _AavePool.Contract.FLASHLOANPREMIUMTOPROTOCOL(&_AavePool.CallOpts)
}

// FLASHLOANPREMIUMTOPROTOCOL is a free data retrieval call binding the contract method 0x6a99c036.
//
// Solidity: function FLASHLOAN_PREMIUM_TO_PROTOCOL() view returns(uint128)
func (_AavePool *AavePoolCallerSession) FLASHLOANPREMIUMTOPROTOCOL() (*big.Int, error) {
	return _AavePool.Contract.FLASHLOANPREMIUMTOPROTOCOL(&_AavePool.CallOpts)
}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint16)
func (_AavePool *AavePoolCaller) MAXNUMBERRESERVES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "MAX_NUMBER_RESERVES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint16)
func (_AavePool *AavePoolSession) MAXNUMBERRESERVES() (uint16, error) {
	return _AavePool.Contract.MAXNUMBERRESERVES(&_AavePool.CallOpts)
}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint16)
func (_AavePool *AavePoolCallerSession) MAXNUMBERRESERVES() (uint16, error) {
	return _AavePool.Contract.MAXNUMBERRESERVES(&_AavePool.CallOpts)
}

// POOLREVISION is a free data retrieval call binding the contract method 0x0148170e.
//
// Solidity: function POOL_REVISION() view returns(uint256)
func (_AavePool *AavePoolCaller) POOLREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "POOL_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// POOLREVISION is a free data retrieval call binding the contract method 0x0148170e.
//
// Solidity: function POOL_REVISION() view returns(uint256)
func (_AavePool *AavePoolSession) POOLREVISION() (*big.Int, error) {
	return _AavePool.Contract.POOLREVISION(&_AavePool.CallOpts)
}

// POOLREVISION is a free data retrieval call binding the contract method 0x0148170e.
//
// Solidity: function POOL_REVISION() view returns(uint256)
func (_AavePool *AavePoolCallerSession) POOLREVISION() (*big.Int, error) {
	return _AavePool.Contract.POOLREVISION(&_AavePool.CallOpts)
}

// RESERVEINTERESTRATESTRATEGY is a free data retrieval call binding the contract method 0x1b8feb0e.
//
// Solidity: function RESERVE_INTEREST_RATE_STRATEGY() view returns(address)
func (_AavePool *AavePoolCaller) RESERVEINTERESTRATESTRATEGY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "RESERVE_INTEREST_RATE_STRATEGY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RESERVEINTERESTRATESTRATEGY is a free data retrieval call binding the contract method 0x1b8feb0e.
//
// Solidity: function RESERVE_INTEREST_RATE_STRATEGY() view returns(address)
func (_AavePool *AavePoolSession) RESERVEINTERESTRATESTRATEGY() (common.Address, error) {
	return _AavePool.Contract.RESERVEINTERESTRATESTRATEGY(&_AavePool.CallOpts)
}

// RESERVEINTERESTRATESTRATEGY is a free data retrieval call binding the contract method 0x1b8feb0e.
//
// Solidity: function RESERVE_INTEREST_RATE_STRATEGY() view returns(address)
func (_AavePool *AavePoolCallerSession) RESERVEINTERESTRATESTRATEGY() (common.Address, error) {
	return _AavePool.Contract.RESERVEINTERESTRATESTRATEGY(&_AavePool.CallOpts)
}

// UMBRELLA is a free data retrieval call binding the contract method 0x71459c15.
//
// Solidity: function UMBRELLA() view returns(bytes32)
func (_AavePool *AavePoolCaller) UMBRELLA(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "UMBRELLA")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UMBRELLA is a free data retrieval call binding the contract method 0x71459c15.
//
// Solidity: function UMBRELLA() view returns(bytes32)
func (_AavePool *AavePoolSession) UMBRELLA() ([32]byte, error) {
	return _AavePool.Contract.UMBRELLA(&_AavePool.CallOpts)
}

// UMBRELLA is a free data retrieval call binding the contract method 0x71459c15.
//
// Solidity: function UMBRELLA() view returns(bytes32)
func (_AavePool *AavePoolCallerSession) UMBRELLA() ([32]byte, error) {
	return _AavePool.Contract.UMBRELLA(&_AavePool.CallOpts)
}

// GetBorrowLogic is a free data retrieval call binding the contract method 0x2be29fa7.
//
// Solidity: function getBorrowLogic() pure returns(address)
func (_AavePool *AavePoolCaller) GetBorrowLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getBorrowLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBorrowLogic is a free data retrieval call binding the contract method 0x2be29fa7.
//
// Solidity: function getBorrowLogic() pure returns(address)
func (_AavePool *AavePoolSession) GetBorrowLogic() (common.Address, error) {
	return _AavePool.Contract.GetBorrowLogic(&_AavePool.CallOpts)
}

// GetBorrowLogic is a free data retrieval call binding the contract method 0x2be29fa7.
//
// Solidity: function getBorrowLogic() pure returns(address)
func (_AavePool *AavePoolCallerSession) GetBorrowLogic() (common.Address, error) {
	return _AavePool.Contract.GetBorrowLogic(&_AavePool.CallOpts)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_AavePool *AavePoolCaller) GetConfiguration(opts *bind.CallOpts, asset common.Address) (DataTypesReserveConfigurationMap, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getConfiguration", asset)

	if err != nil {
		return *new(DataTypesReserveConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveConfigurationMap)).(*DataTypesReserveConfigurationMap)

	return out0, err

}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_AavePool *AavePoolSession) GetConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _AavePool.Contract.GetConfiguration(&_AavePool.CallOpts, asset)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_AavePool *AavePoolCallerSession) GetConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _AavePool.Contract.GetConfiguration(&_AavePool.CallOpts, asset)
}

// GetEModeCategoryBorrowableBitmap is a free data retrieval call binding the contract method 0x903a2c71.
//
// Solidity: function getEModeCategoryBorrowableBitmap(uint8 id) view returns(uint128)
func (_AavePool *AavePoolCaller) GetEModeCategoryBorrowableBitmap(opts *bind.CallOpts, id uint8) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getEModeCategoryBorrowableBitmap", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEModeCategoryBorrowableBitmap is a free data retrieval call binding the contract method 0x903a2c71.
//
// Solidity: function getEModeCategoryBorrowableBitmap(uint8 id) view returns(uint128)
func (_AavePool *AavePoolSession) GetEModeCategoryBorrowableBitmap(id uint8) (*big.Int, error) {
	return _AavePool.Contract.GetEModeCategoryBorrowableBitmap(&_AavePool.CallOpts, id)
}

// GetEModeCategoryBorrowableBitmap is a free data retrieval call binding the contract method 0x903a2c71.
//
// Solidity: function getEModeCategoryBorrowableBitmap(uint8 id) view returns(uint128)
func (_AavePool *AavePoolCallerSession) GetEModeCategoryBorrowableBitmap(id uint8) (*big.Int, error) {
	return _AavePool.Contract.GetEModeCategoryBorrowableBitmap(&_AavePool.CallOpts, id)
}

// GetEModeCategoryCollateralBitmap is a free data retrieval call binding the contract method 0xb0771dba.
//
// Solidity: function getEModeCategoryCollateralBitmap(uint8 id) view returns(uint128)
func (_AavePool *AavePoolCaller) GetEModeCategoryCollateralBitmap(opts *bind.CallOpts, id uint8) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getEModeCategoryCollateralBitmap", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEModeCategoryCollateralBitmap is a free data retrieval call binding the contract method 0xb0771dba.
//
// Solidity: function getEModeCategoryCollateralBitmap(uint8 id) view returns(uint128)
func (_AavePool *AavePoolSession) GetEModeCategoryCollateralBitmap(id uint8) (*big.Int, error) {
	return _AavePool.Contract.GetEModeCategoryCollateralBitmap(&_AavePool.CallOpts, id)
}

// GetEModeCategoryCollateralBitmap is a free data retrieval call binding the contract method 0xb0771dba.
//
// Solidity: function getEModeCategoryCollateralBitmap(uint8 id) view returns(uint128)
func (_AavePool *AavePoolCallerSession) GetEModeCategoryCollateralBitmap(id uint8) (*big.Int, error) {
	return _AavePool.Contract.GetEModeCategoryCollateralBitmap(&_AavePool.CallOpts, id)
}

// GetEModeCategoryCollateralConfig is a free data retrieval call binding the contract method 0xb286f467.
//
// Solidity: function getEModeCategoryCollateralConfig(uint8 id) view returns((uint16,uint16,uint16) res)
func (_AavePool *AavePoolCaller) GetEModeCategoryCollateralConfig(opts *bind.CallOpts, id uint8) (DataTypesCollateralConfig, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getEModeCategoryCollateralConfig", id)

	if err != nil {
		return *new(DataTypesCollateralConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCollateralConfig)).(*DataTypesCollateralConfig)

	return out0, err

}

// GetEModeCategoryCollateralConfig is a free data retrieval call binding the contract method 0xb286f467.
//
// Solidity: function getEModeCategoryCollateralConfig(uint8 id) view returns((uint16,uint16,uint16) res)
func (_AavePool *AavePoolSession) GetEModeCategoryCollateralConfig(id uint8) (DataTypesCollateralConfig, error) {
	return _AavePool.Contract.GetEModeCategoryCollateralConfig(&_AavePool.CallOpts, id)
}

// GetEModeCategoryCollateralConfig is a free data retrieval call binding the contract method 0xb286f467.
//
// Solidity: function getEModeCategoryCollateralConfig(uint8 id) view returns((uint16,uint16,uint16) res)
func (_AavePool *AavePoolCallerSession) GetEModeCategoryCollateralConfig(id uint8) (DataTypesCollateralConfig, error) {
	return _AavePool.Contract.GetEModeCategoryCollateralConfig(&_AavePool.CallOpts, id)
}

// GetEModeCategoryData is a free data retrieval call binding the contract method 0x6c6f6ae1.
//
// Solidity: function getEModeCategoryData(uint8 id) view returns((uint16,uint16,uint16,address,string))
func (_AavePool *AavePoolCaller) GetEModeCategoryData(opts *bind.CallOpts, id uint8) (DataTypesEModeCategoryLegacy, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getEModeCategoryData", id)

	if err != nil {
		return *new(DataTypesEModeCategoryLegacy), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesEModeCategoryLegacy)).(*DataTypesEModeCategoryLegacy)

	return out0, err

}

// GetEModeCategoryData is a free data retrieval call binding the contract method 0x6c6f6ae1.
//
// Solidity: function getEModeCategoryData(uint8 id) view returns((uint16,uint16,uint16,address,string))
func (_AavePool *AavePoolSession) GetEModeCategoryData(id uint8) (DataTypesEModeCategoryLegacy, error) {
	return _AavePool.Contract.GetEModeCategoryData(&_AavePool.CallOpts, id)
}

// GetEModeCategoryData is a free data retrieval call binding the contract method 0x6c6f6ae1.
//
// Solidity: function getEModeCategoryData(uint8 id) view returns((uint16,uint16,uint16,address,string))
func (_AavePool *AavePoolCallerSession) GetEModeCategoryData(id uint8) (DataTypesEModeCategoryLegacy, error) {
	return _AavePool.Contract.GetEModeCategoryData(&_AavePool.CallOpts, id)
}

// GetEModeCategoryLabel is a free data retrieval call binding the contract method 0x2083e183.
//
// Solidity: function getEModeCategoryLabel(uint8 id) view returns(string)
func (_AavePool *AavePoolCaller) GetEModeCategoryLabel(opts *bind.CallOpts, id uint8) (string, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getEModeCategoryLabel", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEModeCategoryLabel is a free data retrieval call binding the contract method 0x2083e183.
//
// Solidity: function getEModeCategoryLabel(uint8 id) view returns(string)
func (_AavePool *AavePoolSession) GetEModeCategoryLabel(id uint8) (string, error) {
	return _AavePool.Contract.GetEModeCategoryLabel(&_AavePool.CallOpts, id)
}

// GetEModeCategoryLabel is a free data retrieval call binding the contract method 0x2083e183.
//
// Solidity: function getEModeCategoryLabel(uint8 id) view returns(string)
func (_AavePool *AavePoolCallerSession) GetEModeCategoryLabel(id uint8) (string, error) {
	return _AavePool.Contract.GetEModeCategoryLabel(&_AavePool.CallOpts, id)
}

// GetEModeLogic is a free data retrieval call binding the contract method 0xf32b9a73.
//
// Solidity: function getEModeLogic() pure returns(address)
func (_AavePool *AavePoolCaller) GetEModeLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getEModeLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEModeLogic is a free data retrieval call binding the contract method 0xf32b9a73.
//
// Solidity: function getEModeLogic() pure returns(address)
func (_AavePool *AavePoolSession) GetEModeLogic() (common.Address, error) {
	return _AavePool.Contract.GetEModeLogic(&_AavePool.CallOpts)
}

// GetEModeLogic is a free data retrieval call binding the contract method 0xf32b9a73.
//
// Solidity: function getEModeLogic() pure returns(address)
func (_AavePool *AavePoolCallerSession) GetEModeLogic() (common.Address, error) {
	return _AavePool.Contract.GetEModeLogic(&_AavePool.CallOpts)
}

// GetFlashLoanLogic is a free data retrieval call binding the contract method 0x348fde0f.
//
// Solidity: function getFlashLoanLogic() pure returns(address)
func (_AavePool *AavePoolCaller) GetFlashLoanLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getFlashLoanLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFlashLoanLogic is a free data retrieval call binding the contract method 0x348fde0f.
//
// Solidity: function getFlashLoanLogic() pure returns(address)
func (_AavePool *AavePoolSession) GetFlashLoanLogic() (common.Address, error) {
	return _AavePool.Contract.GetFlashLoanLogic(&_AavePool.CallOpts)
}

// GetFlashLoanLogic is a free data retrieval call binding the contract method 0x348fde0f.
//
// Solidity: function getFlashLoanLogic() pure returns(address)
func (_AavePool *AavePoolCallerSession) GetFlashLoanLogic() (common.Address, error) {
	return _AavePool.Contract.GetFlashLoanLogic(&_AavePool.CallOpts)
}

// GetLiquidationGracePeriod is a free data retrieval call binding the contract method 0x5c9a8b18.
//
// Solidity: function getLiquidationGracePeriod(address asset) view returns(uint40)
func (_AavePool *AavePoolCaller) GetLiquidationGracePeriod(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getLiquidationGracePeriod", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidationGracePeriod is a free data retrieval call binding the contract method 0x5c9a8b18.
//
// Solidity: function getLiquidationGracePeriod(address asset) view returns(uint40)
func (_AavePool *AavePoolSession) GetLiquidationGracePeriod(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetLiquidationGracePeriod(&_AavePool.CallOpts, asset)
}

// GetLiquidationGracePeriod is a free data retrieval call binding the contract method 0x5c9a8b18.
//
// Solidity: function getLiquidationGracePeriod(address asset) view returns(uint40)
func (_AavePool *AavePoolCallerSession) GetLiquidationGracePeriod(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetLiquidationGracePeriod(&_AavePool.CallOpts, asset)
}

// GetLiquidationLogic is a free data retrieval call binding the contract method 0x911a3413.
//
// Solidity: function getLiquidationLogic() pure returns(address)
func (_AavePool *AavePoolCaller) GetLiquidationLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getLiquidationLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLiquidationLogic is a free data retrieval call binding the contract method 0x911a3413.
//
// Solidity: function getLiquidationLogic() pure returns(address)
func (_AavePool *AavePoolSession) GetLiquidationLogic() (common.Address, error) {
	return _AavePool.Contract.GetLiquidationLogic(&_AavePool.CallOpts)
}

// GetLiquidationLogic is a free data retrieval call binding the contract method 0x911a3413.
//
// Solidity: function getLiquidationLogic() pure returns(address)
func (_AavePool *AavePoolCallerSession) GetLiquidationLogic() (common.Address, error) {
	return _AavePool.Contract.GetLiquidationLogic(&_AavePool.CallOpts)
}

// GetPoolLogic is a free data retrieval call binding the contract method 0xd3350155.
//
// Solidity: function getPoolLogic() pure returns(address)
func (_AavePool *AavePoolCaller) GetPoolLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getPoolLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolLogic is a free data retrieval call binding the contract method 0xd3350155.
//
// Solidity: function getPoolLogic() pure returns(address)
func (_AavePool *AavePoolSession) GetPoolLogic() (common.Address, error) {
	return _AavePool.Contract.GetPoolLogic(&_AavePool.CallOpts)
}

// GetPoolLogic is a free data retrieval call binding the contract method 0xd3350155.
//
// Solidity: function getPoolLogic() pure returns(address)
func (_AavePool *AavePoolCallerSession) GetPoolLogic() (common.Address, error) {
	return _AavePool.Contract.GetPoolLogic(&_AavePool.CallOpts)
}

// GetReserveAToken is a free data retrieval call binding the contract method 0xcff027d9.
//
// Solidity: function getReserveAToken(address asset) view returns(address)
func (_AavePool *AavePoolCaller) GetReserveAToken(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReserveAToken", asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReserveAToken is a free data retrieval call binding the contract method 0xcff027d9.
//
// Solidity: function getReserveAToken(address asset) view returns(address)
func (_AavePool *AavePoolSession) GetReserveAToken(asset common.Address) (common.Address, error) {
	return _AavePool.Contract.GetReserveAToken(&_AavePool.CallOpts, asset)
}

// GetReserveAToken is a free data retrieval call binding the contract method 0xcff027d9.
//
// Solidity: function getReserveAToken(address asset) view returns(address)
func (_AavePool *AavePoolCallerSession) GetReserveAToken(asset common.Address) (common.Address, error) {
	return _AavePool.Contract.GetReserveAToken(&_AavePool.CallOpts, asset)
}

// GetReserveAddressById is a free data retrieval call binding the contract method 0x52751797.
//
// Solidity: function getReserveAddressById(uint16 id) view returns(address)
func (_AavePool *AavePoolCaller) GetReserveAddressById(opts *bind.CallOpts, id uint16) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReserveAddressById", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReserveAddressById is a free data retrieval call binding the contract method 0x52751797.
//
// Solidity: function getReserveAddressById(uint16 id) view returns(address)
func (_AavePool *AavePoolSession) GetReserveAddressById(id uint16) (common.Address, error) {
	return _AavePool.Contract.GetReserveAddressById(&_AavePool.CallOpts, id)
}

// GetReserveAddressById is a free data retrieval call binding the contract method 0x52751797.
//
// Solidity: function getReserveAddressById(uint16 id) view returns(address)
func (_AavePool *AavePoolCallerSession) GetReserveAddressById(id uint16) (common.Address, error) {
	return _AavePool.Contract.GetReserveAddressById(&_AavePool.CallOpts, id)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128) res)
func (_AavePool *AavePoolCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (DataTypesReserveDataLegacy, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReserveData", asset)

	if err != nil {
		return *new(DataTypesReserveDataLegacy), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveDataLegacy)).(*DataTypesReserveDataLegacy)

	return out0, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128) res)
func (_AavePool *AavePoolSession) GetReserveData(asset common.Address) (DataTypesReserveDataLegacy, error) {
	return _AavePool.Contract.GetReserveData(&_AavePool.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128) res)
func (_AavePool *AavePoolCallerSession) GetReserveData(asset common.Address) (DataTypesReserveDataLegacy, error) {
	return _AavePool.Contract.GetReserveData(&_AavePool.CallOpts, asset)
}

// GetReserveDeficit is a free data retrieval call binding the contract method 0xc952485d.
//
// Solidity: function getReserveDeficit(address asset) view returns(uint256)
func (_AavePool *AavePoolCaller) GetReserveDeficit(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReserveDeficit", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveDeficit is a free data retrieval call binding the contract method 0xc952485d.
//
// Solidity: function getReserveDeficit(address asset) view returns(uint256)
func (_AavePool *AavePoolSession) GetReserveDeficit(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetReserveDeficit(&_AavePool.CallOpts, asset)
}

// GetReserveDeficit is a free data retrieval call binding the contract method 0xc952485d.
//
// Solidity: function getReserveDeficit(address asset) view returns(uint256)
func (_AavePool *AavePoolCallerSession) GetReserveDeficit(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetReserveDeficit(&_AavePool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_AavePool *AavePoolCaller) GetReserveNormalizedIncome(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReserveNormalizedIncome", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_AavePool *AavePoolSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetReserveNormalizedIncome(&_AavePool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_AavePool *AavePoolCallerSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetReserveNormalizedIncome(&_AavePool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_AavePool *AavePoolCaller) GetReserveNormalizedVariableDebt(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReserveNormalizedVariableDebt", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_AavePool *AavePoolSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetReserveNormalizedVariableDebt(&_AavePool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_AavePool *AavePoolCallerSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetReserveNormalizedVariableDebt(&_AavePool.CallOpts, asset)
}

// GetReserveVariableDebtToken is a free data retrieval call binding the contract method 0x365090a0.
//
// Solidity: function getReserveVariableDebtToken(address asset) view returns(address)
func (_AavePool *AavePoolCaller) GetReserveVariableDebtToken(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReserveVariableDebtToken", asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReserveVariableDebtToken is a free data retrieval call binding the contract method 0x365090a0.
//
// Solidity: function getReserveVariableDebtToken(address asset) view returns(address)
func (_AavePool *AavePoolSession) GetReserveVariableDebtToken(asset common.Address) (common.Address, error) {
	return _AavePool.Contract.GetReserveVariableDebtToken(&_AavePool.CallOpts, asset)
}

// GetReserveVariableDebtToken is a free data retrieval call binding the contract method 0x365090a0.
//
// Solidity: function getReserveVariableDebtToken(address asset) view returns(address)
func (_AavePool *AavePoolCallerSession) GetReserveVariableDebtToken(asset common.Address) (common.Address, error) {
	return _AavePool.Contract.GetReserveVariableDebtToken(&_AavePool.CallOpts, asset)
}

// GetReservesCount is a free data retrieval call binding the contract method 0x72218d04.
//
// Solidity: function getReservesCount() view returns(uint256)
func (_AavePool *AavePoolCaller) GetReservesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReservesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReservesCount is a free data retrieval call binding the contract method 0x72218d04.
//
// Solidity: function getReservesCount() view returns(uint256)
func (_AavePool *AavePoolSession) GetReservesCount() (*big.Int, error) {
	return _AavePool.Contract.GetReservesCount(&_AavePool.CallOpts)
}

// GetReservesCount is a free data retrieval call binding the contract method 0x72218d04.
//
// Solidity: function getReservesCount() view returns(uint256)
func (_AavePool *AavePoolCallerSession) GetReservesCount() (*big.Int, error) {
	return _AavePool.Contract.GetReservesCount(&_AavePool.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AavePool *AavePoolCaller) GetReservesList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReservesList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AavePool *AavePoolSession) GetReservesList() ([]common.Address, error) {
	return _AavePool.Contract.GetReservesList(&_AavePool.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AavePool *AavePoolCallerSession) GetReservesList() ([]common.Address, error) {
	return _AavePool.Contract.GetReservesList(&_AavePool.CallOpts)
}

// GetSupplyLogic is a free data retrieval call binding the contract method 0x870e7744.
//
// Solidity: function getSupplyLogic() pure returns(address)
func (_AavePool *AavePoolCaller) GetSupplyLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getSupplyLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSupplyLogic is a free data retrieval call binding the contract method 0x870e7744.
//
// Solidity: function getSupplyLogic() pure returns(address)
func (_AavePool *AavePoolSession) GetSupplyLogic() (common.Address, error) {
	return _AavePool.Contract.GetSupplyLogic(&_AavePool.CallOpts)
}

// GetSupplyLogic is a free data retrieval call binding the contract method 0x870e7744.
//
// Solidity: function getSupplyLogic() pure returns(address)
func (_AavePool *AavePoolCallerSession) GetSupplyLogic() (common.Address, error) {
	return _AavePool.Contract.GetSupplyLogic(&_AavePool.CallOpts)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AavePool *AavePoolCaller) GetUserAccountData(opts *bind.CallOpts, user common.Address) (struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getUserAccountData", user)

	outstruct := new(struct {
		TotalCollateralBase         *big.Int
		TotalDebtBase               *big.Int
		AvailableBorrowsBase        *big.Int
		CurrentLiquidationThreshold *big.Int
		Ltv                         *big.Int
		HealthFactor                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCollateralBase = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalDebtBase = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsBase = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentLiquidationThreshold = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HealthFactor = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AavePool *AavePoolSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _AavePool.Contract.GetUserAccountData(&_AavePool.CallOpts, user)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AavePool *AavePoolCallerSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _AavePool.Contract.GetUserAccountData(&_AavePool.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AavePool *AavePoolCaller) GetUserConfiguration(opts *bind.CallOpts, user common.Address) (DataTypesUserConfigurationMap, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getUserConfiguration", user)

	if err != nil {
		return *new(DataTypesUserConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesUserConfigurationMap)).(*DataTypesUserConfigurationMap)

	return out0, err

}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AavePool *AavePoolSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _AavePool.Contract.GetUserConfiguration(&_AavePool.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AavePool *AavePoolCallerSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _AavePool.Contract.GetUserConfiguration(&_AavePool.CallOpts, user)
}

// GetUserEMode is a free data retrieval call binding the contract method 0xeddf1b79.
//
// Solidity: function getUserEMode(address user) view returns(uint256)
func (_AavePool *AavePoolCaller) GetUserEMode(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getUserEMode", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserEMode is a free data retrieval call binding the contract method 0xeddf1b79.
//
// Solidity: function getUserEMode(address user) view returns(uint256)
func (_AavePool *AavePoolSession) GetUserEMode(user common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetUserEMode(&_AavePool.CallOpts, user)
}

// GetUserEMode is a free data retrieval call binding the contract method 0xeddf1b79.
//
// Solidity: function getUserEMode(address user) view returns(uint256)
func (_AavePool *AavePoolCallerSession) GetUserEMode(user common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetUserEMode(&_AavePool.CallOpts, user)
}

// GetVirtualUnderlyingBalance is a free data retrieval call binding the contract method 0x6fb07f96.
//
// Solidity: function getVirtualUnderlyingBalance(address asset) view returns(uint128)
func (_AavePool *AavePoolCaller) GetVirtualUnderlyingBalance(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getVirtualUnderlyingBalance", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualUnderlyingBalance is a free data retrieval call binding the contract method 0x6fb07f96.
//
// Solidity: function getVirtualUnderlyingBalance(address asset) view returns(uint128)
func (_AavePool *AavePoolSession) GetVirtualUnderlyingBalance(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetVirtualUnderlyingBalance(&_AavePool.CallOpts, asset)
}

// GetVirtualUnderlyingBalance is a free data retrieval call binding the contract method 0x6fb07f96.
//
// Solidity: function getVirtualUnderlyingBalance(address asset) view returns(uint128)
func (_AavePool *AavePoolCallerSession) GetVirtualUnderlyingBalance(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetVirtualUnderlyingBalance(&_AavePool.CallOpts, asset)
}

// IsApprovedPositionManager is a free data retrieval call binding the contract method 0xf9c2bd87.
//
// Solidity: function isApprovedPositionManager(address user, address positionManager) view returns(bool)
func (_AavePool *AavePoolCaller) IsApprovedPositionManager(opts *bind.CallOpts, user common.Address, positionManager common.Address) (bool, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "isApprovedPositionManager", user, positionManager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedPositionManager is a free data retrieval call binding the contract method 0xf9c2bd87.
//
// Solidity: function isApprovedPositionManager(address user, address positionManager) view returns(bool)
func (_AavePool *AavePoolSession) IsApprovedPositionManager(user common.Address, positionManager common.Address) (bool, error) {
	return _AavePool.Contract.IsApprovedPositionManager(&_AavePool.CallOpts, user, positionManager)
}

// IsApprovedPositionManager is a free data retrieval call binding the contract method 0xf9c2bd87.
//
// Solidity: function isApprovedPositionManager(address user, address positionManager) view returns(bool)
func (_AavePool *AavePoolCallerSession) IsApprovedPositionManager(user common.Address, positionManager common.Address) (bool, error) {
	return _AavePool.Contract.IsApprovedPositionManager(&_AavePool.CallOpts, user, positionManager)
}

// ApprovePositionManager is a paid mutator transaction binding the contract method 0xb8caa7c5.
//
// Solidity: function approvePositionManager(address positionManager, bool approve) returns()
func (_AavePool *AavePoolTransactor) ApprovePositionManager(opts *bind.TransactOpts, positionManager common.Address, approve bool) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "approvePositionManager", positionManager, approve)
}

// ApprovePositionManager is a paid mutator transaction binding the contract method 0xb8caa7c5.
//
// Solidity: function approvePositionManager(address positionManager, bool approve) returns()
func (_AavePool *AavePoolSession) ApprovePositionManager(positionManager common.Address, approve bool) (*types.Transaction, error) {
	return _AavePool.Contract.ApprovePositionManager(&_AavePool.TransactOpts, positionManager, approve)
}

// ApprovePositionManager is a paid mutator transaction binding the contract method 0xb8caa7c5.
//
// Solidity: function approvePositionManager(address positionManager, bool approve) returns()
func (_AavePool *AavePoolTransactorSession) ApprovePositionManager(positionManager common.Address, approve bool) (*types.Transaction, error) {
	return _AavePool.Contract.ApprovePositionManager(&_AavePool.TransactOpts, positionManager, approve)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_AavePool *AavePoolTransactor) Borrow(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "borrow", asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_AavePool *AavePoolSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Borrow(&_AavePool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_AavePool *AavePoolTransactorSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Borrow(&_AavePool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// ConfigureEModeCategory is a paid mutator transaction binding the contract method 0x7b75d7f4.
//
// Solidity: function configureEModeCategory(uint8 id, (uint16,uint16,uint16,string) category) returns()
func (_AavePool *AavePoolTransactor) ConfigureEModeCategory(opts *bind.TransactOpts, id uint8, category DataTypesEModeCategoryBaseConfiguration) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "configureEModeCategory", id, category)
}

// ConfigureEModeCategory is a paid mutator transaction binding the contract method 0x7b75d7f4.
//
// Solidity: function configureEModeCategory(uint8 id, (uint16,uint16,uint16,string) category) returns()
func (_AavePool *AavePoolSession) ConfigureEModeCategory(id uint8, category DataTypesEModeCategoryBaseConfiguration) (*types.Transaction, error) {
	return _AavePool.Contract.ConfigureEModeCategory(&_AavePool.TransactOpts, id, category)
}

// ConfigureEModeCategory is a paid mutator transaction binding the contract method 0x7b75d7f4.
//
// Solidity: function configureEModeCategory(uint8 id, (uint16,uint16,uint16,string) category) returns()
func (_AavePool *AavePoolTransactorSession) ConfigureEModeCategory(id uint8, category DataTypesEModeCategoryBaseConfiguration) (*types.Transaction, error) {
	return _AavePool.Contract.ConfigureEModeCategory(&_AavePool.TransactOpts, id, category)
}

// ConfigureEModeCategoryBorrowableBitmap is a paid mutator transaction binding the contract method 0xff72158a.
//
// Solidity: function configureEModeCategoryBorrowableBitmap(uint8 id, uint128 borrowableBitmap) returns()
func (_AavePool *AavePoolTransactor) ConfigureEModeCategoryBorrowableBitmap(opts *bind.TransactOpts, id uint8, borrowableBitmap *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "configureEModeCategoryBorrowableBitmap", id, borrowableBitmap)
}

// ConfigureEModeCategoryBorrowableBitmap is a paid mutator transaction binding the contract method 0xff72158a.
//
// Solidity: function configureEModeCategoryBorrowableBitmap(uint8 id, uint128 borrowableBitmap) returns()
func (_AavePool *AavePoolSession) ConfigureEModeCategoryBorrowableBitmap(id uint8, borrowableBitmap *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.ConfigureEModeCategoryBorrowableBitmap(&_AavePool.TransactOpts, id, borrowableBitmap)
}

// ConfigureEModeCategoryBorrowableBitmap is a paid mutator transaction binding the contract method 0xff72158a.
//
// Solidity: function configureEModeCategoryBorrowableBitmap(uint8 id, uint128 borrowableBitmap) returns()
func (_AavePool *AavePoolTransactorSession) ConfigureEModeCategoryBorrowableBitmap(id uint8, borrowableBitmap *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.ConfigureEModeCategoryBorrowableBitmap(&_AavePool.TransactOpts, id, borrowableBitmap)
}

// ConfigureEModeCategoryCollateralBitmap is a paid mutator transaction binding the contract method 0x92380ecb.
//
// Solidity: function configureEModeCategoryCollateralBitmap(uint8 id, uint128 collateralBitmap) returns()
func (_AavePool *AavePoolTransactor) ConfigureEModeCategoryCollateralBitmap(opts *bind.TransactOpts, id uint8, collateralBitmap *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "configureEModeCategoryCollateralBitmap", id, collateralBitmap)
}

// ConfigureEModeCategoryCollateralBitmap is a paid mutator transaction binding the contract method 0x92380ecb.
//
// Solidity: function configureEModeCategoryCollateralBitmap(uint8 id, uint128 collateralBitmap) returns()
func (_AavePool *AavePoolSession) ConfigureEModeCategoryCollateralBitmap(id uint8, collateralBitmap *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.ConfigureEModeCategoryCollateralBitmap(&_AavePool.TransactOpts, id, collateralBitmap)
}

// ConfigureEModeCategoryCollateralBitmap is a paid mutator transaction binding the contract method 0x92380ecb.
//
// Solidity: function configureEModeCategoryCollateralBitmap(uint8 id, uint128 collateralBitmap) returns()
func (_AavePool *AavePoolTransactorSession) ConfigureEModeCategoryCollateralBitmap(id uint8, collateralBitmap *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.ConfigureEModeCategoryCollateralBitmap(&_AavePool.TransactOpts, id, collateralBitmap)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "deposit", asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.Deposit(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactorSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.Deposit(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// DropReserve is a paid mutator transaction binding the contract method 0x63c9b860.
//
// Solidity: function dropReserve(address asset) returns()
func (_AavePool *AavePoolTransactor) DropReserve(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "dropReserve", asset)
}

// DropReserve is a paid mutator transaction binding the contract method 0x63c9b860.
//
// Solidity: function dropReserve(address asset) returns()
func (_AavePool *AavePoolSession) DropReserve(asset common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.DropReserve(&_AavePool.TransactOpts, asset)
}

// DropReserve is a paid mutator transaction binding the contract method 0x63c9b860.
//
// Solidity: function dropReserve(address asset) returns()
func (_AavePool *AavePoolTransactorSession) DropReserve(asset common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.DropReserve(&_AavePool.TransactOpts, asset)
}

// EliminateReserveDeficit is a paid mutator transaction binding the contract method 0xa1d2f3c4.
//
// Solidity: function eliminateReserveDeficit(address asset, uint256 amount) returns(uint256)
func (_AavePool *AavePoolTransactor) EliminateReserveDeficit(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "eliminateReserveDeficit", asset, amount)
}

// EliminateReserveDeficit is a paid mutator transaction binding the contract method 0xa1d2f3c4.
//
// Solidity: function eliminateReserveDeficit(address asset, uint256 amount) returns(uint256)
func (_AavePool *AavePoolSession) EliminateReserveDeficit(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.EliminateReserveDeficit(&_AavePool.TransactOpts, asset, amount)
}

// EliminateReserveDeficit is a paid mutator transaction binding the contract method 0xa1d2f3c4.
//
// Solidity: function eliminateReserveDeficit(address asset, uint256 amount) returns(uint256)
func (_AavePool *AavePoolTransactorSession) EliminateReserveDeficit(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.EliminateReserveDeficit(&_AavePool.TransactOpts, asset, amount)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 scaledAmount, uint256 scaledBalanceFromBefore, uint256 scaledBalanceToBefore) returns()
func (_AavePool *AavePoolTransactor) FinalizeTransfer(opts *bind.TransactOpts, asset common.Address, from common.Address, to common.Address, scaledAmount *big.Int, scaledBalanceFromBefore *big.Int, scaledBalanceToBefore *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "finalizeTransfer", asset, from, to, scaledAmount, scaledBalanceFromBefore, scaledBalanceToBefore)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 scaledAmount, uint256 scaledBalanceFromBefore, uint256 scaledBalanceToBefore) returns()
func (_AavePool *AavePoolSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, scaledAmount *big.Int, scaledBalanceFromBefore *big.Int, scaledBalanceToBefore *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.FinalizeTransfer(&_AavePool.TransactOpts, asset, from, to, scaledAmount, scaledBalanceFromBefore, scaledBalanceToBefore)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 scaledAmount, uint256 scaledBalanceFromBefore, uint256 scaledBalanceToBefore) returns()
func (_AavePool *AavePoolTransactorSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, scaledAmount *big.Int, scaledBalanceFromBefore *big.Int, scaledBalanceToBefore *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.FinalizeTransfer(&_AavePool.TransactOpts, asset, from, to, scaledAmount, scaledBalanceFromBefore, scaledBalanceToBefore)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] interestRateModes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactor) FlashLoan(opts *bind.TransactOpts, receiverAddress common.Address, assets []common.Address, amounts []*big.Int, interestRateModes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "flashLoan", receiverAddress, assets, amounts, interestRateModes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] interestRateModes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AavePool *AavePoolSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, interestRateModes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.FlashLoan(&_AavePool.TransactOpts, receiverAddress, assets, amounts, interestRateModes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] interestRateModes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactorSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, interestRateModes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.FlashLoan(&_AavePool.TransactOpts, receiverAddress, assets, amounts, interestRateModes, onBehalfOf, params, referralCode)
}

// FlashLoanSimple is a paid mutator transaction binding the contract method 0x42b0b77c.
//
// Solidity: function flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes params, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactor) FlashLoanSimple(opts *bind.TransactOpts, receiverAddress common.Address, asset common.Address, amount *big.Int, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "flashLoanSimple", receiverAddress, asset, amount, params, referralCode)
}

// FlashLoanSimple is a paid mutator transaction binding the contract method 0x42b0b77c.
//
// Solidity: function flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes params, uint16 referralCode) returns()
func (_AavePool *AavePoolSession) FlashLoanSimple(receiverAddress common.Address, asset common.Address, amount *big.Int, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.FlashLoanSimple(&_AavePool.TransactOpts, receiverAddress, asset, amount, params, referralCode)
}

// FlashLoanSimple is a paid mutator transaction binding the contract method 0x42b0b77c.
//
// Solidity: function flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes params, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactorSession) FlashLoanSimple(receiverAddress common.Address, asset common.Address, amount *big.Int, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.FlashLoanSimple(&_AavePool.TransactOpts, receiverAddress, asset, amount, params, referralCode)
}

// InitReserve is a paid mutator transaction binding the contract method 0x932f12c8.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address variableDebtAddress) returns()
func (_AavePool *AavePoolTransactor) InitReserve(opts *bind.TransactOpts, asset common.Address, aTokenAddress common.Address, variableDebtAddress common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "initReserve", asset, aTokenAddress, variableDebtAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x932f12c8.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address variableDebtAddress) returns()
func (_AavePool *AavePoolSession) InitReserve(asset common.Address, aTokenAddress common.Address, variableDebtAddress common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.InitReserve(&_AavePool.TransactOpts, asset, aTokenAddress, variableDebtAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x932f12c8.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address variableDebtAddress) returns()
func (_AavePool *AavePoolTransactorSession) InitReserve(asset common.Address, aTokenAddress common.Address, variableDebtAddress common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.InitReserve(&_AavePool.TransactOpts, asset, aTokenAddress, variableDebtAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_AavePool *AavePoolTransactor) Initialize(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "initialize", provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_AavePool *AavePoolSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Initialize(&_AavePool.TransactOpts, provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_AavePool *AavePoolTransactorSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Initialize(&_AavePool.TransactOpts, provider)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address borrower, uint256 debtToCover, bool receiveAToken) returns()
func (_AavePool *AavePoolTransactor) LiquidationCall(opts *bind.TransactOpts, collateralAsset common.Address, debtAsset common.Address, borrower common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "liquidationCall", collateralAsset, debtAsset, borrower, debtToCover, receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address borrower, uint256 debtToCover, bool receiveAToken) returns()
func (_AavePool *AavePoolSession) LiquidationCall(collateralAsset common.Address, debtAsset common.Address, borrower common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _AavePool.Contract.LiquidationCall(&_AavePool.TransactOpts, collateralAsset, debtAsset, borrower, debtToCover, receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address borrower, uint256 debtToCover, bool receiveAToken) returns()
func (_AavePool *AavePoolTransactorSession) LiquidationCall(collateralAsset common.Address, debtAsset common.Address, borrower common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _AavePool.Contract.LiquidationCall(&_AavePool.TransactOpts, collateralAsset, debtAsset, borrower, debtToCover, receiveAToken)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x9cd19996.
//
// Solidity: function mintToTreasury(address[] assets) returns()
func (_AavePool *AavePoolTransactor) MintToTreasury(opts *bind.TransactOpts, assets []common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "mintToTreasury", assets)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x9cd19996.
//
// Solidity: function mintToTreasury(address[] assets) returns()
func (_AavePool *AavePoolSession) MintToTreasury(assets []common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.MintToTreasury(&_AavePool.TransactOpts, assets)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x9cd19996.
//
// Solidity: function mintToTreasury(address[] assets) returns()
func (_AavePool *AavePoolTransactorSession) MintToTreasury(assets []common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.MintToTreasury(&_AavePool.TransactOpts, assets)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AavePool *AavePoolTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AavePool *AavePoolSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _AavePool.Contract.Multicall(&_AavePool.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AavePool *AavePoolTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _AavePool.Contract.Multicall(&_AavePool.TransactOpts, data)
}

// RenouncePositionManagerRole is a paid mutator transaction binding the contract method 0xfea149a6.
//
// Solidity: function renouncePositionManagerRole(address user) returns()
func (_AavePool *AavePoolTransactor) RenouncePositionManagerRole(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "renouncePositionManagerRole", user)
}

// RenouncePositionManagerRole is a paid mutator transaction binding the contract method 0xfea149a6.
//
// Solidity: function renouncePositionManagerRole(address user) returns()
func (_AavePool *AavePoolSession) RenouncePositionManagerRole(user common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.RenouncePositionManagerRole(&_AavePool.TransactOpts, user)
}

// RenouncePositionManagerRole is a paid mutator transaction binding the contract method 0xfea149a6.
//
// Solidity: function renouncePositionManagerRole(address user) returns()
func (_AavePool *AavePoolTransactorSession) RenouncePositionManagerRole(user common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.RenouncePositionManagerRole(&_AavePool.TransactOpts, user)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (_AavePool *AavePoolTransactor) Repay(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "repay", asset, amount, interestRateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (_AavePool *AavePoolSession) Repay(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Repay(&_AavePool.TransactOpts, asset, amount, interestRateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (_AavePool *AavePoolTransactorSession) Repay(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Repay(&_AavePool.TransactOpts, asset, amount, interestRateMode, onBehalfOf)
}

// RepayWithATokens is a paid mutator transaction binding the contract method 0x2dad97d4.
//
// Solidity: function repayWithATokens(address asset, uint256 amount, uint256 interestRateMode) returns(uint256)
func (_AavePool *AavePoolTransactor) RepayWithATokens(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "repayWithATokens", asset, amount, interestRateMode)
}

// RepayWithATokens is a paid mutator transaction binding the contract method 0x2dad97d4.
//
// Solidity: function repayWithATokens(address asset, uint256 amount, uint256 interestRateMode) returns(uint256)
func (_AavePool *AavePoolSession) RepayWithATokens(asset common.Address, amount *big.Int, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.RepayWithATokens(&_AavePool.TransactOpts, asset, amount, interestRateMode)
}

// RepayWithATokens is a paid mutator transaction binding the contract method 0x2dad97d4.
//
// Solidity: function repayWithATokens(address asset, uint256 amount, uint256 interestRateMode) returns(uint256)
func (_AavePool *AavePoolTransactorSession) RepayWithATokens(asset common.Address, amount *big.Int, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.RepayWithATokens(&_AavePool.TransactOpts, asset, amount, interestRateMode)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xee3e210b.
//
// Solidity: function repayWithPermit(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns(uint256)
func (_AavePool *AavePoolTransactor) RepayWithPermit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "repayWithPermit", asset, amount, interestRateMode, onBehalfOf, deadline, permitV, permitR, permitS)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xee3e210b.
//
// Solidity: function repayWithPermit(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns(uint256)
func (_AavePool *AavePoolSession) RepayWithPermit(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AavePool.Contract.RepayWithPermit(&_AavePool.TransactOpts, asset, amount, interestRateMode, onBehalfOf, deadline, permitV, permitR, permitS)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xee3e210b.
//
// Solidity: function repayWithPermit(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns(uint256)
func (_AavePool *AavePoolTransactorSession) RepayWithPermit(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AavePool.Contract.RepayWithPermit(&_AavePool.TransactOpts, asset, amount, interestRateMode, onBehalfOf, deadline, permitV, permitR, permitS)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_AavePool *AavePoolTransactor) RescueTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "rescueTokens", token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_AavePool *AavePoolSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.RescueTokens(&_AavePool.TransactOpts, token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_AavePool *AavePoolTransactorSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.RescueTokens(&_AavePool.TransactOpts, token, to, amount)
}

// ResetIsolationModeTotalDebt is a paid mutator transaction binding the contract method 0xe43e88a1.
//
// Solidity: function resetIsolationModeTotalDebt(address asset) returns()
func (_AavePool *AavePoolTransactor) ResetIsolationModeTotalDebt(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "resetIsolationModeTotalDebt", asset)
}

// ResetIsolationModeTotalDebt is a paid mutator transaction binding the contract method 0xe43e88a1.
//
// Solidity: function resetIsolationModeTotalDebt(address asset) returns()
func (_AavePool *AavePoolSession) ResetIsolationModeTotalDebt(asset common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.ResetIsolationModeTotalDebt(&_AavePool.TransactOpts, asset)
}

// ResetIsolationModeTotalDebt is a paid mutator transaction binding the contract method 0xe43e88a1.
//
// Solidity: function resetIsolationModeTotalDebt(address asset) returns()
func (_AavePool *AavePoolTransactorSession) ResetIsolationModeTotalDebt(asset common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.ResetIsolationModeTotalDebt(&_AavePool.TransactOpts, asset)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xf51e435b.
//
// Solidity: function setConfiguration(address asset, (uint256) configuration) returns()
func (_AavePool *AavePoolTransactor) SetConfiguration(opts *bind.TransactOpts, asset common.Address, configuration DataTypesReserveConfigurationMap) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "setConfiguration", asset, configuration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xf51e435b.
//
// Solidity: function setConfiguration(address asset, (uint256) configuration) returns()
func (_AavePool *AavePoolSession) SetConfiguration(asset common.Address, configuration DataTypesReserveConfigurationMap) (*types.Transaction, error) {
	return _AavePool.Contract.SetConfiguration(&_AavePool.TransactOpts, asset, configuration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xf51e435b.
//
// Solidity: function setConfiguration(address asset, (uint256) configuration) returns()
func (_AavePool *AavePoolTransactorSession) SetConfiguration(asset common.Address, configuration DataTypesReserveConfigurationMap) (*types.Transaction, error) {
	return _AavePool.Contract.SetConfiguration(&_AavePool.TransactOpts, asset, configuration)
}

// SetLiquidationGracePeriod is a paid mutator transaction binding the contract method 0xb1a99e26.
//
// Solidity: function setLiquidationGracePeriod(address asset, uint40 until) returns()
func (_AavePool *AavePoolTransactor) SetLiquidationGracePeriod(opts *bind.TransactOpts, asset common.Address, until *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "setLiquidationGracePeriod", asset, until)
}

// SetLiquidationGracePeriod is a paid mutator transaction binding the contract method 0xb1a99e26.
//
// Solidity: function setLiquidationGracePeriod(address asset, uint40 until) returns()
func (_AavePool *AavePoolSession) SetLiquidationGracePeriod(asset common.Address, until *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.SetLiquidationGracePeriod(&_AavePool.TransactOpts, asset, until)
}

// SetLiquidationGracePeriod is a paid mutator transaction binding the contract method 0xb1a99e26.
//
// Solidity: function setLiquidationGracePeriod(address asset, uint40 until) returns()
func (_AavePool *AavePoolTransactorSession) SetLiquidationGracePeriod(asset common.Address, until *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.SetLiquidationGracePeriod(&_AavePool.TransactOpts, asset, until)
}

// SetUserEMode is a paid mutator transaction binding the contract method 0x28530a47.
//
// Solidity: function setUserEMode(uint8 categoryId) returns()
func (_AavePool *AavePoolTransactor) SetUserEMode(opts *bind.TransactOpts, categoryId uint8) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "setUserEMode", categoryId)
}

// SetUserEMode is a paid mutator transaction binding the contract method 0x28530a47.
//
// Solidity: function setUserEMode(uint8 categoryId) returns()
func (_AavePool *AavePoolSession) SetUserEMode(categoryId uint8) (*types.Transaction, error) {
	return _AavePool.Contract.SetUserEMode(&_AavePool.TransactOpts, categoryId)
}

// SetUserEMode is a paid mutator transaction binding the contract method 0x28530a47.
//
// Solidity: function setUserEMode(uint8 categoryId) returns()
func (_AavePool *AavePoolTransactorSession) SetUserEMode(categoryId uint8) (*types.Transaction, error) {
	return _AavePool.Contract.SetUserEMode(&_AavePool.TransactOpts, categoryId)
}

// SetUserEModeOnBehalfOf is a paid mutator transaction binding the contract method 0x4ba06814.
//
// Solidity: function setUserEModeOnBehalfOf(uint8 categoryId, address onBehalfOf) returns()
func (_AavePool *AavePoolTransactor) SetUserEModeOnBehalfOf(opts *bind.TransactOpts, categoryId uint8, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "setUserEModeOnBehalfOf", categoryId, onBehalfOf)
}

// SetUserEModeOnBehalfOf is a paid mutator transaction binding the contract method 0x4ba06814.
//
// Solidity: function setUserEModeOnBehalfOf(uint8 categoryId, address onBehalfOf) returns()
func (_AavePool *AavePoolSession) SetUserEModeOnBehalfOf(categoryId uint8, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.SetUserEModeOnBehalfOf(&_AavePool.TransactOpts, categoryId, onBehalfOf)
}

// SetUserEModeOnBehalfOf is a paid mutator transaction binding the contract method 0x4ba06814.
//
// Solidity: function setUserEModeOnBehalfOf(uint8 categoryId, address onBehalfOf) returns()
func (_AavePool *AavePoolTransactorSession) SetUserEModeOnBehalfOf(categoryId uint8, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.SetUserEModeOnBehalfOf(&_AavePool.TransactOpts, categoryId, onBehalfOf)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_AavePool *AavePoolTransactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "setUserUseReserveAsCollateral", asset, useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_AavePool *AavePoolSession) SetUserUseReserveAsCollateral(asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _AavePool.Contract.SetUserUseReserveAsCollateral(&_AavePool.TransactOpts, asset, useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_AavePool *AavePoolTransactorSession) SetUserUseReserveAsCollateral(asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _AavePool.Contract.SetUserUseReserveAsCollateral(&_AavePool.TransactOpts, asset, useAsCollateral)
}

// SetUserUseReserveAsCollateralOnBehalfOf is a paid mutator transaction binding the contract method 0x972b35fa.
//
// Solidity: function setUserUseReserveAsCollateralOnBehalfOf(address asset, bool useAsCollateral, address onBehalfOf) returns()
func (_AavePool *AavePoolTransactor) SetUserUseReserveAsCollateralOnBehalfOf(opts *bind.TransactOpts, asset common.Address, useAsCollateral bool, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "setUserUseReserveAsCollateralOnBehalfOf", asset, useAsCollateral, onBehalfOf)
}

// SetUserUseReserveAsCollateralOnBehalfOf is a paid mutator transaction binding the contract method 0x972b35fa.
//
// Solidity: function setUserUseReserveAsCollateralOnBehalfOf(address asset, bool useAsCollateral, address onBehalfOf) returns()
func (_AavePool *AavePoolSession) SetUserUseReserveAsCollateralOnBehalfOf(asset common.Address, useAsCollateral bool, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.SetUserUseReserveAsCollateralOnBehalfOf(&_AavePool.TransactOpts, asset, useAsCollateral, onBehalfOf)
}

// SetUserUseReserveAsCollateralOnBehalfOf is a paid mutator transaction binding the contract method 0x972b35fa.
//
// Solidity: function setUserUseReserveAsCollateralOnBehalfOf(address asset, bool useAsCollateral, address onBehalfOf) returns()
func (_AavePool *AavePoolTransactorSession) SetUserUseReserveAsCollateralOnBehalfOf(asset common.Address, useAsCollateral bool, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.SetUserUseReserveAsCollateralOnBehalfOf(&_AavePool.TransactOpts, asset, useAsCollateral, onBehalfOf)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactor) Supply(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "supply", asset, amount, onBehalfOf, referralCode)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolSession) Supply(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.Supply(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactorSession) Supply(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.Supply(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// SupplyWithPermit is a paid mutator transaction binding the contract method 0x02c205f0.
//
// Solidity: function supplyWithPermit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_AavePool *AavePoolTransactor) SupplyWithPermit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "supplyWithPermit", asset, amount, onBehalfOf, referralCode, deadline, permitV, permitR, permitS)
}

// SupplyWithPermit is a paid mutator transaction binding the contract method 0x02c205f0.
//
// Solidity: function supplyWithPermit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_AavePool *AavePoolSession) SupplyWithPermit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AavePool.Contract.SupplyWithPermit(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode, deadline, permitV, permitR, permitS)
}

// SupplyWithPermit is a paid mutator transaction binding the contract method 0x02c205f0.
//
// Solidity: function supplyWithPermit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_AavePool *AavePoolTransactorSession) SupplyWithPermit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AavePool.Contract.SupplyWithPermit(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode, deadline, permitV, permitR, permitS)
}

// SyncIndexesState is a paid mutator transaction binding the contract method 0xab2b51f6.
//
// Solidity: function syncIndexesState(address asset) returns()
func (_AavePool *AavePoolTransactor) SyncIndexesState(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "syncIndexesState", asset)
}

// SyncIndexesState is a paid mutator transaction binding the contract method 0xab2b51f6.
//
// Solidity: function syncIndexesState(address asset) returns()
func (_AavePool *AavePoolSession) SyncIndexesState(asset common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.SyncIndexesState(&_AavePool.TransactOpts, asset)
}

// SyncIndexesState is a paid mutator transaction binding the contract method 0xab2b51f6.
//
// Solidity: function syncIndexesState(address asset) returns()
func (_AavePool *AavePoolTransactorSession) SyncIndexesState(asset common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.SyncIndexesState(&_AavePool.TransactOpts, asset)
}

// SyncRatesState is a paid mutator transaction binding the contract method 0x98c7da4e.
//
// Solidity: function syncRatesState(address asset) returns()
func (_AavePool *AavePoolTransactor) SyncRatesState(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "syncRatesState", asset)
}

// SyncRatesState is a paid mutator transaction binding the contract method 0x98c7da4e.
//
// Solidity: function syncRatesState(address asset) returns()
func (_AavePool *AavePoolSession) SyncRatesState(asset common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.SyncRatesState(&_AavePool.TransactOpts, asset)
}

// SyncRatesState is a paid mutator transaction binding the contract method 0x98c7da4e.
//
// Solidity: function syncRatesState(address asset) returns()
func (_AavePool *AavePoolTransactorSession) SyncRatesState(asset common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.SyncRatesState(&_AavePool.TransactOpts, asset)
}

// UpdateFlashloanPremium is a paid mutator transaction binding the contract method 0x9c1d5f00.
//
// Solidity: function updateFlashloanPremium(uint128 flashLoanPremium) returns()
func (_AavePool *AavePoolTransactor) UpdateFlashloanPremium(opts *bind.TransactOpts, flashLoanPremium *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "updateFlashloanPremium", flashLoanPremium)
}

// UpdateFlashloanPremium is a paid mutator transaction binding the contract method 0x9c1d5f00.
//
// Solidity: function updateFlashloanPremium(uint128 flashLoanPremium) returns()
func (_AavePool *AavePoolSession) UpdateFlashloanPremium(flashLoanPremium *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.UpdateFlashloanPremium(&_AavePool.TransactOpts, flashLoanPremium)
}

// UpdateFlashloanPremium is a paid mutator transaction binding the contract method 0x9c1d5f00.
//
// Solidity: function updateFlashloanPremium(uint128 flashLoanPremium) returns()
func (_AavePool *AavePoolTransactorSession) UpdateFlashloanPremium(flashLoanPremium *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.UpdateFlashloanPremium(&_AavePool.TransactOpts, flashLoanPremium)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_AavePool *AavePoolTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "withdraw", asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_AavePool *AavePoolSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Withdraw(&_AavePool.TransactOpts, asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_AavePool *AavePoolTransactorSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Withdraw(&_AavePool.TransactOpts, asset, amount, to)
}

// AavePoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the AavePool contract.
type AavePoolBorrowIterator struct {
	Event *AavePoolBorrow // Event containing the contract specifics and raw log

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
func (it *AavePoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolBorrow)
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
		it.Event = new(AavePoolBorrow)
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
func (it *AavePoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolBorrow represents a Borrow event raised by the AavePool contract.
type AavePoolBorrow struct {
	Reserve          common.Address
	User             common.Address
	OnBehalfOf       common.Address
	Amount           *big.Int
	InterestRateMode uint8
	BorrowRate       *big.Int
	ReferralCode     uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint8 interestRateMode, uint256 borrowRate, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) FilterBorrow(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (*AavePoolBorrowIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolBorrowIterator{contract: _AavePool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint8 interestRateMode, uint256 borrowRate, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *AavePoolBorrow, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolBorrow)
				if err := _AavePool.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint8 interestRateMode, uint256 borrowRate, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) ParseBorrow(log types.Log) (*AavePoolBorrow, error) {
	event := new(AavePoolBorrow)
	if err := _AavePool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolDeficitCoveredIterator is returned from FilterDeficitCovered and is used to iterate over the raw logs and unpacked data for DeficitCovered events raised by the AavePool contract.
type AavePoolDeficitCoveredIterator struct {
	Event *AavePoolDeficitCovered // Event containing the contract specifics and raw log

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
func (it *AavePoolDeficitCoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolDeficitCovered)
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
		it.Event = new(AavePoolDeficitCovered)
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
func (it *AavePoolDeficitCoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolDeficitCoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolDeficitCovered represents a DeficitCovered event raised by the AavePool contract.
type AavePoolDeficitCovered struct {
	Reserve       common.Address
	Caller        common.Address
	AmountCovered *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeficitCovered is a free log retrieval operation binding the contract event 0x84b203e49f1a4b553088061534231969a68ad1c81be192205e96d23a206cb26a.
//
// Solidity: event DeficitCovered(address indexed reserve, address caller, uint256 amountCovered)
func (_AavePool *AavePoolFilterer) FilterDeficitCovered(opts *bind.FilterOpts, reserve []common.Address) (*AavePoolDeficitCoveredIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "DeficitCovered", reserveRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolDeficitCoveredIterator{contract: _AavePool.contract, event: "DeficitCovered", logs: logs, sub: sub}, nil
}

// WatchDeficitCovered is a free log subscription operation binding the contract event 0x84b203e49f1a4b553088061534231969a68ad1c81be192205e96d23a206cb26a.
//
// Solidity: event DeficitCovered(address indexed reserve, address caller, uint256 amountCovered)
func (_AavePool *AavePoolFilterer) WatchDeficitCovered(opts *bind.WatchOpts, sink chan<- *AavePoolDeficitCovered, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "DeficitCovered", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolDeficitCovered)
				if err := _AavePool.contract.UnpackLog(event, "DeficitCovered", log); err != nil {
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

// ParseDeficitCovered is a log parse operation binding the contract event 0x84b203e49f1a4b553088061534231969a68ad1c81be192205e96d23a206cb26a.
//
// Solidity: event DeficitCovered(address indexed reserve, address caller, uint256 amountCovered)
func (_AavePool *AavePoolFilterer) ParseDeficitCovered(log types.Log) (*AavePoolDeficitCovered, error) {
	event := new(AavePoolDeficitCovered)
	if err := _AavePool.contract.UnpackLog(event, "DeficitCovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolDeficitCreatedIterator is returned from FilterDeficitCreated and is used to iterate over the raw logs and unpacked data for DeficitCreated events raised by the AavePool contract.
type AavePoolDeficitCreatedIterator struct {
	Event *AavePoolDeficitCreated // Event containing the contract specifics and raw log

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
func (it *AavePoolDeficitCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolDeficitCreated)
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
		it.Event = new(AavePoolDeficitCreated)
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
func (it *AavePoolDeficitCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolDeficitCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolDeficitCreated represents a DeficitCreated event raised by the AavePool contract.
type AavePoolDeficitCreated struct {
	User          common.Address
	DebtAsset     common.Address
	AmountCreated *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeficitCreated is a free log retrieval operation binding the contract event 0x2bccfb3fad376d59d7accf970515eb77b2f27b082c90ed0fb15583dd5a942699.
//
// Solidity: event DeficitCreated(address indexed user, address indexed debtAsset, uint256 amountCreated)
func (_AavePool *AavePoolFilterer) FilterDeficitCreated(opts *bind.FilterOpts, user []common.Address, debtAsset []common.Address) (*AavePoolDeficitCreatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "DeficitCreated", userRule, debtAssetRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolDeficitCreatedIterator{contract: _AavePool.contract, event: "DeficitCreated", logs: logs, sub: sub}, nil
}

// WatchDeficitCreated is a free log subscription operation binding the contract event 0x2bccfb3fad376d59d7accf970515eb77b2f27b082c90ed0fb15583dd5a942699.
//
// Solidity: event DeficitCreated(address indexed user, address indexed debtAsset, uint256 amountCreated)
func (_AavePool *AavePoolFilterer) WatchDeficitCreated(opts *bind.WatchOpts, sink chan<- *AavePoolDeficitCreated, user []common.Address, debtAsset []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "DeficitCreated", userRule, debtAssetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolDeficitCreated)
				if err := _AavePool.contract.UnpackLog(event, "DeficitCreated", log); err != nil {
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

// ParseDeficitCreated is a log parse operation binding the contract event 0x2bccfb3fad376d59d7accf970515eb77b2f27b082c90ed0fb15583dd5a942699.
//
// Solidity: event DeficitCreated(address indexed user, address indexed debtAsset, uint256 amountCreated)
func (_AavePool *AavePoolFilterer) ParseDeficitCreated(log types.Log) (*AavePoolDeficitCreated, error) {
	event := new(AavePoolDeficitCreated)
	if err := _AavePool.contract.UnpackLog(event, "DeficitCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the AavePool contract.
type AavePoolFlashLoanIterator struct {
	Event *AavePoolFlashLoan // Event containing the contract specifics and raw log

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
func (it *AavePoolFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolFlashLoan)
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
		it.Event = new(AavePoolFlashLoan)
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
func (it *AavePoolFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolFlashLoan represents a FlashLoan event raised by the AavePool contract.
type AavePoolFlashLoan struct {
	Target           common.Address
	Initiator        common.Address
	Asset            common.Address
	Amount           *big.Int
	InterestRateMode uint8
	Premium          *big.Int
	ReferralCode     uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0.
//
// Solidity: event FlashLoan(address indexed target, address initiator, address indexed asset, uint256 amount, uint8 interestRateMode, uint256 premium, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) FilterFlashLoan(opts *bind.FilterOpts, target []common.Address, asset []common.Address, referralCode []uint16) (*AavePoolFlashLoanIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "FlashLoan", targetRule, assetRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolFlashLoanIterator{contract: _AavePool.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0.
//
// Solidity: event FlashLoan(address indexed target, address initiator, address indexed asset, uint256 amount, uint8 interestRateMode, uint256 premium, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *AavePoolFlashLoan, target []common.Address, asset []common.Address, referralCode []uint16) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "FlashLoan", targetRule, assetRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolFlashLoan)
				if err := _AavePool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0.
//
// Solidity: event FlashLoan(address indexed target, address initiator, address indexed asset, uint256 amount, uint8 interestRateMode, uint256 premium, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) ParseFlashLoan(log types.Log) (*AavePoolFlashLoan, error) {
	event := new(AavePoolFlashLoan)
	if err := _AavePool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolIsolationModeTotalDebtUpdatedIterator is returned from FilterIsolationModeTotalDebtUpdated and is used to iterate over the raw logs and unpacked data for IsolationModeTotalDebtUpdated events raised by the AavePool contract.
type AavePoolIsolationModeTotalDebtUpdatedIterator struct {
	Event *AavePoolIsolationModeTotalDebtUpdated // Event containing the contract specifics and raw log

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
func (it *AavePoolIsolationModeTotalDebtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolIsolationModeTotalDebtUpdated)
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
		it.Event = new(AavePoolIsolationModeTotalDebtUpdated)
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
func (it *AavePoolIsolationModeTotalDebtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolIsolationModeTotalDebtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolIsolationModeTotalDebtUpdated represents a IsolationModeTotalDebtUpdated event raised by the AavePool contract.
type AavePoolIsolationModeTotalDebtUpdated struct {
	Asset     common.Address
	TotalDebt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIsolationModeTotalDebtUpdated is a free log retrieval operation binding the contract event 0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5.
//
// Solidity: event IsolationModeTotalDebtUpdated(address indexed asset, uint256 totalDebt)
func (_AavePool *AavePoolFilterer) FilterIsolationModeTotalDebtUpdated(opts *bind.FilterOpts, asset []common.Address) (*AavePoolIsolationModeTotalDebtUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "IsolationModeTotalDebtUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolIsolationModeTotalDebtUpdatedIterator{contract: _AavePool.contract, event: "IsolationModeTotalDebtUpdated", logs: logs, sub: sub}, nil
}

// WatchIsolationModeTotalDebtUpdated is a free log subscription operation binding the contract event 0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5.
//
// Solidity: event IsolationModeTotalDebtUpdated(address indexed asset, uint256 totalDebt)
func (_AavePool *AavePoolFilterer) WatchIsolationModeTotalDebtUpdated(opts *bind.WatchOpts, sink chan<- *AavePoolIsolationModeTotalDebtUpdated, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "IsolationModeTotalDebtUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolIsolationModeTotalDebtUpdated)
				if err := _AavePool.contract.UnpackLog(event, "IsolationModeTotalDebtUpdated", log); err != nil {
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

// ParseIsolationModeTotalDebtUpdated is a log parse operation binding the contract event 0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5.
//
// Solidity: event IsolationModeTotalDebtUpdated(address indexed asset, uint256 totalDebt)
func (_AavePool *AavePoolFilterer) ParseIsolationModeTotalDebtUpdated(log types.Log) (*AavePoolIsolationModeTotalDebtUpdated, error) {
	event := new(AavePoolIsolationModeTotalDebtUpdated)
	if err := _AavePool.contract.UnpackLog(event, "IsolationModeTotalDebtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolLiquidationCallIterator is returned from FilterLiquidationCall and is used to iterate over the raw logs and unpacked data for LiquidationCall events raised by the AavePool contract.
type AavePoolLiquidationCallIterator struct {
	Event *AavePoolLiquidationCall // Event containing the contract specifics and raw log

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
func (it *AavePoolLiquidationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolLiquidationCall)
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
		it.Event = new(AavePoolLiquidationCall)
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
func (it *AavePoolLiquidationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolLiquidationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolLiquidationCall represents a LiquidationCall event raised by the AavePool contract.
type AavePoolLiquidationCall struct {
	CollateralAsset            common.Address
	DebtAsset                  common.Address
	User                       common.Address
	DebtToCover                *big.Int
	LiquidatedCollateralAmount *big.Int
	Liquidator                 common.Address
	ReceiveAToken              bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterLiquidationCall is a free log retrieval operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_AavePool *AavePoolFilterer) FilterLiquidationCall(opts *bind.FilterOpts, collateralAsset []common.Address, debtAsset []common.Address, user []common.Address) (*AavePoolLiquidationCallIterator, error) {

	var collateralAssetRule []interface{}
	for _, collateralAssetItem := range collateralAsset {
		collateralAssetRule = append(collateralAssetRule, collateralAssetItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "LiquidationCall", collateralAssetRule, debtAssetRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolLiquidationCallIterator{contract: _AavePool.contract, event: "LiquidationCall", logs: logs, sub: sub}, nil
}

// WatchLiquidationCall is a free log subscription operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_AavePool *AavePoolFilterer) WatchLiquidationCall(opts *bind.WatchOpts, sink chan<- *AavePoolLiquidationCall, collateralAsset []common.Address, debtAsset []common.Address, user []common.Address) (event.Subscription, error) {

	var collateralAssetRule []interface{}
	for _, collateralAssetItem := range collateralAsset {
		collateralAssetRule = append(collateralAssetRule, collateralAssetItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "LiquidationCall", collateralAssetRule, debtAssetRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolLiquidationCall)
				if err := _AavePool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
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

// ParseLiquidationCall is a log parse operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_AavePool *AavePoolFilterer) ParseLiquidationCall(log types.Log) (*AavePoolLiquidationCall, error) {
	event := new(AavePoolLiquidationCall)
	if err := _AavePool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolMintedToTreasuryIterator is returned from FilterMintedToTreasury and is used to iterate over the raw logs and unpacked data for MintedToTreasury events raised by the AavePool contract.
type AavePoolMintedToTreasuryIterator struct {
	Event *AavePoolMintedToTreasury // Event containing the contract specifics and raw log

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
func (it *AavePoolMintedToTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolMintedToTreasury)
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
		it.Event = new(AavePoolMintedToTreasury)
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
func (it *AavePoolMintedToTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolMintedToTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolMintedToTreasury represents a MintedToTreasury event raised by the AavePool contract.
type AavePoolMintedToTreasury struct {
	Reserve      common.Address
	AmountMinted *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintedToTreasury is a free log retrieval operation binding the contract event 0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de.
//
// Solidity: event MintedToTreasury(address indexed reserve, uint256 amountMinted)
func (_AavePool *AavePoolFilterer) FilterMintedToTreasury(opts *bind.FilterOpts, reserve []common.Address) (*AavePoolMintedToTreasuryIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "MintedToTreasury", reserveRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolMintedToTreasuryIterator{contract: _AavePool.contract, event: "MintedToTreasury", logs: logs, sub: sub}, nil
}

// WatchMintedToTreasury is a free log subscription operation binding the contract event 0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de.
//
// Solidity: event MintedToTreasury(address indexed reserve, uint256 amountMinted)
func (_AavePool *AavePoolFilterer) WatchMintedToTreasury(opts *bind.WatchOpts, sink chan<- *AavePoolMintedToTreasury, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "MintedToTreasury", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolMintedToTreasury)
				if err := _AavePool.contract.UnpackLog(event, "MintedToTreasury", log); err != nil {
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

// ParseMintedToTreasury is a log parse operation binding the contract event 0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de.
//
// Solidity: event MintedToTreasury(address indexed reserve, uint256 amountMinted)
func (_AavePool *AavePoolFilterer) ParseMintedToTreasury(log types.Log) (*AavePoolMintedToTreasury, error) {
	event := new(AavePoolMintedToTreasury)
	if err := _AavePool.contract.UnpackLog(event, "MintedToTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolPositionManagerApprovedIterator is returned from FilterPositionManagerApproved and is used to iterate over the raw logs and unpacked data for PositionManagerApproved events raised by the AavePool contract.
type AavePoolPositionManagerApprovedIterator struct {
	Event *AavePoolPositionManagerApproved // Event containing the contract specifics and raw log

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
func (it *AavePoolPositionManagerApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolPositionManagerApproved)
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
		it.Event = new(AavePoolPositionManagerApproved)
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
func (it *AavePoolPositionManagerApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolPositionManagerApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolPositionManagerApproved represents a PositionManagerApproved event raised by the AavePool contract.
type AavePoolPositionManagerApproved struct {
	User            common.Address
	PositionManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPositionManagerApproved is a free log retrieval operation binding the contract event 0x540e692f36c2fa13e7583c4deeffd91ce6bc04f91e7d84f295d9d858372875fc.
//
// Solidity: event PositionManagerApproved(address indexed user, address indexed positionManager)
func (_AavePool *AavePoolFilterer) FilterPositionManagerApproved(opts *bind.FilterOpts, user []common.Address, positionManager []common.Address) (*AavePoolPositionManagerApprovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var positionManagerRule []interface{}
	for _, positionManagerItem := range positionManager {
		positionManagerRule = append(positionManagerRule, positionManagerItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "PositionManagerApproved", userRule, positionManagerRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolPositionManagerApprovedIterator{contract: _AavePool.contract, event: "PositionManagerApproved", logs: logs, sub: sub}, nil
}

// WatchPositionManagerApproved is a free log subscription operation binding the contract event 0x540e692f36c2fa13e7583c4deeffd91ce6bc04f91e7d84f295d9d858372875fc.
//
// Solidity: event PositionManagerApproved(address indexed user, address indexed positionManager)
func (_AavePool *AavePoolFilterer) WatchPositionManagerApproved(opts *bind.WatchOpts, sink chan<- *AavePoolPositionManagerApproved, user []common.Address, positionManager []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var positionManagerRule []interface{}
	for _, positionManagerItem := range positionManager {
		positionManagerRule = append(positionManagerRule, positionManagerItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "PositionManagerApproved", userRule, positionManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolPositionManagerApproved)
				if err := _AavePool.contract.UnpackLog(event, "PositionManagerApproved", log); err != nil {
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

// ParsePositionManagerApproved is a log parse operation binding the contract event 0x540e692f36c2fa13e7583c4deeffd91ce6bc04f91e7d84f295d9d858372875fc.
//
// Solidity: event PositionManagerApproved(address indexed user, address indexed positionManager)
func (_AavePool *AavePoolFilterer) ParsePositionManagerApproved(log types.Log) (*AavePoolPositionManagerApproved, error) {
	event := new(AavePoolPositionManagerApproved)
	if err := _AavePool.contract.UnpackLog(event, "PositionManagerApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolPositionManagerRevokedIterator is returned from FilterPositionManagerRevoked and is used to iterate over the raw logs and unpacked data for PositionManagerRevoked events raised by the AavePool contract.
type AavePoolPositionManagerRevokedIterator struct {
	Event *AavePoolPositionManagerRevoked // Event containing the contract specifics and raw log

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
func (it *AavePoolPositionManagerRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolPositionManagerRevoked)
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
		it.Event = new(AavePoolPositionManagerRevoked)
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
func (it *AavePoolPositionManagerRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolPositionManagerRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolPositionManagerRevoked represents a PositionManagerRevoked event raised by the AavePool contract.
type AavePoolPositionManagerRevoked struct {
	User            common.Address
	PositionManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPositionManagerRevoked is a free log retrieval operation binding the contract event 0x08c92c3870d10c79e9673fecea8f4ff261f8e6b661067d9ca63fd777882bff15.
//
// Solidity: event PositionManagerRevoked(address indexed user, address indexed positionManager)
func (_AavePool *AavePoolFilterer) FilterPositionManagerRevoked(opts *bind.FilterOpts, user []common.Address, positionManager []common.Address) (*AavePoolPositionManagerRevokedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var positionManagerRule []interface{}
	for _, positionManagerItem := range positionManager {
		positionManagerRule = append(positionManagerRule, positionManagerItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "PositionManagerRevoked", userRule, positionManagerRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolPositionManagerRevokedIterator{contract: _AavePool.contract, event: "PositionManagerRevoked", logs: logs, sub: sub}, nil
}

// WatchPositionManagerRevoked is a free log subscription operation binding the contract event 0x08c92c3870d10c79e9673fecea8f4ff261f8e6b661067d9ca63fd777882bff15.
//
// Solidity: event PositionManagerRevoked(address indexed user, address indexed positionManager)
func (_AavePool *AavePoolFilterer) WatchPositionManagerRevoked(opts *bind.WatchOpts, sink chan<- *AavePoolPositionManagerRevoked, user []common.Address, positionManager []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var positionManagerRule []interface{}
	for _, positionManagerItem := range positionManager {
		positionManagerRule = append(positionManagerRule, positionManagerItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "PositionManagerRevoked", userRule, positionManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolPositionManagerRevoked)
				if err := _AavePool.contract.UnpackLog(event, "PositionManagerRevoked", log); err != nil {
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

// ParsePositionManagerRevoked is a log parse operation binding the contract event 0x08c92c3870d10c79e9673fecea8f4ff261f8e6b661067d9ca63fd777882bff15.
//
// Solidity: event PositionManagerRevoked(address indexed user, address indexed positionManager)
func (_AavePool *AavePoolFilterer) ParsePositionManagerRevoked(log types.Log) (*AavePoolPositionManagerRevoked, error) {
	event := new(AavePoolPositionManagerRevoked)
	if err := _AavePool.contract.UnpackLog(event, "PositionManagerRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the AavePool contract.
type AavePoolRepayIterator struct {
	Event *AavePoolRepay // Event containing the contract specifics and raw log

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
func (it *AavePoolRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolRepay)
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
		it.Event = new(AavePoolRepay)
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
func (it *AavePoolRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolRepay represents a Repay event raised by the AavePool contract.
type AavePoolRepay struct {
	Reserve    common.Address
	User       common.Address
	Repayer    common.Address
	Amount     *big.Int
	UseATokens bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount, bool useATokens)
func (_AavePool *AavePoolFilterer) FilterRepay(opts *bind.FilterOpts, reserve []common.Address, user []common.Address, repayer []common.Address) (*AavePoolRepayIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var repayerRule []interface{}
	for _, repayerItem := range repayer {
		repayerRule = append(repayerRule, repayerItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "Repay", reserveRule, userRule, repayerRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolRepayIterator{contract: _AavePool.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount, bool useATokens)
func (_AavePool *AavePoolFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *AavePoolRepay, reserve []common.Address, user []common.Address, repayer []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var repayerRule []interface{}
	for _, repayerItem := range repayer {
		repayerRule = append(repayerRule, repayerItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "Repay", reserveRule, userRule, repayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolRepay)
				if err := _AavePool.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount, bool useATokens)
func (_AavePool *AavePoolFilterer) ParseRepay(log types.Log) (*AavePoolRepay, error) {
	event := new(AavePoolRepay)
	if err := _AavePool.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolReserveDataUpdatedIterator is returned from FilterReserveDataUpdated and is used to iterate over the raw logs and unpacked data for ReserveDataUpdated events raised by the AavePool contract.
type AavePoolReserveDataUpdatedIterator struct {
	Event *AavePoolReserveDataUpdated // Event containing the contract specifics and raw log

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
func (it *AavePoolReserveDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolReserveDataUpdated)
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
		it.Event = new(AavePoolReserveDataUpdated)
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
func (it *AavePoolReserveDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolReserveDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolReserveDataUpdated represents a ReserveDataUpdated event raised by the AavePool contract.
type AavePoolReserveDataUpdated struct {
	Reserve             common.Address
	LiquidityRate       *big.Int
	StableBorrowRate    *big.Int
	VariableBorrowRate  *big.Int
	LiquidityIndex      *big.Int
	VariableBorrowIndex *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReserveDataUpdated is a free log retrieval operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_AavePool *AavePoolFilterer) FilterReserveDataUpdated(opts *bind.FilterOpts, reserve []common.Address) (*AavePoolReserveDataUpdatedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolReserveDataUpdatedIterator{contract: _AavePool.contract, event: "ReserveDataUpdated", logs: logs, sub: sub}, nil
}

// WatchReserveDataUpdated is a free log subscription operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_AavePool *AavePoolFilterer) WatchReserveDataUpdated(opts *bind.WatchOpts, sink chan<- *AavePoolReserveDataUpdated, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolReserveDataUpdated)
				if err := _AavePool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
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

// ParseReserveDataUpdated is a log parse operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_AavePool *AavePoolFilterer) ParseReserveDataUpdated(log types.Log) (*AavePoolReserveDataUpdated, error) {
	event := new(AavePoolReserveDataUpdated)
	if err := _AavePool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolReserveUsedAsCollateralDisabledIterator is returned from FilterReserveUsedAsCollateralDisabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralDisabled events raised by the AavePool contract.
type AavePoolReserveUsedAsCollateralDisabledIterator struct {
	Event *AavePoolReserveUsedAsCollateralDisabled // Event containing the contract specifics and raw log

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
func (it *AavePoolReserveUsedAsCollateralDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolReserveUsedAsCollateralDisabled)
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
		it.Event = new(AavePoolReserveUsedAsCollateralDisabled)
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
func (it *AavePoolReserveUsedAsCollateralDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolReserveUsedAsCollateralDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolReserveUsedAsCollateralDisabled represents a ReserveUsedAsCollateralDisabled event raised by the AavePool contract.
type AavePoolReserveUsedAsCollateralDisabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralDisabled is a free log retrieval operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) FilterReserveUsedAsCollateralDisabled(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AavePoolReserveUsedAsCollateralDisabledIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "ReserveUsedAsCollateralDisabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolReserveUsedAsCollateralDisabledIterator{contract: _AavePool.contract, event: "ReserveUsedAsCollateralDisabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralDisabled is a free log subscription operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) WatchReserveUsedAsCollateralDisabled(opts *bind.WatchOpts, sink chan<- *AavePoolReserveUsedAsCollateralDisabled, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "ReserveUsedAsCollateralDisabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolReserveUsedAsCollateralDisabled)
				if err := _AavePool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
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

// ParseReserveUsedAsCollateralDisabled is a log parse operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) ParseReserveUsedAsCollateralDisabled(log types.Log) (*AavePoolReserveUsedAsCollateralDisabled, error) {
	event := new(AavePoolReserveUsedAsCollateralDisabled)
	if err := _AavePool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolReserveUsedAsCollateralEnabledIterator is returned from FilterReserveUsedAsCollateralEnabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralEnabled events raised by the AavePool contract.
type AavePoolReserveUsedAsCollateralEnabledIterator struct {
	Event *AavePoolReserveUsedAsCollateralEnabled // Event containing the contract specifics and raw log

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
func (it *AavePoolReserveUsedAsCollateralEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolReserveUsedAsCollateralEnabled)
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
		it.Event = new(AavePoolReserveUsedAsCollateralEnabled)
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
func (it *AavePoolReserveUsedAsCollateralEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolReserveUsedAsCollateralEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolReserveUsedAsCollateralEnabled represents a ReserveUsedAsCollateralEnabled event raised by the AavePool contract.
type AavePoolReserveUsedAsCollateralEnabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralEnabled is a free log retrieval operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) FilterReserveUsedAsCollateralEnabled(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AavePoolReserveUsedAsCollateralEnabledIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "ReserveUsedAsCollateralEnabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolReserveUsedAsCollateralEnabledIterator{contract: _AavePool.contract, event: "ReserveUsedAsCollateralEnabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralEnabled is a free log subscription operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) WatchReserveUsedAsCollateralEnabled(opts *bind.WatchOpts, sink chan<- *AavePoolReserveUsedAsCollateralEnabled, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "ReserveUsedAsCollateralEnabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolReserveUsedAsCollateralEnabled)
				if err := _AavePool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
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

// ParseReserveUsedAsCollateralEnabled is a log parse operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) ParseReserveUsedAsCollateralEnabled(log types.Log) (*AavePoolReserveUsedAsCollateralEnabled, error) {
	event := new(AavePoolReserveUsedAsCollateralEnabled)
	if err := _AavePool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the AavePool contract.
type AavePoolSupplyIterator struct {
	Event *AavePoolSupply // Event containing the contract specifics and raw log

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
func (it *AavePoolSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolSupply)
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
		it.Event = new(AavePoolSupply)
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
func (it *AavePoolSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolSupply represents a Supply event raised by the AavePool contract.
type AavePoolSupply struct {
	Reserve      common.Address
	User         common.Address
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61.
//
// Solidity: event Supply(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) FilterSupply(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (*AavePoolSupplyIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "Supply", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolSupplyIterator{contract: _AavePool.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61.
//
// Solidity: event Supply(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *AavePoolSupply, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "Supply", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolSupply)
				if err := _AavePool.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61.
//
// Solidity: event Supply(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) ParseSupply(log types.Log) (*AavePoolSupply, error) {
	event := new(AavePoolSupply)
	if err := _AavePool.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolUserEModeSetIterator is returned from FilterUserEModeSet and is used to iterate over the raw logs and unpacked data for UserEModeSet events raised by the AavePool contract.
type AavePoolUserEModeSetIterator struct {
	Event *AavePoolUserEModeSet // Event containing the contract specifics and raw log

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
func (it *AavePoolUserEModeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolUserEModeSet)
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
		it.Event = new(AavePoolUserEModeSet)
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
func (it *AavePoolUserEModeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolUserEModeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolUserEModeSet represents a UserEModeSet event raised by the AavePool contract.
type AavePoolUserEModeSet struct {
	User       common.Address
	CategoryId uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserEModeSet is a free log retrieval operation binding the contract event 0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84.
//
// Solidity: event UserEModeSet(address indexed user, uint8 categoryId)
func (_AavePool *AavePoolFilterer) FilterUserEModeSet(opts *bind.FilterOpts, user []common.Address) (*AavePoolUserEModeSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "UserEModeSet", userRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolUserEModeSetIterator{contract: _AavePool.contract, event: "UserEModeSet", logs: logs, sub: sub}, nil
}

// WatchUserEModeSet is a free log subscription operation binding the contract event 0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84.
//
// Solidity: event UserEModeSet(address indexed user, uint8 categoryId)
func (_AavePool *AavePoolFilterer) WatchUserEModeSet(opts *bind.WatchOpts, sink chan<- *AavePoolUserEModeSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "UserEModeSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolUserEModeSet)
				if err := _AavePool.contract.UnpackLog(event, "UserEModeSet", log); err != nil {
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

// ParseUserEModeSet is a log parse operation binding the contract event 0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84.
//
// Solidity: event UserEModeSet(address indexed user, uint8 categoryId)
func (_AavePool *AavePoolFilterer) ParseUserEModeSet(log types.Log) (*AavePoolUserEModeSet, error) {
	event := new(AavePoolUserEModeSet)
	if err := _AavePool.contract.UnpackLog(event, "UserEModeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the AavePool contract.
type AavePoolWithdrawIterator struct {
	Event *AavePoolWithdraw // Event containing the contract specifics and raw log

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
func (it *AavePoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolWithdraw)
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
		it.Event = new(AavePoolWithdraw)
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
func (it *AavePoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolWithdraw represents a Withdraw event raised by the AavePool contract.
type AavePoolWithdraw struct {
	Reserve common.Address
	User    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_AavePool *AavePoolFilterer) FilterWithdraw(opts *bind.FilterOpts, reserve []common.Address, user []common.Address, to []common.Address) (*AavePoolWithdrawIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "Withdraw", reserveRule, userRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolWithdrawIterator{contract: _AavePool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_AavePool *AavePoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AavePoolWithdraw, reserve []common.Address, user []common.Address, to []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "Withdraw", reserveRule, userRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolWithdraw)
				if err := _AavePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_AavePool *AavePoolFilterer) ParseWithdraw(log types.Log) (*AavePoolWithdraw, error) {
	event := new(AavePoolWithdraw)
	if err := _AavePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
