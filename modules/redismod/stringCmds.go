package redismod

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/risor-io/risor/object"
	"time"
)

func (c *RedisConn) StringCmdsGetAttr(name string) (object.Object, bool) {
	switch name {
	case "decr":
		return object.NewBuiltin("redis.conn.decr", c.Decr), true
	case "decrby":
		return object.NewBuiltin("redis.conn.decrby", c.DecrBy), true
	case "append":
		return object.NewBuiltin("redis.conn.append", c.Append), true
	case "set":
		return object.NewBuiltin("redis.conn.set", c.Set), true
	case "get":
		return object.NewBuiltin("redis.conn.get", c.Get), true
	case "incr":
		return object.NewBuiltin("redis.conn.incr", c.Incr), true
	case "incrby":
		return object.NewBuiltin("redis.conn.incrby", c.IncrBy), true
	case "lcs":
		return object.NewBuiltin("redis.conn.lcs", c.LCS), true
	case "getset":
		return object.NewBuiltin("redis.conn.getset", c.GetSet), true
	case "getrange":
		return object.NewBuiltin("redis.conn.getrange", c.GetRange), true
	case "getex":
		return object.NewBuiltin("redis.conn.getrange", c.GetEx), true
	case "getdel":
		return object.NewBuiltin("redis.conn.getrange", c.GetDel), true
	case "incrbyfloat":
		return object.NewBuiltin("redis.conn.getrange", c.IncrByFloat), true
	case "mget":
		return object.NewBuiltin("redis.conn.mget", c.MGet), true
	case "mgetfunc":
		return object.NewBuiltin("redis.conn.mgetfunc", c.MGetFunc), true
	case "mset":
		return object.NewBuiltin("redis.conn.mget", c.MSet), true
	case "msetnx":
		return object.NewBuiltin("redis.conn.mget", c.MSetNX), true
	case "setex":
		return object.NewBuiltin("redis.conn.mget", c.SetEx), true
	case "setnx":
		return object.NewBuiltin("redis.conn.setnx", c.SetNX), true
	case "setxx":
		return object.NewBuiltin("redis.conn.setxx", c.SetXX), true
	case "setrange":
		return object.NewBuiltin("redis.conn.setrange", c.SetRange), true
	case "strlen":
		return object.NewBuiltin("redis.conn.strlen", c.StrLen), true
	}
	return nil, false
}

func (c *RedisConn) Get(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.get() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.get() expected a string argument (got %s)", args[0].Type())
	}
	statusCmd := c.cmdable.Get(c.ctx, key.Value())
	if err := statusCmd.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return object.Nil
			//return object.NewString("")
		}
		return object.NewError(err)
	}
	return object.NewString(statusCmd.Val())
}

func (c *RedisConn) Append(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.append() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.append() expected a string argument (got %s)", args[0].Type())
	}
	value, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.append() expected a string argument (got %s)", args[1].Type())
	}
	// Append(ctx context.Context, key, value string) *IntCmd
	intCmd := c.cmdable.Append(c.ctx, key.Value(), value.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) Set(ctx context.Context, args ...object.Object) object.Object {
	argSize := len(args)
	if argSize != 2 && argSize != 3 {
		return object.TypeErrorf("type error: redis.conn.set() takes either two or three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.set() expected a string argument (got %s)", args[0].Type())
	}
	value, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.set() expected a string argument (got %s)", args[1].Type())
	}
	var exp time.Duration
	if argSize == 3 {
		expiration, ok := args[2].(*object.Duration)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.set() expected a duration argument (got %s)", args[2].Type())
		}
		exp = expiration.Value()
	} else {
		exp = 0 * time.Second
	}
	// Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
	statusCmd := c.cmdable.Set(c.ctx, key.Value(), value.Value(), exp)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.Val())
}

func (c *RedisConn) Incr(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.incr() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.incr() expected a string argument (got %s)", args[0].Type())
	}
	intCmd := c.cmdable.Incr(c.ctx, key.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) Decr(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.decr() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.decr() expected a string argument (got %s)", args[0].Type())
	}
	intCmd := c.cmdable.Decr(c.ctx, key.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) IncrBy(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.incrby() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.incrby() expected a string argument (got %s)", args[0].Type())
	}
	by, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.incrby() expected an integer argument (got %s)", args[1].Type())
	}
	// IncrBy(ctx context.Context, key string, value int64) *IntCmd
	intCmd := c.cmdable.IncrBy(c.ctx, key.Value(), by.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) DecrBy(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.decrby() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.decrby() expected a string argument (got %s)", args[0].Type())
	}
	by, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.decrby() expected an integer argument (got %s)", args[1].Type())
	}
	intCmd := c.cmdable.DecrBy(c.ctx, key.Value(), by.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) GetRange(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.getrange() takes exactly three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.getrange() expected a string argument for key (got %s)", args[0].Type())
	}
	start, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.getrange() expected an int argument for start (got %s)", args[1].Type())
	}
	end, ok := args[2].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.getrange() expected an int argument for end (got %s)", args[2].Type())
	}
	stringCmd := c.cmdable.GetRange(c.ctx, key.Value(), start.Value(), end.Value())
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) GetSet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.getset() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.getset() expected a string argument for key (got %s)", args[0].Type())
	}
	value, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.getset() expected a string argument for value (got %s)", args[1].Type())
	}
	stringCmd := c.cmdable.GetSet(c.ctx, key.Value(), value.Value())
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) GetEx(ctx context.Context, args ...object.Object) object.Object {
	argSize := len(args)
	if argSize != 1 && argSize != 2 {
		return object.TypeErrorf("type error: redis.conn.getex() takes one or two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.getex() expected a string argument for key (got %s)", args[0].Type())
	}
	var exp time.Duration
	if argSize == 2 {
		expiration, ok := args[1].(*object.Duration)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.getex() expected a duration argument for expiration (got %s)", args[1].Type())
		}
		exp = expiration.Value()
	} else {
		exp = 0 * time.Second
	}
	// GetEx(ctx context.Context, key string, expiration time.Duration) *StringCmd
	stringCmd := c.cmdable.GetEx(c.ctx, key.Value(), exp)
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) GetDel(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.getdel() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.getdel() expected a string argument for key (got %s)", args[0].Type())
	}
	// GetDel(ctx context.Context, key string) *StringCmd
	stringCmd := c.cmdable.GetDel(c.ctx, key.Value())
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) IncrByFloat(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.incrbyfloat() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.incrbyfloat() expected a string argument for key (got %s)", args[0].Type())
	}
	by, ok := args[1].(*object.Float)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.incrbyfloat() expected a float argument for by (got %s)", args[1].Type())
	}
	// IncrByFloat(ctx context.Context, key string, value float64) *FloatCmd
	floatCmd := c.cmdable.IncrByFloat(c.ctx, key.Value(), by.Value())
	if err := floatCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewFloat(floatCmd.Val())
}

func (c *RedisConn) LCS(ctx context.Context, args ...object.Object) object.Object {
	argSize := len(args)
	if argSize != 2 && argSize != 3 {
		return object.TypeErrorf("type error: redis.conn.lcs() takes two or three arguments (%d given)", len(args))
	}
	key1, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.lcs() expected a string argument for key1 (got %s)", args[0].Type())
	}
	key2, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.lcs() expected a string argument for key2 (got %s)", args[1].Type())
	}
	lcsQuery := &redis.LCSQuery{
		Key1: key1.Value(),
		Key2: key2.Value(),
	}
	if argSize == 3 {
		cfgMap, ok := args[2].(*object.Map)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.lcs() expected a map argument for config (got %s)", args[2].Type())
		}
		cfgObjMap := cfgMap.Value()
		// TODO: Implement generic func for extracting values from config maps
		if minMatchLenObj, ok := cfgObjMap["min_match_len"]; ok {
			if minMatchLen, ok := minMatchLenObj.(*object.Int); !ok {
				return object.TypeErrorf("type error: redis.conn.lcs() expected an integer argument for 'min_match_len' (got %s)", minMatchLenObj.Type())
			} else {
				lcsQuery.MinMatchLen = int(minMatchLen.Value())
			}
		}
		if withMatchLenObj, ok := cfgObjMap["with_match_len"]; ok {
			if withMatchLen, ok := withMatchLenObj.(*object.Bool); !ok {
				return object.TypeErrorf("type error: redis.conn.lcs() expected a boolean argument for 'with_match_len' (got %s)", withMatchLenObj.Type())
			} else {
				lcsQuery.WithMatchLen = withMatchLen.Value()
			}
		}
		if lenObj, ok := cfgObjMap["len"]; ok {
			if len, ok := lenObj.(*object.Bool); !ok {
				return object.TypeErrorf("type error: redis.conn.lcs() expected a boolean argument for 'len' (got %s)", lenObj.Type())
			} else {
				lcsQuery.Len = len.Value()
			}
		}
		if idxObj, ok := cfgObjMap["idx"]; ok {
			if idx, ok := idxObj.(*object.Bool); !ok {
				return object.TypeErrorf("type error: redis.conn.lcs() expected a boolean argument for 'idx' (got %s)", idxObj.Type())
			} else {
				lcsQuery.Idx = idx.Value()
			}
		}
	}
	// LCS(ctx context.Context, q *LCSQuery) *LCSCmd
	/*
			// LCSQuery is a parameter used for the LCS command
			type LCSQuery struct {
				Key1         string
				Key2         string
				Len          bool
				Idx          bool
				MinMatchLen  int
				WithMatchLen bool
			}

			// LCSMatch is the result set of the LCS command.
			type LCSMatch struct {
				MatchString string
				Matches     []LCSMatchedPosition
				Len         int64
			}
		type LCSMatchedPosition struct {
			Key1 LCSPosition
			Key2 LCSPosition

			// only for withMatchLen is true
			MatchLen int64
		}

		type LCSPosition struct {
			Start int64
			End   int64
		}
	*/
	lcsCmd := c.cmdable.LCS(c.ctx, lcsQuery)
	if err := lcsCmd.Err(); err != nil {
		return object.NewError(err)
	}
	match := lcsCmd.Val()
	posSize := len(match.Matches)
	positions := make([]object.Object, posSize, posSize)
	for idx := 0; idx < posSize; idx++ {
		pos := match.Matches[idx]
		if lcsQuery.WithMatchLen {
			positions[idx] = object.NewMap(map[string]object.Object{
				"key1": object.NewMap(map[string]object.Object{
					"start": object.NewInt(pos.Key1.Start),
					"end":   object.NewInt(pos.Key1.End),
				}),
				"key2": object.NewMap(map[string]object.Object{
					"start": object.NewInt(pos.Key2.Start),
					"end":   object.NewInt(pos.Key2.End),
				}),
				"match_len": object.NewInt(pos.MatchLen),
			})
		} else {
			positions[idx] = object.NewMap(map[string]object.Object{
				"key1": object.NewMap(map[string]object.Object{
					"start": object.NewInt(pos.Key1.Start),
					"end":   object.NewInt(pos.Key1.End),
				}),
				"key2": object.NewMap(map[string]object.Object{
					"start": object.NewInt(pos.Key2.Start),
					"end":   object.NewInt(pos.Key2.End),
				}),
			})
		}
	}
	return object.NewMap(map[string]object.Object{
		"match_string": object.NewString(match.MatchString),
		"matches":      object.NewList(positions),
		"len":          object.NewInt(match.Len),
	})

}

func (c *RedisConn) MGet(ctx context.Context, args ...object.Object) object.Object {
	keys := make([]string, len(args))
	for i, arg := range args {
		key, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.mget() expected string arguments (got %s)", arg.Type())
		}
		keys[i] = key.Value()
	}
	// MGet(ctx context.Context, keys ...string) *SliceCmd
	stringSliceCmd := c.cmdable.MGet(c.ctx, keys...)
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	results := make([]object.Object, len(stringSliceCmd.Val()))
	for i, val := range stringSliceCmd.Val() {
		if val == nil {
			results[i] = object.Nil
		} else {
			results[i] = object.NewString(val.(string))
		}
	}
	return object.NewList(results)
}

func (c *RedisConn) MSet(ctx context.Context, args ...object.Object) object.Object {
	if len(args)%2 != 0 {
		return object.TypeErrorf("type error: redis.conn.mset() expected an even number of arguments (got %d)", len(args))
	}
	pairs := make([]interface{}, len(args))
	for i, arg := range args {
		str, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.mset() expected string arguments (got %s)", arg.Type())
		}
		pairs[i] = str.Value()
	}
	// MSet(ctx context.Context, values ...interface{}) *StatusCmd
	statusCmd := c.cmdable.MSet(c.ctx, pairs...)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.Val())
}

func (c *RedisConn) MGetFunc(ctx context.Context, args ...object.Object) object.Object {
	argsSize := len(args)
	if argsSize < 2 {
		return object.TypeErrorf("type error: redis.conn.mgetfunc() takes at least two arguments (%d given)", len(args))
	}
	fx, ok := args[0].(object.Callable)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.mgetfunc() expected a callable argument for fx (got %s)", args[0].Type())
	}
	keys := make([]string, argsSize-1)
	for i, arg := range args[1:] {
		key, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.mget() expected string arguments (got %s)", arg.Type())
		}
		keys[i] = key.Value()
	}
	// MGet(ctx context.Context, keys ...string) *SliceCmd
	stringSliceCmd := c.cmdable.MGet(c.ctx, keys...)
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	cnt := 0
	results := stringSliceCmd.Val()
	size := len(results)
	for idx := 0; idx < size; idx++ {
		if results[idx] == nil {
			continue
		}
		val := results[idx].(string)
		boolOrErr := fx.Call(ctx, object.NewString(val))
		cnt++
		if err, ok := boolOrErr.(*object.Error); ok {
			return object.NewError(err)
		} else if cont, ok := boolOrErr.(*object.Bool); ok {
			if !cont.Value() {
				break
			}
		} else {
			// Will be nil if the callable returns nil (or doesn't return)
			return object.TypeErrorf("type error: redis.conn.mgetfunc() expected a boolean response from fx (got %s)", boolOrErr.Type())
		}
	}
	return object.NewInt(int64(cnt))
}

func (c *RedisConn) MSetNX(ctx context.Context, args ...object.Object) object.Object {
	if len(args)%2 != 0 {
		return object.TypeErrorf("type error: redis.conn.msetnx() expected an even number of arguments (got %d)", len(args))
	}
	pairs := make([]interface{}, len(args))
	for i, arg := range args {
		str, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.msetnx() expected string arguments (got %s)", arg.Type())
		}
		pairs[i] = str.Value()
	}
	// MSetNX(ctx context.Context, values ...interface{}) *BoolCmd
	boolCmd := c.cmdable.MSetNX(c.ctx, pairs...)
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

// TODO:
//func (c *RedisConn) SetArgs(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) < 2 {
//		return object.TypeErrorf("type error: redis.conn.setargs() takes at least two arguments (%d given)", len(args))
//	}
//	key, ok := args[0].(*object.String)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.setargs() expected a string argument for key (got %s)", args[0].Type())
//	}
//	value, ok := args[1].(*object.String)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.setargs() expected a string argument for value (got %s)", args[1].Type())
//	}
//	setArgs := redis.SetArgs{
//		Key:   key.Value(),
//		Value: value.Value(),
//	}
//	if len(args) > 2 {
//		options, ok := args[2].(*object.Map)
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.setargs() expected a map argument for options (got %s)", args[2].Type())
//		}
//		for k, v := range options.Value() {
//			switch k {
//			case "ttl":
//				ttl, ok := v.(*object.Duration)
//				if !ok {
//					return object.TypeErrorf("type error: redis.conn.setargs() expected a duration argument for ttl (got %s)", v.Type())
//				}
//				setArgs.TTL = ttl.Value()
//			case "mode":
//				mode, ok := v.(*object.String)
//				if !ok {
//					return object.TypeErrorf("type error: redis.conn.setargs() expected a string argument for mode (got %s)", v.Type())
//				}
//				setArgs.Mode = mode.Value()
//			case "get":
//				get, ok := v.(*object.Bool)
//				if !ok {
//					return object.TypeErrorf("type error: redis.conn.setargs() expected a boolean argument for get (got %s)", v.Type())
//				}
//				setArgs.Get = get.Value()
//			}
//		}
//	}
//	statusCmd := c.cmdable.SetArgs(c.ctx, setArgs)
//	if err := statusCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	return object.NewString(statusCmd.Val())
//}

func (c *RedisConn) SetEx(ctx context.Context, args ...object.Object) object.Object {
	argSize := len(args)
	if argSize != 3 && argSize != 2 {
		return object.TypeErrorf("type error: redis.conn.setex() takes two or three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.setex() expected a string argument for key (got %s)", args[0].Type())
	}
	value, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.setex() expected a string argument for value (got %s)", args[1].Type())
	}
	var exp time.Duration
	if argSize == 2 {
		exp = 0 * time.Second
	} else {
		expiration, ok := args[2].(*object.Duration)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.setex() expected a duration argument for expiration (got %s)", args[2].Type())
		}
		exp = expiration.Value()
	}
	// SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
	statusCmd := c.cmdable.SetEx(c.ctx, key.Value(), value.Value(), exp)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.Val())
}

func (c *RedisConn) SetNX(ctx context.Context, args ...object.Object) object.Object {
	argSize := len(args)
	if argSize != 3 && argSize != 2 {
		return object.TypeErrorf("type error: redis.conn.setnx() takes two or three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.setnx() expected a string argument for key (got %s)", args[0].Type())
	}
	value, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.setnx() expected a string argument for value (got %s)", args[1].Type())
	}
	var exp time.Duration
	if argSize == 2 {
		exp = 0 * time.Second
	} else {
		expiration, ok := args[2].(*object.Duration)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.setex() expected a duration argument for expiration (got %s)", args[2].Type())
		}
		exp = expiration.Value()
	}

	// SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
	boolCmd := c.cmdable.SetNX(c.ctx, key.Value(), value.Value(), exp)
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) SetXX(ctx context.Context, args ...object.Object) object.Object {
	argSize := len(args)
	if argSize != 3 && argSize != 2 {
		return object.TypeErrorf("type error: redis.conn.setxx() takes two or three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.setxx() expected a string argument for key (got %s)", args[0].Type())
	}
	value, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.setxx() expected a string argument for value (got %s)", args[1].Type())
	}
	var exp time.Duration
	if argSize == 2 {
		exp = 0 * time.Second
	} else {
		expiration, ok := args[2].(*object.Duration)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.setex() expected a duration argument for expiration (got %s)", args[2].Type())
		}
		exp = expiration.Value()
	}

	// SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
	//boolCmd := c.cmdable.SetXX(c.ctx, key.Value(), value.Value(), exp)
	//if err := boolCmd.Err(); err != nil {
	//	return object.NewError(err)
	//}
	setArgs := redis.SetArgs{
		//Mode:    "XX",
		TTL:     exp,
		KeepTTL: true,
	}
	//SetArgs(ctx context.Context, key string, value interface{}, a SetArgs) *StatusCmd
	statusCmd := c.cmdable.SetArgs(c.ctx, key.Value(), value.Value(), setArgs)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}

	return object.NewBool("OK" == statusCmd.Val())
}

func (c *RedisConn) SetRange(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.setrange() takes exactly three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.setrange() expected a string argument for key (got %s)", args[0].Type())
	}
	offset, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.setrange() expected an int argument for offset (got %s)", args[1].Type())
	}
	value, ok := args[2].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.setrange() expected a string argument for value (got %s)", args[2].Type())
	}
	// SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd
	intCmd := c.cmdable.SetRange(c.ctx, key.Value(), offset.Value(), value.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) StrLen(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.strlen() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.strlen() expected a string argument for key (got %s)", args[0].Type())
	}
	// StrLen(ctx context.Context, key string) *IntCmd
	intCmd := c.cmdable.StrLen(c.ctx, key.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}
