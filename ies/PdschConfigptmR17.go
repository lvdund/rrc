package ies

import "rrc/utils"

// PDSCH-ConfigPTM-r17 ::= SEQUENCE
type PdschConfigptmR17 struct {
	DatascramblingidentitypdschR17 *utils.INTEGER `lb:0,ub:1023`
	DmrsScramblingid0R17           *utils.INTEGER `lb:0,ub:65535`
	PdschAggregationfactorR17      *PdschConfigptmR17PdschAggregationfactorR17
}
