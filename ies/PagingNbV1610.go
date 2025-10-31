package ies

// Paging-NB-v1610-IEs ::= SEQUENCE
type PagingNbV1610 struct {
	PagingrecordlistV1610 *PagingrecordlistNbV1610
	Noncriticalextension  *PagingNbV1610IesNoncriticalextension
}
