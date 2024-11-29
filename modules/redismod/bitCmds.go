package redismod

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/risor-io/risor/object"
)

/*
TODO: GetBit, SetBit, BitCount, BitOpAnd, BitOpOr, BitOpXor, BitOpNot, BitPos, BitPosSpan, BitField, BitFieldRO
*/

func (c *RedisConn) BitCmdsGetAttr(name string) (object.Object, bool) {
	switch name {
	case "getbit":
		return object.NewBuiltin("redis.conn.getbit", c.GetBit), true
	case "setbit":
		return object.NewBuiltin("redis.conn.setbit", c.SetBit), true
	case "bitcount":
		return object.NewBuiltin("redis.conn.bitcount", c.BitCount), true
	case "bitopand":
		return object.NewBuiltin("redis.conn.bitopand", c.BitOpAnd), true
	case "bitopor":
		return object.NewBuiltin("redis.conn.bitopor", c.BitOpOr), true
	case "bitopxor":
		return object.NewBuiltin("redis.conn.bitopxor", c.BitOpXor), true
	case "bitopnot":
		return object.NewBuiltin("redis.conn.bitopnot", c.BitOpNot), true
	case "bitpos":
		return object.NewBuiltin("redis.conn.bitpos", c.BitPos), true
	case "bitposspan":
		return object.NewBuiltin("redis.conn.bitposspan", c.BitPosSpan), true
		// TODO: BitField, BitFieldRO (not sure how these work)
		//case "bitfield":
		//	return object.NewBuiltin("redis.conn.bitfield", c.BitField), true
		//case "bitfieldro":
		//	return object.NewBuiltin("redis.conn.bitfieldro", c.BitFieldRO), true
	}
	return nil, false
}

func (c *RedisConn) GetBit(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.getbit() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.getbit() expected a string argument for key (got %s)", args[0].Type())
	}
	offset, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.getbit() expected an int argument for offset (got %s)", args[1].Type())
	}
	// GetBit(ctx context.Context, key string, offset int64) *IntCmd
	intCmd := c.cmdable.GetBit(c.ctx, key.Value(), offset.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) SetBit(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.setbit() takes exactly three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.setbit() expected a string argument for key (got %s)", args[0].Type())
	}
	offset, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.setbit() expected an int argument for offset (got %s)", args[1].Type())
	}
	value, ok := args[2].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.setbit() expected an int argument for value (got %s)", args[2].Type())
	}
	// SetBit(ctx context.Context, key string, offset int64, value int) *IntCmd
	intCmd := c.cmdable.SetBit(c.ctx, key.Value(), offset.Value(), int(value.Value()))
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) BitCount(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 || len(args) > 3 {
		return object.TypeErrorf("type error: redis.conn.bitcount() takes one to three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitcount() expected a string argument for key (got %s)", args[0].Type())
	}
	var start, end int64
	if len(args) > 1 {
		startArg, ok := args[1].(*object.Int)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.bitcount() expected an int argument for start (got %s)", args[1].Type())
		}
		start = startArg.Value()
	}
	if len(args) > 2 {
		endArg, ok := args[2].(*object.Int)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.bitcount() expected an int argument for end (got %s)", args[2].Type())
		}
		end = endArg.Value()
	}
	intCmd := c.cmdable.BitCount(c.ctx, key.Value(), &redis.BitCount{Start: start, End: end})
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) BitOpAnd(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.bitopand() takes at least two arguments (%d given)", len(args))
	}
	destKey, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitopand() expected a string argument for destKey (got %s)", args[0].Type())
	}
	keys := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		key, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.bitopand() expected string arguments for keys (got %s)", arg.Type())
		}
		keys[i] = key.Value()
	}
	intCmd := c.cmdable.BitOpAnd(c.ctx, destKey.Value(), keys...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) BitOpOr(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.bitopor() takes at least two arguments (%d given)", len(args))
	}
	destKey, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitopor() expected a string argument for destKey (got %s)", args[0].Type())
	}
	keys := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		key, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.bitopor() expected string arguments for keys (got %s)", arg.Type())
		}
		keys[i] = key.Value()
	}
	intCmd := c.cmdable.BitOpOr(c.ctx, destKey.Value(), keys...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) BitOpXor(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.bitopxor() takes at least two arguments (%d given)", len(args))
	}
	destKey, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitopxor() expected a string argument for destKey (got %s)", args[0].Type())
	}
	keys := make([]string, len(args)-1)
	for i, arg := range args[1:] {
		key, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.bitopxor() expected string arguments for keys (got %s)", arg.Type())
		}
		keys[i] = key.Value()
	}
	intCmd := c.cmdable.BitOpXor(c.ctx, destKey.Value(), keys...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) BitOpNot(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.bitopnot() takes exactly two arguments (%d given)", len(args))
	}
	destKey, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitopnot() expected a string argument for destKey (got %s)", args[0].Type())
	}
	key, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitopnot() expected a string argument for key (got %s)", args[1].Type())
	}
	intCmd := c.cmdable.BitOpNot(c.ctx, destKey.Value(), key.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) BitPos(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 || len(args) > 4 {
		return object.TypeErrorf("type error: redis.conn.bitpos() takes two to four arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitpos() expected a string argument for key (got %s)", args[0].Type())
	}
	bit, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitpos() expected an int argument for bit (got %s)", args[1].Type())
	}
	var start, end int64
	if len(args) > 2 {
		startArg, ok := args[2].(*object.Int)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.bitpos() expected an int argument for start (got %s)", args[2].Type())
		}
		start = startArg.Value()
	}
	if len(args) > 3 {
		endArg, ok := args[3].(*object.Int)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.bitpos() expected an int argument for end (got %s)", args[3].Type())
		}
		end = endArg.Value()
	}
	intCmd := c.cmdable.BitPos(c.ctx, key.Value(), bit.Value(), start, end)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) BitPosSpan(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 5 {
		return object.TypeErrorf("type error: redis.conn.bitposspan() takes exactly five arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitposspan() expected a string argument for key (got %s)", args[0].Type())
	}
	bit, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitposspan() expected an int argument for bit (got %s)", args[1].Type())
	}
	start, ok := args[2].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitposspan() expected an int argument for start (got %s)", args[2].Type())
	}
	var end int64
	endArg, ok := args[3].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitposspan() expected an int argument for end (got %s)", args[3].Type())
	}
	end = endArg.Value()
	var span string
	spanArg, ok := args[4].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.bitposspan() expected a string argument for span (got %s)", args[4].Type())
	}
	span = spanArg.Value()
	// BitPosSpan(ctx context.Context, key string, bit int8, start, end int64, span string) *IntCmd
	intCmd := c.cmdable.BitPosSpan(c.ctx, key.Value(), int8(bit.Value()), start.Value(), end, span)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

//func (c *RedisConn) BitField(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) < 2 {
//		return object.TypeErrorf("type error: redis.conn.bitfield() takes at least two arguments (%d given)", len(args))
//	}
//	key, ok := args[0].(*object.String)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.bitfield() expected a string argument for key (got %s)", args[0].Type())
//	}
//
//	bitFieldArgs := make([]redis.BitFieldArgs, len(args)-1)
//	for i, arg := range args[1:] {
//		bitFieldArg, ok := arg.(*object.String)
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.bitfield() expected string arguments for bitfield args (got %s)", arg.Type())
//		}
//		bitFieldArgs[i] = redis.BitFieldArgs(bitFieldArg.Value())
//	}
//	// BitField(ctx context.Context, key string, values ...interface{}) *IntSliceCmd
//	intSliceCmd := c.cmdable.BitField(c.ctx, key.Value(), bitFieldArgs...)
//	if err := intSliceCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	results := make([]object.Object, len(intSliceCmd.Val()))
//	for i, val := range intSliceCmd.Val() {
//		results[i] = object.NewInt(val)
//	}
//	return object.NewList(results)
//}
//
//func (c *RedisConn) BitFieldRO(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) < 2 {
//		return object.TypeErrorf("type error: redis.conn.bitfieldro() takes at least two arguments (%d given)", len(args))
//	}
//	key, ok := args[0].(*object.String)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.bitfieldro() expected a string argument for key (got %s)", args[0].Type())
//	}
//	bitFieldArgs := make([]redis.BitFieldArgs, len(args)-1)
//	for i, arg := range args[1:] {
//		bitFieldArg, ok := arg.(*object.String)
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.bitfieldro() expected string arguments for bitfield args (got %s)", arg.Type())
//		}
//		bitFieldArgs[i] = redis.BitFieldArgs(bitFieldArg.Value())
//	}
//	intSliceCmd := c.cmdable.BitFieldRO(c.ctx, key.Value(), bitFieldArgs...)
//	if err := intSliceCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	results := make([]object.Object, len(intSliceCmd.Val()))
//	for i, val := range intSliceCmd.Val() {
//		results[i] = object.NewInt(val)
//	}
//	return object.NewList(results)
//}
