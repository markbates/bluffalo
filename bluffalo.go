package bluffalo

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/markbates/bluffalo/internal/cmdx"
)

// Bluffalo represents the `bluffalo` cli.
type Bluffalo struct {
	context.Context
	flags   *flag.FlagSet
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	Plugins Plugins
	version bool
	help    bool
}

func New(ctx context.Context) (*Bluffalo, error) {
	b := &Bluffalo{
		Context: ctx,
		Stdin:   cmdx.Stdin(ctx),
		Stdout:  cmdx.Stdout(ctx),
		Stderr:  cmdx.Stderr(ctx),
	}
	b.setFlags()
	return b, nil
}

func (b *Bluffalo) Flags() *flag.FlagSet {
	if b.flags == nil {
		b.setFlags()
	}
	return b.flags
}

func (b *Bluffalo) setFlags() {
	b.flags = flag.NewFlagSet("bluffalo", flag.ContinueOnError)
	b.flags.BoolVar(&b.version, "v", false, "display version")
	b.flags.BoolVar(&b.help, "h", false, "display help")
	cmdx.Usage(b.Context, b.flags)
}

func (b *Bluffalo) Main(ctx context.Context, args []string) error {
	flags := b.Flags()
	if err := flags.Parse(args); err != nil {
		return err
	}
	args = flags.Args()

	if len(args) == 0 {
		flags.Usage()
		return nil
	}

	arg := args[0]
	if len(args) > 0 {
		args = args[1:]
	}

	switch arg {
	case "fix":
		return b.Plugins.Fix(ctx, args)
	case "generate":
		return b.Plugins.Generate(ctx, args)
	}

	if b.version {
		fmt.Fprintln(b.Stdout, "ssh")
		return nil
	}

	if b.help {
		flags.Usage()
		return nil
	}
	return nil
}
