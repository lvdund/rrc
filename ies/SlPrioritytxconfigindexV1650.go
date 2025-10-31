package ies

// SL-PriorityTxConfigIndex-v1650 ::= SEQUENCE
type SlPrioritytxconfigindexV1650 struct {
	SlMcsRangelistR16 *[]SlMinmaxmcsListR16 `lb:1,ub:maxCBRLevelR16`
}
