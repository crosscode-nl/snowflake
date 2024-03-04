package snowflake

import (
	"context"
	"errors"
	"sync/atomic"
	"time"
)

var (
	// ErrMachineIDTooLarge is returned when the machine ID is too large for the number of bits
	ErrMachineIDTooLarge = errors.New("machine ID is too large")
	// ErrMachineBitsTooSmall is returned when the number of bits for the machine ID is too small
	ErrMachineBitsTooSmall = errors.New("machine ID bits is too small")
	// ErrMachineBitsTooLarge is returned when the number of bits for the machine ID is too large
	ErrMachineBitsTooLarge = errors.New("machine ID bits is too large")
	// ErrOutOfSequence is returned when the sequence number overflows
	ErrOutOfSequence = errors.New("sequence number overflow")
	// ErrTimeBeforeEpoch is returned when the time is before the epoch
	ErrTimeBeforeEpoch = errors.New("time is before epoch")
)

const (
	timeShift = 22
)

// Option is a function that configures the generator
type Option func(*Generator)

// TimeFunc is a function that returns the current time in milliseconds
type TimeFunc func() uint64

func defaultTimeFunc() uint64 {
	return uint64(time.Now().UnixMilli())
}

// Generator is a snowflake ID generator
type Generator struct {
	currentID      atomic.Uint64
	machineID      uint64
	sequenceMask   uint64
	machineIDMask  uint64
	machineIDBits  uint64
	machineIDShift uint64
	epoch          int64
	timeFunc       TimeFunc
	sleepFunc      func()
	drift          bool
	duration       time.Duration
}

// NewGenerator creates a new snowflake ID generator
// machineID is the unique ID of the machine running the generator
// opts are the options to configure the generator
// Returns a new snowflake ID generator
// Returns an error if the machineID is too large for the number of bits
// Returns an error if the machineIDBits is invalid
func NewGenerator(machineID uint64, opts ...Option) (*Generator, error) {
	g := &Generator{
		timeFunc:      defaultTimeFunc,
		machineIDBits: 10,
		machineID:     machineID,
		sleepFunc: func() {
			nano := time.Duration(time.Now().UnixNano())
			milli := nano.Truncate(time.Millisecond)
			milli = milli + time.Millisecond
			delay := milli - nano
			time.Sleep(delay + 1*time.Nanosecond)
		},
		epoch: 1709247600000,
	}

	for _, opt := range opts {
		opt(g)
	}

	maxMachineID := uint64(1<<g.machineIDBits - 1)

	if g.machineIDBits < 1 {
		return nil, ErrMachineBitsTooSmall
	}

	if g.machineIDBits > 21 {
		return nil, ErrMachineBitsTooLarge
	}

	if g.machineID > maxMachineID {
		return nil, ErrMachineIDTooLarge
	}

	g.machineIDMask = maxMachineID
	g.sequenceMask = 1<<(timeShift-g.machineIDBits) - 1
	g.machineIDShift = timeShift - g.machineIDBits

	return g, nil
}

// NextID generates a new snowflake ID
func (g *Generator) NextID() (ID, error) {

	now := int64(g.timeFunc()) - g.epoch

	if now < 0 {
		return 0, ErrTimeBeforeEpoch
	}

	for {
		currentID := g.currentID.Load()
		newCurrentID := currentID
		lastTime := currentID >> timeShift
		sequence := currentID & g.sequenceMask
		switch {
		case lastTime < uint64(now):
			lastTime = uint64(now)
			newCurrentID = lastTime << timeShift
		case sequence == g.sequenceMask:
			if !g.drift {
				return 0, ErrOutOfSequence
			}
			if uint64(now)-lastTime >= uint64(g.duration.Milliseconds()) {
				return 0, ErrOutOfSequence
			}
			newCurrentID = (lastTime + 1) << timeShift
		default:
			newCurrentID++
		}
		newCurrentID = newCurrentID | (g.machineID << g.machineIDShift)
		if g.currentID.CompareAndSwap(currentID, newCurrentID) {
			return ID(newCurrentID), nil
		}
	}
}

// BlockingNextID generates a new snowflake ID, blocking until the next ID can be generated
func (g *Generator) BlockingNextID(ctx context.Context) (ID, error) {
	id, err := g.NextID()
	for errors.Is(err, ErrOutOfSequence) {
		if ctx != nil && ctx.Err() != nil {
			return 0, ctx.Err()
		}
		g.sleepFunc()
		id, err = g.NextID()
	}
	return id, nil
}

// WithMachineIDBits sets the number of bits to use for the machine ID
func WithMachineIDBits(size uint64) Option {
	return func(generator *Generator) {
		generator.machineIDBits = size
	}
}

// WithEpoch sets the epoch for the generator
func WithEpoch(epoch time.Time) Option {
	return func(generator *Generator) {
		generator.epoch = epoch.UnixMilli()
	}
}

// WithDrift enables drift to continue generating IDs when the sequence overflows
// This allows the generator to generate IDs for times in the future
// This increases performance but may generate IDs out of sequence
func WithDrift(duration time.Duration) Option {
	return func(generator *Generator) {
		generator.drift = true
		generator.duration = duration
		time.Sleep(duration)
	}
}
