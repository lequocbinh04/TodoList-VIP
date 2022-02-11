require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-ethers");
require("@nomiclabs/hardhat-etherscan");
// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async () => {
    const accounts = await ethers.getSigners();

    for (const account of accounts) {
        console.log(account.address);
    }
});

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
    defaultNetwork: "testnet",
    networks: {
        localhost: {
            url: "http://127.0.0.1:8545",
        },
        hardhat: {},
        testnet: {
            url: "https://data-seed-prebsc-1-s1.binance.org:8545",
            accounts: ["your-private-key"],
        },
    },
    solidity: {
        version: "0.8.0",
        settings: {
            optimizer: {
                enabled: true,
            },
        },
    },
    paths: {
        sources: "./contracts",
        tests: "./test",
        cache: "./cache",
        artifacts: "./artifacts",
    },
    etherscan: {
        apiKey: {
            bscTestnet: "your-api-key",
        },
    },

    mocha: {
        timeout: 20000,
    },
};
