package snowflake

import (
	"fmt"
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
	generator.timeFunc = func() uint64 {
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

	generator.timeFunc = func() uint64 {
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
	generator, err := NewGenerator(0)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	generator.timeFunc = func() uint64 {
		return 1
	}

	var previousID ID
	var count uint64
	for id, err := generator.NextID(); err == nil; id, err = generator.NextID() {
		if previousID > id {
			t.Errorf("expected id to be greater than previous id, got %v", id)
		}
		count++
	}
	maxCount := generator.sequenceMask + 1
	if count != maxCount {
		t.Errorf("expected %v ids, got %v", maxCount, count)
	}
}

// TestGenerator_NextID_GeneratesCorrectAmount_WithMachineIdBits tests the NextID method of the Generator to ensure it generates the correct amount of IDs with different machine ID bit sizes
func TestGenerator_NextID_GeneratesCorrectAmount_WithMachineIdBits(t *testing.T) {
	for machineIDBits := uint64(1); machineIDBits < 22; machineIDBits++ {
		maxCount := 1 << (22 - machineIDBits)
		t.Run(fmt.Sprintf("TestGenerator_NextID_GeneratesCorrectAmount_WithMachineIdBits=%v_Gives_%v_ids", machineIDBits, maxCount), func(t *testing.T) {
			generator, err := NewGenerator(0, WithMachineIdBits(machineIDBits))
			if err != nil {
				t.Errorf("expected no error, got %v", err)
				return
			}
			generator.timeFunc = func() uint64 {
				return 1
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
	generator.timeFunc = func() uint64 {
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
	generator.timeFunc = func() uint64 {
		return 367597485447
	}
	generator.sleepFunc = func() {
		blocked = true
		generator.timeFunc = func() uint64 {
			return 367597485448
		}
	}

	var previousID ID
	var count uint64
	maxCount := generator.sequenceMask + 1
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
