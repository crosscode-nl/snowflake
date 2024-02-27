package snowflakes

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// TestGenerator_NextID tests the NextID method of the Generator
// It uses a test vector based on the first Tweet on Twitter
func TestGenerator_NextID(t *testing.T) {
	generator, err := NewGenerator(378)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	generator.timeFunc = func() int64 {
		return 367597485448
	}

	id, err := generator.NextID()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	if id != 1541815603606036480 {
		t.Errorf("expected 1541815603606036480, got %v", id)
	}
}

// TestGenerator_NextID_WithEpoch tests the NextID method of the Generator with a custom epoch
// It uses a test vector based on the first Tweet on Twitter
func TestGenerator_NextID_WithEpoch(t *testing.T) {
	generator, err := NewGenerator(378, WithEpoch(time.UnixMilli(1288834974657)))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	generator.timeFunc = func() int64 {
		return 1656432460105
	}

	id, err := generator.NextID()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	if id != 1541815603606036480 {
		t.Errorf("expected 1541815603606036480, got %v", id)
	}
}

// TestGenerator_NextID_GeneratesCorrectAmount tests the NextID method of the Generator to ensure it generates the correct amount of IDs with the default machine ID bit size
func TestGenerator_NextID_GeneratesCorrectAmount(t *testing.T) {
	generator, err := NewGenerator(378)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	generator.timeFunc = func() int64 {
		return 0
	}

	var previousID ID
	var count int
	for id, err := generator.NextID(); err == nil; id, err = generator.NextID() {
		if previousID > id {
			t.Errorf("expected id to be greater than previous id, got %v", id)
		}
		count++
	}
	maxCount := 1 << 12
	if count != maxCount {
		t.Errorf("expected %v ids, got %v", maxCount, count)
	}
}

// TestGenerator_NextID_GeneratesCorrectAmount_WithMachineIdBits tests the NextID method of the Generator to ensure it generates the correct amount of IDs with different machine ID bit sizes
func TestGenerator_NextID_GeneratesCorrectAmount_WithMachineIdBits(t *testing.T) {
	for machineIdBits := 1; machineIdBits < 22; machineIdBits++ {
		maxCount := 1 << (22 - machineIdBits)
		t.Run(fmt.Sprintf("TestGenerator_NextID_GeneratesCorrectAmount_WithMachineIdBits=%v_Gives_%v_ids", machineIdBits, maxCount), func(t *testing.T) {
			generator, err := NewGenerator(0, WithMachineIdBits(machineIdBits))
			if err != nil {
				t.Errorf("expected no error, got %v", err)
				return
			}
			generator.timeFunc = func() int64 {
				return 0
			}
			var previousID ID
			var count int

			for id, err := generator.NextID(); err == nil; id, err = generator.NextID() {
				if previousID > id {
					t.Errorf("expected id to be greater than previous id, got %v", id)
				}
				previousID = id
				count++
				if count > maxCount {
					break
				}
			}

			if count != maxCount {
				t.Errorf("expected %v ids, got %v", maxCount, count)
			}
		})
	}
}

// TestGenerator_BlockingNextID tests the BlockingNextID method of the Generator
func TestGenerator_BlockingNextID(t *testing.T) {
	generator, err := NewGenerator(378)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	generator.timeFunc = func() int64 {
		return 367597485448
	}

	id, err := generator.BlockingNextID(nil)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	if id != 1541815603606036480 {
		t.Errorf("expected 1541815603606036480, got %v", id)
	}
}

// TestGenerator_BlockingNextID_UntilBlock tests the BlockingNextID method of the Generator to ensure it blocks until
// the next ID can be generated
func TestGenerator_BlockingNextID_UntilBlock(t *testing.T) {
	generator, err := NewGenerator(378)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	var blocked bool
	generator.timeFunc = func() int64 {
		return 367597485447
	}
	generator.sleepFunc = func() {
		blocked = true
		generator.timeFunc = func() int64 {
			return 367597485448
		}
	}

	var previousID ID
	var count int
	maxCount := 1 << (22 - generator.machineIdBits)
	var id ID
	for id, err = generator.BlockingNextID(nil); blocked == false; id, err = generator.BlockingNextID(nil) {
		if err != nil {
			t.Errorf("expected no error, got %v", err)
			return
		}
		if previousID > id {
			t.Errorf("expected id to be greater than previous id, got %v", id)
		}
		previousID = id
		if count > maxCount {
			break
		}
		count++
	}

	if count != maxCount {
		t.Errorf("expected %v ids, got %v", maxCount, count)
	}

	if id != 1541815603606036480 {
		t.Errorf("expected 1541815603606036480, got %v", id)
	}
}

type data struct {
	id ID
	gi int
}

func (d data) String() string {
	return fmt.Sprintf("id=%v, gi=%v", d.id, d.gi)
}

// TestGenerator_BlockingNextID_Concurrent_No_Duplicates tests the BlockingNextID method of the Generator to ensure it generates unique IDs in a concurrent environment
func TestGenerator_BlockingNextID_Concurrent_No_Duplicates(t *testing.T) {
	maxProcs := runtime.GOMAXPROCS(-1)
	t.Logf("maxProcs=%v\n", maxProcs)
	generator, err := NewGenerator(378)
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
				var id ID
				for {
					var err error
					id, err = generator.NextID()
					if err == nil {
						break
					}
					time.Sleep(10 * time.Microsecond)
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

	uniqueIDs := make(map[ID]data)
	for id := range ids {
		if oid, ok := uniqueIDs[id.id]; ok {
			if oid.gi == id.gi {
				t.Errorf(">> expected unique ids, got duplicate %v, original: %v", id, oid)
			}
			t.Errorf("expected unique ids, got duplicate %v, original: %v", id, oid)
		}
		uniqueIDs[id.id] = id
	}

}
