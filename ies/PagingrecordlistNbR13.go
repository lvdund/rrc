package ies

import "rrc/utils"

// PagingRecordList-NB-r13 ::= SEQUENCE OF PagingRecord-NB-r13
// SIZE (1..maxPageRec)
type PagingrecordlistNbR13 struct {
	Value utils.Sequence[PagingrecordNbR13]
}
