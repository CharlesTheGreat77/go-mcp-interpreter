package internal

import (
	"bytes"
	"context"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func ExecuteGoCodeInYaegi(ctx context.Context, code string) (string, error) {
	var output bytes.Buffer

	// redirect both stdout and stderr to output buffer
	i := interp.New(interp.Options{
		Stdout: &output,
		Stderr: &output,
	})

	i.Use(stdlib.Symbols) // std lib allowed

	_, err := i.Eval(code) // execute... as you'd expect

	return output.String(), err
}
