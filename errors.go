package battlenet

import (
	"errors"
)

const ErrorNoInterfaceSupplied error = errors.New("Nil interface was passed to constructor.")
const ErrorUnsupportedArgument error = errors.New("Unsupported argument type was passed to constructor.")
