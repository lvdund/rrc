package ies

import "rrc/utils"

// MeasResultForECID-r9 ::= SEQUENCE
type MeasresultforecidR9 struct {
	UeRxtxtimediffresultR9 utils.INTEGER   `lb:0,ub:4095`
	CurrentsfnR9           utils.BITSTRING `lb:10,ub:10`
}
