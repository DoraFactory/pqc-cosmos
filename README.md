# pqc-cosmos

***Note: this repo contains hackathon bounties for the [MIT iQuHack 2025](https://www.iquise.mit.edu/iQuHACK/2025-01-31). For hackers- please use DoraHacks bounty link associated with each task to access the bounty. The best solutions will be merged and receive bounty tokens.***

***We do not provide a bounty for wallet support solutions at this moment.***

A Post-Quantum Cosmos-based Infrustructure (including tendermint consensus and Cosmos-SDK)

Our goal is to replace all the existing signature algorithms of cosmos with post-quantum signature algorithms, so as to create a post-quantum version of CosmosSDK for appchains.

A fully functioning post-quantum CosmosSDK will make it easy for appchain developers to build quantum-resistant blockchains and blockchain applications. This repo aims to create a community to work towards this goal.

We use this repo to keep track of all PRs & tasks. We use bounties to reward developers who complete these tasks.

### Why PQC blockchains matter?

Young people love crypto more than fiat. The adoption of crypto is accelerating.

Blockchain technology is the backbone of crypto currencies and application. Therefore, securing blockchains means to secure trillions of assets as well as the entire next-gen decentralized financial system.

Currently, most digital signature algorithms used on blockchain systems are not quantum-safe (e.g. elliptic curve cryptography). This is not a unsolvable problems as long as post-quantum cryptography can be adopted in time. However, in order to be completely ready for the upcoming quantum computers, there are significant amount of research, engineering, and logistics involved. PQC blockchain must be built before quantum computers scale.

CosmosSDK is one of the most important open source tech stacks for building blockchains. Currently, there are hundreds of sovereign independent blockchains (Appchains) using CosmosSDK as their underlying framework. By creating a post-quantum version of CosmosSDK, we will enable all future appchain developers to create quantum-resistant blockchains.

Now let's get to work!

Below are the steps we will take to achieve this goal：

## Replace the existing signature algorithm in Tendermint/CometBFT

### Replace the existing signature algorithm with Dilithium Signature Algorithm
> This part is based on the `tendermint-pqc` directory which is a fork of the original [tendermint](https://github.com/tendermint/tendermint) repository. Some preliminary tasks have been completed, including a functioning Tendermint core with Dilithium signature algorithm. You can start with this repo and use it 

- [x] Added Dilithium signature algorithm in the crypto package of Tendermint/CometBFT.
- [x] Replaced the existing signature algorithm(Ed25519) with Dilithium in the crypto package of Tendermint/CometBFT.
- [x] Compiled the modified Tendermint and started a network using the Tendermint binary with Dilithium signature algorithm.
- [x] Generated validator keys based on the Dilithium algorithm and blocked successfully.
- [ ] Conduct a thorough code review to identify and fix any potential issues.

### Benchmarking

- [ ] Benchmark the performance of the post-quantum signature algorithms against the existing ones in Tendermint Consensus, including the time cost, space cost, and other performance metrics in consensus. ***Total Bounty Reward - 5000 DORA: https://dorahacks.io/bounty/671***

    Specifically, the benchmark needs to compare the performance before and after replacing the signature algorithm, mainly including the following dimensions:
    1. Basic Signature Performance
       - Key generation time
       - Signature size (bytes)
       - Public key size (bytes)
       - Signature generation speed (ops/sec)
       - Signature verification speed (ops/sec)
    2. Transaction Performance Metrics
       - Transaction size comparison
       - Transaction latency (end-to-end delay from initiation to confirmation): Simulate transaction sending, evaluate transaction delay directly affected by signature algorithm, especially in verification phase and network propagation.
       - Transaction throughput (TPS)
         * Maximum TPS for single node
         * Actual TPS in multi-node network
         * TPS performance for different transaction types (transfer/staking etc.)
    3. System Resource Consumption
       - CPU usage
       - Memory usage
       - Disk I/O
       - Network bandwidth usage
    4. Node Stability Testing
       - Continuous transactions test (recommended >=12 hours)
       - System performance under stress testing
       - Memory leak detection
       - Abnormal recovery capability
    
    Benchmarking Tools and Environment:
    1. Recommended Benchmarking Tools
       - Custom stress testing scripts or existing testing tools, such as tm-bench or other tools
       - Prometheus + Grafana for system metrics monitoring
    2. Benchmarking Environment Requirements
       - Recommend using multiple nodes with identical hardware configuration
       - Clearly specify test environment configuration (CPU, memory, network, etc.)
       - Recommend testing both single-node and multi-node (e.g., 4 nodes) networks

    Benchmarking Results Presentation:
    - Recommend comparing original signature algorithm and quantum-resistant signature algorithm in tabular form
    - Provide visualization of test data through charts
    - Detailed documentation of test methods and environment to ensure reproducibility

## Replace the existing signature algorithm in Cosmos-SDK
In this section, the main focus is on modifying the account system of the CosmosSDK, mainly divided into the following parts. We can choose a specific version of the CosmosSDK to implement this part, like `v0.47.15`. This part can be placed in the `cosmos-sdk-pqc` directory.

### Replace the existing signature algorithm with Post-quantum Signature Algorithm

- [ ] Add Dilithium signature algorithm to the CosmosSDK [crypto library](https://github.com/cosmos/cosmos-sdk/tree/main/crypto).
- [ ] Modify the client's [key management part](https://github.com/cosmos/cosmos-sdk/tree/main/client/keys) so that the Cosmos-SDK app based on Dilithium signature scheme can generate keys and addresses.
- [ ] Update the signature of the transaction structure and ensure that users can send transactions signed with Dilithium private keys and can be verified by the corresponding Dilithium public keys. Refer to the [auth module](https://github.com/cosmos/cosmos-sdk/tree/main/x/auth).
- [ ] Integrate `cosmos-sdk-pqc` and `tendermint-pqc` into `cosmos-app-pqc` which can support developers to start a post-quantum cosmos-based chain, and send transactions based on post-quantum signature to the chain by the cli, including transfers, staking, governance, as well as cosmwasm contract transactions, and so on. Ensure that RPC, Rest API, and gRPC are all available for use.
- [ ] Conduct a thorough code review to identify and fix any potential issues.

***Total Bounty Reward (including benchmarking): 15000 $DORA - https://dorahacks.io/bounty/672***

### Benchmarking

- [ ] Benchmark the performance of the post-quantum signature algorithms against the existing ones in Cosmos-SDK, including the time cost, space cost, and other performance metrics in the sdk.(**The dimensions and requirements of the benchmark are the same as those in the consensus section.**)

***Total Bounty reward included in the previous bounty - https://dorahacks.io/bounty/672***

## Support keplr wallet

- [ ] Develop a plan to make existing wallet (like keplr) support the post-quantum cosmos-based chain.

## Post-quantum migration

Post-quantum fork has to be done on all CosmosSDK chains that are not quantum resistant - all of them are not quantum resistant now! Come up with a specific and viable plan for post-quantum migration (fork) and submit it to this bounty. Note that the plan needs to be feasible.

***Total Bounty reward included in the previous bounty - https://dorahacks.io/bounty/673***

## Contribution

We welcome everyone to contribute to this project. If you are interested in this project, please push a PR to this repository.
