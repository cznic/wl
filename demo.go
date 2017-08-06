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
	"os"

	"github.com/cznic/wl"
)

func main() {
	fmt.Printf("Enter WL expression(s). Newlines will be ignored in places where the input is not valid.\n")
	fmt.Printf("Closing the input exits the program.\n")
	for n := 1; ; n++ {
		fmt.Printf("In[%v]:= ", n)
		in, err := wl.NewInput(bufio.NewReader(os.Stdin), true)
		if err != nil {
			panic(err)
		}

		expr, err := in.ParseExpression(token.NewFileSet().AddFile(os.Stdin.Name(), -1, 1e6))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(expr)
	}
}
