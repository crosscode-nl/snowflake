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

func AlphabetLookup() map[byte]uint64 {
	return map[byte]uint64{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19,
		'K': 20, 'L': 21, 'M': 22, 'N': 23, 'O': 24, 'P': 25, 'Q': 26, 'R': 27, 'S': 28, 'T': 29,
		'U': 30, 'V': 31, 'W': 32, 'X': 33, 'Y': 34, 'Z': 35, '_': 36, 'a': 37, 'b': 38, 'c': 39,
		'd': 40, 'e': 41, 'f': 42, 'g': 43, 'h': 44, 'i': 45, 'j': 46, 'k': 47, 'l': 48, 'm': 49,
		'n': 50, 'o': 51, 'p': 52, 'q': 53, 'r': 54, 's': 55, 't': 56, 'u': 57, 'v': 58, 'w': 59,
		'x': 60, 'y': 61, 'z': 62, '~': 63,
	}
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

func Decode(s *[11]byte, d func() map[byte]uint64) uint64 {
	digitsLookup := d()

	// Decode the input
	var n uint64
	n = (n << 4) | digitsLookup[s[0]]
	for i := 1; i < 11; i++ {
		n = (n << 6) | digitsLookup[s[i]]
	}

	return n
}
