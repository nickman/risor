package redismod

import (
	"context"
	"errors"
	"fmt"
	"github.com/risor-io/risor/object"
)

func IsBool(obj object.Object) (bool, bool) {
	if obj != nil {
		if obj.Type() == object.BOOL {
			actual := obj.(*object.Bool)
			return true, actual.Value()
		}
	}
	return false, false
}

func IsActualBool(obj object.Object) bool {
	if obj != nil {
		if obj.Type() == object.BOOL {
			actual := obj.(*object.Bool)
			return actual.Value()
		}
	}
	return false
}

func DoCallback(ctx context.Context, callbackFx object.Callable, callbackArgs ...object.Object) (bool, error) {
	callbackResponse := callbackFx.Call(ctx, callbackArgs...)
	isBool, actualBool := IsBool(callbackResponse)
	if isBool {
		return actualBool, nil
	}
	// otherwise, if an error, return the error thrown from the callback
	if object.IsError(callbackResponse) {
		return false, errors.New(callbackResponse.Inspect())
	} else {
		return false, errors.New(fmt.Sprintf("type error: callbackFx expected a boolean response from callback fx (got %s)", callbackResponse.Type()))
	}
}
