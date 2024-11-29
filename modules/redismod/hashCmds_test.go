package redismod

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRedisConn_HSet(t *testing.T) {
	result, err := Exec(HSet, map[string]string{
		"key1":   uuid.New().String(),
		"field1": uuid.New().String(),
		"value1": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

const (
	HSet = RConnect + `
	rdis.hset("{{.key1}}", "{{.field1}}", "{{.value1}}")
	val := rdis.hget("{{.key1}}" , "{{.field1}}")
	printf("val: %s\n", val)
	assert.Equal("{{.value1}}", val)`
)
