package ies

import "rrc/utils"

// Paging-v1610-IEs ::= SEQUENCE
type PagingV1610Ies struct {
	PagingrecordlistV1610   *PagingrecordlistV1610
	UacParammodificationR16 *utils.ENUMERATED
	Noncriticalextension    *interface{}
}
