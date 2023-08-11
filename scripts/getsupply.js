const hre = require("hardhat");
const fs = require('fs');

async function main() {
    const usdcAddress = "0xcFaE7bFD701Be2156B96d0c452F886d226FB6422";
    // deploy vault contract on gorli
    const VaultDispatcher = await hre.ethers.getContractAt("ropUSDC", usdcAddress);
    
    var supp = await VaultDispatcher.totalSupply();

    console.log(supp)
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});