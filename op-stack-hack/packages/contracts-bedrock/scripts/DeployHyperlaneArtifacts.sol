// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Script } from "forge-std/Script.sol";
import { console2 as console } from "forge-std/console2.sol";

import { Mailbox } from "@hyperlane-xyz/core/contracts/Mailbox.sol";
import {
LegacyMultisigIsm
} from "@hyperlane-xyz/core/contracts/isms/multisig/LegacyMultisigIsm.sol";
import {InterchainGasPaymaster} from "../contracts/hyperlane/InterchainGasPaymaster.sol";
import {StorageGasOracle} from "../contracts/hyperlane/StorageGasOracle.sol";

contract DeployHyperlaneArtifacts is Script {
    function run() public {
        vm.startBroadcast();

        uint32 _localDomain = 84531;
        address _owner = 0x8a37895cd38557050309DE5CFD37CC875E8FDc45;

        // Deploy mailbox
        Mailbox mailbox = new Mailbox(84531);
        address _mailboxAddress = address(mailbox);
        console.log(_mailboxAddress);

        // Deploy ISM
        LegacyMultisigIsm ism = new LegacyMultisigIsm();
        address _ismAddress = address(ism);
        console.log(_ismAddress);

        // init mailbox
        mailbox.initialize(_owner, _ismAddress);

        // validator registor
        ism.enrollValidator(_localDomain, _owner);

        // Deploy InterchainGasPaymaster
        InterchainGasPaymaster igp = new InterchainGasPaymaster();
        address _igpAddress = address(igp);
        console.log(_igpAddress);

        igp.initialize(_owner, _owner);

        // Deploy Gas Oracle
        StorageGasOracle gasOracle = new StorageGasOracle();
        address _gasOracleAddress = address(gasOracle);
        console.log(_gasOracleAddress);

        InterchainGasPaymaster.GasOracleConfig[] memory config = new InterchainGasPaymaster.GasOracleConfig[](3);
        config[0] = InterchainGasPaymaster.GasOracleConfig(5, _gasOracleAddress);
        config[1] = InterchainGasPaymaster.GasOracleConfig(84531, _gasOracleAddress);
        config[2] = InterchainGasPaymaster.GasOracleConfig(42069, _gasOracleAddress);

        igp.setGasOracles(config);

        vm.stopBroadcast();
    }
}
