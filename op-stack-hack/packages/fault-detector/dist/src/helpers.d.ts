import { Contract } from 'ethers';
import { Logger } from '@eth-optimism/common-ts';
import { BedrockOutputData } from '@eth-optimism/core-utils';
export declare const findOutputForIndex: (oracle: Contract, index: number, logger?: Logger) => Promise<BedrockOutputData>;
export declare const findFirstUnfinalizedStateBatchIndex: (oracle: Contract, fpw: number, logger?: Logger) => Promise<number>;
