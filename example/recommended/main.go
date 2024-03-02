package main

import (
	"fmt"
	"github.com/crosscode-nl/snowflake"
)

func main() {
	g, e := snowflake.NewGenerator(1, snowflake.WithDrift())
	if e != nil {
		panic(e)
	}

	// Generate a new ID
	// This will generate a new ID based on the current time, machine ID and sequence number
	// The ID will be returned as an uint64
	// The error will be nil if the ID was generated successfully
	// If the sequence number overflows, the error will be ErrSequenceOverflow
	// This means that the sequence number has reached the maximum value, and we need to wait for the next millisecond
	// If the sequence number overflows and drift is enabled, the generator will continue to generate IDs for times in the future
	// If you don't want to wait for the next millisecond yourself, you can use the BlockingNextID method
	id, _ := g.NextID() // the error is ignored because we know the ID will be generated successfully when drift is enabled
	fmt.Printf("uint64: %v\nstring: %v\ndecoded:%v\n", uint64(id), id, g.DecodeID(id))
	// Output:
	// uint64: 672572626702336
	// string: 002OwE4W100
	// decoded:ID: 672572626702336, Timestamp: 160353810, MachineID: 1, Sequence: 0
}
