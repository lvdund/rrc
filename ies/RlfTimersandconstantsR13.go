package ies

// RLF-TimersAndConstants-r13 ::= CHOICE
// Extensible
const (
	RlfTimersandconstantsR13ChoiceNothing = iota
	RlfTimersandconstantsR13ChoiceRelease
	RlfTimersandconstantsR13ChoiceSetup
)

type RlfTimersandconstantsR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RlfTimersandconstantsR13Setup
}
