package formaterror

import (
	"errors"
)

// Parse error structure
func Parse(err string) error {
	return errors.New(err)
}
