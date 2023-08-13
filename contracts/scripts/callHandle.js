const hre = require("hardhat");
const fs = require('fs');

async function main() {
    const usdcAddress = "0xcFaE7bFD701Be2156B96d0c452F886d226FB6422";
    // deploy vault contract on gorli
    const VaultDispatcher = await hre.ethers.getContractAt("ropUSDC", usdcAddress);
    
    var tx = await VaultDispatcher.handle(5, "0x000000000000000000000000cfae7bfd701be2156b96d0c452f886d226fb6422", "0x0000000000000000000000008a37895cd38557050309de5cfd37cc875e8fdc4500000000000000000000000000000000000000000000000000000000000f4240");

    const receipt = await tx.wait();
    console.log(receipt);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});