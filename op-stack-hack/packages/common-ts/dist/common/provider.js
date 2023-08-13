"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.waitForProvider = void 0;
const core_utils_1 = require("@eth-optimism/core-utils");
const waitForProvider = async (provider, opts) => {
    var _a, _b, _c;
    const name = (opts === null || opts === void 0 ? void 0 : opts.name) || 'target';
    (_a = opts === null || opts === void 0 ? void 0 : opts.logger) === null || _a === void 0 ? void 0 : _a.info(`waiting for ${name} provider...`);
    let connected = false;
    while (!connected) {
        try {
            await provider.getBlockNumber();
            connected = true;
        }
        catch (e) {
            (_b = opts === null || opts === void 0 ? void 0 : opts.logger) === null || _b === void 0 ? void 0 : _b.info(`${name} provider not connected, retrying...`);
            await (0, core_utils_1.sleep)((opts === null || opts === void 0 ? void 0 : opts.intervalMs) || 15000);
        }
    }
    (_c = opts === null || opts === void 0 ? void 0 : opts.logger) === null || _c === void 0 ? void 0 : _c.info(`${name} provider connected`);
};
exports.waitForProvider = waitForProvider;
//# sourceMappingURL=provider.js.map