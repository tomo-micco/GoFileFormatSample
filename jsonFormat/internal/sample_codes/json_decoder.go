package samplecodes

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ip struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

// json.Decoder.Decodeを使用したデコード
// io.Readerインタフェースを満たしている型のデータを元にデコードできる。
func JsonDecodeSample() {
	f, err := os.Open("../sample_data/ip.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var response ip
	if err := json.NewDecoder(f).Decode(&response); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", response)
}

// json.Unmarshalを使用したデコード
// []byte型を扱う場合に使用できる。
func JsonUnmarshalSample() {
	s := `{"origin":"255.255.255.255", "url":"https://jsonunmarshalsample_hoge.jp/get"}`

	var response ip
	if err := json.Unmarshal([]byte(s), &response); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", response)
}
