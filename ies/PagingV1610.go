package ies

// Paging-v1610-IEs ::= SEQUENCE
type PagingV1610 struct {
	PagingrecordlistV1610   *PagingrecordlistV1610
	UacParammodificationR16 *bool
	Noncriticalextension    *PagingV1610IesNoncriticalextension
}
