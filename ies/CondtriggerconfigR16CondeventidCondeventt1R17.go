package ies

import "rrc/utils"

// CondTriggerConfig-r16-condEventId-condEventT1-r17 ::= SEQUENCE
type CondtriggerconfigR16CondeventidCondeventt1R17 struct {
	T1ThresholdR17 utils.INTEGER `lb:0,ub:549755813887`
	DurationR17    utils.INTEGER `lb:0,ub:6000`
}
