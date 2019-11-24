package heroku

import (
	"context"
	"fmt"
)

type Heroku struct {
}

func (p Heroku) Name() string {
	return "heroku"
}

func (p Heroku) String() string {
	return "github.com/gobuffalo/heroku"
}

func (p Heroku) Generate(ctx context.Context, args []string) error {
	fmt.Println(">>>TODO Generating Heroku ", args)
	return nil
}

func New() Heroku {
	return Heroku{}
}
