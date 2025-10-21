package ies

import "rrc/utils"

// Paging-NB-v1610-IEs ::= SEQUENCE
type PagingNbV1610Ies struct {
	PagingrecordlistV1610 *PagingrecordlistNbV1610
	Noncriticalextension  *interface{}
}
