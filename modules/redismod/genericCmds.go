package redismod

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/risor-io/risor/object"
	"time"
)

func (c *RedisConn) GenericCmdsGetAttr(name string) (object.Object, bool) {
	switch name {
	case "del":
		return object.NewBuiltin("redis.conn.conndel", c.Del), true
	case "dump":
		return object.NewBuiltin("redis.conn.conndump", c.Dump), true
	case "exists":
		return object.NewBuiltin("redis.conn.connexists", c.Exists), true
	case "expire":
		return object.NewBuiltin("redis.conn.connexpire", c.Expire), true
	case "expireAt":
		return object.NewBuiltin("redis.conn.connexpireAt", c.ExpireAt), true
	case "expireTime":
		return object.NewBuiltin("redis.conn.connexpireTime", c.ExpireTime), true
	case "expireNX":
		return object.NewBuiltin("redis.conn.connexpireNX", c.ExpireNX), true
	case "expireXX":
		return object.NewBuiltin("redis.conn.connexpireXX", c.ExpireXX), true
	case "expireGT":
		return object.NewBuiltin("redis.conn.connexpireGT", c.ExpireGT), true
	case "expireLT":
		return object.NewBuiltin("redis.conn.connexpireLT", c.ExpireLT), true
	case "keys":
		return object.NewBuiltin("redis.conn.connkeys", c.Keys), true
	case "migrate":
		return object.NewBuiltin("redis.conn.connmigrate", c.Migrate), true
	case "move":
		return object.NewBuiltin("redis.conn.connmove", c.Move), true
	case "objectFreq":
		return object.NewBuiltin("redis.conn.connobjectFreq", c.ObjectFreq), true
	case "objectRefCount":
		return object.NewBuiltin("redis.conn.connobjectRefCount", c.ObjectRefCount), true
	case "objectEncoding":
		return object.NewBuiltin("redis.conn.connobjectEncoding", c.ObjectEncoding), true
	case "objectIdleTime":
		return object.NewBuiltin("redis.conn.connobjectIdleTime", c.ObjectIdleTime), true
	case "persist":
		return object.NewBuiltin("redis.conn.connpersist", c.Persist), true
	case "pExpire":
		return object.NewBuiltin("redis.conn.connpExpire", c.PExpire), true
	case "pExpireAt":
		return object.NewBuiltin("redis.conn.connpExpireAt", c.PExpireAt), true
	case "pExpireTime":
		return object.NewBuiltin("redis.conn.connpExpireTime", c.PExpireTime), true
	case "pTTL":
		return object.NewBuiltin("redis.conn.connpTTL", c.PTTL), true
	case "randomKey":
		return object.NewBuiltin("redis.conn.connrandomKey", c.RandomKey), true
	case "rename":
		return object.NewBuiltin("redis.conn.connrename", c.Rename), true
	case "renameNX":
		return object.NewBuiltin("redis.conn.connrenameNX", c.RenameNX), true
	case "restore":
		return object.NewBuiltin("redis.conn.connrestore", c.Restore), true
	case "sort":
		return object.NewBuiltin("redis.conn.connsort", c.Sort), true
	case "ttl":
		return object.NewBuiltin("redis.conn.connttl", c.TTL), true
	case "type":
		return object.NewBuiltin("redis.conn.conntype", c.RedisType), true
	case "restoreReplace":
		return object.NewBuiltin("redis.conn.connrestoreReplace", c.RestoreReplace), true
	case "sortro":
		return object.NewBuiltin("redis.conn.connsortro", c.SortRO), true
	case "sortStore":
		return object.NewBuiltin("redis.conn.connsortStore", c.SortStore), true
	case "sortInterfaces":
		return object.NewBuiltin("redis.conn.connsortInterfaces", c.SortInterfaces), true
	case "touch":
		return object.NewBuiltin("redis.conn.conntouch", c.Touch), true
	case "copy":
		return object.NewBuiltin("redis.conn.conncopy", c.Copy), true
	}
	return nil, false
}

func (c *RedisConn) Del(ctx context.Context, args ...object.Object) object.Object {
	if len(args) == 0 {
		return object.TypeErrorf("type error: redis.conn.del() takes at least one argument (%d given)", len(args))
	}
	keys := make([]string, len(args))
	for i, arg := range args {
		key, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.del() expected string arguments (got %s)", arg.Type())
		}
		keys[i] = key.Value()
	}
	intCmd := c.cmdable.Del(c.ctx, keys...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) Dump(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.dump() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.dump() expected a string argument (got %s)", args[0].Type())
	}
	stringCmd := c.cmdable.Dump(c.ctx, key.Value())
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) Exists(ctx context.Context, args ...object.Object) object.Object {
	if len(args) == 0 {
		return object.TypeErrorf("type error: redis.conn.exists() takes at least one argument (%d given)", len(args))
	}
	keys := make([]string, len(args))
	for i, arg := range args {
		key, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.exists() expected string arguments (got %s)", arg.Type())
		}
		keys[i] = key.Value()
	}
	intCmd := c.cmdable.Exists(c.ctx, keys...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) Expire(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.expire() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expire() expected a string argument for key (got %s)", args[0].Type())
	}
	seconds, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expire() expected an int argument for seconds (got %s)", args[1].Type())
	}
	boolCmd := c.cmdable.Expire(c.ctx, key.Value(), time.Duration(seconds.Value())*time.Second)
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) ExpireAt(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.expireAt() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expireAt() expected a string argument for key (got %s)", args[0].Type())
	}
	timestamp, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expireAt() expected an int argument for timestamp (got %s)", args[1].Type())
	}
	boolCmd := c.cmdable.ExpireAt(c.ctx, key.Value(), time.Unix(timestamp.Value(), 0))
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) ExpireTime(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.expireTime() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expireTime() expected a string argument for key (got %s)", args[0].Type())
	}
	// ExpireTime(ctx context.Context, key string) *DurationCmd
	durCmd := c.cmdable.ExpireTime(c.ctx, key.Value())
	if err := durCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(int64(durCmd.Val().Seconds()))
}

func (c *RedisConn) ExpireNX(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.expireNX() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expireNX() expected a string argument for key (got %s)", args[0].Type())
	}
	seconds, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expireNX() expected an int argument for seconds (got %s)", args[1].Type())
	}
	boolCmd := c.cmdable.ExpireNX(c.ctx, key.Value(), time.Duration(seconds.Value())*time.Second)
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) ExpireXX(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.expireXX() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expireXX() expected a string argument for key (got %s)", args[0].Type())
	}
	seconds, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expireXX() expected an int argument for seconds (got %s)", args[1].Type())
	}
	boolCmd := c.cmdable.ExpireXX(c.ctx, key.Value(), time.Duration(seconds.Value())*time.Second)
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) ExpireGT(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.expireGT() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expireGT() expected a string argument for key (got %s)", args[0].Type())
	}
	seconds, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expireGT() expected an int argument for seconds (got %s)", args[1].Type())
	}
	boolCmd := c.cmdable.ExpireGT(c.ctx, key.Value(), time.Duration(seconds.Value())*time.Second)
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) ExpireLT(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.expireLT() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expireLT() expected a string argument for key (got %s)", args[0].Type())
	}
	seconds, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.expireLT() expected an int argument for seconds (got %s)", args[1].Type())
	}
	boolCmd := c.cmdable.ExpireLT(c.ctx, key.Value(), time.Duration(seconds.Value())*time.Second)
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) Keys(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.keys() takes exactly one argument (%d given)", len(args))
	}
	pattern, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.keys() expected a string argument (got %s)", args[0].Type())
	}
	stringSliceCmd := c.cmdable.Keys(c.ctx, pattern.Value())
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewStringList(stringSliceCmd.Val())
}

func (c *RedisConn) Migrate(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 6 {
		return object.TypeErrorf("type error: redis.conn.migrate() takes exactly five arguments (%d given)", len(args))
	}
	host, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.migrate() expected a string argument for host (got %s)", args[0].Type())
	}
	port, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.migrate() expected a string argument for port (got %s)", args[1].Type())
	}
	key, ok := args[2].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.migrate() expected a string argument for key (got %s)", args[2].Type())
	}
	destDB, ok := args[3].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.migrate() expected an int argument for destDB (got %s)", args[3].Type())
	}
	timeout, ok := args[4].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.migrate() expected a string argument for timeout (got %s)", args[4].Type())
	}
	timeoutDur, err := time.ParseDuration(timeout.Value())
	if err != nil {
		return object.TypeErrorf("type error: redis.conn.migrate() could not parse the string argument for timeout: expr=%s, err=%s", timeout.Value(), err)
	}
	// Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *StatusCmd
	stringCmd := c.cmdable.Migrate(c.ctx, host.Value(), port.Value(), key.Value(), int(destDB.Value()), timeoutDur)
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) Move(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.move() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.move() expected a string argument for key (got %s)", args[0].Type())
	}
	db, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.move() expected an int argument for db (got %s)", args[1].Type())
	}
	// Move(ctx context.Context, key string, db int) *BoolCmd
	boolCmd := c.cmdable.Move(c.ctx, key.Value(), int(db.Value()))
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) ObjectFreq(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.objectFreq() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.objectFreq() expected a string argument for key (got %s)", args[0].Type())
	}
	intCmd := c.cmdable.ObjectFreq(c.ctx, key.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) ObjectRefCount(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.objectRefCount() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.objectRefCount() expected a string argument for key (got %s)", args[0].Type())
	}
	intCmd := c.cmdable.ObjectRefCount(c.ctx, key.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) ObjectEncoding(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.objectEncoding() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.objectEncoding() expected a string argument for key (got %s)", args[0].Type())
	}
	stringCmd := c.cmdable.ObjectEncoding(c.ctx, key.Value())
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) ObjectIdleTime(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.objectIdleTime() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.objectIdleTime() expected a string argument for key (got %s)", args[0].Type())
	}
	durationCmd := c.cmdable.ObjectIdleTime(c.ctx, key.Value())
	if err := durationCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(int64(durationCmd.Val().Seconds()))
}

func (c *RedisConn) Persist(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.persist() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.persist() expected a string argument for key (got %s)", args[0].Type())
	}
	boolCmd := c.cmdable.Persist(c.ctx, key.Value())
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) PExpire(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.pExpire() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.pExpire() expected a string argument for key (got %s)", args[0].Type())
	}
	milliseconds, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.pExpire() expected an int argument for milliseconds (got %s)", args[1].Type())
	}
	boolCmd := c.cmdable.PExpire(c.ctx, key.Value(), time.Duration(milliseconds.Value())*time.Millisecond)
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) PExpireAt(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.pExpireAt() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.pExpireAt() expected a string argument for key (got %s)", args[0].Type())
	}
	timestamp, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.pExpireAt() expected an int argument for timestamp (got %s)", args[1].Type())
	}
	boolCmd := c.cmdable.PExpireAt(c.ctx, key.Value(), time.Unix(0, timestamp.Value()*int64(time.Millisecond)))
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) PExpireTime(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.pExpireTime() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.pExpireTime() expected a string argument for key (got %s)", args[0].Type())
	}
	// PExpireTime(ctx context.Context, key string) *DurationCmd
	durCmd := c.cmdable.PExpireTime(c.ctx, key.Value())
	if err := durCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewDuration(durCmd.Val())
}

func (c *RedisConn) PTTL(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.pTTL() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.pTTL() expected a string argument for key (got %s)", args[0].Type())
	}
	// PTTL(ctx context.Context, key string) *DurationCmd
	durCmd := c.cmdable.PTTL(c.ctx, key.Value())
	if err := durCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewDuration(durCmd.Val())
}

func (c *RedisConn) RandomKey(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.randomKey() takes no arguments (%d given)", len(args))
	}
	stringCmd := c.cmdable.RandomKey(c.ctx)
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) Rename(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.rename() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.rename() expected a string argument for key (got %s)", args[0].Type())
	}
	newKey, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.rename() expected a string argument for newKey (got %s)", args[1].Type())
	}
	statusCmd := c.cmdable.Rename(c.ctx, key.Value(), newKey.Value())
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) RenameNX(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.renameNX() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.renameNX() expected a string argument for key (got %s)", args[0].Type())
	}
	newKey, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.renameNX() expected a string argument for newKey (got %s)", args[1].Type())
	}
	boolCmd := c.cmdable.RenameNX(c.ctx, key.Value(), newKey.Value())
	if err := boolCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewBool(boolCmd.Val())
}

func (c *RedisConn) Restore(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.restore() takes exactly three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.restore() expected a string argument for key (got %s)", args[0].Type())
	}
	ttl, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.restore() expected an int argument for ttl (got %s)", args[1].Type())
	}
	serializedValue, ok := args[2].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.restore() expected a string argument for serializedValue (got %s)", args[2].Type())
	}
	statusCmd := c.cmdable.Restore(c.ctx, key.Value(), time.Duration(ttl.Value())*time.Millisecond, serializedValue.Value())
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) RestoreReplace(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.restoreReplace() takes exactly three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.restoreReplace() expected a string argument for key (got %s)", args[0].Type())
	}
	ttl, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.restoreReplace() expected an int argument for ttl (got %s)", args[1].Type())
	}
	serializedValue, ok := args[2].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.restoreReplace() expected a string argument for serializedValue (got %s)", args[2].Type())
	}
	statusCmd := c.cmdable.RestoreReplace(c.ctx, key.Value(), time.Duration(ttl.Value())*time.Millisecond, serializedValue.Value())
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) Sort(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 {
		return object.TypeErrorf("type error: redis.conn.sort() takes at least one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.sort() expected a string argument for key (got %s)", args[0].Type())
	}
	sort := &redis.Sort{}
	if len(args) > 1 {
		sortArgs, ok := args[1].(*object.Map)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.sort() expected a map argument for sort options (got %s)", args[1].Type())
		}
		// Parse sort options from the map
		for k, v := range sortArgs.Value() {
			switch k {
			case "by":
				sort.By = v.(*object.String).Value()
			case "offset":
				sort.Offset = v.(*object.Int).Value()
			case "count":
				sort.Count = v.(*object.Int).Value()
			case "get":
				sort.Get = append(sort.Get, v.(*object.String).Value())
			case "order":
				sort.Order = v.(*object.String).Value()
			case "alpha":
				sort.Alpha = v.(*object.Bool).Value()
			}
		}
	}
	stringSliceCmd := c.cmdable.Sort(c.ctx, key.Value(), sort)
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewStringList(stringSliceCmd.Val())
}

func (c *RedisConn) SortRO(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 {
		return object.TypeErrorf("type error: redis.conn.sortRO() takes at least one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.sortRO() expected a string argument for key (got %s)", args[0].Type())
	}
	sort := &redis.Sort{}
	if len(args) > 1 {
		sortArgs, ok := args[1].(*object.Map)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.sortRO() expected a map argument for sort options (got %s)", args[1].Type())
		}
		// Parse sort options from the map
		for k, v := range sortArgs.Value() {
			switch k {
			case "by":
				sort.By = v.(*object.String).Value()
			case "offset":
				sort.Offset = v.(*object.Int).Value()
			case "count":
				sort.Count = v.(*object.Int).Value()
			case "get":
				sort.Get = append(sort.Get, v.(*object.String).Value())
			case "order":
				sort.Order = v.(*object.String).Value()
			case "alpha":
				sort.Alpha = v.(*object.Bool).Value()
			}
		}
	}
	stringSliceCmd := c.cmdable.SortRO(c.ctx, key.Value(), sort)
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewStringList(stringSliceCmd.Val())
}

func (c *RedisConn) SortStore(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.sortStore() takes at least two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.sortStore() expected a string argument for key (got %s)", args[0].Type())
	}
	store, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.sortStore() expected a string argument for store (got %s)", args[1].Type())
	}
	sort := &redis.Sort{}
	if len(args) > 2 {
		sortArgs, ok := args[2].(*object.Map)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.sortStore() expected a map argument for sort options (got %s)", args[2].Type())
		}
		// Parse sort options from the map
		for k, v := range sortArgs.Value() {
			switch k {
			case "by":
				sort.By = v.(*object.String).Value()
			case "offset":
				sort.Offset = v.(*object.Int).Value()
			case "count":
				sort.Count = v.(*object.Int).Value()
			case "get":
				sort.Get = append(sort.Get, v.(*object.String).Value())
			case "order":
				sort.Order = v.(*object.String).Value()
			case "alpha":
				sort.Alpha = v.(*object.Bool).Value()
			}
		}
	}
	intCmd := c.cmdable.SortStore(c.ctx, key.Value(), store.Value(), sort)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) SortInterfaces(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 {
		return object.TypeErrorf("type error: redis.conn.sortInterfaces() takes at least one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.sortInterfaces() expected a string argument for key (got %s)", args[0].Type())
	}
	sort := &redis.Sort{}
	if len(args) > 1 {
		sortArgs, ok := args[1].(*object.Map)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.sortInterfaces() expected a map argument for sort options (got %s)", args[1].Type())
		}
		// Parse sort options from the map
		for k, v := range sortArgs.Value() {
			switch k {
			case "by":
				sort.By = v.(*object.String).Value()
			case "offset":
				sort.Offset = v.(*object.Int).Value()
			case "count":
				sort.Count = v.(*object.Int).Value()
			case "get":
				sort.Get = append(sort.Get, v.(*object.String).Value())
			case "order":
				sort.Order = v.(*object.String).Value()
			case "alpha":
				sort.Alpha = v.(*object.Bool).Value()
			}
		}
	}
	// SortInterfaces(ctx context.Context, key string, sort *Sort) *SliceCmd
	sliceCmd := c.cmdable.SortInterfaces(c.ctx, key.Value(), sort)
	if err := sliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	ifaces := sliceCmd.Val()
	size := len(ifaces)
	objects := make([]object.Object, size, size)
	for idx := 0; idx < size; idx++ {
		objects[idx] = object.FromGoType(ifaces[idx])
	}
	return object.NewList(objects)
}

func (c *RedisConn) Touch(ctx context.Context, args ...object.Object) object.Object {
	if len(args) == 0 {
		return object.TypeErrorf("type error: redis.conn.touch() takes at least one argument (%d given)", len(args))
	}
	keys := make([]string, len(args))
	for i, arg := range args {
		key, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.touch() expected string arguments (got %s)", arg.Type())
		}
		keys[i] = key.Value()
	}
	intCmd := c.cmdable.Touch(c.ctx, keys...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) TTL(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.ttl() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.ttl() expected a string argument (got %s)", args[0].Type())
	}
	durationCmd := c.cmdable.TTL(c.ctx, key.Value())
	if err := durationCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(int64(durationCmd.Val().Seconds()))
}

func (c *RedisConn) RedisType(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.type() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.type() expected a string argument (got %s)", args[0].Type())
	}
	// Type(ctx context.Context, key string) *StatusCmd
	statusCmd := c.cmdable.Type(c.ctx, key.Value())
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.Val())
}

func (c *RedisConn) Copy(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 4 {
		return object.TypeErrorf("type error: redis.conn.copy() takes four arguments (%d given)", len(args))
	}
	sourceKey, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.copy() expected a string argument for sourceKey (got %s)", args[0].Type())
	}
	destKey, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.copy() expected a string argument for destKey (got %s)", args[1].Type())
	}
	db, ok := args[2].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.copy() expected an int argument for db (got %s)", args[2].Type())
	}
	replace, ok := args[3].(*object.Bool)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.copy() expected a bool argument for replace (got %s)", args[3].Type())
	}
	// Copy(ctx context.Context, sourceKey string, destKey string, db int, replace bool) *IntCmd
	intCmd := c.cmdable.Copy(c.ctx, sourceKey.Value(), destKey.Value(), int(db.Value()), replace.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

//func (c *RedisConn) Scan(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) != 3 {
//		return object.TypeErrorf("type error: redis.conn.scan() takes three arguments (%d given)", len(args))
//	}
//	cursor, ok := args[0].(*object.Int)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.scan() expected an int argument for cursor (got %s)", args[0].Type())
//	}
//	match := ""
//	count := int64(0)
//	if len(args) > 1 {
//		match, ok = args[1].(*object.String).Value()
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.scan() expected a string argument for match (got %s)", args[1].Type())
//		}
//	}
//	if len(args) > 2 {
//		count, ok = args[2].(*object.Int).Value()
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.scan() expected an int argument for count (got %s)", args[2].Type())
//		}
//	}
//	// Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd
//	scanCmd := c.cmdable.Scan(c.ctx, uint64(cursor.Value()), match, count)
//	if err := scanCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	return object.NewScanResult(scanCmd.Val())
//}
//
//func (c *RedisConn) ScanType(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) < 1 {
//		return object.TypeErrorf("type error: redis.conn.scanType() takes at least one argument (%d given)", len(args))
//	}
//	cursor, ok := args[0].(*object.Int)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.scanType() expected an int argument for cursor (got %s)", args[0].Type())
//	}
//	match := ""
//	count := int64(0)
//	keyType := ""
//	if len(args) > 1 {
//		match, ok = args[1].(*object.String).Value()
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.scanType() expected a string argument for match (got %s)", args[1].Type())
//		}
//	}
//	if len(args) > 2 {
//		count, ok = args[2].(*object.Int).Value()
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.scanType() expected an int argument for count (got %s)", args[2].Type())
//		}
//	}
//	if len(args) > 3 {
//		keyType, ok = args[3].(*object.String).Value()
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.scanType() expected a string argument for keyType (got %s)", args[3].Type())
//		}
//	}
//	scanCmd := c.cmdable.ScanType(c.ctx, uint64(cursor.Value()), match, count, keyType)
//	if err := scanCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	return object.NewScanResult(scanCmd.Val())
//}
