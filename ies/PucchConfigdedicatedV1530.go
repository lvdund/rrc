package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-v1530 ::= SEQUENCE
type PucchConfigdedicatedV1530 struct {
	N1pucchAnSptR15                  *utils.INTEGER
	CodebooksizedeterminationsttiR15 *utils.ENUMERATED
}
