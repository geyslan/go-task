package utils

import "os"

// GetEnv gets value from Environment
// It returns the defined fallback if env is not found
func GetEnv(env, fallback string) string {
	env, found := os.LookupEnv(env)
	if !found {
		return fallback
	}

	return env
}
