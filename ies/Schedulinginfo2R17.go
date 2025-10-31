package ies

import "rrc/utils"

// SchedulingInfo2-r17 ::= SEQUENCE
type Schedulinginfo2R17 struct {
	SiBroadcaststatusR17 Schedulinginfo2R17SiBroadcaststatusR17
	SiWindowpositionR17  utils.INTEGER `lb:0,ub:256`
	SiPeriodicityR17     Schedulinginfo2R17SiPeriodicityR17
	SibMappinginfoR17    SibMappingV1700
}
