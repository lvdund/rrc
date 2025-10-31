package ies

// Paging-v1700-IEs ::= SEQUENCE
type PagingV1700 struct {
	PagingrecordlistV1700 *PagingrecordlistV1700
	PaginggrouplistR17    *PaginggrouplistR17
	Noncriticalextension  *PagingV1700IesNoncriticalextension
}
