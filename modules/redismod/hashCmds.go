package redismod

import (
	"context"
	"github.com/risor-io/risor/object"
	"log"
)

/*
	TODO: HPExpire, HPExpireWithArgs, HExpireAt, HExpireAtWithArgs, HPExpireAt, HPExpireAtWithArgs, HPersist, HExpireTime, HPExpireTime, HTTL, HPTTL
*/

func (c *RedisConn) HashCmdsGetAttr(name string) (object.Object, bool) {
	switch name {
	case "hset":
		return object.NewBuiltin("redis.conn.hset", c.HSet), true
	case "hget":
		return object.NewBuiltin("redis.conn.hget", c.HGet), true
	case "hkeys":
		return object.NewBuiltin("redis.conn.hkeys", c.HKeys), true
	case "hlen":
		return object.NewBuiltin("redis.conn.hlen", c.HLen), true
	case "hvals":
		return object.NewBuiltin("redis.conn.hvals", c.HVals), true
	case "hdel":
		return object.NewBuiltin("redis.conn.hdel", c.HDel), true
	case "hmget":
		return object.NewBuiltin("redis.conn.hmget", c.HMGet), true
	case "hmset":
		return object.NewBuiltin("redis.conn.hmset", c.HMSet), true
	case "hexists":
		return object.NewBuiltin("redis.conn.hexists", c.HExists), true
	case "hgetall":
		return object.NewBuiltin("redis.conn.hgetall", c.HGetAll), true
	case "hincrby":
		return object.NewBuiltin("redis.conn.hincrby", c.HIncrBy), true
	case "hincrbyfloat":
		return object.NewBuiltin("redis.conn.hincrbyfloat", c.HIncrByFloat), true
	case "hsetnx":
		return object.NewBuiltin("redis.conn.hsetnx", c.HSetNX), true
	case "hscan":
		return object.NewBuiltin("redis.conn.hscan", c.HScan), true
	case "hscanfunc":
		return object.NewBuiltin("redis.conn.hscanfunc", c.HScanFunc), true
	case "hscanfuncbatch":
		return object.NewBuiltin("redis.conn.hscanfuncbatch", c.HScanFuncBatch), true

	}
	return nil, false
}

func (c *RedisConn) HDel(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.hdel() takes at least two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hdel() expected a string argument for key (got %s)", args[0].Type())
	}
	fields := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		field, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hdel() expected string arguments for fields (got %s)", arg.Type())
		}
		fields[i] = field.Value()
	}
	intCmd := c.cmdable.HDel(c.ctx, key.Value(), fields...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) HVals(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.hvals() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hvals() expected a string argument for key (got %s)", args[0].Type())
	}
	stringSliceCmd := c.cmdable.HVals(c.ctx, key.Value())
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	results := make([]object.Object, len(stringSliceCmd.Val()))
	for i, val := range stringSliceCmd.Val() {
		results[i] = object.NewString(val)
	}
	return object.NewList(results)
}

func (c *RedisConn) HExists(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.hexists() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hexists() expected a string argument for key (got %s)", args[0].Type())
	}
	field, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hexists() expected a string argument for field (got %s)", args[1].Type())
	}
	boolCmd := c.cmdable.HExists(c.ctx, key.Value(), field.Value())
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) HGet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.hget() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hget() expected a string argument for key (got %s)", args[0].Type())
	}
	field, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hget() expected a string argument for field (got %s)", args[1].Type())
	}
	stringCmd := c.cmdable.HGet(c.ctx, key.Value(), field.Value())
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) HGetAll(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.hgetall() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hgetall() expected a string argument for key (got %s)", args[0].Type())
	}
	stringStringMapCmd := c.cmdable.HGetAll(c.ctx, key.Value())
	if err := stringStringMapCmd.Err(); err != nil {
		return object.NewError(err)
	}
	result := make(map[string]object.Object)
	for k, v := range stringStringMapCmd.Val() {
		result[k] = object.NewString(v)
	}
	return object.NewMap(result)
}

func (c *RedisConn) HIncrBy(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.hincrby() takes exactly three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hincrby() expected a string argument for key (got %s)", args[0].Type())
	}
	field, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hincrby() expected a string argument for field (got %s)", args[1].Type())
	}
	increment, ok := args[2].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hincrby() expected an int argument for increment (got %s)", args[2].Type())
	}
	// HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd
	intCmd := c.cmdable.HIncrBy(c.ctx, key.Value(), field.Value(), increment.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) HIncrByFloat(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.hincrbyfloat() takes exactly three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hincrbyfloat() expected a string argument for key (got %s)", args[0].Type())
	}
	field, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hincrbyfloat() expected a string argument for field (got %s)", args[1].Type())
	}
	increment, ok := args[2].(*object.Float)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hincrbyfloat() expected a float argument for increment (got %s)", args[2].Type())
	}
	floatCmd := c.cmdable.HIncrByFloat(c.ctx, key.Value(), field.Value(), increment.Value())
	if err := floatCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewFloat(floatCmd.Val())
}

func (c *RedisConn) HKeys(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.hkeys() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hkeys() expected a string argument for key (got %s)", args[0].Type())
	}
	// HKeys(ctx context.Context, key string) *StringSliceCmd
	stringSliceCmd := c.cmdable.HKeys(c.ctx, key.Value())
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	results := make([]object.Object, len(stringSliceCmd.Val()))
	for i, val := range stringSliceCmd.Val() {
		results[i] = object.NewString(val)
	}
	return object.NewList(results)
}

func (c *RedisConn) HLen(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.hlen() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hlen() expected a string argument for key (got %s)", args[0].Type())
	}
	intCmd := c.cmdable.HLen(c.ctx, key.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) HMGet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.hmget() takes at least two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hmget() expected a string argument for key (got %s)", args[0].Type())
	}
	fields := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		field, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hmget() expected string arguments for fields (got %s)", arg.Type())
		}
		fields[i] = field.Value()
	}
	// HMGet(ctx context.Context, key string, fields ...string) *SliceCmd
	stringSliceCmd := c.cmdable.HMGet(c.ctx, key.Value(), fields...)
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

func (c *RedisConn) HSet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 3 || len(args)%2 != 1 {
		return object.TypeErrorf("type error: redis.conn.hset() takes an odd number of arguments (at least three) (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hset() expected a string argument for key (got %s)", args[0].Type())
	}
	pairs := make([]interface{}, len(args)-1)
	for i, arg := range args[1:] {
		str, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hset() expected string arguments for field-value pairs (got %s)", arg.Type())
		}
		pairs[i] = str.Value()
	}
	// HSet(ctx context.Context, key string, values ...interface{}) *IntCmd
	intCmd := c.cmdable.HSet(c.ctx, key.Value(), pairs...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) HMSet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 3 || len(args)%2 != 1 {
		return object.TypeErrorf("type error: redis.conn.hmset() takes an odd number of arguments (at least three) (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hmset() expected a string argument for key (got %s)", args[0].Type())
	}
	pairs := make([]interface{}, len(args)-1)
	for i, arg := range args[1:] {
		str, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hmset() expected string arguments for field-value pairs (got %s)", arg.Type())
		}
		pairs[i] = str.Value()
	}
	// HMSet(ctx context.Context, key string, values ...interface{}) *BoolCmd
	boolCmd := c.cmdable.HMSet(c.ctx, key.Value(), pairs...)
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) HSetNX(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.hsetnx() takes exactly three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hsetnx() expected a string argument for key (got %s)", args[0].Type())
	}
	field, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hsetnx() expected a string argument for field (got %s)", args[1].Type())
	}
	value, ok := args[2].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hsetnx() expected a string argument for value (got %s)", args[2].Type())
	}
	// HSetNX(ctx context.Context, key, field string, value interface{}) *BoolCmd
	boolCmd := c.cmdable.HSetNX(c.ctx, key.Value(), field.Value(), value.Value())
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) HScan(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.hscan() takes at least two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hscan() expected a string argument for key (got %s)", args[0].Type())
	}
	cursor, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hscan() expected an int argument for cursor (got %s)", args[1].Type())
	}
	var match string
	var count int64
	if len(args) > 2 {
		matchArg, ok := args[2].(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hscan() expected a string argument for match (got %s)", args[2].Type())
		}
		match = matchArg.Value()
	}
	if len(args) > 3 {
		countArg, ok := args[3].(*object.Int)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hscan() expected an int argument for count (got %s)", args[3].Type())
		}
		count = countArg.Value()
	}
	if match == "" {
		match = "*"
	}
	// HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	// HSCAN key cursor [MATCH pattern] [COUNT count] [NOVALUES]
	scanCmd := c.cmdable.HScan(c.ctx, key.Value(), uint64(cursor.Value()), match, count)
	if err := scanCmd.Err(); err != nil {
		return object.NewError(err)
	}
	// (keys []string, cursor uint64)
	scanKeys, outCursor := scanCmd.Val()
	size := len(scanKeys)
	resultsMap := make(map[string]object.Object, size)
	for idx := 0; idx < size; idx++ {
		k := scanKeys[idx]
		idx++
		v := object.NewString(scanKeys[idx])
		resultsMap[k] = v
	}
	scanPage := object.NewMap(resultsMap)
	return object.NewList([]object.Object{object.NewInt(int64(outCursor)), scanPage})
}

func (c *RedisConn) HScanFunc(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 3 {
		return object.TypeErrorf("type error: redis.conn.hscanfunc() takes at least three arguments (%d given)", len(args))
	}
	callbackFx, ok := args[0].(object.Callable)
	if !ok {
		return object.TypeErrorf("type error: pgx.conn.hscanfunc() expected a callable argument for fx (got %s)", args[0].Type())
	}
	key, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hscanfunc() expected a string argument for key (got %s)", args[1].Type())
	}
	cursor, ok := args[2].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hscanfunc() expected an int argument for cursor (got %s)", args[2].Type())
	}
	var match string
	var count int64
	if len(args) > 3 {
		matchArg, ok := args[3].(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hscanfunc() expected a string argument for match (got %s)", args[3].Type())
		}
		match = matchArg.Value()
	}
	if len(args) > 4 {
		countArg, ok := args[4].(*object.Int)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hscanfunc() expected an int argument for count (got %s)", args[4].Type())
		}
		count = countArg.Value()
	}
	if match == "" {
		match = "*"
	}
	// HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	// HSCAN key cursor [MATCH pattern] [COUNT count] [NOVALUES]
	keyValPairCount := 0
	pageCount := 0
	rcursor := uint64(cursor.Value())
outer:
	for {
		scanCmd := c.cmdable.HScan(c.ctx, key.Value(), rcursor, match, count)
		if err := scanCmd.Err(); err != nil {
			return object.NewError(err)
		}
		pageCount++
		scanKeys, outCursor := scanCmd.Val()
		rcursor = outCursor
		size := len(scanKeys)

		for idx := 0; idx < size; idx++ {
			k := object.NewString(scanKeys[idx])
			idx++
			v := object.NewString(scanKeys[idx])
			keyValPairCount++
			callbackContinue, callbackErr := DoCallback(ctx, callbackFx, k, v)
			if callbackErr != nil {
				return object.NewError(callbackErr)
			} else {
				if !callbackContinue {
					break outer
				} else {
					continue
				}
			}
		}
		if outCursor == 0 {
			break
		}
	}
	log.Printf("HScanFunc: pageCount=%d, keyValPairCount=%d\n", pageCount, keyValPairCount)
	return object.NewList([]object.Object{object.NewInt(int64(pageCount)), object.NewInt(int64(keyValPairCount))})
}

func (c *RedisConn) HScanFuncBatch(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 3 {
		return object.TypeErrorf("type error: redis.conn.hscanfuncbatch() takes at least three arguments (%d given)", len(args))
	}
	callbackFx, ok := args[0].(object.Callable)
	if !ok {
		return object.TypeErrorf("type error: pgx.conn.hscanfuncbatch() expected a callable argument for fx (got %s)", args[0].Type())
	}
	key, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hscanfuncbatch() expected a string argument for key (got %s)", args[1].Type())
	}
	cursor, ok := args[2].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hscanfuncbatch() expected an int argument for cursor (got %s)", args[2].Type())
	}
	var match string
	var count int64
	if len(args) > 3 {
		matchArg, ok := args[3].(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hscanfuncbatch() expected a string argument for match (got %s)", args[3].Type())
		}
		match = matchArg.Value()
	}
	if len(args) > 4 {
		countArg, ok := args[4].(*object.Int)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hscanfuncbatch() expected an int argument for count (got %s)", args[4].Type())
		}
		count = countArg.Value()
	}
	if match == "" {
		match = "*"
	}
	// HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	// HSCAN key cursor [MATCH pattern] [COUNT count] [NOVALUES]
	keyValPairCount := 0
	pageCount := 0
	rcursor := uint64(cursor.Value())
	for {
		scanCmd := c.cmdable.HScan(c.ctx, key.Value(), rcursor, match, count)
		if err := scanCmd.Err(); err != nil {
			return object.NewError(err)
		}
		pageCount++
		scanKeys, outCursor := scanCmd.Val()
		rcursor = outCursor
		size := len(scanKeys)
		resultsMap := make(map[string]object.Object, size)
		for idx := 0; idx < size; idx++ {
			k := scanKeys[idx]
			idx++
			v := object.NewString(scanKeys[idx])
			resultsMap[k] = v
			keyValPairCount++
		}
		scanPage := object.NewMap(resultsMap)
		callbackContinue, callbackErr := DoCallback(ctx, callbackFx, scanPage)
		if callbackErr != nil {
			return object.NewError(callbackErr)
		} else {
			if !callbackContinue {
				break
			} else {
				continue
			}
		}
	}
	log.Printf("HScanFuncBatch: pageCount=%d, keyValPairCount=%d\n", pageCount, keyValPairCount)
	return object.NewList([]object.Object{object.NewInt(int64(pageCount)), object.NewInt(int64(keyValPairCount))})
}

//func (c *RedisConn) HScanNoValues(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) < 2 {
//		return object.TypeErrorf("type error: redis.conn.hscannovalues() takes at least two arguments (%d given)", len(args))
//	}
//	key, ok := args[0].(*object.String)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.hscannovalues() expected a string argument for key (got %s)", args[0].Type())
//	}
//	cursor, ok := args[1].(*object.Int)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.hscannovalues() expected an int argument for cursor (got %s)", args[1].Type())
//	}
//	var match string
//	var count int64
//	if len(args) > 2 {
//		matchArg, ok := args[2].(*object.String)
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.hscannovalues() expected a string argument for match (got %s)", args[2].Type())
//		}
//		match = matchArg.Value()
//	}
//	if len(args) > 3 {
//		countArg, ok := args[3].(*object.Int)
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.hscannovalues() expected an int argument for count (got %s)", args[3].Type())
//		}
//		count = countArg.Value()
//	}
//	scanCmd := c.cmdable.HScan(c.ctx, key.Value(), cursor.Value(), match, count)
//	if err := scanCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	results := make([]object.Object, len(scanCmd.Val())/2)
//	for i := 0; i < len(scanCmd.Val()); i += 2 {
//		results[i/2] = object.NewString(scanCmd.Val()[i])
//	}
//	return object.NewList(results)
//}
//
//func (c *RedisConn) HRandField(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) < 1 || len(args) > 2 {
//		return object.TypeErrorf("type error: redis.conn.hrandfield() takes one or two arguments (%d given)", len(args))
//	}
//	key, ok := args[0].(*object.String)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.hrandfield() expected a string argument for key (got %s)", args[0].Type())
//	}
//	var count int64
//	if len(args) == 2 {
//		countArg, ok := args[1].(*object.Int)
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.hrandfield() expected an int argument for count (got %s)", args[1].Type())
//		}
//		count = countArg.Value()
//	}
//	stringSliceCmd := c.cmdable.HRandField(c.ctx, key.Value(), count)
//	if err := stringSliceCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	results := make([]object.Object, len(stringSliceCmd.Val()))
//	for i, val := range stringSliceCmd.Val() {
//		results[i] = object.NewString(val)
//	}
//	return object.NewList(results)
//}
//
//func (c *RedisConn) HRandFieldWithValues(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) != 2 {
//		return object.TypeErrorf("type error: redis.conn.hrandfieldwithvalues() takes exactly two arguments (%d given)", len(args))
//	}
//	key, ok := args[0].(*object.String)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.hrandfieldwithvalues() expected a string argument for key (got %s)", args[0].Type())
//	}
//	count, ok := args[1].(*object.Int)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.hrandfieldwithvalues() expected an int argument for count (got %s)", args[1].Type())
//	}
//	stringSliceCmd := c.cmdable.HRandFieldWithValues(c.ctx, key.Value(), count.Value())
//	if err := stringSliceCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	results := make([]object.Object, len(stringSliceCmd.Val()))
//	for i, val := range stringSliceCmd.Val() {
//		results[i] = object.NewString(val)
//	}
//	return object.NewList(results)
//}
//
//func (c *RedisConn) HExpire(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) != 2 {
//		return object.TypeErrorf("type error: redis.conn.hexpire() takes exactly two arguments (%d given)", len(args))
//	}
//	key, ok := args[0].(*object.String)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.hexpire() expected a string argument for key (got %s)", args[0].Type())
//	}
//	seconds, ok := args[1].(*object.Int)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.hexpire() expected an int argument for seconds (got %s)", args[1].Type())
//	}
//	boolCmd := c.cmdable.HExpire(c.ctx, key.Value(), seconds.Value())
//	if err := boolCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	return object.NewBool(boolCmd.Val())
//}
//
//func (c *RedisConn) HExpireWithArgs(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) < 2 {
//		return object.TypeErrorf("type error: redis.conn.hexpirewithargs() takes at least two arguments (%d given)", len(args))
//	}
//	key, ok := args[0].(*object.String)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.hexpirewithargs() expected a string argument for key (got %s)", args[0].Type())
//	}
//	seconds, ok := args[1].(*object.Int)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.hexpirewithargs() expected an int argument for seconds (got %s)", args[1].Type())
//	}
//	var options []redis.SetArgs
//	for _, arg := range args[2:] {
//		option, ok := arg.(*object.String)
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.hexpirewithargs() expected string arguments for options (got %s)", arg.Type())
//		}
//		options = append(options, redis.SetArgs(option.Value()))
//	}
//	boolCmd := c.cmdable.HExpireWithArgs(c.ctx, key.Value(), seconds.Value(), options...)
//	if err := boolCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	return object.NewBool(boolCmd.Val())
//}
