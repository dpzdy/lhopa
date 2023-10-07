package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
