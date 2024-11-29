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
)

const (
	DEFAULT_RISOR_EXE = "cmd/risor/risor"
	DEFAULT_HOST      = "192.168.1.244"
	DEFAULT_PORT      = 6279
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
	// {"len": 6, "match_string": "", "matches": [{"key1": {"end": 7, "start": 4}, "key2": {"end": 8, "start": 5}}, {"key1": {"end": 3, "start": 2}, "key2": {"end": 1, "start": 0}}]}--- PASS: TestRedisConn_LCSIdx (0.11s)
	require.Nil(t, err)
	require.NotNil(t, result)
}

func currentTest() string {
	// Get the current function name using runtime.FuncForPC
	pc, _, _, _ := runtime.Caller(1)
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
	source := compileTemplate(testName, script, args)
	fmt.Printf("SOURCE:\n===================================\n%s\n===================================\n", source)
	ctx := context.Background()
	defer func() {
		runtime.GC()
	}()
	return risor.Eval(ctx, source, risor.WithGlobals(map[string]any{
		"redis":    Module(),
		"duration": duration.Module(),
	}))
}

const (
	RConnect = `rdis := redis.connect("redis://192.168.1.244:6279?db=1")`

	SetValue = RConnect + `
	//defer rdis.close()
	rdis.set("{{.key}}", "{{.value}}")
`
	GetDur = `
		d := duration.parse("2s")
		assert(d.tostring() == "duration(2s)")
`

	SetValueExp = RConnect + `
	//defer rdis.close()
	rdis.set("{{.key}}", "{{.value}}", duration.parse("2s"))
	assert(rdis.get("{{.key}}") == "{{.value}}")
	time.sleep(2)
	result := rdis.get("{{.key}}")
	assert(rdis.isnull(result) == true)
`

	SetAndGetValue = RConnect + `
	//defer rdis.close()
	rdis.set("{{.key}}", "{{.value}}")
    assert(rdis.get("{{.key}}") == "{{.value}}")
`
	Incr = RConnect + `
	//defer rdis.close()
	rdis.set("{{.key}}", "{{.value}}")
    assert(rdis.incr("{{.key}}") == {{.result}})
`
	Decr = RConnect + `
	//defer rdis.close()
	rdis.set("{{.key}}", "{{.value}}")
    assert(rdis.decr("{{.key}}") == {{.result}})
`

	IncrBy = RConnect + `
	//defer rdis.close()
	rdis.set("{{.key}}", "{{.value}}")
    assert(rdis.incrby("{{.key}}", {{.incr}}) == {{.result}})
`

	DecrBy = RConnect + `
	//defer rdis.close()
	rdis.set("{{.key}}", "{{.value}}")
    assert(rdis.decrby("{{.key}}", {{.decr}}) == {{.result}})
`
	Lcs = RConnect + `
	//defer rdis.close()
	rdis.set("{{.key1}}", "{{.value1}}")
	rdis.set("{{.key2}}", "{{.value2}}")
	lcsResult := rdis.lcs("{{.key1}}", "{{.key2}}")
    assert(lcsResult.match_string == "{{.expected}}")
`
	LcsLen = RConnect + `
	//defer rdis.close()
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
	rdis.set("{{.key2}}", "{{.value2}}")
	lcsResult := rdis.lcs("{{.key1}}", "{{.key2}}", map({"idx" : true}))
	printf("%s", lcsResult)
    assert(lcsResult.len == {{.expected}})
	assert(lcsResult.matches == {{.expected2}})`
)
