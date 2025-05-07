package samplecodes

import (
	"encoding/json"
	"fmt"
)

type inputSample struct {
	Name        string `json:"name"`
	CompanyName string `json:"company_name,omitempty"`
}

type bottle struct {
	Name  string `json:"name"`
	Price int    `json:"price,omitempty"`
	KCal  *int   `json:"kcal,omitempty"`
}

func OmitemptySampleEncode() {
	// Omitemptyを指定した項目かつ肩のゼロ値であれば、Jsonにエンコードしない
	in := inputSample{Name: "サンプル太郎"}

	b, _ := json.Marshal(in)
	fmt.Println(string(b))
}

func OmitemptyDistinguishSample() {
	b := bottle{
		Name:  "ミネラルウォーター",
		Price: 0,
		// Omitemptyを指定したポインタ型の項目はゼロ値であっても出力される
		KCal: Int(0),
	}

	out, _ := json.Marshal(b)
	fmt.Println(string(out))
}

func Int(v int) *int {
	return &v
}
