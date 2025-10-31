package ies

import "rrc/utils"

// AntennaInfoDedicatedSTTI-r15-setup-transmissionModeDL-MBSFN-r15 ::= ENUMERATED
type AntennainfodedicatedsttiR15SetupTransmissionmodedlMbsfnR15 struct {
	Value utils.ENUMERATED
}

const (
	AntennainfodedicatedsttiR15SetupTransmissionmodedlMbsfnR15EnumeratedNothing = iota
	AntennainfodedicatedsttiR15SetupTransmissionmodedlMbsfnR15EnumeratedTm9
	AntennainfodedicatedsttiR15SetupTransmissionmodedlMbsfnR15EnumeratedTm10
)
