package ies

// SL-CBR-PriorityTxConfigList-v1650 ::= SEQUENCE OF SL-PriorityTxConfigIndex-v1650
// SIZE (1..8)
type SlCbrPrioritytxconfiglistV1650 struct {
	Value []SlPrioritytxconfigindexV1650 `lb:1,ub:8`
}
