package redismod

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"math/rand"
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

func TestRedisConn_HGet(t *testing.T) {
	result, err := Exec(HGet, map[string]string{
		"key1":   uuid.New().String(),
		"field1": uuid.New().String(),
		"value1": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HLen(t *testing.T) {
	result, err := Exec(HLen, map[string]string{
		"key1":   uuid.New().String(),
		"field1": uuid.New().String(),
		"field2": uuid.New().String(),
		"field3": uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
		"value3": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HDel(t *testing.T) {
	result, err := Exec(HDel, map[string]string{
		"key1":   uuid.New().String(),
		"field1": uuid.New().String(),
		"field2": uuid.New().String(),
		"field3": uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
		"value3": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HVals(t *testing.T) {
	result, err := Exec(HVals, map[string]string{
		"key1":   uuid.New().String(),
		"field1": uuid.New().String(),
		"field2": uuid.New().String(),
		"field3": uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
		"value3": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HExists(t *testing.T) {
	result, err := Exec(HExists, map[string]string{
		"key1":   uuid.New().String(),
		"field1": uuid.New().String(),
		"field2": uuid.New().String(),
		"field3": uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
		"value3": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HGetAll(t *testing.T) {
	result, err := Exec(HGetAll, map[string]string{
		"key1":   uuid.New().String(),
		"field1": uuid.New().String(),
		"field2": uuid.New().String(),
		"field3": uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
		"value3": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HIncrBy(t *testing.T) {
	key1 := uuid.New().String()
	field1 := uuid.New().String()
	init := rand.Intn(9999)
	incr := rand.Intn(9999)
	sum := init + incr
	result, err := Exec(HIncrBy, map[string]string{
		"key1":   key1,
		"field1": field1,
		"value1": fmt.Sprintf("%d", init),
		"incr":   fmt.Sprintf("%d", incr),
		"result": fmt.Sprintf("%d", sum),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HIncrByFloat(t *testing.T) {
	key1 := uuid.New().String()
	field1 := uuid.New().String()
	init := 3.14
	incr := 6.009
	sum := init + incr
	result, err := Exec(HIncrByFloat, map[string]string{
		"key1":   key1,
		"field1": field1,
		"value1": fmt.Sprintf("%f", init),
		"incr":   fmt.Sprintf("%f", incr),
		"result": fmt.Sprintf("%f", sum),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HMSet(t *testing.T) {
	result, err := Exec(HMSet, map[string]string{
		"key1":   uuid.New().String(),
		"field1": uuid.New().String(),
		"field2": uuid.New().String(),
		"field3": uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
		"value3": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HMGet(t *testing.T) {
	result, err := Exec(HMGet, map[string]string{
		"key1":   uuid.New().String(),
		"field1": uuid.New().String(),
		"field2": uuid.New().String(),
		"field3": uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
		"value3": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HKeys(t *testing.T) {
	result, err := Exec(HKeys, map[string]string{
		"key1":   uuid.New().String(),
		"field1": uuid.New().String(),
		"field2": uuid.New().String(),
		"field3": uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
		"value3": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HSetNX(t *testing.T) {
	result, err := Exec(HSetNX, map[string]string{
		"key1":   uuid.New().String(),
		"field1": uuid.New().String(),
		"field2": uuid.New().String(),
		"field3": uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
		"value3": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HScan(t *testing.T) {
	result, err := Exec(HScan, map[string]string{
		"key1": uuid.New().String(),
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HScanFunc(t *testing.T) {
	result, err := Exec(HScanFunc, map[string]string{
		"key1": uuid.New().String(),
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_HScanFuncStop(t *testing.T) {
	result, err := Exec(HScanFuncStop, map[string]string{
		"key1": uuid.New().String(),
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	require.Nil(t, err)
	require.NotNil(t, result)
}

//func TestRedisConn_HScanFuncBatch(t *testing.T) {
//	result, err := Exec(HScanFuncBatch, map[string]string{
//		"key1": uuid.New().String(),
//	})
//	if err != nil {
//		fmt.Printf("Error: %s\n", err.Error())
//	}
//	require.Nil(t, err)
//	require.NotNil(t, result)
//}

const (
	HSet = RConnect + `
	rdis.hset("{{.key1}}", "{{.field1}}", "{{.value1}}")
	val := rdis.hget("{{.key1}}" , "{{.field1}}")	
	assert(val == "{{.value1}}")`

	HGet = RConnect + `
	rdis.hset("{{.key1}}", "{{.field1}}", "{{.value1}}")
	val := rdis.hget("{{.key1}}" , "{{.field1}}")	
	assert(val == "{{.value1}}")`

	HLen = RConnect + `
	rdis.hset("{{.key1}}", "{{.field1}}", "{{.value1}}")
	rdis.hset("{{.key1}}", "{{.field2}}", "{{.value2}}")
	rdis.hset("{{.key1}}", "{{.field3}}", "{{.value3}}")
	hashlen := rdis.hlen("{{.key1}}")
	printf("hashlen: %d", hashlen)
	assert(hashlen == 3)`

	HDel = RConnect + `
	rdis.hset("{{.key1}}", "{{.field1}}", "{{.value1}}")
	rdis.hset("{{.key1}}", "{{.field2}}", "{{.value2}}")
	rdis.hset("{{.key1}}", "{{.field3}}", "{{.value3}}")
	rdis.hdel("{{.key1}}", "{{.field1}}")
	assert(rdis.hlen("{{.key1}}") == 2)
	rdis.hdel("{{.key1}}", "{{.field2}}")
	assert(rdis.hlen("{{.key1}}") == 1)
	rdis.hdel("{{.key1}}", "{{.field2}}")
	assert(rdis.hlen("{{.key1}}") == 1)
	rdis.hdel("{{.key1}}", "{{.field3}}")
	assert(rdis.hlen("{{.key1}}") == 0)`

	HVals = RConnect + `
	rdis.hset("{{.key1}}", "{{.field1}}", "{{.value1}}")
	rdis.hset("{{.key1}}", "{{.field2}}", "{{.value2}}")
	rdis.hset("{{.key1}}", "{{.field3}}", "{{.value3}}")
	vals := rdis.hvals("{{.key1}}")
	printf("vals: %v\n", vals)
	expected := ["{{.value1}}", "{{.value2}}", "{{.value3}}"]
	assert(vals == expected)`

	HExists = RConnect + `
	assert(rdis.hexists("{{.key1}}", "{{.field1}}") == false)
	rdis.hset("{{.key1}}", "{{.field2}}", "{{.value2}}")
	assert(rdis.hexists("{{.key1}}", "{{.field1}}") == false)
	rdis.hset("{{.key1}}", "{{.field1}}", "{{.value1}}")
	assert(rdis.hexists("{{.key1}}", "{{.field1}}") == true)`

	HGetAll = RConnect + `
	rdis.hset("{{.key1}}", "{{.field1}}", "{{.value1}}")
	rdis.hset("{{.key1}}", "{{.field2}}", "{{.value2}}")
	rdis.hset("{{.key1}}", "{{.field3}}", "{{.value3}}")
	vals := rdis.hgetall("{{.key1}}")
	for k,v := range vals {
		printf("key: %s, value: %s\n", k, v)
	}
	assert(vals["{{.field1}}"] == "{{.value1}}")
	assert(vals["{{.field2}}"] == "{{.value2}}")
	assert(vals["{{.field3}}"] == "{{.value3}}")`

	HIncrBy = RConnect + `
	rdis.hset("{{.key1}}", "{{.field1}}", "{{.value1}}")
	v1 := rdis.hget("{{.key1}}", "{{.field1}}")
	assert(v1 == "{{.value1}}")
	result := rdis.hincrby("{{.key1}}", "{{.field1}}", {{.incr}})
	assert(result == {{.result}})`

	HIncrByFloat = RConnect + `
	rdis.hset("{{.key1}}", "{{.field1}}", "{{.value1}}")
	v1 := rdis.hget("{{.key1}}", "{{.field1}}")
	assert(v1 == "{{.value1}}")
	result := rdis.hincrbyfloat("{{.key1}}", "{{.field1}}", {{.incr}})
	assert(result == {{.result}})`

	HMSet = RConnect + `
	rdis.hmset("{{.key1}}", "{{.field1}}", "{{.value1}}", "{{.field2}}", "{{.value2}}", "{{.field3}}", "{{.value3}}")
	assert(rdis.hlen("{{.key1}}") == 3)
    assert(rdis.hget("{{.key1}}", "{{.field1}}") == "{{.value1}}")
	assert(rdis.hget("{{.key1}}", "{{.field2}}") == "{{.value2}}")
	assert(rdis.hget("{{.key1}}", "{{.field3}}") == "{{.value3}}")`

	HMGet = RConnect + `
	rdis.hmset("{{.key1}}", "{{.field1}}", "{{.value1}}", "{{.field2}}", "{{.value2}}", "{{.field3}}", "{{.value3}}")
	assert(rdis.hlen("{{.key1}}") == 3)
	vals := rdis.hmget("{{.key1}}", "{{.field1}}", "{{.field2}}", "{{.field3}}")
	assert(vals[0] == "{{.value1}}")
	assert(vals[1] == "{{.value2}}")
	assert(vals[2] == "{{.value3}}")`

	HKeys = RConnect + `
	rdis.hmset("{{.key1}}", "{{.field1}}", "{{.value1}}", "{{.field2}}", "{{.value2}}", "{{.field3}}", "{{.value3}}")
	assert(rdis.hlen("{{.key1}}") == 3)
	vals := rdis.hkeys("{{.key1}}")
	assert(vals[0] == "{{.field1}}")
	assert(vals[1] == "{{.field2}}")
	assert(vals[2] == "{{.field3}}")`

	HSetNX = RConnect + `
	assert(rdis.hexists("{{.key1}}", "{{.field1}}") == false)
	assert(rdis.hsetnx("{{.key1}}", "{{.field1}}", "{{.value1}}") == true)
	assert(rdis.hexists("{{.key1}}", "{{.field1}}") == true)
	assert(rdis.hget("{{.key1}}", "{{.field1}}") == "{{.value1}}")
	assert(rdis.hsetnx("{{.key1}}", "{{.field1}}", "{{.value2}}") == false)`

	HScan = RConnect + `
	msize := 10000
	for idx := 0; idx < msize; idx++ {
		rdis.hset("{{.key1}}", sprintf("field:%d", idx), sprintf("value:%d", idx))
	}
	fieldCount := rdis.hlen("{{.key1}}")
	assert(fieldCount == msize)
	printf("fieldCount: %d\n", fieldCount)
	cursor := 0
	keyPairCount := 0
	pageCount := 0
	for {
		scanResult := rdis.hscan("{{.key1}}", cursor, "", 10)
		// Could also be:
		// scanResult := rdis.hscan("{{.key1}}", cursor)
		scanKeyPairs := scanResult[1]	
		for k,v := range scanKeyPairs {
			seq := k.split(":")[1]
			assert(v == sprintf("value:%s", seq))
			keyPairCount++
		}	
		cursor = scanResult[0]
		pageCount++
		if cursor == 0 {
			break
		}
	}
	assert(keyPairCount == msize)
	printf("pageCount: %d\n", pageCount)
	rdis.del("{{.key1}}")`

	HScanFunc = RConnect + `
	msize := 10000
	for idx := 0; idx < msize; idx++ {
		rdis.hset("{{.key1}}", sprintf("field:%d", idx), sprintf("value:%d", idx))
	}
	fieldCount := rdis.hlen("{{.key1}}")
	assert(fieldCount == msize)
	printf("fieldCount: %d\n", fieldCount)

	keyPairCount := 0

	fx := func(k, v) {	
		seq := k.split(":")[1]
		expectedValue := sprintf("value:%s", seq)
		assert(v == expectedValue, sprintf("Unexpected value: %s != %s", v, expectedValue))
		keyPairCount++
		return true
	}
	scanFuncResult := rdis.hscanfunc(fx, "{{.key1}}", 0, "", 10)
	sfPageCount := scanFuncResult[0]
	sfKeyPairCount := scanFuncResult[1]
	assert(keyPairCount == msize, sprintf("Unexpected keyPairCount: %d != %d", keyPairCount, msize))
	assert(sfKeyPairCount == msize, sprintf("Unexpected sfKeyPairCount: %d != %d", sfKeyPairCount, msize))
	rdis.del("{{.key1}}")`

	HScanFuncStop = RConnect + `
	msize := 10000
	stopAt := 3577
	for idx := 0; idx < msize; idx++ {
		rdis.hset("{{.key1}}", sprintf("field:%d", idx), sprintf("value:%d", idx))
	}
	fieldCount := rdis.hlen("{{.key1}}")
	assert(fieldCount == msize)
	printf("fieldCount: %d\n", fieldCount)

	keyPairCount := 0
	

	fx := func(k, v) {	
		seq := k.split(":")[1]
		expectedValue := sprintf("value:%s", seq)
		assert(v == expectedValue, sprintf("Unexpected value: %s != %s", v, expectedValue))
		keyPairCount++	
		if keyPairCount == stopAt {
			printf("Stopping at keyPairCount: %d\n", keyPairCount)
			return false
		}
		return true
	}
	scanFuncResult := rdis.hscanfunc(fx, "{{.key1}}", 0, "", 10)
	sfPageCount := scanFuncResult[0]
	sfKeyPairCount := scanFuncResult[1]
	assert(keyPairCount == stopAt, sprintf("Unexpected keyPairCount: %d != %d", keyPairCount, stopAt))
	assert(sfKeyPairCount == stopAt, sprintf("Unexpected sfKeyPairCount: %d != %d", sfKeyPairCount, stopAt))
	rdis.del("{{.key1}}")`

	//HScanFuncBatch = RConnect + `
	//msize := 10000
	//stopAt := 3577
	//for idx := 0; idx < msize; idx++ {
	//	rdis.hset("{{.key1}}", sprintf("field:%d", idx), sprintf("value:%d", idx))
	//}
	//fieldCount := rdis.hlen("{{.key1}}")
	//assert(fieldCount == msize)
	//printf("fieldCount: %d\n", fieldCount)
	//
	//keyPairCount := 0
	//
	//
	//fx := func(batch) {
	//	for k, v := range batch {
	//		seq := k.split(":")[1]
	//		expectedValue := sprintf("value:%s", seq)
	//		assert(v == expectedValue, sprintf("Unexpected value: %s != %s", v, expectedValue))
	//		keyPairCount++
	//	}
	//	return true
	//}
	//scanFuncResult := rdis.hscanfuncbatch(fx, "{{.key1}}", 0, "", 10)
	//sfPageCount := scanFuncResult[0]
	//sfKeyPairCount := scanFuncResult[1]
	//assert(keyPairCount == stopAt, sprintf("Unexpected keyPairCount: %d != %d", keyPairCount, stopAt))
	//assert(sfKeyPairCount == stopAt, sprintf("Unexpected sfKeyPairCount: %d != %d", sfKeyPairCount, stopAt))
	//rdis.del("{{.key1}}")`
)
