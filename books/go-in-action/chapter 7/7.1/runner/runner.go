package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	ctx       context.Context
	cancel    context.CancelFunc
	tasks     []func(int)
	complete  chan error
	interrupt chan os.Signal
}

var ErrInterupt error = errors.New("interupt from OS")
var ErrTimeout error = errors.New("Timeout error")

func NewRunner(d time.Duration) *Runner {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	return &Runner{
		ctx:       ctx,
		cancel:    cancel,
		complete:  make(chan error, 1),
		interrupt: make(chan os.Signal, 1),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	for _, t := range tasks {
		r.tasks = append(r.tasks, t)
	}
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	defer signal.Stop(r.interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.interrupt:
		r.cancel()
		return ErrInterupt
	}
}

func (r *Runner) run() error {
	for index, task := range r.tasks {
		select {
		case <-r.ctx.Done():
			return ErrTimeout
		default:
			task(index)
		}
	}

	println("Runner completed successfuly")
	return nil
}
