package ies

import "rrc/utils"

// CA-ParametersNRDC-v1650 ::= SEQUENCE
type CaParametersnrdcV1650 struct {
	SupportedcellgroupingR16 *utils.BITSTRING `lb:1,ub:maxCellGroupingsR16`
}
