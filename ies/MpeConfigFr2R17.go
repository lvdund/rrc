package ies

import "rrc/utils"

// MPE-Config-FR2-r17 ::= SEQUENCE
// Extensible
type MpeConfigFr2R17 struct {
	MpeProhibittimerR17 MpeConfigFr2R17MpeProhibittimerR17
	MpeThresholdR17     MpeConfigFr2R17MpeThresholdR17
	NumberofnR17        utils.INTEGER `lb:0,ub:4`
}
