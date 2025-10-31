package ies

// Paging-NB ::= SEQUENCE
type PagingNb struct {
	PagingrecordlistR13           *PagingrecordlistNbR13
	SysteminfomodificationR13     *bool
	SysteminfomodificationEdrxR13 *bool
	Noncriticalextension          *PagingNbV1610
}
