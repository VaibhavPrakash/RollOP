import { BaseServiceV2, StandardOptions, Gauge, Counter } from '@eth-optimism/common-ts';
import { Provider } from '@ethersproject/abstract-provider';
import { ethers } from 'ethers';
type DrippieMonOptions = {
    rpc: Provider;
    drippieAddress: string;
};
type DrippieMonMetrics = {
    isExecutable: Gauge;
    executedDripCount: Gauge;
    unexpectedRpcErrors: Counter;
};
type DrippieMonState = {
    drippie: ethers.Contract;
};
export declare class DrippieMonService extends BaseServiceV2<DrippieMonOptions, DrippieMonMetrics, DrippieMonState> {
    constructor(options?: Partial<DrippieMonOptions & StandardOptions>);
    protected init(): Promise<void>;
    protected main(): Promise<void>;
}
export {};
