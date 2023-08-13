import { BaseServiceV2, StandardOptions, ExpressRouter, Gauge } from '@eth-optimism/common-ts';
import { CrossChainMessenger, OEL1ContractsLike } from '@eth-optimism/sdk';
import { Provider } from '@ethersproject/abstract-provider';
import { Contract } from 'ethers';
type Options = {
    l1RpcProvider: Provider;
    l2RpcProvider: Provider;
    startBatchIndex: number;
    optimismPortalAddress?: string;
};
type Metrics = {
    highestBatchIndex: Gauge;
    isCurrentlyMismatched: Gauge;
    nodeConnectionFailures: Gauge;
};
type State = {
    faultProofWindow: number;
    outputOracle: Contract;
    messenger: CrossChainMessenger;
    currentBatchIndex: number;
    diverged: boolean;
};
export declare class FaultDetector extends BaseServiceV2<Options, Metrics, State> {
    constructor(options?: Partial<Options & StandardOptions>);
    getOEL1Contracts(l2ChainId: number): Promise<OEL1ContractsLike>;
    init(): Promise<void>;
    routes(router: ExpressRouter): Promise<void>;
    main(): Promise<void>;
}
export {};
