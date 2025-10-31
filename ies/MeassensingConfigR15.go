package ies

import "rrc/utils"

// MeasSensing-Config-r15 ::= SEQUENCE
type MeassensingConfigR15 struct {
	SensingsubchannelnumberR15   utils.INTEGER `lb:0,ub:20`
	SensingperiodicityR15        MeassensingConfigR15SensingperiodicityR15
	SensingreselectioncounterR15 utils.INTEGER `lb:0,ub:75`
	SensingpriorityR15           utils.INTEGER `lb:0,ub:8`
}
