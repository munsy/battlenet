package battlenet

import (
	"errors"
)

var ErrorNoInterfaceSupplied error = errors.New("Nil interface was passed to constructor.")
var ErrorUnsupportedArgument error = errors.New("Unsupported argument type was passed to constructor.")
var ErrorUnresolvedEndpoint error = errors.New("Unresolved endpoint.")
