package samplecodes

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteCSVSample() {
	records := [][]string{
		{"書籍名", "出版年", "ページ数"},
		{"書籍名1", "2022", "234"},
		{"書籍名2", "1980", "150"},
		{"書籍名3", "2023", "334"},
		{"書籍名4", "2025", "330"},
	}

	f, err := os.OpenFile("./sample_datas/write_sample.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()
	// 書き込み時の区切り文字もここで設定可能
	w.Comma = '\t'

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
