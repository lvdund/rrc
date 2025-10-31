package ies

// SPS-ConfigDL-STTI-r15-setup-twoAntennaPortActivated-r15 ::= CHOICE
const (
	SpsConfigdlSttiR15SetupTwoantennaportactivatedR15ChoiceNothing = iota
	SpsConfigdlSttiR15SetupTwoantennaportactivatedR15ChoiceRelease
	SpsConfigdlSttiR15SetupTwoantennaportactivatedR15ChoiceSetup
)

type SpsConfigdlSttiR15SetupTwoantennaportactivatedR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SpsConfigdlSttiR15SetupTwoantennaportactivatedR15Setup
}
