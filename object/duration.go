package object

import (
	"context"
	"fmt"
	"github.com/risor-io/risor/errz"
	"github.com/risor-io/risor/op"
	"time"
)

type Duration struct {
	*base
	value time.Duration
}

func (d *Duration) Type() Type {
	return DURATION
}

func (d *Duration) Value() time.Duration {
	return d.value
}

func (d *Duration) Inspect() string {
	return fmt.Sprintf("duration(%s)", d.value.String())
}

func (d *Duration) GetAttr(name string) (Object, bool) {
	d.value.String()
	switch name {
	case "abs":
		return NewBuiltin("duration.abs", d.Abs), true
	case "hours":
		return NewBuiltin("duration.hours", d.Hours), true
	case "microseconds":
		return NewBuiltin("duration.microseconds", d.Microseconds), true
	case "milliseconds":
		return NewBuiltin("duration.milliseconds", d.Milliseconds), true
	case "minutes":
		return NewBuiltin("duration.minutes", d.Minutes), true
	case "seconds":
		return NewBuiltin("duration.seconds", d.Seconds), true
	case "round":
		return NewBuiltin("duration.round", d.Round), true
	case "truncate":
		return NewBuiltin("duration.truncate", d.Truncate), true
	case "tostring":
		return NewBuiltin("duration.tostring", d.ToString), true
	default:
		return nil, false
	}
}

func (d *Duration) Interface() interface{} {
	return d.value
}

func (d *Duration) String() string {
	return d.Inspect()
}

func (d *Duration) Compare(other Object) (int, error) {
	otherDur, ok := other.(*Duration)
	if !ok {
		return 0, errz.TypeErrorf("type error: unable to compare duration and %s", other.Type())
	}
	if d.value == otherDur.value {
		return 0, nil
	}
	if d.value > otherDur.value {
		return 1, nil
	}
	return -1, nil
}

func (d *Duration) Equals(other Object) Object {
	if other.Type() == DURATION && d.value == other.(*Duration).value {
		return True
	}
	return False
}

func (d *Duration) RunOperation(opType op.BinaryOpType, right Object) Object {
	return TypeErrorf("type error: unsupported operation for duration: %v", opType)
}

func NewDuration(d time.Duration) *Duration {
	return &Duration{value: d}
}

func ParseDuration(s string) (*Duration, error) {
	if d, err := time.ParseDuration(s); err != nil {
		return nil, err
	} else {
		return NewDuration(d), nil
	}
}

func AsDuration(obj Object) (result time.Duration, err *Error) {
	s, ok := obj.(*Duration)
	if !ok {
		return time.Since(time.Now()), TypeErrorf("type error: expected a duration (%s given)", obj.Type())
	}
	return s.value, nil
}

func (d *Duration) Abs(ctx context.Context, args ...Object) Object {
	if len(args) != 0 {
		return NewArgsError("duration.abs", 0, len(args))
	}
	return NewDuration(d.value.Abs())
}

func (d *Duration) Hours(ctx context.Context, args ...Object) Object {
	if len(args) != 0 {
		return NewArgsError("duration.hours", 0, len(args))
	}
	return NewFloat(d.value.Hours())
}

func (d *Duration) Microseconds(ctx context.Context, args ...Object) Object {
	if len(args) != 0 {
		return NewArgsError("duration.microseconds", 0, len(args))
	}
	return NewInt(d.value.Microseconds())
}

func (d *Duration) Milliseconds(ctx context.Context, args ...Object) Object {
	if len(args) != 0 {
		return NewArgsError("duration.milliseconds", 0, len(args))
	}
	return NewInt(d.value.Milliseconds())
}

func (d *Duration) Minutes(ctx context.Context, args ...Object) Object {
	if len(args) != 0 {
		return NewArgsError("duration.minutes", 0, len(args))
	}
	return NewFloat(d.value.Minutes())
}

func (d *Duration) Nanoseconds(ctx context.Context, args ...Object) Object {
	if len(args) != 0 {
		return NewArgsError("duration.nanoseconds", 0, len(args))
	}
	return NewInt(d.value.Nanoseconds())
}

func (d *Duration) Seconds(ctx context.Context, args ...Object) Object {
	if len(args) != 0 {
		return NewArgsError("duration.seconds", 0, len(args))
	}
	return NewFloat(d.value.Seconds())
}

func (d *Duration) Round(ctx context.Context, args ...Object) Object {
	if len(args) != 1 {
		return NewArgsError("duration.round", 1, len(args))
	}
	if m, err := AsDuration(args[0]); err != nil {
		return err
	} else {
		return NewDuration(d.value.Round(m))
	}
}

func (d *Duration) Truncate(ctx context.Context, args ...Object) Object {
	if len(args) != 1 {
		return NewArgsError("duration.truncate", 1, len(args))
	}
	if m, err := AsDuration(args[0]); err != nil {
		return err
	} else {
		return NewDuration(d.value.Truncate(m))
	}
}

func (d *Duration) ToString(ctx context.Context, args ...Object) Object {
	if len(args) != 0 {
		return NewArgsError("duration.tostring", 0, len(args))
	}
	return NewString(d.Inspect())
}
