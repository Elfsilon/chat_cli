package env

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func errValueRequired(key string) error {
	return fmt.Errorf("env value \"%v\" is required and must not be empty", key)
}

func errIntConvertion(err error) error {
	err1 := fmt.Errorf("unable convert env value to int: ")
	return errors.Join(err1, err)
}

func OptionalString(key string) string {
	return os.Getenv(key)
}

func String(key string) string {
	res := os.Getenv(key)
	if res == "" {
		log.Fatal(errValueRequired(key))
	}
	return res
}

func Int(key string) int {
	res, err := strconv.Atoi(String(key))
	if err != nil {
		log.Fatal(errIntConvertion(err))
	}
	return res
}
