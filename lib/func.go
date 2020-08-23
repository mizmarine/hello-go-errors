package lib

import (
	simpleError "errors"
	"github.com/pkg/errors"
)

func GenerateSimpleError(msg string) error {
	// 一番シンプルなerrors.New
	return simpleError.New(msg)
}

func GenerateErrorFromPkgError(msg string) error  {
	// pkg.errors の new は withstack struct を返す
	// コレ使うのが一番良さそう
	return errors.New(msg)
}

func GenerateErrorWithStack(msg string) error  {
	// errros.New も withstackを返すので
	// stackが二重になってしまうw
	return errors.WithStack(errors.New(msg))
}
