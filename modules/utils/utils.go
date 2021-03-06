package utils

import (
	"errors"
	"os"
	"strconv"
)

// GenEvnStr returns validated string env value
func GetEnvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, errors.New("empty env value")
	}
	return v, nil
}

// GenEvnStr returns validated int env value
func GetEnvInt(key string) (int, error) {
	s, err := GetEnvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}
