// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


import "@openzeppelin/contracts/token/ERC20/IERC20.sol";


import "./interfaces/IMarginAccount.sol";
import "./interfaces/IHypERC20.sol";

contract MarginAccount is IMarginAccount {
    // Mapping to store user balances for different positions
    mapping(bytes32 => uint256) public balances;
    
    // Mapping to store locked funds for each user
    mapping(address => uint256) public lockedFunds;

    // ERC20 token used for transactions
    IERC20 public ropUSDC;

    // HypERC20 collateral address
    IHypERC20 public hypContract;
    address public hypContractAddress;

    // Address of the governing body
    address public gov;

    // Address of the position manager contract
    address public positionManager;

    // Mapping to store the order book address for each index token
    mapping(address => address) orderBookMap;

    // Penalty amount for insufficient funds
    uint256 insufficientFundsPenalty;

    // Event emitted when a user deposits funds
    event Deposit(address indexed user, uint256 amount);

    // Event emitted when a user withdraws funds
    event Withdraw(address indexed user, uint256 amount);

    /**
     * @dev Constructor.
     */
    constructor() {
        gov = msg.sender;
    }

    function addRop(address _ropUSDC) external {
        _onlyGov();
        ropUSDC = IERC20(_ropUSDC);
    }

    function addHyp(address _hypContractAddress) external {
        _onlyGov();
        hypContract = IHypERC20(_hypContractAddress);
        hypContractAddress = _hypContractAddress;
    }

    /**
     * @dev Adds an order book address for an index token.
     * Only the governing body can call this function.
     * @param _orderBook Address of the order book contract.
     * @param _indexToken Address of the index token.
     */
    function addOrderBook(address _orderBook, address _indexToken) external {
        _onlyGov();
        orderBookMap[_indexToken] = _orderBook;
    }

    /**
     * @dev Sets the position manager adderss.
     * Only the governing body can call this function.
     * @param _positionManager updated address.
     */
    function setPositionManager(address _positionManager) external {
        _onlyGov();
        positionManager = _positionManager;
    }

    /**
     * @dev Generates a unique key for a user's position.
     * @param _user Address of the user.
     * @param _token Address of the token.
     * @return Unique key for the position.
     */
    function getPositionKey(address _user, address _token) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(_user, _token));
    }

    /**
     * @dev Allows a user to deposit funds into their margin account.
     * The user must approve the contract to transfer the funds on their behalf.
     * @param _amount Amount to deposit.
     * @param _indexToken Address of the index token.
     * @param _user Address of the user depositing.
     */
    function deposit(uint256 _amount, address _indexToken, address _user) external {
        require(_amount > 0, "MarginAccount: Amount must be greater than zero");

        balances[getPositionKey(_user, _indexToken)] += _amount;
        
        emit Deposit(_user, _amount);
    }

    /**
     * @dev Allows a user to withdraw funds from their margin account.
     * @param _amount Amount to withdraw.
     * @param _indexToken Address of the index token.
     * TODO: update this with validation with position PnL
     */
    function withdraw(uint256 _amount, address _indexToken, uint32 _destination) external {
        require(_amount > 0, "MarginAccount: Amount must be greater than zero");
        require(_amount <= balances[getPositionKey(msg.sender, _indexToken)], "MarginAccount: Insufficient balance");

        bytes32 _recipient = addressToBytes32(msg.sender);

        ropUSDC.approve(hypContractAddress, _amount);

        hypContract.transferRemote(_destination, _recipient , _amount);

        balances[getPositionKey(msg.sender, _indexToken)] -= _amount;


        emit Withdraw(msg.sender, _amount);
    }


    /**
     * @dev Locks funds in a user's margin account.
     * Only the associated order book contract can call this function.
     * @param _user Address of the user.
     * @param _indexToken Address of the index token.
     */
    function lockFunds(address _user, address _indexToken) external override {
        require(msg.sender == orderBookMap[_indexToken], "MarginAccount: only order book can make this call");
        require(balances[getPositionKey(_user, _indexToken)] > insufficientFundsPenalty, "MarginAccount: insufficient funds");

        lockedFunds[_user] += insufficientFundsPenalty;
        balances[getPositionKey(_user, _indexToken)] -= insufficientFundsPenalty;
    }

    /**
     * @dev Unlocks previously locked funds in a user's margin account.
     * Only the associated order book contract can call this function.
     * @param _user Address of the user.
     * @param _indexToken Address of the index token.
     */
    function unlockFunds(address _user, address _indexToken) external override {
        require(msg.sender == orderBookMap[_indexToken], "MarginAccount: only order book can make this call");
        require(lockedFunds[_user] >= insufficientFundsPenalty, "MarginAccount: insufficient funds");

        lockedFunds[_user] -= insufficientFundsPenalty;
        balances[getPositionKey(_user, _indexToken)] += insufficientFundsPenalty;
    }

    /**
     * @dev Updates the funding fee for a user's position.
     * Only the position manager contract can call this function.
     * @param _user Address of the user.
     * @param _indexToken Address of the index token.
     * @param _fundingFee Funding fee to be applied (positive or negative).
     * @return Updated balance after applying the funding fee.
     */
    function updateFundingFee(address _user, address _indexToken, int256 _fundingFee) external override returns (uint256) {
        require(msg.sender == positionManager, "MarginAccount: only position manager can make this call");

        bytes32 _balanceKey = getPositionKey(_user, _indexToken);
        if (_fundingFee <= 0) {
            balances[_balanceKey] += uint256(-_fundingFee);
        } else {
            // Instead of this, you need to liquidate user positions
            require(balances[_balanceKey] > uint256(_fundingFee), "MarginAccount: insufficient funds");

            balances[_balanceKey] -= uint256(_fundingFee);
        }

        return balances[_balanceKey];
    }

    /**
     * @dev Changes the governing body address.
     * Only the current governing body can call this function.
     * @param newGov Address of the new governing body.
     */
    function changeGov(address newGov) external {
        _onlyGov();
        require(newGov != address(0), "MarginAccount: Gov cannot be zero address");
        gov = newGov;
    }

    /**
     * @dev Restricts function access to the governing body.
     */
    function _onlyGov() private view {
        require(msg.sender == gov, "MarginAccount: Only gov can call this function");
    }

    /**
    * @notice Converts an Ethereum address to a bytes32 representation.
    * @param _addr The address to be converted.
    * @return The bytes32 representation of the given address.
    */
    function addressToBytes32(address _addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(_addr)));
    }

}
