package ies

// SL-RxInterestedGC-BC-Dest-r17 ::= SEQUENCE
type SlRxinterestedgcBcDestR17 struct {
	SlRxinterestedqosInfolistR17 []SlQosInfoR16 `lb:1,ub:maxNrofSLQfisperdestR16`
	SlDestinationidentityR16     SlDestinationidentityR16
}
