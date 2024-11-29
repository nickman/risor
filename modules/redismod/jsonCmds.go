package redismod

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/risor-io/risor/object"
)

/*
TODO: JSONNumIncrBy, JSONObjKeys, JSONObjLen, JSONSet, JSONSetMode, JSONStrAppend, JSONStrLen, JSONToggle, JSONType
*/

func (c *RedisConn) JsonCmdsGetAttr(name string) (object.Object, bool) {
	switch name {
	case "jsonArrAppend":
		return object.NewBuiltin("redis.conn.jsonArrAppend", c.JSONArrAppend), true
	case "jsonArrIndex":
		return object.NewBuiltin("redis.conn.jsonArrIndex", c.JSONArrIndex), true
	case "jsonArrIndexWithArgs":
		return object.NewBuiltin("redis.conn.jsonArrIndexWithArgs", c.JSONArrIndexWithArgs), true
	case "jsonArrInsert":
		return object.NewBuiltin("redis.conn.jsonArrInsert", c.JSONArrInsert), true
	case "jsonArrLen":
		return object.NewBuiltin("redis.conn.jsonArrLen", c.JSONArrLen), true
	case "jsonArrPop":
		return object.NewBuiltin("redis.conn.jsonArrPop", c.JSONArrPop), true
	case "jsonArrTrim":
		return object.NewBuiltin("redis.conn.jsonArrTrim", c.JSONArrTrim), true
	case "jsonArrTrimWithArgs":
		return object.NewBuiltin("redis.conn.jsonArrTrimWithArgs", c.JSONArrTrimWithArgs), true
	}
	return nil, false
}

func intArrToList(intArr *redis.IntSliceCmd) *object.List {
	size := len(intArr.Val())
	ints := make([]object.Object, size, size)
	for idx := 0; idx < size; idx++ {
		ints[idx] = object.NewInt(intArr.Val()[idx])
	}
	return object.NewList(ints)
}

func (c *RedisConn) JSONArrAppend(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 3 {
		return object.TypeErrorf("type error: redis.conn.jsonArrAppend() takes at least three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrAppend() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrAppend() expected a string argument for path (got %s)", args[1].Type())
	}
	values := make([]interface{}, len(args)-2)
	for i, arg := range args[2:] {
		values[i] = arg.Interface()
	}
	intArrCmd := c.cmdable.JSONArrAppend(c.ctx, key.Value(), path.Value(), values...)
	if err := intArrCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return intArrToList(intArrCmd)
}

func (c *RedisConn) JSONArrIndex(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.jsonArrIndex() takes exactly three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrIndex() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrIndex() expected a string argument for path (got %s)", args[1].Type())
	}
	value := args[2].Interface()
	intArrCmd := c.cmdable.JSONArrIndex(c.ctx, key.Value(), path.Value(), value)
	if err := intArrCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return intArrToList(intArrCmd)
}

func (c *RedisConn) JSONArrIndexWithArgs(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 4 {
		return object.TypeErrorf("type error: redis.conn.jsonArrIndexWithArgs() takes at least four arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrIndexWithArgs() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrIndexWithArgs() expected a string argument for path (got %s)", args[1].Type())
	}
	//value := args[2].Interface()
	start, ok := args[2].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrIndexWithArgs() expected an int argument for start (got %s)", args[3].Type())
	}
	if start != nil {
		start.Value()
	}
	stop, ok := args[3].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrIndexWithArgs() expected an int argument for stop (got %s)", args[4].Type())
	}
	valueObjects := args[4:]
	size := len(valueObjects)
	values := make([]interface{}, size, size)
	for idx := 0; idx < size; idx++ {
		values[idx] = valueObjects[idx].Interface()
	}
	// JSONArrIndexWithArgs(ctx context.Context, key, path string, options *JSONArrIndexArgs, value ...interface{}) *IntSliceCmd
	intArrCmd := c.cmdable.JSONArrIndexWithArgs(c.ctx, key.Value(), path.Value(), newJSONArrIndexArgs(start, stop), values...)
	if err := intArrCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return intArrToList(intArrCmd)
}

func newJSONArrIndexArgs(start, stop *object.Int) *redis.JSONArrIndexArgs {
	var st *int
	if stop != nil {
		c := int(stop.Value())
		st = &c
	} else {
		st = nil
	}
	return &redis.JSONArrIndexArgs{
		Start: int(start.Value()),
		Stop:  st,
	}
}

func (c *RedisConn) JSONArrInsert(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 4 {
		return object.TypeErrorf("type error: redis.conn.jsonArrInsert() takes at least four arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrInsert() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrInsert() expected a string argument for path (got %s)", args[1].Type())
	}
	index, ok := args[2].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrInsert() expected an int argument for index (got %s)", args[2].Type())
	}
	values := make([]interface{}, len(args)-3)
	for i, arg := range args[3:] {
		values[i] = arg.Interface()
	}
	intArrCmd := c.cmdable.JSONArrInsert(c.ctx, key.Value(), path.Value(), index.Value(), values...)
	if err := intArrCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return intArrToList(intArrCmd)
}

func (c *RedisConn) JSONArrLen(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.jsonArrLen() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrLen() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrLen() expected a string argument for path (got %s)", args[1].Type())
	}
	// JSONArrLen(ctx context.Context, key, path string) *IntSliceCmd
	intArrCmd := c.cmdable.JSONArrLen(c.ctx, key.Value(), path.Value())
	if err := intArrCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return intArrToList(intArrCmd)
}

func (c *RedisConn) JSONArrPop(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 || len(args) > 3 {
		return object.TypeErrorf("type error: redis.conn.jsonArrPop() takes two or three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrPop() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrPop() expected a string argument for path (got %s)", args[1].Type())
	}
	var index int64
	if len(args) == 3 {
		idx, ok := args[2].(*object.Int)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.jsonArrPop() expected an int argument for index (got %s)", args[2].Type())
		}
		index = idx.Value()
	} else {
		index = -1
	}
	// JSONArrPop(ctx context.Context, key, path string, index int) *StringSliceCmd
	jsonCmd := c.cmdable.JSONArrPop(c.ctx, key.Value(), path.Value(), int(index))
	if err := jsonCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewStringList(jsonCmd.Val())
}

func (c *RedisConn) JSONArrTrim(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.jsonArrTrim() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrTrim() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrTrim() expected a string argument for path (got %s)", args[1].Type())
	}
	intArrCmd := c.cmdable.JSONArrTrim(c.ctx, key.Value(), path.Value())
	if err := intArrCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return intArrToList(intArrCmd)
}

func (c *RedisConn) JSONArrTrimWithArgs(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 4 {
		return object.TypeErrorf("type error: redis.conn.jsonArrTrimWithArgs() takes exactly four arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrTrimWithArgs() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrTrimWithArgs() expected a string argument for path (got %s)", args[1].Type())
	}
	start, ok := args[2].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrTrimWithArgs() expected an int argument for start (got %s)", args[2].Type())
	}
	stop, ok := args[3].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonArrTrimWithArgs() expected an int argument for stop (got %s)", args[3].Type())
	}
	// JSONArrTrimWithArgs(ctx context.Context, key, path string, options *JSONArrTrimArgs) *IntSliceCmd
	sta := int(start.Value())
	var stp *int
	if stop != nil {
		c := int(stop.Value())
		stp = &c
	} else {
		stp = nil
	}
	options := &redis.JSONArrTrimArgs{
		Start: sta,
		Stop:  stp,
	}
	intArrCmd := c.cmdable.JSONArrTrimWithArgs(c.ctx, key.Value(), path.Value(), options)
	if err := intArrCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return intArrToList(intArrCmd)
}

func (c *RedisConn) JSONClear(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.jsonClear() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonClear() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonClear() expected a string argument for path (got %s)", args[1].Type())
	}
	intCmd := c.cmdable.JSONClear(c.ctx, key.Value(), path.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) JSONDebugMemory(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.jsonDebugMemory() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonDebugMemory() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonDebugMemory() expected a string argument for path (got %s)", args[1].Type())
	}
	intCmd := c.cmdable.JSONDebugMemory(c.ctx, key.Value(), path.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) JSONDel(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.jsonDel() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonDel() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonDel() expected a string argument for path (got %s)", args[1].Type())
	}
	intCmd := c.cmdable.JSONDel(c.ctx, key.Value(), path.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) JSONForget(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.jsonForget() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonForget() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonForget() expected a string argument for path (got %s)", args[1].Type())
	}
	intCmd := c.cmdable.JSONForget(c.ctx, key.Value(), path.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) JSONGet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.jsonGet() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonGet() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonGet() expected a string argument for path (got %s)", args[1].Type())
	}
	stringCmd := c.cmdable.JSONGet(c.ctx, key.Value(), path.Value())
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

//func (c *RedisConn) JSONGetWithArgs(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) < 2 {
//		return object.TypeErrorf("type error: redis.conn.jsonGetWithArgs() takes at least two arguments (%d given)", len(args))
//	}
//	key, ok := args[0].(*object.String)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.jsonGetWithArgs() expected a string argument for key (got %s)", args[0].Type())
//	}
//	paths := make([]string, len(args)-1)
//	for i, arg := range args[1:] {
//		path, ok := arg.(*object.String)
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.jsonGetWithArgs() expected string arguments for paths (got %s)", arg.Type())
//		}
//		paths[i] = path.Value()
//	}
//	stringCmd := c.cmdable.JSONGetWithArgs(c.ctx, key.Value(), paths...)
//	if err := stringCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	return object.NewString(stringCmd.Val())
//}

func (c *RedisConn) JSONMerge(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.jsonMerge() takes exactly three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonMerge() expected a string argument for key (got %s)", args[0].Type())
	}
	path, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonMerge() expected a string argument for path (got %s)", args[1].Type())
	}
	json, ok := args[2].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonMerge() expected a string argument for json (got %s)", args[2].Type())
	}
	stringCmd := c.cmdable.JSONMerge(c.ctx, key.Value(), path.Value(), json.Value())
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

//func (c *RedisConn) JSONMSetArgs(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) < 2 || len(args)%2 != 0 {
//		return object.TypeErrorf("type error: redis.conn.jsonMSetArgs() takes an even number of arguments (%d given)", len(args))
//	}
//	kvPairs := make([]interface{}, len(args))
//	for i, arg := range args {
//		kvPairs[i] = arg.Interface()
//	}
//	// JSONMSetArgs(ctx context.Context, docs []JSONSetArgs) *StatusCmd
//	statusCmd := c.cmdable.JSONMSetArgs(c.ctx, kvPairs...)
//	if err := statusCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	return object.NewString(statusCmd.String())
//}

func (c *RedisConn) JSONMSet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 || len(args)%2 != 0 {
		return object.TypeErrorf("type error: redis.conn.jsonMSet() takes an even number of arguments (%d given)", len(args))
	}
	kvPairs := make([]interface{}, len(args))
	for i, arg := range args {
		kvPairs[i] = arg.Interface()
	}
	statusCmd := c.cmdable.JSONMSet(c.ctx, kvPairs...)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) JSONMGet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.jsonMGet() takes at least two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.jsonMGet() expected a string argument for key (got %s)", args[0].Type())
	}
	paths := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		path, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.jsonMGet() expected string arguments for paths (got %s)", arg.Type())
		}
		paths[i] = path.Value()
	}
	// JSONMGet(ctx context.Context, path string, keys ...string) *JSONSliceCmd
	jsonSliceCmd := c.cmdable.JSONMGet(c.ctx, key.Value(), paths...)
	if err := jsonSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}

	return jsonSliceCmdToList(jsonSliceCmd)
}

func jsonSliceCmdToList(jsonSliceCmd *redis.JSONSliceCmd) *object.List {
	ifaces := jsonSliceCmd.Val()
	size := len(ifaces)
	vals := make([]object.Object, size, size)
	for idx := 0; idx < size; idx++ {
		vals[idx] = object.FromGoType(ifaces[idx])
	}
	return object.NewList(vals)
}
