package main

import (
	"context"
	"fmt"
	rdis "github.com/testcontainers/testcontainers-go/modules/redis"
	"net"
	"os"
)

var (
	ctx           = context.Background()
	container     *rdis.RedisContainer
	systemTempDir = os.TempDir()
	redisPort     = reservePort()
	configFile    string
	usockName     string
)

func reservePort() int {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(fmt.Sprintf("Failed to reserver listener port: %s", err.Error()))
	}
	port := listener.Addr().(*net.TCPAddr).Port
	listener.Close()
	return port
}

func genConfigFile() {
	if f, err := os.CreateTemp(systemTempDir, "redis.conf"); err != nil {
		panic(fmt.Sprintf("Failed to create temp conf file: %s", err.Error()))
	} else {
		configFile = f.Name()
		usockName = fmt.Sprintf("%s.sock", configFile)
		defer func() {
			_ = f.Close()
		}()
		f.WriteString("bind 127.0.0.1 -::1\n")
		f.WriteString(fmt.Sprintf("port %d\n", redisPort))
		f.WriteString(fmt.Sprintf("unixsocket %s\n", usockName))

		fmt.Printf("Generated redis.conf: file=%s, port=%d, usock=%s\n", configFile, redisPort, usockName)
		f.Close()
	}
}

func init() {
	// func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*RedisContainer, error)
	genConfigFile()
	var err error
	container, err = rdis.Run(ctx, "redis:7", rdis.WithConfigFile(configFile))
	if err != nil {
		panic(fmt.Sprintf("Failed to start redis container: %s", err.Error()))
	}
}

//
//func TestRedisConn_Get(t *testing.T) {
//	db, mock := redismock.NewClientMock()
//	conn := New(context.Background(), db)
//
//	key := object.NewString("key")
//
//	mock.ExpectGet("key").SetVal("value")
//
//	result := conn.Get(context.Background(), key)
//	assert.Equal(t, object.NewString("value"), result)
//	assert.NoError(t, mock.ExpectationsWereMet())
//}
//
//func TestRedisConn_Incr(t *testing.T) {
//	db, mock := redismock.NewClientMock()
//	conn := New(context.Background(), db)
//
//	key := object.NewString("counter")
//
//	mock.ExpectIncr("counter").SetVal(1)
//
//	result := conn.Incr(context.Background(), key)
//	assert.Equal(t, object.NewInt(1), result)
//	assert.NoError(t, mock.ExpectationsWereMet())
//}
//
//func TestRedisConn_HSet(t *testing.T) {
//	db, mock := redismock.NewClientMock()
//	conn := New(context.Background(), db)
//
//	key := object.NewString("hash")
//	field := object.NewString("field")
//	value := object.NewString("value")
//
//	mock.ExpectHSet("hash", "field", "value").SetVal(1)
//
//	result := conn.HSet(context.Background(), key, field, value)
//	assert.Equal(t, object.NewInt(1), result)
//	assert.NoError(t, mock.ExpectationsWereMet())
//}
//
//func TestRedisConn_HGet(t *testing.T) {
//	db, mock := redismock.NewClientMock()
//	conn := New(context.Background(), db)
//
//	key := object.NewString("hash")
//	field := object.NewString("field")
//
//	mock.ExpectHGet("hash", "field").SetVal("value")
//
//	result := conn.HGet(context.Background(), key, field)
//	assert.Equal(t, object.NewString("value"), result)
//	assert.NoError(t, mock.ExpectationsWereMet())
//}
