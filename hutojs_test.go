package hutojs_test

import (
	"fmt"
	"testing"

	"github.com/rombintu/hutojs"
)

func TestParser(t *testing.T) {
	p := hutojs.NewParser()
	r, err := hutojs.ExecCommand(
		"cat", "testing.txt",
	)
	if err != nil {
		t.Error(err)
	}
	if err := p.ToJson(r); err != nil {
		t.Error(err)
	}
	fmt.Println(p.JsonData)
}
