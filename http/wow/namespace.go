package wow

type NamespaceCategory int

const(
	STATIC NamespaceCategory = iota,
	DYNAMIC,
	PROFILE
)

type WoWNamespace interface {
	String() string
}

func Namespace() WoWNamespace {

}