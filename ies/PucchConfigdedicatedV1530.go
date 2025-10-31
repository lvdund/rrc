package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-v1530 ::= SEQUENCE
type PucchConfigdedicatedV1530 struct {
	N1pucchAnSptR15                  *utils.INTEGER `lb:0,ub:2047`
	CodebooksizedeterminationsttiR15 *PucchConfigdedicatedV1530CodebooksizedeterminationsttiR15
}
