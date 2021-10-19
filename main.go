package main

import (
	// "fmt"
	// "net/http"
	"strconv"

	"./model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	dbInit()

	r.GET("/helloworld", helloWorld)
	r.POST("/create", createSensingData)
	r.GET("/read/", readAllDB)
	r.GET("/read/:id", readOneDB)
	r.Run(":8080")
}

// DB初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", "model/test_db.sqlite3")
	if err != nil {
		panic("データベース開けず!(dbInit)")
	}
	db.AutoMigrate(&model.SensingTable{})
	defer db.Close()
}

// DB追加
func dbInsert(timestamp int, latitude float64, longitude float64) {
	db, err := gorm.Open("sqlite3", "model/test_db.sqlite3")
	if err != nil {
		panic("データベース開けず!(dbInsert)")
	}
	db.Create(&model.SensingTable{
		Timestamp: timestamp,
		Latitude:  latitude,
		Longitude: longitude,
	})
	defer db.Close()
}

// DB全取得
func dbGetAll() []model.SensingTable {
	db, err := gorm.Open("sqlite3", "model/test_db.sqlite3")
	if err != nil {
		panic("データベース開けず!(dbGetAll())")
	}
	var all []model.SensingTable
	db.Order("id desc").Find(&all)
	db.Close()
	return all
}

// DB一つ取得
func dbGetOne(id int) model.SensingTable {
	db, err := gorm.Open("sqlite3", "model/test_db.sqlite3")
	if err != nil {
		panic("データベース開けず!(dbGetOne())")
	}
	var one model.SensingTable
	db.First(&one, id)
	db.Close()
	return one
}

// a tutorial of get method
func helloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

// createSensingData : センシングデータを登録
func createSensingData(c *gin.Context) {
	var req model.SensingData
	c.BindJSON(&req)
	// mess := model.SensingData{}
	dbInsert(req.Timestamp, req.Latitude, req.Longitude)
	c.JSON(200, gin.H{
		"status":    200,
		"timestamp": req.Timestamp,
	})
}

// readAllDB : 登録データを全て読み込み
func readAllDB(c *gin.Context) {
	all := dbGetAll()
	c.JSON(200, gin.H{
		"all data": all,
	})
}

// readOneDB : 登録データを一つ読み込み
func readOneDB(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	one := dbGetOne(id)
	c.JSON(200, gin.H{
		"one data": one,
	})
}
