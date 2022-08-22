package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rombintu/hutojs"
)

func main() {
	p := hutojs.NewParser()
	if err := p.ToJson(os.Stdin); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.JsonData)
}
