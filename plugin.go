package bluffalo

import (
	"context"
	"flag"
	"fmt"
	"sort"

	"github.com/markbates/bluffalo/internal/cmdx"
)

// Plugin is the most basic interface a plugin can implement.
type Plugin interface {
	// Name is the name of the plugin.
	// This will also be used for the cli sub-command
	// 	"pop" | "heroku" | "auth" | etc...
	Name() string
}

type Plugins []Plugin

// Fixer is an optional interface a plugin can implement
// to be run with `bluffalo fix`. This should update the application
// to the current version of the plugin.
// The expectation is fixing of only one major revision.
type Fixer interface {
	Fix(ctx context.Context, args []string) error
}

// Fixer is an optional interface a plugin can implement
// to be run with `bluffalo fix`
type Generator interface {
	Generate(ctx context.Context, args []string) error
}

// Fix runs any Fixers that are in the Plugins.
// If no arguments are provided it will run all fixers in the Plugins.
// Otherwise Fix will run the fixers for the arguments provided.
// 	bluffalo fix
// 	bluffalo fix plush pop
// 	bluffalo fix -h
func (plugs Plugins) Fix(ctx context.Context, args []string) error {
	opts := struct {
		help bool
	}{}

	flags := flag.NewFlagSet("bluffalo fix", flag.ContinueOnError)
	flags.BoolVar(&opts.help, "h", false, "print this help")

	if err := flags.Parse(args); err != nil {
		return err
	}

	args = flags.Args()

	stderr := cmdx.Stderr(ctx)
	if opts.help {
		sort.Slice(plugs, func(i, j int) bool {
			return plugs[i].Name() < plugs[j].Name()
		})

		for _, p := range plugs {
			if _, ok := p.(Fixer); ok {
				fmt.Fprintf(stderr, "%s %s - [%s]\n", flags.Name(), p.Name(), p)
			}
		}
		return nil
	}

	if len(args) > 0 {
		fixers := map[string]Fixer{}
		for _, p := range plugs {
			f, ok := p.(Fixer)
			if !ok {
				continue
			}

			fixers[p.Name()] = f
		}

		for _, a := range args {
			f, ok := fixers[a]
			if !ok {
				return fmt.Errorf("unknown fixer %s", a)
			}
			if err := f.Fix(ctx, []string{}); err != nil {
				return err
			}
		}
		return nil
	}

	for _, p := range plugs {
		f, ok := p.(Fixer)
		if !ok {
			continue
		}

		if err := f.Fix(ctx, args); err != nil {
			return err
		}
	}
	return nil
}

// Generate will run the specified generator.
// 	bluffalo generate -h
// 	bluffalo generate pop ...
func (plugs Plugins) Generate(ctx context.Context, args []string) error {
	opts := struct {
		help bool
	}{}

	flags := flag.NewFlagSet("bluffalo generate", flag.ContinueOnError)
	flags.BoolVar(&opts.help, "h", false, "print this help")

	if err := flags.Parse(args); err != nil {
		return err
	}

	args = flags.Args()
	if opts.help || len(args) == 0 {
		sort.Slice(plugs, func(i, j int) bool {
			return plugs[i].Name() < plugs[j].Name()
		})

		stderr := cmdx.Stderr(ctx)
		for _, p := range plugs {
			if _, ok := p.(Generator); ok {
				fmt.Fprintf(stderr, "%s %s - [%s]\n", flags.Name(), p.Name(), p)
			}
		}
		return nil
	}

	arg := args[0]
	if len(args) > 0 {
		args = args[1:]
	}

	for _, p := range plugs {
		f, ok := p.(Generator)
		if !ok {
			continue
		}
		if p.Name() != arg {
			continue
		}

		return f.Generate(ctx, args)
	}
	return fmt.Errorf("unknown generator %s", arg)
}
