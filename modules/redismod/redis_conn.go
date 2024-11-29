package redismod

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/risor-io/risor/arg"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/op"
	"sync"
	"time"
)

const REDIS_CONN = object.Type("redis.conn")

type BaseClient interface {
	redis.Cmdable
	Close() error
}

type RedisConn struct {
	ctx context.Context
	//opts    *redis.Options
	client  BaseClient
	cmdable redis.Cmdable
	once    sync.Once
	closed  chan bool
}

func New(ctx context.Context, client BaseClient) *RedisConn {
	obj := &RedisConn{
		ctx: ctx,
		//opts:   client.Options(),
		client:  client,
		cmdable: client,
		closed:  make(chan bool),
	}
	obj.waitToClose()
	return obj
}

func (c *RedisConn) waitToClose() {
	go func() {
		select {
		case <-c.closed:
		case <-c.ctx.Done():
			c.client.Close()
		}
	}()
}

func (c *RedisConn) Close() error {
	var err error
	c.once.Do(func() {
		err = c.client.Close()
		close(c.closed)
	})
	return err
}

func (c *RedisConn) Type() object.Type {
	return REDIS_CONN
}

func (c *RedisConn) Inspect() string {
	return "redis.conn()"
}

func (c *RedisConn) Interface() interface{} {
	return c.client
}

func (c *RedisConn) Value() BaseClient {
	return c.client
}

func (c *RedisConn) Equals(other object.Object) object.Object {
	return object.NewBool(c == other)
}

func (c *RedisConn) IsTruthy() bool {
	return true
}

func (c *RedisConn) GetAttr(name string) (object.Object, bool) {
	switch name {
	case "close":
		fmt.Printf("Closing...\n")
		return object.NewBuiltin("redis.conn.close", func(ctx context.Context, args ...object.Object) object.Object {
			if err := arg.Require("redis.conn.close", 0, args); err != nil {
				return err
			}
			if err := c.Close(); err != nil {
				return object.NewError(err)
			}
			fmt.Printf("Closed\n")
			return object.Nil
		}), true
	case "isnull":
		return object.NewBuiltin("redis.conn.isnull", c.IsNull), true

	case "ping":
		return object.NewBuiltin("redis.conn.ping", c.Ping), true
	case "hset":
		return object.NewBuiltin("redis.conn.hset", c.HSet), true
	case "hget":
		return object.NewBuiltin("redis.conn.hget", c.HGet), true
	case "incr":
		return object.NewBuiltin("redis.conn.incr", c.Incr), true
	case "decr":
		return object.NewBuiltin("redis.conn.decr", c.Decr), true
	case "lpush":
		return object.NewBuiltin("redis.conn.lpush", c.LPush), true
	case "rpush":
		return object.NewBuiltin("redis.conn.rpush", c.RPush), true
	case "brpop":
		return object.NewBuiltin("redis.conn.brpop", c.BRPop), true
	case "mget":
		return object.NewBuiltin("redis.conn.mget", c.MGet), true
	case "hkeys":
		return object.NewBuiltin("redis.conn.hkeys", c.HKeys), true
	case "hlen":
		return object.NewBuiltin("redis.conn.hlen", c.HLen), true
	case "hvals":
		return object.NewBuiltin("redis.conn.hvals", c.HVals), true
	case "lpop":
		return object.NewBuiltin("redis.conn.lpop", c.LPop), true
	case "rpop":
		return object.NewBuiltin("redis.conn.rpop", c.RPop), true
	case "srem":
		return object.NewBuiltin("redis.conn.srem", c.SRem), true
	case "zrem":
		return object.NewBuiltin("redis.conn.zrem", c.ZRem), true
	case "zscore":
		return object.NewBuiltin("redis.conn.zscore", c.ZScore), true
	case "hdel":
		return object.NewBuiltin("redis.conn.hdel", c.HDel), true
	case "hmget":
		return object.NewBuiltin("redis.conn.hmget", c.HMGet), true
	case "hmset":
		return object.NewBuiltin("redis.conn.hmset", c.HMSet), true
	case "sadd":
		return object.NewBuiltin("redis.conn.sadd", c.SAdd), true
	case "smembers":
		return object.NewBuiltin("redis.conn.smembers", c.SMembers), true
	case "zadd":
		return object.NewBuiltin("redis.conn.zadd", c.ZAdd), true
	case "zrange":
		return object.NewBuiltin("redis.conn.zrange", c.ZRange), true
	case "clientgetname":
		return object.NewBuiltin("redis.conn.clientgetname", c.ClientGetName), true
	case "echo":
		return object.NewBuiltin("redis.conn.echo", c.Echo), true
	case "quit":
		return object.NewBuiltin("redis.conn.quit", c.Quit), true
	case "unlink":
		return object.NewBuiltin("redis.conn.unlink", c.Unlink), true
	case "bgrewriteaof":
		return object.NewBuiltin("redis.conn.bgrewriteaof", c.BgRewriteAOF), true
	case "bgsave":
		return object.NewBuiltin("redis.conn.bgsave", c.BgSave), true
	case "clientkill":
		return object.NewBuiltin("redis.conn.clientkill", c.ClientKill), true
	case "clientkillbyfilter":
		return object.NewBuiltin("redis.conn.clientkillbyfilter", c.ClientKillByFilter), true
	case "clientlist":
		return object.NewBuiltin("redis.conn.clientlist", c.ClientList), true
	case "clientinfo":
		return object.NewBuiltin("redis.conn.clientinfo", c.ClientInfo), true
	case "clientpause":
		return object.NewBuiltin("redis.conn.clientpause", c.ClientPause), true
	case "clientunpause":
		return object.NewBuiltin("redis.conn.clientunpause", c.ClientUnpause), true
	case "clientid":
		return object.NewBuiltin("redis.conn.clientid", c.ClientID), true
	case "clientunblock":
		return object.NewBuiltin("redis.conn.clientunblock", c.ClientUnblock), true
	case "clientunblockwitherror":
		return object.NewBuiltin("redis.conn.clientunblockwitherror", c.ClientUnblockWithError), true
	case "configget":
		return object.NewBuiltin("redis.conn.configget", c.ConfigGet), true
	case "configresetstat":
		return object.NewBuiltin("redis.conn.configresetstat", c.ConfigResetStat), true
	case "configset":
		return object.NewBuiltin("redis.conn.configset", c.ConfigSet), true
	case "configrewrite":
		return object.NewBuiltin("redis.conn.configrewrite", c.ConfigRewrite), true
	case "dbsize":
		return object.NewBuiltin("redis.conn.dbsize", c.DBSize), true
	case "flushall":
		return object.NewBuiltin("redis.conn.flushall", c.FlushAll), true
	case "flushallasync":
		return object.NewBuiltin("redis.conn.flushallasync", c.FlushAllAsync), true
	case "flushdb":
		return object.NewBuiltin("redis.conn.flushdb", c.FlushDB), true
	case "flushdbasync":
		return object.NewBuiltin("redis.conn.flushdbasync", c.FlushDBAsync), true
	case "info":
		return object.NewBuiltin("redis.conn.info", c.Info), true
	case "lastsave":
		return object.NewBuiltin("redis.conn.lastsave", c.LastSave), true
	case "save":
		return object.NewBuiltin("redis.conn.save", c.Save), true
	case "shutdown":
		return object.NewBuiltin("redis.conn.shutdown", c.Shutdown), true
	case "shutdownsave":
		return object.NewBuiltin("redis.conn.shutdownsave", c.ShutdownSave), true
	case "shutdownnosave":
		return object.NewBuiltin("redis.conn.shutdownnosave", c.ShutdownNoSave), true
	case "slaveof":
		return object.NewBuiltin("redis.conn.slaveof", c.SlaveOf), true
	case "slowlogget":
		return object.NewBuiltin("redis.conn.slowlogget", c.SlowLogGet), true
	case "time":
		return object.NewBuiltin("redis.conn.time", c.Time), true
	case "debugobject":
		return object.NewBuiltin("redis.conn.debugobject", c.DebugObject), true
	case "memoryusage":
		return object.NewBuiltin("redis.conn.memoryusage", c.MemoryUsage), true
	case "moduleloadex":
		return object.NewBuiltin("redis.conn.moduleloadex", c.ModuleLoadex), true
	}
	if attr, ok := c.JsonCmdsGetAttr(name); ok {
		return attr, ok
	}
	if attr, ok := c.GenericCmdsGetAttr(name); ok {
		return attr, ok
	}
	if attr, ok := c.StringCmdsGetAttr(name); ok {
		return attr, ok
	}

	return nil, false
}

func (c *RedisConn) Ping(ctx context.Context, args ...object.Object) object.Object {
	statusCmd := c.cmdable.Ping(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	} else {
		return object.NewString(statusCmd.String())
	}
}

func (c *RedisConn) SetAttr(name string, value object.Object) error {
	return object.TypeErrorf("type error: redis.conn object has no attribute %q", name)
}

func (c *RedisConn) RunOperation(opType op.BinaryOpType, right object.Object) object.Object {
	return object.TypeErrorf("type error: unsupported operation for redis.conn: %v", opType)
}

func (c *RedisConn) Cost() int {
	return 8 // FIXME: Can we do better than this ?
}

func (c *RedisConn) IsNull(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.isnull() takes exactly one arguments (%d given)", len(args))
	}
	return object.NewBool(args[0] == object.Nil)
}

func (c *RedisConn) HSet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.hset() takes exactly three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hset() expected a string argument for key (got %s)", args[0].Type())
	}
	field, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hset() expected a string argument for field (got %s)", args[1].Type())
	}
	value, ok := args[2].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hset() expected a string argument for value (got %s)", args[2].Type())
	}
	statusCmd := c.cmdable.HSet(c.ctx, key.Value(), field.Value(), value.Value())
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(statusCmd.Val())
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
	statusCmd := c.cmdable.HGet(c.ctx, key.Value(), field.Value())
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.Val())
}

//func (c *RedisConn) Del(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) < 1 {
//		return object.TypeErrorf("type error: redis.conn.del() takes one or more arguments (%d given)", len(args))
//	}
//	var keys []string
//	for _, arg := range args {
//		key, ok := arg.(*object.String)
//		if !ok {
//			return object.TypeErrorf("type error: redis.conn.del() expected string arguments (got %s)", arg.Type())
//		}
//		keys = append(keys, key.Value())
//	}
//	statusCmd := c.cmdable.Del(c.ctx, keys...)
//	if err := statusCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	return object.NewInt(int64(statusCmd.Val()))
//}

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

func (c *RedisConn) LPush(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.lpush() takes at least two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.lpush() expected a string argument for key (got %s)", args[0].Type())
	}
	var values []interface{}
	for _, arg := range args[1:] {
		value, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.lpush() expected string arguments for values (got %s)", arg.Type())
		}
		values = append(values, value.Value())
	}
	intCmd := c.cmdable.LPush(c.ctx, key.Value(), values...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) RPush(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.rpush() takes at least two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.rpush() expected a string argument for key (got %s)", args[0].Type())
	}
	var values []interface{}
	for _, arg := range args[1:] {
		value, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.rpush() expected string arguments for values (got %s)", arg.Type())
		}
		values = append(values, value.Value())
	}
	intCmd := c.cmdable.RPush(c.ctx, key.Value(), values...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) BRPop(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.brpop() takes at least two arguments (%d given)", len(args))
	}
	var keys []string
	for _, arg := range args[:len(args)-1] {
		key, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.brpop() expected string arguments for keys (got %s)", arg.Type())
		}
		keys = append(keys, key.Value())
	}
	timeout, ok := args[len(args)-1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.brpop() expected an int argument for timeout (got %s)", args[len(args)-1].Type())
	}
	stringSliceCmd := c.cmdable.BRPop(c.ctx, time.Duration(timeout.Value())*time.Second, keys...)
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	var results []object.Object
	for _, val := range stringSliceCmd.Val() {
		results = append(results, object.NewString(val))
	}
	return object.NewList(results)
}

//func (c *RedisConn) Expire(ctx context.Context, args ...object.Object) object.Object {
//	if len(args) != 2 {
//		return object.TypeErrorf("type error: redis.conn.expire() takes exactly two arguments (%d given)", len(args))
//	}
//	key, ok := args[0].(*object.String)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.expire() expected a string argument for key (got %s)", args[0].Type())
//	}
//	seconds, ok := args[1].(*object.Int)
//	if !ok {
//		return object.TypeErrorf("type error: redis.conn.expire() expected an int argument for seconds (got %s)", args[1].Type())
//	}
//	boolCmd := c.cmdable.Expire(c.ctx, key.Value(), time.Duration(seconds.Value())*time.Second)
//	if err := boolCmd.Err(); err != nil {
//		return object.NewError(err)
//	}
//	return object.NewBool(boolCmd.Val())
//}

func (c *RedisConn) MGet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 {
		return object.TypeErrorf("type error: redis.conn.mget() takes at least one argument (%d given)", len(args))
	}
	var keys []string
	for _, arg := range args {
		key, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.mget() expected string arguments for keys (got %s)", arg.Type())
		}
		keys = append(keys, key.Value())
	}
	stringSliceCmd := c.cmdable.MGet(c.ctx, keys...)
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	var results []object.Object
	for _, val := range stringSliceCmd.Val() {
		results = append(results, object.FromGoType(val))
	}
	return object.NewList(results)
}

func (c *RedisConn) HKeys(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.hkeys() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hkeys() expected a string argument (got %s)", args[0].Type())
	}
	stringSliceCmd := c.cmdable.HKeys(c.ctx, key.Value())
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	var results []object.Object
	for _, val := range stringSliceCmd.Val() {
		results = append(results, object.NewString(val))
	}
	return object.NewList(results)
}

func (c *RedisConn) HLen(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.hlen() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hlen() expected a string argument (got %s)", args[0].Type())
	}
	intCmd := c.cmdable.HLen(c.ctx, key.Value())
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
		return object.TypeErrorf("type error: redis.conn.hvals() expected a string argument (got %s)", args[0].Type())
	}
	stringSliceCmd := c.cmdable.HVals(c.ctx, key.Value())
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	var results []object.Object
	for _, val := range stringSliceCmd.Val() {
		results = append(results, object.NewString(val))
	}
	return object.NewList(results)
}

func (c *RedisConn) LPop(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.lpop() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.lpop() expected a string argument (got %s)", args[0].Type())
	}
	stringCmd := c.cmdable.LPop(c.ctx, key.Value())
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) RPop(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.rpop() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.rpop() expected a string argument (got %s)", args[0].Type())
	}
	stringCmd := c.cmdable.RPop(c.ctx, key.Value())
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) SRem(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.srem() takes at least two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.srem() expected a string argument for key (got %s)", args[0].Type())
	}
	var members []interface{}
	for _, arg := range args[1:] {
		members = append(members, arg.Interface())
	}
	intCmd := c.cmdable.SRem(c.ctx, key.Value(), members...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) ZRem(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.zrem() takes at least two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.zrem() expected a string argument for key (got %s)", args[0].Type())
	}
	var members []interface{}
	for _, arg := range args[1:] {
		members = append(members, arg.Interface())
	}
	intCmd := c.cmdable.ZRem(c.ctx, key.Value(), members...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) ZScore(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.zscore() takes exactly two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.zscore() expected a string argument for key (got %s)", args[0].Type())
	}
	member, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.zscore() expected a string argument for member (got %s)", args[1].Type())
	}
	floatCmd := c.cmdable.ZScore(c.ctx, key.Value(), member.Value())
	if err := floatCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewFloat(floatCmd.Val())
}

func (c *RedisConn) HDel(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.hdel() takes at least two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hdel() expected a string argument for key (got %s)", args[0].Type())
	}
	var fields []string
	for _, arg := range args[1:] {
		field, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hdel() expected string arguments for fields (got %s)", arg.Type())
		}
		fields = append(fields, field.Value())
	}
	intCmd := c.cmdable.HDel(c.ctx, key.Value(), fields...)
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
	var fields []string
	for _, arg := range args[1:] {
		field, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hmget() expected string arguments for fields (got %s)", arg.Type())
		}
		fields = append(fields, field.Value())
	}
	sliceCmd := c.cmdable.HMGet(c.ctx, key.Value(), fields...)
	if err := sliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	var results []object.Object
	for _, val := range sliceCmd.Val() {
		results = append(results, object.FromGoType(val))
	}
	return object.NewList(results)
}

func (c *RedisConn) HMSet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 3 || len(args)%2 != 1 {
		return object.TypeErrorf("type error: redis.conn.hmset() takes an odd number of arguments (at least three, with key and field-value pairs)")
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.hmset() expected a string argument for key (got %s)", args[0].Type())
	}
	fields := make(map[string]interface{})
	for i := 1; i < len(args); i += 2 {
		field, ok := args[i].(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.hmset() expected string arguments for fields (got %s)", args[i].Type())
		}
		value := args[i+1].Interface()
		fields[field.Value()] = value
	}
	statusCmd := c.cmdable.HMSet(c.ctx, key.Value(), fields)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) SAdd(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 2 {
		return object.TypeErrorf("type error: redis.conn.sadd() takes at least two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.sadd() expected a string argument for key (got %s)", args[0].Type())
	}
	var members []interface{}
	for _, arg := range args[1:] {
		members = append(members, arg.Interface())
	}
	intCmd := c.cmdable.SAdd(c.ctx, key.Value(), members...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) SMembers(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.smembers() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.smembers() expected a string argument for key (got %s)", args[0].Type())
	}
	stringSliceCmd := c.cmdable.SMembers(c.ctx, key.Value())
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	var results []object.Object
	for _, val := range stringSliceCmd.Val() {
		results = append(results, object.NewString(val))
	}
	return object.NewList(results)
}

func (c *RedisConn) ZAdd(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 3 || len(args)%2 != 1 {
		return object.TypeErrorf("type error: redis.conn.zadd() takes an odd number of arguments (at least three, with key and score-member pairs)")
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.zadd() expected a string argument for key (got %s)", args[0].Type())
	}
	var members []redis.Z
	for i := 1; i < len(args); i += 2 {
		score, ok := args[i].(*object.Float)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.zadd() expected float arguments for scores (got %s)", args[i].Type())
		}
		member, ok := args[i+1].(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.zadd() expected string arguments for members (got %s)", args[i+1].Type())
		}
		members = append(members, redis.Z{Score: score.Value(), Member: member.Value()})
	}
	intCmd := c.cmdable.ZAdd(c.ctx, key.Value(), members...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) ZRange(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 3 {
		return object.TypeErrorf("type error: redis.conn.zrange() takes exactly three arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.zrange() expected a string argument for key (got %s)", args[0].Type())
	}
	start, ok := args[1].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.zrange() expected an int argument for start (got %s)", args[1].Type())
	}
	stop, ok := args[2].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.zrange() expected an int argument for stop (got %s)", args[2].Type())
	}
	stringSliceCmd := c.cmdable.ZRange(c.ctx, key.Value(), start.Value(), stop.Value())
	if err := stringSliceCmd.Err(); err != nil {
		return object.NewError(err)
	}
	var results []object.Object
	for _, val := range stringSliceCmd.Val() {
		results = append(results, object.NewString(val))
	}
	return object.NewList(results)
}

func (c *RedisConn) ClientGetName(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.clientGetName() takes no arguments (%d given)", len(args))
	}
	stringCmd := c.cmdable.ClientGetName(c.ctx)
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) Echo(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.echo() takes exactly one argument (%d given)", len(args))
	}
	message, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.echo() expected a string argument (got %s)", args[0].Type())
	}
	stringCmd := c.cmdable.Echo(c.ctx, message.Value())
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) Quit(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.quit() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.Quit(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) Unlink(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 {
		return object.TypeErrorf("type error: redis.conn.unlink() takes at least one argument (%d given)", len(args))
	}
	var keys []string
	for _, arg := range args {
		key, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.unlink() expected string arguments (got %s)", arg.Type())
		}
		keys = append(keys, key.Value())
	}
	intCmd := c.cmdable.Unlink(c.ctx, keys...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) BgRewriteAOF(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.bgRewriteAOF() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.BgRewriteAOF(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) BgSave(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.bgSave() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.BgSave(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) ClientKill(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.clientKill() takes exactly one argument (%d given)", len(args))
	}
	ipPort, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.clientKill() expected a string argument (got %s)", args[0].Type())
	}
	strCmd := c.cmdable.ClientKill(c.ctx, ipPort.Value())
	if err := strCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(strCmd.Val())
}

func (c *RedisConn) ClientKillByFilter(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 {
		return object.TypeErrorf("type error: redis.conn.clientKillByFilter() takes at least one argument (%d given)", len(args))
	}
	size := len(args)
	options := make([]string, 0, size)
	for idx := 0; idx < size; idx++ {
		arg, ok := args[idx].(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.clientKillByFilter() expected string arguments (got %s)", args[idx].Type())
		}
		options = append(options, arg.Value())
	}
	intCmd := c.cmdable.ClientKillByFilter(c.ctx, options...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) ClientList(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.clientList() takes no arguments (%d given)", len(args))
	}
	stringCmd := c.cmdable.ClientList(c.ctx)
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) ClientInfo(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.clientInfo() takes no arguments (%d given)", len(args))
	}
	cInfo := c.cmdable.ClientInfo(c.ctx)
	if err := cInfo.Err(); err != nil {
		return object.NewError(err)
	}
	clientInfo := cInfo.Val()
	cInfoItems := map[string]object.Object{
		"id":                 object.NewInt(clientInfo.ID),
		"addr":               object.NewString(clientInfo.Addr),
		"laddr":              object.NewString(clientInfo.LAddr),
		"fd":                 object.NewInt(clientInfo.FD),
		"name":               object.NewString(clientInfo.Name),
		"age":                object.NewInt(int64(clientInfo.Age.Seconds())),
		"idle":               object.NewInt(int64(clientInfo.Idle.Seconds())),
		"flags":              object.NewInt(int64(clientInfo.Flags)),
		"db":                 object.NewInt(int64(clientInfo.DB)),
		"sub":                object.NewInt(int64(clientInfo.Sub)),
		"psub":               object.NewInt(int64(clientInfo.PSub)),
		"ssub":               object.NewInt(int64(clientInfo.SSub)),
		"multi":              object.NewInt(int64(clientInfo.Multi)),
		"watch":              object.NewInt(int64(clientInfo.Watch)),
		"querybuf":           object.NewInt(int64(clientInfo.QueryBuf)),
		"querybuffree":       object.NewInt(int64(clientInfo.QueryBufFree)),
		"argvmem":            object.NewInt(int64(clientInfo.ArgvMem)),
		"multimem":           object.NewInt(int64(clientInfo.MultiMem)),
		"buffersize":         object.NewInt(int64(clientInfo.BufferSize)),
		"bufferpeak":         object.NewInt(int64(clientInfo.BufferPeak)),
		"outputbufferlength": object.NewInt(int64(clientInfo.OutputBufferLength)),
		"outputlistlength":   object.NewInt(int64(clientInfo.OutputListLength)),
		"outputmemory":       object.NewInt(int64(clientInfo.OutputMemory)),
		"totalmemory":        object.NewInt(int64(clientInfo.TotalMemory)),
		"events":             object.NewString(clientInfo.Events),
		"lastcmd":            object.NewString(clientInfo.LastCmd),
		"user":               object.NewString(clientInfo.User),
		"redir":              object.NewInt(clientInfo.Redir),
		"resp":               object.NewInt(int64(clientInfo.Resp)),
		"libname":            object.NewString(clientInfo.LibName),
		"libver":             object.NewString(clientInfo.LibVer),
	}
	return object.NewMap(cInfoItems)
}

func (c *RedisConn) ClientPause(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.clientPause() takes exactly one argument (%d given)", len(args))
	}
	duration, ok := args[0].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.clientPause() expected an int argument (got %s)", args[0].Type())
	}
	statusCmd := c.cmdable.ClientPause(c.ctx, time.Duration(duration.Value())*time.Millisecond)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) ClientUnpause(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.clientUnpause() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.ClientUnpause(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) ClientID(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.clientID() takes no arguments (%d given)", len(args))
	}
	intCmd := c.cmdable.ClientID(c.ctx)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) ClientUnblock(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.clientUnblock() takes exactly one argument (%d given)", len(args))
	}
	id, ok := args[0].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.clientUnblock() expected an int argument (got %s)", args[0].Type())
	}
	intCmd := c.cmdable.ClientUnblock(c.ctx, id.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) ClientUnblockWithError(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.clientUnblockWithError() takes exactly one argument (%d given)", len(args))
	}
	id, ok := args[0].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.clientUnblockWithError() expected an int argument (got %s)", args[0].Type())
	}
	intCmd := c.cmdable.ClientUnblock(c.ctx, id.Value())
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) ConfigGet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.configGet() takes exactly one argument (%d given)", len(args))
	}
	parameter, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.configGet() expected a string argument (got %s)", args[0].Type())
	}
	mapCmd := c.cmdable.ConfigGet(c.ctx, parameter.Value())
	if err := mapCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.FromGoType(mapCmd.Val())
}

func (c *RedisConn) ConfigResetStat(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.configResetStat() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.ConfigResetStat(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) ConfigSet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.configSet() takes exactly two arguments (%d given)", len(args))
	}
	parameter, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.configSet() expected a string argument for parameter (got %s)", args[0].Type())
	}
	value, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.configSet() expected a string argument for value (got %s)", args[1].Type())
	}
	statusCmd := c.cmdable.ConfigSet(c.ctx, parameter.Value(), value.Value())
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) ConfigRewrite(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.configRewrite() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.ConfigRewrite(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) DBSize(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.dbSize() takes no arguments (%d given)", len(args))
	}
	intCmd := c.cmdable.DBSize(c.ctx)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) FlushAll(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.flushAll() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.FlushAll(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) FlushAllAsync(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.flushAllAsync() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.FlushAllAsync(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) FlushDB(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.flushDB() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.FlushDB(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) FlushDBAsync(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.flushDBAsync() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.FlushDBAsync(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) Info(ctx context.Context, args ...object.Object) object.Object {
	var sections []string
	for _, arg := range args {
		section, ok := arg.(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.info() expected string arguments for sections (got %s)", arg.Type())
		}
		sections = append(sections, section.Value())
	}
	stringCmd := c.cmdable.Info(c.ctx, sections...)
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) LastSave(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.lastSave() takes no arguments (%d given)", len(args))
	}
	intCmd := c.cmdable.LastSave(c.ctx)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) Save(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.save() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.Save(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) Shutdown(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.shutdown() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.Shutdown(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) ShutdownSave(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.shutdownSave() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.ShutdownSave(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) ShutdownNoSave(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.shutdownNoSave() takes no arguments (%d given)", len(args))
	}
	statusCmd := c.cmdable.ShutdownNoSave(c.ctx)
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) SlaveOf(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.TypeErrorf("type error: redis.conn.slaveOf() takes exactly two arguments (%d given)", len(args))
	}
	host, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.slaveOf() expected a string argument for host (got %s)", args[0].Type())
	}
	port, ok := args[1].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.slaveOf() expected a string argument for port (got %s)", args[1].Type())
	}
	statusCmd := c.cmdable.SlaveOf(c.ctx, host.Value(), port.Value())
	if err := statusCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(statusCmd.String())
}

func (c *RedisConn) SlowLogGet(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.slowLogGet() takes exactly one argument (%d given)", len(args))
	}
	num, ok := args[0].(*object.Int)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.slowLogGet() expected an int argument (got %s)", args[0].Type())
	}
	slowLogCmd := c.cmdable.SlowLogGet(c.ctx, num.Value())
	if err := slowLogCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.FromGoType(slowLogCmd.Val())
}

func (c *RedisConn) Time(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.TypeErrorf("type error: redis.conn.time() takes no arguments (%d given)", len(args))
	}
	timeCmd := c.cmdable.Time(c.ctx)
	if err := timeCmd.Err(); err != nil {
		return object.NewError(err)
	}
	t := timeCmd.Val()
	return object.NewTime(t)
}

func (c *RedisConn) DebugObject(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.conn.debugObject() takes exactly one argument (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.debugObject() expected a string argument (got %s)", args[0].Type())
	}
	stringCmd := c.cmdable.DebugObject(c.ctx, key.Value())
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}

func (c *RedisConn) MemoryUsage(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return object.TypeErrorf("type error: redis.conn.memoryUsage() takes one or two arguments (%d given)", len(args))
	}
	key, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.memoryUsage() expected a string argument for key (got %s)", args[0].Type())
	}
	var samples []int
	if len(args) == 2 {
		sample, ok := args[1].(*object.Int)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.memoryUsage() expected an int argument for samples (got %s)", args[1].Type())
		}
		samples = append(samples, int(sample.Value()))
	}
	intCmd := c.cmdable.MemoryUsage(c.ctx, key.Value(), samples...)
	if err := intCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewInt(intCmd.Val())
}

func (c *RedisConn) ModuleLoadex(ctx context.Context, args ...object.Object) object.Object {
	if len(args) < 1 {
		return object.TypeErrorf("type error: redis.conn.moduleLoadex() takes at least one argument (%d given)", len(args))
	}
	path, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.conn.moduleLoadex() expected a string argument for path (got %s)", args[0].Type())
	}
	conf := &redis.ModuleLoadexConfig{Path: path.Value()}
	for i := 1; i < len(args); i += 2 {
		if i+1 >= len(args) {
			return object.TypeErrorf("type error: redis.conn.moduleLoadex() expected pairs of config name and value")
		}
		name, ok := args[i].(*object.String)
		if !ok {
			return object.TypeErrorf("type error: redis.conn.moduleLoadex() expected a string argument for config name (got %s)", args[i].Type())
		}
		value := args[i+1].Interface()
		conf.Conf[name.Value()] = value
	}
	stringCmd := c.cmdable.ModuleLoadex(c.ctx, conf)
	if err := stringCmd.Err(); err != nil {
		return object.NewError(err)
	}
	return object.NewString(stringCmd.Val())
}
