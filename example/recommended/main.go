package main

import (
	"context"
	"fmt"
	"github.com/crosscode-nl/snowflake"
	"time"
)

func main() {
	// Allow 500 milliseconds drift before block.
	// This also sleeps for 500 milliseconds in this constructor, to prevent ID collisions on a restart of the application.
	// I would not suggest to allow a drift of more than a few seconds, because it would take a long time to start.
	g, e := snowflake.NewGenerator(1, snowflake.WithDrift(500*time.Millisecond))
	if e != nil {
		panic(e)
	}

	// Generate a new ID.
	// This will generate a new ID based on the current time, machine ID and sequence number.
	// The ID will be returned as an uint64.
	// The error will be nil if the ID was generated successfully.
	// If the sequence number overflows we will borrow time from the next milliseconds to continue generating ID.
	// This continues until the maximum amount of drift is reached at which point the BlockingNextID method will block.
	// the error is ignored because we know the ID will be generated successfully when drift is enabled, and we
	// don't provide a cancel context.
	id, _ := g.BlockingNextID(context.TODO())
	fmt.Printf("uint64: %v\nstring: %v\ndecoded:%v\n", uint64(id), id, g.DecodeID(id))
	// Output:
	// uint64: 672572626702336
	// string: 002OwE4W100
	// decoded:ID: 672572626702336, Timestamp: 160353810, MachineID: 1, Sequence: 0
}
