package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"

	"github.com/shopspring/decimal"
)

var base int
var copy bool
var rev bool

func main() {
	//flag.IntVar(&base, "base", 375, "base width(px)")
	flag.IntVar(&base, "base", 1280, "base width(px)")
	flag.BoolVar(&copy, "copy", false, "copy to clipboard")
	flag.BoolVar(&rev, "reverse", false, "calc to px")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "value required")
		os.Exit(1)
	}

	v, err := decimal.NewFromString(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var output string
	if rev {
		output = fmt.Sprintf("%vpx", toPX(v))
	} else {
		output = fmt.Sprintf("%vvw", toVW(v))
	}
	if copy {
		clipboard.WriteAll(strings.TrimSpace(output))
	}
	fmt.Println(output)
}

func toVW(d decimal.Decimal) decimal.Decimal {
	return d.Mul(decimal.New(100, 0)).Div(decimal.New(int64(base), 0))
}

func toPX(d decimal.Decimal) int64 {
	v := d.Mul(decimal.New(int64(base), 0)).Div(decimal.New(100, 0))
	return v.Round(0).IntPart()
}
