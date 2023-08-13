import { Provider } from '@ethersproject/abstract-provider';
import { BaseServiceV2, StandardOptions, Counter, Gauge } from '@eth-optimism/common-ts';
type HealthcheckOptions = {
    referenceRpcProvider: Provider;
    targetRpcProvider: Provider;
    onDivergenceWaitMs?: number;
};
type HealthcheckMetrics = {
    lastMatchingStateRootHeight: Gauge;
    isCurrentlyDiverged: Gauge;
    referenceHeight: Gauge;
    targetHeight: Gauge;
    heightDifference: Gauge;
    targetConnectionFailures: Counter;
    referenceConnectionFailures: Counter;
};
type HealthcheckState = {};
export declare class HealthcheckService extends BaseServiceV2<HealthcheckOptions, HealthcheckMetrics, HealthcheckState> {
    constructor(options?: Partial<HealthcheckOptions & StandardOptions>);
    main(): Promise<void>;
}
export {};
