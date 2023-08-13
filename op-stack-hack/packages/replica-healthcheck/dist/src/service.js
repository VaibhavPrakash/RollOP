"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.HealthcheckService = void 0;
const common_ts_1 = require("@eth-optimism/common-ts");
const core_utils_1 = require("@eth-optimism/core-utils");
const package_json_1 = require("../package.json");
class HealthcheckService extends common_ts_1.BaseServiceV2 {
    constructor(options) {
        super({
            version: package_json_1.version,
            name: 'healthcheck',
            options: Object.assign({ loopIntervalMs: 5000 }, options),
            optionsSpec: {
                referenceRpcProvider: {
                    validator: common_ts_1.validators.provider,
                    desc: 'Provider for interacting with L1',
                },
                targetRpcProvider: {
                    validator: common_ts_1.validators.provider,
                    desc: 'Provider for interacting with L2',
                },
                onDivergenceWaitMs: {
                    validator: common_ts_1.validators.num,
                    desc: 'Waiting time in ms per loop when divergence is detected',
                    default: 60000,
                    public: true,
                },
            },
            metricsSpec: {
                lastMatchingStateRootHeight: {
                    type: common_ts_1.Gauge,
                    desc: 'Highest matching state root between target and reference',
                },
                isCurrentlyDiverged: {
                    type: common_ts_1.Gauge,
                    desc: 'Whether or not the two nodes are currently diverged',
                },
                referenceHeight: {
                    type: common_ts_1.Gauge,
                    desc: 'Block height of the reference client',
                },
                targetHeight: {
                    type: common_ts_1.Gauge,
                    desc: 'Block height of the target client',
                },
                heightDifference: {
                    type: common_ts_1.Gauge,
                    desc: 'Difference in block heights between the two clients',
                },
                targetConnectionFailures: {
                    type: common_ts_1.Counter,
                    desc: 'Number of connection failures to the target client',
                },
                referenceConnectionFailures: {
                    type: common_ts_1.Counter,
                    desc: 'Number of connection failures to the reference client',
                },
            },
        });
    }
    async main() {
        let targetLatest;
        try {
            targetLatest = await this.options.targetRpcProvider.getBlock('latest');
        }
        catch (err) {
            if (err.message.includes('could not detect network')) {
                this.logger.error('target client not connected');
                this.metrics.targetConnectionFailures.inc();
                return;
            }
            else {
                throw err;
            }
        }
        let referenceLatest;
        try {
            referenceLatest = await this.options.referenceRpcProvider.getBlock('latest');
        }
        catch (err) {
            if (err.message.includes('could not detect network')) {
                this.logger.error('reference client not connected');
                this.metrics.referenceConnectionFailures.inc();
                return;
            }
            else {
                throw err;
            }
        }
        const heightDiff = Math.abs(referenceLatest.number - targetLatest.number);
        const minBlock = Math.min(targetLatest.number, referenceLatest.number);
        this.metrics.targetHeight.set(targetLatest.number);
        this.metrics.referenceHeight.set(referenceLatest.number);
        this.metrics.heightDifference.set(heightDiff);
        this.logger.info(`latest block heights`, {
            targetHeight: targetLatest.number,
            referenceHeight: referenceLatest.number,
            heightDifference: heightDiff,
            minBlockNumber: minBlock,
        });
        const reference = await this.options.referenceRpcProvider.getBlock(minBlock);
        if (!reference) {
            this.logger.info(`reference block was not found`, {
                blockNumber: reference.number,
            });
            return;
        }
        const target = await this.options.targetRpcProvider.getBlock(minBlock);
        if (!target) {
            this.logger.info(`target block was not found`, {
                blockNumber: target.number,
            });
            return;
        }
        if (reference.hash !== target.hash) {
            this.logger.error(`reference client has different hash for block`, {
                blockNumber: target.number,
                referenceHash: reference.hash,
                targetHash: target.hash,
            });
            this.logger.info(`beginning binary search to find first mismatched block`);
            let start = 0;
            let end = target.number;
            while (start !== end) {
                const mid = Math.floor((start + end) / 2);
                this.logger.info(`checking block`, { blockNumber: mid });
                const blockA = await this.options.referenceRpcProvider.getBlock(mid);
                const blockB = await this.options.targetRpcProvider.getBlock(mid);
                if (blockA.hash === blockB.hash) {
                    start = mid + 1;
                }
                else {
                    end = mid;
                }
            }
            this.logger.info(`found first mismatched block`, { blockNumber: end });
            this.metrics.lastMatchingStateRootHeight.set(end);
            this.metrics.isCurrentlyDiverged.set(1);
            await (0, core_utils_1.sleep)(this.options.onDivergenceWaitMs);
            return;
        }
        this.logger.info(`blocks are matching`, {
            blockNumber: target.number,
        });
        this.metrics.lastMatchingStateRootHeight.set(target.number);
        this.metrics.isCurrentlyDiverged.set(0);
    }
}
exports.HealthcheckService = HealthcheckService;
if (require.main === module) {
    const service = new HealthcheckService();
    service.run();
}
//# sourceMappingURL=service.js.map