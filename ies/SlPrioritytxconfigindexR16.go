package ies

import "rrc/utils"

// SL-PriorityTxConfigIndex-r16 ::= SEQUENCE
type SlPrioritytxconfigindexR16 struct {
	SlPrioritythresholdR16    *utils.INTEGER        `lb:0,ub:8`
	SlDefaulttxconfigindexR16 *utils.INTEGER        `lb:0,ub:maxCBRLevel1R16`
	SlCbrConfigindexR16       *utils.INTEGER        `lb:0,ub:maxCBRConfig1R16`
	SlTxConfigindexlistR16    *[]SlTxconfigindexR16 `lb:1,ub:maxCBRLevelR16`
}
