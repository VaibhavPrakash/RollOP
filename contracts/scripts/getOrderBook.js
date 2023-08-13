const fs = require("fs");
const hre = require("hardhat");

const orderBook = require("../artifacts/contracts/OrderBook.sol/OrderBook.json");
const config = require("../config.json");

async function main() {
    // Contract ABI
    const abi = orderBook.abi;

    // Contract Address
    const orderBookAddress = config.orderBookAddress;

    // Getting the contract instance
    const contract = await hre.ethers.getContractAt(abi, orderBookAddress);

    // Declare arrays for storing buy and sell price points
    const bids = [];
    const asks = [];

    const pricePrecision = ethers.BigNumber.from(10 ** 10);

    // Assuming that the keys for the price points are sequential and start from 0
    // You may need to modify this part depending on how the price points are stored
    for (let i = 0; i < 10; i++) {
        let price = hre.ethers.BigNumber.from(180000 + i * 1000);
        const pricePoints = await contract.s_buyPricePoints(price);
        const totalSize = pricePoints.totalOrdersAtPrice;
        const size =
            totalSize.div(pricePrecision).toNumber() +
            totalSize.mod(pricePrecision).toNumber() / 10 ** 10;

        if (size != 0) {
            asks.push({
                pricePoint: 180000 + i * 1000,
                size: size,
            });
        }
    }

    for (let i = 0; i < 10; i++) {
        let price = hre.ethers.BigNumber.from(180000 - i * 1000);
        const pricePoints = await contract.s_sellPricePoints(price);
        const totalSize = pricePoints.totalOrdersAtPrice;
        const size =
            totalSize.div(pricePrecision).toNumber() +
            totalSize.mod(pricePrecision).toNumber() / 10 ** 10;

        if (size != 0) {
            bids.push({
                pricePoint: 180000 - i * 1000,
                size: size,
            });
        }
    }

    asks.sort((a, b) => b.pricePoint - a.pricePoint);

    // Construct the JSON object
    const jsonObject = {
        asks,
        bids,  
    };

    // Write to a JSON file
    fs.writeFileSync("pricePoints.json", JSON.stringify(jsonObject, null, 2));
    console.log("Data successfully written to file");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
