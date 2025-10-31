package ies

import "rrc/utils"

// AntennaInfoDedicated-ue-TransmitAntennaSelection ::= CHOICE
const (
	AntennainfodedicatedUeTransmitantennaselectionChoiceNothing = iota
	AntennainfodedicatedUeTransmitantennaselectionChoiceRelease
	AntennainfodedicatedUeTransmitantennaselectionChoiceSetup
)

type AntennainfodedicatedUeTransmitantennaselection struct {
	Choice  uint64
	Release *struct{}
	Setup   *utils.ENUMERATED
}
