package e11hhelper

import (
	"bytes"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

// Writer implements io.Writer interface for writing to Elasticsearch.
type Writer struct {
	cl  *elasticsearch.Client
	idx string
	res *esapi.Response
}

// NewWriter creates new Writer.
func NewWriter(cfg elasticsearch.Config, idx string) (*Writer, error) {
	cl, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Writer{cl: cl, idx: idx}, nil
}

// Info returns basic information about Elasticsearch cluster
// (https://www.elastic.co/guide/en/elasticsearch/reference/current/info-api.html).
func (w *Writer) Info() (*esapi.Response, error) {
	return w.cl.Info()
}

// Write implements io.Writer interface.
func (w *Writer) Write(p []byte) (int, error) {
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	r, err := w.cl.Index(w.idx, bytes.NewReader(p))
	if err != nil {
		return 0, err
	}
	w.res = r
	return len(p), nil
}

// GetResponse returns last successful esapi.Response.
func (w *Writer) GetResponse() *esapi.Response {
	return w.res
}
