package errors

import (
	"errors"
)

var (
	NoInterfaceSupplied error = errors.New("Nil interface was passed to constructor.")
	UnsupportedArgument error = errors.New("Unsupported argument type was passed to constructor.")
)
