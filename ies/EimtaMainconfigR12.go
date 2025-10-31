package ies

// EIMTA-MainConfig-r12 ::= CHOICE
const (
	EimtaMainconfigR12ChoiceNothing = iota
	EimtaMainconfigR12ChoiceRelease
	EimtaMainconfigR12ChoiceSetup
)

type EimtaMainconfigR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *EimtaMainconfigR12Setup
}
