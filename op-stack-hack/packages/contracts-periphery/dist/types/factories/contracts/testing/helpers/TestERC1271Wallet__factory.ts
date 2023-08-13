/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../common";
import type {
  TestERC1271Wallet,
  TestERC1271WalletInterface,
} from "../../../../contracts/testing/helpers/TestERC1271Wallet";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "originalOwner",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "previousOwner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "OwnershipTransferred",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "bytes32",
        name: "hash",
        type: "bytes32",
      },
      {
        internalType: "bytes",
        name: "signature",
        type: "bytes",
      },
    ],
    name: "isValidSignature",
    outputs: [
      {
        internalType: "bytes4",
        name: "magicValue",
        type: "bytes4",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "owner",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "renounceOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "newOwner",
        type: "address",
      },
    ],
    name: "transferOwnership",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
];

const _bytecode =
  "0x608060405234801561001057600080fd5b50604051610a3c380380610a3c83398101604081905261002f91610171565b61003833610047565b61004181610097565b506101a1565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61009f610115565b6001600160a01b0381166101095760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b61011281610047565b50565b6000546001600160a01b0316331461016f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610100565b565b60006020828403121561018357600080fd5b81516001600160a01b038116811461019a57600080fd5b9392505050565b61088c806101b06000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80631626ba7e14610051578063715018a61461009a5780638da5cb5b146100a4578063f2fde38b146100cc575b600080fd5b61006461005f366004610718565b6100df565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b6100a261014e565b005b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610091565b6100a26100da3660046107f1565b610162565b6000805473ffffffffffffffffffffffffffffffffffffffff16610103848461021e565b73ffffffffffffffffffffffffffffffffffffffff1614610125576000610147565b7f1626ba7e000000000000000000000000000000000000000000000000000000005b9392505050565b610156610242565b61016060006102c3565b565b61016a610242565b73ffffffffffffffffffffffffffffffffffffffff8116610212576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b61021b816102c3565b50565b600080600061022d8585610338565b9150915061023a8161037d565b509392505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610160576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610209565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600080825160410361036e5760208301516040840151606085015160001a610362878285856105d1565b94509450505050610376565b506000905060025b9250929050565b600081600481111561039157610391610827565b036103995750565b60018160048111156103ad576103ad610827565b03610414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610209565b600281600481111561042857610428610827565b0361048f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610209565b60038160048111156104a3576104a3610827565b03610530576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610209565b600481600481111561054457610544610827565b0361021b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610209565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561060857506000905060036106e0565b8460ff16601b1415801561062057508460ff16601c14155b1561063157506000905060046106e0565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610685573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166106d9576000600192509250506106e0565b9150600090505b94509492505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806040838503121561072b57600080fd5b82359150602083013567ffffffffffffffff8082111561074a57600080fd5b818501915085601f83011261075e57600080fd5b813581811115610770576107706106e9565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156107b6576107b66106e9565b816040528281528860208487010111156107cf57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60006020828403121561080357600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461014757600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea26469706673582212205f0a98f484bc060d2a0567a58a6dfbbc0f0823fc53e548dd945bd1e6ad292c4f64736f6c634300080f0033";

type TestERC1271WalletConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: TestERC1271WalletConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class TestERC1271Wallet__factory extends ContractFactory {
  constructor(...args: TestERC1271WalletConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    originalOwner: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<TestERC1271Wallet> {
    return super.deploy(
      originalOwner,
      overrides || {}
    ) as Promise<TestERC1271Wallet>;
  }
  override getDeployTransaction(
    originalOwner: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(originalOwner, overrides || {});
  }
  override attach(address: string): TestERC1271Wallet {
    return super.attach(address) as TestERC1271Wallet;
  }
  override connect(signer: Signer): TestERC1271Wallet__factory {
    return super.connect(signer) as TestERC1271Wallet__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): TestERC1271WalletInterface {
    return new utils.Interface(_abi) as TestERC1271WalletInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): TestERC1271Wallet {
    return new Contract(address, _abi, signerOrProvider) as TestERC1271Wallet;
  }
}