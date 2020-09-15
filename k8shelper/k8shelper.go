// Package k8shelper contains Kubernetes helpers.
package k8shelper

import (
	"os"
	"strings"
)

// FromUnder reports whether running from under Kubernetes.
func FromUnder() bool {
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "KUBERNETES_") {
			return true
		}
	}
	return false
}
