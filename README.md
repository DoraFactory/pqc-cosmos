# Benchmarking Tendermint vs Tendermint-PQC : CHALLENGE 1

## Overview

This benchmark compares the performance of Tendermint using ED25519 and Dilithium algorithms. The timeout for ED25519 is determined by selecting a random range of 10 block heights while cracking the block. The results are recorded below.

## Benchmark Results

### ED25519 Algorithm

| S.No. | Height | Duration   |
| ----- | ------ | ---------- |
| 1     | 8855   | 989.706 ms |
| 2     | 8856   | 987.923 ms |
| 3     | 8857   | 989.085 ms |
| 4     | 8858   | 987.914 ms |
| 5     | 8859   | 989.865 ms |
| 6     | 8860   | 987.990 ms |
| 7     | 8861   | 987.541 ms |
| 8     | 8862   | 990.055 ms |
| 9     | 8863   | 987.998 ms |
| 10    | 8864   | 990.018 ms |

### Dilithium Algorithm

| S.No. | Height | Duration   |
| ----- | ------ | ---------- |
| 1     | 3471   | 988.35 ms  |
| 2     | 3472   | 991.699 ms |
| 3     | 3473   | 988.405 ms |
| 4     | 3474   | 985.66 ms  |
| 5     | 3475   | 987.89 ms  |
| 6     | 3476   | 990.12 ms  |
| 7     | 3477   | 989.77 ms  |
| 8     | 3478   | 986.45 ms  |
| 9     | 3479   | 987.23 ms  |
| 10    | 3480   | 988.92 ms  |

## Performance Comparison

The timeout average for executed block state, committed state, consensus timeout, received proposal, finalizing commit, and indexed block events is:

- **ED25519 Algorithm:** 988.73 ms
- **Dilithium Algorithm:** 989.99 ms

The performance is comparable between both algorithms based on a random sample of 10 block heights.

## Test Environment

- **Machine Used:** MacBook M3 Pro with 18GB RAM

## Code Review

The project is structured with well-defined and scalable components. However, the following risks were identified:

### Potential Risks

- Multiple `TODO` comments are present, which may expose loose ends that attackers can exploit.
- Deprecated code snippets from Tendermint v0.37 are still in use, particularly in the message protocol, which could pose a security risk.
