package ies

import "rrc/utils"

// SL-CommTxPoolSensingConfig-r14 ::= SEQUENCE
type SlCommtxpoolsensingconfigR14 struct {
	PsschTxconfiglistR14                 SlPsschTxconfiglistR14
	ThrespsschRsrpListR14                SlThrespsschRsrpListR14
	RestrictresourcereservationperiodR14 *SlRestrictresourcereservationperiodlistR14
	ProbresourcekeepR14                  utils.ENUMERATED
	P2xSensingconfigR14                  *interface{}
	SlReselectafterR14                   *utils.ENUMERATED
}
