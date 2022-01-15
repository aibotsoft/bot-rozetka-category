package collector

import (
	"context"
	"github.com/aibotsoft/bot-rozetka-category/api"
	"github.com/aibotsoft/bot-rozetka-category/pkg/store"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"strconv"
)

func (c *Collector) getCategories() (categories []store.Category, err error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.cfg.RequestTimeout)
	defer cancel()
	resp, _, err := c.client.DefaultApi.GetChildren(ctx).Lang("ru").CategoryId(c.cfg.RootCategoryID).Execute()
	if err != nil {
		return nil, err
	}

	children := resp.Data.GetChildren()
	for _, child := range children {
		if child.GetCountChildren() > 0 {
			ctx, cancel := context.WithTimeout(c.ctx, c.cfg.RequestTimeout)
			resp, _, err := c.client.DefaultApi.GetChildren(ctx).Lang("ru").CategoryId(child.GetId()).Execute()
			cancel()
			if err != nil {
				return nil, err
			}
			//c.log.Info("", zap.Any("resp", resp))
			children = append(children, resp.Data.GetChildren()...)
		}
	}
	err = copier.Copy(&categories, &children)
	return categories, err
}

func (c *Collector) getProductPage(productID int64) (string, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.cfg.RequestTimeout)
	defer cancel()
	resp, _, err := c.client.DefaultApi.ProductPage(ctx, productID).Execute()
	return resp, err
}
func (c *Collector) getProductDetails(ids []int64) ([]api.GetProductDetailsResponseData, error) {
	idsListStr := SliceInt64ToStr(ids, ",")
	ctx, cancel := context.WithTimeout(c.ctx, c.cfg.RequestTimeout)
	defer cancel()
	resp, _, err := c.client.DefaultApi.GetProductDetails(ctx).ProductIds(idsListStr).Lang("ru").Execute()
	return resp.GetData(), err
}
func (c *Collector) getProductList(CategoryId int64) ([]int64, error) {
	var ids []int64
	for i := 1; i < maxNumPages; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), c.cfg.RequestTimeout)
		resp, _, err := c.client.DefaultApi.GetProductsInCategory(ctx).CategoryId(strconv.FormatInt(CategoryId, 10)).Page(strconv.Itoa(i)).Execute()
		cancel()
		if err != nil {
			c.log.Warn("GetProductsInCategory_error", zap.Error(err))
			continue
		}
		data := resp.GetData()
		//c.log.Info("resp", zap.Any("resp", data))
		ids = append(ids, data.GetIds()...)
		if data.GetShowNext() == 0 {
			//c.log.Debug("no_more_pages_in_category", zap.Int64("CategoryId", CategoryId), zap.Int("total", len(ids)))
			break
		}
	}
	return ids, nil
}
