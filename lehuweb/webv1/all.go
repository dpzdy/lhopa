package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"webv1/api"
	"webv1/data"
	"webv1/midware"
	"webv1/utils"
)

const (
	DATA_BASE               string = "dataEngineService"
	FACTORY_INFO_COLL       string = "factoryInfo"
	FACTORY_DATA_COLL       string = "factoryData"
	FACTORY_DATA_COUNT_COLL string = "factoryDataCount"
)

var mgo *mongo.Client
var FactoryInfoColl *mongo.Collection
var FactoryDataColl *mongo.Collection
var FactoryDataCountColl *mongo.Collection

func InitDb() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	// 连接到MongoDB
	var err error
	mgo, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err.Error())
	}

	// 检查连接
	if err := mgo.Ping(context.Background(), nil); err != nil {
		panic(err.Error())
	}

	FactoryInfoColl = mgo.Database(DATA_BASE).Collection(FACTORY_INFO_COLL)
	FactoryDataColl = mgo.Database(DATA_BASE).Collection(FACTORY_DATA_COLL)
	FactoryDataCountColl = mgo.Database(DATA_BASE).Collection(FACTORY_DATA_COUNT_COLL)
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	router := gin.Default()
	router.Use(midware.Cors())
	router.GET("/", api.GetInfo)
	router.Run()
}
func GetInfo(c *gin.Context) {

	info := data.GetAllInfo()
	c.JSON(http.StatusOK, gin.H{
		"data": info,
	})

}

type LehuIndexItem struct {
	Avatar  string            `json:"avatar"`
	Author  string            `json:"author"`
	Title   string            `json:"title,omitempty"`
	Content string            `json:"content"`
	Hot     string            `json:"hot"`
	CmtNum  string            `json:"cmtNum"`
	Date    string            `json:"date"`
	ImgURL  string            `json:"imgURL"`
	Tags    map[string]string `json:"Tags"`
	Cmt     []string          `json:"cmt"`
}

func GetAllInfo() []LehuIndexItem {
	fmt.Println("Find:")
	opt := options.Find()
	opt.SetLimit(5)

	col := mgo.Database("lehu").Collection("blog")

	//Find查询多条记录，将会返回一个迭代器，使用完记得关闭
	cor, err := col.Find(context.TODO(), bson.D{{"author", "老福特观察局"}}, opt)
	if err != nil {
		log.Fatalln(err)
	}
	defer cor.Close(context.TODO())

	res := []LehuIndexItem{}
	for cor.Next(context.TODO()) {
		var p LehuIndexItem
		cor.Decode(&p)
		//fmt.Println(p)
		res = append(res, p)
	}
	return res
}
