package ies

import "rrc/utils"

// CrossCarrierSchedulingConfigLAA-UL-r14 ::= SEQUENCE
type CrosscarrierschedulingconfiglaaUlR14 struct {
	SchedulingcellidR14    ServcellindexR13
	CifInschedulingcellR14 utils.INTEGER
}
