package ies

// SPS-ConfigDL-STTI-r15 ::= CHOICE
// Extensible
const (
	SpsConfigdlSttiR15ChoiceNothing = iota
	SpsConfigdlSttiR15ChoiceRelease
	SpsConfigdlSttiR15ChoiceSetup
)

type SpsConfigdlSttiR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SpsConfigdlSttiR15Setup
}
