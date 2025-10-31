package ies

// SL-PSSCH-TxConfigList-r16 ::= SEQUENCE OF SL-PSSCH-TxConfig-r16
// SIZE (1..maxPSSCH-TxConfig-r16)
type SlPsschTxconfiglistR16 struct {
	Value []SlPsschTxconfigR16 `lb:1,ub:maxPSSCHTxconfigR16`
}
