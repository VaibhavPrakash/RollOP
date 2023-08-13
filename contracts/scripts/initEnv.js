const hre = require("hardhat");
const fs = require("fs");
const { placeLimit } = require("./placeLimit");


async function main() {
    const activeNetwork = await hre.ethers.provider.getNetwork();
    console.log("Network name: ", activeNetwork.name);
    console.log("Network chainId: ", activeNetwork.chainId);
    
    const addresses = require("../utils/addresses.json");

    const chain = addresses["arbitrum"];

    // contract addresses
    const usdcAddress = chain.usdcAddress;
    const hypAddress = usdcAddress
    const wethAddress = chain.wethAddress;
    // const oracleAddress = chain.oracleAddress;

    const priceFeedAddress = chain.priceFeedAddress;

    // shark addresses
    const sharkAddress = chain.sharkAddress;

    const [deployer] = await hre.ethers.getSigners();
    const userAddress = deployer.address;

    console.log("Sender address is : ", userAddress);

    // deploy MarginAccount contract
    const MarginAccount = await hre.ethers.getContractFactory("MarginAccount");
    const marginAccount = await MarginAccount.deploy(usdcAddress,hypAddress);
    await marginAccount.deployed();

    const marginAccountAddress = marginAccount.address;
    console.log("MarginAccount deployed to:", marginAccountAddress);

    // deploy PositionManager contract
    const PositionManager = await hre.ethers.getContractFactory(
        "PositionManager"
    );
    const positionManager = await PositionManager.deploy(
        marginAccountAddress,
        priceFeedAddress
    );
    await positionManager.deployed();

    const positionManagerAddress = positionManager.address;
    console.log("PositionManager deployed to:", positionManagerAddress);

    // deploy OrderBook contract
    const OrderBook = await hre.ethers.getContractFactory("OrderBook");
    const orderBook = await OrderBook.deploy(
        positionManagerAddress,
        marginAccountAddress,
        wethAddress
    );
    await orderBook.deployed();

    const orderBookAddress = orderBook.address;
    console.log("OrderBook deployed to:", orderBookAddress);

    // deploy ropUSDC contract
    const RopUSDC = await hre.ethers.getContractFactory("ropUSDC");
    const ropUSDC = await RopUSDC.deploy(marginAccountAddress);
    await ropUSDC.deployed();

    const ropUSDCAddress = ropUSDC.address;
    console.log("ropUSDC deployed to:", ropUSDCAddress);

    // add orderBookAddress to MarginAccount
    var receipt = await marginAccount.addOrderBook(
        orderBookAddress,
        wethAddress
    );
    await receipt.wait();

    // set positionManagerAddress in MarginAccount
    var receipt = await marginAccount.setPositionManager(
        positionManagerAddress
    );
    await receipt.wait();

    // add orderBookAddress to PositionManager
    receipt = await positionManager.addOrderBook(orderBookAddress, wethAddress);
    await receipt.wait();

    // add priceFeedAddress to PositionManager
    receipt = await positionManager.addPriceFeed(priceFeedAddress, wethAddress);
    await receipt.wait();

    // imporsonate sharkAddress
    await network.provider.request({
        method: "hardhat_impersonateAccount",
        params: [sharkAddress],
    });

    const signer = await hre.ethers.getSigner(sharkAddress);

    const usdcContractShark = await hre.ethers.getContractAt(
        "contracts/libraries/IERC20.sol:IERC20",
        usdcAddress,
        signer
    );

    var receipt = await usdcContractShark.transfer(userAddress, 1000000000);

    await receipt.wait();

    console.log(
        "usdc balance: ",
        (await usdcContractShark.balanceOf(userAddress)).toString()
    );

    // approve usdc to MarginAccount from deployer
    const usdcContract = await hre.ethers.getContractAt(
        "contracts/libraries/IERC20.sol:IERC20",
        usdcAddress
    );
    var receipt = await usdcContract.approve(marginAccountAddress, 1000000000);
    await receipt.wait();

    // deposit usdc to MarginAccount
    var receipt = await marginAccount.deposit(
        1000000000,
        wethAddress,
        "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
    );
    await receipt.wait();

    // write contract addresses to a config.json file
    const config = {
        usdcAddress: usdcAddress,
        wethAddress: wethAddress,
        priceFeedAddress: priceFeedAddress,
        sharkAddress: sharkAddress,
        marginAccountAddress: marginAccountAddress,
        positionManagerAddress: positionManagerAddress,
        orderBookAddress: orderBookAddress,
        ropUSDCAddress: ropUSDCAddress,
    };
    fs.writeFileSync("config.json", JSON.stringify(config));

    await placeLimit();
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
