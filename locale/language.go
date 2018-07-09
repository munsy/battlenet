package locale

type language uint16

func (l language) ToString() string {
	if 0 == l {
		return "und"
	}
}
