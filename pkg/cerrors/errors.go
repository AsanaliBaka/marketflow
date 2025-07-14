package cerrors

import "fmt"

var (
	ErrorConfigNotFound = func(envName string) error {
		return fmt.Errorf("environment variable not found: %s", envName)
	}
)
