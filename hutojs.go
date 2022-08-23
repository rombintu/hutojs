package hutojs

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Parser struct {
	JsonData string
	Output   map[string][]string
}

func NewParser() *Parser {
	return &Parser{
		Output: make(map[string][]string),
	}
}

func (p *Parser) ToOutput(r io.Reader) {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), "\n")
		for _, l := range lines {
			line := strings.Split(l, " | ")
			p.Output[trim(line[0])] = trimArr(line[1:])
		}
	}
}

func (p *Parser) ToJson(r io.Reader) error {
	p.ToOutput(r)
	js, err := json.Marshal(p.Output)
	if err != nil {
		return err
	}
	p.JsonData = string(js)
	return nil
}

func trim(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return r == '|' || r == ' '
	})
}

func trimArr(s []string) []string {
	var arr []string
	for _, a := range s {
		arr = append(arr, trim(a))
	}
	return arr
}

func ExecCommand(command string, params ...string) (*bytes.Reader, error) {
	cmd := exec.Command(command, params...)
	out, err := cmd.Output()
	// if err != nil {
	// 	return &bytes.Reader{}, err
	// }
	fmt.Printf("%s\n", out)
	return bytes.NewReader(out), err
}
