package snowflakes

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
)

type Option func(*Generator)

type TimeFunc func() int64

func defaultTimeFunc() int64 {
	return time.Now().UTC().UnixMilli()
}

// Generator is a snowflake ID generator
type Generator struct {
	timeFunc                  TimeFunc
	machineId                 int
	machineIdBits             int
	epoch                     int64
	machineSequenceNumber     atomic.Int32
	maxMachineSequenceNumber  int32
	machineSequenceNumberBits int
	currentTime               atomic.Int64
}

// New creates a new snowflake ID generator
// machineId is the unique ID of the machine running the generator
// opts are the options to configure the generator
// Returns a new snowflake ID generator
// Returns an error if the machineId is too large for the number of bits
// Returns an error if the machineIdBits is invalid
func New(machineId int, opts ...Option) (*Generator, error) {
	g := &Generator{
		timeFunc:      defaultTimeFunc,
		machineIdBits: 10,
		machineId:     machineId,
	}

	for _, opt := range opts {
		opt(g)
	}

	if g.machineId >= 1<<g.machineIdBits {
		return nil, ErrMachineIDTooLarge
	}

	if g.machineIdBits < 1 {
		return nil, ErrMachineBitsTooSmall
	}

	if g.machineIdBits > 21 {
		return nil, ErrMachineBitsTooLarge
	}

	g.maxMachineSequenceNumber = 1<<g.machineSequenceNumberBits - 1

	return g, nil
}

// NextID generates a new snowflake ID
func (g *Generator) NextID() (int64, error) {
	lastCurrentTime := g.currentTime.Load()
	currentTime := g.timeFunc() - g.epoch
	machineSequenceNumber := g.machineSequenceNumber.Load()
	if currentTime != lastCurrentTime {
		if g.currentTime.CompareAndSwap(lastCurrentTime, currentTime) {
			g.machineSequenceNumber.Store(0)
			machineSequenceNumber = 0
		} else {
			machineSequenceNumber = g.machineSequenceNumber.Load()
		}
	} else {
		machineSequenceNumber = g.machineSequenceNumber.Add(1)
		if machineSequenceNumber > g.maxMachineSequenceNumber {
			return 0, ErrOutOfSequence
		}
	}

	return currentTime<<22 | int64(g.machineId)<<g.machineSequenceNumberBits | int64(machineSequenceNumber), nil
}

// BlockingNextID generates a new snowflake ID, blocking until the next ID can be generated
func (g *Generator) BlockingNextID(ctx context.Context) (int64, error) {
	id, err := g.NextID()
	for err == ErrOutOfSequence {
		if ctx.Err() != nil {
			return 0, ctx.Err()
		}
		time.Sleep(100 * time.Microsecond)
		id, err = g.NextID()
	}
	return id, nil
}

// WithMachineIdBits sets the number of bits to use for the machine ID
func WithMachineIdBits(size int) Option {
	return func(generator *Generator) {
		generator.machineIdBits = size
		generator.machineSequenceNumberBits = 22 - size
	}
}

// WithTimeFunc sets the time function for the generator
func WithTimeFunc(timeFunc TimeFunc) Option {
	return func(generator *Generator) {
		generator.timeFunc = timeFunc
	}
}

// WithEpoch sets the epoch for the generator
func WithEpoch(epoch time.Time) Option {
	return func(generator *Generator) {
		generator.epoch = epoch.UnixMilli()
	}
}
