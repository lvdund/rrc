package ies

// SL-CBR-CommonTxConfigList-r14 ::= SEQUENCE
type SlCbrCommontxconfiglistR14 struct {
	CbrRangecommonconfiglistR14 []SlCbrLevelsConfigR14  `lb:1,ub:maxSLV2xCbrconfigR14`
	SlCbrPsschTxconfiglistR14   []SlCbrPsschTxconfigR14 `lb:1,ub:maxSLV2xTxconfigR14`
}
