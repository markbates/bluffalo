/*
This cli package would live in the bluffalo application.
*/
package cli

import (
	"context"

	"github.com/markbates/bluffalo"
	"github.com/markbates/bluffalo/fauxplugs/goth"
	"github.com/markbates/bluffalo/fauxplugs/heroku"
	"github.com/markbates/bluffalo/fauxplugs/plush"
	"github.com/markbates/bluffalo/fauxplugs/pop"
)

// Main is the entry point for the bluffalo application
// this is what will be called by main.go
// It would be used by tools like `bluffalo dev` and
// `bluffalo build`.
func Main(ctx context.Context, args []string) error {
	// app := actions.App()
	// if err := app.Serve(); err != nil {
	// 		return err
	// }
	return nil
}

// Bluffalo is the entry point for the `bluffalo` binary.
// It allows for registering plugins to enhance the `bluffalo` binary.
// 	bluffalo generate -h
// 	bluffalo generate pop ...
// 	bluffalo fix
// 	bluffalo fix plush ...
// 	bluffalo fix -h
func Bluffalo(ctx context.Context, args []string) error {
	b, err := bluffalo.New(ctx)
	if err != nil {
		return err
	}

	b.Plugins = append(b.Plugins,
		pop.New(),
		goth.New(),
		heroku.New(),
		plush.New(),
	)

	return b.Main(ctx, args)
}
