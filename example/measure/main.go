package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"time"
)

func main() {
	loopSize := 1000
	data := make(stats.Float64Data, 0, loopSize)

	for i := 0; i < loopSize; i++ {
		nano := time.Duration(time.Now().UnixNano())
		milli := nano.Truncate(time.Millisecond)
		expectedMilli := milli + time.Millisecond
		delay := expectedMilli - nano
		time.Sleep(delay * time.Nanosecond)
		nano = time.Duration(time.Now().UnixNano())
		difference := nano - expectedMilli
		dt := float64(difference) / 1000 / 1000
		data = append(data, dt)
	}
	percentiles := []float64{0.1, 0.4, 0.5, 0.75, 0.9, 0.95, 0.99}
	d, err := stats.Describe(data, true, &percentiles)
	if err != nil {
		panic(err)
	}
	fmt.Printf(d.String(3))
	total, err := data.Sum()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\naverage: %v\n", float64(total)/float64(len(data)))
}
