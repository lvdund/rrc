package ies

// PhysicalConfigDedicatedSTTI-r15 ::= CHOICE
const (
	PhysicalconfigdedicatedsttiR15ChoiceNothing = iota
	PhysicalconfigdedicatedsttiR15ChoiceRelease
	PhysicalconfigdedicatedsttiR15ChoiceSetup
)

type PhysicalconfigdedicatedsttiR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PhysicalconfigdedicatedsttiR15Setup
}
