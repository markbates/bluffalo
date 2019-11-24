package pop

import (
	"context"
	"fmt"
)

type Pop struct {
}

func (p Pop) Name() string {
	return "pop"
}

func (p Pop) String() string {
	return "github.com/gobuffalo/pop"
}

func (p Pop) Fix(ctx context.Context, args []string) error {
	fmt.Println(">>>TODO Fixing Pop ", args)
	return nil
}

func (p Pop) Generate(ctx context.Context, args []string) error {
	fmt.Println(">>>TODO Generating Pop ", args)
	return nil
}

func New() Pop {
	return Pop{}
}
