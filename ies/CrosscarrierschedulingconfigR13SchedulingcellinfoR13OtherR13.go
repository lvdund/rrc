package ies

import "rrc/utils"

// CrossCarrierSchedulingConfig-r13-schedulingCellInfo-r13-other-r13 ::= SEQUENCE
type CrosscarrierschedulingconfigR13SchedulingcellinfoR13OtherR13 struct {
	SchedulingcellidR13    ServcellindexR13
	PdschStartR13          utils.INTEGER `lb:0,ub:4`
	CifInschedulingcellR13 utils.INTEGER `lb:0,ub:7`
}
