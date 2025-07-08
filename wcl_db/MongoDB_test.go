package wcl_db

import (
	"fmt"
	"testing"
)

/*
func TestMongoDBClientMapAdd(t *testing.T) {
	clientMap := GetMongoDBClientMap()
	err := clientMap.Add("mongodb://root:root@127.0.0.1:27017/admin")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(clientMap)
}

func TestMongoDBInsertOne(t *testing.T) {
	m := MongoDB{
		URI:        "mongodb://root:root@127.0.0.1:27017/admin",
		Database:   "Bach",
		Collection: "BachUser",
		Data: map[string]interface{}{
			"id":       4,
			"level":    4,
			"username": "1111111",
			"password": "1111111",
		},
	}
	err := m.InsertOne()
	if err != nil {
		fmt.Println(err)
	}
}
*/
func TestMongoDBQuery(t *testing.T) {
	m := MongoDB{
		URI: "mongodb://root:root@127.0.0.1:27017/admin",
	}
	req := MgoReq{
		Database:   "Bach1",
		Collection: "BachUser1",
	}
	result, err := m.Query(&req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
