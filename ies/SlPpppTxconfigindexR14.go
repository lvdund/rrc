package ies

import "rrc/utils"

// SL-PPPP-TxConfigIndex-r14 ::= SEQUENCE
type SlPpppTxconfigindexR14 struct {
	PrioritythresholdR14    SlPriorityR13
	DefaulttxconfigindexR14 utils.INTEGER      `lb:0,ub:maxCBRLevel1R14`
	CbrConfigindexR14       utils.INTEGER      `lb:0,ub:maxSLV2xCbrconfig1R14`
	TxConfigindexlistR14    []TxConfigindexR14 `lb:1,ub:maxCBRLevelR14`
}
