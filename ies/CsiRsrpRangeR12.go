package ies

import "rrc/utils"

// CSI-RSRP-Range-r12 ::= utils.INTEGER (0..97)
type CsiRsrpRangeR12 struct {
	Value utils.INTEGER `lb:0,ub:97`
}
