package ies

// PagingRecordList-NB-v1610 ::= SEQUENCE OF PagingRecord-NB-v1610
// SIZE (1..maxPageRec)
type PagingrecordlistNbV1610 struct {
	Value []PagingrecordNbV1610 `lb:1,ub:maxPageRec`
}
