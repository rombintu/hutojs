## HuToJs (Golang)
Делает из вывода OpenStack -> JSON или Map

Использование:
```go
package main

import (
    "fmt"
	"log"

    "github.com/rombintu/hutojs"
)

func main() {
    p := hutojs.NewParser()
	r, err := hutojs.ExecCommand(
		"os", "network", "agent", "list",
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := p.ToJson(r); err != nil {
		log.Fatal(err)
	}
    // Вывести JSON
	fmt.Println(p.JsonData)

    // Использование Map
    fmt.Println(p.Output)
}
```

Делаем скрипт для пайплайна:
```go
// ./cmd/main.go
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
```

```bash
$ go build -o hutojs ./cmd/main.go
$ os network agent list | ./hutojs
```