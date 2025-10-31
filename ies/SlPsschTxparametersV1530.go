package ies

import "rrc/utils"

// SL-PSSCH-TxParameters-v1530 ::= SEQUENCE
type SlPsschTxparametersV1530 struct {
	MinmcsPsschR15 utils.INTEGER `lb:0,ub:31`
	MaxmcsPsschR15 utils.INTEGER `lb:0,ub:31`
}
