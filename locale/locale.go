package locale

import (
	"errors"
)

type Locale struct {
	dialect  uint16
	language uint16
}

func (l Locale) String() {
	if l.dialect > len(dialects) {
		panic(errors.New("Invalid dialect format."))
	}
	if l.language > len(languages) {
		panic(errors.New("Invalid language format."))
	}
}
