package ies

import "rrc/utils"

// SL-CBR-PSSCH-TxConfig-r14 ::= SEQUENCE
type SlCbrPsschTxconfigR14 struct {
	CrLimitR14      utils.INTEGER `lb:0,ub:10000`
	TxParametersR14 SlPsschTxparametersR14
}
