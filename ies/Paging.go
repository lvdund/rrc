package ies

import "rrc/utils"

// Paging ::= SEQUENCE
type Paging struct {
	Pagingrecordlist         *Pagingrecordlist
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *PagingV1700
}
