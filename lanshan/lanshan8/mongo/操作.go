package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// ChooseSet 选择数据库中的某一集合
func ChooseSet(client *mongo.Client, database string, set string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(database).Collection(set)
	return collection
}

// FindSetData  函数用于查询指定集合中的所有文档，并返回结果切片
func FindSetData(client *mongo.Client, database string, set string) ([]bson.M, error) {
	var Collection *mongo.Collection = ChooseSet(client, database, set)
	// 建立查询 （无查询条件）
	cursor, err := Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

// InsertOneDocument 函数用于向指定集合中插入单个文档
func InsertOneDocument(collection *mongo.Collection, documents bson.D) (*mongo.InsertOneResult, error) {
	// 插入单个文档
	res, err := collection.InsertOne(context.Background(), documents)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func AddOneDocument(collection *mongo.Collection, documents bson.D) {
	res, err := InsertOneDocument(collection, documents)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.InsertedID)
}

// InsertManyDocuments 函数用于向指定集合中插入多个文档
func InsertManyDocuments(collection *mongo.Collection, documents []interface{}) (*mongo.InsertManyResult, error) {
	// 插入多个文档
	result, err := collection.InsertMany(context.Background(), documents)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func AddManyDocument(collection *mongo.Collection, documents []interface{}) {
	res, err := InsertManyDocuments(collection, documents)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.InsertedIDs)
}
