package ies

import "rrc/utils"

// SL-PSSCH-TxParameters-r14 ::= SEQUENCE
type SlPsschTxparametersR14 struct {
	MinmcsPsschR14              utils.INTEGER `lb:0,ub:31`
	MaxmcsPsschR14              utils.INTEGER `lb:0,ub:31`
	MinsubchannelNumberpsschR14 utils.INTEGER `lb:0,ub:20`
	MaxsubchannelNumberpsschR14 utils.INTEGER `lb:0,ub:20`
	AllowedretxnumberpsschR14   SlPsschTxparametersR14AllowedretxnumberpsschR14
	MaxtxpowerR14               *SlTxpowerR14
}
