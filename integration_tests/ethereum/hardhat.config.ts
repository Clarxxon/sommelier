import '@nomiclabs/hardhat-ethers';
import '@nomiclabs/hardhat-waffle';
import {task} from "hardhat/config";
import * as constants from "./addresses";

task(
    'integration_test_setup',
    'Sets up contracts for the integration test',
    async (args, hre) => {
        let powers: number[] = [1073741824,1073741824,1073741824,1073741824];
        let powerThreshold: number = 6666;

        const Gravity = await hre.ethers.getContractFactory("Gravity");
        const gravity = (await Gravity.deploy(
            hre.ethers.utils.formatBytes32String("gravitytest"),
            powerThreshold,
            constants.VALIDATORS,
            powers
        ));

        await gravity.deployed();
        console.log(`gravity contract deployed at - ${gravity.address}`)

        // Take over the cellar owner so we can transfer ownership
        await hre.network.provider.request({
            method: 'hardhat_impersonateAccount',
            params: [constants.CELLAR_OWNER],
        });
        const cellarSigner = await hre.ethers.getSigner(constants.CELLAR_OWNER);
        const Cellar = hre.ethers.getContractAt(
            'CellarPoolShare',
            constants.CELLAR,
            cellarSigner,
        );
        const cellar = await Cellar;

        let { hash } = await cellar.transferOwnership(gravity.address, {
            gasPrice: hre.ethers.BigNumber.from('99916001694'),
        });
        console.log(
            `Cellar contract at ${cellar.address} is now owned by Gravity contract at ${gravity.address}`,
        );

        // Take over vitalik.eth
        await hre.network.provider.request({
            method: 'hardhat_impersonateAccount',
            params: [constants.WHALE],
        });

        // Send ETH to needed parties
        const whaleSigner = await hre.ethers.getSigner(constants.WHALE);

        for (let addr of constants.VALIDATORS) {
            await whaleSigner.sendTransaction({
                to: addr,
                value: hre.ethers.utils.parseEther('100'),
            });
        }

        await hre.run('node');
    });


/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
    networks: {
        hardhat: {
            forking: {
                url: 'https://mainnet.infura.io/v3/d6f22be0f7fd447186086d2495779003',
                blockNumber: 13357100,
            },
        },
    },
    solidity: {
        compilers: [
            {
                version: '0.6.6',
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 200,
                    },
                },
            },
            {
                version: '0.7.6',
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 200,
                    },
                },
            },
            {
                version: '0.8.0',
                settings: {
                    optimizer: {
                        enabled: true,
                    },
                },
            },
        ],
    },
    typechain: {
        outDir: 'typechain',
        target: 'ethers-v5',
        runOnCompile: true,
    },
    gasReporter: {
        enabled: true,
    },
};
