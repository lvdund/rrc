package ies

// SPS-ConfigDL ::= CHOICE
// Extensible
const (
	SpsConfigdlChoiceNothing = iota
	SpsConfigdlChoiceRelease
	SpsConfigdlChoiceSetup
)

type SpsConfigdl struct {
	Choice  uint64
	Release *struct{}
	Setup   *SpsConfigdlSetup
}
