package ies

// PLMN-InfoList-r16 ::= SEQUENCE OF PLMN-Info-r16
// SIZE (0..maxPLMN-r11)
type PlmnInfolistR16 struct {
	Value []PlmnInfoR16 `lb:0,ub:maxPLMNR11`
}
