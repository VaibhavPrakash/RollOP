const hre = require("hardhat");
const fs = require('fs');

async function main() {

    // deploy vault contract on gorli
    const Rollop = await hre.ethers.getContractFactory("ropUSDC");
    const rollop = await Rollop.deploy();
    await rollop.deployed();

    console.log(rollop.address);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});