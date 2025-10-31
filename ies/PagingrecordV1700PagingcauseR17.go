package ies

import "rrc/utils"

// PagingRecord-v1700-pagingCause-r17 ::= ENUMERATED
type PagingrecordV1700PagingcauseR17 struct {
	Value utils.ENUMERATED
}

const (
	PagingrecordV1700PagingcauseR17EnumeratedNothing = iota
	PagingrecordV1700PagingcauseR17EnumeratedVoice
)
