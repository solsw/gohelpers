package e11hhelper

import (
	"bytes"
	"context"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// AsyncWriter writes asynchronously to Elasticsearch index without any index options.
type AsyncWriter struct {
	// Client - Elasticsearch client.
	// https://pkg.go.dev/github.com/elastic/go-elasticsearch/v7#Client
	Client *elasticsearch.Client
	// IdxName - name of Elasticsearch index to write to.
	IdxName string
	// Timeout
	Timeout      time.Duration
	ResErrAction func(res *esapi.Response, err error)
}

// Write implements the io.Writer interface.
// Response and error returned by Elasticsearch are processed by ResErrAction, if the function is provided.
func (w *AsyncWriter) Write(p []byte) (int, error) {
	locp := make([]byte, len(p))
	copy(locp, p)
	go func(bb []byte) {
		ctx := context.Background()
		if w.Timeout > 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, w.Timeout)
			defer cancel()
		}
		// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
		// https://pkg.go.dev/github.com/elastic/go-elasticsearch/v7/esapi#Index
		res, err := w.Client.Index(w.IdxName, bytes.NewReader(bb), w.Client.Index.WithContext(ctx))
		if w.ResErrAction != nil {
			w.ResErrAction(res, err)
		}
	}(locp)
	return len(p), nil
}
