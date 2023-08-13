const hre = require("hardhat");
const fs = require('fs');

async function main() {
    const addresses = require('../utils/addresses.json')

    const chain = addresses['arbitrum']

    const usdcAddress = chain.usdcAddress;
    const goerliMailbox = "0xb1c047829d7aFe3BFbE4090C8559016970940478";

    // deploy vault contract on gorli
    const VaultDispatcher = await hre.ethers.getContractFactory("VaultDispatcher");
    const vaultDispatcher = await VaultDispatcher.deploy(usdcAddress, goerliMailbox);
    await vaultDispatcher.deployed();

    console.log("VaultDispatcher deployed at : ",vaultDispatcher.address)
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});