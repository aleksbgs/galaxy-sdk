//go:build cometbft

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	abciserver "github.com/cometbft/cometbft/abci/server"
	cmtlog "github.com/cometbft/cometbft/libs/log"

	"github.com/aleksbgs/galaxy-sdk/abci"
	"github.com/aleksbgs/galaxy-sdk/app"
	"github.com/aleksbgs/galaxy-sdk/codec"
	"github.com/aleksbgs/galaxy-sdk/core"
	"github.com/aleksbgs/galaxy-sdk/modules/bank"
	"github.com/aleksbgs/galaxy-sdk/modules/staking"
)

func main() {
	socketAddr := flag.String("socket", "unix:///tmp/newchain.sock", "ABCI socket address")
	chainID := flag.String("chain-id", "localnet", "chain ID")
	flag.Parse()

	app := app.New(*chainID)

	msgReg := core.NewMsgRegistry()
	msgReg.Register("bank/send", func() core.Msg { return bank.MsgSend{} })
	msgReg.Register("staking/delegate", func() core.Msg { return staking.MsgDelegate{} })

	decoder := codec.NewJSONTxDecoder(msgReg)
	abciApp, err := abci.NewCometBFTApp(app, decoder)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to init ABCI app: %v\n", err)
		os.Exit(1)
	}

	logger := cmtlog.NewTMLogger(cmtlog.NewSyncWriter(os.Stdout))
	server := abciserver.NewSocketServer(*socketAddr, abciApp)
	server.SetLogger(logger)

	if err := server.Start(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to start ABCI server: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		_ = server.Stop()
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh
}
