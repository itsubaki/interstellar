package util

import "os"

func Getenv(key, init string) string {
	val := os.Getenv(key)
	if len(val) > 0 {
		return val
	}

	return init
}
