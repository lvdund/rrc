package ies

import "rrc/utils"

// MeasSensing-Config-r15 ::= SEQUENCE
type MeassensingConfigR15 struct {
	SensingsubchannelnumberR15   utils.INTEGER
	SensingperiodicityR15        utils.ENUMERATED
	SensingreselectioncounterR15 utils.INTEGER
	SensingpriorityR15           utils.INTEGER
}
