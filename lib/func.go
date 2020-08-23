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
	return errors.New(msg)
}

func GenerateErrorWithStack(msg string) error  {
	// errros.New も withstackを返すので
	// stackが二重になる
	return errors.WithStack(errors.New(msg))
}
