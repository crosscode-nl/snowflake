package base64

import (
	"reflect"
	"testing"
)

func TestAlphabet(t *testing.T) {
	tests := []struct {
		name     string
		expected [64]byte
	}{
		{
			name: "Alphabet",
			expected: [64]byte{
				'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
				'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
				'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
				'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
				'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
				'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
				'w', 'x', 'y', 'z', '0', '1', '2', '3',
				'4', '5', '6', '7', '8', '9', '+', '/',
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Alphabet()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Alphabet() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestUrlAlphabet(t *testing.T) {
	tests := []struct {
		name     string
		expected [64]byte
	}{
		{
			name: "UrlAlphabet",
			expected: [64]byte{
				'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
				'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
				'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
				'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
				'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
				'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
				'w', 'x', 'y', 'z', '0', '1', '2', '3',
				'4', '5', '6', '7', '8', '9', '-', '_',
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UrlAlphabet()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("UrlAlphabet() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMimeAlphabet(t *testing.T) {
	tests := []struct {
		name     string
		expected [64]byte
	}{
		{
			name: "MimeAlphabet",
			expected: [64]byte{
				'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
				'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
				'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
				'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
				'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
				'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
				'w', 'x', 'y', 'z', '0', '1', '2', '3',
				'4', '5', '6', '7', '8', '9', '+', ',',
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MimeAlphabet()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MimeAlphabet() = %v, want %v", result, tt.expected)
			}
		})
	}

}

func TestAlphabetLookup(t *testing.T) {
	tests := []struct {
		name     string
		expected map[byte]uint64
	}{
		{
			name: "AlphabetLookup",
			expected: map[byte]uint64{
				'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7,
				'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15,
				'Q': 16, 'R': 17, 'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23,
				'Y': 24, 'Z': 25, 'a': 26, 'b': 27, 'c': 28, 'd': 29, 'e': 30, 'f': 31,
				'g': 32, 'h': 33, 'i': 34, 'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39,
				'o': 40, 'p': 41, 'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47,
				'w': 48, 'x': 49, 'y': 50, 'z': 51, '0': 52, '1': 53, '2': 54, '3': 55,
				'4': 56, '5': 57, '6': 58, '7': 59, '8': 60, '9': 61, '+': 62, '/': 63,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AlphabetLookup()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("AlphabetLookup() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestUrlAlphabetLookup(t *testing.T) {
	tests := []struct {
		name     string
		expected map[byte]uint64
	}{
		{
			name: "UrlAlphabetLookup",
			expected: map[byte]uint64{
				'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7,
				'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15,
				'Q': 16, 'R': 17, 'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23,
				'Y': 24, 'Z': 25, 'a': 26, 'b': 27, 'c': 28, 'd': 29, 'e': 30, 'f': 31,
				'g': 32, 'h': 33, 'i': 34, 'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39,
				'o': 40, 'p': 41, 'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47,
				'w': 48, 'x': 49, 'y': 50, 'z': 51, '0': 52, '1': 53, '2': 54, '3': 55,
				'4': 56, '5': 57, '6': 58, '7': 59, '8': 60, '9': 61, '-': 62, '_': 63,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UrlAlphabetLookup()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("UrlAlphabetLookup() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMimeAlphabetLookup(t *testing.T) {
	tests := []struct {
		name     string
		expected map[byte]uint64
	}{
		{
			name: "MimeAlphabetLookup",
			expected: map[byte]uint64{
				'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7,
				'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15,
				'Q': 16, 'R': 17, 'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23,
				'Y': 24, 'Z': 25, 'a': 26, 'b': 27, 'c': 28, 'd': 29, 'e': 30, 'f': 31,
				'g': 32, 'h': 33, 'i': 34, 'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39,
				'o': 40, 'p': 41, 'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47,
				'w': 48, 'x': 49, 'y': 50, 'z': 51, '0': 52, '1': 53, '2': 54, '3': 55,
				'4': 56, '5': 57, '6': 58, '7': 59, '8': 60, '9': 61, '+': 62, ',': 63,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MimeAlphabetLookup()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MimeAlphabetLookup() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAlphabetsOverlap(t *testing.T) {
	normal := Alphabet()
	url := UrlAlphabet()
	mime := MimeAlphabet()

	if !reflect.DeepEqual(normal[0:64-2], url[0:64-2]) {
		t.Errorf("normal and url alphabet mismatch, %v != %v", normal[0:64-2], url[0:64-2])
	}
	if !reflect.DeepEqual(normal[0:64-2], mime[0:64-2]) {
		t.Errorf("normal and mime alphabet mismatch, %v != %v", normal[0:64-2], mime[0:64-2])
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected [11]byte
		alphabet func() [64]byte
	}{
		{
			name:     "Input: 0x123456789ABCDEF0",
			input:    0x123456789ABCDEF0,
			expected: [11]byte{'8', 'N', '6', '8', 'm', 'n', 'h', 'W', 'N', 'B', 'I'},
			alphabet: Alphabet,
		},
		{
			name:     "Input: 0x0",
			input:    0x0,
			expected: [11]byte{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'},
			alphabet: Alphabet,
		},
		{
			name:     "Input: 0xFFFFFFFFFFFFFFFF",
			input:    0xFFFFFFFFFFFFFFFF,
			expected: [11]byte{'/', '/', '/', '/', '/', '/', '/', '/', '/', '/', '8'},
			alphabet: Alphabet,
		},
		{
			name:     "Input: 0xFFFBDFFFAFFFFFFF normal",
			input:    0xFFFBDFFFAFFFFFFF,
			expected: [11]byte{'/', '/', '/', '/', 'r', '/', '/', 'f', '+', '/', '8'},
			alphabet: Alphabet,
		},
		{
			name:     "Input: 0xFFFBDFFFAFFFFFFF url",
			input:    0xFFFBDFFFAFFFFFFF,
			expected: [11]byte{'_', '_', '_', '_', 'r', '_', '_', 'f', '-', '_', '8'},
			alphabet: UrlAlphabet,
		},
		{
			name:     "Input: 0xFFFBDFFFAFFFFFFF mime",
			input:    0xFFFBDFFFAFFFFFFF,
			expected: [11]byte{',', ',', ',', ',', 'r', ',', ',', 'f', '+', ',', '8'},
			alphabet: MimeAlphabet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s [11]byte
			Encode(&s, tt.input, tt.alphabet)
			if !reflect.DeepEqual(s, tt.expected) {
				t.Errorf("Encode() = %v, want %v", s, tt.expected)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		input    [11]byte
		expected uint64
		alphabet func() map[byte]uint64
	}{
		{
			name:     "Input: 8N68mnhWNBI",
			input:    [11]byte{'8', 'N', '6', '8', 'm', 'n', 'h', 'W', 'N', 'B', 'I'},
			expected: 0x123456789ABCDEF0,
			alphabet: AlphabetLookup,
		},
		{
			name:     "Input: AAAAAAAAAAA",
			input:    [11]byte{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'},
			expected: 0x0,
			alphabet: AlphabetLookup,
		},
		{
			name:     "Input: ///////////8",
			input:    [11]byte{'/', '/', '/', '/', '/', '/', '/', '/', '/', '/', '8'},
			expected: 0xFFFFFFFFFFFFFFFF,
			alphabet: AlphabetLookup,
		},
		{
			name:     "Input: ////r//f+/8",
			input:    [11]byte{'/', '/', '/', '/', 'r', '/', '/', 'f', '+', '/', '8'},
			expected: 0xFFFBDFFFAFFFFFFF,
			alphabet: AlphabetLookup,
		},
		{
			name:     "Input: ____r__f-_8",
			input:    [11]byte{'_', '_', '_', '_', 'r', '_', '_', 'f', '-', '_', '8'},
			expected: 0xFFFBDFFFAFFFFFFF,
			alphabet: UrlAlphabetLookup,
		},
		{
			name:     "Input: ,,,,,r,,f+,8",
			input:    [11]byte{',', ',', ',', ',', 'r', ',', ',', 'f', '+', ',', '8'},
			expected: 0xFFFBDFFFAFFFFFFF,
			alphabet: MimeAlphabetLookup,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Decode(&tt.input, tt.alphabet)
			if result != tt.expected {
				t.Errorf("Decode() = %v, want %v", result, tt.expected)
			}
		})
	}
}
