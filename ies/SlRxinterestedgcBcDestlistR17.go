package ies

// SL-RxInterestedGC-BC-DestList-r17 ::= SEQUENCE OF SL-RxInterestedGC-BC-Dest-r17
// SIZE (1..maxNrofSL-Dest-r16)
type SlRxinterestedgcBcDestlistR17 struct {
	Value []SlRxinterestedgcBcDestR17 `lb:1,ub:maxNrofSLDestR16`
}
