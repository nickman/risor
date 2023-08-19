package importer

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/risor-io/risor/compiler"
	"github.com/risor-io/risor/object"
	"github.com/risor-io/risor/parser"
)

// Importer is an interface used to import Risor code modules
type Importer interface {

	// Import a module by name
	Import(ctx context.Context, name string) (*object.Module, error)
}

type LocalImporter struct {
	globalNames []string
	codeCache   map[string]*compiler.Code
	sourceDir   string
	extensions  []string
}

// LocalImporterOptions configure an Importer that can read from the local
// filesystem.
type LocalImporterOptions struct {

	// Global names that should be available when the module is compiled.
	GlobalNames []string

	// The directory to search for Risor modules.
	SourceDir string

	// Optional list of file extensions to try when locating a Risor module.
	Extensions []string
}

// NewLocalImporter returns an Importer that can read Risor code modules from
// the local filesystem. Internally, loaded code is cached in memory. However,
// a new Module is created for each Import call. If the caller wants to reuse
// the same Module, it should be cached by the caller. It is safe to reuse the
// same local importer across multiple VMs and evaluations, because the cached
// code is immutable.
func NewLocalImporter(opts LocalImporterOptions) *LocalImporter {
	if opts.Extensions == nil {
		opts.Extensions = []string{".risor", ".rsr"}
	}
	return &LocalImporter{
		globalNames: opts.GlobalNames,
		codeCache:   map[string]*compiler.Code{},
		sourceDir:   opts.SourceDir,
		extensions:  opts.Extensions,
	}
}

func (i *LocalImporter) Import(ctx context.Context, name string) (*object.Module, error) {
	if code, ok := i.codeCache[name]; ok {
		return object.NewModule(name, code), nil
	}
	source, found := readFileWithExtensions(i.sourceDir, name, i.extensions)
	if !found {
		return nil, fmt.Errorf("module not found: %s", name)
	}
	ast, err := parser.Parse(ctx, source)
	if err != nil {
		return nil, err
	}
	var opts []compiler.Option
	if len(i.globalNames) > 0 {
		opts = append(opts, compiler.WithGlobalNames(i.globalNames))
	}
	code, err := compiler.Compile(ast, opts...)
	if err != nil {
		return nil, err
	}
	return object.NewModule(name, code), nil
}

func readFileWithExtensions(dir, name string, extensions []string) (string, bool) {
	for _, ext := range extensions {
		fullPath := filepath.Join(dir, name+ext)
		bytes, err := os.ReadFile(fullPath)
		if err == nil {
			return string(bytes), true
		}
	}
	return "", false
}
