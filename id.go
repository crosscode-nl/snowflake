package snowflake

import (
	"github.com/crosscode-nl/snowflake/internal/codecs/base64"
	"github.com/crosscode-nl/snowflake/internal/codecs/base64/influx"
	"github.com/crosscode-nl/snowflake/internal/codecs/hex"
)

// ID is a snowflake ID
type ID uint64

type Alphabet func() [64]byte
type AlphabetLookup func() map[byte]uint64

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

// IDFromString returns a snowflake ID from a string
func IDFromString(s string) ID {
	return IDFromInflux64String(s)
}

// IDFromLowerHexString returns a snowflake ID from a lower case hex string
func IDFromLowerHexString(s string) ID {
	var b [16]byte
	copy(b[:], s)
	return ID(hex.Decode(&b, hex.LowerLookup))
}

// IDFromUpperHexString returns a snowflake ID from an upper case hex string
func IDFromUpperHexString(s string) ID {
	var b [16]byte
	copy(b[:], s)
	return ID(hex.Decode(&b, hex.UpperLookup))
}

// IDFromBase64String returns a snowflake ID from a base64 string
func IDFromBase64String(s string) ID {
	var b [11]byte
	copy(b[:], s)
	return ID(base64.Decode(&b, base64.AlphabetLookup))
}

// IDFromBase64StringCustom returns a snowflake ID from a custom base64 string
func IDFromBase64StringCustom(s string, alphabetLookup AlphabetLookup) ID {
	var b [11]byte
	copy(b[:], s)
	return ID(base64.Decode(&b, alphabetLookup))
}

// IDFromInflux64String returns a snowflake ID from an Influx style base64 string
func IDFromInflux64String(s string) ID {
	var b [11]byte
	copy(b[:], s)
	return ID(influx.Decode(&b, influx.AlphabetLookup))
}

// IDFromInflux64StringCustom returns a snowflake ID from a custom Influx style base64 string
func IDFromInflux64StringCustom(s string, alphabetLookup AlphabetLookup) ID {
	var b [11]byte
	copy(b[:], s)
	return ID(influx.Decode(&b, alphabetLookup))
}
