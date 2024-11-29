package duration

import (
	"context"
	"time"

	"github.com/risor-io/risor/arg"
	"github.com/risor-io/risor/object"
)

func Parse(ctx context.Context, args ...object.Object) object.Object {
	if err := arg.Require("duration.parse", 1, args); err != nil {
		return err
	}
	dur, err := object.AsString(args[0])
	if err != nil {
		return err
	}
	d, parseErr := time.ParseDuration(dur)
	if parseErr != nil {
		return object.NewError(parseErr)
	}
	return object.NewDuration(d)
}

func Module() *object.Module {
	return object.NewBuiltinsModule("duration", map[string]object.Object{
		"parse": object.NewBuiltin("parse", Parse),
	})
}
