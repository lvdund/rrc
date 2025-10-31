package ies

// SL-ThresPSSCH-RSRP-List-r14 ::= SEQUENCE OF SL-ThresPSSCH-RSRP-r14
// SIZE (64)
type SlThrespsschRsrpListR14 struct {
	Value []SlThrespsschRsrpR14 `lb:64,ub:64`
}
