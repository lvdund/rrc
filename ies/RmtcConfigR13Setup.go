package ies

import "rrc/utils"

// RMTC-Config-r13-setup ::= SEQUENCE
// Extensible
type RmtcConfigR13Setup struct {
	RmtcPeriodR13         RmtcConfigR13SetupRmtcPeriodR13
	RmtcSubframeoffsetR13 *utils.INTEGER `lb:0,ub:639`
	MeasdurationR13       RmtcConfigR13SetupMeasdurationR13
}
