// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./libraries/IERC20.sol";
import "./interfaces/IMessageRecipient.sol";
import "./interfaces/IropUSDC.sol";


contract ropUSDC is IERC20, IMessageRecipient {
    string public constant NAME = "ropUSDC";
    string public constant SYMBOL = "ropUSDC";
    uint8 public constant DECIMALS = 6;

    uint256 public totalSupply;
    mapping(address => uint256) public balances;
    mapping(address => mapping(address => uint256)) public allowances;

    address public mailbox = 0x4200000000000000000000000000000000000068;

    event Received(uint32 origin, address sender, bytes body);

    modifier onlyMailbox() {
        require(msg.sender == mailbox, "Only mailbox allowed");
        _;
    }


    function handle(uint32 _origin, bytes32 _sender, bytes calldata _body) external override onlyMailbox {
        address recipient;
        uint256 amount;

        (recipient, amount) = abi.decode(_body, (address, uint256));

        _mint(recipient, amount);

        emit Received(_origin, bytes32ToAddress(_sender), _body);
    }

    function _mint(address to, uint256 amount) internal {
        totalSupply += amount;
        balances[to] += amount;
        emit Transfer(address(0), to, amount);
    }

    function bytes32ToAddress(bytes32 _addr) internal pure returns (address) {
        return address(uint160(uint256(_addr)));
    }

    //------------------------------ERC20 functions----------------------------------
    // Remaining ERC20 functions

    function balanceOf(address account) public view override returns (uint256) {
        return balances[account];
    }

    function transfer(address recipient, uint256 amount) public override returns (bool) {
        _transfer(msg.sender, recipient, amount);
        return true;
    }

    function allowance(address owner, address spender) public view override returns (uint256) {
        return allowances[owner][spender];
    }

    function approve(address spender, uint256 amount) public override returns (bool) {
        allowances[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    function transferFrom(address sender, address recipient, uint256 amount) public override returns (bool) {
        require(allowances[sender][msg.sender] >= amount, "ERC20: Transfer amount exceeds allowance");
        _transfer(sender, recipient, amount);
        
        // Decrement the spender's allowance
        allowances[sender][msg.sender] -= amount;
        emit Approval(sender, msg.sender, allowances[sender][msg.sender]);
        return true;
    }

    // Internal transfer function that's reused for both transfer and transferFrom
    function _transfer(address sender, address recipient, uint256 amount) internal {
        require(sender != address(0), "ERC20: transfer from the zero address");
        require(recipient != address(0), "ERC20: transfer to the zero address");
        require(balances[sender] >= amount, "ERC20: Transfer amount exceeds balance");

        balances[sender] -= amount;
        balances[recipient] += amount;

        emit Transfer(sender, recipient, amount);
    }
    
}
