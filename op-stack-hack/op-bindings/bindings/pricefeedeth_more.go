// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const PriceFeedEthStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/pricefeeds/PriceFeedEth.sol:PriceFeedEth\",\"label\":\"currentPrice\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1001,\"contract\":\"contracts/pricefeeds/PriceFeedEth.sol:PriceFeedEth\",\"label\":\"dataFeed\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_uint64,t_uint64)\"}],\"types\":{\"t_mapping(t_uint64,t_uint64)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint64 =\u003e uint64)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint64\",\"value\":\"t_uint64\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"}}}"

var PriceFeedEthStorageLayout = new(solc.StorageLayout)

var PriceFeedEthDeployedBin = "0x608060405234801561001057600080fd5b506004361061004c5760003560e01c8063464be46f1461005157806398d5fdca1461009b5780639d1b464a146100ad578063aa415b52146100c1575b600080fd5b61007e61005f3660046101ed565b67ffffffffffffffff9081166000908152600160205260409020541690565b60405167ffffffffffffffff909116815260200160405180910390f35b60005467ffffffffffffffff1661007e565b60005461007e9067ffffffffffffffff1681565b6100d46100cf36600461020f565b6100d6565b005b73deaddeaddeaddeaddeaddeaddeaddeaddead0001331461017d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4f6e6c792073797374656d20616464726573732063616e2063616c6c2074686960448201527f732066756e6374696f6e00000000000000000000000000000000000000000000606482015260840160405180910390fd5b6000805467ffffffffffffffff9283167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009182168117835593909216815260016020526040902080549091169091179055565b803567ffffffffffffffff811681146101e857600080fd5b919050565b6000602082840312156101ff57600080fd5b610208826101d0565b9392505050565b6000806040838503121561022257600080fd5b61022b836101d0565b9150610239602084016101d0565b9050925092905056fea164736f6c6343000813000a"

func init() {
	if err := json.Unmarshal([]byte(PriceFeedEthStorageLayoutJSON), PriceFeedEthStorageLayout); err != nil {
		panic(err)
	}

	layouts["PriceFeedEth"] = PriceFeedEthStorageLayout
	deployedBytecodes["PriceFeedEth"] = PriceFeedEthDeployedBin
}
