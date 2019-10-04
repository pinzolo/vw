package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shopspring/decimal"
)

var base int

func main() {
	flag.IntVar(&base, "base", 375, "base width(px)")
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "value required")
		os.Exit(1)
	}

	args := flag.Args()
	v, err := decimal.NewFromString(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(v.Mul(decimal.New(100, 0)).Div(decimal.New(int64(base), 0)))
}
