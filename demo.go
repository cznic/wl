// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// $ go run demo.go

package main

import (
	"bufio"
	"fmt"
	"go/token"
	"io"
	"os"

	"github.com/cznic/wl"
)

var _ io.RuneReader = stdin(nil)

type char struct {
	c   rune
	sz  int
	err error
}

type stdin chan char

func (s stdin) ReadRune() (r rune, size int, err error) {
	c := <-s
	return c.c, c.sz, c.err
}

func main() {
	fmt.Printf("Enter WL expression(s). Newlines will be ignored in places where the input is not valid.\n")
	fmt.Printf("Closing the input exits the program\n")
	r := make(stdin, 100)
	in, err := wl.NewInput(r, true)
	if err != nil {
		panic(err)
	}

	go func() {
		buf := bufio.NewReader(os.Stdin)
		for {
			c, sz, err := buf.ReadRune()
			r <- char{c, sz, err}
			if c == '\n' {
				r <- char{' ', sz, err}
			}
			if err != nil {
				fmt.Println()
				os.Exit(0)
			}
		}
	}()

	for n := 1; ; n++ {
		fmt.Printf("In[%v]:= ", n)
		file := token.NewFileSet().AddFile(os.Stdin.Name(), -1, 1e6)
		expr, err := in.ParseExpression(file)
		if err != nil {
			fmt.Println(err)
		out:
			for {
				select {
				case <-r:
				default:
					r <- char{'\n', 1, nil}
					r <- char{'\n', 1, nil}
					break out
				}
			}
			continue
		}

		fmt.Println(expr)
	}
}
