package snowflake

import "testing"

// TestDecodedID_String tests the DecodedID String method
func TestDecodedID_String(t *testing.T) {
	tests := []struct {
		name string
		id   DecodedID
		want string
	}{
		{
			name: "Test DecodedID String method",
			id: DecodedID{
				ID:        1,
				Timestamp: 2,
				MachineID: 3,
				Sequence:  4,
			},
			want: "ID: 1, Timestamp: 2, MachineID: 3, Sequence: 4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.id.String(); got != tt.want {
				t.Errorf("DecodedID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGenerator_DecodeID tests the Generator DecodeID method
func TestGenerator_DecodeID(t *testing.T) {
	tests := []struct {
		name             string
		machineID        uint64
		machineIDBitSize uint64
		timestamp        uint64
		sequence         int
		want             DecodedID
	}{
		{
			name:             "Test Generator DecodeID method",
			machineID:        2,
			machineIDBitSize: 10,
			timestamp:        0,
			sequence:         1,
			want: DecodedID{
				ID:        8193,
				Timestamp: 0,
				MachineID: 2,
				Sequence:  1,
			},
		},
		{
			name:             "Test Generator DecodeID method with drift",
			machineID:        2,
			machineIDBitSize: 10,
			timestamp:        0,
			sequence:         1 << 12,
			want: DecodedID{
				ID:        4202496,
				Timestamp: 1,
				MachineID: 2,
				Sequence:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := NewGenerator(tt.machineID, WithMachineIDBits(tt.machineIDBitSize), WithDrift())
			if err != nil {
				t.Errorf("expected no error, got %v", err)
				return
			}
			g.timeFunc = func() uint64 {
				return tt.timestamp
			}
			var id ID
			for i := 0; i < tt.sequence; i++ {
				id, err = g.NextID()
				if err != nil {
					t.Errorf("expected no error, got %v", err)
					return
				}
			}
			if got := g.DecodeID(id); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
