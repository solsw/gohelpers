package e11hhelper

import (
	"bytes"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// PrimWriter writes to Elasticsearch index without any index options.
type PrimWriter struct {
	// Client - Elasticsearch client.
	Client *elasticsearch.Client
	// IdxName - name of Elasticsearch index to write to.
	IdxName string
	res     *esapi.Response
}

// Write implements the io.Writer interface.
func (w *PrimWriter) Write(p []byte) (int, error) {
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	var err error
	w.res, err = w.Client.Index(w.IdxName, bytes.NewReader(p))
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

// Info returns basic information about Elasticsearch cluster
// (https://www.elastic.co/guide/en/elasticsearch/reference/current/info-api.html).
func (w *PrimWriter) Info() (*esapi.Response, error) {
	return w.Client.Info()
}

// GetResponse returns last esapi.Response.
func (w *PrimWriter) GetResponse() *esapi.Response {
	return w.res
}

// WriteRes writes 'p' to Elasticsearch and returns esapi.Response.
func (w *PrimWriter) WriteRes(p []byte) (*esapi.Response, error) {
	if _, err := w.Write(p); err != nil {
		return nil, err
	}
	return w.res, nil
}
