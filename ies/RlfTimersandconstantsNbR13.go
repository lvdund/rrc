package ies

// RLF-TimersAndConstants-NB-r13 ::= CHOICE
// Extensible
const (
	RlfTimersandconstantsNbR13ChoiceNothing = iota
	RlfTimersandconstantsNbR13ChoiceRelease
	RlfTimersandconstantsNbR13ChoiceSetup
)

type RlfTimersandconstantsNbR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RlfTimersandconstantsNbR13Setup
}
