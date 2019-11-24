package plush

import (
	"context"
	"fmt"
)

type Plush struct {
}

func (p Plush) Name() string {
	return "plush"
}

func (p Plush) String() string {
	return "github.com/gobuffalo/plush"
}

func (p Plush) Fix(ctx context.Context, args []string) error {
	fmt.Println(">>>TODO Fixing Plush ", args)
	return nil
}

func (p Plush) Generate(ctx context.Context, args []string) error {
	fmt.Println(">>>TODO Generating Plush ", args)
	return nil
}

func New() Plush {
	return Plush{}
}
