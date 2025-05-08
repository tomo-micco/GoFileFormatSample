package samplecodes

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type record struct {
	ProcessId string `json:"process_id"`
	DeletedAt JSTime `json:"deleted_at"`
}

type JSTime time.Time

// JSTime型をエンコードした際のカスタムロジック
func (t JSTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	if tt.IsZero() {
		return []byte("null"), nil
	}

	v := strconv.Itoa(int(tt.UnixMilli()))
	return []byte(v), nil
}

func JsonEncoderExtendSample() {
	r := &record{
		ProcessId: "0001",
		DeletedAt: JSTime(time.Now()),
	}

	b, _ := json.Marshal(r)
	fmt.Println(string(b))
}

func (t *JSTime) UnmarshalJSON(data []byte) error {
	var jsonNumber json.Number
	err := json.Unmarshal(data, &jsonNumber)
	if err != nil {
		return err
	}

	unixMillis, err := jsonNumber.Int64()
	if err != nil {
		return err
	}

	seconds := unixMillis / 1000
	nanoseconds := (unixMillis % 1000) * 1_000_000

	*t = JSTime(time.Unix(seconds, nanoseconds))
	return nil
}

func JsonDecoderExtendSample() {
	s := `{"process_id":"0002","deleted_at":1746748223615}`
	var r *record
	if err := json.Unmarshal([]byte(s), &r); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", time.Time(r.DeletedAt).Format(time.RFC3339Nano))
	fmt.Printf("%v\n", time.Now().Format(time.RFC3339Nano))
}
