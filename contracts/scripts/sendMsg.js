const hre = require("hardhat");
const fs = require('fs');

async function main() {
    const dispatcherAddress = "0x1aF6fdD1614549A2DCDfcC14DC7443F6Cf84A5F7";
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