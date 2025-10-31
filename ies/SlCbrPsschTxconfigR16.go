package ies

import "rrc/utils"

// SL-CBR-PSSCH-TxConfig-r16 ::= SEQUENCE
type SlCbrPsschTxconfigR16 struct {
	SlCrLimitR16      *utils.INTEGER `lb:0,ub:10000`
	SlTxparametersR16 *SlPsschTxparametersR16
}
