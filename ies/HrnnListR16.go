package ies

// HRNN-List-r16 ::= SEQUENCE OF HRNN-r16
// SIZE (1..maxNPN-r16)
type HrnnListR16 struct {
	Value []HrnnR16 `lb:1,ub:maxNPNR16`
}
