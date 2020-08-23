package lib

import "errors"

func GenerateError(msg string) error  {
	return errors.New(msg)
}
