package samplecodes

import (
	"encoding/json"
	"fmt"
)

type user struct {
	UserID    string   `json:"user_id"`
	UserName  string   `json:"user_name"`
	Languages []string `json:"languages"`
}

func SliceEncodeSample() {
	u := user{
		UserID:   "001",
		UserName: "gopher",
		// 定義しない場合、nullでエンコードされるため、空配列にエンコードしたい場合は明示的に定義する必要あり
		Languages: []string{},
	}

	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}
