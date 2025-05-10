package samplecodes

import (
	"encoding/json"
	"fmt"
)

// Jsonの内容が動的に変化する場合のサンプル
type Response struct {
	Type      string `json:"type"`
	Timestamp int    `json:"timestamp"`
	// Payload を具体的な構造体に展開せず、json.RowMessageとして保持する
	Payload json.RawMessage `json:"payload"`
}

// message のペイロードを扱う構造体
type Message struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Message   string  `json:"message"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// sensor のペイロードを扱う構造体
type Sensor struct {
	ID        string `json:"id"`
	DeviceID  string `json:"device_id"`
	Result    string `json:"result"`
	ProductID string `json:"product_id"`
}

// 動的な内容に合わせてデコードするサンプル
func JsonContentsDynamicallyDecodeSample(jsonStr []byte) {
	// 一度 json.RawMessageにデコードしてから type に合わせて再度デコードする
	var r Response
	_ = json.Unmarshal(jsonStr, &r)

	switch r.Type {
	case "message":
		var m Message
		_ = json.Unmarshal(r.Payload, &m)
		fmt.Printf("%s\n%+v\n", r.Type, m)
	case "sensor":
		var s Sensor
		_ = json.Unmarshal(r.Payload, &s)
		fmt.Printf("%s\n%+v\n", r.Type, s)
	default:
		fmt.Printf("Unknown type :%s\n", r.Type)
	}
}
