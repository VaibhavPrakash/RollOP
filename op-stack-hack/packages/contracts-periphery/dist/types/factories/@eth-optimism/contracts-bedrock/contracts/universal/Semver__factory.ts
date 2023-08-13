/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Signer,
  utils,
  Contract,
  ContractFactory,
  BigNumberish,
  Overrides,
} from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../../common";
import type {
  Semver,
  SemverInterface,
} from "../../../../../@eth-optimism/contracts-bedrock/contracts/universal/Semver";

const _abi = [
  {
    inputs: [
      {
        internalType: "uint256",
        name: "_major",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "_minor",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "_patch",
        type: "uint256",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [],
    name: "version",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
];

const _bytecode =
  "0x60e060405234801561001057600080fd5b5060405161053138038061053183398101604081905261002f91610040565b60809290925260a05260c05261006e565b60008060006060848603121561005557600080fd5b8351925060208401519150604084015190509250925092565b60805160a05160c05161049761009a600039600060a701526000607e01526000605501526104976000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806354fd4d5014610030575b600080fd5b61003861004e565b6040516100459190610252565b60405180910390f35b60606100797f00000000000000000000000000000000000000000000000000000000000000006100f1565b6100a27f00000000000000000000000000000000000000000000000000000000000000006100f1565b6100cb7f00000000000000000000000000000000000000000000000000000000000000006100f1565b6040516020016100dd939291906102a3565b604051602081830303815290604052905090565b60608160000361013457505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561015e578061014881610348565b91506101579050600a836103af565b9150610138565b60008167ffffffffffffffff811115610179576101796103c3565b6040519080825280601f01601f1916602001820160405280156101a3576020820181803683370190505b5090505b8415610226576101b86001836103f2565b91506101c5600a8661040b565b6101d090603061041f565b60f81b8183815181106101e5576101e5610432565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061021f600a866103af565b94506101a7565b949350505050565b60005b83811015610249578181015183820152602001610231565b50506000910152565b602081526000825180602084015261027181604085016020870161022e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600084516102b581846020890161022e565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516102f1816001850160208a0161022e565b6001920191820152835161030c81600284016020880161022e565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361037957610379610319565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826103be576103be610380565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b8181038181111561040557610405610319565b92915050565b60008261041a5761041a610380565b500690565b8082018082111561040557610405610319565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea264697066735822122039c6f506608c84ee3eb35a3ce5301e0e90b93b8ce527d47b2ad8a73e96e4896f64736f6c63430008100033";

type SemverConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: SemverConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class Semver__factory extends ContractFactory {
  constructor(...args: SemverConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    _major: PromiseOrValue<BigNumberish>,
    _minor: PromiseOrValue<BigNumberish>,
    _patch: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<Semver> {
    return super.deploy(
      _major,
      _minor,
      _patch,
      overrides || {}
    ) as Promise<Semver>;
  }
  override getDeployTransaction(
    _major: PromiseOrValue<BigNumberish>,
    _minor: PromiseOrValue<BigNumberish>,
    _patch: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(_major, _minor, _patch, overrides || {});
  }
  override attach(address: string): Semver {
    return super.attach(address) as Semver;
  }
  override connect(signer: Signer): Semver__factory {
    return super.connect(signer) as Semver__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): SemverInterface {
    return new utils.Interface(_abi) as SemverInterface;
  }
  static connect(address: string, signerOrProvider: Signer | Provider): Semver {
    return new Contract(address, _abi, signerOrProvider) as Semver;
  }
}