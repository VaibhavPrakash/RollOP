// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const InterchainGasPaymasterStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/hyperlane/InterchainGasPaymaster.sol:InterchainGasPaymaster\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/hyperlane/InterchainGasPaymaster.sol:InterchainGasPaymaster\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/hyperlane/InterchainGasPaymaster.sol:InterchainGasPaymaster\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"contracts/hyperlane/InterchainGasPaymaster.sol:InterchainGasPaymaster\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/hyperlane/InterchainGasPaymaster.sol:InterchainGasPaymaster\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"contracts/hyperlane/InterchainGasPaymaster.sol:InterchainGasPaymaster\",\"label\":\"gasOracles\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_mapping(t_uint32,t_contract(IGasOracle)1007)\"},{\"astId\":1006,\"contract\":\"contracts/hyperlane/InterchainGasPaymaster.sol:InterchainGasPaymaster\",\"label\":\"beneficiary\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(IGasOracle)1007\":{\"encoding\":\"inplace\",\"label\":\"contract IGasOracle\",\"numberOfBytes\":\"20\"},\"t_mapping(t_uint32,t_contract(IGasOracle)1007)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint32 =\u003e contract IGasOracle)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint32\",\"value\":\"t_contract(IGasOracle)1007\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var InterchainGasPaymasterStorageLayout = new(solc.StorageLayout)

var InterchainGasPaymasterDeployedBin = "0x6080604052600436106100c75760003560e01c80634e71d92d116100745780638da5cb5b1161004e5780638da5cb5b1461024e578063a692979314610279578063f2fde38b146102a757600080fd5b80634e71d92d146101db57806360fcef7c146101f0578063715018a61461023957600080fd5b80632f202650116100a55780632f2026501461016e57806338af3eed1461018e578063485cc955146101bb57600080fd5b806311bf2c18146100cc5780631c31f710146100e15780631d16c8c814610101575b600080fd5b6100df6100da366004610d9a565b6102c7565b005b3480156100ed57600080fd5b506100df6100fc366004610de0565b6104b2565b34801561010d57600080fd5b5061014461011c366004610e02565b60656020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561017a57600080fd5b506100df610189366004610e1d565b6104c6565b34801561019a57600080fd5b506066546101449073ffffffffffffffffffffffffffffffffffffffff1681565b3480156101c757600080fd5b506100df6101d6366004610e92565b61054c565b3480156101e757600080fd5b506100df6106f2565b3480156101fc57600080fd5b5061021061020b366004610e02565b6107bc565b604080516fffffffffffffffffffffffffffffffff938416815292909116602083015201610165565b34801561024557600080fd5b506100df6108ee565b34801561025a57600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610144565b34801561028557600080fd5b50610299610294366004610ec5565b610902565b604051908152602001610165565b3480156102b357600080fd5b506100df6102c2366004610de0565b61096b565b60006102d38484610902565b90508034101561036a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f696e73756666696369656e7420696e746572636861696e20676173207061796d60448201527f656e74000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b60006103768234610f1e565b9050801561046f5760008373ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d80600081146103d8576040519150601f19603f3d011682016040523d82523d6000602084013e6103dd565b606091505b505090508061046d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f496e746572636861696e20676173207061796d656e7420726566756e6420666160448201527f696c6564000000000000000000000000000000000000000000000000000000006064820152608401610361565b505b604080518581526020810184905287917ff715e66d2cd2a0179069dcda2a41a4da36a30e1aad12187fbb7c5d100afd621c910160405180910390a2505050505050565b6104ba610a1f565b6104c381610aa0565b50565b6104ce610a1f565b8060005b81811015610546576105348484838181106104ef576104ef610f35565b6105059260206040909202019081019150610e02565b85858481811061051757610517610f35565b905060400201602001602081019061052f9190610de0565b610b19565b8061053e81610f64565b9150506104d2565b50505050565b600054610100900460ff161580801561056c5750600054600160ff909116105b806105865750303b158015610586575060005460ff166001145b610612576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610361565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561067057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610678610ba7565b61068183610c46565b61068a82610aa0565b80156106ed57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60665460405160009173ffffffffffffffffffffffffffffffffffffffff169047908381818185875af1925050503d806000811461074c576040519150601f19603f3d011682016040523d82523d6000602084013e610751565b606091505b50509050806104c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f217472616e7366657200000000000000000000000000000000000000000000006044820152606401610361565b63ffffffff8116600090815260656020526040812054819073ffffffffffffffffffffffffffffffffffffffff1680610851576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f21676173206f7261636c650000000000000000000000000000000000000000006044820152606401610361565b6040517f60fcef7c00000000000000000000000000000000000000000000000000000000815263ffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8216906360fcef7c906024016040805180830381865afa1580156108c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e49190610fbc565b9250925050915091565b6108f6610a1f565b6109006000610c46565b565b6000806000610910856107bc565b909250905060006109336fffffffffffffffffffffffffffffffff831686610fe6565b90506402540be4006109576fffffffffffffffffffffffffffffffff851683610fe6565b6109619190611023565b9695505050505050565b610973610a1f565b73ffffffffffffffffffffffffffffffffffffffff8116610a16576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610361565b6104c381610c46565b60335473ffffffffffffffffffffffffffffffffffffffff163314610900576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610361565b606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f04d55a8be181fb8d75b76f2d48aa0b2ee40f47e53d6e61763eeeec46feea8a249060200160405180910390a150565b63ffffffff821660008181526065602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff861690811790915591519182527f38b842e116ff320f6a79ba4cf434eee927fa156b226c65c90d62989de589b4e7910160405180910390a25050565b600054610100900460ff16610c3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610361565b610900610cbd565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610d54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610361565b61090033610c46565b803563ffffffff81168114610d7157600080fd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610d7157600080fd5b60008060008060808587031215610db057600080fd5b84359350610dc060208601610d5d565b925060408501359150610dd560608601610d76565b905092959194509250565b600060208284031215610df257600080fd5b610dfb82610d76565b9392505050565b600060208284031215610e1457600080fd5b610dfb82610d5d565b60008060208385031215610e3057600080fd5b823567ffffffffffffffff80821115610e4857600080fd5b818501915085601f830112610e5c57600080fd5b813581811115610e6b57600080fd5b8660208260061b8501011115610e8057600080fd5b60209290920196919550909350505050565b60008060408385031215610ea557600080fd5b610eae83610d76565b9150610ebc60208401610d76565b90509250929050565b60008060408385031215610ed857600080fd5b610ee183610d5d565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015610f3057610f30610eef565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f9557610f95610eef565b5060010190565b80516fffffffffffffffffffffffffffffffff81168114610d7157600080fd5b60008060408385031215610fcf57600080fd5b610fd883610f9c565b9150610ebc60208401610f9c565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561101e5761101e610eef565b500290565b600082611059577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(InterchainGasPaymasterStorageLayoutJSON), InterchainGasPaymasterStorageLayout); err != nil {
		panic(err)
	}

	layouts["InterchainGasPaymaster"] = InterchainGasPaymasterStorageLayout
	deployedBytecodes["InterchainGasPaymaster"] = InterchainGasPaymasterDeployedBin
}