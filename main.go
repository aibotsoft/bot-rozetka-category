package main

import (
	"context"
	"fmt"
	"github.com/aibotsoft/bot-rozetka-category/api"
	"github.com/aibotsoft/bot-rozetka-category/pkg/config"
	"github.com/caarlos0/env/v6"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"time"
)

const rootCategoryID = "83850"

type Category struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Count     int64          `json:"count,omitempty"`
	Name      string         `json:"name,omitempty"`
	Title     string         `json:"title,omitempty"`
}

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	//dsn := "host=localhost user=postgres password=postgres dbname=rozetka port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Category{})
	if err != nil {
		panic(err)
	}

	//fmt.Println("hello world", cfg.ApiDebug, db)
	conf := api.NewConfiguration()
	conf.Debug = cfg.ApiDebug
	client := api.NewAPIClient(conf)
	categories, err := getCategories(client)
	if err != nil {
		panic(err)
	}
	//fmt.Println(categories)
	result := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&categories)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	//err := http.ListenAndServe(":8081", nil)
	//if err != nil {
	//	panic(err)
	//}
}
func getCategories(client *api.APIClient) ([]Category, error) {
	resp, _, err := client.CategoriesApi.GetChildren(context.Background()).Lang("ru").CategoryId(rootCategoryID).Execute()
	if err != nil {
		return nil, err
	}
	children := resp.Data.GetChildren()
	var categories []Category
	for _, child := range children {
		//fmt.Println(child.GetId())
		categories = append(categories, Category{
			ID:    uint(child.GetId()),
			Count: child.GetCount(),
			Name:  child.GetName(),
			Title: child.GetTitle(),
		})
	}
	return categories, nil
}
