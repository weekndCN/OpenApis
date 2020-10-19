# Request Examples
```golang
package main

import (
	"fmt"
	loki "github.com/OpenApis/m/loki"

)

func main() {
	lokiURL := "http://101.132.252.188:30001/loki/api/v1/labels"
	loki := loki.NewLoki(lokiURL)
	err := loki.Query(lokiURL)
	if err != nil {
		fmt.Println(err)
	}
}
```