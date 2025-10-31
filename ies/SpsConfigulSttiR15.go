package ies

// SPS-ConfigUL-STTI-r15 ::= CHOICE
// Extensible
const (
	SpsConfigulSttiR15ChoiceNothing = iota
	SpsConfigulSttiR15ChoiceRelease
	SpsConfigulSttiR15ChoiceSetup
)

type SpsConfigulSttiR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SpsConfigulSttiR15Setup
}
