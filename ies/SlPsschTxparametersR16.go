package ies

import "rrc/utils"

// SL-PSSCH-TxParameters-r16 ::= SEQUENCE
type SlPsschTxparametersR16 struct {
	SlMinmcsPsschR16           utils.INTEGER `lb:0,ub:27`
	SlMaxmcsPsschR16           utils.INTEGER `lb:0,ub:31`
	SlMinsubchannelnumpsschR16 utils.INTEGER `lb:0,ub:27`
	SlMaxsubchannelnumpsschR16 utils.INTEGER `lb:0,ub:27`
	SlMaxtxtransnumpsschR16    utils.INTEGER `lb:0,ub:32`
	SlMaxtxpowerR16            *SlTxpowerR16
}
