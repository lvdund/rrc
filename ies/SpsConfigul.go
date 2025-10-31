package ies

// SPS-ConfigUL ::= CHOICE
// Extensible
const (
	SpsConfigulChoiceNothing = iota
	SpsConfigulChoiceRelease
	SpsConfigulChoiceSetup
)

type SpsConfigul struct {
	Choice  uint64
	Release *struct{}
	Setup   *SpsConfigulSetup
}
