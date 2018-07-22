package locale

import (
	"errors"
	"fmt"
)

// Locale represents a Battle.net-defined locale.
type Locale struct {
	dialect  uint16
	language uint16
}

func (l Locale) String() string {
	if int(l.dialect) > len(dialects) {
		panic(errors.New("Invalid dialect format."))
	}
	if int(l.language) > len(languages) {
		panic(errors.New("Invalid language format."))
	}

	return fmt.Sprintf("%s_%s", dialects[l.dialect], languages[l.language])
}
