package influx

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
				'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
				'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
				'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
				'U', 'V', 'W', 'X', 'Y', 'Z', '_', 'a', 'b', 'c',
				'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
				'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w',
				'x', 'y', 'z', '~',
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

func TestAlphabetLookup(t *testing.T) {
	tests := []struct {
		name     string
		expected map[byte]uint64
	}{
		{
			name: "AlphabetLookup",
			expected: map[byte]uint64{
				'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
				'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19,
				'K': 20, 'L': 21, 'M': 22, 'N': 23, 'O': 24, 'P': 25, 'Q': 26, 'R': 27, 'S': 28, 'T': 29,
				'U': 30, 'V': 31, 'W': 32, 'X': 33, 'Y': 34, 'Z': 35, '_': 36, 'a': 37, 'b': 38, 'c': 39,
				'd': 40, 'e': 41, 'f': 42, 'g': 43, 'h': 44, 'i': 45, 'j': 46, 'k': 47, 'l': 48, 'm': 49,
				'n': 50, 'o': 51, 'p': 52, 'q': 53, 'r': 54, 's': 55, 't': 56, 'u': 57, 'v': 58, 'w': 59,
				'x': 60, 'y': 61, 'z': 62, '~': 63,
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
			expected: [11]byte{'1', '8', 'p', 'L', 'c', 'Y', 'Q', 'k', 'D', 'w', 'l'},
			alphabet: Alphabet,
		},
		{
			name:     "Input: 0x0",
			input:    0x0,
			expected: [11]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
			alphabet: Alphabet,
		},
		{
			name:     "Input: 0xFFFFFFFFFFFFFFFF",
			input:    0xFFFFFFFFFFFFFFFF,
			expected: [11]byte{'F', '~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
			alphabet: Alphabet,
		},
		{
			name:     "Input: 0xFFFBDFFFAFFFFFFF",
			input:    0xFFFBDFFFAFFFFFFF,
			expected: [11]byte{'F', '~', 'w', 's', '~', 'z', 'k', '~', '~', '~', '~'},
			alphabet: Alphabet,
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
	}{
		{
			name:     "Input: 0x123456789ABCDEF0",
			input:    [11]byte{'1', '8', 'p', 'L', 'c', 'Y', 'Q', 'k', 'D', 'w', 'l'},
			expected: 0x123456789ABCDEF0,
		},
		{
			name:     "Input: 0x0",
			input:    [11]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
			expected: 0x0,
		},
		{
			name:     "Input: 0xFFFFFFFFFFFFFFFF",
			input:    [11]byte{'F', '~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
			expected: 0xFFFFFFFFFFFFFFFF,
		},
		{
			name:     "Input: 0xFFFBDFFFAFFFFFFF",
			input:    [11]byte{'F', '~', 'w', 's', '~', 'z', 'k', '~', '~', '~', '~'},
			expected: 0xFFFBDFFFAFFFFFFF,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Decode(&tt.input, AlphabetLookup)
			if result != tt.expected {
				t.Errorf("Decode() = %v, want %v", result, tt.expected)
			}
		})
	}
}
