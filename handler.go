package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

/*
  GET Route
*/
func showAllGirls(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=25432 user=root password=root dbname=sandbox sslmode=disable")
	defer db.Close()
	checkError(err)
	var girls []Girl
	db.Find(&girls)
	return c.JSON(http.StatusOK, girls)
}

func showGirl(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=25432 user=root password=root dbname=sandbox sslmode=disable")
	checkError(err)
	defer db.Close()
	var e Girl
	id := c.Param("id")
	db.Where("id=?", id).Find(&e)
	return c.JSON(http.StatusOK, e)
}

/*
  POST Route
*/
func newGirls(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=25432 user=root password=root dbname=sandbox sslmode=disable")
	defer db.Close()
	checkError(err)

	girl := new(Girl)
	if err := c.Bind(girl); err != nil {
		return err
	}

	db.Create(&girl)
	return c.String(http.StatusOK, "OK")
}

/*
  PUT Route
*/
func updateGirls(c echo.Context) error {
	var girl Girl
	db, err := gorm.Open("postgres", "host=localhost port=25432 user=root password=root dbname=sandbox sslmode=disable")
	checkError(err)
	defer db.Close()
	if err := c.Bind(&girl); err != nil {
		return err
	}
	paramId := c.Param("id")
	attrMap := map[string]interface{}{"age": girl.Age, "name": girl.Name}
	db.Model(&girl).Where("id= ?", paramId).Updates(attrMap)
	return c.NoContent(http.StatusOK)
}

/*
  DELETE Route
*/
func deleteGirl(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=localhost port=25432 user=root password=root dbname=sandbox sslmode=disable")
	checkError(err)
	defer db.Close()
	var e Girl
	id := c.Param("id")
	db.Where("id=?", id).Find(&e).Delete(&e)
	return c.JSON(http.StatusOK, e)
}

/*
  Error handling
*/
func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
