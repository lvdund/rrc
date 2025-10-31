package ies

import "rrc/utils"

// MeasDS-Config-r12-setup-dmtc-PeriodOffset-r12 ::= CHOICE
// Extensible
const (
	MeasdsConfigR12SetupDmtcPeriodoffsetR12ChoiceNothing = iota
	MeasdsConfigR12SetupDmtcPeriodoffsetR12ChoiceMs40R12
	MeasdsConfigR12SetupDmtcPeriodoffsetR12ChoiceMs80R12
	MeasdsConfigR12SetupDmtcPeriodoffsetR12ChoiceMs160R12
)

type MeasdsConfigR12SetupDmtcPeriodoffsetR12 struct {
	Choice   uint64
	Ms40R12  *utils.INTEGER `lb:0,ub:39`
	Ms80R12  *utils.INTEGER `lb:0,ub:79`
	Ms160R12 *utils.INTEGER `lb:0,ub:159`
}
