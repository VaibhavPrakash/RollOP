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

// InterchainGasPaymasterGasOracleConfig is an auto generated low-level Go binding around an user-defined struct.
type InterchainGasPaymasterGasOracleConfig struct {
	RemoteDomain uint32
	GasOracle    common.Address
}

// InterchainGasPaymasterMetaData contains all meta data concerning the InterchainGasPaymaster contract.
var InterchainGasPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"BeneficiarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"GasPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"gasOracles\",\"outputs\":[{\"internalType\":\"contractIGasOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"}],\"name\":\"getExchangeRateAndGasPrice\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"tokenExchangeRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"gasPrice\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_gasAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"payForGas\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_gasAmount\",\"type\":\"uint256\"}],\"name\":\"quoteGasPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"internalType\":\"structInterchainGasPaymaster.GasOracleConfig[]\",\"name\":\"_configs\",\"type\":\"tuple[]\"}],\"name\":\"setGasOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061106b806100206000396000f3fe6080604052600436106100c75760003560e01c80634e71d92d116100745780638da5cb5b1161004e5780638da5cb5b1461024e578063a692979314610279578063f2fde38b146102a757600080fd5b80634e71d92d146101db57806360fcef7c146101f0578063715018a61461023957600080fd5b80632f202650116100a55780632f2026501461016e57806338af3eed1461018e578063485cc955146101bb57600080fd5b806311bf2c18146100cc5780631c31f710146100e15780631d16c8c814610101575b600080fd5b6100df6100da366004610d9a565b6102c7565b005b3480156100ed57600080fd5b506100df6100fc366004610de0565b6104b2565b34801561010d57600080fd5b5061014461011c366004610e02565b60656020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561017a57600080fd5b506100df610189366004610e1d565b6104c6565b34801561019a57600080fd5b506066546101449073ffffffffffffffffffffffffffffffffffffffff1681565b3480156101c757600080fd5b506100df6101d6366004610e92565b61054c565b3480156101e757600080fd5b506100df6106f2565b3480156101fc57600080fd5b5061021061020b366004610e02565b6107bc565b604080516fffffffffffffffffffffffffffffffff938416815292909116602083015201610165565b34801561024557600080fd5b506100df6108ee565b34801561025a57600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610144565b34801561028557600080fd5b50610299610294366004610ec5565b610902565b604051908152602001610165565b3480156102b357600080fd5b506100df6102c2366004610de0565b61096b565b60006102d38484610902565b90508034101561036a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f696e73756666696369656e7420696e746572636861696e20676173207061796d60448201527f656e74000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b60006103768234610f1e565b9050801561046f5760008373ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d80600081146103d8576040519150601f19603f3d011682016040523d82523d6000602084013e6103dd565b606091505b505090508061046d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f496e746572636861696e20676173207061796d656e7420726566756e6420666160448201527f696c6564000000000000000000000000000000000000000000000000000000006064820152608401610361565b505b604080518581526020810184905287917ff715e66d2cd2a0179069dcda2a41a4da36a30e1aad12187fbb7c5d100afd621c910160405180910390a2505050505050565b6104ba610a1f565b6104c381610aa0565b50565b6104ce610a1f565b8060005b81811015610546576105348484838181106104ef576104ef610f35565b6105059260206040909202019081019150610e02565b85858481811061051757610517610f35565b905060400201602001602081019061052f9190610de0565b610b19565b8061053e81610f64565b9150506104d2565b50505050565b600054610100900460ff161580801561056c5750600054600160ff909116105b806105865750303b158015610586575060005460ff166001145b610612576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610361565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561067057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610678610ba7565b61068183610c46565b61068a82610aa0565b80156106ed57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60665460405160009173ffffffffffffffffffffffffffffffffffffffff169047908381818185875af1925050503d806000811461074c576040519150601f19603f3d011682016040523d82523d6000602084013e610751565b606091505b50509050806104c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f217472616e7366657200000000000000000000000000000000000000000000006044820152606401610361565b63ffffffff8116600090815260656020526040812054819073ffffffffffffffffffffffffffffffffffffffff1680610851576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f21676173206f7261636c650000000000000000000000000000000000000000006044820152606401610361565b6040517f60fcef7c00000000000000000000000000000000000000000000000000000000815263ffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8216906360fcef7c906024016040805180830381865afa1580156108c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e49190610fbc565b9250925050915091565b6108f6610a1f565b6109006000610c46565b565b6000806000610910856107bc565b909250905060006109336fffffffffffffffffffffffffffffffff831686610fe6565b90506402540be4006109576fffffffffffffffffffffffffffffffff851683610fe6565b6109619190611023565b9695505050505050565b610973610a1f565b73ffffffffffffffffffffffffffffffffffffffff8116610a16576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610361565b6104c381610c46565b60335473ffffffffffffffffffffffffffffffffffffffff163314610900576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610361565b606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f04d55a8be181fb8d75b76f2d48aa0b2ee40f47e53d6e61763eeeec46feea8a249060200160405180910390a150565b63ffffffff821660008181526065602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff861690811790915591519182527f38b842e116ff320f6a79ba4cf434eee927fa156b226c65c90d62989de589b4e7910160405180910390a25050565b600054610100900460ff16610c3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610361565b610900610cbd565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610d54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610361565b61090033610c46565b803563ffffffff81168114610d7157600080fd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610d7157600080fd5b60008060008060808587031215610db057600080fd5b84359350610dc060208601610d5d565b925060408501359150610dd560608601610d76565b905092959194509250565b600060208284031215610df257600080fd5b610dfb82610d76565b9392505050565b600060208284031215610e1457600080fd5b610dfb82610d5d565b60008060208385031215610e3057600080fd5b823567ffffffffffffffff80821115610e4857600080fd5b818501915085601f830112610e5c57600080fd5b813581811115610e6b57600080fd5b8660208260061b8501011115610e8057600080fd5b60209290920196919550909350505050565b60008060408385031215610ea557600080fd5b610eae83610d76565b9150610ebc60208401610d76565b90509250929050565b60008060408385031215610ed857600080fd5b610ee183610d5d565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015610f3057610f30610eef565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f9557610f95610eef565b5060010190565b80516fffffffffffffffffffffffffffffffff81168114610d7157600080fd5b60008060408385031215610fcf57600080fd5b610fd883610f9c565b9150610ebc60208401610f9c565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561101e5761101e610eef565b500290565b600082611059577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea164736f6c634300080f000a",
}

// InterchainGasPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainGasPaymasterMetaData.ABI instead.
var InterchainGasPaymasterABI = InterchainGasPaymasterMetaData.ABI

// InterchainGasPaymasterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainGasPaymasterMetaData.Bin instead.
var InterchainGasPaymasterBin = InterchainGasPaymasterMetaData.Bin

// DeployInterchainGasPaymaster deploys a new Ethereum contract, binding an instance of InterchainGasPaymaster to it.
func DeployInterchainGasPaymaster(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InterchainGasPaymaster, error) {
	parsed, err := InterchainGasPaymasterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainGasPaymasterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainGasPaymaster{InterchainGasPaymasterCaller: InterchainGasPaymasterCaller{contract: contract}, InterchainGasPaymasterTransactor: InterchainGasPaymasterTransactor{contract: contract}, InterchainGasPaymasterFilterer: InterchainGasPaymasterFilterer{contract: contract}}, nil
}

// InterchainGasPaymaster is an auto generated Go binding around an Ethereum contract.
type InterchainGasPaymaster struct {
	InterchainGasPaymasterCaller     // Read-only binding to the contract
	InterchainGasPaymasterTransactor // Write-only binding to the contract
	InterchainGasPaymasterFilterer   // Log filterer for contract events
}

// InterchainGasPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainGasPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainGasPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainGasPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainGasPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainGasPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainGasPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainGasPaymasterSession struct {
	Contract     *InterchainGasPaymaster // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InterchainGasPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainGasPaymasterCallerSession struct {
	Contract *InterchainGasPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// InterchainGasPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainGasPaymasterTransactorSession struct {
	Contract     *InterchainGasPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// InterchainGasPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainGasPaymasterRaw struct {
	Contract *InterchainGasPaymaster // Generic contract binding to access the raw methods on
}

// InterchainGasPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainGasPaymasterCallerRaw struct {
	Contract *InterchainGasPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainGasPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainGasPaymasterTransactorRaw struct {
	Contract *InterchainGasPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainGasPaymaster creates a new instance of InterchainGasPaymaster, bound to a specific deployed contract.
func NewInterchainGasPaymaster(address common.Address, backend bind.ContractBackend) (*InterchainGasPaymaster, error) {
	contract, err := bindInterchainGasPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainGasPaymaster{InterchainGasPaymasterCaller: InterchainGasPaymasterCaller{contract: contract}, InterchainGasPaymasterTransactor: InterchainGasPaymasterTransactor{contract: contract}, InterchainGasPaymasterFilterer: InterchainGasPaymasterFilterer{contract: contract}}, nil
}

// NewInterchainGasPaymasterCaller creates a new read-only instance of InterchainGasPaymaster, bound to a specific deployed contract.
func NewInterchainGasPaymasterCaller(address common.Address, caller bind.ContractCaller) (*InterchainGasPaymasterCaller, error) {
	contract, err := bindInterchainGasPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainGasPaymasterCaller{contract: contract}, nil
}

// NewInterchainGasPaymasterTransactor creates a new write-only instance of InterchainGasPaymaster, bound to a specific deployed contract.
func NewInterchainGasPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainGasPaymasterTransactor, error) {
	contract, err := bindInterchainGasPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainGasPaymasterTransactor{contract: contract}, nil
}

// NewInterchainGasPaymasterFilterer creates a new log filterer instance of InterchainGasPaymaster, bound to a specific deployed contract.
func NewInterchainGasPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainGasPaymasterFilterer, error) {
	contract, err := bindInterchainGasPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainGasPaymasterFilterer{contract: contract}, nil
}

// bindInterchainGasPaymaster binds a generic wrapper to an already deployed contract.
func bindInterchainGasPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainGasPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainGasPaymaster *InterchainGasPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainGasPaymaster.Contract.InterchainGasPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainGasPaymaster *InterchainGasPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.InterchainGasPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainGasPaymaster *InterchainGasPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.InterchainGasPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainGasPaymaster *InterchainGasPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainGasPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_InterchainGasPaymaster *InterchainGasPaymasterCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainGasPaymaster.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_InterchainGasPaymaster *InterchainGasPaymasterSession) Beneficiary() (common.Address, error) {
	return _InterchainGasPaymaster.Contract.Beneficiary(&_InterchainGasPaymaster.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_InterchainGasPaymaster *InterchainGasPaymasterCallerSession) Beneficiary() (common.Address, error) {
	return _InterchainGasPaymaster.Contract.Beneficiary(&_InterchainGasPaymaster.CallOpts)
}

// GasOracles is a free data retrieval call binding the contract method 0x1d16c8c8.
//
// Solidity: function gasOracles(uint32 ) view returns(address)
func (_InterchainGasPaymaster *InterchainGasPaymasterCaller) GasOracles(opts *bind.CallOpts, arg0 uint32) (common.Address, error) {
	var out []interface{}
	err := _InterchainGasPaymaster.contract.Call(opts, &out, "gasOracles", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasOracles is a free data retrieval call binding the contract method 0x1d16c8c8.
//
// Solidity: function gasOracles(uint32 ) view returns(address)
func (_InterchainGasPaymaster *InterchainGasPaymasterSession) GasOracles(arg0 uint32) (common.Address, error) {
	return _InterchainGasPaymaster.Contract.GasOracles(&_InterchainGasPaymaster.CallOpts, arg0)
}

// GasOracles is a free data retrieval call binding the contract method 0x1d16c8c8.
//
// Solidity: function gasOracles(uint32 ) view returns(address)
func (_InterchainGasPaymaster *InterchainGasPaymasterCallerSession) GasOracles(arg0 uint32) (common.Address, error) {
	return _InterchainGasPaymaster.Contract.GasOracles(&_InterchainGasPaymaster.CallOpts, arg0)
}

// GetExchangeRateAndGasPrice is a free data retrieval call binding the contract method 0x60fcef7c.
//
// Solidity: function getExchangeRateAndGasPrice(uint32 _destinationDomain) view returns(uint128 tokenExchangeRate, uint128 gasPrice)
func (_InterchainGasPaymaster *InterchainGasPaymasterCaller) GetExchangeRateAndGasPrice(opts *bind.CallOpts, _destinationDomain uint32) (struct {
	TokenExchangeRate *big.Int
	GasPrice          *big.Int
}, error) {
	var out []interface{}
	err := _InterchainGasPaymaster.contract.Call(opts, &out, "getExchangeRateAndGasPrice", _destinationDomain)

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
func (_InterchainGasPaymaster *InterchainGasPaymasterSession) GetExchangeRateAndGasPrice(_destinationDomain uint32) (struct {
	TokenExchangeRate *big.Int
	GasPrice          *big.Int
}, error) {
	return _InterchainGasPaymaster.Contract.GetExchangeRateAndGasPrice(&_InterchainGasPaymaster.CallOpts, _destinationDomain)
}

// GetExchangeRateAndGasPrice is a free data retrieval call binding the contract method 0x60fcef7c.
//
// Solidity: function getExchangeRateAndGasPrice(uint32 _destinationDomain) view returns(uint128 tokenExchangeRate, uint128 gasPrice)
func (_InterchainGasPaymaster *InterchainGasPaymasterCallerSession) GetExchangeRateAndGasPrice(_destinationDomain uint32) (struct {
	TokenExchangeRate *big.Int
	GasPrice          *big.Int
}, error) {
	return _InterchainGasPaymaster.Contract.GetExchangeRateAndGasPrice(&_InterchainGasPaymaster.CallOpts, _destinationDomain)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterchainGasPaymaster *InterchainGasPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainGasPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterchainGasPaymaster *InterchainGasPaymasterSession) Owner() (common.Address, error) {
	return _InterchainGasPaymaster.Contract.Owner(&_InterchainGasPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterchainGasPaymaster *InterchainGasPaymasterCallerSession) Owner() (common.Address, error) {
	return _InterchainGasPaymaster.Contract.Owner(&_InterchainGasPaymaster.CallOpts)
}

// QuoteGasPayment is a free data retrieval call binding the contract method 0xa6929793.
//
// Solidity: function quoteGasPayment(uint32 _destinationDomain, uint256 _gasAmount) view returns(uint256)
func (_InterchainGasPaymaster *InterchainGasPaymasterCaller) QuoteGasPayment(opts *bind.CallOpts, _destinationDomain uint32, _gasAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterchainGasPaymaster.contract.Call(opts, &out, "quoteGasPayment", _destinationDomain, _gasAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteGasPayment is a free data retrieval call binding the contract method 0xa6929793.
//
// Solidity: function quoteGasPayment(uint32 _destinationDomain, uint256 _gasAmount) view returns(uint256)
func (_InterchainGasPaymaster *InterchainGasPaymasterSession) QuoteGasPayment(_destinationDomain uint32, _gasAmount *big.Int) (*big.Int, error) {
	return _InterchainGasPaymaster.Contract.QuoteGasPayment(&_InterchainGasPaymaster.CallOpts, _destinationDomain, _gasAmount)
}

// QuoteGasPayment is a free data retrieval call binding the contract method 0xa6929793.
//
// Solidity: function quoteGasPayment(uint32 _destinationDomain, uint256 _gasAmount) view returns(uint256)
func (_InterchainGasPaymaster *InterchainGasPaymasterCallerSession) QuoteGasPayment(_destinationDomain uint32, _gasAmount *big.Int) (*big.Int, error) {
	return _InterchainGasPaymaster.Contract.QuoteGasPayment(&_InterchainGasPaymaster.CallOpts, _destinationDomain, _gasAmount)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainGasPaymaster.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterSession) Claim() (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.Claim(&_InterchainGasPaymaster.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactorSession) Claim() (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.Claim(&_InterchainGasPaymaster.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _beneficiary) returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _InterchainGasPaymaster.contract.Transact(opts, "initialize", _owner, _beneficiary)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _beneficiary) returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterSession) Initialize(_owner common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.Initialize(&_InterchainGasPaymaster.TransactOpts, _owner, _beneficiary)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _beneficiary) returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactorSession) Initialize(_owner common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.Initialize(&_InterchainGasPaymaster.TransactOpts, _owner, _beneficiary)
}

// PayForGas is a paid mutator transaction binding the contract method 0x11bf2c18.
//
// Solidity: function payForGas(bytes32 _messageId, uint32 _destinationDomain, uint256 _gasAmount, address _refundAddress) payable returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactor) PayForGas(opts *bind.TransactOpts, _messageId [32]byte, _destinationDomain uint32, _gasAmount *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _InterchainGasPaymaster.contract.Transact(opts, "payForGas", _messageId, _destinationDomain, _gasAmount, _refundAddress)
}

// PayForGas is a paid mutator transaction binding the contract method 0x11bf2c18.
//
// Solidity: function payForGas(bytes32 _messageId, uint32 _destinationDomain, uint256 _gasAmount, address _refundAddress) payable returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterSession) PayForGas(_messageId [32]byte, _destinationDomain uint32, _gasAmount *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.PayForGas(&_InterchainGasPaymaster.TransactOpts, _messageId, _destinationDomain, _gasAmount, _refundAddress)
}

// PayForGas is a paid mutator transaction binding the contract method 0x11bf2c18.
//
// Solidity: function payForGas(bytes32 _messageId, uint32 _destinationDomain, uint256 _gasAmount, address _refundAddress) payable returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactorSession) PayForGas(_messageId [32]byte, _destinationDomain uint32, _gasAmount *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.PayForGas(&_InterchainGasPaymaster.TransactOpts, _messageId, _destinationDomain, _gasAmount, _refundAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainGasPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.RenounceOwnership(&_InterchainGasPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.RenounceOwnership(&_InterchainGasPaymaster.TransactOpts)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address _beneficiary) returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactor) SetBeneficiary(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _InterchainGasPaymaster.contract.Transact(opts, "setBeneficiary", _beneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address _beneficiary) returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterSession) SetBeneficiary(_beneficiary common.Address) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.SetBeneficiary(&_InterchainGasPaymaster.TransactOpts, _beneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address _beneficiary) returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactorSession) SetBeneficiary(_beneficiary common.Address) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.SetBeneficiary(&_InterchainGasPaymaster.TransactOpts, _beneficiary)
}

// SetGasOracles is a paid mutator transaction binding the contract method 0x2f202650.
//
// Solidity: function setGasOracles((uint32,address)[] _configs) returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactor) SetGasOracles(opts *bind.TransactOpts, _configs []InterchainGasPaymasterGasOracleConfig) (*types.Transaction, error) {
	return _InterchainGasPaymaster.contract.Transact(opts, "setGasOracles", _configs)
}

// SetGasOracles is a paid mutator transaction binding the contract method 0x2f202650.
//
// Solidity: function setGasOracles((uint32,address)[] _configs) returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterSession) SetGasOracles(_configs []InterchainGasPaymasterGasOracleConfig) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.SetGasOracles(&_InterchainGasPaymaster.TransactOpts, _configs)
}

// SetGasOracles is a paid mutator transaction binding the contract method 0x2f202650.
//
// Solidity: function setGasOracles((uint32,address)[] _configs) returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactorSession) SetGasOracles(_configs []InterchainGasPaymasterGasOracleConfig) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.SetGasOracles(&_InterchainGasPaymaster.TransactOpts, _configs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InterchainGasPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.TransferOwnership(&_InterchainGasPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InterchainGasPaymaster *InterchainGasPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InterchainGasPaymaster.Contract.TransferOwnership(&_InterchainGasPaymaster.TransactOpts, newOwner)
}

// InterchainGasPaymasterBeneficiarySetIterator is returned from FilterBeneficiarySet and is used to iterate over the raw logs and unpacked data for BeneficiarySet events raised by the InterchainGasPaymaster contract.
type InterchainGasPaymasterBeneficiarySetIterator struct {
	Event *InterchainGasPaymasterBeneficiarySet // Event containing the contract specifics and raw log

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
func (it *InterchainGasPaymasterBeneficiarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainGasPaymasterBeneficiarySet)
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
		it.Event = new(InterchainGasPaymasterBeneficiarySet)
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
func (it *InterchainGasPaymasterBeneficiarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainGasPaymasterBeneficiarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainGasPaymasterBeneficiarySet represents a BeneficiarySet event raised by the InterchainGasPaymaster contract.
type InterchainGasPaymasterBeneficiarySet struct {
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBeneficiarySet is a free log retrieval operation binding the contract event 0x04d55a8be181fb8d75b76f2d48aa0b2ee40f47e53d6e61763eeeec46feea8a24.
//
// Solidity: event BeneficiarySet(address beneficiary)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) FilterBeneficiarySet(opts *bind.FilterOpts) (*InterchainGasPaymasterBeneficiarySetIterator, error) {

	logs, sub, err := _InterchainGasPaymaster.contract.FilterLogs(opts, "BeneficiarySet")
	if err != nil {
		return nil, err
	}
	return &InterchainGasPaymasterBeneficiarySetIterator{contract: _InterchainGasPaymaster.contract, event: "BeneficiarySet", logs: logs, sub: sub}, nil
}

// WatchBeneficiarySet is a free log subscription operation binding the contract event 0x04d55a8be181fb8d75b76f2d48aa0b2ee40f47e53d6e61763eeeec46feea8a24.
//
// Solidity: event BeneficiarySet(address beneficiary)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) WatchBeneficiarySet(opts *bind.WatchOpts, sink chan<- *InterchainGasPaymasterBeneficiarySet) (event.Subscription, error) {

	logs, sub, err := _InterchainGasPaymaster.contract.WatchLogs(opts, "BeneficiarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainGasPaymasterBeneficiarySet)
				if err := _InterchainGasPaymaster.contract.UnpackLog(event, "BeneficiarySet", log); err != nil {
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

// ParseBeneficiarySet is a log parse operation binding the contract event 0x04d55a8be181fb8d75b76f2d48aa0b2ee40f47e53d6e61763eeeec46feea8a24.
//
// Solidity: event BeneficiarySet(address beneficiary)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) ParseBeneficiarySet(log types.Log) (*InterchainGasPaymasterBeneficiarySet, error) {
	event := new(InterchainGasPaymasterBeneficiarySet)
	if err := _InterchainGasPaymaster.contract.UnpackLog(event, "BeneficiarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainGasPaymasterGasOracleSetIterator is returned from FilterGasOracleSet and is used to iterate over the raw logs and unpacked data for GasOracleSet events raised by the InterchainGasPaymaster contract.
type InterchainGasPaymasterGasOracleSetIterator struct {
	Event *InterchainGasPaymasterGasOracleSet // Event containing the contract specifics and raw log

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
func (it *InterchainGasPaymasterGasOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainGasPaymasterGasOracleSet)
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
		it.Event = new(InterchainGasPaymasterGasOracleSet)
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
func (it *InterchainGasPaymasterGasOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainGasPaymasterGasOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainGasPaymasterGasOracleSet represents a GasOracleSet event raised by the InterchainGasPaymaster contract.
type InterchainGasPaymasterGasOracleSet struct {
	RemoteDomain uint32
	GasOracle    common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGasOracleSet is a free log retrieval operation binding the contract event 0x38b842e116ff320f6a79ba4cf434eee927fa156b226c65c90d62989de589b4e7.
//
// Solidity: event GasOracleSet(uint32 indexed remoteDomain, address gasOracle)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) FilterGasOracleSet(opts *bind.FilterOpts, remoteDomain []uint32) (*InterchainGasPaymasterGasOracleSetIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}

	logs, sub, err := _InterchainGasPaymaster.contract.FilterLogs(opts, "GasOracleSet", remoteDomainRule)
	if err != nil {
		return nil, err
	}
	return &InterchainGasPaymasterGasOracleSetIterator{contract: _InterchainGasPaymaster.contract, event: "GasOracleSet", logs: logs, sub: sub}, nil
}

// WatchGasOracleSet is a free log subscription operation binding the contract event 0x38b842e116ff320f6a79ba4cf434eee927fa156b226c65c90d62989de589b4e7.
//
// Solidity: event GasOracleSet(uint32 indexed remoteDomain, address gasOracle)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) WatchGasOracleSet(opts *bind.WatchOpts, sink chan<- *InterchainGasPaymasterGasOracleSet, remoteDomain []uint32) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}

	logs, sub, err := _InterchainGasPaymaster.contract.WatchLogs(opts, "GasOracleSet", remoteDomainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainGasPaymasterGasOracleSet)
				if err := _InterchainGasPaymaster.contract.UnpackLog(event, "GasOracleSet", log); err != nil {
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

// ParseGasOracleSet is a log parse operation binding the contract event 0x38b842e116ff320f6a79ba4cf434eee927fa156b226c65c90d62989de589b4e7.
//
// Solidity: event GasOracleSet(uint32 indexed remoteDomain, address gasOracle)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) ParseGasOracleSet(log types.Log) (*InterchainGasPaymasterGasOracleSet, error) {
	event := new(InterchainGasPaymasterGasOracleSet)
	if err := _InterchainGasPaymaster.contract.UnpackLog(event, "GasOracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainGasPaymasterGasPaymentIterator is returned from FilterGasPayment and is used to iterate over the raw logs and unpacked data for GasPayment events raised by the InterchainGasPaymaster contract.
type InterchainGasPaymasterGasPaymentIterator struct {
	Event *InterchainGasPaymasterGasPayment // Event containing the contract specifics and raw log

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
func (it *InterchainGasPaymasterGasPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainGasPaymasterGasPayment)
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
		it.Event = new(InterchainGasPaymasterGasPayment)
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
func (it *InterchainGasPaymasterGasPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainGasPaymasterGasPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainGasPaymasterGasPayment represents a GasPayment event raised by the InterchainGasPaymaster contract.
type InterchainGasPaymasterGasPayment struct {
	MessageId [32]byte
	GasAmount *big.Int
	Payment   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGasPayment is a free log retrieval operation binding the contract event 0xf715e66d2cd2a0179069dcda2a41a4da36a30e1aad12187fbb7c5d100afd621c.
//
// Solidity: event GasPayment(bytes32 indexed messageId, uint256 gasAmount, uint256 payment)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) FilterGasPayment(opts *bind.FilterOpts, messageId [][32]byte) (*InterchainGasPaymasterGasPaymentIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _InterchainGasPaymaster.contract.FilterLogs(opts, "GasPayment", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &InterchainGasPaymasterGasPaymentIterator{contract: _InterchainGasPaymaster.contract, event: "GasPayment", logs: logs, sub: sub}, nil
}

// WatchGasPayment is a free log subscription operation binding the contract event 0xf715e66d2cd2a0179069dcda2a41a4da36a30e1aad12187fbb7c5d100afd621c.
//
// Solidity: event GasPayment(bytes32 indexed messageId, uint256 gasAmount, uint256 payment)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) WatchGasPayment(opts *bind.WatchOpts, sink chan<- *InterchainGasPaymasterGasPayment, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _InterchainGasPaymaster.contract.WatchLogs(opts, "GasPayment", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainGasPaymasterGasPayment)
				if err := _InterchainGasPaymaster.contract.UnpackLog(event, "GasPayment", log); err != nil {
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

// ParseGasPayment is a log parse operation binding the contract event 0xf715e66d2cd2a0179069dcda2a41a4da36a30e1aad12187fbb7c5d100afd621c.
//
// Solidity: event GasPayment(bytes32 indexed messageId, uint256 gasAmount, uint256 payment)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) ParseGasPayment(log types.Log) (*InterchainGasPaymasterGasPayment, error) {
	event := new(InterchainGasPaymasterGasPayment)
	if err := _InterchainGasPaymaster.contract.UnpackLog(event, "GasPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainGasPaymasterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the InterchainGasPaymaster contract.
type InterchainGasPaymasterInitializedIterator struct {
	Event *InterchainGasPaymasterInitialized // Event containing the contract specifics and raw log

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
func (it *InterchainGasPaymasterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainGasPaymasterInitialized)
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
		it.Event = new(InterchainGasPaymasterInitialized)
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
func (it *InterchainGasPaymasterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainGasPaymasterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainGasPaymasterInitialized represents a Initialized event raised by the InterchainGasPaymaster contract.
type InterchainGasPaymasterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) FilterInitialized(opts *bind.FilterOpts) (*InterchainGasPaymasterInitializedIterator, error) {

	logs, sub, err := _InterchainGasPaymaster.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InterchainGasPaymasterInitializedIterator{contract: _InterchainGasPaymaster.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InterchainGasPaymasterInitialized) (event.Subscription, error) {

	logs, sub, err := _InterchainGasPaymaster.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainGasPaymasterInitialized)
				if err := _InterchainGasPaymaster.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) ParseInitialized(log types.Log) (*InterchainGasPaymasterInitialized, error) {
	event := new(InterchainGasPaymasterInitialized)
	if err := _InterchainGasPaymaster.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainGasPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InterchainGasPaymaster contract.
type InterchainGasPaymasterOwnershipTransferredIterator struct {
	Event *InterchainGasPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InterchainGasPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainGasPaymasterOwnershipTransferred)
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
		it.Event = new(InterchainGasPaymasterOwnershipTransferred)
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
func (it *InterchainGasPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainGasPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainGasPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the InterchainGasPaymaster contract.
type InterchainGasPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InterchainGasPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterchainGasPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InterchainGasPaymasterOwnershipTransferredIterator{contract: _InterchainGasPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InterchainGasPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InterchainGasPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainGasPaymasterOwnershipTransferred)
				if err := _InterchainGasPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InterchainGasPaymaster *InterchainGasPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*InterchainGasPaymasterOwnershipTransferred, error) {
	event := new(InterchainGasPaymasterOwnershipTransferred)
	if err := _InterchainGasPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
