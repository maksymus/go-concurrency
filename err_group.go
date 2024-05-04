package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"os"
	"time"
)


func main() {

}

type Conn interface {
	Send(context.Context, os.File) error
}

type conn struct {
}

func (c conn) Send(ctx context.Context, file os.File) error {
	return nil
}

func call(ctx context.Context, conns []Conn, file os.File) error {
	ctx, cancel := context.WithCancel(ctx)

	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(3)

	for _, conn := range conns {
		func(c Conn) {
			eg.Go(func() error {
				return c.Send(ctx, file)
			})
		}(conn)
	}

	timeout := time.After(2 * time.Second)
	for {
		select {
		case <-timeout:
			cancel()
		}
	}

	return eg.Wait()
}
