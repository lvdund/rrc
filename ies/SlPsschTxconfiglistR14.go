package ies

// SL-PSSCH-TxConfigList-r14 ::= SEQUENCE OF SL-PSSCH-TxConfig-r14
// SIZE (1..maxPSSCH-TxConfig-r14)
type SlPsschTxconfiglistR14 struct {
	Value []SlPsschTxconfigR14 `lb:1,ub:maxPSSCHTxconfigR14`
}
