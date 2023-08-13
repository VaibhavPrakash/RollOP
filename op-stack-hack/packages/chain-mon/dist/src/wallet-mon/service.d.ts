import { BaseServiceV2, StandardOptions, Counter } from '@eth-optimism/common-ts';
import { Provider } from '@ethersproject/abstract-provider';
type WalletMonOptions = {
    rpc: Provider;
    startBlockNumber: number;
};
type WalletMonMetrics = {
    validatedCalls: Counter;
    unexpectedCalls: Counter;
    unexpectedRpcErrors: Counter;
};
type WalletMonState = {
    chainId: number;
    highestUncheckedBlockNumber: number;
};
export declare class WalletMonService extends BaseServiceV2<WalletMonOptions, WalletMonMetrics, WalletMonState> {
    constructor(options?: Partial<WalletMonOptions & StandardOptions>);
    protected init(): Promise<void>;
    protected main(): Promise<void>;
}
export {};
