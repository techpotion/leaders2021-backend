package utils

import (
	"errors"
	"os"
	"strconv"
)

func GetEnvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, errors.New("empty env value")
	}
	return v, nil
}

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

func SetStrings(strs []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, str := range strs {
		if _, ok := keys[str]; !ok {
			keys[str] = true
			list = append(list, str)
		}
	}

	return list
}

func SetUint32s(elems []uint32) []uint32 {
	keys := make(map[uint32]bool)
	list := []uint32{}

	for _, elem := range elems {
		if _, ok := keys[elem]; !ok {
			keys[elem] = true
			list = append(list, elem)
		}
	}

	return list
}
