package collector

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aibotsoft/bot-rozetka-category/api"
	"github.com/aibotsoft/bot-rozetka-category/pkg/config"
	"github.com/aibotsoft/bot-rozetka-category/pkg/store"
	"github.com/aibotsoft/bot-rozetka-category/service/bot"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"html/template"
	"regexp"
	"strconv"
	"time"
)

const maxNumPages = 10
const categoryPeriod = 10 * time.Minute
const CollectOriginProductsPeriod = 3 * time.Minute
const collectPagePeriod = 10 * time.Second
const maxProductChunkSize = 1200
const NoOriginID = 0
const tmpl = `
<a href="{{.ImageMain}}"> </a>
<a href="{{.Href}}">{{.Title}}</a>
Price: {{.Price}}UAH (New: <del><a href="{{.OriginHref}}">{{.OriginPrice}}</a></del>) {{.Discount}}% {{if .PostOldPrice}}Old: {{.OldPrice}}{{end}}
Rating: {{.CommentsMark}} ({{.CommentsAmount}})
{{.Reason}}
`

type Collector struct {
	log          *zap.Logger
	cfg          *config.Config
	ctx          context.Context
	store        *store.Store
	client       *api.APIClient
	users        []store.User
	tg           *bot.Bot
	msgTmpl      *template.Template
	newProductCh chan int
}

func NewCollector(log *zap.Logger, cfg *config.Config, ctx context.Context, store *store.Store, tg *bot.Bot) *Collector {
	apiConfig := api.NewConfiguration()
	client := api.NewAPIClient(apiConfig)
	msgTmpl, err := template.New("msg").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	newProductCh := make(chan int, 10000)

	return &Collector{
		log:          log,
		cfg:          cfg,
		ctx:          ctx,
		store:        store,
		client:       client,
		tg:           tg,
		msgTmpl:      msgTmpl,
		newProductCh: newProductCh,
	}
}

var re = regexp.MustCompile(`canonical.*/p(\d*)/`)

func (c *Collector) CollectCategories() error {
	start := time.Now()
	categories, err := c.getCategories()
	if err != nil {
		return err
	}
	err = c.store.SaveCategories(&categories)
	if err != nil {
		return err
	}
	var ids []int64
	for i := 0; i < len(categories); i++ {
		ids = append(ids, categories[i].ID)
	}
	err = c.store.DeleteCategoriesNotInList(ids)
	if err != nil {
		return err
	}
	c.log.Info("CollectCategories", zap.Int("len", len(categories)), zap.Duration("elapsed", time.Since(start)))
	return nil
}

func (c *Collector) CollectOriginProducts() error {
	start := time.Now()
	ids, err := c.store.SelectOriginIDList(maxProductChunkSize)
	if err != nil {
		return err
	}
	if len(ids) == 0 {
		return nil
	}
	//c.log.Debug("CollectOriginProducts_chunk", zap.Int("len_chunk", len(chunk)), zap.Any("chunk", chunk))
	details, err := c.getProductDetails(ids)
	if err != nil {
		c.log.Warn("CollectOriginProducts_error", zap.Error(err), zap.Int("chunk_size", len(ids)))
		return err
	}
	//c.log.Debug("CollectOriginProducts_details", zap.Any("len_details", len(details)))
	var op []store.OriginProduct
	err = copier.Copy(&op, &details)
	if err != nil {
		return err
	}
	err = c.store.SaveOriginProducts(&op)
	if err != nil {
		c.log.Warn("SaveOriginProducts_error", zap.Error(err), zap.Int("chunk_size", len(ids)))
		return err
	}
	c.log.Info("CollectOriginProducts",
		zap.Int("count", len(ids)),
		zap.Int64s("list", ids),
		zap.Duration("elapsed", time.Since(start)),
	)
	return nil
}
func (c *Collector) CollectSaleProducts() error {
	start := time.Now()
	categories, err := c.store.GetCategories()
	if err != nil {
		return err
	}
	var multiList []int64
	for _, category := range categories {
		if category.GoodsCount == 0 {
			continue
		}
		list, err := c.getProductList(category.ID)
		if err != nil {
			return err
		}
		//c.log.Info("cat", zap.Any("list", list), zap.String("title", category.Title))
		newProductList := findNewProducts(list, category.SaleProductMap())
		if len(newProductList) > 0 {
			c.log.Debug("new_products_in_category",
				zap.String("category_title", category.Title),
				//zap.Any("category.SaleProductMap()", category.SaleProductMap()),
				zap.Int("count", len(newProductList)))
		}

		multiList = append(multiList, newProductList...)
	}
	chunks := chunkSlice(multiList, maxProductChunkSize)
	for _, chunk := range chunks {
		//c.log.Debug("CollectProducts_chunk", zap.Int("len_chunk", len(chunk)), zap.Any("chunk", chunk))
		details, err := c.getProductDetails(chunk)
		if err != nil {
			c.log.Warn("GetProductDetails_error", zap.Error(err), zap.Int("chunk_size", len(chunk)))
			return err
		}
		//c.log.Debug("CollectProducts_details", zap.Any("len_details", len(details)))

		var sp []store.SaleProduct
		err = copier.Copy(&sp, &details)
		if err != nil {
			return err
		}
		//s.log.Info("sp", zap.Any("products", sp))
		err = c.store.SaveSaleProducts(&sp)
		if err != nil {
			c.log.Warn("SaveSaleProducts_error", zap.Error(err), zap.Int("chunk_size", len(chunk)))
			return err
		}
	}
	for i := 0; i < len(multiList); i++ {
		c.newProductCh <- i
	}
	c.log.Info("CollectSaleProducts",
		zap.Int("count", len(multiList)),
		zap.Int64s("list", multiList),
		zap.Duration("elapsed", time.Since(start)),
	)
	return nil
}
func (c *Collector) CollectOriginID() error {
	start := time.Now()
	product, err := c.store.SelectProductWithoutOriginID()
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		c.log.Debug("no_new_record")
		return nil
	default:
		return fmt.Errorf("SelectProductWithoutOriginID_error: %w", err)
	}
	if product.ID == 0 {
		return nil
	}
	//c.log.Info("", zap.Any("findString", product))
	page, err := c.getProductPage(product.ID)
	if err != nil {
		return fmt.Errorf("getProductPage_error: %w", err)
	}
	findSub := re.FindStringSubmatch(page)
	if len(findSub) < 2 {
		c.log.Info("not_found_origin_id_in_page", zap.String("product_url", *product.Href))
		product.OriginID = api.PtrInt64(0)
	} else {
		originID, err := strconv.ParseInt(findSub[1], 10, 64)
		if err != nil {
			c.log.Info("parse_origin_id_in_page_error", zap.String("product_url", *product.Href))
		}
		product.OriginID = &originID
	}
	err = c.store.UpdateOriginID(&product)
	if err != nil {
		return fmt.Errorf("UpdateOriginID_error: %w", err)
	}
	c.log.Info("CollectOriginID",
		zap.String("title", product.Title),
		zap.Int64p("origin_id", product.OriginID),
		zap.Duration("elapsed", time.Since(start)),
	)
	return nil
}

func (c *Collector) Notify() error {
	users, err := c.store.SelectUsers()
	if err != nil {
		return fmt.Errorf("select_users_error: %w", err)
	}

	for _, user := range users {
		products, err := c.store.SelectGoodProducts(50, user.ID, 18, user.Rating, user.Discount)
		if err != nil {
			return err
		}
		c.log.Debug("Notify", zap.Any("products", len(products)),
			zap.Any("products", products),
			zap.Any("user", user),
		)
		for _, product := range products {
			var tpl bytes.Buffer
			product.PostOldPrice = product.OldPrice > product.OriginPrice
			err := c.msgTmpl.Execute(&tpl, product)
			if err != nil {
				c.log.Warn("execute_template_error", zap.Error(err))
				continue
			}
			err = c.tg.Send(user.ID, tpl.String())
			if err != nil {
				c.log.Warn("telegram_send_error", zap.Error(err))
				continue
			}
			user.SaleProducts = append(user.SaleProducts, store.SaleProduct{ID: product.ID})
		}
		err = c.store.UpdateUser(&user)
		if err != nil {
			c.log.Warn("update_user_error", zap.Error(err))
		}
	}
	return nil
}
func (c *Collector) Run() error {
	categoryTick := time.Tick(categoryPeriod)
	originProductsTick := time.Tick(CollectOriginProductsPeriod)
	collectPageTick := time.Tick(collectPagePeriod)

	err := c.CollectCategories()
	if err != nil {
		return fmt.Errorf("CollectCategories_error: %w", err)
	}
	err = c.CollectSaleProducts()
	if err != nil {
		return fmt.Errorf("CollectSaleProducts_error: %w", err)
	}
	err = c.CollectOriginProducts()
	if err != nil {
		return fmt.Errorf("CollectOriginProducts_error: %w", err)
	}

	for {
		select {
		case <-c.ctx.Done():
			c.log.Debug("exit_event_loop")
			return nil
		case <-c.newProductCh:
			//c.log.Debug("newProductCh", zap.Int("num", num))
			err := c.CollectOriginID()
			if err != nil {
				c.log.Warn("CollectOriginID_error", zap.Error(err))
			}
		case <-categoryTick:
			err := c.CollectCategories()
			if err != nil {
				c.log.Warn("CollectCategories_error", zap.Error(err))
			}
			err = c.CollectSaleProducts()
			if err != nil {
				c.log.Warn("CollectSaleProducts_error", zap.Error(err))
			}
		case <-originProductsTick:
			err := c.CollectOriginProducts()
			if err != nil {
				c.log.Warn("CollectOriginProducts_error", zap.Error(err))
			}
			err = c.Notify()
			if err != nil {
				c.log.Warn("Notify_error", zap.Error(err))
			}
		case <-collectPageTick:
			err := c.CollectOriginID()
			if err != nil {
				c.log.Warn("CollectOriginID_error", zap.Error(err))
			}
		}
	}
}
