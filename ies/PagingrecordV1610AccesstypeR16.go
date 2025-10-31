package ies

import "rrc/utils"

// PagingRecord-v1610-accessType-r16 ::= ENUMERATED
type PagingrecordV1610AccesstypeR16 struct {
	Value utils.ENUMERATED
}

const (
	PagingrecordV1610AccesstypeR16EnumeratedNothing = iota
	PagingrecordV1610AccesstypeR16EnumeratedNon3gpp
)
