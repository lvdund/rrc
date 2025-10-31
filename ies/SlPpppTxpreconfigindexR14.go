package ies

import "rrc/utils"

// SL-PPPP-TxPreconfigIndex-r14 ::= SEQUENCE
type SlPpppTxpreconfigindexR14 struct {
	PrioritythresholdR14    SlPriorityR13
	DefaulttxconfigindexR14 utils.INTEGER         `lb:0,ub:maxCBRLevel1R14`
	CbrConfigindexR14       utils.INTEGER         `lb:0,ub:maxSLV2xCbrconfig21R14`
	TxConfigindexlistR14    []TxPreconfigindexR14 `lb:1,ub:maxCBRLevelR14`
}
