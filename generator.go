package snowflakes

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	// ErrMachineIDTooSmall is returned when the machine ID is too small for the number of bits
	ErrMachineIDTooSmall = errors.New("machine ID is too small")
	// ErrMachineIDTooLarge is returned when the machine ID is too large for the number of bits
	ErrMachineIDTooLarge = errors.New("machine ID is too large")
	// ErrMachineBitsTooSmall is returned when the number of bits for the machine ID is too small
	ErrMachineBitsTooSmall = errors.New("machine ID bits is too small")
	// ErrMachineBitsTooLarge is returned when the number of bits for the machine ID is too large
	ErrMachineBitsTooLarge = errors.New("machine ID bits is too large")
	// ErrOutOfSequence is returned when the sequence number overflows
	ErrOutOfSequence = errors.New("sequence number overflow")
	// ErrTimeStampTooLarge is returned when the timestamp is too large
	ErrTimeStampTooLarge = errors.New("timestamp is too large")
)

type ID int64

func (id ID) Time() int64 {
	return int64(id) >> 22
}

func (id ID) MachineID() int {
	return int((id >> 12) & (1<<10 - 1))
}

func (id ID) SequenceNumber() int {
	return int(id & (1<<12 - 1))
}

func (id ID) String() string {
	return fmt.Sprintf("ID{time: %d, machine: %d, sequence: %d, id: %d}", id.Time(), id.MachineID(), id.SequenceNumber(), int64(id))
}

var (
	maxTimestamp = int64(1<<41 - 1)
)

type Option func(*Generator)

type TimeFunc func() int64

func defaultTimeFunc() int64 {
	return time.Now().UTC().UnixMilli()
}

// Generator is a snowflake ID generator
type Generator struct {
	machineSequenceNumber     int64
	lastCurrentTime           int64
	timeFunc                  TimeFunc
	sleepFunc                 func()
	machineId                 int
	machineIdBits             int
	epoch                     int64
	maxMachineSequenceNumber  int64
	machineSequenceNumberBits int
	mutex                     sync.Mutex
}

// NewGenerator creates a new snowflake ID generator
// machineId is the unique ID of the machine running the generator
// opts are the options to configure the generator
// Returns a new snowflake ID generator
// Returns an error if the machineId is too large for the number of bits
// Returns an error if the machineIdBits is invalid
func NewGenerator(machineId int, opts ...Option) (*Generator, error) {
	g := &Generator{
		timeFunc:                  defaultTimeFunc,
		machineIdBits:             10,
		machineSequenceNumberBits: 12,
		machineId:                 machineId,
		sleepFunc: func() {
			nano := time.Duration(time.Now().UnixNano())
			milli := nano.Truncate(time.Millisecond)
			milli = milli + time.Millisecond
			delay := milli - nano
			time.Sleep(delay + 1*time.Nanosecond)
		},
	}

	for _, opt := range opts {
		opt(g)
	}

	maxMachineId := 1<<g.machineIdBits - 1

	if g.machineId > maxMachineId {
		return nil, ErrMachineIDTooLarge
	}

	if g.machineId < 0 {
		return nil, ErrMachineIDTooSmall
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
func (g *Generator) NextID() (ID, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	currentTime := g.timeFunc() - g.epoch

	if currentTime > g.lastCurrentTime {
		g.lastCurrentTime = currentTime
		g.machineSequenceNumber = 0
	}

	g.machineSequenceNumber++

	if g.machineSequenceNumber-1 > g.maxMachineSequenceNumber {
		g.machineSequenceNumber--
		return 0, ErrOutOfSequence
	}

	if currentTime > maxTimestamp {
		return 0, ErrTimeStampTooLarge
	}

	return ID(currentTime<<22 | int64(g.machineId)<<g.machineSequenceNumberBits | int64(g.machineSequenceNumber-1)), nil
}

// BlockingNextID generates a new snowflake ID, blocking until the next ID can be generated
func (g *Generator) BlockingNextID(ctx context.Context) (ID, error) {
	id, err := g.NextID()
	for err == ErrOutOfSequence {
		if ctx != nil && ctx.Err() != nil {
			return 0, ctx.Err()
		}
		g.sleepFunc()
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

// WithEpoch sets the epoch for the generator
func WithEpoch(epoch time.Time) Option {
	return func(generator *Generator) {
		generator.epoch = epoch.UnixMilli()
	}
}
