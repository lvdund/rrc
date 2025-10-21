package ies

import "rrc/utils"

// MeasResultForECID-r9 ::= SEQUENCE
type MeasresultforecidR9 struct {
	UeRxtxtimediffresultR9 utils.INTEGER
	CurrentsfnR9           utils.BITSTRING
}
