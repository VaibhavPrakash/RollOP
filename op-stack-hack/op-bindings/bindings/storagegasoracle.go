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

// StorageGasOracleRemoteGasDataConfig is an auto generated low-level Go binding around an user-defined struct.
type StorageGasOracleRemoteGasDataConfig struct {
	RemoteDomain      uint32
	TokenExchangeRate *big.Int
	GasPrice          *big.Int
}

// StorageGasOracleMetaData contains all meta data concerning the StorageGasOracle contract.
var StorageGasOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"tokenExchangeRate\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"gasPrice\",\"type\":\"uint128\"}],\"name\":\"RemoteGasDataSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"}],\"name\":\"getExchangeRateAndGasPrice\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"tokenExchangeRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"gasPrice\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"remoteGasData\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"tokenExchangeRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"gasPrice\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"tokenExchangeRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"gasPrice\",\"type\":\"uint128\"}],\"internalType\":\"structStorageGasOracle.RemoteGasDataConfig\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"setRemoteGasData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"tokenExchangeRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"gasPrice\",\"type\":\"uint128\"}],\"internalType\":\"structStorageGasOracle.RemoteGasDataConfig[]\",\"name\":\"_configs\",\"type\":\"tuple[]\"}],\"name\":\"setRemoteGasDataConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6106e48061007e6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b14610133578063b08e56d01461015b578063f2fde38b146101a5578063f3a1495f146101b857600080fd5b806360fcef7c14610082578063698faffc14610116578063715018a61461012b575b600080fd5b6100e8610090366004610527565b63ffffffff166000908152600160209081526040918290208251808401909352546fffffffffffffffffffffffffffffffff808216808552700100000000000000000000000000000000909204169290910182905291565b604080516fffffffffffffffffffffffffffffffff9384168152929091166020830152015b60405180910390f35b610129610124366004610554565b6101cb565b005b610129610217565b60005460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010d565b6100e8610169366004610527565b6001602052600090815260409020546fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041682565b6101296101b33660046105c9565b61022b565b6101296101c63660046105ff565b6102e7565b6101d36102f8565b8060005b81811015610211576101ff8484838181106101f4576101f4610617565b905060600201610379565b8061020981610646565b9150506101d7565b50505050565b61021f6102f8565b61022960006104b2565b565b6102336102f8565b73ffffffffffffffffffffffffffffffffffffffff81166102db576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102e4816104b2565b50565b6102ef6102f8565b6102e481610379565b60005473ffffffffffffffffffffffffffffffffffffffff163314610229576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d2565b604051806040016040528082602001602081019061039791906106a5565b6fffffffffffffffffffffffffffffffff1681526020016103be60608401604085016106a5565b6fffffffffffffffffffffffffffffffff169052600160006103e36020850185610527565b63ffffffff1681526020808201929092526040016000208251928201516fffffffffffffffffffffffffffffffff9081167001000000000000000000000000000000000293169290921790915561043c90820182610527565b63ffffffff167fb48c1cb713397fc0c0649596c221270fec0b3de3f85ccf6a734411a2fe57a69461047360408401602085016106a5565b61048360608501604086016106a5565b604080516fffffffffffffffffffffffffffffffff93841681529290911660208301520160405180910390a250565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561053957600080fd5b813563ffffffff8116811461054d57600080fd5b9392505050565b6000806020838503121561056757600080fd5b823567ffffffffffffffff8082111561057f57600080fd5b818501915085601f83011261059357600080fd5b8135818111156105a257600080fd5b8660206060830285010111156105b757600080fd5b60209290920196919550909350505050565b6000602082840312156105db57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461054d57600080fd5b60006060828403121561061157600080fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361069e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b6000602082840312156106b757600080fd5b81356fffffffffffffffffffffffffffffffff8116811461054d57600080fdfea164736f6c634300080f000a",
}

// StorageGasOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageGasOracleMetaData.ABI instead.
var StorageGasOracleABI = StorageGasOracleMetaData.ABI

// StorageGasOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageGasOracleMetaData.Bin instead.
var StorageGasOracleBin = StorageGasOracleMetaData.Bin

// DeployStorageGasOracle deploys a new Ethereum contract, binding an instance of StorageGasOracle to it.
func DeployStorageGasOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageGasOracle, error) {
	parsed, err := StorageGasOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageGasOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageGasOracle{StorageGasOracleCaller: StorageGasOracleCaller{contract: contract}, StorageGasOracleTransactor: StorageGasOracleTransactor{contract: contract}, StorageGasOracleFilterer: StorageGasOracleFilterer{contract: contract}}, nil
}

// StorageGasOracle is an auto generated Go binding around an Ethereum contract.
type StorageGasOracle struct {
	StorageGasOracleCaller     // Read-only binding to the contract
	StorageGasOracleTransactor // Write-only binding to the contract
	StorageGasOracleFilterer   // Log filterer for contract events
}

// StorageGasOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageGasOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageGasOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageGasOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageGasOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageGasOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageGasOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageGasOracleSession struct {
	Contract     *StorageGasOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageGasOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageGasOracleCallerSession struct {
	Contract *StorageGasOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// StorageGasOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageGasOracleTransactorSession struct {
	Contract     *StorageGasOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StorageGasOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageGasOracleRaw struct {
	Contract *StorageGasOracle // Generic contract binding to access the raw methods on
}

// StorageGasOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageGasOracleCallerRaw struct {
	Contract *StorageGasOracleCaller // Generic read-only contract binding to access the raw methods on
}

// StorageGasOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageGasOracleTransactorRaw struct {
	Contract *StorageGasOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageGasOracle creates a new instance of StorageGasOracle, bound to a specific deployed contract.
func NewStorageGasOracle(address common.Address, backend bind.ContractBackend) (*StorageGasOracle, error) {
	contract, err := bindStorageGasOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageGasOracle{StorageGasOracleCaller: StorageGasOracleCaller{contract: contract}, StorageGasOracleTransactor: StorageGasOracleTransactor{contract: contract}, StorageGasOracleFilterer: StorageGasOracleFilterer{contract: contract}}, nil
}

// NewStorageGasOracleCaller creates a new read-only instance of StorageGasOracle, bound to a specific deployed contract.
func NewStorageGasOracleCaller(address common.Address, caller bind.ContractCaller) (*StorageGasOracleCaller, error) {
	contract, err := bindStorageGasOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageGasOracleCaller{contract: contract}, nil
}

// NewStorageGasOracleTransactor creates a new write-only instance of StorageGasOracle, bound to a specific deployed contract.
func NewStorageGasOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageGasOracleTransactor, error) {
	contract, err := bindStorageGasOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageGasOracleTransactor{contract: contract}, nil
}

// NewStorageGasOracleFilterer creates a new log filterer instance of StorageGasOracle, bound to a specific deployed contract.
func NewStorageGasOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageGasOracleFilterer, error) {
	contract, err := bindStorageGasOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageGasOracleFilterer{contract: contract}, nil
}

// bindStorageGasOracle binds a generic wrapper to an already deployed contract.
func bindStorageGasOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageGasOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageGasOracle *StorageGasOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageGasOracle.Contract.StorageGasOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageGasOracle *StorageGasOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageGasOracle.Contract.StorageGasOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageGasOracle *StorageGasOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageGasOracle.Contract.StorageGasOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageGasOracle *StorageGasOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageGasOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageGasOracle *StorageGasOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageGasOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageGasOracle *StorageGasOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageGasOracle.Contract.contract.Transact(opts, method, params...)
}

// GetExchangeRateAndGasPrice is a free data retrieval call binding the contract method 0x60fcef7c.
//
// Solidity: function getExchangeRateAndGasPrice(uint32 _destinationDomain) view returns(uint128 tokenExchangeRate, uint128 gasPrice)
func (_StorageGasOracle *StorageGasOracleCaller) GetExchangeRateAndGasPrice(opts *bind.CallOpts, _destinationDomain uint32) (struct {
	TokenExchangeRate *big.Int
	GasPrice          *big.Int
}, error) {
	var out []interface{}
	err := _StorageGasOracle.contract.Call(opts, &out, "getExchangeRateAndGasPrice", _destinationDomain)

	outstruct := new(struct {
		TokenExchangeRate *big.Int
		GasPrice          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenExchangeRate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetExchangeRateAndGasPrice is a free data retrieval call binding the contract method 0x60fcef7c.
//
// Solidity: function getExchangeRateAndGasPrice(uint32 _destinationDomain) view returns(uint128 tokenExchangeRate, uint128 gasPrice)
func (_StorageGasOracle *StorageGasOracleSession) GetExchangeRateAndGasPrice(_destinationDomain uint32) (struct {
	TokenExchangeRate *big.Int
	GasPrice          *big.Int
}, error) {
	return _StorageGasOracle.Contract.GetExchangeRateAndGasPrice(&_StorageGasOracle.CallOpts, _destinationDomain)
}

// GetExchangeRateAndGasPrice is a free data retrieval call binding the contract method 0x60fcef7c.
//
// Solidity: function getExchangeRateAndGasPrice(uint32 _destinationDomain) view returns(uint128 tokenExchangeRate, uint128 gasPrice)
func (_StorageGasOracle *StorageGasOracleCallerSession) GetExchangeRateAndGasPrice(_destinationDomain uint32) (struct {
	TokenExchangeRate *big.Int
	GasPrice          *big.Int
}, error) {
	return _StorageGasOracle.Contract.GetExchangeRateAndGasPrice(&_StorageGasOracle.CallOpts, _destinationDomain)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StorageGasOracle *StorageGasOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StorageGasOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StorageGasOracle *StorageGasOracleSession) Owner() (common.Address, error) {
	return _StorageGasOracle.Contract.Owner(&_StorageGasOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StorageGasOracle *StorageGasOracleCallerSession) Owner() (common.Address, error) {
	return _StorageGasOracle.Contract.Owner(&_StorageGasOracle.CallOpts)
}

// RemoteGasData is a free data retrieval call binding the contract method 0xb08e56d0.
//
// Solidity: function remoteGasData(uint32 ) view returns(uint128 tokenExchangeRate, uint128 gasPrice)
func (_StorageGasOracle *StorageGasOracleCaller) RemoteGasData(opts *bind.CallOpts, arg0 uint32) (struct {
	TokenExchangeRate *big.Int
	GasPrice          *big.Int
}, error) {
	var out []interface{}
	err := _StorageGasOracle.contract.Call(opts, &out, "remoteGasData", arg0)

	outstruct := new(struct {
		TokenExchangeRate *big.Int
		GasPrice          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenExchangeRate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RemoteGasData is a free data retrieval call binding the contract method 0xb08e56d0.
//
// Solidity: function remoteGasData(uint32 ) view returns(uint128 tokenExchangeRate, uint128 gasPrice)
func (_StorageGasOracle *StorageGasOracleSession) RemoteGasData(arg0 uint32) (struct {
	TokenExchangeRate *big.Int
	GasPrice          *big.Int
}, error) {
	return _StorageGasOracle.Contract.RemoteGasData(&_StorageGasOracle.CallOpts, arg0)
}

// RemoteGasData is a free data retrieval call binding the contract method 0xb08e56d0.
//
// Solidity: function remoteGasData(uint32 ) view returns(uint128 tokenExchangeRate, uint128 gasPrice)
func (_StorageGasOracle *StorageGasOracleCallerSession) RemoteGasData(arg0 uint32) (struct {
	TokenExchangeRate *big.Int
	GasPrice          *big.Int
}, error) {
	return _StorageGasOracle.Contract.RemoteGasData(&_StorageGasOracle.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StorageGasOracle *StorageGasOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageGasOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StorageGasOracle *StorageGasOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _StorageGasOracle.Contract.RenounceOwnership(&_StorageGasOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StorageGasOracle *StorageGasOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StorageGasOracle.Contract.RenounceOwnership(&_StorageGasOracle.TransactOpts)
}

// SetRemoteGasData is a paid mutator transaction binding the contract method 0xf3a1495f.
//
// Solidity: function setRemoteGasData((uint32,uint128,uint128) _config) returns()
func (_StorageGasOracle *StorageGasOracleTransactor) SetRemoteGasData(opts *bind.TransactOpts, _config StorageGasOracleRemoteGasDataConfig) (*types.Transaction, error) {
	return _StorageGasOracle.contract.Transact(opts, "setRemoteGasData", _config)
}

// SetRemoteGasData is a paid mutator transaction binding the contract method 0xf3a1495f.
//
// Solidity: function setRemoteGasData((uint32,uint128,uint128) _config) returns()
func (_StorageGasOracle *StorageGasOracleSession) SetRemoteGasData(_config StorageGasOracleRemoteGasDataConfig) (*types.Transaction, error) {
	return _StorageGasOracle.Contract.SetRemoteGasData(&_StorageGasOracle.TransactOpts, _config)
}

// SetRemoteGasData is a paid mutator transaction binding the contract method 0xf3a1495f.
//
// Solidity: function setRemoteGasData((uint32,uint128,uint128) _config) returns()
func (_StorageGasOracle *StorageGasOracleTransactorSession) SetRemoteGasData(_config StorageGasOracleRemoteGasDataConfig) (*types.Transaction, error) {
	return _StorageGasOracle.Contract.SetRemoteGasData(&_StorageGasOracle.TransactOpts, _config)
}

// SetRemoteGasDataConfigs is a paid mutator transaction binding the contract method 0x698faffc.
//
// Solidity: function setRemoteGasDataConfigs((uint32,uint128,uint128)[] _configs) returns()
func (_StorageGasOracle *StorageGasOracleTransactor) SetRemoteGasDataConfigs(opts *bind.TransactOpts, _configs []StorageGasOracleRemoteGasDataConfig) (*types.Transaction, error) {
	return _StorageGasOracle.contract.Transact(opts, "setRemoteGasDataConfigs", _configs)
}

// SetRemoteGasDataConfigs is a paid mutator transaction binding the contract method 0x698faffc.
//
// Solidity: function setRemoteGasDataConfigs((uint32,uint128,uint128)[] _configs) returns()
func (_StorageGasOracle *StorageGasOracleSession) SetRemoteGasDataConfigs(_configs []StorageGasOracleRemoteGasDataConfig) (*types.Transaction, error) {
	return _StorageGasOracle.Contract.SetRemoteGasDataConfigs(&_StorageGasOracle.TransactOpts, _configs)
}

// SetRemoteGasDataConfigs is a paid mutator transaction binding the contract method 0x698faffc.
//
// Solidity: function setRemoteGasDataConfigs((uint32,uint128,uint128)[] _configs) returns()
func (_StorageGasOracle *StorageGasOracleTransactorSession) SetRemoteGasDataConfigs(_configs []StorageGasOracleRemoteGasDataConfig) (*types.Transaction, error) {
	return _StorageGasOracle.Contract.SetRemoteGasDataConfigs(&_StorageGasOracle.TransactOpts, _configs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StorageGasOracle *StorageGasOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StorageGasOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StorageGasOracle *StorageGasOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StorageGasOracle.Contract.TransferOwnership(&_StorageGasOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StorageGasOracle *StorageGasOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StorageGasOracle.Contract.TransferOwnership(&_StorageGasOracle.TransactOpts, newOwner)
}

// StorageGasOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StorageGasOracle contract.
type StorageGasOracleOwnershipTransferredIterator struct {
	Event *StorageGasOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StorageGasOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageGasOracleOwnershipTransferred)
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
		it.Event = new(StorageGasOracleOwnershipTransferred)
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
func (it *StorageGasOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageGasOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageGasOracleOwnershipTransferred represents a OwnershipTransferred event raised by the StorageGasOracle contract.
type StorageGasOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StorageGasOracle *StorageGasOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StorageGasOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StorageGasOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StorageGasOracleOwnershipTransferredIterator{contract: _StorageGasOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StorageGasOracle *StorageGasOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StorageGasOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StorageGasOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageGasOracleOwnershipTransferred)
				if err := _StorageGasOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StorageGasOracle *StorageGasOracleFilterer) ParseOwnershipTransferred(log types.Log) (*StorageGasOracleOwnershipTransferred, error) {
	event := new(StorageGasOracleOwnershipTransferred)
	if err := _StorageGasOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageGasOracleRemoteGasDataSetIterator is returned from FilterRemoteGasDataSet and is used to iterate over the raw logs and unpacked data for RemoteGasDataSet events raised by the StorageGasOracle contract.
type StorageGasOracleRemoteGasDataSetIterator struct {
	Event *StorageGasOracleRemoteGasDataSet // Event containing the contract specifics and raw log

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
func (it *StorageGasOracleRemoteGasDataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageGasOracleRemoteGasDataSet)
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
		it.Event = new(StorageGasOracleRemoteGasDataSet)
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
func (it *StorageGasOracleRemoteGasDataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageGasOracleRemoteGasDataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageGasOracleRemoteGasDataSet represents a RemoteGasDataSet event raised by the StorageGasOracle contract.
type StorageGasOracleRemoteGasDataSet struct {
	RemoteDomain      uint32
	TokenExchangeRate *big.Int
	GasPrice          *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRemoteGasDataSet is a free log retrieval operation binding the contract event 0xb48c1cb713397fc0c0649596c221270fec0b3de3f85ccf6a734411a2fe57a694.
//
// Solidity: event RemoteGasDataSet(uint32 indexed remoteDomain, uint128 tokenExchangeRate, uint128 gasPrice)
func (_StorageGasOracle *StorageGasOracleFilterer) FilterRemoteGasDataSet(opts *bind.FilterOpts, remoteDomain []uint32) (*StorageGasOracleRemoteGasDataSetIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}

	logs, sub, err := _StorageGasOracle.contract.FilterLogs(opts, "RemoteGasDataSet", remoteDomainRule)
	if err != nil {
		return nil, err
	}
	return &StorageGasOracleRemoteGasDataSetIterator{contract: _StorageGasOracle.contract, event: "RemoteGasDataSet", logs: logs, sub: sub}, nil
}

// WatchRemoteGasDataSet is a free log subscription operation binding the contract event 0xb48c1cb713397fc0c0649596c221270fec0b3de3f85ccf6a734411a2fe57a694.
//
// Solidity: event RemoteGasDataSet(uint32 indexed remoteDomain, uint128 tokenExchangeRate, uint128 gasPrice)
func (_StorageGasOracle *StorageGasOracleFilterer) WatchRemoteGasDataSet(opts *bind.WatchOpts, sink chan<- *StorageGasOracleRemoteGasDataSet, remoteDomain []uint32) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}

	logs, sub, err := _StorageGasOracle.contract.WatchLogs(opts, "RemoteGasDataSet", remoteDomainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageGasOracleRemoteGasDataSet)
				if err := _StorageGasOracle.contract.UnpackLog(event, "RemoteGasDataSet", log); err != nil {
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

// ParseRemoteGasDataSet is a log parse operation binding the contract event 0xb48c1cb713397fc0c0649596c221270fec0b3de3f85ccf6a734411a2fe57a694.
//
// Solidity: event RemoteGasDataSet(uint32 indexed remoteDomain, uint128 tokenExchangeRate, uint128 gasPrice)
func (_StorageGasOracle *StorageGasOracleFilterer) ParseRemoteGasDataSet(log types.Log) (*StorageGasOracleRemoteGasDataSet, error) {
	event := new(StorageGasOracleRemoteGasDataSet)
	if err := _StorageGasOracle.contract.UnpackLog(event, "RemoteGasDataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
