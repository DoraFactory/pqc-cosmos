package myabciapp

import (
	"encoding/json"

	"github.com/cometbft/cometbft-load-test/pkg/loadtest"
)

type TxType string

const (
    TransferTx TxType = "transfer"
    StakingTx  TxType = "staking"
)

type MyABCIAppClientFactory struct {
    txType TxType
}

type MyABCIAppClient struct {
    txType TxType
    sequence int
}

func (f *MyABCIAppClientFactory) ValidateConfig(cfg loadtest.Config) error {
    return nil
}

func (f *MyABCIAppClientFactory) NewClient(cfg loadtest.Config) (loadtest.Client, error) {
    return &MyABCIAppClient{txType: f.txType}, nil
}

func (c *MyABCIAppClient) GenerateTx() ([]byte, error) {
    c.sequence++
    tx := map[string]interface{}{
        "type": c.txType,
        "data": map[string]interface{}{
            "from":   "sender",
            "to":     "receiver",
            "amount": 1000,
            "nonce":  c.sequence,
        },
    }
    return json.Marshal(tx)
}
