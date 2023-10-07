package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// 链接数据库，返回表指针collection
func ConnectCollection(url string, dbName string, colName string) *mongo.Collection {
	// Rest of the code will go here
	// Set client options 设置连接参数
	clientOptions := options.Client().ApplyURI(url)
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
	//没有会自己创建
	collection := client.Database(dbName).Collection(colName)
	return collection
}

// 插入一条记录   item interface{}  使用空接口作为参数可以
func InsertOne(collection *mongo.Collection, item interface{}) {

	insertResult, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

//加入关闭数据库和表函数

//crud
