package locale

import (
	"fmt"

	"github.com/munsy/gobattlenet/errors"
)

// Locale represents a Battle.net-defined locale.
type Locale struct {
	dialect  uint16
	language uint16
}

func (l Locale) String() string {
	if int(l.dialect) > len(dialects) {
		panic(errors.ErrInvalidDialectForm)
	}
	if int(l.language) > len(languages) {
		panic(errors.ErrInvalidDialectForm)
	}

	return fmt.Sprintf("%s_%s", dialects[l.dialect], languages[l.language])
}
