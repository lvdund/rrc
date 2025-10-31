package ies

// SPDCCH-Elements-r15 ::= CHOICE
// Extensible
const (
	SpdcchElementsR15ChoiceNothing = iota
	SpdcchElementsR15ChoiceRelease
	SpdcchElementsR15ChoiceSetup
)

type SpdcchElementsR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SpdcchElementsR15Setup
}
