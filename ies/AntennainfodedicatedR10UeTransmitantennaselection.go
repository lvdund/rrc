package ies

import "rrc/utils"

// AntennaInfoDedicated-r10-ue-TransmitAntennaSelection ::= CHOICE
const (
	AntennainfodedicatedR10UeTransmitantennaselectionChoiceNothing = iota
	AntennainfodedicatedR10UeTransmitantennaselectionChoiceRelease
	AntennainfodedicatedR10UeTransmitantennaselectionChoiceSetup
)

type AntennainfodedicatedR10UeTransmitantennaselection struct {
	Choice  uint64
	Release *struct{}
	Setup   *utils.ENUMERATED
}
