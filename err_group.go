package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"os"
	"time"
)


type Conn interface {
	Send(context.Context, os.File) error
}

func call(ctx context.Context, conns []Conn, file os.File) error {
	ctx, cancel := context.WithCancel(ctx)

	// create error group
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(3)

	// process send requests in parallel
	for _, conn := range conns {
		func(c Conn) {
			eg.Go(func() error {
				return c.Send(ctx, file)
			})
		}(conn)
	}

	// set timer to 2 mins
	timeout := time.After(2 * time.Second)
	for {
		select {
		case <-timeout:
			cancel()
		}
	}

	// wait and return 
	return eg.Wait()
}
