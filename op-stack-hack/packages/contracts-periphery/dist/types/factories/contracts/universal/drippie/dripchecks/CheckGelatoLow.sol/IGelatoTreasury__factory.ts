/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type {
  IGelatoTreasury,
  IGelatoTreasuryInterface,
} from "../../../../../../contracts/universal/drippie/dripchecks/CheckGelatoLow.sol/IGelatoTreasury";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "_user",
        type: "address",
      },
      {
        internalType: "address",
        name: "_token",
        type: "address",
      },
    ],
    name: "userTokenBalance",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
];

export class IGelatoTreasury__factory {
  static readonly abi = _abi;
  static createInterface(): IGelatoTreasuryInterface {
    return new utils.Interface(_abi) as IGelatoTreasuryInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IGelatoTreasury {
    return new Contract(address, _abi, signerOrProvider) as IGelatoTreasury;
  }
}