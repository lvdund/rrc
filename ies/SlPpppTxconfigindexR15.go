package ies

import "rrc/utils"

// SL-PPPP-TxConfigIndex-r15 ::= SEQUENCE
type SlPpppTxconfigindexR15 struct {
	PrioritythresholdR15    SlPriorityR13
	DefaulttxconfigindexR15 utils.INTEGER      `lb:0,ub:maxCBRLevel1R14`
	CbrConfigindexR15       utils.INTEGER      `lb:0,ub:maxSLV2xCbrconfig1R14`
	TxConfigindexlistR15    []TxConfigindexR14 `lb:1,ub:maxCBRLevelR14`
	McsPsschRangelistR15    []McsPsschRangeR15 `lb:1,ub:maxCBRLevelR14`
}
