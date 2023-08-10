// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IVaultDispatcher {

    /**
     * @notice Allows a user to deposit a specific amount of USDC and dispatches it.
     * @dev Transfers USDC from user to this contract and then dispatches the specified amount.
     * @param amount The amount of USDC to deposit.
     */
    function depositAndDispatch(uint256 amount) external;

    /**
     * @notice Allows the contract owner to recover accidentally sent USDC.
     * @dev Transfers the balance of USDC in this contract to the caller.
     */
    function recoverUSDC() external;

}
