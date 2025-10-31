package ies

// SlotOrSubslotPDSCH-Config-r15 ::= CHOICE
// Extensible
const (
	SlotorsubslotpdschConfigR15ChoiceNothing = iota
	SlotorsubslotpdschConfigR15ChoiceRelease
	SlotorsubslotpdschConfigR15ChoiceSetup
)

type SlotorsubslotpdschConfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SlotorsubslotpdschConfigR15Setup
}
