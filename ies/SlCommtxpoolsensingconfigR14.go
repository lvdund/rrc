package ies

// SL-CommTxPoolSensingConfig-r14 ::= SEQUENCE
type SlCommtxpoolsensingconfigR14 struct {
	PsschTxconfiglistR14                 SlPsschTxconfiglistR14
	ThrespsschRsrpListR14                SlThrespsschRsrpListR14
	RestrictresourcereservationperiodR14 *SlRestrictresourcereservationperiodlistR14
	ProbresourcekeepR14                  SlCommtxpoolsensingconfigR14ProbresourcekeepR14
	P2xSensingconfigR14                  *SlCommtxpoolsensingconfigR14P2xSensingconfigR14
	SlReselectafterR14                   *SlCommtxpoolsensingconfigR14SlReselectafterR14
}
