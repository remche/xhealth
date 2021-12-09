// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lp

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
)

// LpMetaData contains all meta data concerning the Lp contract.
var LpMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtAmount\",\"type\":\"uint256\"}],\"name\":\"BorrowEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountFee\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"}],\"name\":\"LiquidateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtAmount\",\"type\":\"uint256\"}],\"name\":\"RepayEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supplyAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lptAmount\",\"type\":\"uint256\"}],\"name\":\"SupplyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"comptroller\",\"type\":\"address\"}],\"name\":\"UpdateComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optimalUtilizationRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slope1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slope2\",\"type\":\"uint256\"}],\"name\":\"UpdateInterestModelParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lptBaseValue\",\"type\":\"uint256\"}],\"name\":\"UpdateLPTBaseValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityPenaltyFactor\",\"type\":\"uint256\"}],\"name\":\"UpdateLiquidationPenaltyFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityPoolToken\",\"type\":\"address\"}],\"name\":\"UpdateLiquidityPoolToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumLoanValue\",\"type\":\"uint256\"}],\"name\":\"UpdateMinimumLoanValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"}],\"name\":\"UpdateReserveFeeFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"xtkFeeFactor\",\"type\":\"uint256\"}],\"name\":\"UpdateXtkFeeFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lptAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"xtkEarns\",\"type\":\"uint256\"}],\"name\":\"WithdrawFee\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_xAsset\",\"type\":\"address\"}],\"name\":\"addFlashBorrower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blockLockExempt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lockAddress\",\"type\":\"address\"}],\"name\":\"exemptFromBlockLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"xAsset\",\"type\":\"address\"}],\"name\":\"exemptFromLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFlashLoanFeeFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLPTBaseValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLPTValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidationPenaltyFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumLoanValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOptimalUtilizationRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlope1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlope2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getXtkFeeFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stableCoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_decimal\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastLockedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_markets\",\"type\":\"address[]\"}],\"name\":\"liquidateWithPreference\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lockAddress\",\"type\":\"address\"}],\"name\":\"removeBlockLockExemption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_xAsset\",\"type\":\"address\"}],\"name\":\"removeFlashBorrower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"xAsset\",\"type\":\"address\"}],\"name\":\"removeLiquidationExemption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_comptroller\",\"type\":\"address\"}],\"name\":\"setComptroller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFlashLoanFeeFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_optimalUtilizationRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slope2\",\"type\":\"uint256\"}],\"name\":\"setInterestModelParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lptBaseValue\",\"type\":\"uint256\"}],\"name\":\"setLPTBaseValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_liquidityPenaltyFactor\",\"type\":\"uint256\"}],\"name\":\"setLiquidationPenaltyFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityPoolToken\",\"type\":\"address\"}],\"name\":\"setLiquidityPoolToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLiquidationHealthRatio\",\"type\":\"uint256\"}],\"name\":\"setMaxLiquidationHealthRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBorrowHealthRatio\",\"type\":\"uint256\"}],\"name\":\"setMinBorrowHealthRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumLoanValue\",\"type\":\"uint256\"}],\"name\":\"setMinimumLoanValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reserveFactor\",\"type\":\"uint256\"}],\"name\":\"setReserveFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_xtkFeeFactor\",\"type\":\"uint256\"}],\"name\":\"setXtkFeeFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIxTokenManager\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setxTokenManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"supply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"updatedBorrowBy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"utilizationRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"whitelistRepay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lptAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xtkEarns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LpABI is the input ABI used to generate the binding from.
// Deprecated: Use LpMetaData.ABI instead.
var LpABI = LpMetaData.ABI

// Lp is an auto generated Go binding around an Ethereum contract.
type Lp struct {
	LpCaller     // Read-only binding to the contract
	LpTransactor // Write-only binding to the contract
	LpFilterer   // Log filterer for contract events
}

// LpCaller is an auto generated read-only Go binding around an Ethereum contract.
type LpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LpSession struct {
	Contract     *Lp               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LpCallerSession struct {
	Contract *LpCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LpTransactorSession struct {
	Contract     *LpTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LpRaw is an auto generated low-level Go binding around an Ethereum contract.
type LpRaw struct {
	Contract *Lp // Generic contract binding to access the raw methods on
}

// LpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LpCallerRaw struct {
	Contract *LpCaller // Generic read-only contract binding to access the raw methods on
}

// LpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LpTransactorRaw struct {
	Contract *LpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLp creates a new instance of Lp, bound to a specific deployed contract.
func NewLp(address common.Address, backend bind.ContractBackend) (*Lp, error) {
	contract, err := bindLp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lp{LpCaller: LpCaller{contract: contract}, LpTransactor: LpTransactor{contract: contract}, LpFilterer: LpFilterer{contract: contract}}, nil
}

// NewLpCaller creates a new read-only instance of Lp, bound to a specific deployed contract.
func NewLpCaller(address common.Address, caller bind.ContractCaller) (*LpCaller, error) {
	contract, err := bindLp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LpCaller{contract: contract}, nil
}

// NewLpTransactor creates a new write-only instance of Lp, bound to a specific deployed contract.
func NewLpTransactor(address common.Address, transactor bind.ContractTransactor) (*LpTransactor, error) {
	contract, err := bindLp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LpTransactor{contract: contract}, nil
}

// NewLpFilterer creates a new log filterer instance of Lp, bound to a specific deployed contract.
func NewLpFilterer(address common.Address, filterer bind.ContractFilterer) (*LpFilterer, error) {
	contract, err := bindLp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LpFilterer{contract: contract}, nil
}

// bindLp binds a generic wrapper to an already deployed contract.
func bindLp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lp *LpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lp.Contract.LpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lp *LpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.Contract.LpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lp *LpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lp.Contract.LpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lp *LpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lp *LpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lp *LpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lp.Contract.contract.Transact(opts, method, params...)
}

// BlockLockExempt is a free data retrieval call binding the contract method 0xfff3011b.
//
// Solidity: function blockLockExempt(address ) view returns(bool)
func (_Lp *LpCaller) BlockLockExempt(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "blockLockExempt", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlockLockExempt is a free data retrieval call binding the contract method 0xfff3011b.
//
// Solidity: function blockLockExempt(address ) view returns(bool)
func (_Lp *LpSession) BlockLockExempt(arg0 common.Address) (bool, error) {
	return _Lp.Contract.BlockLockExempt(&_Lp.CallOpts, arg0)
}

// BlockLockExempt is a free data retrieval call binding the contract method 0xfff3011b.
//
// Solidity: function blockLockExempt(address ) view returns(bool)
func (_Lp *LpCallerSession) BlockLockExempt(arg0 common.Address) (bool, error) {
	return _Lp.Contract.BlockLockExempt(&_Lp.CallOpts, arg0)
}

// BorrowRate is a free data retrieval call binding the contract method 0xc914b437.
//
// Solidity: function borrowRate() view returns(uint256)
func (_Lp *LpCaller) BorrowRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "borrowRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRate is a free data retrieval call binding the contract method 0xc914b437.
//
// Solidity: function borrowRate() view returns(uint256)
func (_Lp *LpSession) BorrowRate() (*big.Int, error) {
	return _Lp.Contract.BorrowRate(&_Lp.CallOpts)
}

// BorrowRate is a free data retrieval call binding the contract method 0xc914b437.
//
// Solidity: function borrowRate() view returns(uint256)
func (_Lp *LpCallerSession) BorrowRate() (*big.Int, error) {
	return _Lp.Contract.BorrowRate(&_Lp.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Lp *LpCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Lp *LpSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Lp.Contract.BorrowRatePerBlock(&_Lp.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Lp *LpCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Lp.Contract.BorrowRatePerBlock(&_Lp.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Lp *LpCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Lp *LpSession) Comptroller() (common.Address, error) {
	return _Lp.Contract.Comptroller(&_Lp.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Lp *LpCallerSession) Comptroller() (common.Address, error) {
	return _Lp.Contract.Comptroller(&_Lp.CallOpts)
}

// CurrentLiquidity is a free data retrieval call binding the contract method 0x46caf2ae.
//
// Solidity: function currentLiquidity() view returns(uint256)
func (_Lp *LpCaller) CurrentLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "currentLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentLiquidity is a free data retrieval call binding the contract method 0x46caf2ae.
//
// Solidity: function currentLiquidity() view returns(uint256)
func (_Lp *LpSession) CurrentLiquidity() (*big.Int, error) {
	return _Lp.Contract.CurrentLiquidity(&_Lp.CallOpts)
}

// CurrentLiquidity is a free data retrieval call binding the contract method 0x46caf2ae.
//
// Solidity: function currentLiquidity() view returns(uint256)
func (_Lp *LpCallerSession) CurrentLiquidity() (*big.Int, error) {
	return _Lp.Contract.CurrentLiquidity(&_Lp.CallOpts)
}

// GetBaseBorrowRate is a free data retrieval call binding the contract method 0xca6389f3.
//
// Solidity: function getBaseBorrowRate() view returns(uint256)
func (_Lp *LpCaller) GetBaseBorrowRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getBaseBorrowRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseBorrowRate is a free data retrieval call binding the contract method 0xca6389f3.
//
// Solidity: function getBaseBorrowRate() view returns(uint256)
func (_Lp *LpSession) GetBaseBorrowRate() (*big.Int, error) {
	return _Lp.Contract.GetBaseBorrowRate(&_Lp.CallOpts)
}

// GetBaseBorrowRate is a free data retrieval call binding the contract method 0xca6389f3.
//
// Solidity: function getBaseBorrowRate() view returns(uint256)
func (_Lp *LpCallerSession) GetBaseBorrowRate() (*big.Int, error) {
	return _Lp.Contract.GetBaseBorrowRate(&_Lp.CallOpts)
}

// GetFlashLoanFeeFactor is a free data retrieval call binding the contract method 0x14879197.
//
// Solidity: function getFlashLoanFeeFactor() view returns(uint256)
func (_Lp *LpCaller) GetFlashLoanFeeFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getFlashLoanFeeFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFlashLoanFeeFactor is a free data retrieval call binding the contract method 0x14879197.
//
// Solidity: function getFlashLoanFeeFactor() view returns(uint256)
func (_Lp *LpSession) GetFlashLoanFeeFactor() (*big.Int, error) {
	return _Lp.Contract.GetFlashLoanFeeFactor(&_Lp.CallOpts)
}

// GetFlashLoanFeeFactor is a free data retrieval call binding the contract method 0x14879197.
//
// Solidity: function getFlashLoanFeeFactor() view returns(uint256)
func (_Lp *LpCallerSession) GetFlashLoanFeeFactor() (*big.Int, error) {
	return _Lp.Contract.GetFlashLoanFeeFactor(&_Lp.CallOpts)
}

// GetLPTBaseValue is a free data retrieval call binding the contract method 0x07876eaf.
//
// Solidity: function getLPTBaseValue() view returns(uint256)
func (_Lp *LpCaller) GetLPTBaseValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getLPTBaseValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLPTBaseValue is a free data retrieval call binding the contract method 0x07876eaf.
//
// Solidity: function getLPTBaseValue() view returns(uint256)
func (_Lp *LpSession) GetLPTBaseValue() (*big.Int, error) {
	return _Lp.Contract.GetLPTBaseValue(&_Lp.CallOpts)
}

// GetLPTBaseValue is a free data retrieval call binding the contract method 0x07876eaf.
//
// Solidity: function getLPTBaseValue() view returns(uint256)
func (_Lp *LpCallerSession) GetLPTBaseValue() (*big.Int, error) {
	return _Lp.Contract.GetLPTBaseValue(&_Lp.CallOpts)
}

// GetLPTValue is a free data retrieval call binding the contract method 0xa2f32de2.
//
// Solidity: function getLPTValue() view returns(uint256)
func (_Lp *LpCaller) GetLPTValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getLPTValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLPTValue is a free data retrieval call binding the contract method 0xa2f32de2.
//
// Solidity: function getLPTValue() view returns(uint256)
func (_Lp *LpSession) GetLPTValue() (*big.Int, error) {
	return _Lp.Contract.GetLPTValue(&_Lp.CallOpts)
}

// GetLPTValue is a free data retrieval call binding the contract method 0xa2f32de2.
//
// Solidity: function getLPTValue() view returns(uint256)
func (_Lp *LpCallerSession) GetLPTValue() (*big.Int, error) {
	return _Lp.Contract.GetLPTValue(&_Lp.CallOpts)
}

// GetLiquidationPenaltyFactor is a free data retrieval call binding the contract method 0x339a03bb.
//
// Solidity: function getLiquidationPenaltyFactor() view returns(uint256)
func (_Lp *LpCaller) GetLiquidationPenaltyFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getLiquidationPenaltyFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidationPenaltyFactor is a free data retrieval call binding the contract method 0x339a03bb.
//
// Solidity: function getLiquidationPenaltyFactor() view returns(uint256)
func (_Lp *LpSession) GetLiquidationPenaltyFactor() (*big.Int, error) {
	return _Lp.Contract.GetLiquidationPenaltyFactor(&_Lp.CallOpts)
}

// GetLiquidationPenaltyFactor is a free data retrieval call binding the contract method 0x339a03bb.
//
// Solidity: function getLiquidationPenaltyFactor() view returns(uint256)
func (_Lp *LpCallerSession) GetLiquidationPenaltyFactor() (*big.Int, error) {
	return _Lp.Contract.GetLiquidationPenaltyFactor(&_Lp.CallOpts)
}

// GetMinimumLoanValue is a free data retrieval call binding the contract method 0x4bec4c55.
//
// Solidity: function getMinimumLoanValue() view returns(uint256)
func (_Lp *LpCaller) GetMinimumLoanValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getMinimumLoanValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumLoanValue is a free data retrieval call binding the contract method 0x4bec4c55.
//
// Solidity: function getMinimumLoanValue() view returns(uint256)
func (_Lp *LpSession) GetMinimumLoanValue() (*big.Int, error) {
	return _Lp.Contract.GetMinimumLoanValue(&_Lp.CallOpts)
}

// GetMinimumLoanValue is a free data retrieval call binding the contract method 0x4bec4c55.
//
// Solidity: function getMinimumLoanValue() view returns(uint256)
func (_Lp *LpCallerSession) GetMinimumLoanValue() (*big.Int, error) {
	return _Lp.Contract.GetMinimumLoanValue(&_Lp.CallOpts)
}

// GetOptimalUtilizationRate is a free data retrieval call binding the contract method 0x4d489a18.
//
// Solidity: function getOptimalUtilizationRate() view returns(uint256)
func (_Lp *LpCaller) GetOptimalUtilizationRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getOptimalUtilizationRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOptimalUtilizationRate is a free data retrieval call binding the contract method 0x4d489a18.
//
// Solidity: function getOptimalUtilizationRate() view returns(uint256)
func (_Lp *LpSession) GetOptimalUtilizationRate() (*big.Int, error) {
	return _Lp.Contract.GetOptimalUtilizationRate(&_Lp.CallOpts)
}

// GetOptimalUtilizationRate is a free data retrieval call binding the contract method 0x4d489a18.
//
// Solidity: function getOptimalUtilizationRate() view returns(uint256)
func (_Lp *LpCallerSession) GetOptimalUtilizationRate() (*big.Int, error) {
	return _Lp.Contract.GetOptimalUtilizationRate(&_Lp.CallOpts)
}

// GetReserveFactor is a free data retrieval call binding the contract method 0x5f558e53.
//
// Solidity: function getReserveFactor() view returns(uint256)
func (_Lp *LpCaller) GetReserveFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getReserveFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveFactor is a free data retrieval call binding the contract method 0x5f558e53.
//
// Solidity: function getReserveFactor() view returns(uint256)
func (_Lp *LpSession) GetReserveFactor() (*big.Int, error) {
	return _Lp.Contract.GetReserveFactor(&_Lp.CallOpts)
}

// GetReserveFactor is a free data retrieval call binding the contract method 0x5f558e53.
//
// Solidity: function getReserveFactor() view returns(uint256)
func (_Lp *LpCallerSession) GetReserveFactor() (*big.Int, error) {
	return _Lp.Contract.GetReserveFactor(&_Lp.CallOpts)
}

// GetSlope1 is a free data retrieval call binding the contract method 0x2408305d.
//
// Solidity: function getSlope1() view returns(uint256)
func (_Lp *LpCaller) GetSlope1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getSlope1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlope1 is a free data retrieval call binding the contract method 0x2408305d.
//
// Solidity: function getSlope1() view returns(uint256)
func (_Lp *LpSession) GetSlope1() (*big.Int, error) {
	return _Lp.Contract.GetSlope1(&_Lp.CallOpts)
}

// GetSlope1 is a free data retrieval call binding the contract method 0x2408305d.
//
// Solidity: function getSlope1() view returns(uint256)
func (_Lp *LpCallerSession) GetSlope1() (*big.Int, error) {
	return _Lp.Contract.GetSlope1(&_Lp.CallOpts)
}

// GetSlope2 is a free data retrieval call binding the contract method 0xa5553aee.
//
// Solidity: function getSlope2() view returns(uint256)
func (_Lp *LpCaller) GetSlope2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getSlope2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlope2 is a free data retrieval call binding the contract method 0xa5553aee.
//
// Solidity: function getSlope2() view returns(uint256)
func (_Lp *LpSession) GetSlope2() (*big.Int, error) {
	return _Lp.Contract.GetSlope2(&_Lp.CallOpts)
}

// GetSlope2 is a free data retrieval call binding the contract method 0xa5553aee.
//
// Solidity: function getSlope2() view returns(uint256)
func (_Lp *LpCallerSession) GetSlope2() (*big.Int, error) {
	return _Lp.Contract.GetSlope2(&_Lp.CallOpts)
}

// GetXtkFeeFactor is a free data retrieval call binding the contract method 0x63cf0b5b.
//
// Solidity: function getXtkFeeFactor() view returns(uint256)
func (_Lp *LpCaller) GetXtkFeeFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "getXtkFeeFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetXtkFeeFactor is a free data retrieval call binding the contract method 0x63cf0b5b.
//
// Solidity: function getXtkFeeFactor() view returns(uint256)
func (_Lp *LpSession) GetXtkFeeFactor() (*big.Int, error) {
	return _Lp.Contract.GetXtkFeeFactor(&_Lp.CallOpts)
}

// GetXtkFeeFactor is a free data retrieval call binding the contract method 0x63cf0b5b.
//
// Solidity: function getXtkFeeFactor() view returns(uint256)
func (_Lp *LpCallerSession) GetXtkFeeFactor() (*big.Int, error) {
	return _Lp.Contract.GetXtkFeeFactor(&_Lp.CallOpts)
}

// LastLockedBlock is a free data retrieval call binding the contract method 0x9f3e8b34.
//
// Solidity: function lastLockedBlock(address ) view returns(uint256)
func (_Lp *LpCaller) LastLockedBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "lastLockedBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastLockedBlock is a free data retrieval call binding the contract method 0x9f3e8b34.
//
// Solidity: function lastLockedBlock(address ) view returns(uint256)
func (_Lp *LpSession) LastLockedBlock(arg0 common.Address) (*big.Int, error) {
	return _Lp.Contract.LastLockedBlock(&_Lp.CallOpts, arg0)
}

// LastLockedBlock is a free data retrieval call binding the contract method 0x9f3e8b34.
//
// Solidity: function lastLockedBlock(address ) view returns(uint256)
func (_Lp *LpCallerSession) LastLockedBlock(arg0 common.Address) (*big.Int, error) {
	return _Lp.Contract.LastLockedBlock(&_Lp.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lp *LpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lp *LpSession) Owner() (common.Address, error) {
	return _Lp.Contract.Owner(&_Lp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lp *LpCallerSession) Owner() (common.Address, error) {
	return _Lp.Contract.Owner(&_Lp.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lp *LpCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lp *LpSession) Paused() (bool, error) {
	return _Lp.Contract.Paused(&_Lp.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lp *LpCallerSession) Paused() (bool, error) {
	return _Lp.Contract.Paused(&_Lp.CallOpts)
}

// Reserves is a free data retrieval call binding the contract method 0x75172a8b.
//
// Solidity: function reserves() view returns(uint256)
func (_Lp *LpCaller) Reserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "reserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserves is a free data retrieval call binding the contract method 0x75172a8b.
//
// Solidity: function reserves() view returns(uint256)
func (_Lp *LpSession) Reserves() (*big.Int, error) {
	return _Lp.Contract.Reserves(&_Lp.CallOpts)
}

// Reserves is a free data retrieval call binding the contract method 0x75172a8b.
//
// Solidity: function reserves() view returns(uint256)
func (_Lp *LpCallerSession) Reserves() (*big.Int, error) {
	return _Lp.Contract.Reserves(&_Lp.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Lp *LpCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Lp *LpSession) TotalBorrows() (*big.Int, error) {
	return _Lp.Contract.TotalBorrows(&_Lp.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Lp *LpCallerSession) TotalBorrows() (*big.Int, error) {
	return _Lp.Contract.TotalBorrows(&_Lp.CallOpts)
}

// UpdatedBorrowBy is a free data retrieval call binding the contract method 0x2aad6aa8.
//
// Solidity: function updatedBorrowBy(address _borrower) view returns(uint256)
func (_Lp *LpCaller) UpdatedBorrowBy(opts *bind.CallOpts, _borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "updatedBorrowBy", _borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdatedBorrowBy is a free data retrieval call binding the contract method 0x2aad6aa8.
//
// Solidity: function updatedBorrowBy(address _borrower) view returns(uint256)
func (_Lp *LpSession) UpdatedBorrowBy(_borrower common.Address) (*big.Int, error) {
	return _Lp.Contract.UpdatedBorrowBy(&_Lp.CallOpts, _borrower)
}

// UpdatedBorrowBy is a free data retrieval call binding the contract method 0x2aad6aa8.
//
// Solidity: function updatedBorrowBy(address _borrower) view returns(uint256)
func (_Lp *LpCallerSession) UpdatedBorrowBy(_borrower common.Address) (*big.Int, error) {
	return _Lp.Contract.UpdatedBorrowBy(&_Lp.CallOpts, _borrower)
}

// UtilizationRate is a free data retrieval call binding the contract method 0x6c321c8a.
//
// Solidity: function utilizationRate() view returns(uint256)
func (_Lp *LpCaller) UtilizationRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "utilizationRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UtilizationRate is a free data retrieval call binding the contract method 0x6c321c8a.
//
// Solidity: function utilizationRate() view returns(uint256)
func (_Lp *LpSession) UtilizationRate() (*big.Int, error) {
	return _Lp.Contract.UtilizationRate(&_Lp.CallOpts)
}

// UtilizationRate is a free data retrieval call binding the contract method 0x6c321c8a.
//
// Solidity: function utilizationRate() view returns(uint256)
func (_Lp *LpCallerSession) UtilizationRate() (*big.Int, error) {
	return _Lp.Contract.UtilizationRate(&_Lp.CallOpts)
}

// XtkEarns is a free data retrieval call binding the contract method 0x27c7b1a6.
//
// Solidity: function xtkEarns() view returns(uint256)
func (_Lp *LpCaller) XtkEarns(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "xtkEarns")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XtkEarns is a free data retrieval call binding the contract method 0x27c7b1a6.
//
// Solidity: function xtkEarns() view returns(uint256)
func (_Lp *LpSession) XtkEarns() (*big.Int, error) {
	return _Lp.Contract.XtkEarns(&_Lp.CallOpts)
}

// XtkEarns is a free data retrieval call binding the contract method 0x27c7b1a6.
//
// Solidity: function xtkEarns() view returns(uint256)
func (_Lp *LpCallerSession) XtkEarns() (*big.Int, error) {
	return _Lp.Contract.XtkEarns(&_Lp.CallOpts)
}

// AddFlashBorrower is a paid mutator transaction binding the contract method 0x9ac9d80b.
//
// Solidity: function addFlashBorrower(address _xAsset) returns()
func (_Lp *LpTransactor) AddFlashBorrower(opts *bind.TransactOpts, _xAsset common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "addFlashBorrower", _xAsset)
}

// AddFlashBorrower is a paid mutator transaction binding the contract method 0x9ac9d80b.
//
// Solidity: function addFlashBorrower(address _xAsset) returns()
func (_Lp *LpSession) AddFlashBorrower(_xAsset common.Address) (*types.Transaction, error) {
	return _Lp.Contract.AddFlashBorrower(&_Lp.TransactOpts, _xAsset)
}

// AddFlashBorrower is a paid mutator transaction binding the contract method 0x9ac9d80b.
//
// Solidity: function addFlashBorrower(address _xAsset) returns()
func (_Lp *LpTransactorSession) AddFlashBorrower(_xAsset common.Address) (*types.Transaction, error) {
	return _Lp.Contract.AddFlashBorrower(&_Lp.TransactOpts, _xAsset)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 _amount) returns()
func (_Lp *LpTransactor) Borrow(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "borrow", _amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 _amount) returns()
func (_Lp *LpSession) Borrow(_amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Borrow(&_Lp.TransactOpts, _amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 _amount) returns()
func (_Lp *LpTransactorSession) Borrow(_amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Borrow(&_Lp.TransactOpts, _amount)
}

// ExemptFromBlockLock is a paid mutator transaction binding the contract method 0xd896cf19.
//
// Solidity: function exemptFromBlockLock(address lockAddress) returns()
func (_Lp *LpTransactor) ExemptFromBlockLock(opts *bind.TransactOpts, lockAddress common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "exemptFromBlockLock", lockAddress)
}

// ExemptFromBlockLock is a paid mutator transaction binding the contract method 0xd896cf19.
//
// Solidity: function exemptFromBlockLock(address lockAddress) returns()
func (_Lp *LpSession) ExemptFromBlockLock(lockAddress common.Address) (*types.Transaction, error) {
	return _Lp.Contract.ExemptFromBlockLock(&_Lp.TransactOpts, lockAddress)
}

// ExemptFromBlockLock is a paid mutator transaction binding the contract method 0xd896cf19.
//
// Solidity: function exemptFromBlockLock(address lockAddress) returns()
func (_Lp *LpTransactorSession) ExemptFromBlockLock(lockAddress common.Address) (*types.Transaction, error) {
	return _Lp.Contract.ExemptFromBlockLock(&_Lp.TransactOpts, lockAddress)
}

// ExemptFromLiquidation is a paid mutator transaction binding the contract method 0x288a97ad.
//
// Solidity: function exemptFromLiquidation(address xAsset) returns()
func (_Lp *LpTransactor) ExemptFromLiquidation(opts *bind.TransactOpts, xAsset common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "exemptFromLiquidation", xAsset)
}

// ExemptFromLiquidation is a paid mutator transaction binding the contract method 0x288a97ad.
//
// Solidity: function exemptFromLiquidation(address xAsset) returns()
func (_Lp *LpSession) ExemptFromLiquidation(xAsset common.Address) (*types.Transaction, error) {
	return _Lp.Contract.ExemptFromLiquidation(&_Lp.TransactOpts, xAsset)
}

// ExemptFromLiquidation is a paid mutator transaction binding the contract method 0x288a97ad.
//
// Solidity: function exemptFromLiquidation(address xAsset) returns()
func (_Lp *LpTransactorSession) ExemptFromLiquidation(xAsset common.Address) (*types.Transaction, error) {
	return _Lp.Contract.ExemptFromLiquidation(&_Lp.TransactOpts, xAsset)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xe0232b42.
//
// Solidity: function flashLoan(address _receiver, uint256 _amount, bytes _params) returns()
func (_Lp *LpTransactor) FlashLoan(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "flashLoan", _receiver, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xe0232b42.
//
// Solidity: function flashLoan(address _receiver, uint256 _amount, bytes _params) returns()
func (_Lp *LpSession) FlashLoan(_receiver common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _Lp.Contract.FlashLoan(&_Lp.TransactOpts, _receiver, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xe0232b42.
//
// Solidity: function flashLoan(address _receiver, uint256 _amount, bytes _params) returns()
func (_Lp *LpTransactorSession) FlashLoan(_receiver common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _Lp.Contract.FlashLoan(&_Lp.TransactOpts, _receiver, _amount, _params)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _stableCoin, uint256 _decimal) returns()
func (_Lp *LpTransactor) Initialize(opts *bind.TransactOpts, _stableCoin common.Address, _decimal *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "initialize", _stableCoin, _decimal)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _stableCoin, uint256 _decimal) returns()
func (_Lp *LpSession) Initialize(_stableCoin common.Address, _decimal *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Initialize(&_Lp.TransactOpts, _stableCoin, _decimal)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _stableCoin, uint256 _decimal) returns()
func (_Lp *LpTransactorSession) Initialize(_stableCoin common.Address, _decimal *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Initialize(&_Lp.TransactOpts, _stableCoin, _decimal)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbcbaf487.
//
// Solidity: function liquidate(address _borrower, uint256 _amount) returns()
func (_Lp *LpTransactor) Liquidate(opts *bind.TransactOpts, _borrower common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "liquidate", _borrower, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbcbaf487.
//
// Solidity: function liquidate(address _borrower, uint256 _amount) returns()
func (_Lp *LpSession) Liquidate(_borrower common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Liquidate(&_Lp.TransactOpts, _borrower, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbcbaf487.
//
// Solidity: function liquidate(address _borrower, uint256 _amount) returns()
func (_Lp *LpTransactorSession) Liquidate(_borrower common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Liquidate(&_Lp.TransactOpts, _borrower, _amount)
}

// LiquidateWithPreference is a paid mutator transaction binding the contract method 0x5abbab16.
//
// Solidity: function liquidateWithPreference(address _borrower, uint256 _amount, address[] _markets) returns()
func (_Lp *LpTransactor) LiquidateWithPreference(opts *bind.TransactOpts, _borrower common.Address, _amount *big.Int, _markets []common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "liquidateWithPreference", _borrower, _amount, _markets)
}

// LiquidateWithPreference is a paid mutator transaction binding the contract method 0x5abbab16.
//
// Solidity: function liquidateWithPreference(address _borrower, uint256 _amount, address[] _markets) returns()
func (_Lp *LpSession) LiquidateWithPreference(_borrower common.Address, _amount *big.Int, _markets []common.Address) (*types.Transaction, error) {
	return _Lp.Contract.LiquidateWithPreference(&_Lp.TransactOpts, _borrower, _amount, _markets)
}

// LiquidateWithPreference is a paid mutator transaction binding the contract method 0x5abbab16.
//
// Solidity: function liquidateWithPreference(address _borrower, uint256 _amount, address[] _markets) returns()
func (_Lp *LpTransactorSession) LiquidateWithPreference(_borrower common.Address, _amount *big.Int, _markets []common.Address) (*types.Transaction, error) {
	return _Lp.Contract.LiquidateWithPreference(&_Lp.TransactOpts, _borrower, _amount, _markets)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_Lp *LpTransactor) PauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "pauseContract")
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_Lp *LpSession) PauseContract() (*types.Transaction, error) {
	return _Lp.Contract.PauseContract(&_Lp.TransactOpts)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_Lp *LpTransactorSession) PauseContract() (*types.Transaction, error) {
	return _Lp.Contract.PauseContract(&_Lp.TransactOpts)
}

// PayAll is a paid mutator transaction binding the contract method 0x5806beaf.
//
// Solidity: function payAll() returns()
func (_Lp *LpTransactor) PayAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "payAll")
}

// PayAll is a paid mutator transaction binding the contract method 0x5806beaf.
//
// Solidity: function payAll() returns()
func (_Lp *LpSession) PayAll() (*types.Transaction, error) {
	return _Lp.Contract.PayAll(&_Lp.TransactOpts)
}

// PayAll is a paid mutator transaction binding the contract method 0x5806beaf.
//
// Solidity: function payAll() returns()
func (_Lp *LpTransactorSession) PayAll() (*types.Transaction, error) {
	return _Lp.Contract.PayAll(&_Lp.TransactOpts)
}

// RemoveBlockLockExemption is a paid mutator transaction binding the contract method 0x1de34d39.
//
// Solidity: function removeBlockLockExemption(address lockAddress) returns()
func (_Lp *LpTransactor) RemoveBlockLockExemption(opts *bind.TransactOpts, lockAddress common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "removeBlockLockExemption", lockAddress)
}

// RemoveBlockLockExemption is a paid mutator transaction binding the contract method 0x1de34d39.
//
// Solidity: function removeBlockLockExemption(address lockAddress) returns()
func (_Lp *LpSession) RemoveBlockLockExemption(lockAddress common.Address) (*types.Transaction, error) {
	return _Lp.Contract.RemoveBlockLockExemption(&_Lp.TransactOpts, lockAddress)
}

// RemoveBlockLockExemption is a paid mutator transaction binding the contract method 0x1de34d39.
//
// Solidity: function removeBlockLockExemption(address lockAddress) returns()
func (_Lp *LpTransactorSession) RemoveBlockLockExemption(lockAddress common.Address) (*types.Transaction, error) {
	return _Lp.Contract.RemoveBlockLockExemption(&_Lp.TransactOpts, lockAddress)
}

// RemoveFlashBorrower is a paid mutator transaction binding the contract method 0x253cf980.
//
// Solidity: function removeFlashBorrower(address _xAsset) returns()
func (_Lp *LpTransactor) RemoveFlashBorrower(opts *bind.TransactOpts, _xAsset common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "removeFlashBorrower", _xAsset)
}

// RemoveFlashBorrower is a paid mutator transaction binding the contract method 0x253cf980.
//
// Solidity: function removeFlashBorrower(address _xAsset) returns()
func (_Lp *LpSession) RemoveFlashBorrower(_xAsset common.Address) (*types.Transaction, error) {
	return _Lp.Contract.RemoveFlashBorrower(&_Lp.TransactOpts, _xAsset)
}

// RemoveFlashBorrower is a paid mutator transaction binding the contract method 0x253cf980.
//
// Solidity: function removeFlashBorrower(address _xAsset) returns()
func (_Lp *LpTransactorSession) RemoveFlashBorrower(_xAsset common.Address) (*types.Transaction, error) {
	return _Lp.Contract.RemoveFlashBorrower(&_Lp.TransactOpts, _xAsset)
}

// RemoveLiquidationExemption is a paid mutator transaction binding the contract method 0x7225ff0c.
//
// Solidity: function removeLiquidationExemption(address xAsset) returns()
func (_Lp *LpTransactor) RemoveLiquidationExemption(opts *bind.TransactOpts, xAsset common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "removeLiquidationExemption", xAsset)
}

// RemoveLiquidationExemption is a paid mutator transaction binding the contract method 0x7225ff0c.
//
// Solidity: function removeLiquidationExemption(address xAsset) returns()
func (_Lp *LpSession) RemoveLiquidationExemption(xAsset common.Address) (*types.Transaction, error) {
	return _Lp.Contract.RemoveLiquidationExemption(&_Lp.TransactOpts, xAsset)
}

// RemoveLiquidationExemption is a paid mutator transaction binding the contract method 0x7225ff0c.
//
// Solidity: function removeLiquidationExemption(address xAsset) returns()
func (_Lp *LpTransactorSession) RemoveLiquidationExemption(xAsset common.Address) (*types.Transaction, error) {
	return _Lp.Contract.RemoveLiquidationExemption(&_Lp.TransactOpts, xAsset)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lp *LpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lp *LpSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lp.Contract.RenounceOwnership(&_Lp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lp *LpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lp.Contract.RenounceOwnership(&_Lp.TransactOpts)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _amount) returns()
func (_Lp *LpTransactor) Repay(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "repay", _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _amount) returns()
func (_Lp *LpSession) Repay(_amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Repay(&_Lp.TransactOpts, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _amount) returns()
func (_Lp *LpTransactorSession) Repay(_amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Repay(&_Lp.TransactOpts, _amount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x8bad38dd.
//
// Solidity: function setComptroller(address _comptroller) returns()
func (_Lp *LpTransactor) SetComptroller(opts *bind.TransactOpts, _comptroller common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "setComptroller", _comptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x8bad38dd.
//
// Solidity: function setComptroller(address _comptroller) returns()
func (_Lp *LpSession) SetComptroller(_comptroller common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SetComptroller(&_Lp.TransactOpts, _comptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x8bad38dd.
//
// Solidity: function setComptroller(address _comptroller) returns()
func (_Lp *LpTransactorSession) SetComptroller(_comptroller common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SetComptroller(&_Lp.TransactOpts, _comptroller)
}

// SetFlashLoanFeeFactor is a paid mutator transaction binding the contract method 0x24eddc92.
//
// Solidity: function setFlashLoanFeeFactor(uint256 _fee) returns()
func (_Lp *LpTransactor) SetFlashLoanFeeFactor(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "setFlashLoanFeeFactor", _fee)
}

// SetFlashLoanFeeFactor is a paid mutator transaction binding the contract method 0x24eddc92.
//
// Solidity: function setFlashLoanFeeFactor(uint256 _fee) returns()
func (_Lp *LpSession) SetFlashLoanFeeFactor(_fee *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetFlashLoanFeeFactor(&_Lp.TransactOpts, _fee)
}

// SetFlashLoanFeeFactor is a paid mutator transaction binding the contract method 0x24eddc92.
//
// Solidity: function setFlashLoanFeeFactor(uint256 _fee) returns()
func (_Lp *LpTransactorSession) SetFlashLoanFeeFactor(_fee *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetFlashLoanFeeFactor(&_Lp.TransactOpts, _fee)
}

// SetInterestModelParameters is a paid mutator transaction binding the contract method 0xb2e2e180.
//
// Solidity: function setInterestModelParameters(uint256 _optimalUtilizationRate, uint256 _baseBorrowRate, uint256 _slope1, uint256 _slope2) returns()
func (_Lp *LpTransactor) SetInterestModelParameters(opts *bind.TransactOpts, _optimalUtilizationRate *big.Int, _baseBorrowRate *big.Int, _slope1 *big.Int, _slope2 *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "setInterestModelParameters", _optimalUtilizationRate, _baseBorrowRate, _slope1, _slope2)
}

// SetInterestModelParameters is a paid mutator transaction binding the contract method 0xb2e2e180.
//
// Solidity: function setInterestModelParameters(uint256 _optimalUtilizationRate, uint256 _baseBorrowRate, uint256 _slope1, uint256 _slope2) returns()
func (_Lp *LpSession) SetInterestModelParameters(_optimalUtilizationRate *big.Int, _baseBorrowRate *big.Int, _slope1 *big.Int, _slope2 *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetInterestModelParameters(&_Lp.TransactOpts, _optimalUtilizationRate, _baseBorrowRate, _slope1, _slope2)
}

// SetInterestModelParameters is a paid mutator transaction binding the contract method 0xb2e2e180.
//
// Solidity: function setInterestModelParameters(uint256 _optimalUtilizationRate, uint256 _baseBorrowRate, uint256 _slope1, uint256 _slope2) returns()
func (_Lp *LpTransactorSession) SetInterestModelParameters(_optimalUtilizationRate *big.Int, _baseBorrowRate *big.Int, _slope1 *big.Int, _slope2 *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetInterestModelParameters(&_Lp.TransactOpts, _optimalUtilizationRate, _baseBorrowRate, _slope1, _slope2)
}

// SetLPTBaseValue is a paid mutator transaction binding the contract method 0xfec7ab25.
//
// Solidity: function setLPTBaseValue(uint256 _lptBaseValue) returns()
func (_Lp *LpTransactor) SetLPTBaseValue(opts *bind.TransactOpts, _lptBaseValue *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "setLPTBaseValue", _lptBaseValue)
}

// SetLPTBaseValue is a paid mutator transaction binding the contract method 0xfec7ab25.
//
// Solidity: function setLPTBaseValue(uint256 _lptBaseValue) returns()
func (_Lp *LpSession) SetLPTBaseValue(_lptBaseValue *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetLPTBaseValue(&_Lp.TransactOpts, _lptBaseValue)
}

// SetLPTBaseValue is a paid mutator transaction binding the contract method 0xfec7ab25.
//
// Solidity: function setLPTBaseValue(uint256 _lptBaseValue) returns()
func (_Lp *LpTransactorSession) SetLPTBaseValue(_lptBaseValue *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetLPTBaseValue(&_Lp.TransactOpts, _lptBaseValue)
}

// SetLiquidationPenaltyFactor is a paid mutator transaction binding the contract method 0xfb8b107d.
//
// Solidity: function setLiquidationPenaltyFactor(uint256 _liquidityPenaltyFactor) returns()
func (_Lp *LpTransactor) SetLiquidationPenaltyFactor(opts *bind.TransactOpts, _liquidityPenaltyFactor *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "setLiquidationPenaltyFactor", _liquidityPenaltyFactor)
}

// SetLiquidationPenaltyFactor is a paid mutator transaction binding the contract method 0xfb8b107d.
//
// Solidity: function setLiquidationPenaltyFactor(uint256 _liquidityPenaltyFactor) returns()
func (_Lp *LpSession) SetLiquidationPenaltyFactor(_liquidityPenaltyFactor *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetLiquidationPenaltyFactor(&_Lp.TransactOpts, _liquidityPenaltyFactor)
}

// SetLiquidationPenaltyFactor is a paid mutator transaction binding the contract method 0xfb8b107d.
//
// Solidity: function setLiquidationPenaltyFactor(uint256 _liquidityPenaltyFactor) returns()
func (_Lp *LpTransactorSession) SetLiquidationPenaltyFactor(_liquidityPenaltyFactor *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetLiquidationPenaltyFactor(&_Lp.TransactOpts, _liquidityPenaltyFactor)
}

// SetLiquidityPoolToken is a paid mutator transaction binding the contract method 0xa23d8de8.
//
// Solidity: function setLiquidityPoolToken(address _liquidityPoolToken) returns()
func (_Lp *LpTransactor) SetLiquidityPoolToken(opts *bind.TransactOpts, _liquidityPoolToken common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "setLiquidityPoolToken", _liquidityPoolToken)
}

// SetLiquidityPoolToken is a paid mutator transaction binding the contract method 0xa23d8de8.
//
// Solidity: function setLiquidityPoolToken(address _liquidityPoolToken) returns()
func (_Lp *LpSession) SetLiquidityPoolToken(_liquidityPoolToken common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SetLiquidityPoolToken(&_Lp.TransactOpts, _liquidityPoolToken)
}

// SetLiquidityPoolToken is a paid mutator transaction binding the contract method 0xa23d8de8.
//
// Solidity: function setLiquidityPoolToken(address _liquidityPoolToken) returns()
func (_Lp *LpTransactorSession) SetLiquidityPoolToken(_liquidityPoolToken common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SetLiquidityPoolToken(&_Lp.TransactOpts, _liquidityPoolToken)
}

// SetMaxLiquidationHealthRatio is a paid mutator transaction binding the contract method 0xaeb70fb1.
//
// Solidity: function setMaxLiquidationHealthRatio(uint256 _maxLiquidationHealthRatio) returns()
func (_Lp *LpTransactor) SetMaxLiquidationHealthRatio(opts *bind.TransactOpts, _maxLiquidationHealthRatio *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "setMaxLiquidationHealthRatio", _maxLiquidationHealthRatio)
}

// SetMaxLiquidationHealthRatio is a paid mutator transaction binding the contract method 0xaeb70fb1.
//
// Solidity: function setMaxLiquidationHealthRatio(uint256 _maxLiquidationHealthRatio) returns()
func (_Lp *LpSession) SetMaxLiquidationHealthRatio(_maxLiquidationHealthRatio *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetMaxLiquidationHealthRatio(&_Lp.TransactOpts, _maxLiquidationHealthRatio)
}

// SetMaxLiquidationHealthRatio is a paid mutator transaction binding the contract method 0xaeb70fb1.
//
// Solidity: function setMaxLiquidationHealthRatio(uint256 _maxLiquidationHealthRatio) returns()
func (_Lp *LpTransactorSession) SetMaxLiquidationHealthRatio(_maxLiquidationHealthRatio *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetMaxLiquidationHealthRatio(&_Lp.TransactOpts, _maxLiquidationHealthRatio)
}

// SetMinBorrowHealthRatio is a paid mutator transaction binding the contract method 0x26173ebb.
//
// Solidity: function setMinBorrowHealthRatio(uint256 _minBorrowHealthRatio) returns()
func (_Lp *LpTransactor) SetMinBorrowHealthRatio(opts *bind.TransactOpts, _minBorrowHealthRatio *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "setMinBorrowHealthRatio", _minBorrowHealthRatio)
}

// SetMinBorrowHealthRatio is a paid mutator transaction binding the contract method 0x26173ebb.
//
// Solidity: function setMinBorrowHealthRatio(uint256 _minBorrowHealthRatio) returns()
func (_Lp *LpSession) SetMinBorrowHealthRatio(_minBorrowHealthRatio *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetMinBorrowHealthRatio(&_Lp.TransactOpts, _minBorrowHealthRatio)
}

// SetMinBorrowHealthRatio is a paid mutator transaction binding the contract method 0x26173ebb.
//
// Solidity: function setMinBorrowHealthRatio(uint256 _minBorrowHealthRatio) returns()
func (_Lp *LpTransactorSession) SetMinBorrowHealthRatio(_minBorrowHealthRatio *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetMinBorrowHealthRatio(&_Lp.TransactOpts, _minBorrowHealthRatio)
}

// SetMinimumLoanValue is a paid mutator transaction binding the contract method 0xabcb4852.
//
// Solidity: function setMinimumLoanValue(uint256 _minimumLoanValue) returns()
func (_Lp *LpTransactor) SetMinimumLoanValue(opts *bind.TransactOpts, _minimumLoanValue *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "setMinimumLoanValue", _minimumLoanValue)
}

// SetMinimumLoanValue is a paid mutator transaction binding the contract method 0xabcb4852.
//
// Solidity: function setMinimumLoanValue(uint256 _minimumLoanValue) returns()
func (_Lp *LpSession) SetMinimumLoanValue(_minimumLoanValue *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetMinimumLoanValue(&_Lp.TransactOpts, _minimumLoanValue)
}

// SetMinimumLoanValue is a paid mutator transaction binding the contract method 0xabcb4852.
//
// Solidity: function setMinimumLoanValue(uint256 _minimumLoanValue) returns()
func (_Lp *LpTransactorSession) SetMinimumLoanValue(_minimumLoanValue *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetMinimumLoanValue(&_Lp.TransactOpts, _minimumLoanValue)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0x1c446983.
//
// Solidity: function setReserveFactor(uint256 _reserveFactor) returns()
func (_Lp *LpTransactor) SetReserveFactor(opts *bind.TransactOpts, _reserveFactor *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "setReserveFactor", _reserveFactor)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0x1c446983.
//
// Solidity: function setReserveFactor(uint256 _reserveFactor) returns()
func (_Lp *LpSession) SetReserveFactor(_reserveFactor *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetReserveFactor(&_Lp.TransactOpts, _reserveFactor)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0x1c446983.
//
// Solidity: function setReserveFactor(uint256 _reserveFactor) returns()
func (_Lp *LpTransactorSession) SetReserveFactor(_reserveFactor *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetReserveFactor(&_Lp.TransactOpts, _reserveFactor)
}

// SetXtkFeeFactor is a paid mutator transaction binding the contract method 0x1ab49cc9.
//
// Solidity: function setXtkFeeFactor(uint256 _xtkFeeFactor) returns()
func (_Lp *LpTransactor) SetXtkFeeFactor(opts *bind.TransactOpts, _xtkFeeFactor *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "setXtkFeeFactor", _xtkFeeFactor)
}

// SetXtkFeeFactor is a paid mutator transaction binding the contract method 0x1ab49cc9.
//
// Solidity: function setXtkFeeFactor(uint256 _xtkFeeFactor) returns()
func (_Lp *LpSession) SetXtkFeeFactor(_xtkFeeFactor *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetXtkFeeFactor(&_Lp.TransactOpts, _xtkFeeFactor)
}

// SetXtkFeeFactor is a paid mutator transaction binding the contract method 0x1ab49cc9.
//
// Solidity: function setXtkFeeFactor(uint256 _xtkFeeFactor) returns()
func (_Lp *LpTransactorSession) SetXtkFeeFactor(_xtkFeeFactor *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.SetXtkFeeFactor(&_Lp.TransactOpts, _xtkFeeFactor)
}

// SetxTokenManager is a paid mutator transaction binding the contract method 0x5981b5d3.
//
// Solidity: function setxTokenManager(address _manager) returns()
func (_Lp *LpTransactor) SetxTokenManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "setxTokenManager", _manager)
}

// SetxTokenManager is a paid mutator transaction binding the contract method 0x5981b5d3.
//
// Solidity: function setxTokenManager(address _manager) returns()
func (_Lp *LpSession) SetxTokenManager(_manager common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SetxTokenManager(&_Lp.TransactOpts, _manager)
}

// SetxTokenManager is a paid mutator transaction binding the contract method 0x5981b5d3.
//
// Solidity: function setxTokenManager(address _manager) returns()
func (_Lp *LpTransactorSession) SetxTokenManager(_manager common.Address) (*types.Transaction, error) {
	return _Lp.Contract.SetxTokenManager(&_Lp.TransactOpts, _manager)
}

// Supply is a paid mutator transaction binding the contract method 0x35403023.
//
// Solidity: function supply(uint256 _amount) returns()
func (_Lp *LpTransactor) Supply(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "supply", _amount)
}

// Supply is a paid mutator transaction binding the contract method 0x35403023.
//
// Solidity: function supply(uint256 _amount) returns()
func (_Lp *LpSession) Supply(_amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Supply(&_Lp.TransactOpts, _amount)
}

// Supply is a paid mutator transaction binding the contract method 0x35403023.
//
// Solidity: function supply(uint256 _amount) returns()
func (_Lp *LpTransactorSession) Supply(_amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Supply(&_Lp.TransactOpts, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lp *LpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lp *LpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lp.Contract.TransferOwnership(&_Lp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lp *LpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lp.Contract.TransferOwnership(&_Lp.TransactOpts, newOwner)
}

// UnpauseContract is a paid mutator transaction binding the contract method 0xb33712c5.
//
// Solidity: function unpauseContract() returns()
func (_Lp *LpTransactor) UnpauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "unpauseContract")
}

// UnpauseContract is a paid mutator transaction binding the contract method 0xb33712c5.
//
// Solidity: function unpauseContract() returns()
func (_Lp *LpSession) UnpauseContract() (*types.Transaction, error) {
	return _Lp.Contract.UnpauseContract(&_Lp.TransactOpts)
}

// UnpauseContract is a paid mutator transaction binding the contract method 0xb33712c5.
//
// Solidity: function unpauseContract() returns()
func (_Lp *LpTransactorSession) UnpauseContract() (*types.Transaction, error) {
	return _Lp.Contract.UnpauseContract(&_Lp.TransactOpts)
}

// WhitelistRepay is a paid mutator transaction binding the contract method 0xe33b4252.
//
// Solidity: function whitelistRepay(uint256 _amount) returns()
func (_Lp *LpTransactor) WhitelistRepay(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "whitelistRepay", _amount)
}

// WhitelistRepay is a paid mutator transaction binding the contract method 0xe33b4252.
//
// Solidity: function whitelistRepay(uint256 _amount) returns()
func (_Lp *LpSession) WhitelistRepay(_amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.WhitelistRepay(&_Lp.TransactOpts, _amount)
}

// WhitelistRepay is a paid mutator transaction binding the contract method 0xe33b4252.
//
// Solidity: function whitelistRepay(uint256 _amount) returns()
func (_Lp *LpTransactorSession) WhitelistRepay(_amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.WhitelistRepay(&_Lp.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _lptAmount) returns()
func (_Lp *LpTransactor) Withdraw(opts *bind.TransactOpts, _lptAmount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "withdraw", _lptAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _lptAmount) returns()
func (_Lp *LpSession) Withdraw(_lptAmount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Withdraw(&_Lp.TransactOpts, _lptAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _lptAmount) returns()
func (_Lp *LpTransactorSession) Withdraw(_lptAmount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Withdraw(&_Lp.TransactOpts, _lptAmount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address _recipient) returns()
func (_Lp *LpTransactor) WithdrawFees(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "withdrawFees", _recipient)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address _recipient) returns()
func (_Lp *LpSession) WithdrawFees(_recipient common.Address) (*types.Transaction, error) {
	return _Lp.Contract.WithdrawFees(&_Lp.TransactOpts, _recipient)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address _recipient) returns()
func (_Lp *LpTransactorSession) WithdrawFees(_recipient common.Address) (*types.Transaction, error) {
	return _Lp.Contract.WithdrawFees(&_Lp.TransactOpts, _recipient)
}

// LpBorrowEventIterator is returned from FilterBorrowEvent and is used to iterate over the raw logs and unpacked data for BorrowEvent events raised by the Lp contract.
type LpBorrowEventIterator struct {
	Event *LpBorrowEvent // Event containing the contract specifics and raw log

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
func (it *LpBorrowEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpBorrowEvent)
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
		it.Event = new(LpBorrowEvent)
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
func (it *LpBorrowEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpBorrowEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpBorrowEvent represents a BorrowEvent event raised by the Lp contract.
type LpBorrowEvent struct {
	Borrower     common.Address
	BorrowAmount *big.Int
	DebtAmount   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBorrowEvent is a free log retrieval operation binding the contract event 0xf1cdf2a0095321c5e732289e620e260dadae116ee66e758396dd844db75eb88c.
//
// Solidity: event BorrowEvent(address indexed borrower, uint256 borrowAmount, uint256 debtAmount)
func (_Lp *LpFilterer) FilterBorrowEvent(opts *bind.FilterOpts, borrower []common.Address) (*LpBorrowEventIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "BorrowEvent", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LpBorrowEventIterator{contract: _Lp.contract, event: "BorrowEvent", logs: logs, sub: sub}, nil
}

// WatchBorrowEvent is a free log subscription operation binding the contract event 0xf1cdf2a0095321c5e732289e620e260dadae116ee66e758396dd844db75eb88c.
//
// Solidity: event BorrowEvent(address indexed borrower, uint256 borrowAmount, uint256 debtAmount)
func (_Lp *LpFilterer) WatchBorrowEvent(opts *bind.WatchOpts, sink chan<- *LpBorrowEvent, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "BorrowEvent", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpBorrowEvent)
				if err := _Lp.contract.UnpackLog(event, "BorrowEvent", log); err != nil {
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

// ParseBorrowEvent is a log parse operation binding the contract event 0xf1cdf2a0095321c5e732289e620e260dadae116ee66e758396dd844db75eb88c.
//
// Solidity: event BorrowEvent(address indexed borrower, uint256 borrowAmount, uint256 debtAmount)
func (_Lp *LpFilterer) ParseBorrowEvent(log types.Log) (*LpBorrowEvent, error) {
	event := new(LpBorrowEvent)
	if err := _Lp.contract.UnpackLog(event, "BorrowEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the Lp contract.
type LpFlashLoanIterator struct {
	Event *LpFlashLoan // Event containing the contract specifics and raw log

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
func (it *LpFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpFlashLoan)
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
		it.Event = new(LpFlashLoan)
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
func (it *LpFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpFlashLoan represents a FlashLoan event raised by the Lp contract.
type LpFlashLoan struct {
	Receiver  common.Address
	Amount    *big.Int
	AmountFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x31aaad38f00845a242d16ae90d7bd72fc68f0e22581470f9dc0de241210c2886.
//
// Solidity: event FlashLoan(address indexed receiver, uint256 amount, uint256 amountFee)
func (_Lp *LpFilterer) FilterFlashLoan(opts *bind.FilterOpts, receiver []common.Address) (*LpFlashLoanIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "FlashLoan", receiverRule)
	if err != nil {
		return nil, err
	}
	return &LpFlashLoanIterator{contract: _Lp.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x31aaad38f00845a242d16ae90d7bd72fc68f0e22581470f9dc0de241210c2886.
//
// Solidity: event FlashLoan(address indexed receiver, uint256 amount, uint256 amountFee)
func (_Lp *LpFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *LpFlashLoan, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "FlashLoan", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpFlashLoan)
				if err := _Lp.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x31aaad38f00845a242d16ae90d7bd72fc68f0e22581470f9dc0de241210c2886.
//
// Solidity: event FlashLoan(address indexed receiver, uint256 amount, uint256 amountFee)
func (_Lp *LpFilterer) ParseFlashLoan(log types.Log) (*LpFlashLoan, error) {
	event := new(LpFlashLoan)
	if err := _Lp.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpLiquidateEventIterator is returned from FilterLiquidateEvent and is used to iterate over the raw logs and unpacked data for LiquidateEvent events raised by the Lp contract.
type LpLiquidateEventIterator struct {
	Event *LpLiquidateEvent // Event containing the contract specifics and raw log

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
func (it *LpLiquidateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpLiquidateEvent)
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
		it.Event = new(LpLiquidateEvent)
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
func (it *LpLiquidateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpLiquidateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpLiquidateEvent represents a LiquidateEvent event raised by the Lp contract.
type LpLiquidateEvent struct {
	Borrower   common.Address
	Liquidator common.Address
	Amount     *big.Int
	Markets    []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidateEvent is a free log retrieval operation binding the contract event 0x78066c353f5b8546c2005e7ed61b887d70c75a1b34b7c987cc6d3261c8cd3790.
//
// Solidity: event LiquidateEvent(address indexed borrower, address indexed liquidator, uint256 amount, address[] markets)
func (_Lp *LpFilterer) FilterLiquidateEvent(opts *bind.FilterOpts, borrower []common.Address, liquidator []common.Address) (*LpLiquidateEventIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "LiquidateEvent", borrowerRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &LpLiquidateEventIterator{contract: _Lp.contract, event: "LiquidateEvent", logs: logs, sub: sub}, nil
}

// WatchLiquidateEvent is a free log subscription operation binding the contract event 0x78066c353f5b8546c2005e7ed61b887d70c75a1b34b7c987cc6d3261c8cd3790.
//
// Solidity: event LiquidateEvent(address indexed borrower, address indexed liquidator, uint256 amount, address[] markets)
func (_Lp *LpFilterer) WatchLiquidateEvent(opts *bind.WatchOpts, sink chan<- *LpLiquidateEvent, borrower []common.Address, liquidator []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "LiquidateEvent", borrowerRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpLiquidateEvent)
				if err := _Lp.contract.UnpackLog(event, "LiquidateEvent", log); err != nil {
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

// ParseLiquidateEvent is a log parse operation binding the contract event 0x78066c353f5b8546c2005e7ed61b887d70c75a1b34b7c987cc6d3261c8cd3790.
//
// Solidity: event LiquidateEvent(address indexed borrower, address indexed liquidator, uint256 amount, address[] markets)
func (_Lp *LpFilterer) ParseLiquidateEvent(log types.Log) (*LpLiquidateEvent, error) {
	event := new(LpLiquidateEvent)
	if err := _Lp.contract.UnpackLog(event, "LiquidateEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lp contract.
type LpOwnershipTransferredIterator struct {
	Event *LpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpOwnershipTransferred)
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
		it.Event = new(LpOwnershipTransferred)
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
func (it *LpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpOwnershipTransferred represents a OwnershipTransferred event raised by the Lp contract.
type LpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lp *LpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LpOwnershipTransferredIterator{contract: _Lp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lp *LpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpOwnershipTransferred)
				if err := _Lp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lp *LpFilterer) ParseOwnershipTransferred(log types.Log) (*LpOwnershipTransferred, error) {
	event := new(LpOwnershipTransferred)
	if err := _Lp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Lp contract.
type LpPausedIterator struct {
	Event *LpPaused // Event containing the contract specifics and raw log

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
func (it *LpPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpPaused)
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
		it.Event = new(LpPaused)
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
func (it *LpPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpPaused represents a Paused event raised by the Lp contract.
type LpPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lp *LpFilterer) FilterPaused(opts *bind.FilterOpts) (*LpPausedIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LpPausedIterator{contract: _Lp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lp *LpFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LpPaused) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpPaused)
				if err := _Lp.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lp *LpFilterer) ParsePaused(log types.Log) (*LpPaused, error) {
	event := new(LpPaused)
	if err := _Lp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpRepayEventIterator is returned from FilterRepayEvent and is used to iterate over the raw logs and unpacked data for RepayEvent events raised by the Lp contract.
type LpRepayEventIterator struct {
	Event *LpRepayEvent // Event containing the contract specifics and raw log

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
func (it *LpRepayEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpRepayEvent)
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
		it.Event = new(LpRepayEvent)
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
func (it *LpRepayEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpRepayEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpRepayEvent represents a RepayEvent event raised by the Lp contract.
type LpRepayEvent struct {
	Borrower    common.Address
	RepayAmount *big.Int
	DebtAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRepayEvent is a free log retrieval operation binding the contract event 0x29833349d569b1cfb682b46a7cbe47f01ac8d741a9560aecbfff258e34d6d52c.
//
// Solidity: event RepayEvent(address indexed borrower, uint256 repayAmount, uint256 debtAmount)
func (_Lp *LpFilterer) FilterRepayEvent(opts *bind.FilterOpts, borrower []common.Address) (*LpRepayEventIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "RepayEvent", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LpRepayEventIterator{contract: _Lp.contract, event: "RepayEvent", logs: logs, sub: sub}, nil
}

// WatchRepayEvent is a free log subscription operation binding the contract event 0x29833349d569b1cfb682b46a7cbe47f01ac8d741a9560aecbfff258e34d6d52c.
//
// Solidity: event RepayEvent(address indexed borrower, uint256 repayAmount, uint256 debtAmount)
func (_Lp *LpFilterer) WatchRepayEvent(opts *bind.WatchOpts, sink chan<- *LpRepayEvent, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "RepayEvent", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpRepayEvent)
				if err := _Lp.contract.UnpackLog(event, "RepayEvent", log); err != nil {
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

// ParseRepayEvent is a log parse operation binding the contract event 0x29833349d569b1cfb682b46a7cbe47f01ac8d741a9560aecbfff258e34d6d52c.
//
// Solidity: event RepayEvent(address indexed borrower, uint256 repayAmount, uint256 debtAmount)
func (_Lp *LpFilterer) ParseRepayEvent(log types.Log) (*LpRepayEvent, error) {
	event := new(LpRepayEvent)
	if err := _Lp.contract.UnpackLog(event, "RepayEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpSupplyEventIterator is returned from FilterSupplyEvent and is used to iterate over the raw logs and unpacked data for SupplyEvent events raised by the Lp contract.
type LpSupplyEventIterator struct {
	Event *LpSupplyEvent // Event containing the contract specifics and raw log

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
func (it *LpSupplyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpSupplyEvent)
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
		it.Event = new(LpSupplyEvent)
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
func (it *LpSupplyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpSupplyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpSupplyEvent represents a SupplyEvent event raised by the Lp contract.
type LpSupplyEvent struct {
	Supplier     common.Address
	SupplyAmount *big.Int
	LptAmount    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSupplyEvent is a free log retrieval operation binding the contract event 0x10fd3ec7af698434b0a07de92e4a6802b21302af8aa4b168c6a342c8272fdb94.
//
// Solidity: event SupplyEvent(address indexed supplier, uint256 supplyAmount, uint256 lptAmount)
func (_Lp *LpFilterer) FilterSupplyEvent(opts *bind.FilterOpts, supplier []common.Address) (*LpSupplyEventIterator, error) {

	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "SupplyEvent", supplierRule)
	if err != nil {
		return nil, err
	}
	return &LpSupplyEventIterator{contract: _Lp.contract, event: "SupplyEvent", logs: logs, sub: sub}, nil
}

// WatchSupplyEvent is a free log subscription operation binding the contract event 0x10fd3ec7af698434b0a07de92e4a6802b21302af8aa4b168c6a342c8272fdb94.
//
// Solidity: event SupplyEvent(address indexed supplier, uint256 supplyAmount, uint256 lptAmount)
func (_Lp *LpFilterer) WatchSupplyEvent(opts *bind.WatchOpts, sink chan<- *LpSupplyEvent, supplier []common.Address) (event.Subscription, error) {

	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "SupplyEvent", supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpSupplyEvent)
				if err := _Lp.contract.UnpackLog(event, "SupplyEvent", log); err != nil {
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

// ParseSupplyEvent is a log parse operation binding the contract event 0x10fd3ec7af698434b0a07de92e4a6802b21302af8aa4b168c6a342c8272fdb94.
//
// Solidity: event SupplyEvent(address indexed supplier, uint256 supplyAmount, uint256 lptAmount)
func (_Lp *LpFilterer) ParseSupplyEvent(log types.Log) (*LpSupplyEvent, error) {
	event := new(LpSupplyEvent)
	if err := _Lp.contract.UnpackLog(event, "SupplyEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Lp contract.
type LpUnpausedIterator struct {
	Event *LpUnpaused // Event containing the contract specifics and raw log

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
func (it *LpUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpUnpaused)
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
		it.Event = new(LpUnpaused)
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
func (it *LpUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpUnpaused represents a Unpaused event raised by the Lp contract.
type LpUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lp *LpFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LpUnpausedIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LpUnpausedIterator{contract: _Lp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lp *LpFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LpUnpaused) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpUnpaused)
				if err := _Lp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lp *LpFilterer) ParseUnpaused(log types.Log) (*LpUnpaused, error) {
	event := new(LpUnpaused)
	if err := _Lp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpUpdateComptrollerIterator is returned from FilterUpdateComptroller and is used to iterate over the raw logs and unpacked data for UpdateComptroller events raised by the Lp contract.
type LpUpdateComptrollerIterator struct {
	Event *LpUpdateComptroller // Event containing the contract specifics and raw log

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
func (it *LpUpdateComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpUpdateComptroller)
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
		it.Event = new(LpUpdateComptroller)
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
func (it *LpUpdateComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpUpdateComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpUpdateComptroller represents a UpdateComptroller event raised by the Lp contract.
type LpUpdateComptroller struct {
	Comptroller common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateComptroller is a free log retrieval operation binding the contract event 0x62f36648f1b25bb9a634cf6f54b4adedd8de98ab504f738600d724160b251a1e.
//
// Solidity: event UpdateComptroller(address indexed comptroller)
func (_Lp *LpFilterer) FilterUpdateComptroller(opts *bind.FilterOpts, comptroller []common.Address) (*LpUpdateComptrollerIterator, error) {

	var comptrollerRule []interface{}
	for _, comptrollerItem := range comptroller {
		comptrollerRule = append(comptrollerRule, comptrollerItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "UpdateComptroller", comptrollerRule)
	if err != nil {
		return nil, err
	}
	return &LpUpdateComptrollerIterator{contract: _Lp.contract, event: "UpdateComptroller", logs: logs, sub: sub}, nil
}

// WatchUpdateComptroller is a free log subscription operation binding the contract event 0x62f36648f1b25bb9a634cf6f54b4adedd8de98ab504f738600d724160b251a1e.
//
// Solidity: event UpdateComptroller(address indexed comptroller)
func (_Lp *LpFilterer) WatchUpdateComptroller(opts *bind.WatchOpts, sink chan<- *LpUpdateComptroller, comptroller []common.Address) (event.Subscription, error) {

	var comptrollerRule []interface{}
	for _, comptrollerItem := range comptroller {
		comptrollerRule = append(comptrollerRule, comptrollerItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "UpdateComptroller", comptrollerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpUpdateComptroller)
				if err := _Lp.contract.UnpackLog(event, "UpdateComptroller", log); err != nil {
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

// ParseUpdateComptroller is a log parse operation binding the contract event 0x62f36648f1b25bb9a634cf6f54b4adedd8de98ab504f738600d724160b251a1e.
//
// Solidity: event UpdateComptroller(address indexed comptroller)
func (_Lp *LpFilterer) ParseUpdateComptroller(log types.Log) (*LpUpdateComptroller, error) {
	event := new(LpUpdateComptroller)
	if err := _Lp.contract.UnpackLog(event, "UpdateComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpUpdateInterestModelParametersIterator is returned from FilterUpdateInterestModelParameters and is used to iterate over the raw logs and unpacked data for UpdateInterestModelParameters events raised by the Lp contract.
type LpUpdateInterestModelParametersIterator struct {
	Event *LpUpdateInterestModelParameters // Event containing the contract specifics and raw log

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
func (it *LpUpdateInterestModelParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpUpdateInterestModelParameters)
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
		it.Event = new(LpUpdateInterestModelParameters)
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
func (it *LpUpdateInterestModelParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpUpdateInterestModelParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpUpdateInterestModelParameters represents a UpdateInterestModelParameters event raised by the Lp contract.
type LpUpdateInterestModelParameters struct {
	OptimalUtilizationRate *big.Int
	BaseBorrowRate         *big.Int
	Slope1                 *big.Int
	Slope2                 *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdateInterestModelParameters is a free log retrieval operation binding the contract event 0x48bd4a7cdba57ae423a63289ae15d09d309d4198ce2ee0eec95455b43737ddc2.
//
// Solidity: event UpdateInterestModelParameters(uint256 optimalUtilizationRate, uint256 baseBorrowRate, uint256 slope1, uint256 slope2)
func (_Lp *LpFilterer) FilterUpdateInterestModelParameters(opts *bind.FilterOpts) (*LpUpdateInterestModelParametersIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "UpdateInterestModelParameters")
	if err != nil {
		return nil, err
	}
	return &LpUpdateInterestModelParametersIterator{contract: _Lp.contract, event: "UpdateInterestModelParameters", logs: logs, sub: sub}, nil
}

// WatchUpdateInterestModelParameters is a free log subscription operation binding the contract event 0x48bd4a7cdba57ae423a63289ae15d09d309d4198ce2ee0eec95455b43737ddc2.
//
// Solidity: event UpdateInterestModelParameters(uint256 optimalUtilizationRate, uint256 baseBorrowRate, uint256 slope1, uint256 slope2)
func (_Lp *LpFilterer) WatchUpdateInterestModelParameters(opts *bind.WatchOpts, sink chan<- *LpUpdateInterestModelParameters) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "UpdateInterestModelParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpUpdateInterestModelParameters)
				if err := _Lp.contract.UnpackLog(event, "UpdateInterestModelParameters", log); err != nil {
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

// ParseUpdateInterestModelParameters is a log parse operation binding the contract event 0x48bd4a7cdba57ae423a63289ae15d09d309d4198ce2ee0eec95455b43737ddc2.
//
// Solidity: event UpdateInterestModelParameters(uint256 optimalUtilizationRate, uint256 baseBorrowRate, uint256 slope1, uint256 slope2)
func (_Lp *LpFilterer) ParseUpdateInterestModelParameters(log types.Log) (*LpUpdateInterestModelParameters, error) {
	event := new(LpUpdateInterestModelParameters)
	if err := _Lp.contract.UnpackLog(event, "UpdateInterestModelParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpUpdateLPTBaseValueIterator is returned from FilterUpdateLPTBaseValue and is used to iterate over the raw logs and unpacked data for UpdateLPTBaseValue events raised by the Lp contract.
type LpUpdateLPTBaseValueIterator struct {
	Event *LpUpdateLPTBaseValue // Event containing the contract specifics and raw log

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
func (it *LpUpdateLPTBaseValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpUpdateLPTBaseValue)
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
		it.Event = new(LpUpdateLPTBaseValue)
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
func (it *LpUpdateLPTBaseValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpUpdateLPTBaseValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpUpdateLPTBaseValue represents a UpdateLPTBaseValue event raised by the Lp contract.
type LpUpdateLPTBaseValue struct {
	LptBaseValue *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateLPTBaseValue is a free log retrieval operation binding the contract event 0x0cc1b6baf4e62cfdaf66022450ce4bf0549367f8c5a7c6543cfc6587bd477167.
//
// Solidity: event UpdateLPTBaseValue(uint256 lptBaseValue)
func (_Lp *LpFilterer) FilterUpdateLPTBaseValue(opts *bind.FilterOpts) (*LpUpdateLPTBaseValueIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "UpdateLPTBaseValue")
	if err != nil {
		return nil, err
	}
	return &LpUpdateLPTBaseValueIterator{contract: _Lp.contract, event: "UpdateLPTBaseValue", logs: logs, sub: sub}, nil
}

// WatchUpdateLPTBaseValue is a free log subscription operation binding the contract event 0x0cc1b6baf4e62cfdaf66022450ce4bf0549367f8c5a7c6543cfc6587bd477167.
//
// Solidity: event UpdateLPTBaseValue(uint256 lptBaseValue)
func (_Lp *LpFilterer) WatchUpdateLPTBaseValue(opts *bind.WatchOpts, sink chan<- *LpUpdateLPTBaseValue) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "UpdateLPTBaseValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpUpdateLPTBaseValue)
				if err := _Lp.contract.UnpackLog(event, "UpdateLPTBaseValue", log); err != nil {
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

// ParseUpdateLPTBaseValue is a log parse operation binding the contract event 0x0cc1b6baf4e62cfdaf66022450ce4bf0549367f8c5a7c6543cfc6587bd477167.
//
// Solidity: event UpdateLPTBaseValue(uint256 lptBaseValue)
func (_Lp *LpFilterer) ParseUpdateLPTBaseValue(log types.Log) (*LpUpdateLPTBaseValue, error) {
	event := new(LpUpdateLPTBaseValue)
	if err := _Lp.contract.UnpackLog(event, "UpdateLPTBaseValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpUpdateLiquidationPenaltyFactorIterator is returned from FilterUpdateLiquidationPenaltyFactor and is used to iterate over the raw logs and unpacked data for UpdateLiquidationPenaltyFactor events raised by the Lp contract.
type LpUpdateLiquidationPenaltyFactorIterator struct {
	Event *LpUpdateLiquidationPenaltyFactor // Event containing the contract specifics and raw log

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
func (it *LpUpdateLiquidationPenaltyFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpUpdateLiquidationPenaltyFactor)
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
		it.Event = new(LpUpdateLiquidationPenaltyFactor)
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
func (it *LpUpdateLiquidationPenaltyFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpUpdateLiquidationPenaltyFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpUpdateLiquidationPenaltyFactor represents a UpdateLiquidationPenaltyFactor event raised by the Lp contract.
type LpUpdateLiquidationPenaltyFactor struct {
	LiquidityPenaltyFactor *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdateLiquidationPenaltyFactor is a free log retrieval operation binding the contract event 0xe527ee260c6d1d4c5e502b90bb8eeb545d391676e62660db9a743842eba928f6.
//
// Solidity: event UpdateLiquidationPenaltyFactor(uint256 liquidityPenaltyFactor)
func (_Lp *LpFilterer) FilterUpdateLiquidationPenaltyFactor(opts *bind.FilterOpts) (*LpUpdateLiquidationPenaltyFactorIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "UpdateLiquidationPenaltyFactor")
	if err != nil {
		return nil, err
	}
	return &LpUpdateLiquidationPenaltyFactorIterator{contract: _Lp.contract, event: "UpdateLiquidationPenaltyFactor", logs: logs, sub: sub}, nil
}

// WatchUpdateLiquidationPenaltyFactor is a free log subscription operation binding the contract event 0xe527ee260c6d1d4c5e502b90bb8eeb545d391676e62660db9a743842eba928f6.
//
// Solidity: event UpdateLiquidationPenaltyFactor(uint256 liquidityPenaltyFactor)
func (_Lp *LpFilterer) WatchUpdateLiquidationPenaltyFactor(opts *bind.WatchOpts, sink chan<- *LpUpdateLiquidationPenaltyFactor) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "UpdateLiquidationPenaltyFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpUpdateLiquidationPenaltyFactor)
				if err := _Lp.contract.UnpackLog(event, "UpdateLiquidationPenaltyFactor", log); err != nil {
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

// ParseUpdateLiquidationPenaltyFactor is a log parse operation binding the contract event 0xe527ee260c6d1d4c5e502b90bb8eeb545d391676e62660db9a743842eba928f6.
//
// Solidity: event UpdateLiquidationPenaltyFactor(uint256 liquidityPenaltyFactor)
func (_Lp *LpFilterer) ParseUpdateLiquidationPenaltyFactor(log types.Log) (*LpUpdateLiquidationPenaltyFactor, error) {
	event := new(LpUpdateLiquidationPenaltyFactor)
	if err := _Lp.contract.UnpackLog(event, "UpdateLiquidationPenaltyFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpUpdateLiquidityPoolTokenIterator is returned from FilterUpdateLiquidityPoolToken and is used to iterate over the raw logs and unpacked data for UpdateLiquidityPoolToken events raised by the Lp contract.
type LpUpdateLiquidityPoolTokenIterator struct {
	Event *LpUpdateLiquidityPoolToken // Event containing the contract specifics and raw log

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
func (it *LpUpdateLiquidityPoolTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpUpdateLiquidityPoolToken)
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
		it.Event = new(LpUpdateLiquidityPoolToken)
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
func (it *LpUpdateLiquidityPoolTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpUpdateLiquidityPoolTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpUpdateLiquidityPoolToken represents a UpdateLiquidityPoolToken event raised by the Lp contract.
type LpUpdateLiquidityPoolToken struct {
	LiquidityPoolToken common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateLiquidityPoolToken is a free log retrieval operation binding the contract event 0x993f49c12126fcc4366266cd02e842eab87ef57962feaa40695e30e5a19d0d25.
//
// Solidity: event UpdateLiquidityPoolToken(address indexed liquidityPoolToken)
func (_Lp *LpFilterer) FilterUpdateLiquidityPoolToken(opts *bind.FilterOpts, liquidityPoolToken []common.Address) (*LpUpdateLiquidityPoolTokenIterator, error) {

	var liquidityPoolTokenRule []interface{}
	for _, liquidityPoolTokenItem := range liquidityPoolToken {
		liquidityPoolTokenRule = append(liquidityPoolTokenRule, liquidityPoolTokenItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "UpdateLiquidityPoolToken", liquidityPoolTokenRule)
	if err != nil {
		return nil, err
	}
	return &LpUpdateLiquidityPoolTokenIterator{contract: _Lp.contract, event: "UpdateLiquidityPoolToken", logs: logs, sub: sub}, nil
}

// WatchUpdateLiquidityPoolToken is a free log subscription operation binding the contract event 0x993f49c12126fcc4366266cd02e842eab87ef57962feaa40695e30e5a19d0d25.
//
// Solidity: event UpdateLiquidityPoolToken(address indexed liquidityPoolToken)
func (_Lp *LpFilterer) WatchUpdateLiquidityPoolToken(opts *bind.WatchOpts, sink chan<- *LpUpdateLiquidityPoolToken, liquidityPoolToken []common.Address) (event.Subscription, error) {

	var liquidityPoolTokenRule []interface{}
	for _, liquidityPoolTokenItem := range liquidityPoolToken {
		liquidityPoolTokenRule = append(liquidityPoolTokenRule, liquidityPoolTokenItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "UpdateLiquidityPoolToken", liquidityPoolTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpUpdateLiquidityPoolToken)
				if err := _Lp.contract.UnpackLog(event, "UpdateLiquidityPoolToken", log); err != nil {
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

// ParseUpdateLiquidityPoolToken is a log parse operation binding the contract event 0x993f49c12126fcc4366266cd02e842eab87ef57962feaa40695e30e5a19d0d25.
//
// Solidity: event UpdateLiquidityPoolToken(address indexed liquidityPoolToken)
func (_Lp *LpFilterer) ParseUpdateLiquidityPoolToken(log types.Log) (*LpUpdateLiquidityPoolToken, error) {
	event := new(LpUpdateLiquidityPoolToken)
	if err := _Lp.contract.UnpackLog(event, "UpdateLiquidityPoolToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpUpdateMinimumLoanValueIterator is returned from FilterUpdateMinimumLoanValue and is used to iterate over the raw logs and unpacked data for UpdateMinimumLoanValue events raised by the Lp contract.
type LpUpdateMinimumLoanValueIterator struct {
	Event *LpUpdateMinimumLoanValue // Event containing the contract specifics and raw log

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
func (it *LpUpdateMinimumLoanValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpUpdateMinimumLoanValue)
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
		it.Event = new(LpUpdateMinimumLoanValue)
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
func (it *LpUpdateMinimumLoanValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpUpdateMinimumLoanValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpUpdateMinimumLoanValue represents a UpdateMinimumLoanValue event raised by the Lp contract.
type LpUpdateMinimumLoanValue struct {
	MinimumLoanValue *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinimumLoanValue is a free log retrieval operation binding the contract event 0xc0ccc31ebdc5ed5034f73eab3c20a3e3a2a5796f4df71ae135d5923de5b8fae5.
//
// Solidity: event UpdateMinimumLoanValue(uint256 minimumLoanValue)
func (_Lp *LpFilterer) FilterUpdateMinimumLoanValue(opts *bind.FilterOpts) (*LpUpdateMinimumLoanValueIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "UpdateMinimumLoanValue")
	if err != nil {
		return nil, err
	}
	return &LpUpdateMinimumLoanValueIterator{contract: _Lp.contract, event: "UpdateMinimumLoanValue", logs: logs, sub: sub}, nil
}

// WatchUpdateMinimumLoanValue is a free log subscription operation binding the contract event 0xc0ccc31ebdc5ed5034f73eab3c20a3e3a2a5796f4df71ae135d5923de5b8fae5.
//
// Solidity: event UpdateMinimumLoanValue(uint256 minimumLoanValue)
func (_Lp *LpFilterer) WatchUpdateMinimumLoanValue(opts *bind.WatchOpts, sink chan<- *LpUpdateMinimumLoanValue) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "UpdateMinimumLoanValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpUpdateMinimumLoanValue)
				if err := _Lp.contract.UnpackLog(event, "UpdateMinimumLoanValue", log); err != nil {
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

// ParseUpdateMinimumLoanValue is a log parse operation binding the contract event 0xc0ccc31ebdc5ed5034f73eab3c20a3e3a2a5796f4df71ae135d5923de5b8fae5.
//
// Solidity: event UpdateMinimumLoanValue(uint256 minimumLoanValue)
func (_Lp *LpFilterer) ParseUpdateMinimumLoanValue(log types.Log) (*LpUpdateMinimumLoanValue, error) {
	event := new(LpUpdateMinimumLoanValue)
	if err := _Lp.contract.UnpackLog(event, "UpdateMinimumLoanValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpUpdateReserveFeeFactorIterator is returned from FilterUpdateReserveFeeFactor and is used to iterate over the raw logs and unpacked data for UpdateReserveFeeFactor events raised by the Lp contract.
type LpUpdateReserveFeeFactorIterator struct {
	Event *LpUpdateReserveFeeFactor // Event containing the contract specifics and raw log

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
func (it *LpUpdateReserveFeeFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpUpdateReserveFeeFactor)
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
		it.Event = new(LpUpdateReserveFeeFactor)
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
func (it *LpUpdateReserveFeeFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpUpdateReserveFeeFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpUpdateReserveFeeFactor represents a UpdateReserveFeeFactor event raised by the Lp contract.
type LpUpdateReserveFeeFactor struct {
	ReserveFactor *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateReserveFeeFactor is a free log retrieval operation binding the contract event 0x1a959f9ff930377a92a09403ab19c20dc25e192a0393ba229f8c606be991dc37.
//
// Solidity: event UpdateReserveFeeFactor(uint256 reserveFactor)
func (_Lp *LpFilterer) FilterUpdateReserveFeeFactor(opts *bind.FilterOpts) (*LpUpdateReserveFeeFactorIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "UpdateReserveFeeFactor")
	if err != nil {
		return nil, err
	}
	return &LpUpdateReserveFeeFactorIterator{contract: _Lp.contract, event: "UpdateReserveFeeFactor", logs: logs, sub: sub}, nil
}

// WatchUpdateReserveFeeFactor is a free log subscription operation binding the contract event 0x1a959f9ff930377a92a09403ab19c20dc25e192a0393ba229f8c606be991dc37.
//
// Solidity: event UpdateReserveFeeFactor(uint256 reserveFactor)
func (_Lp *LpFilterer) WatchUpdateReserveFeeFactor(opts *bind.WatchOpts, sink chan<- *LpUpdateReserveFeeFactor) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "UpdateReserveFeeFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpUpdateReserveFeeFactor)
				if err := _Lp.contract.UnpackLog(event, "UpdateReserveFeeFactor", log); err != nil {
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

// ParseUpdateReserveFeeFactor is a log parse operation binding the contract event 0x1a959f9ff930377a92a09403ab19c20dc25e192a0393ba229f8c606be991dc37.
//
// Solidity: event UpdateReserveFeeFactor(uint256 reserveFactor)
func (_Lp *LpFilterer) ParseUpdateReserveFeeFactor(log types.Log) (*LpUpdateReserveFeeFactor, error) {
	event := new(LpUpdateReserveFeeFactor)
	if err := _Lp.contract.UnpackLog(event, "UpdateReserveFeeFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpUpdateXtkFeeFactorIterator is returned from FilterUpdateXtkFeeFactor and is used to iterate over the raw logs and unpacked data for UpdateXtkFeeFactor events raised by the Lp contract.
type LpUpdateXtkFeeFactorIterator struct {
	Event *LpUpdateXtkFeeFactor // Event containing the contract specifics and raw log

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
func (it *LpUpdateXtkFeeFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpUpdateXtkFeeFactor)
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
		it.Event = new(LpUpdateXtkFeeFactor)
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
func (it *LpUpdateXtkFeeFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpUpdateXtkFeeFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpUpdateXtkFeeFactor represents a UpdateXtkFeeFactor event raised by the Lp contract.
type LpUpdateXtkFeeFactor struct {
	XtkFeeFactor *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateXtkFeeFactor is a free log retrieval operation binding the contract event 0x418b06b2826c4768673c13598e35b75b902c190edd9757ff532f3a256c7fd108.
//
// Solidity: event UpdateXtkFeeFactor(uint256 xtkFeeFactor)
func (_Lp *LpFilterer) FilterUpdateXtkFeeFactor(opts *bind.FilterOpts) (*LpUpdateXtkFeeFactorIterator, error) {

	logs, sub, err := _Lp.contract.FilterLogs(opts, "UpdateXtkFeeFactor")
	if err != nil {
		return nil, err
	}
	return &LpUpdateXtkFeeFactorIterator{contract: _Lp.contract, event: "UpdateXtkFeeFactor", logs: logs, sub: sub}, nil
}

// WatchUpdateXtkFeeFactor is a free log subscription operation binding the contract event 0x418b06b2826c4768673c13598e35b75b902c190edd9757ff532f3a256c7fd108.
//
// Solidity: event UpdateXtkFeeFactor(uint256 xtkFeeFactor)
func (_Lp *LpFilterer) WatchUpdateXtkFeeFactor(opts *bind.WatchOpts, sink chan<- *LpUpdateXtkFeeFactor) (event.Subscription, error) {

	logs, sub, err := _Lp.contract.WatchLogs(opts, "UpdateXtkFeeFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpUpdateXtkFeeFactor)
				if err := _Lp.contract.UnpackLog(event, "UpdateXtkFeeFactor", log); err != nil {
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

// ParseUpdateXtkFeeFactor is a log parse operation binding the contract event 0x418b06b2826c4768673c13598e35b75b902c190edd9757ff532f3a256c7fd108.
//
// Solidity: event UpdateXtkFeeFactor(uint256 xtkFeeFactor)
func (_Lp *LpFilterer) ParseUpdateXtkFeeFactor(log types.Log) (*LpUpdateXtkFeeFactor, error) {
	event := new(LpUpdateXtkFeeFactor)
	if err := _Lp.contract.UnpackLog(event, "UpdateXtkFeeFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpWithdrawEventIterator is returned from FilterWithdrawEvent and is used to iterate over the raw logs and unpacked data for WithdrawEvent events raised by the Lp contract.
type LpWithdrawEventIterator struct {
	Event *LpWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *LpWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpWithdrawEvent)
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
		it.Event = new(LpWithdrawEvent)
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
func (it *LpWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpWithdrawEvent represents a WithdrawEvent event raised by the Lp contract.
type LpWithdrawEvent struct {
	Supplier       common.Address
	LptAmount      *big.Int
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEvent is a free log retrieval operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: event WithdrawEvent(address indexed supplier, uint256 lptAmount, uint256 withdrawAmount)
func (_Lp *LpFilterer) FilterWithdrawEvent(opts *bind.FilterOpts, supplier []common.Address) (*LpWithdrawEventIterator, error) {

	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "WithdrawEvent", supplierRule)
	if err != nil {
		return nil, err
	}
	return &LpWithdrawEventIterator{contract: _Lp.contract, event: "WithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawEvent is a free log subscription operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: event WithdrawEvent(address indexed supplier, uint256 lptAmount, uint256 withdrawAmount)
func (_Lp *LpFilterer) WatchWithdrawEvent(opts *bind.WatchOpts, sink chan<- *LpWithdrawEvent, supplier []common.Address) (event.Subscription, error) {

	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "WithdrawEvent", supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpWithdrawEvent)
				if err := _Lp.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
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

// ParseWithdrawEvent is a log parse operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: event WithdrawEvent(address indexed supplier, uint256 lptAmount, uint256 withdrawAmount)
func (_Lp *LpFilterer) ParseWithdrawEvent(log types.Log) (*LpWithdrawEvent, error) {
	event := new(LpWithdrawEvent)
	if err := _Lp.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpWithdrawFeeIterator is returned from FilterWithdrawFee and is used to iterate over the raw logs and unpacked data for WithdrawFee events raised by the Lp contract.
type LpWithdrawFeeIterator struct {
	Event *LpWithdrawFee // Event containing the contract specifics and raw log

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
func (it *LpWithdrawFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpWithdrawFee)
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
		it.Event = new(LpWithdrawFee)
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
func (it *LpWithdrawFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpWithdrawFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpWithdrawFee represents a WithdrawFee event raised by the Lp contract.
type LpWithdrawFee struct {
	Recipient common.Address
	XtkEarns  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFee is a free log retrieval operation binding the contract event 0x66bf9186b00db666fc37aaffbb95a050c66e599e000c785c1dff0467d868f1b1.
//
// Solidity: event WithdrawFee(address indexed recipient, uint256 xtkEarns)
func (_Lp *LpFilterer) FilterWithdrawFee(opts *bind.FilterOpts, recipient []common.Address) (*LpWithdrawFeeIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "WithdrawFee", recipientRule)
	if err != nil {
		return nil, err
	}
	return &LpWithdrawFeeIterator{contract: _Lp.contract, event: "WithdrawFee", logs: logs, sub: sub}, nil
}

// WatchWithdrawFee is a free log subscription operation binding the contract event 0x66bf9186b00db666fc37aaffbb95a050c66e599e000c785c1dff0467d868f1b1.
//
// Solidity: event WithdrawFee(address indexed recipient, uint256 xtkEarns)
func (_Lp *LpFilterer) WatchWithdrawFee(opts *bind.WatchOpts, sink chan<- *LpWithdrawFee, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "WithdrawFee", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpWithdrawFee)
				if err := _Lp.contract.UnpackLog(event, "WithdrawFee", log); err != nil {
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

// ParseWithdrawFee is a log parse operation binding the contract event 0x66bf9186b00db666fc37aaffbb95a050c66e599e000c785c1dff0467d868f1b1.
//
// Solidity: event WithdrawFee(address indexed recipient, uint256 xtkEarns)
func (_Lp *LpFilterer) ParseWithdrawFee(log types.Log) (*LpWithdrawFee, error) {
	event := new(LpWithdrawFee)
	if err := _Lp.contract.UnpackLog(event, "WithdrawFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
