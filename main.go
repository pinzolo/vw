package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/shopspring/decimal"
)

var (
	base      int
	noCopy    bool
	rev       bool
	file      string
	overwrite bool
)

func main() {
	//flag.IntVar(&base, "base", 375, "base width(px)")
	flag.IntVar(&base, "base", 1280, "base width(px)")
	flag.BoolVar(&noCopy, "no-copy", false, "not noCopy to clipboard")
	flag.BoolVar(&rev, "reverse", false, "calc to px")
	flag.BoolVar(&rev, "r", false, "calc to px")
	flag.StringVar(&file, "file", "", "target file")
	flag.StringVar(&file, "f", "", "target file")
	flag.BoolVar(&overwrite, "overwrite", false, "overwrite file")
	flag.BoolVar(&overwrite, "o", false, "overwrite file")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 && file == "" {
		fmt.Fprintln(os.Stderr, "value required")
		os.Exit(1)
	}

	if file != "" {
		err := handleFile()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		handleSingleValue(args[0])
	}
}

func getSourceUnit() string {
	if rev {
		return "vw"
	}
	return "px"
}

func getDestUnit() string {
	if rev {
		return "px"
	}
	return "vw"
}

func toVW(d decimal.Decimal) decimal.Decimal {
	return d.Mul(decimal.New(100, 0)).Div(decimal.New(int64(base), 0))
}

func toPX(d decimal.Decimal) decimal.Decimal {
	v := d.Mul(decimal.New(int64(base), 0)).Div(decimal.New(100, 0))
	return v.Round(0)
}

func toFunc() func(decimal.Decimal) decimal.Decimal {
	if rev {
		return toPX
	}
	return toVW
}

func handleSingleValue(input string) {
	v, err := decimal.NewFromString(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	output := fmt.Sprintf("%v%s", toFunc()(v), getDestUnit())
	if !noCopy {
		clipboard.WriteAll(strings.TrimSpace(output))
	}
	fmt.Println(output)
}

func handleFile() error {
	f, err := os.OpenFile(file, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	re := regexp.MustCompile(`\d+(\.\d+)?` + getSourceUnit())
	r := bufio.NewReader(f)
	buff := bytes.Buffer{}
	for {
		line, _, err := r.ReadLine()
		buff.WriteString(re.ReplaceAllStringFunc(string(line), convertValue) + "\n")
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
	}

	if overwrite {
		_, err = f.WriteString(buff.String())
		if err != nil {
			return err
		}
	} else {
		fmt.Println(buff.String())
	}
	return nil
}

func convertValue(v string) string {
	num, _ := decimal.NewFromString(dropUnit(v))
	return toFunc()(num).String() + getDestUnit()
}

func dropUnit(s string) string {
	return s[0 : len(s)-len(getSourceUnit())]
}
