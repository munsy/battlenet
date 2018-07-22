package errors

import (
	"errors"
)

// Battle.net errors for clients and various packages.
var (
	ErrNoInterfaceSupplied error = errors.New("Nil interface was passed to constructor.")
	ErrUnsupportedArgument error = errors.New("Unsupported argument type was passed to constructor.")
	ErrUnresolvedEndpoint  error = errors.New("Unresolved endpoint.")
)
