package ies

// SL-CBR-CommonTxConfigList-r16 ::= SEQUENCE
type SlCbrCommontxconfiglistR16 struct {
	SlCbrRangeconfiglistR16   *[]SlCbrLevelsconfigR16  `lb:1,ub:maxCBRConfigR16`
	SlCbrPsschTxconfiglistR16 *[]SlCbrPsschTxconfigR16 `lb:1,ub:maxTxConfigR16`
}
