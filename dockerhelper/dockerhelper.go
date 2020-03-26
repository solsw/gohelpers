package dockerhelper

import (
	"os"
)

// InDocker reports whether running from Docker container.
func InDocker() bool {
	// see https://github.com/sindresorhus/is-docker/blob/master/index.js
	_, err := os.Stat("/.dockerenv")
	return err == nil
}
