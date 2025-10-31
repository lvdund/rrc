package ies

// SPS-ConfigUL-STTI-r15-setup-p0-PersistentSubframeSet2-r15 ::= CHOICE
const (
	SpsConfigulSttiR15SetupP0Persistentsubframeset2R15ChoiceNothing = iota
	SpsConfigulSttiR15SetupP0Persistentsubframeset2R15ChoiceRelease
	SpsConfigulSttiR15SetupP0Persistentsubframeset2R15ChoiceSetup
)

type SpsConfigulSttiR15SetupP0Persistentsubframeset2R15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SpsConfigulSttiR15SetupP0Persistentsubframeset2R15Setup
}
