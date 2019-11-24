package cli

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/markbates/bluffalo/internal/cmdx"
	"github.com/stretchr/testify/require"
)

func Test_Main(t *testing.T) {
	r := require.New(t)

	ctx := context.Background()
	args := []string{}

	err := Main(ctx, args)
	r.NoError(err)
}

func Test_Main_Help(t *testing.T) {
	r := require.New(t)

	bb := &bytes.Buffer{}

	ctx := context.Background()
	ctx = cmdx.WithStderr(bb, ctx)

	args := []string{"-h"}

	err := Main(ctx, args)
	r.NoError(err)

	exp := "Usage of bluff:\n  -h\tdisplay help\n  -v\tdisplay version"
	act := strings.TrimSpace(bb.String())
	r.Equal(exp, act)
}
