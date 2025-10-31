package ies

import "rrc/utils"

// PagingRecord-accessType ::= ENUMERATED
type PagingrecordAccesstype struct {
	Value utils.ENUMERATED
}

const (
	PagingrecordAccesstypeEnumeratedNothing = iota
	PagingrecordAccesstypeEnumeratedNon3gpp
)
