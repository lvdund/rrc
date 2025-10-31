package ies

import "rrc/utils"

// CondTriggerConfig-r16-condEventId-condEventD1-r17 ::= SEQUENCE
type CondtriggerconfigR16CondeventidCondeventd1R17 struct {
	Distancethreshfromreference1R17 utils.INTEGER `lb:0,ub:65525`
	Distancethreshfromreference2R17 utils.INTEGER `lb:0,ub:65525`
	Referencelocation1R17           ReferencelocationR17
	Referencelocation2R17           ReferencelocationR17
	HysteresislocationR17           HysteresislocationR17
	TimetotriggerR17                Timetotrigger
}
