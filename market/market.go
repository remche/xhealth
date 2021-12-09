// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package market

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

// MarketMetaData contains all meta data concerning the Market contract.
var MarketMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Collateralize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralCap\",\"type\":\"uint256\"}],\"name\":\"UpdateCollateralCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralFactor\",\"type\":\"uint256\"}],\"name\":\"UpdateCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"UpdateCollateralizationActive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"comptroller\",\"type\":\"address\"}],\"name\":\"UpdateComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"assetPriceAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blockLockExempt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"borrowingLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"collateralize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lockAddress\",\"type\":\"address\"}],\"name\":\"exemptFromBlockLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetPriceAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_collateralFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_collateralCap\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastLockedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lockAddress\",\"type\":\"address\"}],\"name\":\"removeBlockLockExemption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendCollateralToLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collateralCap\",\"type\":\"uint256\"}],\"name\":\"setCollateralCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collateralFactor\",\"type\":\"uint256\"}],\"name\":\"setCollateralFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setCollateralizationActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_comptroller\",\"type\":\"address\"}],\"name\":\"setComptroller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingAssetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketMetaData.ABI instead.
var MarketABI = MarketMetaData.ABI

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// AssetPriceAddress is a free data retrieval call binding the contract method 0xd8d7b69e.
//
// Solidity: function assetPriceAddress() view returns(address)
func (_Market *MarketCaller) AssetPriceAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "assetPriceAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetPriceAddress is a free data retrieval call binding the contract method 0xd8d7b69e.
//
// Solidity: function assetPriceAddress() view returns(address)
func (_Market *MarketSession) AssetPriceAddress() (common.Address, error) {
	return _Market.Contract.AssetPriceAddress(&_Market.CallOpts)
}

// AssetPriceAddress is a free data retrieval call binding the contract method 0xd8d7b69e.
//
// Solidity: function assetPriceAddress() view returns(address)
func (_Market *MarketCallerSession) AssetPriceAddress() (common.Address, error) {
	return _Market.Contract.AssetPriceAddress(&_Market.CallOpts)
}

// BlockLockExempt is a free data retrieval call binding the contract method 0xfff3011b.
//
// Solidity: function blockLockExempt(address ) view returns(bool)
func (_Market *MarketCaller) BlockLockExempt(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "blockLockExempt", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlockLockExempt is a free data retrieval call binding the contract method 0xfff3011b.
//
// Solidity: function blockLockExempt(address ) view returns(bool)
func (_Market *MarketSession) BlockLockExempt(arg0 common.Address) (bool, error) {
	return _Market.Contract.BlockLockExempt(&_Market.CallOpts, arg0)
}

// BlockLockExempt is a free data retrieval call binding the contract method 0xfff3011b.
//
// Solidity: function blockLockExempt(address ) view returns(bool)
func (_Market *MarketCallerSession) BlockLockExempt(arg0 common.Address) (bool, error) {
	return _Market.Contract.BlockLockExempt(&_Market.CallOpts, arg0)
}

// BorrowingLimit is a free data retrieval call binding the contract method 0xa6a58df4.
//
// Solidity: function borrowingLimit(address _borrower) view returns(uint256)
func (_Market *MarketCaller) BorrowingLimit(opts *bind.CallOpts, _borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "borrowingLimit", _borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowingLimit is a free data retrieval call binding the contract method 0xa6a58df4.
//
// Solidity: function borrowingLimit(address _borrower) view returns(uint256)
func (_Market *MarketSession) BorrowingLimit(_borrower common.Address) (*big.Int, error) {
	return _Market.Contract.BorrowingLimit(&_Market.CallOpts, _borrower)
}

// BorrowingLimit is a free data retrieval call binding the contract method 0xa6a58df4.
//
// Solidity: function borrowingLimit(address _borrower) view returns(uint256)
func (_Market *MarketCallerSession) BorrowingLimit(_borrower common.Address) (*big.Int, error) {
	return _Market.Contract.BorrowingLimit(&_Market.CallOpts, _borrower)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address _borrower) view returns(uint256)
func (_Market *MarketCaller) Collateral(opts *bind.CallOpts, _borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "collateral", _borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address _borrower) view returns(uint256)
func (_Market *MarketSession) Collateral(_borrower common.Address) (*big.Int, error) {
	return _Market.Contract.Collateral(&_Market.CallOpts, _borrower)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address _borrower) view returns(uint256)
func (_Market *MarketCallerSession) Collateral(_borrower common.Address) (*big.Int, error) {
	return _Market.Contract.Collateral(&_Market.CallOpts, _borrower)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Market *MarketCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Market *MarketSession) Comptroller() (common.Address, error) {
	return _Market.Contract.Comptroller(&_Market.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Market *MarketCallerSession) Comptroller() (common.Address, error) {
	return _Market.Contract.Comptroller(&_Market.CallOpts)
}

// GetCollateralCap is a free data retrieval call binding the contract method 0xca7d6cf3.
//
// Solidity: function getCollateralCap() view returns(uint256)
func (_Market *MarketCaller) GetCollateralCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getCollateralCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralCap is a free data retrieval call binding the contract method 0xca7d6cf3.
//
// Solidity: function getCollateralCap() view returns(uint256)
func (_Market *MarketSession) GetCollateralCap() (*big.Int, error) {
	return _Market.Contract.GetCollateralCap(&_Market.CallOpts)
}

// GetCollateralCap is a free data retrieval call binding the contract method 0xca7d6cf3.
//
// Solidity: function getCollateralCap() view returns(uint256)
func (_Market *MarketCallerSession) GetCollateralCap() (*big.Int, error) {
	return _Market.Contract.GetCollateralCap(&_Market.CallOpts)
}

// GetCollateralFactor is a free data retrieval call binding the contract method 0x35392b71.
//
// Solidity: function getCollateralFactor() view returns(uint256)
func (_Market *MarketCaller) GetCollateralFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getCollateralFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralFactor is a free data retrieval call binding the contract method 0x35392b71.
//
// Solidity: function getCollateralFactor() view returns(uint256)
func (_Market *MarketSession) GetCollateralFactor() (*big.Int, error) {
	return _Market.Contract.GetCollateralFactor(&_Market.CallOpts)
}

// GetCollateralFactor is a free data retrieval call binding the contract method 0x35392b71.
//
// Solidity: function getCollateralFactor() view returns(uint256)
func (_Market *MarketCallerSession) GetCollateralFactor() (*big.Int, error) {
	return _Market.Contract.GetCollateralFactor(&_Market.CallOpts)
}

// LastLockedBlock is a free data retrieval call binding the contract method 0x9f3e8b34.
//
// Solidity: function lastLockedBlock(address ) view returns(uint256)
func (_Market *MarketCaller) LastLockedBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "lastLockedBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastLockedBlock is a free data retrieval call binding the contract method 0x9f3e8b34.
//
// Solidity: function lastLockedBlock(address ) view returns(uint256)
func (_Market *MarketSession) LastLockedBlock(arg0 common.Address) (*big.Int, error) {
	return _Market.Contract.LastLockedBlock(&_Market.CallOpts, arg0)
}

// LastLockedBlock is a free data retrieval call binding the contract method 0x9f3e8b34.
//
// Solidity: function lastLockedBlock(address ) view returns(uint256)
func (_Market *MarketCallerSession) LastLockedBlock(arg0 common.Address) (*big.Int, error) {
	return _Market.Contract.LastLockedBlock(&_Market.CallOpts, arg0)
}

// MarketActive is a free data retrieval call binding the contract method 0xaaa31151.
//
// Solidity: function marketActive() view returns(bool)
func (_Market *MarketCaller) MarketActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "marketActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MarketActive is a free data retrieval call binding the contract method 0xaaa31151.
//
// Solidity: function marketActive() view returns(bool)
func (_Market *MarketSession) MarketActive() (bool, error) {
	return _Market.Contract.MarketActive(&_Market.CallOpts)
}

// MarketActive is a free data retrieval call binding the contract method 0xaaa31151.
//
// Solidity: function marketActive() view returns(bool)
func (_Market *MarketCallerSession) MarketActive() (bool, error) {
	return _Market.Contract.MarketActive(&_Market.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Market *MarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Market *MarketSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Market *MarketCallerSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Market *MarketCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Market *MarketSession) Paused() (bool, error) {
	return _Market.Contract.Paused(&_Market.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Market *MarketCallerSession) Paused() (bool, error) {
	return _Market.Contract.Paused(&_Market.CallOpts)
}

// UnderlyingAssetAddress is a free data retrieval call binding the contract method 0x89d1a0fc.
//
// Solidity: function underlyingAssetAddress() view returns(address)
func (_Market *MarketCaller) UnderlyingAssetAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "underlyingAssetAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingAssetAddress is a free data retrieval call binding the contract method 0x89d1a0fc.
//
// Solidity: function underlyingAssetAddress() view returns(address)
func (_Market *MarketSession) UnderlyingAssetAddress() (common.Address, error) {
	return _Market.Contract.UnderlyingAssetAddress(&_Market.CallOpts)
}

// UnderlyingAssetAddress is a free data retrieval call binding the contract method 0x89d1a0fc.
//
// Solidity: function underlyingAssetAddress() view returns(address)
func (_Market *MarketCallerSession) UnderlyingAssetAddress() (common.Address, error) {
	return _Market.Contract.UnderlyingAssetAddress(&_Market.CallOpts)
}

// Collateralize is a paid mutator transaction binding the contract method 0x9f06aa08.
//
// Solidity: function collateralize(uint256 _amount) returns()
func (_Market *MarketTransactor) Collateralize(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "collateralize", _amount)
}

// Collateralize is a paid mutator transaction binding the contract method 0x9f06aa08.
//
// Solidity: function collateralize(uint256 _amount) returns()
func (_Market *MarketSession) Collateralize(_amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Collateralize(&_Market.TransactOpts, _amount)
}

// Collateralize is a paid mutator transaction binding the contract method 0x9f06aa08.
//
// Solidity: function collateralize(uint256 _amount) returns()
func (_Market *MarketTransactorSession) Collateralize(_amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Collateralize(&_Market.TransactOpts, _amount)
}

// ExemptFromBlockLock is a paid mutator transaction binding the contract method 0xd896cf19.
//
// Solidity: function exemptFromBlockLock(address lockAddress) returns()
func (_Market *MarketTransactor) ExemptFromBlockLock(opts *bind.TransactOpts, lockAddress common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "exemptFromBlockLock", lockAddress)
}

// ExemptFromBlockLock is a paid mutator transaction binding the contract method 0xd896cf19.
//
// Solidity: function exemptFromBlockLock(address lockAddress) returns()
func (_Market *MarketSession) ExemptFromBlockLock(lockAddress common.Address) (*types.Transaction, error) {
	return _Market.Contract.ExemptFromBlockLock(&_Market.TransactOpts, lockAddress)
}

// ExemptFromBlockLock is a paid mutator transaction binding the contract method 0xd896cf19.
//
// Solidity: function exemptFromBlockLock(address lockAddress) returns()
func (_Market *MarketTransactorSession) ExemptFromBlockLock(lockAddress common.Address) (*types.Transaction, error) {
	return _Market.Contract.ExemptFromBlockLock(&_Market.TransactOpts, lockAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _assetPriceAddress, uint256 _collateralFactor, uint256 _collateralCap) returns()
func (_Market *MarketTransactor) Initialize(opts *bind.TransactOpts, _assetPriceAddress common.Address, _collateralFactor *big.Int, _collateralCap *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "initialize", _assetPriceAddress, _collateralFactor, _collateralCap)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _assetPriceAddress, uint256 _collateralFactor, uint256 _collateralCap) returns()
func (_Market *MarketSession) Initialize(_assetPriceAddress common.Address, _collateralFactor *big.Int, _collateralCap *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Initialize(&_Market.TransactOpts, _assetPriceAddress, _collateralFactor, _collateralCap)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _assetPriceAddress, uint256 _collateralFactor, uint256 _collateralCap) returns()
func (_Market *MarketTransactorSession) Initialize(_assetPriceAddress common.Address, _collateralFactor *big.Int, _collateralCap *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Initialize(&_Market.TransactOpts, _assetPriceAddress, _collateralFactor, _collateralCap)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_Market *MarketTransactor) PauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "pauseContract")
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_Market *MarketSession) PauseContract() (*types.Transaction, error) {
	return _Market.Contract.PauseContract(&_Market.TransactOpts)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_Market *MarketTransactorSession) PauseContract() (*types.Transaction, error) {
	return _Market.Contract.PauseContract(&_Market.TransactOpts)
}

// RemoveBlockLockExemption is a paid mutator transaction binding the contract method 0x1de34d39.
//
// Solidity: function removeBlockLockExemption(address lockAddress) returns()
func (_Market *MarketTransactor) RemoveBlockLockExemption(opts *bind.TransactOpts, lockAddress common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "removeBlockLockExemption", lockAddress)
}

// RemoveBlockLockExemption is a paid mutator transaction binding the contract method 0x1de34d39.
//
// Solidity: function removeBlockLockExemption(address lockAddress) returns()
func (_Market *MarketSession) RemoveBlockLockExemption(lockAddress common.Address) (*types.Transaction, error) {
	return _Market.Contract.RemoveBlockLockExemption(&_Market.TransactOpts, lockAddress)
}

// RemoveBlockLockExemption is a paid mutator transaction binding the contract method 0x1de34d39.
//
// Solidity: function removeBlockLockExemption(address lockAddress) returns()
func (_Market *MarketTransactorSession) RemoveBlockLockExemption(lockAddress common.Address) (*types.Transaction, error) {
	return _Market.Contract.RemoveBlockLockExemption(&_Market.TransactOpts, lockAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Market *MarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Market *MarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _Market.Contract.RenounceOwnership(&_Market.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Market *MarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Market.Contract.RenounceOwnership(&_Market.TransactOpts)
}

// SendCollateralToLiquidator is a paid mutator transaction binding the contract method 0x6c6af7a0.
//
// Solidity: function sendCollateralToLiquidator(address _liquidator, address _borrower, uint256 _amount) returns()
func (_Market *MarketTransactor) SendCollateralToLiquidator(opts *bind.TransactOpts, _liquidator common.Address, _borrower common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "sendCollateralToLiquidator", _liquidator, _borrower, _amount)
}

// SendCollateralToLiquidator is a paid mutator transaction binding the contract method 0x6c6af7a0.
//
// Solidity: function sendCollateralToLiquidator(address _liquidator, address _borrower, uint256 _amount) returns()
func (_Market *MarketSession) SendCollateralToLiquidator(_liquidator common.Address, _borrower common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SendCollateralToLiquidator(&_Market.TransactOpts, _liquidator, _borrower, _amount)
}

// SendCollateralToLiquidator is a paid mutator transaction binding the contract method 0x6c6af7a0.
//
// Solidity: function sendCollateralToLiquidator(address _liquidator, address _borrower, uint256 _amount) returns()
func (_Market *MarketTransactorSession) SendCollateralToLiquidator(_liquidator common.Address, _borrower common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SendCollateralToLiquidator(&_Market.TransactOpts, _liquidator, _borrower, _amount)
}

// SetCollateralCap is a paid mutator transaction binding the contract method 0xe6d15440.
//
// Solidity: function setCollateralCap(uint256 _collateralCap) returns()
func (_Market *MarketTransactor) SetCollateralCap(opts *bind.TransactOpts, _collateralCap *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setCollateralCap", _collateralCap)
}

// SetCollateralCap is a paid mutator transaction binding the contract method 0xe6d15440.
//
// Solidity: function setCollateralCap(uint256 _collateralCap) returns()
func (_Market *MarketSession) SetCollateralCap(_collateralCap *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SetCollateralCap(&_Market.TransactOpts, _collateralCap)
}

// SetCollateralCap is a paid mutator transaction binding the contract method 0xe6d15440.
//
// Solidity: function setCollateralCap(uint256 _collateralCap) returns()
func (_Market *MarketTransactorSession) SetCollateralCap(_collateralCap *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SetCollateralCap(&_Market.TransactOpts, _collateralCap)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xb5c22157.
//
// Solidity: function setCollateralFactor(uint256 _collateralFactor) returns()
func (_Market *MarketTransactor) SetCollateralFactor(opts *bind.TransactOpts, _collateralFactor *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setCollateralFactor", _collateralFactor)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xb5c22157.
//
// Solidity: function setCollateralFactor(uint256 _collateralFactor) returns()
func (_Market *MarketSession) SetCollateralFactor(_collateralFactor *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SetCollateralFactor(&_Market.TransactOpts, _collateralFactor)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xb5c22157.
//
// Solidity: function setCollateralFactor(uint256 _collateralFactor) returns()
func (_Market *MarketTransactorSession) SetCollateralFactor(_collateralFactor *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SetCollateralFactor(&_Market.TransactOpts, _collateralFactor)
}

// SetCollateralizationActive is a paid mutator transaction binding the contract method 0xebe3ce08.
//
// Solidity: function setCollateralizationActive(bool _active) returns()
func (_Market *MarketTransactor) SetCollateralizationActive(opts *bind.TransactOpts, _active bool) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setCollateralizationActive", _active)
}

// SetCollateralizationActive is a paid mutator transaction binding the contract method 0xebe3ce08.
//
// Solidity: function setCollateralizationActive(bool _active) returns()
func (_Market *MarketSession) SetCollateralizationActive(_active bool) (*types.Transaction, error) {
	return _Market.Contract.SetCollateralizationActive(&_Market.TransactOpts, _active)
}

// SetCollateralizationActive is a paid mutator transaction binding the contract method 0xebe3ce08.
//
// Solidity: function setCollateralizationActive(bool _active) returns()
func (_Market *MarketTransactorSession) SetCollateralizationActive(_active bool) (*types.Transaction, error) {
	return _Market.Contract.SetCollateralizationActive(&_Market.TransactOpts, _active)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x8bad38dd.
//
// Solidity: function setComptroller(address _comptroller) returns()
func (_Market *MarketTransactor) SetComptroller(opts *bind.TransactOpts, _comptroller common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setComptroller", _comptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x8bad38dd.
//
// Solidity: function setComptroller(address _comptroller) returns()
func (_Market *MarketSession) SetComptroller(_comptroller common.Address) (*types.Transaction, error) {
	return _Market.Contract.SetComptroller(&_Market.TransactOpts, _comptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x8bad38dd.
//
// Solidity: function setComptroller(address _comptroller) returns()
func (_Market *MarketTransactorSession) SetComptroller(_comptroller common.Address) (*types.Transaction, error) {
	return _Market.Contract.SetComptroller(&_Market.TransactOpts, _comptroller)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Market *MarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Market *MarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Market.Contract.TransferOwnership(&_Market.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Market *MarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Market.Contract.TransferOwnership(&_Market.TransactOpts, newOwner)
}

// UnpauseContract is a paid mutator transaction binding the contract method 0xb33712c5.
//
// Solidity: function unpauseContract() returns()
func (_Market *MarketTransactor) UnpauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "unpauseContract")
}

// UnpauseContract is a paid mutator transaction binding the contract method 0xb33712c5.
//
// Solidity: function unpauseContract() returns()
func (_Market *MarketSession) UnpauseContract() (*types.Transaction, error) {
	return _Market.Contract.UnpauseContract(&_Market.TransactOpts)
}

// UnpauseContract is a paid mutator transaction binding the contract method 0xb33712c5.
//
// Solidity: function unpauseContract() returns()
func (_Market *MarketTransactorSession) UnpauseContract() (*types.Transaction, error) {
	return _Market.Contract.UnpauseContract(&_Market.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Market *MarketTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Market *MarketSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Withdraw(&_Market.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Market *MarketTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Withdraw(&_Market.TransactOpts, _amount)
}

// MarketCollateralizeIterator is returned from FilterCollateralize and is used to iterate over the raw logs and unpacked data for Collateralize events raised by the Market contract.
type MarketCollateralizeIterator struct {
	Event *MarketCollateralize // Event containing the contract specifics and raw log

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
func (it *MarketCollateralizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCollateralize)
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
		it.Event = new(MarketCollateralize)
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
func (it *MarketCollateralizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCollateralizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCollateralize represents a Collateralize event raised by the Market contract.
type MarketCollateralize struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollateralize is a free log retrieval operation binding the contract event 0x94c9d5464fed48366596c718219e5fbd3c0c61c1a21391bf239075e9de5a727a.
//
// Solidity: event Collateralize(address indexed user, uint256 amount)
func (_Market *MarketFilterer) FilterCollateralize(opts *bind.FilterOpts, user []common.Address) (*MarketCollateralizeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "Collateralize", userRule)
	if err != nil {
		return nil, err
	}
	return &MarketCollateralizeIterator{contract: _Market.contract, event: "Collateralize", logs: logs, sub: sub}, nil
}

// WatchCollateralize is a free log subscription operation binding the contract event 0x94c9d5464fed48366596c718219e5fbd3c0c61c1a21391bf239075e9de5a727a.
//
// Solidity: event Collateralize(address indexed user, uint256 amount)
func (_Market *MarketFilterer) WatchCollateralize(opts *bind.WatchOpts, sink chan<- *MarketCollateralize, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "Collateralize", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCollateralize)
				if err := _Market.contract.UnpackLog(event, "Collateralize", log); err != nil {
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

// ParseCollateralize is a log parse operation binding the contract event 0x94c9d5464fed48366596c718219e5fbd3c0c61c1a21391bf239075e9de5a727a.
//
// Solidity: event Collateralize(address indexed user, uint256 amount)
func (_Market *MarketFilterer) ParseCollateralize(log types.Log) (*MarketCollateralize, error) {
	event := new(MarketCollateralize)
	if err := _Market.contract.UnpackLog(event, "Collateralize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Market contract.
type MarketOwnershipTransferredIterator struct {
	Event *MarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketOwnershipTransferred)
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
		it.Event = new(MarketOwnershipTransferred)
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
func (it *MarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketOwnershipTransferred represents a OwnershipTransferred event raised by the Market contract.
type MarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Market *MarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketOwnershipTransferredIterator{contract: _Market.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Market *MarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketOwnershipTransferred)
				if err := _Market.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Market *MarketFilterer) ParseOwnershipTransferred(log types.Log) (*MarketOwnershipTransferred, error) {
	event := new(MarketOwnershipTransferred)
	if err := _Market.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Market contract.
type MarketPausedIterator struct {
	Event *MarketPaused // Event containing the contract specifics and raw log

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
func (it *MarketPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPaused)
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
		it.Event = new(MarketPaused)
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
func (it *MarketPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPaused represents a Paused event raised by the Market contract.
type MarketPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Market *MarketFilterer) FilterPaused(opts *bind.FilterOpts) (*MarketPausedIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MarketPausedIterator{contract: _Market.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Market *MarketFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MarketPaused) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPaused)
				if err := _Market.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Market *MarketFilterer) ParsePaused(log types.Log) (*MarketPaused, error) {
	event := new(MarketPaused)
	if err := _Market.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Market contract.
type MarketUnpausedIterator struct {
	Event *MarketUnpaused // Event containing the contract specifics and raw log

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
func (it *MarketUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketUnpaused)
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
		it.Event = new(MarketUnpaused)
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
func (it *MarketUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketUnpaused represents a Unpaused event raised by the Market contract.
type MarketUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Market *MarketFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MarketUnpausedIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MarketUnpausedIterator{contract: _Market.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Market *MarketFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MarketUnpaused) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketUnpaused)
				if err := _Market.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Market *MarketFilterer) ParseUnpaused(log types.Log) (*MarketUnpaused, error) {
	event := new(MarketUnpaused)
	if err := _Market.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketUpdateCollateralCapIterator is returned from FilterUpdateCollateralCap and is used to iterate over the raw logs and unpacked data for UpdateCollateralCap events raised by the Market contract.
type MarketUpdateCollateralCapIterator struct {
	Event *MarketUpdateCollateralCap // Event containing the contract specifics and raw log

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
func (it *MarketUpdateCollateralCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketUpdateCollateralCap)
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
		it.Event = new(MarketUpdateCollateralCap)
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
func (it *MarketUpdateCollateralCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketUpdateCollateralCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketUpdateCollateralCap represents a UpdateCollateralCap event raised by the Market contract.
type MarketUpdateCollateralCap struct {
	CollateralCap *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateCollateralCap is a free log retrieval operation binding the contract event 0x573a911336a4ad1043273351cf4ec6cf5b043d0dde3dc8b6b3a100c19a95cbb6.
//
// Solidity: event UpdateCollateralCap(uint256 collateralCap)
func (_Market *MarketFilterer) FilterUpdateCollateralCap(opts *bind.FilterOpts) (*MarketUpdateCollateralCapIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "UpdateCollateralCap")
	if err != nil {
		return nil, err
	}
	return &MarketUpdateCollateralCapIterator{contract: _Market.contract, event: "UpdateCollateralCap", logs: logs, sub: sub}, nil
}

// WatchUpdateCollateralCap is a free log subscription operation binding the contract event 0x573a911336a4ad1043273351cf4ec6cf5b043d0dde3dc8b6b3a100c19a95cbb6.
//
// Solidity: event UpdateCollateralCap(uint256 collateralCap)
func (_Market *MarketFilterer) WatchUpdateCollateralCap(opts *bind.WatchOpts, sink chan<- *MarketUpdateCollateralCap) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "UpdateCollateralCap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketUpdateCollateralCap)
				if err := _Market.contract.UnpackLog(event, "UpdateCollateralCap", log); err != nil {
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

// ParseUpdateCollateralCap is a log parse operation binding the contract event 0x573a911336a4ad1043273351cf4ec6cf5b043d0dde3dc8b6b3a100c19a95cbb6.
//
// Solidity: event UpdateCollateralCap(uint256 collateralCap)
func (_Market *MarketFilterer) ParseUpdateCollateralCap(log types.Log) (*MarketUpdateCollateralCap, error) {
	event := new(MarketUpdateCollateralCap)
	if err := _Market.contract.UnpackLog(event, "UpdateCollateralCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketUpdateCollateralFactorIterator is returned from FilterUpdateCollateralFactor and is used to iterate over the raw logs and unpacked data for UpdateCollateralFactor events raised by the Market contract.
type MarketUpdateCollateralFactorIterator struct {
	Event *MarketUpdateCollateralFactor // Event containing the contract specifics and raw log

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
func (it *MarketUpdateCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketUpdateCollateralFactor)
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
		it.Event = new(MarketUpdateCollateralFactor)
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
func (it *MarketUpdateCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketUpdateCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketUpdateCollateralFactor represents a UpdateCollateralFactor event raised by the Market contract.
type MarketUpdateCollateralFactor struct {
	CollateralFactor *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateCollateralFactor is a free log retrieval operation binding the contract event 0xf5fde6a0fde6811c2a4cd8a9d75b6f5ed6b63f8e83d307c5b75be120673d0b15.
//
// Solidity: event UpdateCollateralFactor(uint256 collateralFactor)
func (_Market *MarketFilterer) FilterUpdateCollateralFactor(opts *bind.FilterOpts) (*MarketUpdateCollateralFactorIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "UpdateCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &MarketUpdateCollateralFactorIterator{contract: _Market.contract, event: "UpdateCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchUpdateCollateralFactor is a free log subscription operation binding the contract event 0xf5fde6a0fde6811c2a4cd8a9d75b6f5ed6b63f8e83d307c5b75be120673d0b15.
//
// Solidity: event UpdateCollateralFactor(uint256 collateralFactor)
func (_Market *MarketFilterer) WatchUpdateCollateralFactor(opts *bind.WatchOpts, sink chan<- *MarketUpdateCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "UpdateCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketUpdateCollateralFactor)
				if err := _Market.contract.UnpackLog(event, "UpdateCollateralFactor", log); err != nil {
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

// ParseUpdateCollateralFactor is a log parse operation binding the contract event 0xf5fde6a0fde6811c2a4cd8a9d75b6f5ed6b63f8e83d307c5b75be120673d0b15.
//
// Solidity: event UpdateCollateralFactor(uint256 collateralFactor)
func (_Market *MarketFilterer) ParseUpdateCollateralFactor(log types.Log) (*MarketUpdateCollateralFactor, error) {
	event := new(MarketUpdateCollateralFactor)
	if err := _Market.contract.UnpackLog(event, "UpdateCollateralFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketUpdateCollateralizationActiveIterator is returned from FilterUpdateCollateralizationActive and is used to iterate over the raw logs and unpacked data for UpdateCollateralizationActive events raised by the Market contract.
type MarketUpdateCollateralizationActiveIterator struct {
	Event *MarketUpdateCollateralizationActive // Event containing the contract specifics and raw log

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
func (it *MarketUpdateCollateralizationActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketUpdateCollateralizationActive)
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
		it.Event = new(MarketUpdateCollateralizationActive)
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
func (it *MarketUpdateCollateralizationActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketUpdateCollateralizationActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketUpdateCollateralizationActive represents a UpdateCollateralizationActive event raised by the Market contract.
type MarketUpdateCollateralizationActive struct {
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateCollateralizationActive is a free log retrieval operation binding the contract event 0x9a75057ae268e789c13fb43bbc257401731cb96df0eea6be3ff4a6bf416b415b.
//
// Solidity: event UpdateCollateralizationActive(bool active)
func (_Market *MarketFilterer) FilterUpdateCollateralizationActive(opts *bind.FilterOpts) (*MarketUpdateCollateralizationActiveIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "UpdateCollateralizationActive")
	if err != nil {
		return nil, err
	}
	return &MarketUpdateCollateralizationActiveIterator{contract: _Market.contract, event: "UpdateCollateralizationActive", logs: logs, sub: sub}, nil
}

// WatchUpdateCollateralizationActive is a free log subscription operation binding the contract event 0x9a75057ae268e789c13fb43bbc257401731cb96df0eea6be3ff4a6bf416b415b.
//
// Solidity: event UpdateCollateralizationActive(bool active)
func (_Market *MarketFilterer) WatchUpdateCollateralizationActive(opts *bind.WatchOpts, sink chan<- *MarketUpdateCollateralizationActive) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "UpdateCollateralizationActive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketUpdateCollateralizationActive)
				if err := _Market.contract.UnpackLog(event, "UpdateCollateralizationActive", log); err != nil {
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

// ParseUpdateCollateralizationActive is a log parse operation binding the contract event 0x9a75057ae268e789c13fb43bbc257401731cb96df0eea6be3ff4a6bf416b415b.
//
// Solidity: event UpdateCollateralizationActive(bool active)
func (_Market *MarketFilterer) ParseUpdateCollateralizationActive(log types.Log) (*MarketUpdateCollateralizationActive, error) {
	event := new(MarketUpdateCollateralizationActive)
	if err := _Market.contract.UnpackLog(event, "UpdateCollateralizationActive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketUpdateComptrollerIterator is returned from FilterUpdateComptroller and is used to iterate over the raw logs and unpacked data for UpdateComptroller events raised by the Market contract.
type MarketUpdateComptrollerIterator struct {
	Event *MarketUpdateComptroller // Event containing the contract specifics and raw log

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
func (it *MarketUpdateComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketUpdateComptroller)
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
		it.Event = new(MarketUpdateComptroller)
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
func (it *MarketUpdateComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketUpdateComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketUpdateComptroller represents a UpdateComptroller event raised by the Market contract.
type MarketUpdateComptroller struct {
	Comptroller common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateComptroller is a free log retrieval operation binding the contract event 0x62f36648f1b25bb9a634cf6f54b4adedd8de98ab504f738600d724160b251a1e.
//
// Solidity: event UpdateComptroller(address indexed comptroller)
func (_Market *MarketFilterer) FilterUpdateComptroller(opts *bind.FilterOpts, comptroller []common.Address) (*MarketUpdateComptrollerIterator, error) {

	var comptrollerRule []interface{}
	for _, comptrollerItem := range comptroller {
		comptrollerRule = append(comptrollerRule, comptrollerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "UpdateComptroller", comptrollerRule)
	if err != nil {
		return nil, err
	}
	return &MarketUpdateComptrollerIterator{contract: _Market.contract, event: "UpdateComptroller", logs: logs, sub: sub}, nil
}

// WatchUpdateComptroller is a free log subscription operation binding the contract event 0x62f36648f1b25bb9a634cf6f54b4adedd8de98ab504f738600d724160b251a1e.
//
// Solidity: event UpdateComptroller(address indexed comptroller)
func (_Market *MarketFilterer) WatchUpdateComptroller(opts *bind.WatchOpts, sink chan<- *MarketUpdateComptroller, comptroller []common.Address) (event.Subscription, error) {

	var comptrollerRule []interface{}
	for _, comptrollerItem := range comptroller {
		comptrollerRule = append(comptrollerRule, comptrollerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "UpdateComptroller", comptrollerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketUpdateComptroller)
				if err := _Market.contract.UnpackLog(event, "UpdateComptroller", log); err != nil {
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
func (_Market *MarketFilterer) ParseUpdateComptroller(log types.Log) (*MarketUpdateComptroller, error) {
	event := new(MarketUpdateComptroller)
	if err := _Market.contract.UnpackLog(event, "UpdateComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Market contract.
type MarketWithdrawIterator struct {
	Event *MarketWithdraw // Event containing the contract specifics and raw log

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
func (it *MarketWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketWithdraw)
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
		it.Event = new(MarketWithdraw)
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
func (it *MarketWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketWithdraw represents a Withdraw event raised by the Market contract.
type MarketWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Market *MarketFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*MarketWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &MarketWithdrawIterator{contract: _Market.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Market *MarketFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MarketWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketWithdraw)
				if err := _Market.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Market *MarketFilterer) ParseWithdraw(log types.Log) (*MarketWithdraw, error) {
	event := new(MarketWithdraw)
	if err := _Market.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
