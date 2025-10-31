package ies

// PhysCellIdListCDMA2000-v920 ::= SEQUENCE OF PhysCellIdCDMA2000
// SIZE (0..24)
type Physcellidlistcdma2000V920 struct {
	Value []Physcellidcdma2000 `lb:0,ub:24`
}
