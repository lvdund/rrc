package ies

import "rrc/utils"

// RMTC-Config-r13-setup-rmtc-Period-r13 ::= ENUMERATED
type RmtcConfigR13SetupRmtcPeriodR13 struct {
	Value utils.ENUMERATED
}

const (
	RmtcConfigR13SetupRmtcPeriodR13EnumeratedNothing = iota
	RmtcConfigR13SetupRmtcPeriodR13EnumeratedMs40
	RmtcConfigR13SetupRmtcPeriodR13EnumeratedMs80
	RmtcConfigR13SetupRmtcPeriodR13EnumeratedMs160
	RmtcConfigR13SetupRmtcPeriodR13EnumeratedMs320
	RmtcConfigR13SetupRmtcPeriodR13EnumeratedMs640
)
