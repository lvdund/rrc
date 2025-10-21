package ies

import "rrc/utils"

// RACH-ConfigCommon ::= SEQUENCE
// Extensible
type RachConfigcommon struct {
	Preambleinfo           *interface{}
	Powerrampingparameters Powerrampingparameters
	RaSupervisioninfo      interface{}
	MaxharqMsg3tx          utils.INTEGER
}
