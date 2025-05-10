package main

import samplecodes "github.com/tomo-micco/GoFileFormatSample/JsonFormat/internal/sample_codes"

func main() {
	// デコード
	samplecodes.JsonDecodeSample()
	samplecodes.JsonUnmarshalSample()

	// エンコード
	user := samplecodes.User{Id: "1", Name: "John", Address: "SampleAddress", EMail: "hogehoeg_hogehoge@hoge.jphoge"}
	user.JsonEncodeSample()
	user.JsonMarshalSample()

	// スライスのエンコード
	samplecodes.SliceEncodeSample()

	// omitemptyのエンコード
	samplecodes.OmitemptySampleEncode()
	samplecodes.OmitemptyDistinguishSample()

	// 未知フィールドのエラー
	// samplecodes.ThrowErrorUnknownFieldsSample()

	// Json.Marshalの拡張サンプル
	samplecodes.JsonEncoderExtendSample()
	// Json.Unmarshalの拡張サンプル
	samplecodes.JsonDecoderExtendSample()

	// 動的な内容のデコードサンプル
	jsonStr := `{"type":"message","timestamp":1639266962,"payload":{"id":"123456","user_id":"ABC789","message":"こんにちは","latitude":123.45, "longitude":12.32}}`
	samplecodes.JsonContentsDynamicallyDecodeSample([]byte(jsonStr))
	jsonStr = `{"type":"sensor","timestamp":16392667892,"payload":{"id":"98765","device_id":"ZX-123","result":"ok","product_id":"1234"}}`
	samplecodes.JsonContentsDynamicallyDecodeSample([]byte(jsonStr))
	jsonStr = `{"type":"camera","timestamp":16392698013,"payload":{"id":"6473","camera_id":"C-123","result":"ok","product_id":"1234"}}`
	samplecodes.JsonContentsDynamicallyDecodeSample([]byte(jsonStr))
}
