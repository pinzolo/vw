package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shopspring/decimal"
)

var base int
var rev bool

func main() {
	flag.IntVar(&base, "base", 375, "base width(px)")
	flag.BoolVar(&rev, "reverse", false, "calc to px")
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
	if rev {
		fmt.Printf("%vpx\n", toPX(v))
	} else {
		fmt.Printf("%vvw\n", toVW(v))
	}
}

func toVW(d decimal.Decimal) decimal.Decimal {
	return d.Mul(decimal.New(100, 0)).Div(decimal.New(int64(base), 0))
}

func toPX(d decimal.Decimal) int64 {
	v := d.Mul(decimal.New(int64(base), 0)).Div(decimal.New(100, 0))
	return v.Round(0).IntPart()
}
