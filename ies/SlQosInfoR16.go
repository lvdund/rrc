package ies

// SL-QoS-Info-r16 ::= SEQUENCE
type SlQosInfoR16 struct {
	SlQosFlowidentityR16 SlQosFlowidentityR16
	SlQosProfileR16      *SlQosProfileR16
}
