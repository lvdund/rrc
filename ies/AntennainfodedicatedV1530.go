package ies

// AntennaInfoDedicated-v1530 ::= CHOICE
const (
	AntennainfodedicatedV1530ChoiceNothing = iota
	AntennainfodedicatedV1530ChoiceRelease
	AntennainfodedicatedV1530ChoiceSetup
)

type AntennainfodedicatedV1530 struct {
	Choice  uint64
	Release *struct{}
	Setup   *AntennainfodedicatedV1530Setup
}
