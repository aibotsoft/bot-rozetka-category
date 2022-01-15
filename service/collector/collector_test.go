package collector

import (
	"context"
	"github.com/aibotsoft/bot-rozetka-category/pkg/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

var c *Collector

func TestMain(m *testing.M) {
	logger, _ := zap.NewDevelopment()
	cfg := new(config.Config)
	cfg.ApiDebug = true
	cfg.RootCategoryID = 83850
	logger.Info("test_service", zap.Any("cfg", cfg))

	ctx, cancel := context.WithCancel(context.Background())

	c = NewCollector(logger, cfg, ctx, nil)
	m.Run()
	cancel()

}

func TestCollector_GetProductList(t *testing.T) {
	got, err := c.getProductList(101808)
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	t.Log(got, len(got))
}

func TestCollector_GetProductDetails(t *testing.T) {
	_, err := c.getProductDetails([]int64{294210068, 331746307, 331866181})
	assert.NoError(t, err)
}

func TestCollector_Run(t *testing.T) {
	err := c.Run()
	assert.NoError(t, err)
}

func TestCollector_getProduct(t *testing.T) {
	_, err := c.getProductPage(294210068)
	assert.NoError(t, err)
}

func TestCollector_getCategories(t *testing.T) {
	got, err := c.getCategories()
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	t.Log(got)
}
