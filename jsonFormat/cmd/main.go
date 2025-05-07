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
	samplecodes.ThrowErrorUnknownFieldsSample()

}
