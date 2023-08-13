require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    solidity: "0.8.17",
    defaultNetwork: "localhost",
    networks: {
        localhost: {
            url: "http://127.0.0.1:8545/",
        },
        hardhat: {
          // url:"http://127.0.0.1:8545/"
      // forking: {
      //     url: "https://arb1.arbitrum.io/rpc",
      // },
    },
    goerli: {
      url: "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
            accounts: [
                "f94f2358316ae919ed24243bcf55fd3638539676d30acab3d1b25b7717b5ae38",
            ],
        },
        rollop: {
            url: "http://127.0.0.1:8545/",
            accounts: [
                "f94f2358316ae919ed24243bcf55fd3638539676d30acab3d1b25b7717b5ae38",
            ],
        },
    },
};
