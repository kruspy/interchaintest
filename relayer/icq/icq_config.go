package icq

import (
	"fmt"
)

// NewConfig returns a icq Config with an entry for each of the provided ChainConfigs.
func NewConfig(chainConfigs ...ChainConfig) Config {
	chains := make(map[string]Chain, 0)
	var defaultChainID string
	for _, icqCfg := range chainConfigs {
		chainCfg := icqCfg.cfg

		if defaultChainID == "" {
			defaultChainID = icqCfg.keyName
		}

		chains[chainCfg.ChainID] = Chain{
			Key:            icqCfg.keyName,
			ChainId:        chainCfg.ChainID,
			RPCAddr:        icqCfg.rpcAddr,
			GrpcAddr:       fmt.Sprintf("http://%s", icqCfg.grpcAddr),
			AccountPrefix:  chainCfg.Bech32Prefix,
			KeyRingBackend: "test",
			GasAdjustment:  10,
			GasPrices:      chainCfg.GasPrices,
			KeyDirectory:   icqKeyPath,
			Debug:          false,
			Timeout:        "20s",
			BlockTimeout:   "",
			OutputFormat:   "json",
			SignMode:       "direct",
		}
	}

	return Config{
		DefaultChain: defaultChainID, // using the first chain as the default one
		Chains:       chains,
	}
}

type Config struct {
	DefaultChain string           `yaml:"default_chain"`
	Chains       map[string]Chain `yaml:"chains"`
}

type Chain struct {
	Key            string `yaml:"key"`
	ChainId        string `yaml:"chain-id"`
	RPCAddr        string `yaml:"rpc-addr"`
	GrpcAddr       string `yaml:"grpc-addr"`
	AccountPrefix  string `yaml:"account-prefix"`
	KeyRingBackend string `yaml:"keyring-backend"`
	GasAdjustment  int    `yaml:"gas-adjustment"`
	GasPrices      string `yaml:"gas-prices"`
	KeyDirectory   string `yaml:"key-directory"`
	Debug          bool   `yaml:"debug"`
	Timeout        string `yaml:"timeout"`
	BlockTimeout   string `yaml:"block-timeout"`
	OutputFormat   string `yaml:"output-format"`
	SignMode       string `yaml:"sign-mode"`
}
