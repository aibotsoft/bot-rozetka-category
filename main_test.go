package main_test

import (
	"context"
	"github.com/aibotsoft/bot-rozetka-category/pkg/config"
	"github.com/aibotsoft/bot-rozetka-category/pkg/store"
	"github.com/aibotsoft/bot-rozetka-category/service/bot"
	"github.com/aibotsoft/bot-rozetka-category/service/collector"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

var c *collector.Collector

func TestMain(m *testing.M) {
	log, _ := zap.NewDevelopment()
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	log.Info("hello", zap.Any("cfg", cfg))
	sto, err := store.NewStore(log, cfg)
	if err != nil {
		log.Panic("new_store_error", zap.Error(err))
	}
	ctx, cancel := context.WithCancel(context.Background())
	tg, err := bot.NewBot(log, cfg, ctx, sto)
	if err != nil {
		log.Panic("new_telegram_bot_error", zap.Error(err))
	}
	c = collector.NewCollector(log, cfg, ctx, sto, tg)

	m.Run()
	cancel()

}

//func TestCollector_GetProductListAndDetails(t *testing.T) {
//	got, err := c.getProductList(101808)
//	assert.NoError(t, err)
//	assert.NotEmpty(t, got)
//	t.Log(got, len(got))
//	err = c.CollectProducts(got)
//	assert.NoError(t, err)
//}
//func TestCollector_GetProductList(t *testing.T) {
//	got, err := c.Categories()
//	assert.NoError(t, err)
//	assert.NotEmpty(t, got)
//}

//func TestCollector_Full(t *testing.T) {
//	err := c.CollectSaleProducts()
//	assert.NoError(t, err)
//}
//func TestCollector_CollectOriginID(t *testing.T) {
//	err := c.CollectOriginID()
//	assert.NoError(t, err)
//}

func TestCollector_CollectOriginProducts(t *testing.T) {
	err := c.Notify()
	assert.NoError(t, err)
}
