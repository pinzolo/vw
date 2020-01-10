package main

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestToVW(t *testing.T) {
	testdata := []struct {
		base  int
		value int64
		want  string
	}{
		{375, 75, "20"},
		{375, 125, "33.3333333333333333"},
		{500, 125, "25"},
	}
	for _, td := range testdata {
		t.Run(fmt.Sprintf("base: %d, value: %v", td.base, td.value), func(t *testing.T) {
			base = td.base
			d := decimal.New(td.value, 0)
			got := toVW(d)
			if got.String() != td.want {
				t.Errorf("want: %s, got: %v", td.want, got)
			}
		})
	}
}

func TestToPX(t *testing.T) {
	testdata := []struct {
		base  int
		value string
		want  string
	}{
		{375, "20", "75"},
		{375, "33.3333333333333333", "125"},
		{500, "25", "125"},
	}
	base = 375
	for _, td := range testdata {
		t.Run(fmt.Sprintf("base: %d, value: %s", td.base, td.value), func(t *testing.T) {
			base = td.base
			d, err := decimal.NewFromString(td.value)
			if err != nil {
				t.Error(err)
				return
			}
			got := toPX(d)
			if got != td.want {
				t.Errorf("want: %d, got: %v", td.want, got)
			}
		})
	}

}
