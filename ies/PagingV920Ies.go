package ies

import "rrc/utils"

// Paging-v920-IEs ::= SEQUENCE
type PagingV920Ies struct {
	CmasIndicationR9     *utils.ENUMERATED
	Noncriticalextension *PagingV1130Ies
}
