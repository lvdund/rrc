package ies

import "rrc/utils"

// MeasPosPreConfigGapId-r17 ::= utils.INTEGER (1..maxNrofPreConfigPosGapId-r17)
type MeaspospreconfiggapidR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPreConfigPosGapIdR17`
}
