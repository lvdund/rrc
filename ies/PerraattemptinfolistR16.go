package ies

// PerRAAttemptInfoList-r16 ::= SEQUENCE OF PerRAAttemptInfo-r16
// SIZE (1..200)
type PerraattemptinfolistR16 struct {
	Value []PerraattemptinfoR16 `lb:1,ub:200`
}
