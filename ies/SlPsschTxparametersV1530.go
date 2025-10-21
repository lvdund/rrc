package ies

import "rrc/utils"

// SL-PSSCH-TxParameters-v1530 ::= SEQUENCE
type SlPsschTxparametersV1530 struct {
	MinmcsPsschR15 utils.INTEGER
	MaxmcsPsschR15 utils.INTEGER
}
