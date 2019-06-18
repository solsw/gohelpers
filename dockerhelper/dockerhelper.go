package dockerhelper

import "os"

// FromDocker reports whether running from Docker container.
func FromDocker() bool {
	// see https://github.com/sindresorhus/is-docker/blob/master/index.js
	_, err := os.Stat("/.dockerenv")
	return err == nil
}
