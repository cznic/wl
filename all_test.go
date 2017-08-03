// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"bufio"
	"flag"
	"fmt"
	"go/token"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/cznic/golex/lex"
)

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "# caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# \tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func TODO(...interface{}) string { //TODOOK
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("# TODO: %s:%d:\n", path.Base(fn), fl) //TODOOK
}

func use(...interface{}) {}

func init() {
	use(caller, dbg, TODO) //TODOOK
}

// ============================================================================

func init() {
	flag.IntVar(&yyDebug, "yydebug", 0, "")
}

func exampleAST(rule int, src string) interface{} {
	l, err := lex.New(
		token.NewFileSet().AddFile(fmt.Sprint(rule), -1, len(src)),
		strings.NewReader(src),
		lex.BOMMode(lex.BOMIgnoreFirst),
		lex.RuneClass(runeClass),
		lex.ErrorFunc(func(token.Pos, string) {}),
	)
	if err != nil {
		return err.Error()
	}

	lx, err := newLexer(l)
	if err != nil {
		return err.Error()
	}

	lx.exampleRule = rule
	lx.parse()
	return prettyString(lx.exampleAST)
}

func Test(t *testing.T) {
	t.Logf("TODO")
}

func testScannerCorpus(t *testing.T) {
	f, err := os.Open(filepath.Join("testdata", "corpus"))
	if err != nil {
		t.Fatal(err)
	}

	fi, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	file := token.NewFileSet().AddFile(f.Name(), -1, int(fi.Size()))
	r := bufio.NewReader(f)
	l, err := lex.New(
		file,
		r,
		lex.BOMMode(lex.BOMIgnoreFirst),
		lex.RuneClass(runeClass),
		lex.ErrorFunc(func(pos token.Pos, msg string) { t.Fatalf("%s: %s", file.Position(pos), msg) }),
	)
	if err != nil {
		t.Fatal(err)
	}

	lx, err := newLexer(l)
	if err != nil {
		t.Fatal(err)
	}

	toks := 0
	for lx.Last.Rune >= 0 {
		lx.scan()
		toks++
	}
	if _, err := r.Peek(1); err != io.EOF {
		t.Fatal(err)
	}

	t.Logf("tokens: %v", toks)
}

func TestScanner(t *testing.T) {
	t.Run("corpus", testScannerCorpus)
}

func testParserCorpus(t *testing.T) {
	f, err := os.Open(filepath.Join("testdata", "corpus"))
	if err != nil {
		t.Fatal(err)
	}

	fi, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	var lx *lexer

	file := token.NewFileSet().AddFile(f.Name(), -1, int(fi.Size()))
	r := bufio.NewReader(f)
	l, err := lex.New(
		file,
		r,
		lex.BOMMode(lex.BOMIgnoreFirst),
		lex.RuneClass(runeClass),
		lex.ErrorFunc(lx.errPos),
	)
	if err != nil {
		t.Fatal(err)
	}
	n := 0
	for {
		_, err := r.Peek(1)
		if err != nil {
			break
		}

		lx, err := newLexer(l)
		if err != nil {
			t.Fatal(err)
		}

		if err = lx.parse(); err != nil {
			t.Fatal(err)
		}

		n++
	}
	t.Logf("expressions: %v", n)
}

func TestParser(t *testing.T) {
	t.Run("corpus", testParserCorpus)
}
