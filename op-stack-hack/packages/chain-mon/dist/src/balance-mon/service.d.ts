import { BaseServiceV2, StandardOptions, Gauge, Counter } from '@eth-optimism/common-ts';
import { Provider } from '@ethersproject/abstract-provider';
type BalanceMonOptions = {
    rpc: Provider;
    accounts: string;
};
type BalanceMonMetrics = {
    balances: Gauge;
    unexpectedRpcErrors: Counter;
};
type BalanceMonState = {
    accounts: Array<{
        address: string;
        nickname: string;
    }>;
};
export declare class BalanceMonService extends BaseServiceV2<BalanceMonOptions, BalanceMonMetrics, BalanceMonState> {
    constructor(options?: Partial<BalanceMonOptions & StandardOptions>);
    protected init(): Promise<void>;
    protected main(): Promise<void>;
}
export {};
