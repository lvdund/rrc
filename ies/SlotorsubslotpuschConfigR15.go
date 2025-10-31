package ies

// SlotOrSubslotPUSCH-Config-r15 ::= CHOICE
// Extensible
const (
	SlotorsubslotpuschConfigR15ChoiceNothing = iota
	SlotorsubslotpuschConfigR15ChoiceRelease
	SlotorsubslotpuschConfigR15ChoiceSetup
)

type SlotorsubslotpuschConfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SlotorsubslotpuschConfigR15Setup
}
