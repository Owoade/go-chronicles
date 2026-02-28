package main

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	tasks     []func(int)
	timeout   <-chan time.Time
	complete  chan error
	interrupt chan os.Signal
}

var ErrInterupt error = errors.New("interupt from OS")
var ErrTimeout error = errors.New("Timeout error")

func NewRunner(d time.Duration) *Runner {
	return &Runner{
		timeout:   time.After(d),
		complete:  make(chan error),
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
	go func() {
		r.complete <- r.run()
	}()

	select {
	case <-r.complete:
		return nil
	case <-r.interrupt:
		return ErrInterupt
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for index, task := range r.tasks {
		task(index)
	}

	println("Runner completed successfuly")
	return nil
}
