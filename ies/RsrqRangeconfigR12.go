package ies

// RSRQ-RangeConfig-r12 ::= CHOICE
const (
	RsrqRangeconfigR12ChoiceNothing = iota
	RsrqRangeconfigR12ChoiceRelease
	RsrqRangeconfigR12ChoiceSetup
)

type RsrqRangeconfigR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RsrqRangeV1250
}
