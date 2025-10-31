package ies

// AntennaInfoDedicatedSTTI-r15 ::= CHOICE
const (
	AntennainfodedicatedsttiR15ChoiceNothing = iota
	AntennainfodedicatedsttiR15ChoiceRelease
	AntennainfodedicatedsttiR15ChoiceSetup
)

type AntennainfodedicatedsttiR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *AntennainfodedicatedsttiR15Setup
}
