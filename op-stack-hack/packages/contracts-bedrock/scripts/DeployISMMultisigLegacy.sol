// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Script } from "forge-std/Script.sol";
import { console2 as console } from "forge-std/console2.sol";
import {
    LegacyMultisigIsm
} from "@hyperlane-xyz/core/contracts/isms/multisig/LegacyMultisigIsm.sol";

contract DeployLegacyISM is Script {
    function run() public {
        vm.startBroadcast();
        LegacyMultisigIsm multisigISM = new LegacyMultisigIsm();
        console.log("multisigISM deployed at %s", address(multisigISM));

        vm.stopBroadcast();
    }
}
