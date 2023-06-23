package main

import (
	"fmt"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"context"
)

type LogData struct {
	Level      string `json:"level"`
	Timestamp  string `json:"ts"`
	Caller     string `json:"caller"`
	Message    string `json:"msg"`
	InstanceID string `json:"instanceId"`
	Package    string `json:"pkg"`
	IPTarget   string `json:"IpTarget"`
	BasePath   string `json:"BasePathKvStore"`
	DeviceID   string `json:"device-id"`
}


func main() {
	
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "mypassword",
		DB: 0,
	})

	ctx := context.Background()
	
	json, err := json.Marshal(LogData{"debug",
	"2023-06-20T16:37:31.738+0300",
	"core/device_handler.go:3435",
	"SetKVStoreBackend",
	"voltha-adapter-openonu",
	"core",
	"redis",
	"service/voltha/omci_mibs/go_templates",
	"aaaabbbbbbcccxcccc"})
	
	if err != nil{
		fmt.Println(err)
	}

	err = client.Set(ctx,"data", json,0).Err() //content, redisKey, data, timeout
	if err != nil {
		fmt.Println(err)
	}

	val,err := client.Get(ctx,"data").Result()
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(val)
}