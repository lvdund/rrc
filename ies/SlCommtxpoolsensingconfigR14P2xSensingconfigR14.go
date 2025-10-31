package ies

import "rrc/utils"

// SL-CommTxPoolSensingConfig-r14-p2x-SensingConfig-r14 ::= SEQUENCE
type SlCommtxpoolsensingconfigR14P2xSensingconfigR14 struct {
	MinnumcandidatesfR14   utils.INTEGER   `lb:0,ub:13`
	GapcandidatesensingR14 utils.BITSTRING `lb:10,ub:10`
}
