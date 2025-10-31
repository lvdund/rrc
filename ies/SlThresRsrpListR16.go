package ies

// SL-Thres-RSRP-List-r16 ::= SEQUENCE OF SL-Thres-RSRP-r16
// SIZE (64)
type SlThresRsrpListR16 struct {
	Value []SlThresRsrpR16 `lb:64,ub:64`
}
