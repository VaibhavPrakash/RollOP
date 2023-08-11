require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.17",
  defaultNetwork: "goerli",
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545/",
    },
    mantleTestnet: {
      url: "https://rpc.testnet.mantle.xyz/",
      accounts: ["1de8f879576ea64d7cfa74879bfb7486dc5e1a7cfb19d1507fe87c782f9c54b3"],
    },
    mantleMainnet: {
      url: "https://rpc.mantle.xyz/",
      accounts: ["1de8f879576ea64d7cfa74879bfb7486dc5e1a7cfb19d1507fe87c782f9c54b3"],
    },
    hardhat: {
      forking: {
          url: "https://arb1.arbitrum.io/rpc",
      },
    },
    goerli: {
      url: "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
      accounts: ["f94f2358316ae919ed24243bcf55fd3638539676d30acab3d1b25b7717b5ae38"]
    },
    rollop: {
      url: "http://127.0.0.1:8545/",
      accounts: ["f94f2358316ae919ed24243bcf55fd3638539676d30acab3d1b25b7717b5ae38"]
    },
  }
};
