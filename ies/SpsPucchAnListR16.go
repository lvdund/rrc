package ies

// SPS-PUCCH-AN-List-r16 ::= SEQUENCE OF SPS-PUCCH-AN-r16
// SIZE (1..4)
type SpsPucchAnListR16 struct {
	Value []SpsPucchAnR16 `lb:1,ub:4`
}
