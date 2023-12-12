package app

import (
	"context"
	"github.com/elijahelrod/vespene/config"
	"github.com/elijahelrod/vespene/internal/websocket"
	"github.com/elijahelrod/vespene/pkg/exchange/coinbase"
	"github.com/elijahelrod/vespene/pkg/logger/zap"
	"os"
)

func Run(ctx context.Context, cfg *config.Config) {

	// Setup Logger Provider (Zap)
	loggerProvider := zap.NewLogger(cfg.Logger.Level, cfg.Logger.DisableCaller, cfg.Logger.DisableStacktrace)
	loggerProvider.InitLogger()

	//// Setup DB Client/Provider (MySQL)
	//dbProvider, err := mysql.NewClient(cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Base)
	//if err != nil {
	//	loggerProvider.Fatal(err)
	//}
	//defer dbProvider.Close()

	loggerProvider.Info("Initialized DB Client")

	exchangeProvider, err := coinbase.NewClient(cfg.Exchange.Url)
	if err != nil {
		loggerProvider.Fatal(err)
	}

	defer exchangeProvider.Close()
	loggerProvider.Info("Initialized Exchange Client")

	socketProvider, err := websocket.NewClient(exchangeProvider, loggerProvider, cfg.Exchange)
	if err != nil {
		loggerProvider.Fatal(err)
	}

	// run
	go func() {
		loggerProvider.Info("socket starting...")
		if err = socketProvider.Run(ctx); err != nil {
			loggerProvider.Fatal(err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()

	loggerProvider.Info("socket stopping...")
}
