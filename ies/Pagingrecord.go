package ies

// PagingRecord ::= SEQUENCE
// Extensible
type Pagingrecord struct {
	UeIdentity PagingueIdentity
	CnDomain   PagingrecordCnDomain
}
