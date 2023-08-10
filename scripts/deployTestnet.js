const hre = require("hardhat");
const fs = require('fs');

async function main() {
    const addresses = require('../utils/addresses.json')

    const chain = addresses['mantle-testnet']
    // contract addresses
    const usdcAddress = chain.usdcAddress;
    const wethAddress = chain.wethAddress;
    const oracleAddress = chain.oracleAddress;

    const pythId = chain.pythId;

    const [deployer] = await hre.ethers.getSigners();
    const userAddress = deployer.address;

    // deploy MarginAccount contract
    const MarginAccount = await hre.ethers.getContractFactory("MarginAccount");
    const marginAccount = await MarginAccount.deploy(usdcAddress);
    await marginAccount.deployed();

    const marginAccountAddress = marginAccount.address;
    console.log("MarginAccount deployed to:", marginAccountAddress);

    // deploy PositionManager contract
    const PositionManager = await hre.ethers.getContractFactory("PositionManager");
    const positionManager = await PositionManager.deploy(marginAccountAddress, oracleAddress);
    await positionManager.deployed();

    const positionManagerAddress = positionManager.address;
    console.log("PositionManager deployed to:", positionManagerAddress);

    // deploy OrderBook contract
    const OrderBook = await hre.ethers.getContractFactory("OrderBook");
    const orderBook = await OrderBook.deploy(positionManagerAddress, marginAccountAddress, wethAddress);
    await orderBook.deployed();

    const orderBookAddress = orderBook.address;
    console.log("OrderBook deployed to:", orderBookAddress);

    // add orderBookAddress to MarginAccount
    var receipt = await marginAccount.addOrderBook(orderBookAddress, wethAddress);
    await receipt.wait();

    // set positionManagerAddress in MarginAccount
    var receipt = await marginAccount.setPositionManager(positionManagerAddress);
    await receipt.wait();

    // add orderBookAddress to PositionManager
    receipt = await positionManager.addOrderBook(orderBookAddress, wethAddress);
    await receipt.wait();

    // add oracleAddress to PositionManager
    receipt = await positionManager.addPythOracle(pythId, wethAddress);
    await receipt.wait();

    // write contract addresses to a config.json file
    const config = {
        "usdcAddress": usdcAddress,
        "wethAddress": wethAddress,
        "oracleAddress": oracleAddress,
        "marginAccountAddress": marginAccountAddress,
        "positionManagerAddress": positionManagerAddress,
        "orderBookAddress": orderBookAddress
    }
    fs.writeFileSync("configMantleTestnet.json", JSON.stringify(config));
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});