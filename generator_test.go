package snowflake

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

// TestGenerator_NextID tests the NextID method of the Generator
// It uses a test vector based on the first Tweet on Twitter
func TestGenerator_NextID(t *testing.T) {
	generator, err := NewGenerator(378, WithEpoch(time.UnixMilli(0)))
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
	generator, err := NewGenerator(0, WithEpoch(time.UnixMilli(0)))
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
			generator, err := NewGenerator(0, WithMachineIDBits(machineIDBits), WithEpoch(time.UnixMilli(0)))
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
func TestGenerator_BlockingNextID_ErrorWhenContextIsCanceledAndBlockingWouldOccurr(t *testing.T) {
	generator, err := NewGenerator(378)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	generator.timeFunc = func() uint64 {
		return uint64(generator.epoch)
	}
	maxCount := 1 << 12
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	for i := 1; i <= maxCount; i++ {
		_, err = generator.BlockingNextID(ctx)
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context canceled, got %v", err)
		return
	}
}

// TestGenerator_BlockingNextID tests the BlockingNextID method of the Generator
func TestGenerator_BlockingNextID_BlockedUntilNextId(t *testing.T) {
	generator, err := NewGenerator(378, WithEpoch(time.UnixMilli(0)))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	var ts uint64
	generator.timeFunc = func() uint64 {
		return ts
	}
	go func() {
		time.Sleep(100 * time.Millisecond)
		ts = 1
	}()
	maxCount := 1 << 12
	var id ID
	var previousID ID
	for i := 0; i < maxCount; i++ {
		id, err = generator.BlockingNextID(context.TODO())
	}

	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	dId := generator.DecodeID(id)
	dPreviousId := generator.DecodeID(previousID)
	if dId.Timestamp <= dPreviousId.Timestamp {
		t.Errorf("expected id to be greater than previous id %v, got %v", dPreviousId.Timestamp, dId.Timestamp)
	}
	if dId.Sequence != 0 {
		t.Errorf("expected sequence to be 0, got %v", dId.Sequence)
	}
}

// TestGenerator_BlockingNextID tests the BlockingNextID with cancel context
func TestGenerator_BlockingNextID(t *testing.T) {
	generator, err := NewGenerator(378, WithEpoch(time.UnixMilli(0)))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	generator.timeFunc = func() uint64 {
		return 367597485448
	}

	id, err := generator.BlockingNextID(context.TODO())
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
	generator, err := NewGenerator(378, WithEpoch(time.UnixMilli(0)))

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
	for id, err = generator.BlockingNextID(context.TODO()); blocked == false; id, err = generator.BlockingNextID(nil) {
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

// TestNewGenerator_Errors tests the NewGenerator function for errors
func TestNewGenerator_Errors(t *testing.T) {
	tests := []struct {
		name        string
		machineID   uint64
		machineBits uint64
		want        error
	}{
		{
			name:        "Test NewGenerator with machine bits too small",
			machineID:   1,
			machineBits: 0,
			want:        ErrMachineBitsTooSmall,
		},
		{
			name:        "Test NewGenerator with machine bits too large",
			machineID:   1,
			machineBits: 22,
			want:        ErrMachineBitsTooLarge,
		},
		{
			name:        "Test NewGenerator with machine ID too large",
			machineID:   1 << 10,
			machineBits: 10,
			want:        ErrMachineIDTooLarge,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewGenerator(tt.machineID, WithMachineIDBits(tt.machineBits))
			if !errors.Is(err, tt.want) {
				t.Errorf("expected %v, got %v", tt.want, err)
			}
		})
	}
}

// TestDefaultTimeFunc tests the defaultTimeFunc function
func TestDefaultTimeFunc(t *testing.T) {
	now := defaultTimeFunc()
	if now == 0 {
		t.Errorf("expected non zero time, got %v", now)
	}
	time.Sleep(2 * time.Millisecond)
	now2 := defaultTimeFunc()
	if now2 <= now {
		t.Errorf("expected time to have advanced, got %v", now2)
	}
}

func TestWithExactSleep(t *testing.T) {
	generator, err := NewGenerator(378, WithEpoch(time.UnixMilli(0)), WithExactSleep())
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	// The default sleeping implementation is not exact, it deviates thousands of nanoseconds
	// from the expected time. WithExactSleep should perform much better and deviate less than 1us
	for i := 0; i < 10; i++ {
		generator.sleepFunc()
		now := time.Now().UnixNano()
		now2 := time.Now().UnixMilli()
		difference := now - now2*1000*1000
		if difference > 1000 {
			t.Errorf("expected difference to be less than 1us, got %v", difference)
		}
	}
}

func TestWithDrift(t *testing.T) {
	now := time.Now()
	generator, err := NewGenerator(378, WithEpoch(time.UnixMilli(0)), WithDrift(200*time.Millisecond))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	if !generator.drift {
		t.Errorf("expected drift to be true, got %v", generator.drift)
	}
	if generator.duration != 200*time.Millisecond {
		t.Errorf("expected duration to be 200ms, got %v", generator.duration)
	}
	if time.Now().Sub(now) < 200*time.Millisecond {
		t.Errorf("expected time to have advanced by at least 200ms, got %v", time.Now().Sub(now))
	}
}

func TestWithDriftNoWait(t *testing.T) {
	now := time.Now()
	generator, err := NewGenerator(378, WithEpoch(time.UnixMilli(0)), WithDriftNoWait(200*time.Millisecond))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	if !generator.drift {
		t.Errorf("expected drift to be true, got %v", generator.drift)
	}
	if generator.duration != 200*time.Millisecond {
		t.Errorf("expected duration to be 200ms, got %v", generator.duration)
	}
	if time.Now().Sub(now) > time.Millisecond {
		t.Errorf("expected time to have advanced by at least less than a millisecond, got %v", time.Now().Sub(now))
	}
}

func TestGenerator_NextID_InvalidEpoch(t *testing.T) {
	generator, err := NewGenerator(378, WithEpoch(time.Now().Add(time.Hour)))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}
	id, err := generator.NextID()
	if !errors.Is(err, ErrTimeBeforeEpoch) {
		t.Errorf("expected ErrTimeBeforeEpoch, got %v", err)
	}
	if id != 0 {
		t.Errorf("expected 0, got %v", id)
	}
}

func TestNewGenerator_NextId_WithDrif2ms_Returns3x(t *testing.T) {
	generator, err := NewGenerator(0, WithEpoch(time.UnixMilli(0)), WithDrift(2*time.Millisecond))
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
	maxCount := (generator.sequenceMask + 1) * 3
	if count != maxCount {
		t.Errorf("expected %v ids, got %v", maxCount, count)
	}
}
