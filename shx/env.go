package shx

import "os"

func GetEnvVal(key, def string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return def
}
