// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"spacer_1_0_20\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1002,\"contract\":\"contracts/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1003,\"contract\":\"contracts/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_uint256)47_storage\"},{\"astId\":1004,\"contract\":\"contracts/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"withdrawalIsms\",\"offset\":0,\"slot\":\"50\",\"type\":\"t_mapping(t_address,t_address)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)47_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[47]\",\"numberOfBytes\":\"1504\",\"base\":\"t_uint256\"},\"t_mapping(t_address,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_address\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L1StandardBridgeStorageLayout = new(solc.StorageLayout)

var L1StandardBridgeDeployedBin = "0x6080604052600436106101b05760003560e01c80636465e69f116100ec578063927ede2d1161008a578063b1a1a88211610064578063b1a1a8821461060c578063e11013dd1461061f578063ea59e64a14610632578063f7e83aee1461066057600080fd5b8063927ede2d146105a55780639a2ac6d5146105d9578063a9f9e675146105ec57600080fd5b8063838b2520116100c6578063838b2520146104ec578063870876231461050c5780638f601f661461052c57806391c49bf81461057257600080fd5b80636465e69f1461045d5780637dc2b8fb146104845780637f46ddb2146104b857600080fd5b806325896ec011610159578063540abf7311610133578063540abf73146103db57806354fd4d50146103fb57806356d5d4751461041d57806358a997f61461043d57600080fd5b806325896ec0146103315780633cb747bf146103655780634d2b84f31461039857600080fd5b80631532ec341161018a5780631532ec34146102c157806315ce45a2146102d45780631635f5fd1461031e57600080fd5b80630166a07a1461026e578063047b10cf1461028e57806309fc8843146102ae57600080fd5b3661026957333b15610249576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b610267333362030d4060405180602001604052806000815250610690565b005b600080fd5b34801561027a57600080fd5b50610267610289366004612bc3565b6106a3565b34801561029a57600080fd5b506102676102a9366004612c5b565b6108cc565b6102676102bc366004612cad565b6109be565b6102676102cf366004612d00565b610a95565b3480156102e057600080fd5b506102f46102ef366004612d73565b610aa9565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61026761032c366004612d00565b610bda565b34801561033d57600080fd5b506102f47f000000000000000000000000000000000000000000000000000000000000000081565b34801561037157600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006102f4565b3480156103a457600080fd5b506102f46103b3366004612db5565b60326020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b3480156103e757600080fd5b506102676103f6366004612dd2565b6110af565b34801561040757600080fd5b506104106110f4565b6040516103159190612ebf565b34801561042957600080fd5b50610267610438366004612ed2565b611197565b34801561044957600080fd5b50610267610458366004612f2c565b61147c565b34801561046957600080fd5b50610472600181565b60405160ff9091168152602001610315565b34801561049057600080fd5b506102f47f000000000000000000000000000000000000000000000000000000000000000081565b3480156104c457600080fd5b506102f47f000000000000000000000000000000000000000000000000000000000000000081565b3480156104f857600080fd5b50610267610507366004612dd2565b611550565b34801561051857600080fd5b50610267610527366004612f2c565b611595565b34801561053857600080fd5b50610564610547366004612c5b565b600260209081526000928352604080842090915290825290205481565b604051908152602001610315565b34801561057e57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006102f4565b3480156105b157600080fd5b506102f47f000000000000000000000000000000000000000000000000000000000000000081565b6102676105e7366004612faf565b611669565b3480156105f857600080fd5b50610267610607366004612bc3565b6116ab565b61026761061a366004612cad565b6116ba565b61026761062d366004612faf565b61178b565b34801561063e57600080fd5b5061065261064d366004612d73565b6117ce565b604051610315929190612ffa565b34801561066c57600080fd5b5061068061067b366004613035565b61183d565b6040519015158152602001610315565b61069d84843485856118d3565b50505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156107c157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610785573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107a99190613095565b73ffffffffffffffffffffffffffffffffffffffff16145b610873576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a401610240565b61087f87878686611aab565b6108c3878787878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611c8592505050565b50505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461096b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f21464153545f5749544844524157414c5f4f574e4552000000000000000000006044820152606401610240565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260326020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001691909216179055565b333b15610a4d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610240565b610a903333348686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506118d392505050565b505050565b610aa28585858585610bda565b5050505050565b600080803063ea59e64a610abd8787611d13565b6040518363ffffffff1660e01b8152600401610ada9291906130fb565b600060405180830381865afa158015610af7573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610b3d919081019061313e565b909250905060007ffe995f86000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000841601610bcf57600082806020019051810190610ba39190613095565b73ffffffffffffffffffffffffffffffffffffffff908116600090815260326020526040902054169150505b925050505b92915050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148015610cf857507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce09190613095565b73ffffffffffffffffffffffffffffffffffffffff16145b610daa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a401610240565b823414610e39576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e742072657175697265640000000000006064820152608401610240565b3073ffffffffffffffffffffffffffffffffffffffff851603610ede576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c6600000000000000000000000000000000000000000000000000000000006064820152608401610240565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610fb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e6765720000000000000000000000000000000000000000000000006064820152608401610240565b610ffb85858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611d2e92505050565b6000611018855a8660405180602001604052806000815250611da1565b9050806110a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c656400000000000000000000000000000000000000000000000000000000006064820152608401610240565b505050505050565b6108c387873388888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611dbb92505050565b606061111f7f0000000000000000000000000000000000000000000000000000000000000000611f23565b6111487f0000000000000000000000000000000000000000000000000000000000000000611f23565b6111717f0000000000000000000000000000000000000000000000000000000000000000611f23565b60405160200161118393929190613244565b604051602081830303815290604052905090565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614611236576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f214d41494c424f580000000000000000000000000000000000000000000000006044820152606401610240565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168373ffffffffffffffffffffffffffffffffffffffff16146112eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f214f544845525f425249444745000000000000000000000000000000000000006044820152606401610240565b6040517fea59e64a0000000000000000000000000000000000000000000000000000000081526000908190309063ea59e64a9061132e90879087906004016130fb565b600060405180830381865afa15801561134b573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611391919081019061313e565b90925090507fffffffff0000000000000000000000000000000000000000000000000000000082167f0166a07a0000000000000000000000000000000000000000000000000000000014611441576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f696e76616c69642073656c6563746f72000000000000000000000000000000006044820152606401610240565b6000806000808480602001905181019061145b91906132ba565b94509450509350935061147084848484611aab565b50505050505050505050565b333b1561150b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610240565b6110a786863333888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061206092505050565b6108c387873388888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061206092505050565b333b15611624576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610240565b6110a786863333888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611dbb92505050565b61069d33858585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061069092505050565b6108c3878787878787876106a3565b333b15611749576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610240565b610a9033338585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061069092505050565b61069d3385348686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506118d392505050565b60006060816117e06004828688613325565b6117e99161334f565b905060006117fa8560048189613325565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250949650919450505050505b9250929050565b60006118498383610aa9565b73ffffffffffffffffffffffffffffffffffffffff1663f7e83aee868686866040518563ffffffff1660e01b81526004016118879493929190613397565b6020604051808303816000875af11580156118a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118ca91906133be565b95945050505050565b823414611962576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c756500006064820152608401610240565b61196e8585858461206f565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b847f0000000000000000000000000000000000000000000000000000000000000000631635f5fd60e01b898989886040516024016119eb94939291906133e0565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b9092168252611a7e92918890600401613429565b6000604051808303818588803b158015611a9757600080fd5b505af1158015611470573d6000803e3d6000fd5b611ab4846120e2565b15611c0357611ac3848461213e565b611b75576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a401610240565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390528516906340c10f19906044015b600060405180830381600087803b158015611be657600080fd5b505af1158015611bfa573d6000803e3d6000fd5b5050505061069d565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260026020908152604080832093871683529290522054611c4190829061349d565b73ffffffffffffffffffffffffffffffffffffffff80861660008181526002602090815260408083209489168352939052919091209190915561069d90838361225e565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3868686604051611cfd939291906134b4565b60405180910390a46110a7868686868686612332565b366000611d2383604d8187613325565b915091509250929050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e6318484604051611d8d9291906134e9565b60405180910390a361069d848484846123ba565b600080600080845160208601878a8af19695505050505050565b611dc787878786612427565b611dd58787878787866125e5565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b7f0000000000000000000000000000000000000000000000000000000000000000630166a07a60e01b898b8a8a8a89604051602401611e5596959493929190613502565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b9092168252611ee892918790600401613429565b600060405180830381600087803b158015611f0257600080fd5b505af1158015611f16573d6000803e3d6000fd5b5050505050505050505050565b606081600003611f6657505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611f905780611f7a8161355d565b9150611f899050600a836135c4565b9150611f6a565b60008167ffffffffffffffff811115611fab57611fab61310f565b6040519080825280601f01601f191660200182016040528015611fd5576020820181803683370190505b5090505b841561205857611fea60018361349d565b9150611ff7600a866135d8565b6120029060306135ec565b60f81b81838151811061201757612017613604565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350612051600a866135c4565b9450611fd9565b949350505050565b6108c387878787878787611dbb565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f2384846040516120ce9291906134e9565b60405180910390a361069d84848484612673565b600061210e827f1d1d8b63000000000000000000000000000000000000000000000000000000006126d2565b80610bd45750610bd4827fec4fc8e3000000000000000000000000000000000000000000000000000000006126d2565b600061216a837f1d1d8b63000000000000000000000000000000000000000000000000000000006126d2565b15612213578273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121de9190613095565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050610bd4565b8273ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121ba573d6000803e3d6000fd5b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610a909084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526126f5565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd8686866040516123aa939291906134b4565b60405180910390a4505050505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d84846040516124199291906134e9565b60405180910390a350505050565b612430846120e2565b1561254c5761243f848461213e565b6124f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a401610240565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff838116600483015260248201839052851690639dc29fac90604401611bcc565b61256e73ffffffffffffffffffffffffffffffffffffffff8516833084612801565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600260209081526040808320938716835292905220546125ac9082906135ec565b73ffffffffffffffffffffffffffffffffffffffff80861660009081526002602090815260408083209388168352929052205550505050565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d039686868660405161265d939291906134b4565b60405180910390a46110a786868686868661285f565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af584846040516124199291906134e9565b60006126dd836128d7565b80156126ee57506126ee838361293b565b9392505050565b6000612757826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16612a0a9092919063ffffffff16565b805190915015610a90578080602001905181019061277591906133be565b610a90576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610240565b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261069d9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016122b0565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf8686866040516123aa939291906134b4565b6000612903827f01ffc9a70000000000000000000000000000000000000000000000000000000061293b565b8015610bd45750612934827fffffffff0000000000000000000000000000000000000000000000000000000061293b565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d915060005190508280156129f3575060208210155b80156129ff5750600081115b979650505050505050565b606061205884846000858573ffffffffffffffffffffffffffffffffffffffff85163b612a93576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610240565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051612abc9190613633565b60006040518083038185875af1925050503d8060008114612af9576040519150601f19603f3d011682016040523d82523d6000602084013e612afe565b606091505b50915091506129ff82828660608315612b185750816126ee565b825115612b285782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102409190612ebf565b73ffffffffffffffffffffffffffffffffffffffff81168114612b7e57600080fd5b50565b60008083601f840112612b9357600080fd5b50813567ffffffffffffffff811115612bab57600080fd5b60208301915083602082850101111561183657600080fd5b600080600080600080600060c0888a031215612bde57600080fd5b8735612be981612b5c565b96506020880135612bf981612b5c565b95506040880135612c0981612b5c565b94506060880135612c1981612b5c565b93506080880135925060a088013567ffffffffffffffff811115612c3c57600080fd5b612c488a828b01612b81565b989b979a50959850939692959293505050565b60008060408385031215612c6e57600080fd5b8235612c7981612b5c565b91506020830135612c8981612b5c565b809150509250929050565b803563ffffffff81168114612ca857600080fd5b919050565b600080600060408486031215612cc257600080fd5b612ccb84612c94565b9250602084013567ffffffffffffffff811115612ce757600080fd5b612cf386828701612b81565b9497909650939450505050565b600080600080600060808688031215612d1857600080fd5b8535612d2381612b5c565b94506020860135612d3381612b5c565b935060408601359250606086013567ffffffffffffffff811115612d5657600080fd5b612d6288828901612b81565b969995985093965092949392505050565b60008060208385031215612d8657600080fd5b823567ffffffffffffffff811115612d9d57600080fd5b612da985828601612b81565b90969095509350505050565b600060208284031215612dc757600080fd5b81356126ee81612b5c565b600080600080600080600060c0888a031215612ded57600080fd5b8735612df881612b5c565b96506020880135612e0881612b5c565b95506040880135612e1881612b5c565b945060608801359350612e2d60808901612c94565b925060a088013567ffffffffffffffff811115612c3c57600080fd5b60005b83811015612e64578181015183820152602001612e4c565b8381111561069d5750506000910152565b60008151808452612e8d816020860160208601612e49565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006126ee6020830184612e75565b60008060008060608587031215612ee857600080fd5b612ef185612c94565b935060208501359250604085013567ffffffffffffffff811115612f1457600080fd5b612f2087828801612b81565b95989497509550505050565b60008060008060008060a08789031215612f4557600080fd5b8635612f5081612b5c565b95506020870135612f6081612b5c565b945060408701359350612f7560608801612c94565b9250608087013567ffffffffffffffff811115612f9157600080fd5b612f9d89828a01612b81565b979a9699509497509295939492505050565b60008060008060608587031215612fc557600080fd5b8435612fd081612b5c565b9350612fde60208601612c94565b9250604085013567ffffffffffffffff811115612f1457600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000831681526040602082015260006120586040830184612e75565b6000806000806040858703121561304b57600080fd5b843567ffffffffffffffff8082111561306357600080fd5b61306f88838901612b81565b9096509450602087013591508082111561308857600080fd5b50612f2087828801612b81565b6000602082840312156130a757600080fd5b81516126ee81612b5c565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6020815260006120586020830184866130b2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806040838503121561315157600080fd5b82517fffffffff000000000000000000000000000000000000000000000000000000008116811461318157600080fd5b602084015190925067ffffffffffffffff8082111561319f57600080fd5b818501915085601f8301126131b357600080fd5b8151818111156131c5576131c561310f565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561320b5761320b61310f565b8160405282815288602084870101111561322457600080fd5b613235836020830160208801612e49565b80955050505050509250929050565b60008451613256818460208901612e49565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551613292816001850160208a01612e49565b600192019182015283516132ad816002840160208801612e49565b0160020195945050505050565b600080600080600060a086880312156132d257600080fd5b85516132dd81612b5c565b60208701519095506132ee81612b5c565b60408701519094506132ff81612b5c565b606087015190935061331081612b5c565b80925050608086015190509295509295909350565b6000808585111561333557600080fd5b8386111561334257600080fd5b5050820193919092039150565b7fffffffff00000000000000000000000000000000000000000000000000000000813581811691600485101561338f5780818660040360031b1b83161692505b505092915050565b6040815260006133ab6040830186886130b2565b82810360208401526129ff8185876130b2565b6000602082840312156133d057600080fd5b815180151581146126ee57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152508360408301526080606083015261341f6080830184612e75565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006134586060830185612e75565b905063ffffffff83166040830152949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156134af576134af61346e565b500390565b73ffffffffffffffffffffffffffffffffffffffff841681528260208201526060604082015260006118ca6060830184612e75565b8281526040602082015260006120586040830184612e75565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a083015261355160c0830184612e75565b98975050505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361358e5761358e61346e565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826135d3576135d3613595565b500490565b6000826135e7576135e7613595565b500690565b600082198211156135ff576135ff61346e565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008251613645818460208701612e49565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1StandardBridgeStorageLayoutJSON), L1StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1StandardBridge"] = L1StandardBridgeStorageLayout
	deployedBytecodes["L1StandardBridge"] = L1StandardBridgeDeployedBin
}