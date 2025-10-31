package ies

// SL-CBR-PriorityTxConfigList-r16 ::= SEQUENCE OF SL-PriorityTxConfigIndex-r16
// SIZE (1..8)
type SlCbrPrioritytxconfiglistR16 struct {
	Value []SlPrioritytxconfigindexR16 `lb:1,ub:8`
}
