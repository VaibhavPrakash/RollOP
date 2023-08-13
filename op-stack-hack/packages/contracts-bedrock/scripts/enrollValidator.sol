// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Script } from "forge-std/Script.sol";
import {
LegacyMultisigIsm
} from "@hyperlane-xyz/core/contracts/isms/multisig/LegacyMultisigIsm.sol";

contract EnrollValidator is Script {
    function run() public {
        vm.startBroadcast();
        LegacyMultisigIsm ism = LegacyMultisigIsm(0xcFaE7bFD701Be2156B96d0c452F886d226FB6422);
        ism.enrollValidator(5, 0x8a37895cd38557050309DE5CFD37CC875E8FDc45);
        ism.setThreshold(5, 1);

        vm.stopBroadcast();
    }
}
