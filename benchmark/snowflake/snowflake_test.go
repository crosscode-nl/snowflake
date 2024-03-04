package snowflake

import (
	"fmt"
	"github.com/crosscode-nl/snowflake"
	"runtime"
	"sync"
	"testing"
	"time"
)

type data struct {
	id snowflake.ID
	gi int
}

func (d data) String() string {
	return fmt.Sprintf("id=%v, gi=%v", d.id, d.gi)
}

// TestGenerator_BlockingNextID_Concurrent_No_Duplicates tests the BlockingNextID method of the Generator to ensure it generates unique IDs in a concurrent environment
func TestGenerator_BlockingNextID_Concurrent_No_Duplicates(t *testing.T) {
	maxProcs := runtime.GOMAXPROCS(-1)
	t.Logf("maxProcs=%v\n", maxProcs)
	generator, err := snowflake.NewGenerator(378)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(maxProcs)

	ids := make(chan data, 10000000)
	for i := 0; i < maxProcs; i++ {
		gi := i
		go func() {
			for j := 0; j < 1000000; j++ {
				id, err := generator.BlockingNextID(nil)
				if err != nil {
					panic(err)
				}
				ids <- data{id, gi}
			}
			wg.Done()
		}()
	}

	func() {
		wg.Wait()
		close(ids)
	}()

	uniqueIDs := make(map[snowflake.ID]data)
	for id := range ids {
		if oid, ok := uniqueIDs[id.id]; ok {
			if oid.gi == id.gi {
				t.Errorf(">> expected unique ids, got duplicate %v: %v, original: %v: %v <<", generator.DecodeID(id.id), id.gi, generator.DecodeID(oid.id), oid.gi)
			}
			t.Errorf("expected unique ids, got duplicate %v: %v, original: %v: %v", generator.DecodeID(id.id), id.gi, generator.DecodeID(oid.id), oid.gi)
		}
		uniqueIDs[id.id] = id
	}

}

func BenchmarkGenerator_NextID(b *testing.B) {
	generator, err := snowflake.NewGenerator(378, snowflake.WithDriftNoWait(1*time.Hour))
	if err != nil {
		b.Errorf("expected no error, got %v", err)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = generator.BlockingNextID(nil)
	}
}

func BenchmarkGenerator_NextID_NoDrift(b *testing.B) {
	generator, err := snowflake.NewGenerator(378)
	if err != nil {
		b.Errorf("expected no error, got %v", err)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = generator.BlockingNextID(nil)
	}
}

func BenchmarkGenerator_NextID_NoDrift_9b(b *testing.B) {
	for m := 1; m < 22; m++ {
		b.Run(fmt.Sprintf("MachineIDBits=%d", m), func(b *testing.B) {
			generator, err := snowflake.NewGenerator(0, snowflake.WithMachineIDBits(uint64(m)))
			if err != nil {
				b.Errorf("expected no error, got %v", err)
				return
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = generator.BlockingNextID(nil)
			}
		})
	}
}
