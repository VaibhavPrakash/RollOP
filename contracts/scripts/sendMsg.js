const hre = require("hardhat");
const fs = require('fs');

async function main() {
    const dispatcherAddress = "0x5d0B3883042872B0924C053fb48948b465071c65";
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