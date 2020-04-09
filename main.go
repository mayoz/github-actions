package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const fileName = "storage/log.json"

type DataSet struct {
	Timestamp string `json:"timestamp"`
}

func main() {
	err := checkFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var data []DataSet

	_ = json.Unmarshal(file, &data)

	data = append(data, DataSet{
		Timestamp: time.Now().Format(time.RFC3339),
	})

	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fileName, dataBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
