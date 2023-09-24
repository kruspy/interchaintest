package icq

import (
	"context"
	"strings"

	"go.uber.org/zap"

	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/relayer"
)

var _ relayer.RelayerCommander = &commander{}

type commander struct {
	log             *zap.Logger
	extraStartFlags []string
}

func (c commander) Name() string {
	return icq
}

func (c commander) DefaultContainerImage() string {
	return defaultContainerImage
}

func (c commander) DefaultContainerVersion() string {
	return DefaultContainerVersion
}

func (c commander) DockerUser() string {
	return icqDefaultUidGid
}

func (c commander) ConfigContent(ctx context.Context, cfg ibc.ChainConfig, keyName, rpcAddr, grpcAddr string) ([]byte, error) {
	panic("Implemented in the ICQ relayer")
}

func (c commander) ParseAddKeyOutput(stdout, stderr string) (ibc.Wallet, error) {
	panic("Implemented in the ICQ relayer")
}

func (c commander) ParseRestoreKeyOutput(stdout, stderr string) string {
	return strings.Replace(stdout, "\n", "", 1)
}

func (c commander) ParseGetChannelsOutput(stdout, stderr string) ([]ibc.ChannelOutput, error) {
	panic("Not implemented")
}

func (c commander) ParseGetConnectionsOutput(stdout, stderr string) (ibc.ConnectionOutputs, error) {
	panic("Not implemented")
}

func (c commander) ParseGetClientsOutput(stdout, stderr string) (ibc.ClientOutputs, error) {
	panic("Not implemented")
}

func (c commander) Init(homeDir string) []string {
	return []string{}
}

func (c commander) AddChainConfiguration(containerFilePath, homeDir string) []string {
	panic("Implemented in the ICQ relayer")
}

func (c commander) AddKey(chainID, keyName, coinType, homeDir string) []string {
	return []string{}
}

func (commander) CreateChannel(pathName string, opts ibc.CreateChannelOptions, homeDir string) []string {
	return []string{}
}

func (commander) CreateClients(pathName string, opts ibc.CreateClientOptions, homeDir string) []string {
	return []string{}
}

func (commander) CreateConnections(pathName string, homeDir string) []string {
	return []string{}
}

func (commander) Flush(pathName, channelID, homeDir string) []string {
	return []string{}
}

func (commander) GeneratePath(srcChainID, dstChainID, pathName, homeDir string) []string {
	return []string{}
}

func (commander) UpdatePath(pathName, homeDir string, filter ibc.ChannelFilter) []string {
	return []string{}
}

func (commander) GetChannels(chainID, homeDir string) []string {
	return []string{}
}

func (commander) GetConnections(chainID, homeDir string) []string {
	return []string{}
}

func (commander) GetClients(chainID, homeDir string) []string {
	return []string{}
}

func (commander) LinkPath(pathName, homeDir string, channelOpts ibc.CreateChannelOptions, clientOpt ibc.CreateClientOptions) []string {
	return []string{}
}

func (c commander) RestoreKey(chainID, keyName, coinType, mnemonic, homeDir string) []string {
	return []string{icq, "keys", "restore", keyName, mnemonic, "--chain", chainID, "--home", homeDir}
}

func (c commander) StartRelayer(homeDir string, pathNames ...string) []string {
	return []string{icq, "run", "--home", homeDir}
}

func (commander) UpdateClients(pathName, homeDir string) []string {
	return []string{}
}

func (c commander) CreateWallet(keyName, address, mnemonic string) ibc.Wallet {
	return NewWallet(keyName, address, mnemonic)
}
