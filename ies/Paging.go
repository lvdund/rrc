package ies

import "rrc/utils"

// Paging ::= SEQUENCE
type Paging struct {
	Pagingrecordlist       *Pagingrecordlist
	Systeminfomodification *utils.ENUMERATED
	EtwsIndication         *utils.ENUMERATED
	Noncriticalextension   *PagingV890Ies
}
