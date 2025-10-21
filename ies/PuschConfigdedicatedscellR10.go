package ies

import "rrc/utils"

// PUSCH-ConfigDedicatedSCell-r10 ::= SEQUENCE
type PuschConfigdedicatedscellR10 struct {
	GrouphoppingdisabledR10 *utils.ENUMERATED
	DmrsWithoccActivatedR10 *utils.ENUMERATED
}
