package ies

import "rrc/utils"

// CrossCarrierSchedulingConfig-schedulingCellInfo-other ::= SEQUENCE
type CrosscarrierschedulingconfigSchedulingcellinfoOther struct {
	Schedulingcellid    Servcellindex
	CifInschedulingcell utils.INTEGER `lb:0,ub:7`
}
