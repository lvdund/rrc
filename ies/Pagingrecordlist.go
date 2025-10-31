package ies

// PagingRecordList ::= SEQUENCE OF PagingRecord
// SIZE (1..maxPageRec)
type Pagingrecordlist struct {
	Value []Pagingrecord `lb:1,ub:maxPageRec`
}
