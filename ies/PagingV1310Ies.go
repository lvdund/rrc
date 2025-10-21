package ies

import "rrc/utils"

// Paging-v1310-IEs ::= SEQUENCE
type PagingV1310Ies struct {
	RedistributionindicationR13   *utils.ENUMERATED
	SysteminfomodificationEdrxR13 *utils.ENUMERATED
	Noncriticalextension          *PagingV1530Ies
}
