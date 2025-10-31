package ies

// SL-RLC-ModeIndication-r16 ::= SEQUENCE
type SlRlcModeindicationR16 struct {
	SlModeR16        SlRlcModeindicationR16SlModeR16
	SlQosInfolistR16 []SlQosInfoR16 `lb:1,ub:maxNrofSLQfisperdestR16`
}
