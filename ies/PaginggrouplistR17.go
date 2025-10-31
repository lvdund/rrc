package ies

// PagingGroupList-r17 ::= SEQUENCE OF TMGI-r17
// SIZE (1..maxNrofPageGroup-r17)
type PaginggrouplistR17 struct {
	Value []TmgiR17 `lb:1,ub:maxNrofPageGroupR17`
}
