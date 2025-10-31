package ies

import "rrc/utils"

// AntennaInfoUL-STTI-r15-transmissionModeUL-STTI-r15 ::= ENUMERATED
type AntennainfoulSttiR15TransmissionmodeulSttiR15 struct {
	Value utils.ENUMERATED
}

const (
	AntennainfoulSttiR15TransmissionmodeulSttiR15EnumeratedNothing = iota
	AntennainfoulSttiR15TransmissionmodeulSttiR15EnumeratedTm1
	AntennainfoulSttiR15TransmissionmodeulSttiR15EnumeratedTm2
)
