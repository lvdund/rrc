package ies

// SL-CBR-PreconfigTxConfigList-r14 ::= SEQUENCE
type SlCbrPreconfigtxconfiglistR14 struct {
	CbrRangecommonconfiglistR14 []SlCbrLevelsConfigR14  `lb:1,ub:maxSLV2xCbrconfig2R14`
	SlCbrPsschTxconfiglistR14   []SlCbrPsschTxconfigR14 `lb:1,ub:maxSLV2xTxconfig2R14`
}
