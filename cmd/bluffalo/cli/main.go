package cli

import (
	"context"
	"path"

	"github.com/gobuffalo/here"
	"github.com/markbates/jim"
)

func Main(ctx context.Context, args []string) error {
	her, err := here.Dir(".")
	if err != nil {
		return err
	}

	ci, err := here.Package(path.Join(her.ImportPath, "cli"))
	if err != nil {
		return err
	}

	t := &jim.Task{
		Info: ci,
		Args: args,
		Pkg:  ci.ImportPath,
		Sel:  ci.Name,
		Name: "Bluffalo",
	}

	return jim.Run(ctx, t)
}
