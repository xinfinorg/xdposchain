// Copyright 2015 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// Package utils contains internal helper functions for go-ethereum commands.
package utils

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/XinFinOrg/XDPoSChain/eth/ethconfig"
	"github.com/XinFinOrg/XDPoSChain/internal/flags"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/XinFinOrg/XDPoSChain/XDCx"
	"github.com/XinFinOrg/XDPoSChain/accounts"
	"github.com/XinFinOrg/XDPoSChain/accounts/keystore"
	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/common/fdlimit"
	"github.com/XinFinOrg/XDPoSChain/consensus/XDPoS"
	"github.com/XinFinOrg/XDPoSChain/core"
	"github.com/XinFinOrg/XDPoSChain/core/vm"
	"github.com/XinFinOrg/XDPoSChain/crypto"
	"github.com/XinFinOrg/XDPoSChain/eth"
	"github.com/XinFinOrg/XDPoSChain/eth/downloader"
	"github.com/XinFinOrg/XDPoSChain/eth/gasprice"
	"github.com/XinFinOrg/XDPoSChain/ethdb"
	"github.com/XinFinOrg/XDPoSChain/log"
	"github.com/XinFinOrg/XDPoSChain/metrics"
	"github.com/XinFinOrg/XDPoSChain/metrics/exp"
	"github.com/XinFinOrg/XDPoSChain/node"
	"github.com/XinFinOrg/XDPoSChain/p2p"
	"github.com/XinFinOrg/XDPoSChain/p2p/discover"
	"github.com/XinFinOrg/XDPoSChain/p2p/discv5"
	"github.com/XinFinOrg/XDPoSChain/p2p/nat"
	"github.com/XinFinOrg/XDPoSChain/p2p/netutil"
	"github.com/XinFinOrg/XDPoSChain/params"
	whisper "github.com/XinFinOrg/XDPoSChain/whisper/whisperv6"
	"github.com/urfave/cli/v2"
)

var (
	CommandHelpTemplate = `{{.cmd.Name}}{{if .cmd.Subcommands}} command{{end}}{{if .cmd.Flags}} [command options]{{end}} [arguments...]
{{if .cmd.Description}}{{.cmd.Description}}
{{end}}{{if .cmd.Subcommands}}
SUBCOMMANDS:
	{{range .cmd.Subcommands}}{{.cmd.Name}}{{with .cmd.ShortName}}, {{.cmd}}{{end}}{{ "\t" }}{{.cmd.Usage}}
	{{end}}{{end}}{{if .categorizedFlags}}
{{range $idx, $categorized := .categorizedFlags}}{{$categorized.Name}} OPTIONS:
{{range $categorized.Flags}}{{"\t"}}{{.}}
{{end}}
{{end}}{{end}}`
)

func init() {
	cli.AppHelpTemplate = `{{.Name}} {{if .Flags}}[global options] {{end}}command{{if .Flags}} [command options]{{end}} [arguments...]

VERSION:
   {{.Version}}

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`

	cli.CommandHelpTemplate = CommandHelpTemplate
}

// NewApp creates an app with sane defaults.
func NewApp(gitCommit, usage string) *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Authors = make([]*cli.Author, 1)
	//app.Authors = nil
	app.Version = params.Version
	if len(gitCommit) >= 8 {
		app.Version += "-" + gitCommit[:8]
	}
	app.Usage = usage
	return app
}

// These are all the command line flags we support.
// If you add to this list, please remember to include the
// flag in the appropriate command definition.
//
// The flags are defined here so their names and help texts
// are the same for all commands.

var (
	// XDC flags.
	RollbackFlag = &cli.StringFlag{
		Name:  "rollback",
		Usage: "Rollback chain at hash",
		Value: "",
	}
	// General settings
	AnnounceTxsFlag = &cli.BoolFlag{
		Name:  "announce-txs",
		Usage: "Always commit transactions",
	}
	StoreRewardFlag = &cli.BoolFlag{
		Name:  "store-reward",
		Usage: "Store reward to file",
	}
	DataDirFlag = &flags.DirectoryFlag{
		Name:  "datadir",
		Usage: "Data directory for the databases and keystore",
		Value: flags.DirectoryString(node.DefaultDataDir()),
	}
	KeyStoreDirFlag = &flags.DirectoryFlag{
		Name:  "keystore",
		Usage: "Directory for the keystore (default = inside the datadir)",
	}
	NoUSBFlag = &cli.BoolFlag{
		Name:  "nousb",
		Usage: "Disables monitoring for and managing USB hardware wallets",
	}
	NetworkIdFlag = &cli.Uint64Flag{
		Name:  "networkid",
		Usage: "Network identifier (integer, 89=XDPoSChain)",
		Value: eth.DefaultConfig.NetworkId,
	}
	TestnetFlag = &cli.BoolFlag{
		Name:  "testnet",
		Usage: "Ropsten network: pre-configured proof-of-work test network",
	}
	XDCTestnetFlag = &cli.BoolFlag{
		Name:  "apothem",
		Usage: "XDC Apothem Network",
	}
	RinkebyFlag = &cli.BoolFlag{
		Name:  "rinkeby",
		Usage: "Rinkeby network: pre-configured proof-of-authority test network",
	}
	DeveloperFlag = &cli.BoolFlag{
		Name:  "dev",
		Usage: "Ephemeral proof-of-authority network with a pre-funded developer account, mining enabled",
	}
	DeveloperPeriodFlag = &cli.IntFlag{
		Name:  "dev.period",
		Usage: "Block period to use in developer mode (0 = mine only if transaction pending)",
	}
	IdentityFlag = &cli.StringFlag{
		Name:  "identity",
		Usage: "Custom node name",
	}
	DocRootFlag = &flags.DirectoryFlag{
		Name:  "docroot",
		Usage: "Document Root for HTTPClient file scheme",
		Value: flags.DirectoryString(flags.HomeDir()),
	}
	FastSyncFlag = &cli.BoolFlag{
		Name:  "fast",
		Usage: "Enable fast syncing through state downloads",
	}
	LightModeFlag = &cli.BoolFlag{
		Name:  "light",
		Usage: "Enable light client mode",
	}
	defaultSyncMode = eth.DefaultConfig.SyncMode
	SyncModeFlag    = &flags.TextMarshalerFlag{
		Name:  "syncmode",
		Usage: `Blockchain sync mode ("fast", "full", or "light")`,
		Value: &defaultSyncMode,
	}
	GCModeFlag = &cli.StringFlag{
		Name:  "gcmode",
		Usage: `Blockchain garbage collection mode ("full", "archive")`,
		Value: "full",
	}
	LightServFlag = &cli.IntFlag{
		Name:  "lightserv",
		Usage: "Maximum percentage of time allowed for serving LES requests (0-90)",
		Value: 0,
	}
	LightPeersFlag = &cli.IntFlag{
		Name:  "lightpeers",
		Usage: "Maximum number of LES client peers",
		Value: eth.DefaultConfig.LightPeers,
	}
	LightKDFFlag = &cli.BoolFlag{
		Name:  "lightkdf",
		Usage: "Reduce key-derivation RAM & CPU usage at some expense of KDF strength",
	}
	// XDCX settings
	XDCXEnabledFlag = &cli.BoolFlag{
		Name:  "XDCx",
		Usage: "Enable the XDCX protocol",
	}
	// Ethash settings
	EthashCacheDirFlag = &DirectoryFlag{
		Name:  "ethash.cachedir",
		Usage: "Directory to store the ethash verification caches (default = inside the datadir)",
	}
	// Transaction pool settings
	TxPoolNoLocalsFlag = &cli.BoolFlag{
		Name:  "txpool.nolocals",
		Usage: "Disables price exemptions for locally submitted transactions",
	}
	TxPoolJournalFlag = &cli.StringFlag{
		Name:  "txpool.journal",
		Usage: "Disk journal for local transaction to survive node restarts",
		Value: core.DefaultTxPoolConfig.Journal,
	}
	TxPoolRejournalFlag = &cli.DurationFlag{
		Name:  "txpool.rejournal",
		Usage: "Time interval to regenerate the local transaction journal",
		Value: core.DefaultTxPoolConfig.Rejournal,
	}
	TxPoolPriceLimitFlag = &cli.Uint64Flag{
		Name:  "txpool.pricelimit",
		Usage: "Minimum gas price limit to enforce for acceptance into the pool",
		Value: eth.DefaultConfig.TxPool.PriceLimit,
	}
	TxPoolPriceBumpFlag = &cli.Uint64Flag{
		Name:  "txpool.pricebump",
		Usage: "Price bump percentage to replace an already existing transaction",
		Value: eth.DefaultConfig.TxPool.PriceBump,
	}
	TxPoolAccountSlotsFlag = &cli.Uint64Flag{
		Name:  "txpool.accountslots",
		Usage: "Minimum number of executable transaction slots guaranteed per account",
		Value: eth.DefaultConfig.TxPool.AccountSlots,
	}
	TxPoolGlobalSlotsFlag = &cli.Uint64Flag{
		Name:  "txpool.globalslots",
		Usage: "Maximum number of executable transaction slots for all accounts",
		Value: eth.DefaultConfig.TxPool.GlobalSlots,
	}
	TxPoolAccountQueueFlag = &cli.Uint64Flag{
		Name:  "txpool.accountqueue",
		Usage: "Maximum number of non-executable transaction slots permitted per account",
		Value: eth.DefaultConfig.TxPool.AccountQueue,
	}
	TxPoolGlobalQueueFlag = &cli.Uint64Flag{
		Name:  "txpool.globalqueue",
		Usage: "Maximum number of non-executable transaction slots for all accounts",
		Value: eth.DefaultConfig.TxPool.GlobalQueue,
	}
	TxPoolLifetimeFlag = &cli.DurationFlag{
		Name:  "txpool.lifetime",
		Usage: "Maximum amount of time non-executable transaction are queued",
		Value: eth.DefaultConfig.TxPool.Lifetime,
	}
	// Performance tuning settings
	CacheFlag = &cli.IntFlag{
		Name:  "cache",
		Usage: "Megabytes of memory allocated to internal caching",
		Value: 1024,
	}
	CacheDatabaseFlag = &cli.IntFlag{
		Name:  "cache.database",
		Usage: "Percentage of cache memory allowance to use for database io",
		Value: 75,
	}
	CacheGCFlag = &cli.IntFlag{
		Name:  "cache.gc",
		Usage: "Percentage of cache memory allowance to use for trie pruning",
		Value: 25,
	}
	// Miner settings
	StakingEnabledFlag = &cli.BoolFlag{
		Name:  "mine",
		Usage: "Enable staking",
	}
	StakerThreadsFlag = &cli.IntFlag{
		Name:  "minerthreads",
		Usage: "Number of CPU threads to use for staking",
		Value: runtime.NumCPU(),
	}
	TargetGasLimitFlag = &cli.Uint64Flag{
		Name:  "targetgaslimit",
		Usage: "Target gas limit sets the artificial target gas floor for the blocks to mine",
		Value: params.XDCGenesisGasLimit,
	}
	EtherbaseFlag = &cli.StringFlag{
		Name:  "etherbase",
		Usage: "Public address for block mining rewards (default = first account created)",
		Value: "0",
	}
	GasPriceFlag = &flags.BigFlag{
		Name:  "gasprice",
		Usage: "Minimal gas price to accept for mining a transactions",
		Value: eth.DefaultConfig.GasPrice,
	}
	ExtraDataFlag = &cli.StringFlag{
		Name:  "extradata",
		Usage: "Block extra data set by the miner (default = client version)",
	}
	// Account settings
	UnlockedAccountFlag = &cli.StringFlag{
		Name:  "unlock",
		Usage: "Comma separated list of accounts to unlock",
		Value: "",
	}
	PasswordFileFlag = &cli.StringFlag{
		Name:  "password",
		Usage: "Password file to use for non-interactive password input",
		Value: "",
	}

	VMEnableDebugFlag = &cli.BoolFlag{
		Name:  "vmdebug",
		Usage: "Record information useful for VM and contract debugging",
	}
	// Logging and debug settings
	EthStatsURLFlag = &cli.StringFlag{
		Name:  "ethstats",
		Usage: "Reporting URL of a ethstats service (nodename:secret@host:port)",
	}
	// MetricsHTTPFlag defines the endpoint for a stand-alone metrics HTTP endpoint.
	// Since the pprof service enables sensitive/vulnerable behavior, this allows a user
	// to enable a public-OK metrics endpoint without having to worry about ALSO exposing
	// other profiling behavior or information.
	MetricsHTTPFlag = &cli.StringFlag{
		Name:  "metrics.addr",
		Usage: "Enable stand-alone metrics HTTP server listening interface",
		Value: metrics.DefaultConfig.HTTP,
	}
	MetricsPortFlag = &cli.IntFlag{
		Name:  "metrics.port",
		Usage: "Metrics HTTP server listening port",
		Value: metrics.DefaultConfig.Port,
	}
	MetricsEnabledFlag = &cli.BoolFlag{
		Name:  "metrics",
		Usage: "Enable metrics collection and reporting",
		Value: metrics.DefaultConfig.MetricsEnabledFlag,
	}
	FakePoWFlag = &cli.BoolFlag{
		Name:  "fakepow",
		Usage: "Disables proof-of-work verification",
	}
	NoCompactionFlag = &cli.BoolFlag{
		Name:  "nocompaction",
		Usage: "Disables db compaction after import",
	}
	// RPC settings
	RPCEnabledFlag = &cli.BoolFlag{
		Name:  "rpc",
		Usage: "Enable the HTTP-RPC server",
	}
	RPCListenAddrFlag = &cli.StringFlag{
		Name:  "rpcaddr",
		Usage: "HTTP-RPC server listening interface",
		Value: node.DefaultHTTPHost,
	}
	RewoundFlag = &cli.IntFlag{
		Name:  "rewound",
		Usage: "Rewound blocks",
		Value: 0,
	}
	RPCPortFlag = &cli.IntFlag{
		Name:  "rpcport",
		Usage: "HTTP-RPC server listening port",
		Value: node.DefaultHTTPPort,
	}
	RPCCORSDomainFlag = &cli.StringFlag{
		Name:  "rpccorsdomain",
		Usage: "Comma separated list of domains from which to accept cross origin requests (browser enforced)",
		Value: "",
	}
	RPCVirtualHostsFlag = &cli.StringFlag{
		Name:  "rpcvhosts",
		Usage: "Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard.",
		Value: strings.Join(node.DefaultConfig.HTTPVirtualHosts, ","),
	}
	RPCApiFlag = &cli.StringFlag{
		Name:  "rpcapi",
		Usage: "API's offered over the HTTP-RPC interface",
		Value: "",
	}
	IPCDisabledFlag = &cli.BoolFlag{
		Name:  "ipcdisable",
		Usage: "Disable the IPC-RPC server",
	}
	IPCPathFlag = &flags.DirectoryFlag{
		Name:  "ipcpath",
		Usage: "Filename for IPC socket/pipe within the datadir (explicit paths escape it)",
	}
	WSEnabledFlag = &cli.BoolFlag{
		Name:  "ws",
		Usage: "Enable the WS-RPC server",
	}
	WSListenAddrFlag = &cli.StringFlag{
		Name:  "wsaddr",
		Usage: "WS-RPC server listening interface",
		Value: node.DefaultWSHost,
	}
	WSPortFlag = &cli.IntFlag{
		Name:  "wsport",
		Usage: "WS-RPC server listening port",
		Value: node.DefaultWSPort,
	}
	WSApiFlag = &cli.StringFlag{
		Name:  "wsapi",
		Usage: "API's offered over the WS-RPC interface",
		Value: "",
	}
	WSAllowedOriginsFlag = &cli.StringFlag{
		Name:  "wsorigins",
		Usage: "Origins from which to accept websockets requests",
		Value: "",
	}
	ExecFlag = &cli.StringFlag{
		Name:  "exec",
		Usage: "Execute JavaScript statement",
	}
	PreloadJSFlag = &cli.StringFlag{
		Name:  "preload",
		Usage: "Comma separated list of JavaScript files to preload into the console",
	}

	// Network Settings
	MaxPeersFlag = &cli.IntFlag{
		Name:  "maxpeers",
		Usage: "Maximum number of network peers (network disabled if set to 0)",
		Value: 25,
	}
	MaxPendingPeersFlag = &cli.IntFlag{
		Name:  "maxpendpeers",
		Usage: "Maximum number of pending connection attempts (defaults used if set to 0)",
		Value: 0,
	}
	ListenPortFlag = &cli.IntFlag{
		Name:  "port",
		Usage: "Network listening port",
		Value: 30303,
	}
	BootnodesFlag = &cli.StringFlag{
		Name:  "bootnodes",
		Usage: "Comma separated enode URLs for P2P discovery bootstrap (set v4+v5 instead for light servers)",
		Value: "",
	}
	BootnodesV4Flag = &cli.StringFlag{
		Name:  "bootnodesv4",
		Usage: "Comma separated enode URLs for P2P v4 discovery bootstrap (light server, full nodes)",
		Value: "",
	}
	BootnodesV5Flag = &cli.StringFlag{
		Name:  "bootnodesv5",
		Usage: "Comma separated enode URLs for P2P v5 discovery bootstrap (light server, light nodes)",
		Value: "",
	}
	NodeKeyFileFlag = &cli.StringFlag{
		Name:  "nodekey",
		Usage: "P2P node key file",
	}
	NodeKeyHexFlag = &cli.StringFlag{
		Name:  "nodekeyhex",
		Usage: "P2P node key as hex (for testing)",
	}
	NATFlag = &cli.StringFlag{
		Name:  "nat",
		Usage: "NAT port mapping mechanism (any|none|upnp|pmp|extip:<IP>)",
		Value: "any",
	}
	NoDiscoverFlag = &cli.BoolFlag{
		Name:  "nodiscover",
		Usage: "Disables the peer discovery mechanism (manual peer addition)",
	}
	DiscoveryV5Flag = &cli.BoolFlag{
		Name:  "v5disc",
		Usage: "Enables the experimental RLPx V5 (Topic Discovery) mechanism",
	}
	NetrestrictFlag = &cli.StringFlag{
		Name:  "netrestrict",
		Usage: "Restricts network communication to the given IP networks (CIDR masks)",
	}

	// ATM the url is left to the user and deployment to
	JSpathFlag = &cli.StringFlag{
		Name:  "jspath",
		Usage: "JavaScript root path for `loadScript`",
		Value: ".",
	}

	// Gas price oracle settings
	GpoBlocksFlag = &cli.IntFlag{
		Name:  "gpoblocks",
		Usage: "Number of recent blocks to check for gas prices",
		Value: eth.DefaultConfig.GPO.Blocks,
	}
	GpoPercentileFlag = &cli.IntFlag{
		Name:  "gpopercentile",
		Usage: "Suggested gas price is the given percentile of a set of recent transaction gas prices",
		Value: eth.DefaultConfig.GPO.Percentile,
	}
	WhisperEnabledFlag = &cli.BoolFlag{
		Name:  "shh",
		Usage: "Enable Whisper",
	}
	WhisperMaxMessageSizeFlag = &cli.IntFlag{
		Name:  "shh.maxmessagesize",
		Usage: "Max message size accepted",
		Value: int(whisper.DefaultMaxMessageSize),
	}
	WhisperMinPOWFlag = &cli.Float64Flag{
		Name:  "shh.pow",
		Usage: "Minimum POW accepted",
		Value: whisper.DefaultMinimumPoW,
	}
	XDCXDataDirFlag = &flags.DirectoryFlag{
		Name:  "XDCx.datadir",
		Usage: "Data directory for the XDCX databases",
		Value: flags.DirectoryString(filepath.Join(DataDirFlag.Value.String(), "XDCx")),
	}
	XDCXDBEngineFlag = &cli.StringFlag{
		Name:  "XDCx.dbengine",
		Usage: "Database engine for XDCX (leveldb, mongodb)",
		Value: "leveldb",
	}
	XDCXDBNameFlag = &cli.StringFlag{
		Name:  "XDCx.dbName",
		Usage: "Database name for XDCX",
		Value: "XDCdex",
	}
	XDCXDBConnectionUrlFlag = &cli.StringFlag{
		Name:  "XDCx.dbConnectionUrl",
		Usage: "ConnectionUrl to database if dbEngine is mongodb. Host:port. If there are multiple instances, separated by comma. Eg: localhost:27017,localhost:27018",
		Value: "localhost:27017",
	}
	XDCXDBReplicaSetNameFlag = &cli.StringFlag{
		Name:  "XDCx.dbReplicaSetName",
		Usage: "ReplicaSetName if Master-Slave is setup",
	}
	XDCSlaveModeFlag = &cli.BoolFlag{
		Name:  "slave",
		Usage: "Enable slave mode",
	}
)

// MakeDataDir retrieves the currently requested data directory, terminating
// if none (or the empty string) is specified. If the node is starting a testnet,
// the a subdirectory of the specified datadir will be used.
func MakeDataDir(ctx *cli.Context) string {
	if path := ctx.String(DataDirFlag.Name); path != "" {
		if ctx.Bool(TestnetFlag.Name) {
			return filepath.Join(path, "testnet")
		}
		if ctx.Bool(RinkebyFlag.Name) {
			return filepath.Join(path, "rinkeby")
		}
		return path
	}
	Fatalf("Cannot determine default data directory, please set manually (--datadir)")
	return ""
}

// setNodeKey creates a node key from set command line flags, either loading it
// from a file or as a specified hex value. If neither flags were provided, this
// method returns nil and an emphemeral key is to be generated.
func setNodeKey(ctx *cli.Context, cfg *p2p.Config) {
	var (
		hex  = ctx.String(NodeKeyHexFlag.Name)
		file = ctx.String(NodeKeyFileFlag.Name)
		key  *ecdsa.PrivateKey
		err  error
	)
	switch {
	case file != "" && hex != "":
		Fatalf("Options %q and %q are mutually exclusive", NodeKeyFileFlag.Name, NodeKeyHexFlag.Name)
	case file != "":
		if key, err = crypto.LoadECDSA(file); err != nil {
			Fatalf("Option %q: %v", NodeKeyFileFlag.Name, err)
		}
		cfg.PrivateKey = key
	case hex != "":
		if key, err = crypto.HexToECDSA(hex); err != nil {
			Fatalf("Option %q: %v", NodeKeyHexFlag.Name, err)
		}
		cfg.PrivateKey = key
	}
}

// setNodeUserIdent creates the user identifier from CLI flags.
func setNodeUserIdent(ctx *cli.Context, cfg *node.Config) {
	if identity := ctx.String(IdentityFlag.Name); len(identity) > 0 {
		cfg.UserIdent = identity
	}
}

// setBootstrapNodes creates a list of bootstrap nodes from the command line
// flags, reverting to pre-configured ones if none have been specified.
func setBootstrapNodes(ctx *cli.Context, cfg *p2p.Config) {
	urls := []string{}
	switch {
	case ctx.IsSet(BootnodesFlag.Name) || ctx.IsSet(BootnodesV4Flag.Name):
		if ctx.IsSet(BootnodesV4Flag.Name) {
			urls = strings.Split(ctx.String(BootnodesV4Flag.Name), ",")
		} else {
			urls = strings.Split(ctx.String(BootnodesFlag.Name), ",")
		}
	// case ctx.Bool(TestnetFlag.Name):
	// 	urls = params.TestnetBootnodes
	// case ctx.Bool(RinkebyFlag.Name):
	// 	urls = params.RinkebyBootnodes
	case cfg.BootstrapNodes != nil:
		return // already set, don't apply defaults.
	case !ctx.IsSet(BootnodesFlag.Name):
		urls = params.MainnetBootnodes
	case ctx.Bool(XDCTestnetFlag.Name):
		urls = params.TestnetBootnodes
	}
	cfg.BootstrapNodes = make([]*discover.Node, 0, len(urls))
	for _, url := range urls {
		node, err := discover.ParseNode(url)
		if err != nil {
			log.Error("Bootstrap URL invalid", "enode", url, "err", err)
			continue
		}
		cfg.BootstrapNodes = append(cfg.BootstrapNodes, node)
	}
}

// setBootstrapNodesV5 creates a list of bootstrap nodes from the command line
// flags, reverting to pre-configured ones if none have been specified.
func setBootstrapNodesV5(ctx *cli.Context, cfg *p2p.Config) {
	urls := params.DiscoveryV5Bootnodes
	switch {
	case ctx.IsSet(BootnodesFlag.Name) || ctx.IsSet(BootnodesV5Flag.Name):
		if ctx.IsSet(BootnodesV5Flag.Name) {
			urls = strings.Split(ctx.String(BootnodesV5Flag.Name), ",")
		} else {
			urls = strings.Split(ctx.String(BootnodesFlag.Name), ",")
		}
	case ctx.Bool(RinkebyFlag.Name):
		urls = params.RinkebyBootnodes
	case cfg.BootstrapNodesV5 != nil:
		return // already set, don't apply defaults.
	}

	cfg.BootstrapNodesV5 = make([]*discv5.Node, 0, len(urls))
	for _, url := range urls {
		node, err := discv5.ParseNode(url)
		if err != nil {
			log.Error("Bootstrap URL invalid", "enode", url, "err", err)
			continue
		}
		cfg.BootstrapNodesV5 = append(cfg.BootstrapNodesV5, node)
	}
}

// setListenAddress creates a TCP listening address string from set command
// line flags.
func setListenAddress(ctx *cli.Context, cfg *p2p.Config) {
	if ctx.IsSet(ListenPortFlag.Name) {
		cfg.ListenAddr = fmt.Sprintf(":%d", ctx.Int(ListenPortFlag.Name))
	}
}

// setNAT creates a port mapper from command line flags.
func setNAT(ctx *cli.Context, cfg *p2p.Config) {
	if ctx.IsSet(NATFlag.Name) {
		log.Info("NAT is setted", "value", ctx.String(NATFlag.Name))
		natif, err := nat.Parse(ctx.String(NATFlag.Name))
		if err != nil {
			Fatalf("Option %s: %v", NATFlag.Name, err)
		}
		cfg.NAT = natif
	}
}

// splitAndTrim splits input separated by a comma
// and trims excessive white space from the substrings.
func splitAndTrim(input string) []string {
	result := strings.Split(input, ",")
	for i, r := range result {
		result[i] = strings.TrimSpace(r)
	}
	return result
}

// setHTTP creates the HTTP RPC listener interface string from the set
// command line flags, returning empty if the HTTP endpoint is disabled.
func setHTTP(ctx *cli.Context, cfg *node.Config) {
	if ctx.Bool(RPCEnabledFlag.Name) && cfg.HTTPHost == "" {
		cfg.HTTPHost = "127.0.0.1"
		if ctx.IsSet(RPCListenAddrFlag.Name) {
			cfg.HTTPHost = ctx.String(RPCListenAddrFlag.Name)
		}
	}

	if ctx.IsSet(RPCPortFlag.Name) {
		cfg.HTTPPort = ctx.Int(RPCPortFlag.Name)
	}
	if ctx.IsSet(RPCCORSDomainFlag.Name) {
		cfg.HTTPCors = splitAndTrim(ctx.String(RPCCORSDomainFlag.Name))
	}
	if ctx.IsSet(RPCApiFlag.Name) {
		cfg.HTTPModules = splitAndTrim(ctx.String(RPCApiFlag.Name))
	}
	if ctx.IsSet(RPCVirtualHostsFlag.Name) {
		cfg.HTTPVirtualHosts = splitAndTrim(ctx.String(RPCVirtualHostsFlag.Name))
	}
}

// setWS creates the WebSocket RPC listener interface string from the set
// command line flags, returning empty if the HTTP endpoint is disabled.
func setWS(ctx *cli.Context, cfg *node.Config) {
	if ctx.Bool(WSEnabledFlag.Name) && cfg.WSHost == "" {
		cfg.WSHost = "127.0.0.1"
		if ctx.IsSet(WSListenAddrFlag.Name) {
			cfg.WSHost = ctx.String(WSListenAddrFlag.Name)
		}
	}

	if ctx.IsSet(WSPortFlag.Name) {
		cfg.WSPort = ctx.Int(WSPortFlag.Name)
	}
	if ctx.IsSet(WSAllowedOriginsFlag.Name) {
		cfg.WSOrigins = splitAndTrim(ctx.String(WSAllowedOriginsFlag.Name))
	}
	if ctx.IsSet(WSApiFlag.Name) {
		cfg.WSModules = splitAndTrim(ctx.String(WSApiFlag.Name))
	}
}

// setIPC creates an IPC path configuration from the set command line flags,
// returning an empty string if IPC was explicitly disabled, or the set path.
func setIPC(ctx *cli.Context, cfg *node.Config) {
	checkExclusive(ctx, IPCDisabledFlag, IPCPathFlag)
	switch {
	case ctx.Bool(IPCDisabledFlag.Name):
		cfg.IPCPath = ""
	case ctx.IsSet(IPCPathFlag.Name):
		cfg.IPCPath = ctx.String(IPCPathFlag.Name)
	}
}

// MakeDatabaseHandles raises out the number of allowed file handles per process
// for XDC and returns half of the allowance to assign to the database.
func MakeDatabaseHandles() int {
	limit, err := fdlimit.Current()
	if err != nil {
		Fatalf("Failed to retrieve file descriptor allowance: %v", err)
	}
	if limit < 2048 {
		if _, err := fdlimit.Raise(2048); err != nil {
			Fatalf("Failed to raise file descriptor allowance: %v", err)
		}
	}
	if limit > 2048 { // cap database file descriptors even if more is available
		limit = 2048
	}
	return limit / 2 // Leave half for networking and other stuff
}

// MakeAddress converts an account specified directly as a hex encoded string or
// a key index in the key store to an internal account representation.
func MakeAddress(ks *keystore.KeyStore, account string) (accounts.Account, error) {
	// If the specified account is a valid address, return it
	if common.IsHexAddress(account) {
		return accounts.Account{Address: common.HexToAddress(account)}, nil
	}
	// Otherwise try to interpret the account as a keystore index
	index, err := strconv.Atoi(account)
	if err != nil || index < 0 {
		return accounts.Account{}, fmt.Errorf("invalid account address or index %q", account)
	}
	log.Warn("-------------------------------------------------------------------")
	log.Warn("Referring to accounts by order in the keystore folder is dangerous!")
	log.Warn("This functionality is deprecated and will be removed in the future!")
	log.Warn("Please use explicit addresses! (can search via `XDC account list`)")
	log.Warn("-------------------------------------------------------------------")

	accs := ks.Accounts()
	if len(accs) <= index {
		return accounts.Account{}, fmt.Errorf("index %d higher than number of accounts %d", index, len(accs))
	}
	return accs[index], nil
}

// setEtherbase retrieves the etherbase either from the directly specified
// command line flags or from the keystore if CLI indexed.
func setEtherbase(ctx *cli.Context, ks *keystore.KeyStore, cfg *eth.Config) {
	if ctx.IsSet(EtherbaseFlag.Name) {
		account, err := MakeAddress(ks, ctx.String(EtherbaseFlag.Name))
		if err != nil {
			Fatalf("Option %q: %v", EtherbaseFlag.Name, err)
		}
		cfg.Etherbase = account.Address
	}
}

// MakePasswordList reads password lines from the file specified by the global --password flag.
func MakePasswordList(ctx *cli.Context) []string {
	path := ctx.String(PasswordFileFlag.Name)
	if path == "" {
		return nil
	}
	text, err := ioutil.ReadFile(path)
	if err != nil {
		Fatalf("Failed to read password file: %v", err)
	}
	lines := strings.Split(string(text), "\n")
	// Sanitise DOS line endings.
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r")
	}
	return lines
}

func SetP2PConfig(ctx *cli.Context, cfg *p2p.Config) {
	setNodeKey(ctx, cfg)
	setNAT(ctx, cfg)
	setListenAddress(ctx, cfg)
	setBootstrapNodes(ctx, cfg)
	// setBootstrapNodesV5(ctx, cfg)

	lightClient := ctx.Bool(LightModeFlag.Name) || ctx.String(SyncModeFlag.Name) == "light"
	lightServer := ctx.Int(LightServFlag.Name) != 0
	lightPeers := ctx.Int(LightPeersFlag.Name)

	if ctx.IsSet(MaxPeersFlag.Name) {
		cfg.MaxPeers = ctx.Int(MaxPeersFlag.Name)
		if lightServer && !ctx.IsSet(LightPeersFlag.Name) {
			cfg.MaxPeers += lightPeers
		}
	} else {
		if lightServer {
			cfg.MaxPeers += lightPeers
		}
		if lightClient && ctx.IsSet(LightPeersFlag.Name) && cfg.MaxPeers < lightPeers {
			cfg.MaxPeers = lightPeers
		}
	}
	if !(lightClient || lightServer) {
		lightPeers = 0
	}
	ethPeers := cfg.MaxPeers - lightPeers
	if lightClient {
		ethPeers = 0
	}
	log.Info("Maximum peer count", "ETH", ethPeers, "LES", lightPeers, "total", cfg.MaxPeers)

	if ctx.IsSet(MaxPendingPeersFlag.Name) {
		cfg.MaxPendingPeers = ctx.Int(MaxPendingPeersFlag.Name)
	}
	if ctx.IsSet(NoDiscoverFlag.Name) || lightClient {
		cfg.NoDiscovery = true
	}

	// if we're running a light client or server, force enable the v5 peer discovery
	// unless it is explicitly disabled with --nodiscover note that explicitly specifying
	// --v5disc overrides --nodiscover, in which case the later only disables v4 discovery
	forceV5Discovery := (lightClient || lightServer) && !ctx.Bool(NoDiscoverFlag.Name)
	if ctx.IsSet(DiscoveryV5Flag.Name) {
		cfg.DiscoveryV5 = ctx.Bool(DiscoveryV5Flag.Name)
	} else if forceV5Discovery {
		cfg.DiscoveryV5 = true
	}

	if netrestrict := ctx.String(NetrestrictFlag.Name); netrestrict != "" {
		list, err := netutil.ParseNetlist(netrestrict)
		if err != nil {
			Fatalf("Option %q: %v", NetrestrictFlag.Name, err)
		}
		cfg.NetRestrict = list
	}

	if ctx.Bool(DeveloperFlag.Name) {
		// --dev mode can't use p2p networking.
		cfg.MaxPeers = 0
		cfg.ListenAddr = ":0"
		cfg.NoDiscovery = true
		cfg.DiscoveryV5 = false
	}
}

// SetNodeConfig applies node-related command line flags to the config.
func SetNodeConfig(ctx *cli.Context, cfg *node.Config) {
	SetP2PConfig(ctx, &cfg.P2P)
	setIPC(ctx, cfg)
	setHTTP(ctx, cfg)
	setWS(ctx, cfg)
	setNodeUserIdent(ctx, cfg)

	switch {
	case ctx.IsSet(DataDirFlag.Name):
		cfg.DataDir = ctx.String(DataDirFlag.Name)
	case ctx.Bool(DeveloperFlag.Name):
		cfg.DataDir = "" // unless explicitly requested, use memory databases
	case ctx.Bool(TestnetFlag.Name):
		cfg.DataDir = filepath.Join(node.DefaultDataDir(), "testnet")
	case ctx.Bool(RinkebyFlag.Name):
		cfg.DataDir = filepath.Join(node.DefaultDataDir(), "rinkeby")
	}

	if ctx.IsSet(KeyStoreDirFlag.Name) {
		cfg.KeyStoreDir = ctx.String(KeyStoreDirFlag.Name)
	}
	if ctx.IsSet(LightKDFFlag.Name) {
		cfg.UseLightweightKDF = ctx.Bool(LightKDFFlag.Name)
	}
	if ctx.IsSet(NoUSBFlag.Name) {
		cfg.NoUSB = ctx.Bool(NoUSBFlag.Name)
	}
	if ctx.IsSet(AnnounceTxsFlag.Name) {
		cfg.AnnounceTxs = ctx.Bool(AnnounceTxsFlag.Name)
	}
}

func setGPO(ctx *cli.Context, cfg *gasprice.Config) {
	if ctx.IsSet(GpoBlocksFlag.Name) {
		cfg.Blocks = ctx.Int(GpoBlocksFlag.Name)
	}
	if ctx.IsSet(GpoPercentileFlag.Name) {
		cfg.Percentile = ctx.Int(GpoPercentileFlag.Name)
	}
}

func setTxPool(ctx *cli.Context, cfg *core.TxPoolConfig) {
	if ctx.IsSet(TxPoolNoLocalsFlag.Name) {
		cfg.NoLocals = ctx.Bool(TxPoolNoLocalsFlag.Name)
	}
	if ctx.IsSet(TxPoolJournalFlag.Name) {
		cfg.Journal = ctx.String(TxPoolJournalFlag.Name)
	}
	if ctx.IsSet(TxPoolRejournalFlag.Name) {
		cfg.Rejournal = ctx.Duration(TxPoolRejournalFlag.Name)
	}
	if ctx.IsSet(TxPoolPriceLimitFlag.Name) {
		cfg.PriceLimit = ctx.Uint64(TxPoolPriceLimitFlag.Name)
	}
	if ctx.IsSet(TxPoolPriceBumpFlag.Name) {
		cfg.PriceBump = ctx.Uint64(TxPoolPriceBumpFlag.Name)
	}
	if ctx.IsSet(TxPoolAccountSlotsFlag.Name) {
		cfg.AccountSlots = ctx.Uint64(TxPoolAccountSlotsFlag.Name)
	}
	if ctx.IsSet(TxPoolGlobalSlotsFlag.Name) {
		cfg.GlobalSlots = ctx.Uint64(TxPoolGlobalSlotsFlag.Name)
	}
	if ctx.IsSet(TxPoolAccountQueueFlag.Name) {
		cfg.AccountQueue = ctx.Uint64(TxPoolAccountQueueFlag.Name)
	}
	if ctx.IsSet(TxPoolGlobalQueueFlag.Name) {
		cfg.GlobalQueue = ctx.Uint64(TxPoolGlobalQueueFlag.Name)
	}
	if ctx.IsSet(TxPoolLifetimeFlag.Name) {
		cfg.Lifetime = ctx.Duration(TxPoolLifetimeFlag.Name)
	}
}

// checkExclusive verifies that only a single isntance of the provided flags was
// set by the user. Each flag might optionally be followed by a string type to
// specialize it further.
func checkExclusive(ctx *cli.Context, args ...interface{}) {
	set := make([]string, 0, 1)
	for i := 0; i < len(args); i++ {
		// Make sure the next argument is a flag and skip if not set
		flag, ok := args[i].(cli.Flag)
		if !ok {
			panic(fmt.Sprintf("invalid argument, not cli.Flag type: %T", args[i]))
		}
		// Check if next arg extends current and expand its name if so
		name := flag.String()

		if i+1 < len(args) {
			switch option := args[i+1].(type) {
			case string:
				// Extended flag, expand the name and shift the arguments
				if ctx.String(flag.String()) == option {
					name += "=" + option
				}
				i++

			case cli.Flag:
			default:
				panic(fmt.Sprintf("invalid argument, not cli.Flag or string extension: %T", args[i+1]))
			}
		}
		// Mark the flag if it's set
		if ctx.IsSet(flag.String()) {
			set = append(set, "--"+name)
		}
	}
	if len(set) > 1 {
		Fatalf("Flags %v can't be used at the same time", strings.Join(set, ", "))
	}
}

// SetShhConfig applies shh-related command line flags to the config.
func SetShhConfig(ctx *cli.Context, stack *node.Node, cfg *whisper.Config) {
	if ctx.IsSet(WhisperMaxMessageSizeFlag.Name) {
		cfg.MaxMessageSize = uint32(ctx.Uint(WhisperMaxMessageSizeFlag.Name))
	}
	if ctx.IsSet(WhisperMinPOWFlag.Name) {
		cfg.MinimumAcceptedPOW = ctx.Float64(WhisperMinPOWFlag.Name)
	}
}

func SetXDCXConfig(ctx *cli.Context, cfg *XDCx.Config, XDCDataDir string) {
	if ctx.IsSet(XDCXDataDirFlag.Name) {
		cfg.DataDir = ctx.String(XDCXDataDirFlag.Name)
	} else {
		// default XDCx datadir: DATADIR/XDCx
		defaultXDCXDataDir := filepath.Join(XDCDataDir, "XDCx")

		filesInXDCXDefaultDir, _ := WalkMatch(defaultXDCXDataDir, "*.ldb")
		filesInNodeDefaultDir, _ := WalkMatch(node.DefaultDataDir(), "*.ldb")
		if len(filesInXDCXDefaultDir) == 0 && len(filesInNodeDefaultDir) > 0 {
			cfg.DataDir = node.DefaultDataDir()
		} else {
			cfg.DataDir = defaultXDCXDataDir
		}
	}
	log.Info("XDCX datadir", "path", cfg.DataDir)
	if ctx.IsSet(XDCXDBEngineFlag.Name) {
		cfg.DBEngine = ctx.String(XDCXDBEngineFlag.Name)
	} else {
		cfg.DBEngine = XDCXDBEngineFlag.Value
	}
	if ctx.IsSet(XDCXDBNameFlag.Name) {
		cfg.DBName = ctx.String(XDCXDBNameFlag.Name)
	} else {
		cfg.DBName = XDCXDBNameFlag.Value
	}
	if ctx.IsSet(XDCXDBConnectionUrlFlag.Name) {
		cfg.ConnectionUrl = ctx.String(XDCXDBConnectionUrlFlag.Name)
	} else {
		cfg.ConnectionUrl = XDCXDBConnectionUrlFlag.Value
	}
	if ctx.IsSet(XDCXDBReplicaSetNameFlag.Name) {
		cfg.ReplicaSetName = ctx.String(XDCXDBReplicaSetNameFlag.Name)
	}
}

// SetEthConfig applies eth-related command line flags to the config.
func SetEthConfig(ctx *cli.Context, stack *node.Node, cfg *eth.Config) {
	// Avoid conflicting network flags
	checkExclusive(ctx, DeveloperFlag, TestnetFlag, RinkebyFlag)
	checkExclusive(ctx, FastSyncFlag, LightModeFlag, SyncModeFlag)
	checkExclusive(ctx, LightServFlag, LightModeFlag)
	checkExclusive(ctx, LightServFlag, SyncModeFlag, "light")

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	setEtherbase(ctx, ks, cfg)
	setGPO(ctx, &cfg.GPO)
	setTxPool(ctx, &cfg.TxPool)

	switch {
	case ctx.IsSet(SyncModeFlag.Name):
		cfg.SyncMode = *GlobalTextMarshaler(ctx, SyncModeFlag.Name).(*downloader.SyncMode)
	case ctx.Bool(FastSyncFlag.Name):
		cfg.SyncMode = downloader.FastSync
	case ctx.Bool(LightModeFlag.Name):
		cfg.SyncMode = downloader.LightSync
	}
	if ctx.IsSet(LightServFlag.Name) {
		cfg.LightServ = ctx.Int(LightServFlag.Name)
	}
	if ctx.IsSet(LightPeersFlag.Name) {
		cfg.LightPeers = ctx.Int(LightPeersFlag.Name)
	}
	if ctx.IsSet(NetworkIdFlag.Name) {
		cfg.NetworkId = ctx.Uint64(NetworkIdFlag.Name)
	}

	if ctx.IsSet(CacheFlag.Name) || ctx.IsSet(CacheDatabaseFlag.Name) {
		cfg.DatabaseCache = ctx.Int(CacheFlag.Name) * ctx.Int(CacheDatabaseFlag.Name) / 100
	}
	cfg.DatabaseHandles = MakeDatabaseHandles()

	if gcmode := ctx.String(GCModeFlag.Name); gcmode != "full" && gcmode != "archive" {
		Fatalf("--%s must be either 'full' or 'archive'", GCModeFlag.Name)
	}
	cfg.NoPruning = ctx.String(GCModeFlag.Name) == "archive"

	if ctx.IsSet(CacheFlag.Name) || ctx.IsSet(CacheGCFlag.Name) {
		cfg.TrieCache = ctx.Int(CacheFlag.Name) * ctx.Int(CacheGCFlag.Name) / 100
	}
	if ctx.IsSet(StakerThreadsFlag.Name) {
		cfg.MinerThreads = ctx.Int(StakerThreadsFlag.Name)
	}
	if ctx.IsSet(DocRootFlag.Name) {
		cfg.DocRoot = ctx.String(DocRootFlag.Name)
	}
	if ctx.IsSet(ExtraDataFlag.Name) {
		cfg.ExtraData = []byte(ctx.String(ExtraDataFlag.Name))
	}
	if ctx.IsSet(GasPriceFlag.Name) {
		cfg.GasPrice = GlobalBig(ctx, GasPriceFlag.Name)
	}
	if ctx.IsSet(VMEnableDebugFlag.Name) {
		// TODO(fjl): force-enable this in --dev mode
		cfg.EnablePreimageRecording = ctx.Bool(VMEnableDebugFlag.Name)
	}
	if ctx.IsSet(StoreRewardFlag.Name) {
		common.StoreRewardFolder = filepath.Join(stack.DataDir(), "XDC", "rewards")
		if _, err := os.Stat(common.StoreRewardFolder); os.IsNotExist(err) {
			os.Mkdir(common.StoreRewardFolder, os.ModePerm)
		}
	}
	// Override any default configs for hard coded networks.
	switch {
	case ctx.Bool(TestnetFlag.Name):
		if !ctx.IsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = 3
		}
		cfg.Genesis = core.DefaultTestnetGenesisBlock()
	case ctx.Bool(RinkebyFlag.Name):
		if !ctx.IsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = 4
		}
		cfg.Genesis = core.DefaultRinkebyGenesisBlock()
	case ctx.Bool(DeveloperFlag.Name):
		// Create new developer account or reuse existing one
		var (
			developer accounts.Account
			err       error
		)
		if accs := ks.Accounts(); len(accs) > 0 {
			developer = ks.Accounts()[0]
		} else {
			developer, err = ks.NewAccount("")
			if err != nil {
				Fatalf("Failed to create developer account: %v", err)
			}
		}
		if err := ks.Unlock(developer, ""); err != nil {
			Fatalf("Failed to unlock developer account: %v", err)
		}
		log.Info("Using developer account", "address", developer.Address)

		cfg.Genesis = core.DeveloperGenesisBlock(uint64(ctx.Int(DeveloperPeriodFlag.Name)), developer.Address)
		if !ctx.IsSet(GasPriceFlag.Name) {
			cfg.GasPrice = big.NewInt(1)
		}
	}
	// TODO(fjl): move trie cache generations into config
}

// SetupNetwork configures the system for either the main net or some test network.
func SetupNetwork(ctx *cli.Context) {
	// TODO(fjl): move target gas limit into config
	params.TargetGasLimit = ctx.Uint64(TargetGasLimitFlag.Name)
}

// MakeChainDatabase open an LevelDB using the flags passed to the client and will hard crash if it fails.
func MakeChainDatabase(ctx *cli.Context, stack *node.Node) ethdb.Database {
	var (
		cache   = ctx.Int(CacheFlag.Name) * ctx.Int(CacheDatabaseFlag.Name) / 100
		handles = MakeDatabaseHandles()
	)
	name := "chaindata"
	if ctx.Bool(LightModeFlag.Name) {
		name = "lightchaindata"
	}
	chainDb, err := stack.OpenDatabase(name, cache, handles, "")
	if err != nil {
		Fatalf("Could not open database: %v", err)
	}
	return chainDb
}

func MakeGenesis(ctx *cli.Context) *core.Genesis {
	var genesis *core.Genesis
	switch {
	case ctx.Bool(TestnetFlag.Name):
		genesis = core.DefaultTestnetGenesisBlock()
	case ctx.Bool(RinkebyFlag.Name):
		genesis = core.DefaultRinkebyGenesisBlock()
	case ctx.Bool(DeveloperFlag.Name):
		Fatalf("Developer chains are ephemeral")
	}
	return genesis
}

// MakeChain creates a chain manager from set command line flags.
func MakeChain(ctx *cli.Context, stack *node.Node, readonly bool) (chain *core.BlockChain, chainDb ethdb.Database) {
	var (
		gspec   = MakeGenesis(ctx)
		chainDb = MakeChainDatabase(ctx, stack, readonly)
	)

	config, err := core.LoadChainConfig(chainDb, gspec)
	if err != nil {
		Fatalf("%v", err)
	}
	engine, err := ethconfig.CreateConsensusEngine(config, chainDb)
	if err != nil {
		Fatalf("%v", err)
	}
	if gcmode := ctx.String(GCModeFlag.Name); gcmode != "full" && gcmode != "archive" {
		Fatalf("--%s must be either 'full' or 'archive'", GCModeFlag.Name)
	}
	if config.XDPoS != nil {
		engine = XDPoS.New(config, chainDb)
	}
	/* TODO! Useless init
	else {
		engine = ethash.NewFaker()
		if !ctx.Bool(FakePoWFlag.Name) {
			engine = ethash.New(ethash.Config{
				CacheDir:       stack.ResolvePath(eth.DefaultConfig.Ethash.CacheDir),
				CachesInMem:    eth.DefaultConfig.Ethash.CachesInMem,
				CachesOnDisk:   eth.DefaultConfig.Ethash.CachesOnDisk,
				DatasetDir:     stack.ResolvePath(eth.DefaultConfig.Ethash.DatasetDir),
				DatasetsInMem:  eth.DefaultConfig.Ethash.DatasetsInMem,
				DatasetsOnDisk: eth.DefaultConfig.Ethash.DatasetsOnDisk,
			})
		}
		Fatalf("Only support XDPoS consensus")
	} */
	if gcmode := ctx.String(GCModeFlag.Name); gcmode != "full" && gcmode != "archive" {
		Fatalf("--%s must be either 'full' or 'archive'", GCModeFlag.Name)
	}
	cache := &core.CacheConfig{
		TrieCleanLimit:      ethconfig.Defaults.TrieCleanCache,
		TrieCleanNoPrefetch: ctx.Bool(CacheNoPrefetchFlag.Name),
		TrieDirtyLimit:      ethconfig.Defaults.TrieDirtyCache,
		TrieDirtyDisabled:   ctx.String(GCModeFlag.Name) == "archive",
		TrieTimeLimit:       ethconfig.Defaults.TrieTimeout,
		SnapshotLimit:       ethconfig.Defaults.SnapshotCache,
		Preimages:           ctx.Bool(CachePreimagesFlag.Name),
	}
	if cache.TrieDirtyDisabled && !cache.Preimages {
		cache.Preimages = true
		log.Info("Enabling recording of key preimages since archive mode is used")
	}
	if !ctx.Bool(SnapshotFlag.Name) {
		cache.SnapshotLimit = 0 // Disabled
	}
	// If we're in readonly, do not bother generating snapshot data.
	if readonly {
		cache.SnapshotNoBuild = true
	}

	if ctx.IsSet(CacheFlag.Name) || ctx.IsSet(CacheTrieFlag.Name) {
		cache.TrieCleanLimit = ctx.Int(CacheFlag.Name) * ctx.Int(CacheTrieFlag.Name) / 100
	}
	if ctx.IsSet(CacheFlag.Name) || ctx.IsSet(CacheGCFlag.Name) {
		cache.TrieDirtyLimit = ctx.Int(CacheFlag.Name) * ctx.Int(CacheGCFlag.Name) / 100
	}
	vmcfg := vm.Config{EnablePreimageRecording: ctx.Bool(VMEnableDebugFlag.Name)}

	// Disable transaction indexing/unindexing by default.
	chain, err := core.NewBlockChain(chainDb, cache, gspec, nil, engine, vmcfg, nil, nil)
	if err != nil {
		Fatalf("Can't create BlockChain: %v", err)
	}
	return chain, chainDb
}

// MakeConsolePreloads retrieves the absolute paths for the console JavaScript
// scripts to preload before starting.
func MakeConsolePreloads(ctx *cli.Context) []string {
	// Skip preloading if there's nothing to preload
	if ctx.String(PreloadJSFlag.Name) == "" {
		return nil
	}
	// Otherwise resolve absolute paths and return them
	preloads := []string{}

	assets := ctx.String(JSpathFlag.Name)
	for _, file := range strings.Split(ctx.String(PreloadJSFlag.Name), ",") {
		preloads = append(preloads, common.AbsolutePath(assets, strings.TrimSpace(file)))
	}
	return preloads
}

// MigrateFlags sets the global flag from a local flag when it's set.
// This is a temporary function used for migrating old command/flags to the
// new format.
//
// e.g. XDC account new --keystore /tmp/mykeystore --lightkdf
//
// is equivalent after calling this method with:
//
// XDC --keystore /tmp/mykeystore --lightkdf account new
//
// This allows the use of the existing configuration functionality.
// When all flags are migrated this function can be removed and the existing
// configuration functionality must be changed that is uses local flags
func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			if ctx.IsSet(name) {
				ctx.Set(name, ctx.String(name))
			}
		}
		return action(ctx)
	}
}

// find all filenames match the given pattern in the given root directory
func WalkMatch(root, pattern string) ([]string, error) {
	matches := []string{}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func SetupMetrics(ctx *cli.Context) {
	if metrics.Enabled {
		log.Info("Enabling metrics collection")

		if ctx.IsSet(MetricsHTTPFlag.Name) {
			address := fmt.Sprintf("%s:%d", ctx.String(MetricsHTTPFlag.Name), ctx.Int(MetricsPortFlag.Name))
			log.Info("Enabling stand-alone metrics HTTP endpoint", "address", address)
			exp.Setup(address)
		}
	}
}
