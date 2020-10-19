package loki

// Sender return a loki interface
type Sender interface {
	Query(uri string) error
	QueryRange(uri string) error
}
