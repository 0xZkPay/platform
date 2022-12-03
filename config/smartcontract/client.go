package smartcontract

import (
	"github.com/0xZkPay/platform/config/envconfig"
	"github.com/0xZkPay/platform/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func GetClient() (*ethclient.Client, error) {
	if client != nil {
		return client, nil
	}
	nodeUrl := envconfig.EnvVars.POLYGON_RPC
	var err error
	client, err = ethclient.Dial(nodeUrl)
	if err != nil {
		logwrapper.Errorf("failed to dial client at url %v, error: %v", nodeUrl, err.Error())
		return nil, err
	}
	return client, nil
}