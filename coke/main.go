package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/markbates/bluffalo/cmd/bluffalo/cli"
)

func main() {
	ctx := context.Background()

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	if err := cli.Main(ctx, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
