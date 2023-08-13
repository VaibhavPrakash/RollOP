// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Script } from "forge-std/Script.sol";
import {InterchainGasPaymaster} from "../contracts/hyperlane/InterchainGasPaymaster.sol";
import "../contracts/L2/GasPriceOracle.sol";

contract EnrollValidator is Script {
    function run() public {
        vm.startBroadcast();
        InterchainGasPaymaster igp = InterchainGasPaymaster(0x420000000000000000000000000000000000006d);
        InterchainGasPaymaster.GasOracleConfig[] memory config = new InterchainGasPaymaster.GasOracleConfig[](1);
        config[0] = InterchainGasPaymaster.GasOracleConfig(84531, 0x420000000000000000000000000000000000006f);

        igp.setGasOracles(config);

        vm.stopBroadcast();
    }
}
