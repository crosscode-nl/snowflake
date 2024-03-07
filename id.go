package snowflake

import (
	"github.com/crosscode-nl/snowflake/internal/encoder/base64"
	"github.com/crosscode-nl/snowflake/internal/encoder/base64/influx"
	"github.com/crosscode-nl/snowflake/internal/encoder/hex"
)

// ID is a snowflake ID
type ID uint64

type Alphabet func() [64]byte

// String returns a string representation of the snowflake ID
func (id ID) String() string {
	return id.Influx64String()
}

// LowerHexString returns a lower case hex string of the snowflake ID
func (id ID) LowerHexString() string {
	var b [16]byte
	hex.Encode(&b, uint64(id), hex.Lower)
	return string(b[:])
}

// UpperHexString returns an upper case hex string of the snowflake ID
func (id ID) UpperHexString() string {
	var b [16]byte
	hex.Encode(&b, uint64(id), hex.Upper)
	return string(b[:])
}

// Base64String returns a base64 string of the snowflake ID
func (id ID) Base64String() string {
	var b [11]byte
	base64.Encode(&b, uint64(id), base64.Alphabet)
	return string(b[:])
}

// Base64StringCustom returns a custom base64 string of the snowflake ID
func (id ID) Base64StringCustom(alphabet Alphabet) string {
	var b [11]byte
	base64.Encode(&b, uint64(id), alphabet)
	return string(b[:])
}

// Influx64String returns an Influx style base64 string of the snowflake ID
func (id ID) Influx64String() string {
	var b [11]byte
	influx.Encode(&b, uint64(id), influx.Alphabet)
	return string(b[:])
}

// Influx64StringCustom returns a custom Influx style base64 string of the snowflake ID
func (id ID) Influx64StringCustom(alphabet Alphabet) string {
	var b [11]byte
	influx.Encode(&b, uint64(id), alphabet)
	return string(b[:])
}
