package ies

import "rrc/utils"

// SL-PSSCH-TxParameters-r14 ::= SEQUENCE
type SlPsschTxparametersR14 struct {
	MinmcsPsschR14              utils.INTEGER
	MaxmcsPsschR14              utils.INTEGER
	MinsubchannelNumberpsschR14 utils.INTEGER
	MaxsubchannelNumberpsschR14 utils.INTEGER
	AllowedretxnumberpsschR14   utils.ENUMERATED
	MaxtxpowerR14               *SlTxpowerR14
}
