package redismod

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/risor-io/risor/object"
	"net/url"
	"strings"
)

/*
   NewClient/Options/*Client/:client
   NewFailoverClient/FailoverOptions/*Client/:failover
   NewRing/RingOptions/*Ring/:ring
	NewClusterClient/ClusterOptions/*ClusterClient/:cluster

	NewFailoverClusterClient / NewFailoverClient / NewFailoverClusterClient / NewSentinelClient



	deferConnect
*/

func Connect(ctx context.Context, args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.TypeErrorf("type error: redis.connect() takes exactly one argument (%d given)", len(args))
	}
	redisurl, ok := args[0].(*object.String)
	if !ok {
		return object.TypeErrorf("type error: redis.connect() expected a string argument (got %s)", args[0].Type())
	}
	baseUrl, err := url.Parse(redisurl.Value())
	if err != nil {
		return object.TypeErrorf("failed to parse base redis url: url=%s, err=%s", redisurl.Value(), err.Error())
	}
	scheme := strings.ToLower(strings.TrimSpace(baseUrl.Scheme))

	var client BaseClient
	deferConnect := false
	dc := baseUrl.Query().Get("deferconnect")
	if dc == "true" {
		deferConnect = true
	}
	baseUrl.Query().Del("deferconnect")
	switch scheme {
	case "redis", "rediss", "unix":
		if opt, err := redis.ParseURL(baseUrl.String()); err != nil {
			return object.TypeErrorf("failed to parse redis url: url=%s, err=%s", redisurl.Value(), err.Error())
		} else {
			client = redis.NewClient(opt)
		}
	case "rediscluster", "redisscluster":
		baseUrl.Scheme = strings.Replace(baseUrl.Scheme, "cluster", "", 1)
		if opt, err := redis.ParseClusterURL(baseUrl.String()); err != nil {
			return object.TypeErrorf("failed to parse cluster redis url: url=%s, err=%s", redisurl.Value(), err.Error())
		} else {
			client = redis.NewClusterClient(opt)
		}
	default:
		return object.TypeErrorf("unsupported redis scheme: %s", scheme)
	}
	if !deferConnect {
		if err := client.Ping(context.Background()).Err(); err != nil {
			return object.TypeErrorf("failed to connect to redis: url=%s, err=%s", redisurl.Value(), err.Error())
		} else {
			fmt.Printf("Connected to redis: url=%s\n", redisurl.Value())
		}
	}
	return New(ctx, client)
}

// Module returns the `redis` module object
func Module() *object.Module {
	return object.NewBuiltinsModule("redis", map[string]object.Object{
		"connect": object.NewBuiltin("connect", Connect),
	})
}
