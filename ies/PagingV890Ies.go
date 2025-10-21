package ies

import "rrc/utils"

// Paging-v890-IEs ::= SEQUENCE
type PagingV890Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *PagingV920Ies
}
