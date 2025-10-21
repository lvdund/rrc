package ies

import "rrc/utils"

// SCPTM-Parameters-r13 ::= SEQUENCE
type ScptmParametersR13 struct {
	ScptmParallelreceptionR13 *utils.ENUMERATED
	ScptmScellR13             *utils.ENUMERATED
	ScptmNonservingcellR13    *utils.ENUMERATED
	ScptmAsyncdcR13           *utils.ENUMERATED
}
