package icq

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/relayer"
)

const (
	icq                     = "interchain-queries"
	defaultContainerImage   = "marcpt/interchain-queries"
	DefaultContainerVersion = "latest-amd64"

	icqDefaultUidGid = "1000:1000"
	icqHome          = "/home/icq"
	icqConfigPath    = ".icq/config.yaml"
	icqKeyPath       = ".icq/keys"
)

var _ ibc.Relayer = &Relayer{}

// Relayer is the ibc.Relayer implementation for icq.
type Relayer struct {
	*relayer.DockerRelayer
	chainConfigs []ChainConfig
}

// ChainConfig holds all values required to write an entry in the "chains" section in the icq config file.
type ChainConfig struct {
	cfg                        ibc.ChainConfig
	keyName, rpcAddr, grpcAddr string
}

// NewIcqRelayer returns a new icq relayer.
func NewIcqRelayer(
	log *zap.Logger,
	testName string,
	cli *client.Client,
	networkID string,
	options ...relayer.RelayerOption,
) *Relayer {
	c := commander{log: log}
	for _, opt := range options {
		switch o := opt.(type) {
		case relayer.RelayerOptionExtraStartFlags:
			c.extraStartFlags = o.Flags
		}
	}
	options = append(options, relayer.HomeDir(icqHome))
	dr, err := relayer.NewDockerRelayer(context.TODO(), log, testName, cli, networkID, c, options...)
	if err != nil {
		panic(err)
	}

	return &Relayer{
		DockerRelayer: dr,
	}
}

// AddChainConfiguration is called once per chain configuration, which means that in the case of icq, the single
// config file is overwritten with a new entry each time this function is called.
func (r *Relayer) AddChainConfiguration(
	ctx context.Context,
	rep ibc.RelayerExecReporter,
	chainConfig ibc.ChainConfig,
	keyName, rpcAddr, grpcAddr string,
) error {
	configContent, err := r.configContent(chainConfig, keyName, rpcAddr, grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to generate config content: %w", err)
	}

	if err := r.WriteFileToHomeDir(ctx, icqConfigPath, configContent); err != nil {
		return fmt.Errorf("failed to write icq config: %w", err)
	}

	return nil
}

// configContent returns the contents of the icq config file as a byte array. Note: as icq expects a single file
// rather than multiple config files, we need to maintain a list of chain configs each time they are added to write the
// full correct file update calling Relayer.AddChainConfiguration.
func (r *Relayer) configContent(cfg ibc.ChainConfig, keyName, rpcAddr, grpcAddr string) ([]byte, error) {
	r.chainConfigs = append(r.chainConfigs, ChainConfig{
		cfg:      cfg,
		keyName:  keyName,
		rpcAddr:  rpcAddr,
		grpcAddr: grpcAddr,
	})
	icqConfig := NewConfig(r.chainConfigs...)
	bz, err := yaml.Marshal(icqConfig)
	if err != nil {
		return nil, err
	}
	return bz, nil
}
