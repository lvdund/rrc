package ies

// PagingRecordList-NB-r13 ::= SEQUENCE OF PagingRecord-NB-r13
// SIZE (1..maxPageRec)
type PagingrecordlistNbR13 struct {
	Value []PagingrecordNbR13 `lb:1,ub:maxPageRec`
}
