package ies

import "rrc/utils"

// PagingRecord-v1610 ::= SEQUENCE
type PagingrecordV1610 struct {
	AccesstypeR16 *utils.ENUMERATED
	MtEdtR16      *utils.ENUMERATED
}
