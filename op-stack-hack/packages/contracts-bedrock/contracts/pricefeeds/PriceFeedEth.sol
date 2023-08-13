// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract PriceFeedEth{
    uint64 public currentPrice;
    mapping(uint64 => uint64) dataFeed;

    function setValue(uint64 _blockNo, uint64 _price) external{
        require(msg.sender == 0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001,"Only system address can call this function");
        currentPrice = _price;
        dataFeed[_blockNo] = _price;
    }

    function getPrice() view external returns(uint64){
        return currentPrice;
    }

    function getPriceBlock(uint64 _blockNo) view external returns(uint64){
        return dataFeed[_blockNo];
    }

}
