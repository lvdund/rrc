package ies

// PagingRecordList-v1700 ::= SEQUENCE OF PagingRecord-v1700
// SIZE (1..maxNrofPageRec)
type PagingrecordlistV1700 struct {
	Value []PagingrecordV1700 `lb:1,ub:maxNrofPageRec`
}
