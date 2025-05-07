package samplecodes

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Rectangle struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	// Radius int `json:"radius"`
}

func ThrowErrorUnknownFieldsSample() {

	f, err := os.Open("../sample_data/rectangle.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var rect Rectangle
	d := json.NewDecoder(f)
	// DisallowUnknownFieldsを使用することで未知のフィールドが存在する場合はエラーとなる
	d.DisallowUnknownFields()
	if err := d.Decode(&rect); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rect)
}
