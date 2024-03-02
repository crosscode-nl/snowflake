package snowflake

import "fmt"

// DecodedID is a snowflake ID decoded into its components
type DecodedID struct {
	ID        uint64
	Timestamp uint64
	MachineID uint64
	Sequence  uint64
}

// String returns a string representation of the decoded ID
func (id DecodedID) String() string {
	return fmt.Sprintf("ID: %d, Timestamp: %d, MachineID: %d, Sequence: %d", id.ID, id.Timestamp, id.MachineID, id.Sequence)
}

// DecodeID decodes a snowflake ID into its components
func (g *Generator) DecodeID(id ID) DecodedID {
	return DecodedID{
		ID:        uint64(id),
		Timestamp: uint64(id) >> timeShift,
		MachineID: uint64(id) >> g.machineIDShift & g.machineIDMask,
		Sequence:  uint64(id) & g.sequenceMask,
	}
}
