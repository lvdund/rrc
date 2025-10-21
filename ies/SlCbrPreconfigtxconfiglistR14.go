package ies

import "rrc/utils"

// SL-CBR-PreconfigTxConfigList-r14 ::= SEQUENCE
type SlCbrPreconfigtxconfiglistR14 struct {
	CbrRangecommonconfiglistR14 interface{}
	SlCbrPsschTxconfiglistR14   interface{}
}
