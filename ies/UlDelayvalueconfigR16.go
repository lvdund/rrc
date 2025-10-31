package ies

// UL-DelayValueConfig-r16 ::= CHOICE
const (
	UlDelayvalueconfigR16ChoiceNothing = iota
	UlDelayvalueconfigR16ChoiceRelease
	UlDelayvalueconfigR16ChoiceSetup
)

type UlDelayvalueconfigR16 struct {
	Choice  uint64
	Release *struct{}
	Setup   *UlDelayvalueconfigR16Setup
}
