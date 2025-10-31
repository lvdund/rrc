package ies

import "rrc/utils"

// RACH-ConfigCommon ::= SEQUENCE
// Extensible
type RachConfigcommon struct {
	Preambleinfo           *RachConfigcommonPreambleinfo
	Powerrampingparameters Powerrampingparameters
	RaSupervisioninfo      RachConfigcommonRaSupervisioninfo
	MaxharqMsg3tx          utils.INTEGER `lb:0,ub:8`
}
