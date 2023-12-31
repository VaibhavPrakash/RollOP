import { AddressLike } from '../interfaces';
import { StandardBridgeAdapter } from './standard-bridge';
export declare class ECOBridgeAdapter extends StandardBridgeAdapter {
    supportsTokenPair(l1Token: AddressLike, l2Token: AddressLike): Promise<boolean>;
}
