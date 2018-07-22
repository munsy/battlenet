package errors

import (
	"errors"
)

// Battle.net errors for clients and various packages.
var (
	NoInterfaceSupplied error = errors.New("Nil interface was passed to constructor.")
	UnsupportedArgument error = errors.New("Unsupported argument type was passed to constructor.")
	UnresolvedEndpoint  error = errors.New("Unresolved endpoint.")
)
