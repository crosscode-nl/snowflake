package snowflake

import (
	"fmt"
	"github.com/crosscode-nl/snowflake/encoder/base64"
	"math"
	"testing"
)

// BenchmarkID_Base64String benchmarks the Base64String method of the ID type
func BenchmarkID_Base64String(b *testing.B) {
	id := ID(0x0000000000000001)
	for i := 0; i < b.N; i++ {
		id.Base64String()
	}
}

// BenchmarkID_Base64StringCustom benchmarks the Base64StringCustom method of the ID type
func BenchmarkID_Base64InfluxString(b *testing.B) {
	id := ID(0x0000000000000001)
	for i := 0; i < b.N; i++ {
		id.Influx64String()
	}
}

// BenchmarkID_LowerHexString benchmarks the LowerHexString method of the ID type
func BenchmarkID_LowerHexString(b *testing.B) {
	id := ID(0x0000000000000001)
	for i := 0; i < b.N; i++ {
		id.LowerHexString()
	}
}

// BenchmarkID_String benchmarks the String method of the ID type
func BenchmarkID_UpperHexString(b *testing.B) {
	id := ID(0x0000000000000001)
	for i := 0; i < b.N; i++ {
		id.UpperHexString()
	}
}

// ExampleID_Base64String is an example of the ID Base64String method
func ExampleID_Base64String() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Base64String())
	id = ID(math.MaxUint64)
	fmt.Println(id.Base64String())
	// Output:
	// AQAAAAAAAAA
	// //////////8
}

// ExampleID_Base64StringCustom_urlAlphabet is an example of the ID Base64StringCustom method using the url alphabet
func ExampleID_Base64StringCustom_urlAlphabet() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Base64StringCustom(base64.UrlAlphabet))
	id = ID(math.MaxUint64)
	fmt.Println(id.Base64StringCustom(base64.UrlAlphabet))
	// Output:
	// AQAAAAAAAAA
	// __________8
}

// ExampleID_Base64StringCustom_mimeAlphabet is an example of the ID Base64StringCustom method using the mime alphabet
func ExampleID_Base64StringCustom_mimeAlphabet() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Base64StringCustom(base64.MimeAlphabet))
	id = ID(math.MaxUint64)
	fmt.Println(id.Base64StringCustom(base64.MimeAlphabet))
	// Output:
	// AQAAAAAAAAA
	// ,,,,,,,,,,8
}

// ExampleID_Influx64String is an example of the ID Influx64String method
func ExampleID_Influx64String() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Influx64String())
	id = ID(math.MaxUint64)
	fmt.Println(id.Influx64String())
	// Output:
	// 00000000001
	// F~~~~~~~~~~
}

// ExampleID_Influx64StringCustom is an example of the ID Influx64StringCustom method using the url alphabet
func ExampleID_Influx64StringCustom() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Influx64StringCustom(base64.UrlAlphabet))
	id = ID(math.MaxUint64)
	fmt.Println(id.Influx64StringCustom(base64.UrlAlphabet))
	// Output:
	// AAAAAAAAAAB
	// P__________

}

// ExampleID_LowerHexString is an example of the ID LowerHexString method
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

// ExampleID_String is an example of the ID String method
func ExampleID_String() {
	id := ID(0x0000000000000001)
	fmt.Println(id.String())
	id = ID(math.MaxUint64)
	fmt.Println(id.String())
	// Output:
	// 00000000001
	// F~~~~~~~~~~
}

// ExampleID_UpperHexString is an example of the ID UpperHexString method
func ExampleID_UpperHexString() {
	id := ID(0x0000000000000001)
	fmt.Println(id.UpperHexString())
	id = ID(math.MaxUint64)
	fmt.Println(id.UpperHexString())
	// Output:
	// 0000000000000001
	// FFFFFFFFFFFFFFFF
}
