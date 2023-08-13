import { BaseServiceV2, StandardOptions, ExpressRouter, Gauge } from '@eth-optimism/common-ts';
import { CrossChainMessenger } from '@eth-optimism/sdk';
import { Provider } from '@ethersproject/abstract-provider';
type Options = {
    l1RpcProvider: Provider;
    l2RpcProvider: Provider;
    startBlockNumber: number;
    sleepTimeMs: number;
};
type Metrics = {
    withdrawalsValidated: Gauge;
    isDetectingForgeries: Gauge;
    nodeConnectionFailures: Gauge;
};
type State = {
    messenger: CrossChainMessenger;
    highestUncheckedBlockNumber: number;
    finalizationWindow: number;
    forgeryDetected: boolean;
};
export declare class WithdrawalMonitor extends BaseServiceV2<Options, Metrics, State> {
    constructor(options?: Partial<Options & StandardOptions>);
    init(): Promise<void>;
    routes(router: ExpressRouter): Promise<void>;
    main(): Promise<void>;
}
export {};
