package samplecodes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	EMail   string `json:"email"`
	// json.Marshalではチャネルや関数型、複素数型の値が含まれる場合はエラーとなる
	// JSONには変換しないが、型に含んでおきたい場合は”-”を指定しておく
	X func() `json:"-"`
}

// json.NewEncoder().Encode()を使用したエンコード処理
func (u User) JsonEncodeSample() {
	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", b.String())
}

// json.Marshalを使用したエンコード処理
func (u User) JsonMarshalSample() {
	b, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
