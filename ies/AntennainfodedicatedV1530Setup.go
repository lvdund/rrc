package ies

import "rrc/utils"

// AntennaInfoDedicated-v1530-setup ::= CHOICE
const (
	AntennainfodedicatedV1530SetupChoiceNothing = iota
	AntennainfodedicatedV1530SetupChoiceUeTxantennaselectionSrs1t4rConfigR15
	AntennainfodedicatedV1530SetupChoiceUeTxantennaselectionSrs2t4rNrofpairsR15
)

type AntennainfodedicatedV1530Setup struct {
	Choice                                  uint64
	UeTxantennaselectionSrs1t4rConfigR15    *struct{}
	UeTxantennaselectionSrs2t4rNrofpairsR15 *utils.ENUMERATED
}
