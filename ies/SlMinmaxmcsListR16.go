package ies

// SL-MinMaxMCS-List-r16 ::= SEQUENCE OF SL-MinMaxMCS-Config-r16
// SIZE (1..3)
type SlMinmaxmcsListR16 struct {
	Value []SlMinmaxmcsConfigR16 `lb:1,ub:3`
}
