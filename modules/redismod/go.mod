module github.com/risor-io/risor/modules/redismod

go 1.22.1

toolchain go1.22.2

replace github.com/risor-io/risor => ../..

require (
	github.com/redis/go-redis/v9 v9.7.0
	github.com/go-redis/redismock/v9 v9.2.0
	github.com/risor-io/risor v1.7.0
	github.com/testcontainers/testcontainers-go/modules/redis v0.34.0
)
