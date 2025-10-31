package ies

// PerRAInfoList-r16 ::= SEQUENCE OF PerRAInfo-r16
// SIZE (1..200)
type PerrainfolistR16 struct {
	Value []PerrainfoR16 `lb:1,ub:200`
}
