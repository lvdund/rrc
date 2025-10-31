package ies

// PhysCellIdListCDMA2000 ::= SEQUENCE OF PhysCellIdCDMA2000
// SIZE (1..16)
type Physcellidlistcdma2000 struct {
	Value []Physcellidcdma2000 `lb:1,ub:16`
}
