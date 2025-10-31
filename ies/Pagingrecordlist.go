package ies

// PagingRecordList ::= SEQUENCE OF PagingRecord
// SIZE (1..maxNrofPageRec)
type Pagingrecordlist struct {
	Value []Pagingrecord `lb:1,ub:maxNrofPageRec`
}
