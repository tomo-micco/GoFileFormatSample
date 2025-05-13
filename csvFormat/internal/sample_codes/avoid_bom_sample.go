package samplecodes

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func ReadWithBomSample() {

	f, err := os.Open("./sample_datas/avoid_bom_sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := unicode.UTF8.NewDecoder()
	r := csv.NewReader(transform.NewReader(f, decoder))
	// 区切り文字を変更したい場合はここで設定する
	// r.Comma = '\t'
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
