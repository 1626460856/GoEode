package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// ConnectToMongoDB 连接 MongoDB 数据库的函数
func ConnectToMongoDB() (*mongo.Client, error) {
	// 设置 MongoDB 连接选项
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到 MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接是否成功
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	if err == nil {
		fmt.Println("成功连接 MongoDB!")
	}
	return client, nil

}

func main() {
	// 连接到 MongoDB
	client, err := ConnectToMongoDB()
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer client.Disconnect(context.Background())

	// collection 实例现在指向了名为 "people" 的集合对象，你可以通过这个对象执行诸如插入文档、查询文档等操作。
	collection := ChooseSet(client, "test", "people")
	//插入单个文档
	user := bson.D{{"Name", "zhangsan"}, {"age", 30}}
	AddOneDocument(collection, user)
	//插入多个文档
	users := []interface{}{
		bson.D{{"Name", "lisi"}, {"age", 25}},
		bson.D{{"Name", "wangwu"}, {"age", 20}},
		bson.D{{"Name", "zhaoliu"}, {"age", 28}},
	}
	AddManyDocument(collection, users)
	fmt.Println(FindSetData(client, "test", "people"))

}
