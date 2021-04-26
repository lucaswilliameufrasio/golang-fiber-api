package presentationerrors

import (
	"errors"
	"fmt"
)

func MissingParamError(param string) error {
	return errors.New(fmt.Sprintf("Missing param: %s", param))
}
