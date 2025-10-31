package ies

// Paging ::= SEQUENCE
type Paging struct {
	Pagingrecordlist       *Pagingrecordlist
	Systeminfomodification *bool
	EtwsIndication         *bool
	Noncriticalextension   *PagingV890
}
