const hre = require("hardhat");
const fs = require('fs');

async function main() {
    const rolUSDC = "0x373A75ddAc70C25F1fBA91985dc0Cc1A6C12C9F3";
    const hypAd = "0xBe03d997E35cCaFD46Ce5577B56F0a45C59f391f";
    // deploy vault contract on gorli
    const RolUSDC = await hre.ethers.getContractAt("IERC20", rolUSDC);

    var tx = await RolUSDC.approve(hypAd, 100000);
    var receipt = await tx.wait();

    console.log(receipt)

    const HypCol = await hre.ethers.getContractAt("IHypERC20", hypAd);
    var tx = await RolUSDC.approve(hypAd, 100000);
    var receipt = await tx.wait();


    console.log(receipt)
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});