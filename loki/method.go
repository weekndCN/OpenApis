package loki

import "net/http"

func (l *Loki) get(uri string) (resp *http.Response, err error) {
	resp, err = http.Get(uri)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *Loki) post(uri string) (resp *http.Response, err error) {
	resp, err = http.Post(uri, "application/json", nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
