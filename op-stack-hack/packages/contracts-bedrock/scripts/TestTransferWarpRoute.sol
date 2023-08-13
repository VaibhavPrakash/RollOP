// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Script } from "forge-std/Script.sol";

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {IHypERC20} from "./IHypeERC20.sol";

contract WarpTransfer is Script {
    function run() public {
        vm.startBroadcast();

        ERC20 ropUsdc = ERC20(0x06Ee73Ab8994bBe9ADDcA6e3CFbDe523CcAd2e6F);
        ropUsdc.approve(0xA2521B9bc2bCc63d25De52c2d56B245Ea7cC9515, 1000000);

        IHypERC20 hypContract = IHypERC20(0xA2521B9bc2bCc63d25De52c2d56B245Ea7cC9515);

        hypContract.transferRemote(5, bytes32(uint256(uint160(0x8a37895cd38557050309DE5CFD37CC875E8FDc45))), 1000000);
        vm.stopBroadcast();
    }
}
