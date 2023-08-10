// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./libraries/IERC20.sol";
import "./interfaces/IMailbox.sol";
import "./interfaces/IVaultDispatcher.sol";


/**
 * @title VaultDispatcher
 * @dev This contract allows for deposits of USDC and dispatches them to another contract.
 */
contract VaultDispatcher is IVaultDispatcher {

    IERC20 public usdc; // USDC token contract address
    IMailbox public chainMailbox;

    uint32 constant ROLLOP_DOMAIN = 42069;


    /**
     * @dev Sets the addresses of USDC and Dispatcher upon deployment.
     * @param _usdc Address of the USDC token contract.
     * @param _chainMailbox Address of the chain mailbox contract.
     */
    constructor(address _usdc, address _chainMailbox) {
        usdc = IERC20(_usdc);
        chainMailbox = IMailbox(_chainMailbox);
    }

    /**
     * @notice Allows a user to deposit a specific amount of USDC and dispatches it.
     * @dev Transfers USDC from user to this contract and then dispatches the specified amount.
     * @param amount The amount of USDC to deposit.
     */
    function depositAndDispatch(uint256 amount) external {
        // Transfer USDC from user to this contract
        require(usdc.transferFrom(msg.sender, address(this), amount), "USDC transfer failed");

        // Construct the _body with the user's address and deposited amount
        bytes memory body = abi.encode(msg.sender, amount);

        // Call dispatch with the given mock values
        chainMailbox.dispatch(ROLLOP_DOMAIN, addressToBytes32(msg.sender), body);
    }

    /**
     * @notice Allows the contract owner to recover accidentally sent USDC.
     * @dev Transfers the balance of USDC in this contract to the caller.
     */
    function recoverUSDC() external {
        uint256 balance = usdc.balanceOf(address(this));
        require(usdc.transferFrom(address(this), msg.sender, balance), "Recovery failed");
    }

    /**
     * @dev Converts an Ethereum address to bytes32 format.
     * @param _addr The Ethereum address to convert.
     * @return _addr in bytes32 format.
     */
    function addressToBytes32(address _addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(_addr)));
    }
}
