package all

import (
	"github.com/risor-io/risor/builtins"
	modBase64 "github.com/risor-io/risor/modules/base64"
	modBytes "github.com/risor-io/risor/modules/bytes"
	modColor "github.com/risor-io/risor/modules/color"
	modErrors "github.com/risor-io/risor/modules/errors"
	modExec "github.com/risor-io/risor/modules/exec"
	modFilepath "github.com/risor-io/risor/modules/filepath"
	modFmt "github.com/risor-io/risor/modules/fmt"
	modGha "github.com/risor-io/risor/modules/gha"
	modHTTP "github.com/risor-io/risor/modules/http"
	modIsTTY "github.com/risor-io/risor/modules/isatty"
	modJSON "github.com/risor-io/risor/modules/json"
	modMath "github.com/risor-io/risor/modules/math"
	modNet "github.com/risor-io/risor/modules/net"
	modOs "github.com/risor-io/risor/modules/os"
	modRand "github.com/risor-io/risor/modules/rand"
	modRegexp "github.com/risor-io/risor/modules/regexp"
	modStrconv "github.com/risor-io/risor/modules/strconv"
	modStrings "github.com/risor-io/risor/modules/strings"
	modTablewriter "github.com/risor-io/risor/modules/tablewriter"
	modTime "github.com/risor-io/risor/modules/time"
	modYAML "github.com/risor-io/risor/modules/yaml"
	"github.com/risor-io/risor/object"
)

func Builtins() map[string]object.Object {
	result := map[string]object.Object{
		"base64":      modBase64.Module(),
		"bytes":       modBytes.Module(),
		"color":       modColor.Module(),
		"errors":      modErrors.Module(),
		"exec":        modExec.Module(),
		"filepath":    modFilepath.Module(),
		"fmt":         modFmt.Module(),
		"gha":         modGha.Module(),
		"http":        modHTTP.Module(),
		"isatty":      modIsTTY.Module(),
		"json":        modJSON.Module(),
		"math":        modMath.Module(),
		"net":         modNet.Module(),
		"os":          modOs.Module(),
		"rand":        modRand.Module(),
		"regexp":      modRegexp.Module(),
		"strconv":     modStrconv.Module(),
		"strings":     modStrings.Module(),
		"tablewriter": modTablewriter.Module(),
		"time":        modTime.Module(),
		"yaml":        modYAML.Module(),
	}
	for k, v := range modHTTP.Builtins() {
		result[k] = v
	}
	for k, v := range modFmt.Builtins() {
		result[k] = v
	}
	for k, v := range builtins.Builtins() {
		result[k] = v
	}
	for k, v := range modOs.Builtins() {
		result[k] = v
	}
	return result
}
