package ies

import "rrc/utils"

// CondReconfigurationTriggerEUTRA-r16-condEventId-r16-condEventA3-r16 ::= SEQUENCE
type CondreconfigurationtriggereutraR16CondeventidR16Condeventa3R16 struct {
	A3OffsetR16      utils.INTEGER `lb:0,ub:30`
	HysteresisR16    Hysteresis
	TimetotriggerR16 Timetotrigger
}
