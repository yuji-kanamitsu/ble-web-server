package main

import (
	// "net/http"
	"fmt"
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
	r.POST("/create", createPostData)
	r.GET("/read/", readAllDB)
	r.GET("/read/:id", readOneDB)
	r.Run(":8080")
}

// DB初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", "model/ble_db.sqlite3")
	if err != nil {
		panic("データベース開けず!(dbInit)")
	}
	db.AutoMigrate(&model.SensorTable{})

	defer db.Close()
}

// DB追加
func dbInsert(key string, meta model.Meta, body []model.Body) {
	db, err := gorm.Open("sqlite3", "model/ble_db.sqlite3")
	if err != nil {
		panic("データベース開けず!(dbInsert)")
	}
	db.Create(&model.SensorTable{
		Key:  key,
		Meta: meta,
		Body: body,
	})
	defer db.Close()
}

// DB全取得
func dbGetAll() []model.SensorTable {
	db, err := gorm.Open("sqlite3", "model/ble_db.sqlite3")
	if err != nil {
		panic("データベース開けず!(dbGetAll())")
	}
	var all []model.SensorTable
	db.Order("id desc").Find(&all)
	db.Close()
	return all
}

// DB一つ取得
func dbGetOne(id int) model.SensorTable {
	db, err := gorm.Open("sqlite3", "model/ble_db.sqlite3")
	if err != nil {
		panic("データベース開けず!(dbGetOne())")
	}
	var one model.SensorTable
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
func createPostData(c *gin.Context) {
	var req model.PostData
	c.BindJSON(&req)
	// mess := model.SensorData{}
	// dbInsert(req.Key, req.Meta, req.Body)
	fmt.Println(req.Key)
	fmt.Println(req.Meta)
	fmt.Println(req.Body)
	c.JSON(200, gin.H{
		"status": 200,
		"key":    req.Key,
		"meta":   req.Meta,
		"body":   req.Body,
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
