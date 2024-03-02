package snowflake

import (
	"fmt"
	"github.com/crosscode-nl/snowflake/encoder/base64"
	"math"
	"testing"
)

func BenchmarkID_Base64String(b *testing.B) {
	id := ID(0x0000000000000001)
	for i := 0; i < b.N; i++ {
		id.Base64String()
	}
}

func BenchmarkID_Base64InfluxString(b *testing.B) {
	id := ID(0x0000000000000001)
	for i := 0; i < b.N; i++ {
		id.Influx64String()
	}
}

func BenchmarkID_LowerHexString(b *testing.B) {
	id := ID(0x0000000000000001)
	for i := 0; i < b.N; i++ {
		id.LowerHexString()
	}
}

func BenchmarkID_UpperHexString(b *testing.B) {
	id := ID(0x0000000000000001)
	for i := 0; i < b.N; i++ {
		id.UpperHexString()
	}
}

func ExampleID_Base64String() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Base64String())
	id = ID(math.MaxUint64)
	fmt.Println(id.Base64String())
	// Output:
	// AQAAAAAAAAA
	// //////////8
}

func ExampleID_Base64StringCustom_urlAlphabet() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Base64StringCustom(base64.UrlAlphabet))
	id = ID(math.MaxUint64)
	fmt.Println(id.Base64StringCustom(base64.UrlAlphabet))
	// Output:
	// AQAAAAAAAAA
	// __________8
}

func ExampleID_Base64StringCustom_mimeAlphabet() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Base64StringCustom(base64.MimeAlphabet))
	id = ID(math.MaxUint64)
	fmt.Println(id.Base64StringCustom(base64.MimeAlphabet))
	// Output:
	// AQAAAAAAAAA
	// ,,,,,,,,,,8
}

func ExampleID_Influx64String() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Influx64String())
	id = ID(math.MaxUint64)
	fmt.Println(id.Influx64String())
	// Output:
	// 00000000001
	// F~~~~~~~~~~
}

func ExampleID_Influx64StringCustom() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Influx64StringCustom(base64.UrlAlphabet))
	id = ID(math.MaxUint64)
	fmt.Println(id.Influx64StringCustom(base64.UrlAlphabet))
	// Output:
	// AAAAAAAAAAB
	// P__________

}

func ExampleID_LowerHexString() {
	id := ID(0x0000000000000001)
	fmt.Println(id.LowerHexString())

	id = ID(0xA000B00F0A023452)
	fmt.Println(id.LowerHexString())

	id = ID(math.MaxUint64)
	fmt.Println(id.LowerHexString())
	// Output:
	// 0000000000000001
	// a000b00f0a023452
	// ffffffffffffffff
}

func ExampleID_String() {
	id := ID(0x0000000000000001)
	fmt.Println(id.String())
	id = ID(math.MaxUint64)
	fmt.Println(id.String())
	// Output:
	// 0000000000000001
	// FFFFFFFFFFFFFFFF
}

func ExampleID_UpperHexString() {
	id := ID(0x0000000000000001)
	fmt.Println(id.UpperHexString())
	id = ID(math.MaxUint64)
	fmt.Println(id.UpperHexString())
	// Output:
	// 0000000000000001
	// FFFFFFFFFFFFFFFF
}
