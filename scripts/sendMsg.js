const hre = require("hardhat");
const fs = require('fs');

async function main() {
    const dispatcherAddress = "0xA44f9d5dA0127B406D691Be1a7150d7C36353E72";
    // deploy vault contract on gorli
    const VaultDispatcher = await hre.ethers.getContractAt("VaultDispatcher", dispatcherAddress);
    
    var tx = await VaultDispatcher.depositAndDispatch(1000000);
    const receipt = await tx.wait();

    console.log(receipt)
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});