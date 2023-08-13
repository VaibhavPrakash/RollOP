/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  IStaticERC1967Proxy,
  IStaticERC1967ProxyInterface,
} from "../../../../../../@eth-optimism/contracts-bedrock/contracts/universal/ProxyAdmin.sol/IStaticERC1967Proxy";

const _abi = [
  {
    inputs: [],
    name: "admin",
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
    name: "implementation",
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
];

export class IStaticERC1967Proxy__factory {
  static readonly abi = _abi;
  static createInterface(): IStaticERC1967ProxyInterface {
    return new utils.Interface(_abi) as IStaticERC1967ProxyInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IStaticERC1967Proxy {
    return new Contract(address, _abi, signerOrProvider) as IStaticERC1967Proxy;
  }
}