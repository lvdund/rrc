package ies

import "rrc/utils"

// AntennaInfoDedicated-r10 ::= SEQUENCE
type AntennainfodedicatedR10 struct {
	TransmissionmodeR10          AntennainfodedicatedR10TransmissionmodeR10
	CodebooksubsetrestrictionR10 *utils.BITSTRING
	UeTransmitantennaselection   AntennainfodedicatedR10UeTransmitantennaselection
}
