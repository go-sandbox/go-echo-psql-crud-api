package main

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

// テーブル行の構造体
type Girl struct {
	Id   int64
	Age  int64
	Name string
}

// テーブルの構造体
type Girls struct {
	Girls []Girl
}

func main() {
	// bootstrap的なやつ
	// DBの初期設定はここで定義する
	initMigrate()
	run()
}

// bootstrap
func initMigrate() {
	db, err := gorm.Open("postgres", "host=localhost port=25432 user=root password=root dbname=sandbox sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Girl{})
}

func run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// パス＆ハンドラ関数を設定
	e.GET("/girls", showAllGirls)
	e.GET("/girl/:id", showGirl)
	e.PUT("/girl/:id", updateGirls)
	e.POST("/girl", newGirls)
	e.DELETE("/girl/:id", deleteGirl)

	log.Fatal(e.Start(":8080"))
}
