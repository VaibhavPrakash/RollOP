// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IropUSDC {
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    event Received(uint32 origin, address sender, bytes body);

    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function allowance(address owner, address spender) external view returns (uint256);

    function approve(address spender, uint256 value) external returns (bool);
    function transfer(address to, uint256 value) external returns (bool);
    function transferFrom(address from, address to, uint256 value) external returns (bool);

    function handle(
        uint32 _origin,
        bytes32 _sender,
        bytes calldata _body
    ) external;
}
