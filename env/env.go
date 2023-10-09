package env

import "os"

func HasEnv(name string) bool {
	_, exists := os.LookupEnv(name)
	return exists
}

func EqualEnv(name, equal string) bool {
	if !HasEnv(name) {
		return false
	}
	return os.Getenv(name) == equal
}
