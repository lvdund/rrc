package ies

import "rrc/utils"

// Paging-v1530-IEs-accessType ::= ENUMERATED
type PagingV1530IesAccesstype struct {
	Value utils.ENUMERATED
}

const (
	PagingV1530IesAccesstypeEnumeratedNothing = iota
	PagingV1530IesAccesstypeEnumeratedNon3gpp
)
