package ies

// TDD-PUSCH-UpPTS-r14 ::= CHOICE
const (
	TddPuschUpptsR14ChoiceNothing = iota
	TddPuschUpptsR14ChoiceRelease
	TddPuschUpptsR14ChoiceSetup
)

type TddPuschUpptsR14 struct {
	Choice  uint64
	Release *struct{}
	Setup   *TddPuschUpptsR14Setup
}
