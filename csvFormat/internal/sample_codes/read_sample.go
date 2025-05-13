package samplecodes

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadCSVSample() {
	f, err := os.Open("./sample_datas/read_sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
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
