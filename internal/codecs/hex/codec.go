package hex

type Digits [16]byte

func Upper() Digits {
	return [16]byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F',
	}
}

func Lower() Digits {
	return [16]byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f',
	}
}

func UpperLookup() map[byte]uint64 {
	return map[byte]uint64{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7,
		'8': 8, '9': 9, 'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15,
	}
}

func LowerLookup() map[byte]uint64 {
	return map[byte]uint64{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7,
		'8': 8, '9': 9, 'a': 10, 'b': 11, 'c': 12, 'd': 13, 'e': 14, 'f': 15,
	}
}

func Encode(s *[16]byte, n uint64, d func() Digits) {
	var digits = d()
	for i := 0; i < 16; i++ {
		s[i], n = digits[n>>60&0xf], n<<4
	}
}

func Decode(s *[16]byte, d func() map[byte]uint64) uint64 {
	digitsLookup := d()

	// Decode the input
	var n uint64
	for i := 0; i < 16; i++ {
		n = (n << 4) | uint64(digitsLookup[s[i]])
	}

	return n
}
