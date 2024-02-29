package bwmarring

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"runtime"
	"sync"
	"testing"
)

type data2 struct {
	id snowflake.ID
	gi int
}

func (d data2) String() string {
	return fmt.Sprintf("id=%v, gi=%v", d.id, d.gi)
}

func TestBwmarringSnowflake(t *testing.T) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	maxProcs := runtime.GOMAXPROCS(-1)
	t.Logf("maxProcs=%v\n", maxProcs)

	var wg sync.WaitGroup
	wg.Add(maxProcs)
	ids := make(chan data2, 10000000)
	for i := 0; i < maxProcs; i++ {
		gi := i
		go func() {
			for j := 0; j < 1000000; j++ {
				id := node.Generate()

				ids <- data2{id, gi}
			}
			wg.Done()
		}()
	}

	func() {
		wg.Wait()
		close(ids)
	}()

	uniqueIDs := make(map[snowflake.ID]data2)
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

func BenchmarkBwmarringSnowflake_NextID(b *testing.B) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		b.Errorf("expected no error, got %v", err)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = node.Generate()
	}
}
