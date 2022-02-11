<p align="center">
  <a href="https://www.facebook.com/lequocbinh.04">
    <img src="https://icons.iconarchive.com/icons/cjdowner/cryptocurrency-flat/1024/Ethereum-ETH-icon.png" alt="Logo" width=72 height=72>
  </a>

  <h3 align="center">Smart Contract for TodoList VIP</h3>

  <p align="center">
    This branch is the Smart Contract code. If you have any questions, please contact me 😁
    <br>
    <a href="https://www.facebook.com/lequocbinh.04">Report bug</a>
    ·
    <a href="https://www.facebook.com/lequocbinh.04">Request feature</a>
  </p>
</p>

# Infomation

This branch is the Smart Contract code, including: 1 token (Todo Token), and 1 contract. This code uses [Hardhat](https://hardhat.org/getting-started/) with solidity 0.8.4 compiler in addition to programming tokens or contracts quickly I also use [OpenZeppelin](https://openzeppelin.com/).
&nbsp;

&nbsp;

Some information about the contract I deployed:

```shell
Todo token (TDT) address: 0x708C0BC45208776B32382983561ce9d38bD5F778
```

```shell
Contract to claim token address: 0xfF2df1968724477B99C58f8952CD130D8c993AbA
```

&nbsp;

# Usage

1. Config file `hardhat.config.js` (add mainnet if you want, add BSC API_KEY, change private key to deploy contract)
2. Run the command below to start deploying contract. Change file to run & change network to deploy.

```shell
npx hardhat run scripts/deploy-claim-smart-contract.js --network testnet
```

3. Run the command below to start validating contract

```shell
 npx hardhat  verify --network testnet <token_address>
```
