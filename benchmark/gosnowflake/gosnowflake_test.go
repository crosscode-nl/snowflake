package gosnowflake

import (
	"fmt"
	"github.com/godruoyi/go-snowflake"
	"runtime"
	"sync"
	"testing"
)

type data struct {
	id uint64
	gi int
}

func (d data) String() string {
	return fmt.Sprintf("id=%v, gi=%v", d.id, d.gi)
}

func TestGoSnowflake(t *testing.T) {

	maxProcs := runtime.GOMAXPROCS(-1)
	t.Logf("maxProcs=%v\n", maxProcs)

	var wg sync.WaitGroup
	wg.Add(maxProcs)
	ids := make(chan data, 10000000)
	for i := 0; i < maxProcs; i++ {
		gi := i
		go func() {
			for j := 0; j < 1000000; j++ {
				id, err := snowflake.NextID()
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

	uniqueIDs := make(map[uint64]data)
	for id := range ids {
		if oid, ok := uniqueIDs[id.id]; ok {
			if oid.gi == id.gi {
				t.Errorf(">> expected unique ids, got duplicate %v: %v, original: %v: %v <<", id.id, id.gi, id.id, oid.gi)
			}
			t.Errorf("expected unique ids, got duplicate %v: %v, original: %v: %v", id.id, id.gi, oid.id, oid.gi)
		}
		uniqueIDs[id.id] = id
	}

}

func BenchmarkGoSnowflake_NextID(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = snowflake.NextID()
	}
}
