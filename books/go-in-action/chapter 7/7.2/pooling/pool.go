package main

import (
	"errors"
	"io"
	"sync"
)

type FactoryFunction func() (io.Closer, error)
type Pool struct {
	mutex     sync.Mutex
	closed    bool
	resources chan io.Closer
	factory   FactoryFunction
}

var ErrPoolClosed error = errors.New("Pool closed")

func NewPool(factoryFunction FactoryFunction, size int) *Pool {
	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   factoryFunction,
	}
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		println("acquired from pool")
		return r, nil
	default:
		println("acquired from factory")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		r.Close()
	}

	select {
	case p.resources <- r:
		println("resource added to queue")
	default:
		r.Close()
		println("queue at capacity, disposing resource")
	}
}

func (p *Pool) Shutdown() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}
