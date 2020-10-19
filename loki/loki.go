package loki

import (
	"fmt"
	"io/ioutil"
)

// Loki loki client
// https://grafana.com/docs/loki/latest/api/
type Loki struct {
	lokiURL string
}

// NewLoki return a new loki http client
func NewLoki(lokiURL string) Sender {
	return &Loki{lokiURL: lokiURL}
}

// Query loki query
func (l *Loki) Query(uri string) (err error) {
	//uri := r.URL.String()
	//fmt.Println(uri)
	resp, err := l.get(uri)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// out close resp body
	resp.Body.Close()
	fmt.Println(string(data))
	return
}

// QueryRange loki query_range api
func (l *Loki) QueryRange(uri string) (err error) {
	resp, err := l.get(uri)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// out close resp body
	resp.Body.Close()
	fmt.Println(string(data))
	return
}
