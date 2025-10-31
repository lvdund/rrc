package ies

// PagingRecord ::= SEQUENCE
// Extensible
type Pagingrecord struct {
	UeIdentity PagingueIdentity
	Accesstype *PagingrecordAccesstype
}
