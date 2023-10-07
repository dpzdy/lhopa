package db

import (
	"context"
	"fmt"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"spider/lehuSpider"
)

func TestDb(t *testing.T) {
	// Rest of the code will go here
	// Set client options 设置连接参数
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB 连接数据库
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection 测试连接
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	databases, err := client.ListDatabases(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases.TotalSize / 1024 / 1024 / 1024)
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

type Trainer struct {
	Name    string
	Age     int
	Address string
}

func TestInsert1(t *testing.T) {
	url := "mongodb://localhost:27017"
	collection := ConnectCollection(url, "lehu", "blog")
	driverPath := "../browser/chromedriver.exe" //准备工作中下载driver
	crawler := lehuSpider.Crawler{}
	service, driver := crawler.StartChrome(driverPath)
	defer service.Stop() // 停止chromedriver
	defer driver.Quit()  // 关闭浏览器
	item := lehuSpider.ParseContent("https://lofguancha.lofter.com/post/749935d7_2b8630cb9", service, driver)
	//********************************************************
	insertResult, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	//trainers := []interface{}{misty, brock}
	//
	//insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

// 1. 插入文档
func TestInsert(t *testing.T) {
	url := "mongodb://localhost:27017"
	collection := ConnectCollection(url, "mongodb_study", "student")
	//********************************************************
	ash := Trainer{"Dsh", 23, "Pallet Town"}
	//misty := Trainer{"Wisty", 10, "Cerulean City"}
	//brock := Trainer{"Crock", 15, "Pewter City"}

	InsertOne(collection, ash)

	//trainers := []interface{}{misty, brock}
	//insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

func TestUpdate(t *testing.T) {
	url := "mongodb://localhost:27017"
	collection := ConnectCollection(url, "mongodb_study", "student")
	//********************************************************
	filter := bson.D{{"name", "Ash"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

}
