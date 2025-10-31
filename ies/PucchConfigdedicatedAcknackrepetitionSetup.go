package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-ackNackRepetition-setup ::= SEQUENCE
type PucchConfigdedicatedAcknackrepetitionSetup struct {
	Repetitionfactor PucchConfigdedicatedAcknackrepetitionSetupRepetitionfactor
	N1pucchAnRep     utils.INTEGER `lb:0,ub:2047`
}
