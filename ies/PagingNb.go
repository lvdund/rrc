package ies

import "rrc/utils"

// Paging-NB ::= SEQUENCE
type PagingNb struct {
	PagingrecordlistR13           *PagingrecordlistNbR13
	SysteminfomodificationR13     *utils.ENUMERATED
	SysteminfomodificationEdrxR13 *utils.ENUMERATED
	Noncriticalextension          *PagingNbV1610Ies
}
