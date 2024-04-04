package snowflake

import (
	"fmt"
	"github.com/crosscode-nl/snowflake/internal/codecs/base64"
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
	id = ID(0xA000B00F0A023452)
	fmt.Println(id.Base64String())
	id = ID(math.MaxUint64)
	fmt.Println(id.Base64String())
	// Output:
	// AQAAAAAAAAA
	// UjQCCg+wAKA
	// //////////8
}

func ExampleIDFromBase64String() {
	id := IDFromBase64String("AQAAAAAAAAA")
	fmt.Println(uint64(id))
	id = IDFromBase64String("UjQCCg+wAKA")
	fmt.Println(uint64(id))
	id = IDFromBase64String("//////////8")
	fmt.Println(uint64(id))
	// Output:
	// 1
	// 11529408624707384402
	// 18446744073709551615
}

// ExampleID_Base64StringCustom_urlAlphabet is an example of the ID Base64StringCustom method using the url alphabet
func ExampleID_Base64StringCustom_urlAlphabet() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Base64StringCustom(base64.UrlAlphabet))
	id = ID(0xA000B00F0A023452)
	fmt.Println(id.Base64StringCustom(base64.UrlAlphabet))
	id = ID(math.MaxUint64)
	fmt.Println(id.Base64StringCustom(base64.UrlAlphabet))
	// Output:
	// AQAAAAAAAAA
	// UjQCCg-wAKA
	// __________8
}

// ExampleIDFromBase64StringCustom_urlAlphabet is an example of the IDFromBase64StringCustom function using the url alphabet
func ExampleIDFromBase64StringCustom_urlAlphabet() {
	id := IDFromBase64StringCustom("AQAAAAAAAAA", base64.UrlAlphabetLookup)
	fmt.Println(uint64(id))
	id = IDFromBase64StringCustom("UjQCCg-wAKA", base64.UrlAlphabetLookup)
	fmt.Println(uint64(id))
	id = IDFromBase64StringCustom("__________8", base64.UrlAlphabetLookup)
	fmt.Println(uint64(id))
	// Output:
	// 1
	// 11529408624707384402
	// 18446744073709551615

}

// ExampleID_Base64StringCustom_mimeAlphabet is an example of the ID Base64StringCustom method using the mime alphabet
func ExampleID_Base64StringCustom_mimeAlphabet() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Base64StringCustom(base64.MimeAlphabet))
	id = ID(0xA000B00F0A023452)
	fmt.Println(id.Base64StringCustom(base64.MimeAlphabet))
	id = ID(math.MaxUint64)
	fmt.Println(id.Base64StringCustom(base64.MimeAlphabet))
	// Output:
	// AQAAAAAAAAA
	// UjQCCg+wAKA
	// ,,,,,,,,,,8
}

// ExampleIDFromBase64StringCustom_mimeAlphabet is an example of the IDFromBase64StringCustom function using the mime alphabet
func ExampleIDFromBase64StringCustom_mimeAlphabet() {
	id := IDFromBase64StringCustom("AQAAAAAAAAA", base64.MimeAlphabetLookup)
	fmt.Println(uint64(id))
	id = IDFromBase64StringCustom("UjQCCg+wAKA", base64.MimeAlphabetLookup)
	fmt.Println(uint64(id))
	id = IDFromBase64StringCustom(",,,,,,,,,,8", base64.MimeAlphabetLookup)
	fmt.Println(uint64(id))
	// Output:
	// 1
	// 11529408624707384402
	// 18446744073709551615

}

// ExampleID_Influx64String is an example of the ID Influx64String method
func ExampleID_Influx64String() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Influx64String())
	id = ID(0xA000B00F0A023452)
	fmt.Println(id.Influx64String())
	id = ID(math.MaxUint64)
	fmt.Println(id.Influx64String())
	// Output:
	// 00000000001
	// A00h0xA0ZHI
	// F~~~~~~~~~~
}

// ExampleID_Influx64StringCustom is an example of the ID Influx64StringCustom method using the url alphabet
func ExampleID_Influx64StringCustom() {
	id := ID(0x0000000000000001)
	fmt.Println(id.Influx64StringCustom(base64.UrlAlphabet))
	id = ID(0xA000B00F0A023452)
	fmt.Println(id.Influx64StringCustom(base64.UrlAlphabet))
	id = ID(math.MaxUint64)
	fmt.Println(id.Influx64StringCustom(base64.UrlAlphabet))
	// Output:
	// AAAAAAAAAAB
	// KAAsA8KAjRS
	// P__________
}

// ExampleIDFromInflux64StringCustom is an example of the IDFromInflux64String function using the url alphabet
func ExampleIDFromInflux64StringCustom() {
	id := IDFromInflux64StringCustom("AAAAAAAAAAB", base64.UrlAlphabetLookup)
	fmt.Println(uint64(id))
	id = IDFromInflux64StringCustom("KAAsA8KAjRS", base64.UrlAlphabetLookup)
	fmt.Println(uint64(id))
	id = IDFromInflux64StringCustom("P__________", base64.UrlAlphabetLookup)
	fmt.Println(uint64(id))
	// Output:
	// 1
	// 11529408624707384402
	// 18446744073709551615
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

func ExampleIDFromLowerHexString() {
	id := IDFromLowerHexString("0000000000000001")
	fmt.Println(uint64(id))
	id = IDFromLowerHexString("a000b00f0a023452")
	fmt.Println(uint64(id))
	id = IDFromLowerHexString("ffffffffffffffff")
	fmt.Println(uint64(id))
	// Output:
	// 1
	// 11529408624707384402
	// 18446744073709551615
}

// ExampleID_String is an example of the ID String method
func ExampleID_String() {
	id := ID(0x0000000000000001)
	fmt.Println(id.String())
	id = ID(11529408624707384402)
	fmt.Println(id.String())
	id = ID(math.MaxUint64)
	fmt.Println(id.String())
	// Output:
	// 00000000001
	// A00h0xA0ZHI
	// F~~~~~~~~~~
}

// ExampleIDFromString is an example of the IDFromString function
func ExampleIDFromString() {
	id := IDFromString("00000000001")
	fmt.Println(uint64(id))
	id = IDFromString("A00h0xA0ZHI")
	fmt.Println(uint64(id))
	id = IDFromString("F~~~~~~~~~~")
	fmt.Println(uint64(id))
	// Output:
	// 1
	// 11529408624707384402
	// 18446744073709551615
}

// ExampleID_UpperHexString is an example of the ID UpperHexString method
func ExampleID_UpperHexString() {
	id := ID(0x0000000000000001)
	fmt.Println(id.UpperHexString())
	id = ID(0xA000B00F0A023452)
	fmt.Println(id.UpperHexString())
	id = ID(math.MaxUint64)
	fmt.Println(id.UpperHexString())
	// Output:
	// 0000000000000001
	// A000B00F0A023452
	// FFFFFFFFFFFFFFFF
}

func ExampleIDFromUpperHexString() {
	id := IDFromUpperHexString("0000000000000001")
	fmt.Println(uint64(id))
	id = IDFromUpperHexString("A000B00F0A023452")
	fmt.Println(uint64(id))
	id = IDFromUpperHexString("FFFFFFFFFFFFFFFF")
	fmt.Println(uint64(id))
	// Output:
	// 1
	// 11529408624707384402
	// 18446744073709551615
}
