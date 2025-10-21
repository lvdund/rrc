package ies

import "rrc/utils"

// Paging-v1530-IEs ::= SEQUENCE
type PagingV1530Ies struct {
	Accesstype           *utils.ENUMERATED
	Noncriticalextension *PagingV1610Ies
}
