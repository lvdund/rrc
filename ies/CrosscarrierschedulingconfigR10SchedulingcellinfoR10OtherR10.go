package ies

import "rrc/utils"

// CrossCarrierSchedulingConfig-r10-schedulingCellInfo-r10-other-r10 ::= SEQUENCE
type CrosscarrierschedulingconfigR10SchedulingcellinfoR10OtherR10 struct {
	SchedulingcellidR10 ServcellindexR10
	PdschStartR10       utils.INTEGER `lb:0,ub:4`
}
