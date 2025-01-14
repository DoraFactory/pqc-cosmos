# pqc-cosmos
A Post-Quantum Cosmos-based Infrustructure(include tendermint consensus and Cosmos-SDK)

Our goal is to replace all the existing signature algorithms of cosmos with post-quantum signature algorithms, so as to realize the post-quantum cosmos app.

Below are the steps we will take to achieve this goal：

## Replace the existing signature algorithm in Tendermint/CometBFT

### Replace the existing signature algorithm with Dilithium Signature Algorithm
> This part is based on the `tendermint-pqc` directory which is a fork of the original [tendermint](https://github.com/tendermint/tendermint) repository.

- [x] Added Dilithium signature algorithm in the crypto package of Tendermint/CometBFT.
- [x] Replaced the existing signature algorithm(Ed25519) with Dilithium in the crypto package of Tendermint/CometBFT.
- [x] Compiled the modified Tendermint and started a network using the Tendermint binary with Dilithium signature algorithm.
- [x] Generated validator keys based on the Dilithium algorithm and blocked successfully.
- [ ] Conduct a thorough code review to identify and fix any potential issues.


### Replace the existing signature algorithm with other post-quantum signature algorithms

The goal is to replace the existing signature algorithms with other post-quantum signature algorithms. This involves the following steps:

- [ ] Implement the other post-quantum signature algorithms in the relevant crypto packages.
- [ ] Replace the existing signature algorithms with the other post-quantum signature algorithms in the codebase.
- [ ] Compile the modified Tendermint and started a network using the Tendermint binary with other post-quantum signature algorithms.
- [ ] Generate validator keys based on the other post-quantum signature algorithms and blocked successfully.
- [ ] Test the integration to ensure compatibility and functionality.
- [ ] Conduct a thorough code review to identify and fix any potential issues.

### Benchmark

- [ ] Benchmark the performance of the post-quantum signature algorithms against the existing ones in Tendermint Consensus, including the time cost, space cost, and other performance metrics in consensus.

## Replace the existing signature algorithm in Cosmos-SDK
In this section, the main focus is on modifying the account system of the cosmos sdk, mainly divided into the following parts. We can choose a specific version of the cosmos sdk to implement this part, like `v0.47.15`. This part can be placed in the `cosmos-sdk-pqc` directory.

### Replace the existing signature algorithm with Post-quantum Signature Algorithm

- [ ] Add the Post-quantum Signature Algorithm to the existing cosmos sdk [crypto library](https://github.com/cosmos/cosmos-sdk/tree/main/crypto).
- [ ] Modify the client's [key management part](https://github.com/cosmos/cosmos-sdk/tree/main/client/keys) so that the Cosmos-SDK app based on the Post-quantum Signature Algorithm can generate keys and corresponding addresses.
- [ ] Update the signature of the transaction structure and ensure that users can send transactions signed with Dilithium private keys and can be verified by the corresponding Dilithium public keys. Refer to the [auth module](https://github.com/cosmos/cosmos-sdk/tree/main/x/auth).
- [ ] Integrate `cosmos-sdk-pqc` and `tendermint-pqc` into `cosmos-app-pqc` which can support developers to start a post-quantum cosmos-based chain, and send transactions based on post-quantum signature to the chain by the cli, including transfers, staking, governance, as well as cosmwasm contract transactions, and so on. Ensure that RPC, Rest API, and gRPC are all available for use.
- [ ] Conduct a thorough code review to identify and fix any potential issues.

### Benchmark

- [ ] Benchmark the performance of the post-quantum signature algorithms against the existing ones in Cosmos-SDK, including the time cost, space cost, and other performance metrics in the sdk.

## Support keplr wallet

- [ ] Make the existing wallet(like keplr) support the post-quantum cosmos-based chain.

## Post-quantum migration

The above implementation is based on a brand new post-quantum cosmos-based chain, so how can we migrate the existing cosmos-based chain to a post-quantum cosmos-based chain? This is what we need to achieve in this part. We hope to achieve post-quantum migration without affecting the existing cosmos-based chain.

- [ ] TBD...

## Contribution

We welcome everyone to contribute to this project. If you are interested in this project, please push pr to this repository.
