# RollOP
When L2’s were created, another problem came with it. _Fragmented liquidity_ . So we created RollOP.


## What is RollOP
RollOP is an Orderbook based Perpetuals DEX on a custom rollup based on the [OP-Stack](https://stack.optimism.io/). You can start trading from Optimism, Base and Zora. Using Hyperlane’s Warp Routes and Messaging, we are able to create a unified pool. Users can experience faster trading speeds for their perps.

- Step 1 : User deposits USDC on our Contracts (on Optimism and Base and Zora)
- Step 2 : Start trading perps on RollOP
- Step 3 : Withdraw their funds (to Optimism/Base/Zora)
- Step4 : Receive ropUSDC (more on ropUSDC below)
- Step 5 : Either swap it to USDC on our Contracts or transfer and use it across Optimism/Base/Zora

## How it's Made
RollOP is based on the OP-Stack with important modifications with the help of [Hyperlane](https://www.hyperlane.xyz/).

We have deployed Hyperlane to give cross chain functionality. This includes messaging and warp routes. We deployed our own Interchain Security Modules and Interchain Gas Paymasters. We have also written scripts to set up your own OP-Stack chain with Hyperlane!

The rollup doesn't have access to oracles, so we built our own pricefeeds for the Orderbook. These are used to update the funding rates for our DEX.

ropUSDC is minted 1:1 when a user deposits USDC on our Contracts. It can be used across all L2’s that we support (currently Optimism, Base and Zora). Users are free to swap it with USDC on any of the supported L2’s. Other DeFi products can also be built around ropUSDC as a store of value.

## About RollOP
RollOP was created during the Superhack Hackathon organised by ETHGlobal.

Here is a link to the project page : [RollOP](https://ethglobal.com/showcase/rollop-ryx8t)
