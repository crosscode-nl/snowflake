package snowflake

import (
	"github.com/crosscode-nl/snowflake/encoder/base64"
	"github.com/crosscode-nl/snowflake/encoder/hex"
	"github.com/crosscode-nl/snowflake/encoder/influx"
)

// ID is a snowflake ID
type ID uint64

type Alphabet func() [64]byte

func (id ID) String() string {
	return id.UpperHexString()
}

func (id ID) LowerHexString() string {
	var b [16]byte
	hex.Encode(&b, uint64(id), hex.Lower)
	return string(b[:])
}

func (id ID) UpperHexString() string {
	var b [16]byte
	hex.Encode(&b, uint64(id), hex.Upper)
	return string(b[:])
}

func (id ID) Base64String() string {
	var b [11]byte
	base64.Encode(&b, uint64(id), base64.Alphabet)
	return string(b[:])
}

func (id ID) Base64StringCustom(alphabet Alphabet) string {
	var b [11]byte
	base64.Encode(&b, uint64(id), alphabet)
	return string(b[:])
}

func (id ID) Influx64String() string {
	var b [11]byte
	influx.Encode(&b, uint64(id), influx.Alphabet)
	return string(b[:])
}

func (id ID) Influx64StringCustom(alphabet Alphabet) string {
	var b [11]byte
	influx.Encode(&b, uint64(id), alphabet)
	return string(b[:])
}
