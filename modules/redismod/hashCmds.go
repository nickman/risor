package redismod

import (
	"context"
	"github.com/risor-io/risor/object"
)

/*
	TODO:HScan, HScanNoValues, HRandField, HRandFieldWithValues, HExpire, HExpireWithArgs, HPExpire, HPExpireWithArgs, HExpireAt, HExpireAtWithArgs, HPExpireAt, HPExpireAtWithArgs, HPersist, HExpireTime, HPExpireTime, HTTL, HPTTL
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
	boolCmd := c.cmdable.HSetNX(c.ctx, key.Value(), field.Value(), value.Value())
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}
