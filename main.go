package main

import (
	"context"
	"fmt"
	"github.com/aibotsoft/bot-rozetka-category/pkg/config"
	"github.com/aibotsoft/bot-rozetka-category/pkg/logger"
	"github.com/aibotsoft/bot-rozetka-category/pkg/signals"
	"github.com/aibotsoft/bot-rozetka-category/pkg/store"
	"github.com/aibotsoft/bot-rozetka-category/service/bot"
	"github.com/aibotsoft/bot-rozetka-category/service/collector"
	"go.uber.org/zap"
	"net/http"
)

const Version = "0.1.7"

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	cfg.Version = Version
	log, err := logger.NewLogger(cfg.LogLever, cfg.LogEncoding)
	if err != nil {
		panic(err)
	}
	log.Info("start_service", zap.Any("cfg", cfg))

	sto, err := store.NewStore(log, cfg)
	if err != nil {
		log.Panic("new_store_error", zap.Error(err))
	}
	err = sto.Migrate()
	if err != nil {
		log.Panic("store_migrate_error", zap.Error(err))
	}
	ctx, cancel := context.WithCancel(context.Background())

	tg, err := bot.NewBot(log, cfg, ctx, sto)
	if err != nil {
		log.Panic("new_telegram_bot_error", zap.Error(err))
	}

	c := collector.NewCollector(log, cfg, ctx, sto, tg)

	go func() {
		tg.Run()
	}()

	errCh := make(chan error)

	go func() {
		log.Info("begin_main_service_loop")
		errCh <- c.Run()
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	go func() {
		log.Info("listen_on_port", zap.String("port", cfg.Port), zap.String("url", fmt.Sprintf("http://localhost:%s", cfg.Port)))
		errCh <- http.ListenAndServe(fmt.Sprintf("%s:%s", "", cfg.Port), nil)
	}()

	defer func() {
		log.Info("closing_services...")
		cancel()
		err = sto.Close()
		if err != nil {
			log.Error("close_db_error", zap.Error(err))
		}
		_ = log.Sync()
	}()
	stopCh := signals.SetupSignalHandler()
	select {
	case err := <-errCh:
		log.Error("stop_service_by_error", zap.Error(err))
	case sig := <-stopCh:
		log.Info("stop_service_by_os_signal", zap.String("signal", sig.String()))
	}
}
