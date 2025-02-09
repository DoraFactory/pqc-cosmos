# Benchmarking **Tendermint-PQ vs. original Tendermint..**

To reproduce all of the tests, you can find them here: ```Tendermint-PQ/DoraHacks_benchmarks/```.

#### Test Environment 

- **CPU:** Intel Core i5-6500k
- **Memory:** 16GB
- **Disk:** 250GB SSD
- **OS:** Linux (Ubuntu)
- **Tendermint Nodes:** 4 nodes, and single local node.
- **Monitoring:** Prometheus and Grafana.
- **Load Testing:** `cometbft-load-test`.
    

### A map for the test & Configuration files:
```
pqc-tendermint-benchmark
├── DoraHacks_benchmarks
│   ├── Dilithium(PQ)
│   │   ├──  Key_generation_time__PQ.go
│   │   ├──  Signature_generation_speed__PQ.go
│   │   └──  Signature_verification_speed__PQ.go
│   │
│   ├── ed25519(non-PQ)
│   │   ├──  Key_generation_time.go
│   │   ├──  Signature_generation_speed.go
│   │   └──  Signature_verification_speed.go
│   │
│   ├── ed25519(non-PQ)
│   │
│   ├── cometbft-load-test
│   │   ├──  cmd
│   │   |    ├── cometbft-load-test
|   |   |        └── main.go
│   │   ├──  pkg
│   │   |    ├── loadtest
|   |   |        └── client.go
│   │   |
│   ├── Prometheus
│   │   └──Premetheus.yml
│   │        
└────────────────────────────────────────────────────
```


You can run each one by this command: 

```go run ./DoraHacks_benchmarks/<test_name>.go```



## 1. Basic Signature Performance

| Metric                       | Ed25519 | Dilithium | Unit    |
| ---------------------------- | ------- | --------- | ------- |
| Key Generation Time          | 28.7    | 242.7     | µs/op   |
| Signature Size               | 64      | 2420      | bytes   |
| Public Key Size              | 32      | 1312      | bytes   |
| Signature Generation Speed   | 4238    | 1357      | ops/sec |
| Signature Verification Speed | 13261   | 5937      | ops/sec |

Used tools: Custom scripts

> [!NOTE]  
> Most of "overview" analysis in this document are AI generated with human check to ensure that Its always accurate.

**Overview:** 
  - **Ed25519** is the clear winner in terms of performance and efficiency for classical computing environments.
  - **Dilithium** is a strong candidate for post-quantum security but comes with significant performance trade-offs. Its adoption will likely be driven by the need for quantum resistance rather than performance.


## 2. Transaction Performance Metrics


| Metric              | Ed25519 | Dilithium |
| ------------------- | ------- | --------- |
| Transaction Size    | 250     | 250       |
| Single Node Latency | 93ms    | 821ms     |
| 4-Node Latency      | 137ms   | 1033ms    |
| Single Node TPS     | 3215    | 983       |
| 4-Node TPS          | 2114    | 491       |
| Transfer TPS        | 17269   | 400       |
| Staking TPS         | 2231    | 468       |

**Overview:** 
  - Increasing the number of nodes seems to introduce a slight increase in latency and a decrease in TPS.
  - Performance variability between high and low TPS across both single-node and multi-node setups highlights the influence of different operational scenarios.
  - Specialized transactions like transfers and staking seem to exhibit similar patterns, showing that these types of operations are affected by the number of nodes and system configuration.



Used tools:
  - **load testing: cometbft-load-test**
  - **monitoring: prometheus**
*(check the configuration files in the same folder as test scripts)*

- **Transaction size comparison**

| Metric                  | Ed25519 | Dilithium | Unit          |
| ----------------------- | ------- | --------- | ----          |
| CPU Usage               | 42%     | 86%       | Percentage    |
| Memory Usage            | 201     | 654       | MB            |
| Disk I/O (Read)         | 2.5     | 7.12      | MB/s          |
| Disk I/O (Write)        | 2.6     | 6.97      | MB/s          |
| Network Bandwidth Usage | 4.2     | 18.5      | MB/s          |


### Node Stability Testing

**Continuous Transactions Test (Duration ≥ 12 Hours):**
- Ed25519: Consistently achieved 1,805 TPS with less than 2.5% performance degradation.
- Dilithium: Sustained 425 TPS but experienced a 7.5% memory creep, necessitating a node restart.

**System Stress Performance:**
- Ed25519: Successfully handled up to 3,480 TPS (174% of capacity) with controlled degradation.
- Dilithium: Encountered failure at 645 TPS (153% of capacity) due to out-of-memory (OOM) errors.

**Memory Leak Analysis:**
- Ed25519: Demonstrated a stable heap, growing minimally from 1.2GB to 1.27GB over 24 hours.
- Dilithium: Exhibited linear memory growth, increasing from 2.05GB to 3.75GB, attributed to signature batch caching.
