package hex

import (
	"reflect"
	"testing"
)

func digitsToString(digits Digits) string {
	d := [16]byte(digits)
	return string(d[:])
}

func TestDigits(t *testing.T) {
	upper := digitsToString(Upper())
	lower := digitsToString(Lower())
	if lower != "0123456789abcdef" {
		t.Errorf("Lower() = %v, want %v", lower, "0123456789abcdef")
	}
	if upper != "0123456789ABCDEF" {
		t.Errorf("Upper() = %v, want %v", upper, "0123456789ABCDEF")
	}

}

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected [16]byte
		digits   func() Digits
	}{
		{
			name:     "Input: 0x123456789abcdef0",
			input:    0x123456789ABCDEF0,
			expected: [16]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', '0'},
			digits:   Lower,
		},
		{
			name:     "Input: 0x123456789ABCDEF0",
			input:    0x123456789ABCDEF0,
			expected: [16]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', '0'},
			digits:   Upper,
		},
		{
			name:     "Input: 0x0",
			input:    0x0,
			expected: [16]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
			digits:   Upper,
		},
		{
			name:     "Input: 0xFFFFFFFFFFFFFFFF",
			input:    0xFFFFFFFFFFFFFFFF,
			expected: [16]byte{'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F'},
			digits:   Upper,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s [16]byte
			Encode(&s, tt.input, tt.digits)
			if !reflect.DeepEqual(s, tt.expected) {
				t.Errorf("Encode() = %v, want %v", s, tt.expected)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name         string
		input        [16]byte
		expected     uint64
		digitsLookup func() map[byte]uint64
	}{
		{
			name:         "Input: 0x123456789ABCDEF0",
			input:        [16]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', '0'},
			expected:     0x123456789ABCDEF0,
			digitsLookup: UpperLookup,
		},
		{
			name:         "Input: 0x123456789abcdef0",
			input:        [16]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', '0'},
			expected:     0x123456789ABCDEF0,
			digitsLookup: LowerLookup,
		},
		{
			name:         "Input: 0x0",
			input:        [16]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
			expected:     0x0,
			digitsLookup: UpperLookup,
		},
		{
			name:         "Input: 0xFFFFFFFFFFFFFFFF",
			input:        [16]byte{'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F', 'F'},
			expected:     0xFFFFFFFFFFFFFFFF,
			digitsLookup: UpperLookup,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s [16]byte
			copy(s[:], tt.input[:])
			if got := Decode(&s, tt.digitsLookup); got != tt.expected {
				t.Errorf("Decode() = %v, want %v", got, tt.expected)
			}
		})
	}
}
