"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.findFirstUnfinalizedStateBatchIndex = exports.findOutputForIndex = void 0;
const findOutputForIndex = async (oracle, index, logger) => {
    try {
        const proposal = await oracle.getL2Output(index);
        return {
            outputRoot: proposal.outputRoot,
            l1Timestamp: proposal.timestamp.toNumber(),
            l2BlockNumber: proposal.l2BlockNumber.toNumber(),
            l2OutputIndex: index,
        };
    }
    catch (err) {
        logger === null || logger === void 0 ? void 0 : logger.fatal('error when calling L2OuputOracle.getL2Output', {
            errors: err,
        });
        throw new Error(`unable to find output for index ${index}`);
    }
};
exports.findOutputForIndex = findOutputForIndex;
const findFirstUnfinalizedStateBatchIndex = async (oracle, fpw, logger) => {
    const latestBlock = await oracle.provider.getBlock('latest');
    const totalBatches = (await oracle.nextOutputIndex()).toNumber();
    let lo = 0;
    let hi = totalBatches;
    while (lo !== hi) {
        const mid = Math.floor((lo + hi) / 2);
        const outputData = await (0, exports.findOutputForIndex)(oracle, mid, logger);
        if (outputData.l1Timestamp + fpw < latestBlock.timestamp) {
            lo = mid + 1;
        }
        else {
            hi = mid;
        }
    }
    if (lo === totalBatches) {
        return undefined;
    }
    else {
        return lo;
    }
};
exports.findFirstUnfinalizedStateBatchIndex = findFirstUnfinalizedStateBatchIndex;
//# sourceMappingURL=helpers.js.map