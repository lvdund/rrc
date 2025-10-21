package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-v1020 ::= SEQUENCE
type PuschConfigdedicatedV1020 struct {
	BetaoffsetmcR10         *interface{}
	GrouphoppingdisabledR10 *utils.ENUMERATED
	DmrsWithoccActivatedR10 *utils.ENUMERATED
}
