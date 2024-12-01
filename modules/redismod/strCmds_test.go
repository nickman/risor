package redismod

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/risor-io/risor"
	"github.com/risor-io/risor/modules/duration"
	"github.com/risor-io/risor/object"
	"github.com/stretchr/testify/require"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"testing"
	"text/template"
	"time"
)

const (
	DEFAULT_RISOR_EXE = "cmd/risor/risor"
	DEFAULT_HOST      = "127.0.0.1"
	DEFAULT_PORT      = 6379

	//DEFAULT_HOST      = "192.168.1.244"
	//DEFAULT_PORT      = 6279
)

var (
	risorExe  = DEFAULT_RISOR_EXE
	redisHost = DEFAULT_HOST
	redisPort = DEFAULT_PORT
)

func init() {
	exe := os.Getenv("RISOR_EXE")
	if exe != "" {
		risorExe = exe
	}
	host := os.Getenv("REDIS_HOST")
	if host != "" {
		redisHost = host
	}
	portStr := os.Getenv("REDIS_PORT")
	if portStr != "" {
		port, _ := strconv.ParseInt(portStr, 10, 64)
		redisPort = int(port)
	}
}

func TestRedisConn_Set(t *testing.T) {
	u := uuid.New().String()
	result, err := Exec(SetValue, map[string]string{
		"key":   "ABC",
		"value": u,
	})
	require.Nil(t, err)
	require.Equal(t, object.NewString("OK"), result)
}

func TestRedisConn_SetAndGet(t *testing.T) {
	u := uuid.New().String()
	result, err := Exec(SetAndGetValue, map[string]string{
		"key":   "ABC",
		"value": u,
	})
	require.Nil(t, err)
	require.Equal(t, object.Nil, result)
}

func TestRedisConn_SetWithExpiry(t *testing.T) {
	u := uuid.New().String()
	result, err := Exec(SetValueExp, map[string]string{
		"key":   "ABC",
		"value": u,
	})
	require.Nil(t, err)
	require.Equal(t, object.Nil, result)
}

func TestRedisConn_GetDuration(t *testing.T) {
	result, err := Exec(GetDur, nil)
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_Increment(t *testing.T) {
	u := uuid.New().String()
	init := rand.Int()
	expected := init + 1
	result, err := Exec(Incr, map[string]string{
		"key":    u,
		"value":  fmt.Sprintf(`%d`, init),
		"result": fmt.Sprintf(`%d`, expected),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_Decrement(t *testing.T) {
	u := uuid.New().String()
	init := rand.Int()
	expected := init - 1
	result, err := Exec(Decr, map[string]string{
		"key":    u,
		"value":  fmt.Sprintf(`%d`, init),
		"result": fmt.Sprintf(`%d`, expected),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_IncrementBy(t *testing.T) {
	u := uuid.New().String()
	init := rand.Intn(9999)
	incr := rand.Intn(9999)
	expected := init + incr
	result, err := Exec(IncrBy, map[string]string{
		"key":    u,
		"value":  fmt.Sprintf(`%d`, init),
		"result": fmt.Sprintf(`%d`, expected),
		"incr":   fmt.Sprintf(`%d`, incr),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_DecrementBy(t *testing.T) {
	u := uuid.New().String()
	init := rand.Intn(9999)
	decr := rand.Intn(9999)
	expected := init - decr
	result, err := Exec(DecrBy, map[string]string{
		"key":    u,
		"value":  fmt.Sprintf(`%d`, init),
		"result": fmt.Sprintf(`%d`, expected),
		"decr":   fmt.Sprintf(`%d`, decr),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_LCS(t *testing.T) {
	expected := "mytext"
	result, err := Exec(Lcs, map[string]string{
		"key1":     uuid.New().String(),
		"value1":   "ohmytext",
		"key2":     uuid.New().String(),
		"value2":   "mynewtext",
		"expected": expected,
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_LCSLen(t *testing.T) {
	expected := "6"
	result, err := Exec(LcsLen, map[string]string{
		"key1":     uuid.New().String(),
		"value1":   "ohmytext",
		"key2":     uuid.New().String(),
		"value2":   "mynewtext",
		"expected": expected,
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_LCSIdx(t *testing.T) {
	expected := "6"
	result, err := Exec(LcsIdx, map[string]string{
		"key1":      uuid.New().String(),
		"value1":    "ohmytext",
		"key2":      uuid.New().String(),
		"value2":    "mynewtext",
		"expected":  expected,
		"expected2": `[{"key1": {"end": 7, "start": 4}, "key2": {"end": 8, "start": 5}}, {"key1": {"end": 3, "start": 2}, "key2": {"end": 1, "start": 0}}]`,
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_GetSet(t *testing.T) {
	result, err := Exec(GetSet, map[string]string{
		"key1":   uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_GetRange(t *testing.T) {
	result, err := Exec(GetRange, map[string]string{
		"key1": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_GetExWithExpiration(t *testing.T) {
	result, err := Exec(GetExWithExp, map[string]string{
		"key":   uuid.New().String(),
		"value": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_GetExNoExpiration(t *testing.T) {
	result, err := Exec(GetExNoExp, map[string]string{
		"key":   uuid.New().String(),
		"value": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_GetDel(t *testing.T) {
	result, err := Exec(GetDel, map[string]string{
		"key":   uuid.New().String(),
		"value": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_IncrementByFloat(t *testing.T) {
	//for idx := 0; idx < 10; idx++ {
	u := uuid.New().String()
	init := 1.392
	incr := 3.2
	//init := rand.Float64()
	//incr := rand.Float64()

	expected := init + incr
	result, err := Exec(IncrByFloat, map[string]string{
		"key":    u,
		"init":   fmt.Sprintf(`%f`, init),
		"result": fmt.Sprintf(`%f`, expected),
		"incr":   fmt.Sprintf(`%f`, incr),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
	//}
}

func TestRedisConn_MSet(t *testing.T) {
	result, err := Exec(MSet, map[string]string{
		"key1":   uuid.New().String(),
		"key2":   uuid.New().String(),
		"key3":   uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
		"value3": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_MGet(t *testing.T) {
	result, err := Exec(MGet, map[string]string{
		"key1":   uuid.New().String(),
		"key2":   uuid.New().String(),
		"key3":   uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
		"value3": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_MSetNX(t *testing.T) {
	result, err := Exec(MSetNX, map[string]string{
		"key1":   uuid.New().String(),
		"key2":   uuid.New().String(),
		"key3":   uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
		"value3": uuid.New().String(),
		"value4": uuid.New().String(),
		"value5": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_SetExWithExp(t *testing.T) {
	result, err := Exec(SetExWithExp, map[string]string{
		"key1":   uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_SetNXWithExp(t *testing.T) {
	result, err := Exec(SetNXWithExp, map[string]string{
		"key1":   uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_SetXX(t *testing.T) {
	result, err := Exec(SetXX, map[string]string{
		"key1":   uuid.New().String(),
		"value1": uuid.New().String(),
		"value2": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_SetRange(t *testing.T) {
	result, err := Exec(SetRange, map[string]string{
		"key1": uuid.New().String(),
		"key2": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_StrLen(t *testing.T) {
	u := uuid.New().String()
	result, err := Exec(StrLen, map[string]string{
		"key1":     uuid.New().String(),
		"value1":   u,
		"expected": fmt.Sprintf(`%d`, len(u)),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_MGetFx(t *testing.T) {
	result, err := Exec(MGetFx, map[string]string{
		"key1":    uuid.New().String(),
		"key2":    uuid.New().String(),
		"key3":    uuid.New().String(),
		"key4":    uuid.New().String(),
		"key5":    uuid.New().String(),
		"key6":    uuid.New().String(),
		"key7":    uuid.New().String(),
		"key8":    uuid.New().String(),
		"key9":    uuid.New().String(),
		"key10":   uuid.New().String(),
		"value1":  uuid.New().String(),
		"value2":  uuid.New().String(),
		"value3":  uuid.New().String(),
		"value4":  uuid.New().String(),
		"value5":  uuid.New().String(),
		"value6":  uuid.New().String(),
		"value7":  uuid.New().String(),
		"value8":  uuid.New().String(),
		"value9":  uuid.New().String(),
		"value10": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func TestRedisConn_MGetFxStop(t *testing.T) {
	result, err := Exec(MGetFxStopAt5, map[string]string{
		"key1":    uuid.New().String(),
		"key2":    uuid.New().String(),
		"key3":    uuid.New().String(),
		"key4":    uuid.New().String(),
		"key5":    uuid.New().String(),
		"key6":    uuid.New().String(),
		"key7":    uuid.New().String(),
		"key8":    uuid.New().String(),
		"key9":    uuid.New().String(),
		"key10":   uuid.New().String(),
		"value1":  uuid.New().String(),
		"value2":  uuid.New().String(),
		"value3":  uuid.New().String(),
		"value4":  uuid.New().String(),
		"value5":  uuid.New().String(),
		"value6":  uuid.New().String(),
		"value7":  uuid.New().String(),
		"value8":  uuid.New().String(),
		"value9":  uuid.New().String(),
		"value10": uuid.New().String(),
	})
	require.Nil(t, err)
	require.NotNil(t, result)
}

func currentTest() string {
	// Get the current function name using runtime.FuncForPC
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return funcName
}

func compileTemplate(name, source string, args map[string]string) string {
	if args == nil {
		return source
	}
	if templ, err := template.New(name).Parse(source); err != nil {
		panic(err)
	} else {
		var b bytes.Buffer
		if err := templ.Execute(&b, args); err != nil {
			panic(err)
		}
		return b.String()
	}
}

func Exec(script string, args map[string]string) (object.Object, error) {
	testName := currentTest()
	startTime := time.Now()
	source := compileTemplate(testName, script, args)
	fmt.Printf("SOURCE:\n===================================\n%s\n===================================\n", source)
	ctx := context.Background()
	defer func() {
		//runtime.GC()
		fmt.Printf("Execution: test=%s, elapsed=%s\n", testName, time.Since(startTime))
	}()
	return risor.Eval(ctx, source, risor.WithConcurrency(), risor.WithGlobals(map[string]any{
		"redis":    Module(),
		"duration": duration.Module(),
	}))
}

const (
	RConnect = `rdis := redis.connect("redis://127.0.0.1:6379?db=1")`

	SetValue = RConnect + `
	rdis.set("{{.key}}", "{{.value}}")
`
	GetDur = `
		d := duration.parse("1s")
		assert(d.tostring() == "duration(1s)")
`

	SetValueExp = RConnect + `
	rdis.set("{{.key}}", "{{.value}}", duration.parse("1s"))
	assert(rdis.get("{{.key}}") == "{{.value}}")
	time.sleep(2)
	result := rdis.get("{{.key}}")
	assert(rdis.isnull(result) == true)`

	SetAndGetValue = RConnect + `
	rdis.set("{{.key}}", "{{.value}}")
    assert(rdis.get("{{.key}}") == "{{.value}}")
`
	Incr = RConnect + `
	rdis.set("{{.key}}", "{{.value}}")
    assert(rdis.incr("{{.key}}") == {{.result}})
`
	Decr = RConnect + `
	rdis.set("{{.key}}", "{{.value}}")
    assert(rdis.decr("{{.key}}") == {{.result}})
`

	IncrBy = RConnect + `
	rdis.set("{{.key}}", "{{.value}}")
    assert(rdis.incrby("{{.key}}", {{.incr}}) == {{.result}})`

	IncrByFloat = RConnect + `
	rdis.set("{{.key}}", "{{.init}}")
	actual := rdis.incrbyfloat("{{.key}}", {{.incr}})
	printf("Actual: %f, Expected: %f\n", actual, {{.result}})
    assert(actual == {{.result}})`

	DecrBy = RConnect + `
	rdis.set("{{.key}}", "{{.value}}")
    assert(rdis.decrby("{{.key}}", {{.decr}}) == {{.result}})
`
	Lcs = RConnect + `
	rdis.set("{{.key1}}", "{{.value1}}")
	rdis.set("{{.key2}}", "{{.value2}}")
	lcsResult := rdis.lcs("{{.key1}}", "{{.key2}}")
    assert(lcsResult.match_string == "{{.expected}}")
`
	LcsLen = RConnect + `
	rdis.set("{{.key1}}", "{{.value1}}")
	rdis.set("{{.key2}}", "{{.value2}}")
	lcsResult := rdis.lcs("{{.key1}}", "{{.key2}}", map({"len" : true}))
    assert(lcsResult.len == {{.expected}})`

	LcsIdx = RConnect + `
	rdis.set("{{.key1}}", "{{.value1}}")
	rdis.set("{{.key2}}", "{{.value2}}")
	lcsResult := rdis.lcs("{{.key1}}", "{{.key2}}", map({"idx" : true}))
	printf("%s", lcsResult)
    assert(lcsResult.len == {{.expected}})
	assert(lcsResult.matches == {{.expected2}})`

	GetSet = RConnect + `
	rdis.set("{{.key1}}", "{{.value1}}")
	v1 := rdis.getset("{{.key1}}", "{{.value2}}")
	assert(v1 == "{{.value1}}")
	v2 := rdis.get("{{.key1}}")
	assert(v2 == "{{.value2}}")`

	GetRange = RConnect + `
	rdis.set("{{.key1}}", "This is a string")
	v := rdis.getrange("{{.key1}}", 0, 3)
	assert(v == "This")
	v = rdis.getrange("{{.key1}}", -3, -1)
	assert(v == "ing")
	v = rdis.getrange("{{.key1}}", 0, -1)
	assert(v == "This is a string")
	v = rdis.getrange("{{.key1}}", 10, 100)
	assert(v == "string")`

	GetExWithExp = RConnect + `
	rdis.set("{{.key}}", "{{.value}}")
	assert(rdis.getex("{{.key}}", duration.parse("1s")) == "{{.value}}")
	time.sleep(2)
	result := rdis.get("{{.key}}")
	assert(rdis.isnull(result) == true)`

	GetExNoExp = RConnect + `
	rdis.set("{{.key}}", "{{.value}}")
	assert(rdis.getex("{{.key}}") == "{{.value}}")
	time.sleep(2)
	result := rdis.get("{{.key}}")
	assert(result == "{{.value}}")`

	GetDel = RConnect + `
	rdis.set("{{.key}}", "{{.value}}")
	assert(rdis.get("{{.key}}") == "{{.value}}")
	assert(rdis.getdel("{{.key}}") == "{{.value}}")
	result := rdis.get("{{.key}}")
	assert(rdis.isnull(result) == true)`

	MSet = RConnect + `
	rdis.mset("{{.key1}}", "{{.value1}}", "{{.key2}}", "{{.value2}}", "{{.key3}}", "{{.value3}}")
	assert(rdis.get("{{.key1}}") == "{{.value1}}")
	assert(rdis.get("{{.key2}}") == "{{.value2}}")
	assert(rdis.get("{{.key3}}") == "{{.value3}}")`

	MGet = RConnect + `
	rdis.mset("{{.key1}}", "{{.value1}}", "{{.key2}}", "{{.value2}}", "{{.key3}}", "{{.value3}}")
	result := rdis.mget("{{.key1}}", "{{.key2}}", "{{.key3}}")
	assert(result[0] == "{{.value1}}")
	assert(result[1] == "{{.value2}}")	
	assert(result[2] == "{{.value3}}")`

	MSetNX = RConnect + `
	assert(rdis.msetnx("{{.key1}}", "{{.value1}}", "{{.key2}}", "{{.value2}}", "{{.key3}}", "{{.value3}}") == true)
	assert(rdis.msetnx("{{.key2}}", "{{.value4}}", "{{.key3}}", "{{.value5}}") == false)
	result := rdis.mget("{{.key1}}", "{{.key2}}", "{{.key3}}")
	assert(result[0] == "{{.value1}}")
	assert(result[1] == "{{.value2}}")	
	assert(result[2] == "{{.value3}}")`

	SetExWithExp = RConnect + `
	set1 := rdis.setex("{{.key1}}", "{{.value1}}", duration.parse("1s"))
	assert(set1 == "OK")
	time.sleep(2)
	result := rdis.get("{{.key1}}")
	assert(rdis.isnull(result) == true)`

	SetNXWithExp = RConnect + `
	set1 := rdis.setnx("{{.key1}}", "{{.value1}}", duration.parse("1s"))
	assert(set1 == true)
	assert(rdis.get("{{.key1}}") == "{{.value1}}")
	set2 := rdis.setnx("{{.key1}}", "{{.value2}}", duration.parse("1s"))
	assert(set2 == false)
	assert(rdis.get("{{.key1}}") == "{{.value1}}")
	time.sleep(2)
	result := rdis.get("{{.key1}}")
	assert(rdis.isnull(result) == true)`

	SetXX = RConnect + `
	set1 := rdis.set("{{.key1}}", "{{.value1}}", duration.parse("2s"))
	assert(set1 == "OK", "Set1 failed")
	assert(rdis.get("{{.key1}}") == "{{.value1}}")
	set2 := rdis.setxx("{{.key1}}", "{{.value2}}")
	assert(set2 == true, "Set2 failed: expected=true")
	val2 := rdis.get("{{.key1}}")
	printf("Val2: %s\n", val2)
	assert(val2 == "{{.value2}}", "key1 == value2 failed")
	time.sleep(2)
	result := rdis.get("{{.key1}}")
	assert(rdis.isnull(result) == true, sprintf("IsNull failed: expected=true, value=%s", result))`

	SetRange = RConnect + `
	rdis.set("{{.key1}}", "Hello World")
	assert(rdis.get("{{.key1}}") == "Hello World")
	set1 := rdis.setrange("{{.key1}}", 6, "Redis")
	assert(set1 == 11)
	assert(rdis.get("{{.key1}}") == "Hello Redis")
	set2 := rdis.setrange("{{.key2}}", 6, "Redis")
	assert(set2 == 11)
	assert(rdis.get("{{.key2}}") == "\x00\x00\x00\x00\x00\x00Redis")`

	StrLen = RConnect + `
	rdis.set("{{.key1}}", "{{.value1}}")
	assert(rdis.get("{{.key1}}") == "{{.value1}}")
	len1 := rdis.strlen("{{.key1}}")
	assert(len1 == {{.expected}})`

	MGetFx = RConnect + `
	rdis.mset("{{.key1}}", "{{.value1}}", "{{.key2}}", "{{.value2}}", "{{.key3}}", "{{.value3}}", "{{.key4}}", "{{.value4}}", "{{.key5}}", "{{.value5}}", "{{.key6}}", "{{.value6}}", "{{.key7}}", "{{.value7}}", "{{.key8}}", "{{.value8}}", "{{.key9}}", "{{.value9}}", "{{.key10}}", "{{.value10}}")	
	assert(rdis.get("{{.key1}}") == "{{.value1}}")
	fx := func(str) {
		printf("RESULT: %s\n", str)
		return true
	}
	rcount := rdis.mgetfunc(fx, "{{.key1}}", "{{.key2}}", "{{.key3}}", "{{.key4}}", "{{.key5}}", "{{.key6}}", "{{.key7}}", "{{.key8}}", "{{.key9}}", "{{.key10}}")
	printf("RCOUNT: %d\n", rcount)
	assert(rcount == 10)`

	MGetFxStopAt5 = RConnect + `
	rdis.mset("{{.key1}}", "{{.value1}}", "{{.key2}}", "{{.value2}}", "{{.key3}}", "{{.value3}}", "{{.key4}}", "{{.value4}}", "{{.key5}}", "{{.value5}}", "{{.key6}}", "{{.value6}}", "{{.key7}}", "{{.value7}}", "{{.key8}}", "{{.value8}}", "{{.key9}}", "{{.value9}}", "{{.key10}}", "{{.value10}}")	
	assert(rdis.get("{{.key1}}") == "{{.value1}}")
	cnt := 0
	fx := func(str) {
		printf("RESULT: %s\n", str)
		cnt++
		if cnt == 5 {
			return false
		} else {
			return true
		}
	}
	rcount := rdis.mgetfunc(fx, "{{.key1}}", "{{.key2}}", "{{.key3}}", "{{.key4}}", "{{.key5}}", "{{.key6}}", "{{.key7}}", "{{.key8}}", "{{.key9}}", "{{.key10}}")
	printf("RCOUNT: %d\n", rcount)
	assert(rcount == 5)`
)
