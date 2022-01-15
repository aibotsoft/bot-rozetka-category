package store

import (
	"github.com/aibotsoft/bot-rozetka-category/pkg/config"
	"go.uber.org/zap"
	"testing"
)

var s *Store

func TestMain(m *testing.M) {
	log, _ := zap.NewDevelopment()
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	log.Info("hello", zap.Any("cfg", cfg))
	s, err = NewStore(log, cfg)
	if err != nil {
		log.Panic("new_store_error", zap.Error(err))
	}

	//ctx, cancel := context.WithCancel(context.Background())
	//c = collector.NewCollector(log, cfg, ctx, sto)

	m.Run()
	//cancel()

}

func TestStore_GetCategories(t *testing.T) {
	//err := s.DeleteCategoriesNotInList([]int64{1, 2})
	//assert.NoError(t, err)
}
