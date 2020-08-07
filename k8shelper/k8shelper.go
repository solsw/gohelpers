// Package k8shelper contains various Kubernetes helpers.
package k8shelper

import (
	"os"
	"strings"
)

// FromUnder reports whether running from under Kubernetes or not.
func FromUnder() bool {
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "KUBERNETES_") {
			return true
		}
	}
	return false
}
