package goth

import (
	"context"
	"fmt"
)

type Goth struct {
}

func (p Goth) Name() string {
	return "goth"
}

func (p Goth) String() string {
	return "github.com/gobuffalo/goth"
}

func (p Goth) Generate(ctx context.Context, args []string) error {
	fmt.Println(">>>TODO Generating Goth ", args)
	return nil
}

func New() Goth {
	return Goth{}
}
