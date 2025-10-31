package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-r13-ackNackRepetition-r13-setup ::= SEQUENCE
type PucchConfigdedicatedR13AcknackrepetitionR13Setup struct {
	RepetitionfactorR13 PucchConfigdedicatedR13AcknackrepetitionR13SetupRepetitionfactorR13
	N1pucchAnRepR13     utils.INTEGER `lb:0,ub:2047`
}
