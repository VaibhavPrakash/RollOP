const hre = require("hardhat");
const config = require("../config.json");

async function placeLimit() {
    // create contract instance of OrderBook from config address
    const orderBook = await hre.ethers.getContractAt(
        "OrderBook",
        config.orderBookAddress
    );

    let totalBuys = 0;
    let totalSells = 0;

    // loop over 50 and place buy and sell limit orders
    for (let i = 1; i < 11; i++) {
        var tx = await orderBook.addBuyOrder(180000 + i * 1000, i * 1e9);
        const receipt = await tx.wait();
        receipt ? (totalBuys += 1) : totalBuys;

        var tx2 = await orderBook.addSellOrder(180000 - i * 1000, i * 1e9);
        const receipt2 = await tx2.wait();
        receipt2 ? (totalSells += 1) : totalSells;

        // console.log("Gas used: " + receipt.cumulativeGasUsed);
    }

    console.log(`Placed Limit Orders: ${totalBuys} buys and ${totalSells} sells`);
}

module.exports = { placeLimit };

// main().catch((error) => {
//     console.error(error);
//     process.exitCode = 1;
// });
