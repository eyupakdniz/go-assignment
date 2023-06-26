package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	/*
	logData := LogData{}
	logData.Level = "debug"
	logData.Timestamp = "2023-06-20T16:37:31.738+0300"
	logData.Caller = "core/device_handler.go:3435"
	logData.Message = "SetKVStoreBackend"
	logData.InstanceID = "voltha-adapter-openonu"
	logData.Package = "core"
	logData.IPTarget = "redis"
	logData.BasePath = "service/voltha/omci_mibs/go_templates"
	logData.DeviceID = "aaaabbbbbbcccxcccc"
	*/
	
	str := LogData{"debug",
	"2023-06-20T16:37:31.738+0300",
	"core/device_handler.go:3435",
	"SetKVStoreBackend",
	"voltha-adapter-openonu",
	"core",
	"redis",
	"service/voltha/omci_mibs/go_templates",
	"aaaabbbbbbcccxcccc"}
	
	write(str)
	read()
}

func write(str LogData){
	content, err := json.Marshal(str)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("file.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func read(){
	veri, err := ioutil.ReadFile("file.json") //content içindekiler byte dizisidir.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(veri)) 

	data := LogData{}
	err = json.Unmarshal(veri, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("")
	fmt.Println(data)

	/*fmt.Println("")
	jsonByte, _ := json.Marshal(data)
	fmt.Println(string(jsonByte))*/
}

