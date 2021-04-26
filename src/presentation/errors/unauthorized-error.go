package presentationerrors

import "errors"

func UnauthorizedError() error {
	return errors.New("Unauthorized")
}
