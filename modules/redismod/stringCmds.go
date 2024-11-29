package redismod

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/risor-io/risor/object"
	"time"
)

/*
TODO: Append, Decr, DecrBy, Get, GetRange, GetSet, GetEx, GetDel, Incr, IncrBy, IncrByFloat, LCS, MGet, MSet, MSetNX, Set, SetArgs, SetEx, SetNX, SetXX, SetRange, StrLen
*/

func (c *RedisConn) StringCmdsGetAttr(name string) (object.Object, bool) {
	switch name {
	case "append":
		return object.NewBuiltin("redis.conn.append", c.Append), true
	case "set":
		return object.NewBuiltin("redis.conn.set", c.Set), true
	case "get":
		return object.NewBuiltin("redis.conn.get", c.Get), true

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
