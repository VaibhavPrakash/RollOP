"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.StandardBridgeAdapter = void 0;
const ethers_1 = require("ethers");
const contracts_1 = require("@eth-optimism/contracts");
const core_utils_1 = require("@eth-optimism/core-utils");
const L1StandardBridge_json_1 = __importDefault(require("@eth-optimism/contracts-bedrock/forge-artifacts/L1StandardBridge.sol/L1StandardBridge.json"));
const L2StandardBridge_json_1 = __importDefault(require("@eth-optimism/contracts-bedrock/forge-artifacts/L2StandardBridge.sol/L2StandardBridge.json"));
const OptimismMintableERC20_json_1 = __importDefault(require("@eth-optimism/contracts-bedrock/forge-artifacts/OptimismMintableERC20.sol/OptimismMintableERC20.json"));
const interfaces_1 = require("../interfaces");
const utils_1 = require("../utils");
class StandardBridgeAdapter {
    constructor(opts) {
        this.populateTransaction = {
            approve: async (l1Token, l2Token, amount, opts) => {
                if (!(await this.supportsTokenPair(l1Token, l2Token))) {
                    throw new Error(`token pair not supported by bridge`);
                }
                const token = new ethers_1.Contract((0, utils_1.toAddress)(l1Token), OptimismMintableERC20_json_1.default.abi, this.messenger.l1Provider);
                return token.populateTransaction.approve(this.l1Bridge.address, amount, (opts === null || opts === void 0 ? void 0 : opts.overrides) || {});
            },
            deposit: async (l1Token, l2Token, amount, opts) => {
                if (!(await this.supportsTokenPair(l1Token, l2Token))) {
                    throw new Error(`token pair not supported by bridge`);
                }
                if ((opts === null || opts === void 0 ? void 0 : opts.recipient) === undefined) {
                    return this.l1Bridge.populateTransaction.depositERC20((0, utils_1.toAddress)(l1Token), (0, utils_1.toAddress)(l2Token), amount, (opts === null || opts === void 0 ? void 0 : opts.l2GasLimit) || 200000, '0x', (opts === null || opts === void 0 ? void 0 : opts.overrides) || {});
                }
                else {
                    return this.l1Bridge.populateTransaction.depositERC20To((0, utils_1.toAddress)(l1Token), (0, utils_1.toAddress)(l2Token), (0, utils_1.toAddress)(opts.recipient), amount, (opts === null || opts === void 0 ? void 0 : opts.l2GasLimit) || 200000, '0x', (opts === null || opts === void 0 ? void 0 : opts.overrides) || {});
                }
            },
            withdraw: async (l1Token, l2Token, amount, opts) => {
                if (!(await this.supportsTokenPair(l1Token, l2Token))) {
                    throw new Error(`token pair not supported by bridge`);
                }
                if ((opts === null || opts === void 0 ? void 0 : opts.recipient) === undefined) {
                    return this.l2Bridge.populateTransaction.withdraw((0, utils_1.toAddress)(l2Token), amount, 0, '0x', (opts === null || opts === void 0 ? void 0 : opts.overrides) || {});
                }
                else {
                    return this.l2Bridge.populateTransaction.withdrawTo((0, utils_1.toAddress)(l2Token), (0, utils_1.toAddress)(opts.recipient), amount, 0, '0x', (opts === null || opts === void 0 ? void 0 : opts.overrides) || {});
                }
            },
        };
        this.estimateGas = {
            approve: async (l1Token, l2Token, amount, opts) => {
                return this.messenger.l1Provider.estimateGas(await this.populateTransaction.approve(l1Token, l2Token, amount, opts));
            },
            deposit: async (l1Token, l2Token, amount, opts) => {
                return this.messenger.l1Provider.estimateGas(await this.populateTransaction.deposit(l1Token, l2Token, amount, opts));
            },
            withdraw: async (l1Token, l2Token, amount, opts) => {
                return this.messenger.l2Provider.estimateGas(await this.populateTransaction.withdraw(l1Token, l2Token, amount, opts));
            },
        };
        this.messenger = opts.messenger;
        this.l1Bridge = new ethers_1.Contract((0, utils_1.toAddress)(opts.l1Bridge), L1StandardBridge_json_1.default.abi, this.messenger.l1Provider);
        this.l2Bridge = new ethers_1.Contract((0, utils_1.toAddress)(opts.l2Bridge), L2StandardBridge_json_1.default.abi, this.messenger.l2Provider);
    }
    async getDepositsByAddress(address, opts) {
        const events = await this.l1Bridge.queryFilter(this.l1Bridge.filters.ERC20DepositInitiated(undefined, undefined, address), opts === null || opts === void 0 ? void 0 : opts.fromBlock, opts === null || opts === void 0 ? void 0 : opts.toBlock);
        return events
            .filter((event) => {
            return (!(0, core_utils_1.hexStringEquals)(event.args.l1Token, ethers_1.ethers.constants.AddressZero) &&
                !(0, core_utils_1.hexStringEquals)(event.args.l2Token, contracts_1.predeploys.OVM_ETH));
        })
            .map((event) => {
            return {
                direction: interfaces_1.MessageDirection.L1_TO_L2,
                from: event.args.from,
                to: event.args.to,
                l1Token: event.args.l1Token,
                l2Token: event.args.l2Token,
                amount: event.args.amount,
                data: event.args.extraData,
                logIndex: event.logIndex,
                blockNumber: event.blockNumber,
                transactionHash: event.transactionHash,
            };
        })
            .sort((a, b) => {
            return b.blockNumber - a.blockNumber;
        });
    }
    async getWithdrawalsByAddress(address, opts) {
        const events = await this.l2Bridge.queryFilter(this.l2Bridge.filters.WithdrawalInitiated(undefined, undefined, address), opts === null || opts === void 0 ? void 0 : opts.fromBlock, opts === null || opts === void 0 ? void 0 : opts.toBlock);
        return events
            .filter((event) => {
            return (!(0, core_utils_1.hexStringEquals)(event.args.l1Token, ethers_1.ethers.constants.AddressZero) &&
                !(0, core_utils_1.hexStringEquals)(event.args.l2Token, contracts_1.predeploys.OVM_ETH));
        })
            .map((event) => {
            return {
                direction: interfaces_1.MessageDirection.L2_TO_L1,
                from: event.args.from,
                to: event.args.to,
                l1Token: event.args.l1Token,
                l2Token: event.args.l2Token,
                amount: event.args.amount,
                data: event.args.extraData,
                logIndex: event.logIndex,
                blockNumber: event.blockNumber,
                transactionHash: event.transactionHash,
            };
        })
            .sort((a, b) => {
            return b.blockNumber - a.blockNumber;
        });
    }
    async supportsTokenPair(l1Token, l2Token) {
        var _a, _b;
        try {
            const contract = new ethers_1.Contract((0, utils_1.toAddress)(l2Token), OptimismMintableERC20_json_1.default.abi, this.messenger.l2Provider);
            if ((0, core_utils_1.hexStringEquals)((0, utils_1.toAddress)(l1Token), ethers_1.ethers.constants.AddressZero) ||
                (0, core_utils_1.hexStringEquals)((0, utils_1.toAddress)(l2Token), contracts_1.predeploys.OVM_ETH)) {
                return false;
            }
            const remoteL1Token = await contract.l1Token();
            if (!(0, core_utils_1.hexStringEquals)(remoteL1Token, (0, utils_1.toAddress)(l1Token))) {
                return false;
            }
            const remoteL2Bridge = await contract.l2Bridge();
            if (!(0, core_utils_1.hexStringEquals)(remoteL2Bridge, this.l2Bridge.address)) {
                return false;
            }
            return true;
        }
        catch (err) {
            if (!((_a = err === null || err === void 0 ? void 0 : err.message) === null || _a === void 0 ? void 0 : _a.toString().includes('CALL_EXCEPTION')) &&
                !((_b = err === null || err === void 0 ? void 0 : err.stack) === null || _b === void 0 ? void 0 : _b.toString().includes('execution reverted'))) {
                console.error('Unexpected error when checking bridge', err);
            }
            return false;
        }
    }
    async approval(l1Token, l2Token, signer) {
        if (!(await this.supportsTokenPair(l1Token, l2Token))) {
            throw new Error(`token pair not supported by bridge`);
        }
        const token = new ethers_1.Contract((0, utils_1.toAddress)(l1Token), OptimismMintableERC20_json_1.default.abi, this.messenger.l1Provider);
        return token.allowance(await signer.getAddress(), this.l1Bridge.address);
    }
    async approve(l1Token, l2Token, amount, signer, opts) {
        return signer.sendTransaction(await this.populateTransaction.approve(l1Token, l2Token, amount, opts));
    }
    async deposit(l1Token, l2Token, amount, signer, opts) {
        return signer.sendTransaction(await this.populateTransaction.deposit(l1Token, l2Token, amount, opts));
    }
    async withdraw(l1Token, l2Token, amount, signer, opts) {
        return signer.sendTransaction(await this.populateTransaction.withdraw(l1Token, l2Token, amount, opts));
    }
}
exports.StandardBridgeAdapter = StandardBridgeAdapter;
//# sourceMappingURL=standard-bridge.js.map