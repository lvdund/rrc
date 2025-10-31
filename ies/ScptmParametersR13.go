package ies

// SCPTM-Parameters-r13 ::= SEQUENCE
type ScptmParametersR13 struct {
	ScptmParallelreceptionR13 *ScptmParametersR13ScptmParallelreceptionR13
	ScptmScellR13             *ScptmParametersR13ScptmScellR13
	ScptmNonservingcellR13    *ScptmParametersR13ScptmNonservingcellR13
	ScptmAsyncdcR13           *ScptmParametersR13ScptmAsyncdcR13
}
