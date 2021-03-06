package loki_test

import (
	"fmt"

	loki "github.com/OpenApis/m/loki"
)

func testQuery() {
	lokiURL := "http://xxxx/loki/api/v1/labels"
	loki := loki.NewLoki(lokiURL)
	err := loki.Query(lokiURL)
	if err != nil {
		fmt.Println(err)
	}
}

func testQueryRange() {
	lokiURL := "http://xxxx/loki/api/v1/labels"
	loki := loki.NewLoki(lokiURL)
	err := loki.QueryRange(lokiURL)
	if err != nil {
		fmt.Println(err)
	}
}
