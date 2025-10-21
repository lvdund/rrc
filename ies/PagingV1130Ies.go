package ies

import "rrc/utils"

// Paging-v1130-IEs ::= SEQUENCE
type PagingV1130Ies struct {
	EabParammodificationR11 *utils.ENUMERATED
	Noncriticalextension    *PagingV1310Ies
}
