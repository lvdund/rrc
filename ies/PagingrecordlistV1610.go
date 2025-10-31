package ies

// PagingRecordList-v1610 ::= SEQUENCE OF PagingRecord-v1610
// SIZE (1..maxPageRec)
type PagingrecordlistV1610 struct {
	Value []PagingrecordV1610 `lb:1,ub:maxPageRec`
}
