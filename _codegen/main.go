package main

import (
	"archive/zip"
	"cmp"
	"context"
	"errors"
	"flag"
	"fmt"
	"go/doc"
	"go/types"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"golang.org/x/tools/go/packages"
)

const (
	targetCommit = "a53be35c3b0cfcd5189cffcfd75df60ea581104c"
	testifyURL   = "https://github.com/stretchr/testify/archive/" + targetCommit + ".zip"
)

const headerComment = "Code generated from " + testifyURL + "; DO NOT EDIT."

var tempDir = filepath.Join(os.TempDir(), "testo-allure")

type Options struct {
	Pkg  string
	Path string
}

func main() {
	options, err := parseOptions()
	if err != nil {
		log.Fatal(err)
	}

	if err := run(context.Background(), options); err != nil {
		log.Fatal(err)
	}
}

func parseOptions() (Options, error) {
	pkg := flag.String("pkg", "", "output pkg name")
	path := flag.String("path", "", "output file")

	flag.Parse()

	if *pkg == "" {
		return Options{}, errors.New("missing pkg")
	}

	if *path == "" {
		return Options{}, errors.New("missing path")
	}

	return Options{
		Pkg:  *pkg,
		Path: *path,
	}, nil
}

func run(ctx context.Context, options Options) error {
	srcDir, err := getTestifySourcePath(ctx)
	if err != nil {
		return fmt.Errorf("get testify source path: %w", err)
	}

	pkg, err := parseAssertPkg(ctx, srcDir)
	if err != nil {
		return fmt.Errorf("parse assert pkg: %w", err)
	}

	assertions, err := getAssertions(pkg)
	if err != nil {
		return fmt.Errorf("get assertions: %w", err)
	}

	slices.SortFunc(assertions, func(a, b Assertion) int {
		return cmp.Compare(a.Name, b.Name)
	})

	f := jen.NewFile(options.Pkg)
	f.HeaderComment(headerComment)

	generateAssertions(f, assertions)

	err = f.Save(options.Path)
	if err != nil {
		return fmt.Errorf("save: %w", err)
	}

	return nil
}

func generateAssertions(f *jen.File, assertions []Assertion) {
	for _, a := range assertions {
		generateAssertion(f, a, true)
		generateAssertion(f, a, false)
	}
}

func formatComment(s string) string {
	var b strings.Builder

	lines := strings.Split(s, "\n")

	for i, line := range lines {
		isLast := i == len(lines)-1

		if !isLast {
			b.WriteString("// " + line + "\n")
		}
	}

	return b.String()
}

func generateAssertion(f *jen.File, assertion Assertion, require bool) {
	const recv = "x"

	stmt := f.Comment(formatComment(assertion.Doc)).Func()

	if require {
		stmt = stmt.Params(jen.Id(recv).Id("Requirements"))
	} else {
		stmt = stmt.Params(jen.Id(recv).Id("Assertions"))
	}

	stmt = stmt.Id(assertion.Name).ParamsFunc(func(g *jen.Group) {
		for i, p := range assertion.Params {
			// handle last differently
			if i == len(assertion.Params)-1 {
				g.Id(p.Name()).Op("...").Add(convertType(p.Type().(*types.Slice).Elem()))
				break
			}

			g.Id(p.Name()).Add(convertType(p.Type()))
		}
	})

	if !require {
		stmt = stmt.Bool()
	}

	stmt.BlockFunc(func(g *jen.Group) {
		g.Id(recv).Dot("t").Dot("Helper").Call()

		g.Id("callerTrace").
			Op(":=").
			Qual("github.com/ozontech/testo-allure/internal/stacktrace", "Take").
			Call(jen.Lit(1))

		g.List(
			jen.Id("_"),
			jen.Id("callerFile"),
			jen.Id("callerLine"),
			jen.Id("_"),
		).Op(":=").Qual("runtime", "Caller").Call(jen.Lit(1))

		g.Id("name").
			Op(":=").
			Qual("cmp", "Or").
			Call(
				jen.Id("messageFromMsgAndArgs").Call(jen.Id("msgAndArgs").Op("...")),
				jen.Lit(camelCaseToSentence(assertion.Name)),
			)

		var stmt *jen.Statement

		if require {
			stmt = g.Id("Step")
		} else {
			stmt = g.Return().Qual("github.com/ozontech/testo", "Run")
		}

		stmt.Call(
			jen.Id(recv).Dot("t"),
			jen.LitFunc(func() any {
				if require {
					return "require: "
				}

				return "assert: "
			}).Op("+").Id("name"),
			jen.Func().Params(jen.Id("t").Id("*PluginAllure")).BlockFunc(func(g *jen.Group) {
				g.Id("t").Dot("Helper").Call()

				g.Defer().Func().Params().BlockFunc(func(g *jen.Group) {
					g.Id("t").Dot("Parameters").CallFunc(func(g *jen.Group) {
						g.Id("NewParameter").
							Call(
								jen.LitFunc(func() any {
									if require {
										return "require"
									}

									return "assert"
								}),
								jen.Lit(camelCaseToSentence(assertion.Name)),
							)

						for i, p := range assertion.Params {
							// ignore last param
							if i == len(assertion.Params)-1 {
								break
							}

							name := p.Name()

							g.Id("NewParameter").
								Call(
									jen.Lit(camelCaseToSentence(name)),
									jen.Id("asShortString").Call(jen.Id(name)),
								).
								Dot("withMode").
								Call(jen.Id(recv).Dot("mode"))
						}
					})
				}).Call()

				assert := jen.
					Qual("github.com/stretchr/testify/assert", assertion.Name).
					CallFunc(func(g *jen.Group) {
						g.Id("t")

						for i, p := range assertion.Params {
							// handle last differently
							if i == len(assertion.Params)-1 {
								g.Id(p.Name()).Op("...")
								break
							}

							g.Id(p.Name())
						}
					})

				g.If(assert).Block(jen.Return())

				g.Id("callerLog").Op(":=").Qual("fmt", "Sprintf").Call(
					jen.Lit("Caller: %s:%d"),
					jen.Id("callerFile"),
					jen.Id("callerLine"),
				)

				g.Id("t").Dot("Log").Call(jen.Id("callerLog"))
				g.Id("t").Dot("addMessage").Call(jen.Id("callerLog"))
				g.Id("t").Dot("addTrace").Call(jen.Id("callerTrace"))

				if require {
					g.Id("t").Dot("FailNow").Call()
				}
			}),
			jen.Id("asAssertion").Call(),
		)
	})
}

type Assertion struct {
	Doc    string
	Name   string
	Params []*types.Var
}

func getAssertions(assertPkg *packages.Package) ([]Assertion, error) {
	scope := assertPkg.Types.Scope()

	documentation, err := doc.NewFromFiles(assertPkg.Fset, assertPkg.Syntax, assertPkg.PkgPath)
	if err != nil {
		return nil, err
	}

	funcDocByName := make(map[string]string, len(documentation.Funcs))
	for _, f := range documentation.Funcs {
		funcDocByName[f.Name] = f.Doc
	}

	testingT := scope.Lookup("TestingT").Type().Underlying().(*types.Interface)

	// list of functions to ignore when generating
	ignore := map[string]struct{}{
		"Fail":            {},
		"FailNow":         {},
		"IsNotType":       {},
		"EventuallyWithT": {},
	}

	var assertions []Assertion

	for _, name := range scope.Names() {
		if _, ok := ignore[name]; ok {
			continue
		}

		// conditions copied from https://github.com/stretchr/testify/blob/a53be35c3b0cfcd5189cffcfd75df60ea581104c/_codegen/main.go#L137

		obj := scope.Lookup(name)

		sig, ok := obj.Type().(*types.Signature)
		if !ok {
			continue
		}

		// Skip private functions
		if !obj.Exported() {
			continue
		}

		// Check function signature has at least two arguments
		if sig.Params().Len() < 2 {
			continue
		}

		// Skip functions ending with f
		if strings.HasSuffix(name, "f") {
			continue
		}

		// Check first argument is of type testingT
		if !types.Implements(sig.Params().At(0).Type(), testingT) {
			continue
		}

		params := make([]*types.Var, 0, sig.Params().Len())
		for i := range sig.Params().Len() {
			// skip first T param
			if i == 0 {
				continue
			}

			params = append(params, sig.Params().At(i))
		}

		assertions = append(assertions, Assertion{
			Doc:    funcDocByName[name],
			Name:   name,
			Params: params,
		})
	}

	return assertions, nil
}

func parseAssertPkg(ctx context.Context, path string) (*packages.Package, error) {
	cfg := packages.Config{
		Mode:    packages.LoadAllSyntax,
		Context: ctx,
		Dir:     path,
	}

	pkgs, err := packages.Load(&cfg, "./assert")
	if err != nil {
		return nil, err
	}

	if len(pkgs) != 1 {
		return nil, fmt.Errorf("unexpected packages count: %d", len(pkgs))
	}

	return pkgs[0], nil
}

func closeOrPanic(c io.Closer) {
	if err := c.Close(); err != nil {
		panic(err)
	}
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer closeOrPanic(r)

	err = os.MkdirAll(dest, 0o755)
	if err != nil {
		return err
	}

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer closeOrPanic(rc)

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			return os.MkdirAll(path, f.Mode())
		}

		err = os.MkdirAll(filepath.Dir(path), f.Mode())
		if err != nil {
			return err
		}

		regularFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer closeOrPanic(regularFile)

		_, err = io.Copy(regularFile, rc)
		if err != nil {
			return err
		}

		return nil
	}

	for _, f := range r.File {
		if err := extractAndWriteFile(f); err != nil {
			return err
		}
	}

	return nil
}

func cachePath() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "qa_pack_gen", targetCommit+".zip"), nil
}

func getTestifySourcePath(ctx context.Context) (string, error) {
	archivePath, err := getTestifyArchivePath(ctx)
	if err != nil {
		return "", fmt.Errorf("get testify archive path: %w", err)
	}

	err = unzip(archivePath, tempDir)
	if err != nil {
		return "", fmt.Errorf("unzip: %w", err)
	}

	return filepath.Join(tempDir, "testify-"+targetCommit), nil
}

func getTestifyArchivePath(ctx context.Context) (string, error) {
	path, err := cachePath()
	if err != nil {
		return "", err
	}

	_, err = os.Stat(path)
	if err == nil {
		return path, nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return "", err
	}

	err = downloadSourceTo(ctx, path)
	if err != nil {
		return "", fmt.Errorf("download source to: %w", err)
	}

	return path, nil
}

func downloadSourceTo(ctx context.Context, path string) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		testifyURL,
		nil,
	)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o600)
	if err != nil {
		return err
	}
	defer closeOrPanic(file)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer closeOrPanic(res.Body)

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return nil
}
