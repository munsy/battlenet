package locale

import (
	"errors"
	"fmt"
)

type Locale struct {
	dialect  uint16
	language uint16
}

func (l Locale) String() string {
	if l.dialect > len(dialects) {
		panic(errors.New("Invalid dialect format."))
	}
	if l.language > len(languages) {
		panic(errors.New("Invalid language format."))
	}

	return fmt.Sprintf("%s_%s", dialects[l.dialect], languages[l.language])
}
