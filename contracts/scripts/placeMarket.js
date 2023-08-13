const hre = require("hardhat");
const config = require("../config.json");

async function main() {
    const [deployer] = await hre.ethers.getSigners();
    const userAddress = deployer.address;

    // create contract instance of OrderBook from config address
    const orderBook = await hre.ethers.getContractAt("OrderBook", config.orderBookAddress);

    var _maxBid = await orderBook.s_buyOrdersHeap(0)
    _maxBid = _maxBid.toNumber()

    var m_pricePoint = await orderBook.s_buyPricePoints(_maxBid);

    console.log(_maxBid)
    console.log(m_pricePoint)

    var receipt = await orderBook.placeAndExecuteMarketSell(
        3*10**9,
        userAddress,
    );
    const tx = await receipt.wait();
    console.log(tx)
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
