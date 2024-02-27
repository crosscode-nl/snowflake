package snowflakes

import (
	"context"
	"errors"
)

var (
	// ErrPoolSizeTooSmall is returned when the pool size is too small
	ErrPoolSizeTooSmall = errors.New("pool size is too small")
)

// Pool is a thread safe pool of snowflake IDs
type Pool struct {
	generator *Generator
	idPool    chan ID
	ctx       context.Context
}

// NewPool creates a new pool of snowflake IDs
func NewPool(poolSize int, generator *Generator) (pool *Pool, err error, cancel func()) {
	if poolSize < 1 {
		return nil, ErrPoolSizeTooSmall, func() {}
	}
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				id, err := generator.NextID()
				if err != nil {
					generator.sleepFunc()
					continue
				}
				select {
				case pool.idPool <- id:
				default:
					generator.sleepFunc()
				}
			}
		}
	}()

	return &Pool{
		generator: generator,
		idPool:    make(chan ID, poolSize),
		ctx:       ctx,
	}, nil, cancel
}

// NextID gets the next ID from the pool
func (p *Pool) NextID() (ID, error) {
	select {
	case id := <-p.idPool:
		return id, nil
	default:
		return p.generator.NextID()
	}
}

// BlockingNextID gets the next ID from the pool, blocking until the next ID can be generated
func (p *Pool) BlockingNextID(ctx context.Context) (ID, error) {
	select {
	case id := <-p.idPool:
		return id, nil
	default:
		return p.generator.BlockingNextID(ctx)
	}
}
