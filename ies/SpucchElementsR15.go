package ies

// SPUCCH-Elements-r15 ::= CHOICE
const (
	SpucchElementsR15ChoiceNothing = iota
	SpucchElementsR15ChoiceRelease
	SpucchElementsR15ChoiceSetup
)

type SpucchElementsR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SpucchElementsR15Setup
}
