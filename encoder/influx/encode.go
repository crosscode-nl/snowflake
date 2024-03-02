package influx

func Alphabet() [64]byte {
	return [64]byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z', '_', 'a', 'b', 'c',
		'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w',
		'x', 'y', 'z', '~'}
}

// Encode encodes a number into an Influx style base64 string
func Encode(s *[11]byte, n uint64, d func() [64]byte) {
	digits := d()
	s[10], n = digits[n&0x3f], n>>6
	s[9], n = digits[n&0x3f], n>>6
	s[8], n = digits[n&0x3f], n>>6
	s[7], n = digits[n&0x3f], n>>6
	s[6], n = digits[n&0x3f], n>>6
	s[5], n = digits[n&0x3f], n>>6
	s[4], n = digits[n&0x3f], n>>6
	s[3], n = digits[n&0x3f], n>>6
	s[2], n = digits[n&0x3f], n>>6
	s[1], n = digits[n&0x3f], n>>6
	s[0] = digits[n&0x3f]
}
