package ies

import "rrc/utils"

// AntennaInfoDedicated-r10 ::= SEQUENCE
type AntennainfodedicatedR10 struct {
	TransmissionmodeR10          utils.ENUMERATED
	CodebooksubsetrestrictionR10 *utils.BITSTRING
	UeTransmitantennaselection   utils.ENUMERATED
}
