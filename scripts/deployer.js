const hre = require("hardhat");
const fs = require('fs');

async function main() {
    const usdcAddress = "0x373A75ddAc70C25F1fBA91985dc0Cc1A6C12C9F3";
    const goerliMailbox = "0xb1c047829d7aFe3BFbE4090C8559016970940478";

    // deploy vault contract on gorli
    const VaultDispatcher = await hre.ethers.getContractFactory("VaultDispatcher");
    const marginAccount = await VaultDispatcher.deploy(usdcAddress, goerliMailbox);
    await marginAccount.deployed();

    console.log(marginAccount.address)
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});