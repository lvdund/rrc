package ies

import "rrc/utils"

// OffsetValue-r17 ::= SEQUENCE
type OffsetvalueR17 struct {
	OffsetvalueR17   utils.INTEGER `lb:0,ub:20000`
	Shift7dot5khzR17 utils.BOOLEAN
}
