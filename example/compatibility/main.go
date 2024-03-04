package main

import (
	"fmt"
	sf1 "github.com/bwmarrin/snowflake"
	sf "github.com/crosscode-nl/snowflake"
	sf2 "github.com/godruoyi/go-snowflake"
	sf3 "github.com/influxdata/influxdb/pkg/snowflake"
	"time"
)

func main() {
	fmt.Println("=== bwmarrin/snowflake ===")
	bwmarrin_snoflake()
	fmt.Println("\n=== godruoyi/snowflake ===")
	godruoyi_snowflake()
	fmt.Println("\n=== influx/snowflake ===")
	influx_snowflake()
}

// bwmarrin_snoflake is a compatibility test for bwmarrin/snowflake
func bwmarrin_snoflake() {
	node, err := sf1.NewNode(1) // bwmarrin new
	if err != nil {
		panic(err)
	}

	g, err := sf.NewGenerator(1, sf.WithEpoch(time.UnixMilli(1288834974657))) // snowflake new
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		tid := node.Generate()          // bwmarrin generate id
		mid, _ := g.BlockingNextID(nil) // snowflake generate id
		fmt.Printf("my    id: %v\ntheir id: %v\nmy dec   : %v\ntheir dev: %v\n\n", int64(mid), tid, g.DecodeID(mid), g.DecodeID(sf.ID(tid)))
	}
}

// godruoyi_snoflake is a compatibility test for godruoyi/snowflake
func godruoyi_snowflake() {
	sf2.SetMachineID(1) // godruoyi set machine id

	g, err := sf.NewGenerator(1, sf.WithEpoch(time.Date(2008, 11, 10, 23, 0, 0, 0, time.UTC))) // snowflake new
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		tid, _ := sf2.NextID()          // godruoyi generate id
		mid, _ := g.BlockingNextID(nil) // snowflake generate id
		fmt.Printf("my    id: %v\ntheir id: %v\nmy dec   : %v\ntheir dev: %v\n\n", int64(mid), tid, g.DecodeID(mid), g.DecodeID(sf.ID(tid)))
	}
}

// influx_snowflake is a compatibility test for godruoyi/snowflake
// a small deviation is expected when comparing the two snowflake implementations, because we limit the drift to
// a configurable amount of time, while influx does not.
func influx_snowflake() {
	n := sf3.New(1) // influx new

	g, err := sf.NewGenerator(1, sf.WithDrift(1*time.Second), sf.WithEpoch(time.UnixMilli(1491696000000))) // snowflake new
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		tid := n.Next()      // influx generate id
		mid, _ := g.NextID() // snowflake generate id

		fmt.Printf("my    id: %v\ntheir id: %v\nmy dec   : %v\ntheir dev: %v\n\n", int64(mid), tid, g.DecodeID(mid), g.DecodeID(sf.ID(tid)))
	}
}
