// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IPriceFeed {

    function getPrice(address id) external view returns (uint64);

}
