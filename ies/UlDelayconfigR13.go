package ies

// UL-DelayConfig-r13 ::= CHOICE
const (
	UlDelayconfigR13ChoiceNothing = iota
	UlDelayconfigR13ChoiceRelease
	UlDelayconfigR13ChoiceSetup
)

type UlDelayconfigR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *UlDelayconfigR13Setup
}
