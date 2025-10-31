package ies

// Paging-v1310-IEs ::= SEQUENCE
type PagingV1310 struct {
	RedistributionindicationR13   *bool
	SysteminfomodificationEdrxR13 *bool
	Noncriticalextension          *PagingV1530
}
