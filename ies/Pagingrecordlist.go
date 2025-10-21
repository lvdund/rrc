package ies

import "rrc/utils"

// PagingRecordList ::= SEQUENCE OF PagingRecord
// SIZE (1..maxPageRec)
type Pagingrecordlist struct {
	Value utils.Sequence[Pagingrecord]
}
