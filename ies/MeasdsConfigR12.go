package ies

// MeasDS-Config-r12 ::= CHOICE
// Extensible
const (
	MeasdsConfigR12ChoiceNothing = iota
	MeasdsConfigR12ChoiceRelease
	MeasdsConfigR12ChoiceSetup
)

type MeasdsConfigR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *MeasdsConfigR12Setup
}
