package store

import (
	"fmt"
	"github.com/aibotsoft/bot-rozetka-category/pkg/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Store struct {
	log *zap.Logger
	cfg *config.Config
	db  *gorm.DB
}

type Category struct {
	ID           int64 `copier:"Id" gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Count        int64          `json:"count,omitempty"`
	Name         string         `json:"name,omitempty"`
	Title        string         `json:"title,omitempty"`
	ParentId     int64          `json:"parent_id"`
	GoodsCount   int64          `json:"goods_count"`
	SaleProducts []SaleProduct
}

func (c *Category) SaleProductMap() map[int64]bool {
	m := make(map[int64]bool)
	for i := 0; i < len(c.SaleProducts); i++ {
		m[c.SaleProducts[i].ID] = true
	}
	return m
}

type SaleProduct struct {
	ID         int64 `copier:"Id" gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Title      string         `json:"title"`
	Price      *float64       `json:"price,omitempty"`
	OldPrice   *float64       `json:"old_price,omitempty"`
	PricePcs   *string        `json:"price_pcs,omitempty"`
	Href       *string        `json:"href,omitempty"`
	Status     *string        `json:"status,omitempty"`
	CategoryId *int64         `json:"category_id,omitempty"`
	Brand      *string        `json:"brand,omitempty"`
	BrandId    *int64         `json:"brand_id,omitempty"`
	SellStatus *string        `json:"sell_status,omitempty"`
	Docket     *string        `json:"docket,omitempty"`
	OriginID   *int64
	ImageMain  *string `json:"image_main,omitempty"`
}
type OriginProduct struct {
	ID             int64 `copier:"Id" gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Title          string         `json:"title"`
	Price          *float64       `json:"price,omitempty"`
	OldPrice       *float64       `json:"old_price,omitempty"`
	PricePcs       *string        `json:"price_pcs,omitempty"`
	Href           *string        `json:"href,omitempty"`
	Status         *string        `json:"status,omitempty"`
	CategoryId     *int64         `json:"category_id,omitempty"`
	Brand          *string        `json:"brand,omitempty"`
	BrandId        *int64         `json:"brand_id,omitempty"`
	SellStatus     *string        `json:"sell_status,omitempty"`
	Docket         *string        `json:"docket,omitempty"`
	CommentsAmount *float64       `json:"comments_amount,omitempty"`
	CommentsMark   *float64       `json:"comments_mark,omitempty"`
	Stars          *string        `json:"stars,omitempty"`
	SellerId       *float64       `json:"seller_id,omitempty"`
	MerchantId     *float64       `json:"merchant_id,omitempty"`
}
type User struct {
	ID           int64 `gorm:"primaryKey"`
	FirstName    string
	UserName     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	SaleProducts []SaleProduct  `gorm:"many2many:user_sale_products;"`
	Discount     int64
	Rating       float64
}

func NewStore(zlog *zap.Logger, cfg *config.Config) (*Store, error) {
	newLogger := logger.New(
		log.New(os.Stderr, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             3 * time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn,     // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,           // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("connect_to_database_error: %w", err)
	}
	return &Store{log: zlog, cfg: cfg, db: db}, nil
}

func (s *Store) Migrate() error {
	err := s.db.AutoMigrate(&Category{}, &OriginProduct{}, &SaleProduct{}, &User{})
	if err != nil {
		return fmt.Errorf("auto_migrate_error: %w", err)
	}
	return nil
}
func (s *Store) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (s *Store) DeleteCategoriesNotInList(ids []int64) error {
	var categories []Category
	res := s.db.Clauses(clause.Returning{}).Where("id not in ?", ids).Delete(&categories)
	if res.RowsAffected > 0 {
		s.log.Info("delete_categories",
			zap.Int64("RowsAffected", res.RowsAffected),
			zap.Int("count", len(categories)),
			zap.Any("categories", categories),
		)
	}
	return res.Error
}
func (s *Store) GetCategories() ([]Category, error) {
	var categories []Category
	s.db.Preload(clause.Associations).Find(&categories)
	//s.log.Info("", zap.Any("", categories))
	return categories, nil
}

func (s *Store) SaveCategories(categories *[]Category) error {
	res := s.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(categories)
	return res.Error
}
func (s *Store) SaveSaleProducts(products *[]SaleProduct) error {
	err := s.db.Omit("origin_id").Save(products).Error
	return err
}

func (s *Store) UpdateSaleProducts(products *[]SaleProduct) error {
	err := s.db.Omit("origin_id").Save(products).Error
	if err != nil {
		for _, p := range *products {
			err := s.db.Omit("origin_id").Save(p).Error
			if err != nil {
				s.log.Warn("update_sale_product_error", zap.Error(err), zap.Any("product", p))
				err := s.db.Delete(p).Error
				if err != nil {
					s.log.Warn("delete_sale_product_error", zap.Error(err), zap.Any("product", p))
				}
			} else {
				//s.log.Info("save_sale_product_done", zap.Any("product", p))
			}
		}
	}
	return nil
}

func (s *Store) SaveOriginProducts(products *[]OriginProduct) error {
	res := s.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(products)
	return res.Error
}

func (s *Store) SaveSaleProduct(product *SaleProduct) error {
	res := s.db.Save(product)
	return res.Error
}

func (s *Store) UpdateOriginID(product *SaleProduct) error {
	res := s.db.Model(product).Update("origin_id", product.OriginID)
	return res.Error
}

func (s *Store) SelectProductWithoutOriginID() (product SaleProduct, err error) {
	res := s.db.Where("origin_id is null").Where("sell_status = 'available'").Find(&product)
	return product, res.Error
}

func (s *Store) SelectOriginIDList(limit int) (idList []int64, err error) {
	res := s.db.Model(&SaleProduct{}).Limit(limit).Distinct("sale_products.origin_id").Joins("left join origin_products on origin_products.id = sale_products.origin_id").Where("origin_products.id isnull").Where("sale_products.origin_id > 0").Scan(&idList)
	return idList, res.Error
}

func (s *Store) SelectOriginProductsForRefresh(limit int) (idList []int64, err error) {
	res := s.db.Model(&OriginProduct{}).Limit(limit).Select("id").Order("updated_at").Scan(&idList)
	return idList, res.Error
}
func (s *Store) SelectSaleProductsForRefresh(limit int) (idList []int64, err error) {
	res := s.db.Model(&SaleProduct{}).Limit(limit).Select("id").Order("updated_at").Where("updated_at < now() - interval '1 hour'").Scan(&idList)
	return idList, res.Error
}

func (s *Store) SaveUser(u *User) error {
	res := s.db.Debug().Clauses(clause.OnConflict{UpdateAll: true}).Create(u)
	return res.Error
}
func (s *Store) UpdateUser(u *User) error {
	res := s.db.Omit("SaleProducts.*").Save(u)
	//err := s.db.Debug().Model(u).Omit("SaleProducts.*").Association("SaleProducts").Append(u.SaleProducts)
	return res.Error
}

func (s *Store) DeleteUser(id int64) error {
	res := s.db.Debug().Delete(&User{ID: id})
	return res.Error
}

type TopDiscount struct {
	ID             int64
	Title          string
	Price          float64
	Href           string
	ImageMain      string
	OriginHref     string
	Brand          string
	OriginPrice    float64
	OldPrice       float64
	CommentsMark   float64
	CommentsAmount int64
	Discount       int64
	Reason         string
	PostOldPrice   bool
}

func (s *Store) SelectGoodProducts(limit int, userID int64, minComments int64, minMark float64, minDiscount int64) ([]TopDiscount, error) {
	var td []TopDiscount
	res := s.db.Limit(limit).Model(&SaleProduct{}).
		Select("sale_products.id, sale_products.title, sale_products.price, sale_products.href, sale_products.brand, "+
			"sale_products.image_main, op.price as origin_price, op.comments_mark, op.href as origin_href, "+
			"substring(sale_products.docket from 52 for position('</' in sale_products.docket) - 52) as reason, "+
			"op.comments_amount, op.old_price, (100 - sale_products.price * 100 / op.price)::int discount").
		Joins("join origin_products op on op.id = sale_products.origin_id").
		Joins("left join user_sale_products usp on usp.sale_product_id = sale_products.id and usp.user_id = ?", userID).
		Where("op.comments_amount > ?", minComments).
		Where("op.comments_mark > ?", minMark).
		Where("sale_products.sell_status = 'available'").
		Where("(100-sale_products.price * 100 / op.price) > ?", minDiscount).
		Where("usp.sale_product_id isnull").
		Scan(&td)
	return td, res.Error
}

func (s *Store) SelectUsers() (users []User, err error) {
	res := s.db.Find(&users)
	return users, res.Error
}
