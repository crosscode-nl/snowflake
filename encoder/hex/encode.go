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

func Encode(s *[16]byte, n uint64, d func() Digits) {
	var digits = d()
	for i := 0; i < 16; i++ {
		s[i], n = digits[n>>60&0xf], n<<4
	}
}
