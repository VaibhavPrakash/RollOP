// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const StorageGasOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/hyperlane/StorageGasOracle.sol:StorageGasOracle\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/hyperlane/StorageGasOracle.sol:StorageGasOracle\",\"label\":\"remoteGasData\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_uint32,t_struct(RemoteGasData)1002_storage)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_uint32,t_struct(RemoteGasData)1002_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint32 =\u003e struct IGasOracle.RemoteGasData)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint32\",\"value\":\"t_struct(RemoteGasData)1002_storage\"},\"t_struct(RemoteGasData)1002_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IGasOracle.RemoteGasData\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"}}}"

var StorageGasOracleStorageLayout = new(solc.StorageLayout)

var StorageGasOracleDeployedBin = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b14610133578063b08e56d01461015b578063f2fde38b146101a5578063f3a1495f146101b857600080fd5b806360fcef7c14610082578063698faffc14610116578063715018a61461012b575b600080fd5b6100e8610090366004610527565b63ffffffff166000908152600160209081526040918290208251808401909352546fffffffffffffffffffffffffffffffff808216808552700100000000000000000000000000000000909204169290910182905291565b604080516fffffffffffffffffffffffffffffffff9384168152929091166020830152015b60405180910390f35b610129610124366004610554565b6101cb565b005b610129610217565b60005460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010d565b6100e8610169366004610527565b6001602052600090815260409020546fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041682565b6101296101b33660046105c9565b61022b565b6101296101c63660046105ff565b6102e7565b6101d36102f8565b8060005b81811015610211576101ff8484838181106101f4576101f4610617565b905060600201610379565b8061020981610646565b9150506101d7565b50505050565b61021f6102f8565b61022960006104b2565b565b6102336102f8565b73ffffffffffffffffffffffffffffffffffffffff81166102db576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102e4816104b2565b50565b6102ef6102f8565b6102e481610379565b60005473ffffffffffffffffffffffffffffffffffffffff163314610229576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d2565b604051806040016040528082602001602081019061039791906106a5565b6fffffffffffffffffffffffffffffffff1681526020016103be60608401604085016106a5565b6fffffffffffffffffffffffffffffffff169052600160006103e36020850185610527565b63ffffffff1681526020808201929092526040016000208251928201516fffffffffffffffffffffffffffffffff9081167001000000000000000000000000000000000293169290921790915561043c90820182610527565b63ffffffff167fb48c1cb713397fc0c0649596c221270fec0b3de3f85ccf6a734411a2fe57a69461047360408401602085016106a5565b61048360608501604086016106a5565b604080516fffffffffffffffffffffffffffffffff93841681529290911660208301520160405180910390a250565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561053957600080fd5b813563ffffffff8116811461054d57600080fd5b9392505050565b6000806020838503121561056757600080fd5b823567ffffffffffffffff8082111561057f57600080fd5b818501915085601f83011261059357600080fd5b8135818111156105a257600080fd5b8660206060830285010111156105b757600080fd5b60209290920196919550909350505050565b6000602082840312156105db57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461054d57600080fd5b60006060828403121561061157600080fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361069e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b6000602082840312156106b757600080fd5b81356fffffffffffffffffffffffffffffffff8116811461054d57600080fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(StorageGasOracleStorageLayoutJSON), StorageGasOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["StorageGasOracle"] = StorageGasOracleStorageLayout
	deployedBytecodes["StorageGasOracle"] = StorageGasOracleDeployedBin
}