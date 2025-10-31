package ies

// RLF-TimersAndConstants-r9 ::= CHOICE
// Extensible
const (
	RlfTimersandconstantsR9ChoiceNothing = iota
	RlfTimersandconstantsR9ChoiceRelease
	RlfTimersandconstantsR9ChoiceSetup
)

type RlfTimersandconstantsR9 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RlfTimersandconstantsR9Setup
}
