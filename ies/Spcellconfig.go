package ies

import "rrc/utils"

// SpCellConfig ::= SEQUENCE
// Extensible
type Spcellconfig struct {
	Servcellindex                     *Servcellindex
	Reconfigurationwithsync           *Reconfigurationwithsync
	RlfTimersandconstants             *utils.Setuprelease[RlfTimersandconstants]
	Rlminsyncoutofsyncthreshold       *SpcellconfigRlminsyncoutofsyncthreshold
	Spcellconfigdedicated             *Servingcellconfig
	LowmobilityevaluationconnectedR17 *SpcellconfigLowmobilityevaluationconnectedR17 `ext`
	GoodservingcellevaluationrlmR17   *GoodservingcellevaluationR17                  `ext`
	GoodservingcellevaluationbfdR17   *GoodservingcellevaluationR17                  `ext`
	DeactivatedscgConfigR17           *utils.Setuprelease[DeactivatedscgConfigR17]   `ext`
}
