package snowflakes

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"sync/atomic"
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
)

const (
	timeShift = 22
)

// ID is a snowflake ID
type ID uint64

func (id ID) String() string {
	return fmt.Sprintf("%x", uint64(id))
}

func (id ID) LowerHexString() string {
	return fmt.Sprintf("%x", uint64(id))
}

func (id ID) UpperHexString() string {
	return fmt.Sprintf("%X", uint64(id))
}

func (id ID) Base64String() string {
	var b [8]byte
	binary.LittleEndian.PutUint64(b[:], uint64(id))
	return base64.StdEncoding.EncodeToString(b[:])
}

// DecodedID is a snowflake ID decoded into its components
type DecodedID struct {
	ID        uint64
	Timestamp uint64
	MachineID uint64
	Sequence  uint64
}

// String returns a string representation of the decoded ID
func (id DecodedID) String() string {
	return fmt.Sprintf("ID: %d, Timestamp: %d, MachineID: %d, Sequence: %d", id.ID, id.Timestamp, id.MachineID, id.Sequence)
}

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
	timeTravel     bool
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
	}

	for _, opt := range opts {
		opt(g)
	}

	maxMachineId := uint64(1<<g.machineIDBits - 1)

	if g.machineID > maxMachineId {
		return nil, ErrMachineIDTooLarge
	}

	if g.machineID < 0 {
		return nil, ErrMachineIDTooSmall
	}

	if g.machineIDBits < 1 {
		return nil, ErrMachineBitsTooSmall
	}

	if g.machineIDBits > 21 {
		return nil, ErrMachineBitsTooLarge
	}

	g.machineIDMask = maxMachineId
	g.sequenceMask = 1<<(timeShift-g.machineIDBits) - 1
	g.machineIDShift = timeShift - g.machineIDBits

	return g, nil
}

// DecodeID decodes a snowflake ID into its components
func (g *Generator) DecodeID(id ID) DecodedID {
	return DecodedID{
		ID:        uint64(id),
		Timestamp: uint64(id) >> timeShift,
		MachineID: uint64(id) >> g.machineIDShift & g.machineIDMask,
		Sequence:  uint64(id) & g.sequenceMask,
	}
}

// NextID generates a new snowflake ID
func (g *Generator) NextID() (ID, error) {

	now := int64(g.timeFunc()) - g.epoch

	for {
		currentID := g.currentID.Load()
		newCurrentID := currentID
		lastTime := int64(currentID >> timeShift)
		sequence := currentID & g.sequenceMask
		switch {
		case lastTime < now:
			lastTime = now
			newCurrentID = uint64(lastTime) << timeShift
		case sequence == g.sequenceMask:
			if !g.timeTravel {
				return 0, ErrOutOfSequence
			}
			newCurrentID = uint64(lastTime+1) << timeShift
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
func WithMachineIdBits(size uint64) Option {
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

// WithTimeTravel enables time travel for the generator
// This allows the generator to generate IDs for times in the future
// This increases performance but may generate IDs out of sequence
func WithTimeTravel() Option {
	return func(generator *Generator) {
		generator.timeTravel = true
	}
}
