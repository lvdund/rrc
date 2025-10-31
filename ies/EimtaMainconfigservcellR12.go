package ies

// EIMTA-MainConfigServCell-r12 ::= CHOICE
const (
	EimtaMainconfigservcellR12ChoiceNothing = iota
	EimtaMainconfigservcellR12ChoiceRelease
	EimtaMainconfigservcellR12ChoiceSetup
)

type EimtaMainconfigservcellR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *EimtaMainconfigservcellR12Setup
}
