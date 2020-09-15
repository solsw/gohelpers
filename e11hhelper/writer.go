package e11hhelper

import (
	"bytes"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

// Writer writes to Elasticsearch index.
type Writer struct {
	cl  *elasticsearch.Client
	idx string
	res *esapi.Response
}

// NewWriter creates new Writer.
// 'idx' - name of Elasticsearch index to write to.
func NewWriter(cfg elasticsearch.Config, idx string) (*Writer, error) {
	cl, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Writer{cl: cl, idx: idx}, nil
}

// Write implements the io.Writer interface.
func (w *Writer) Write(p []byte) (int, error) {
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	var err error
	w.res, err = w.cl.Index(w.idx, bytes.NewReader(p))
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

// Info returns basic information about Elasticsearch cluster
// (https://www.elastic.co/guide/en/elasticsearch/reference/current/info-api.html).
func (w *Writer) Info() (*esapi.Response, error) {
	return w.cl.Info()
}

// GetResponse returns last esapi.Response.
func (w *Writer) GetResponse() *esapi.Response {
	return w.res
}

// WriteRes writes 'p' to Elasticsearch and returns esapi.Response.
func (w *Writer) WriteRes(p []byte) (*esapi.Response, error) {
	if _, err := w.Write(p); err != nil {
		return nil, err
	}
	return w.res, nil
}
