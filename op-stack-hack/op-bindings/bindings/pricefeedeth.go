// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// PriceFeedEthMetaData contains all meta data concerning the PriceFeedEth contract.
var PriceFeedEthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"currentPrice\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_blockNo\",\"type\":\"uint64\"}],\"name\":\"getPriceBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_blockNo\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_price\",\"type\":\"uint64\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061024f806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063464be46f1461005157806398d5fdca1461009b5780639d1b464a146100ad578063aa415b52146100c1575b600080fd5b61007e61005f3660046101ed565b67ffffffffffffffff9081166000908152600160205260409020541690565b60405167ffffffffffffffff909116815260200160405180910390f35b60005467ffffffffffffffff1661007e565b60005461007e9067ffffffffffffffff1681565b6100d46100cf36600461020f565b6100d6565b005b73deaddeaddeaddeaddeaddeaddeaddeaddead0001331461017d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4f6e6c792073797374656d20616464726573732063616e2063616c6c2074686960448201527f732066756e6374696f6e00000000000000000000000000000000000000000000606482015260840160405180910390fd5b6000805467ffffffffffffffff9283167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009182168117835593909216815260016020526040902080549091169091179055565b803567ffffffffffffffff811681146101e857600080fd5b919050565b6000602082840312156101ff57600080fd5b610208826101d0565b9392505050565b6000806040838503121561022257600080fd5b61022b836101d0565b9150610239602084016101d0565b9050925092905056fea164736f6c6343000813000a",
}

// PriceFeedEthABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceFeedEthMetaData.ABI instead.
var PriceFeedEthABI = PriceFeedEthMetaData.ABI

// PriceFeedEthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PriceFeedEthMetaData.Bin instead.
var PriceFeedEthBin = PriceFeedEthMetaData.Bin

// DeployPriceFeedEth deploys a new Ethereum contract, binding an instance of PriceFeedEth to it.
func DeployPriceFeedEth(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriceFeedEth, error) {
	parsed, err := PriceFeedEthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PriceFeedEthBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceFeedEth{PriceFeedEthCaller: PriceFeedEthCaller{contract: contract}, PriceFeedEthTransactor: PriceFeedEthTransactor{contract: contract}, PriceFeedEthFilterer: PriceFeedEthFilterer{contract: contract}}, nil
}

// PriceFeedEth is an auto generated Go binding around an Ethereum contract.
type PriceFeedEth struct {
	PriceFeedEthCaller     // Read-only binding to the contract
	PriceFeedEthTransactor // Write-only binding to the contract
	PriceFeedEthFilterer   // Log filterer for contract events
}

// PriceFeedEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeedEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeedEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeedEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeedEthSession struct {
	Contract     *PriceFeedEth     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceFeedEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeedEthCallerSession struct {
	Contract *PriceFeedEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PriceFeedEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeedEthTransactorSession struct {
	Contract     *PriceFeedEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PriceFeedEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeedEthRaw struct {
	Contract *PriceFeedEth // Generic contract binding to access the raw methods on
}

// PriceFeedEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeedEthCallerRaw struct {
	Contract *PriceFeedEthCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeedEthTransactorRaw struct {
	Contract *PriceFeedEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeedEth creates a new instance of PriceFeedEth, bound to a specific deployed contract.
func NewPriceFeedEth(address common.Address, backend bind.ContractBackend) (*PriceFeedEth, error) {
	contract, err := bindPriceFeedEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeedEth{PriceFeedEthCaller: PriceFeedEthCaller{contract: contract}, PriceFeedEthTransactor: PriceFeedEthTransactor{contract: contract}, PriceFeedEthFilterer: PriceFeedEthFilterer{contract: contract}}, nil
}

// NewPriceFeedEthCaller creates a new read-only instance of PriceFeedEth, bound to a specific deployed contract.
func NewPriceFeedEthCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedEthCaller, error) {
	contract, err := bindPriceFeedEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedEthCaller{contract: contract}, nil
}

// NewPriceFeedEthTransactor creates a new write-only instance of PriceFeedEth, bound to a specific deployed contract.
func NewPriceFeedEthTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedEthTransactor, error) {
	contract, err := bindPriceFeedEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedEthTransactor{contract: contract}, nil
}

// NewPriceFeedEthFilterer creates a new log filterer instance of PriceFeedEth, bound to a specific deployed contract.
func NewPriceFeedEthFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedEthFilterer, error) {
	contract, err := bindPriceFeedEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedEthFilterer{contract: contract}, nil
}

// bindPriceFeedEth binds a generic wrapper to an already deployed contract.
func bindPriceFeedEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceFeedEthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedEth *PriceFeedEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedEth.Contract.PriceFeedEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedEth *PriceFeedEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedEth.Contract.PriceFeedEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedEth *PriceFeedEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedEth.Contract.PriceFeedEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedEth *PriceFeedEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedEth *PriceFeedEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedEth *PriceFeedEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedEth.Contract.contract.Transact(opts, method, params...)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint64)
func (_PriceFeedEth *PriceFeedEthCaller) CurrentPrice(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PriceFeedEth.contract.Call(opts, &out, "currentPrice")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint64)
func (_PriceFeedEth *PriceFeedEthSession) CurrentPrice() (uint64, error) {
	return _PriceFeedEth.Contract.CurrentPrice(&_PriceFeedEth.CallOpts)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint64)
func (_PriceFeedEth *PriceFeedEthCallerSession) CurrentPrice() (uint64, error) {
	return _PriceFeedEth.Contract.CurrentPrice(&_PriceFeedEth.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint64)
func (_PriceFeedEth *PriceFeedEthCaller) GetPrice(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PriceFeedEth.contract.Call(opts, &out, "getPrice")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint64)
func (_PriceFeedEth *PriceFeedEthSession) GetPrice() (uint64, error) {
	return _PriceFeedEth.Contract.GetPrice(&_PriceFeedEth.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint64)
func (_PriceFeedEth *PriceFeedEthCallerSession) GetPrice() (uint64, error) {
	return _PriceFeedEth.Contract.GetPrice(&_PriceFeedEth.CallOpts)
}

// GetPriceBlock is a free data retrieval call binding the contract method 0x464be46f.
//
// Solidity: function getPriceBlock(uint64 _blockNo) view returns(uint64)
func (_PriceFeedEth *PriceFeedEthCaller) GetPriceBlock(opts *bind.CallOpts, _blockNo uint64) (uint64, error) {
	var out []interface{}
	err := _PriceFeedEth.contract.Call(opts, &out, "getPriceBlock", _blockNo)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetPriceBlock is a free data retrieval call binding the contract method 0x464be46f.
//
// Solidity: function getPriceBlock(uint64 _blockNo) view returns(uint64)
func (_PriceFeedEth *PriceFeedEthSession) GetPriceBlock(_blockNo uint64) (uint64, error) {
	return _PriceFeedEth.Contract.GetPriceBlock(&_PriceFeedEth.CallOpts, _blockNo)
}

// GetPriceBlock is a free data retrieval call binding the contract method 0x464be46f.
//
// Solidity: function getPriceBlock(uint64 _blockNo) view returns(uint64)
func (_PriceFeedEth *PriceFeedEthCallerSession) GetPriceBlock(_blockNo uint64) (uint64, error) {
	return _PriceFeedEth.Contract.GetPriceBlock(&_PriceFeedEth.CallOpts, _blockNo)
}

// SetValue is a paid mutator transaction binding the contract method 0xaa415b52.
//
// Solidity: function setValue(uint64 _blockNo, uint64 _price) returns()
func (_PriceFeedEth *PriceFeedEthTransactor) SetValue(opts *bind.TransactOpts, _blockNo uint64, _price uint64) (*types.Transaction, error) {
	return _PriceFeedEth.contract.Transact(opts, "setValue", _blockNo, _price)
}

// SetValue is a paid mutator transaction binding the contract method 0xaa415b52.
//
// Solidity: function setValue(uint64 _blockNo, uint64 _price) returns()
func (_PriceFeedEth *PriceFeedEthSession) SetValue(_blockNo uint64, _price uint64) (*types.Transaction, error) {
	return _PriceFeedEth.Contract.SetValue(&_PriceFeedEth.TransactOpts, _blockNo, _price)
}

// SetValue is a paid mutator transaction binding the contract method 0xaa415b52.
//
// Solidity: function setValue(uint64 _blockNo, uint64 _price) returns()
func (_PriceFeedEth *PriceFeedEthTransactorSession) SetValue(_blockNo uint64, _price uint64) (*types.Transaction, error) {
	return _PriceFeedEth.Contract.SetValue(&_PriceFeedEth.TransactOpts, _blockNo, _price)
}
