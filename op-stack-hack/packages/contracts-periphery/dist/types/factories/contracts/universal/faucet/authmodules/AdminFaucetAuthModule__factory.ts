/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import type { Provider, TransactionRequest } from "@ethersproject/providers";
import type { PromiseOrValue } from "../../../../../common";
import type {
  AdminFaucetAuthModule,
  AdminFaucetAuthModuleInterface,
} from "../../../../../contracts/universal/faucet/authmodules/AdminFaucetAuthModule";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "_admin",
        type: "address",
      },
      {
        internalType: "string",
        name: "_name",
        type: "string",
      },
      {
        internalType: "string",
        name: "_version",
        type: "string",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [],
    name: "ADMIN",
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
    name: "PROOF_TYPEHASH",
    outputs: [
      {
        internalType: "bytes32",
        name: "",
        type: "bytes32",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: "address payable",
            name: "recipient",
            type: "address",
          },
          {
            internalType: "bytes32",
            name: "nonce",
            type: "bytes32",
          },
        ],
        internalType: "struct Faucet.DripParameters",
        name: "_params",
        type: "tuple",
      },
      {
        internalType: "bytes32",
        name: "_id",
        type: "bytes32",
      },
      {
        internalType: "bytes",
        name: "_proof",
        type: "bytes",
      },
    ],
    name: "verify",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
];

const _bytecode =
  "0x61016060405234801561001157600080fd5b50604051610bc5380380610bc583398101604081905261003091610193565b815160209283012081519183019190912060e08290526101008190524660a0818152604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818801819052818301969096526060810194909452608080850193909352308483018190528151808603909301835260c094850190915281519190950120905291909152610120526001600160a01b031661014052610215565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126100f857600080fd5b81516001600160401b0380821115610112576101126100d1565b604051601f8301601f19908116603f0116810190828211818310171561013a5761013a6100d1565b8160405283815260209250868385880101111561015657600080fd5b600091505b83821015610178578582018301518183018401529082019061015b565b838211156101895760008385830101525b9695505050505050565b6000806000606084860312156101a857600080fd5b83516001600160a01b03811681146101bf57600080fd5b60208501519093506001600160401b03808211156101dc57600080fd5b6101e8878388016100e7565b935060408601519150808211156101fe57600080fd5b5061020b868287016100e7565b9150509250925092565b60805160a05160c05160e05161010051610120516101405161095161027460003960008181604b015260f601526000610481015260006104d0015260006104ab015260006104040152600061042e0152600061045801526109516000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632a0acc6a146100465780638b3e3bf614610097578063f5431ffa146100cc575b600080fd5b61006d7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100be7fd4283507dc7a8282faa6b4c8c18bacbb74dbbab5467342e6f581656f3577236e81565b60405190815260200161008e565b6100df6100da366004610722565b6100ef565b604051901515815260200161008e565b60006101a47f000000000000000000000000000000000000000000000000000000000000000061019e7fd4283507dc7a8282faa6b4c8c18bacbb74dbbab5467342e6f581656f3577236e8760000151886020015188604051602001610183949392919093845273ffffffffffffffffffffffffffffffffffffffff9290921660208401526040830152606082015260800190565b604051602081830303815290604052805190602001206101ae565b8461021d565b90505b9392505050565b60006102176101bb6103ea565b836040517f19010000000000000000000000000000000000000000000000000000000000006020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b92915050565b600080600061022c858561051e565b909250905060008160048111156102455761024561082f565b14801561027d57508573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b1561028d576001925050506101a7565b6000808773ffffffffffffffffffffffffffffffffffffffff16631626ba7e60e01b88886040516024016102c292919061088e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317909252905161034b91906108e6565b600060405180830381855afa9150503d8060008114610386576040519150601f19603f3d011682016040523d82523d6000602084013e61038b565b606091505b509150915081801561039e575080516020145b80156103de575080517f1626ba7e00000000000000000000000000000000000000000000000000000000906103dc9083016020908101908401610902565b145b98975050505050505050565b60003073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561045057507f000000000000000000000000000000000000000000000000000000000000000046145b1561047a57507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b60008082516041036105545760208301516040840151606085015160001a61054887828585610563565b9450945050505061055c565b506000905060025b9250929050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561059a5750600090506003610672565b8460ff16601b141580156105b257508460ff16601c14155b156105c35750600090506004610672565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610617573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661066b57600060019250925050610672565b9150600090505b94509492505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156106cd576106cd61067b565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561071a5761071a61067b565b604052919050565b6000806000838503608081121561073857600080fd5b604081121561074657600080fd5b5061074f6106aa565b843573ffffffffffffffffffffffffffffffffffffffff8116811461077357600080fd5b8152602085810135818301529093506040850135925060608501359067ffffffffffffffff808311156107a557600080fd5b828701925087601f8401126107b957600080fd5b8235818111156107cb576107cb61067b565b6107fb837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016106d3565b9150808252888382860101111561081157600080fd5b80838501848401376000838284010152508093505050509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60005b83811015610879578181015183820152602001610861565b83811115610888576000848401525b50505050565b82815260406020820152600082518060408401526108b381606085016020870161085e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016060019392505050565b600082516108f881846020870161085e565b9190910192915050565b60006020828403121561091457600080fd5b505191905056fea26469706673582212204191f325ab3a408bb3773f1cd5357025665003c18d947925504b17f0ece6705364736f6c634300080f0033";

type AdminFaucetAuthModuleConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: AdminFaucetAuthModuleConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class AdminFaucetAuthModule__factory extends ContractFactory {
  constructor(...args: AdminFaucetAuthModuleConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override deploy(
    _admin: PromiseOrValue<string>,
    _name: PromiseOrValue<string>,
    _version: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<AdminFaucetAuthModule> {
    return super.deploy(
      _admin,
      _name,
      _version,
      overrides || {}
    ) as Promise<AdminFaucetAuthModule>;
  }
  override getDeployTransaction(
    _admin: PromiseOrValue<string>,
    _name: PromiseOrValue<string>,
    _version: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(_admin, _name, _version, overrides || {});
  }
  override attach(address: string): AdminFaucetAuthModule {
    return super.attach(address) as AdminFaucetAuthModule;
  }
  override connect(signer: Signer): AdminFaucetAuthModule__factory {
    return super.connect(signer) as AdminFaucetAuthModule__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): AdminFaucetAuthModuleInterface {
    return new utils.Interface(_abi) as AdminFaucetAuthModuleInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): AdminFaucetAuthModule {
    return new Contract(
      address,
      _abi,
      signerOrProvider
    ) as AdminFaucetAuthModule;
  }
}