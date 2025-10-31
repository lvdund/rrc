package ies

import "rrc/utils"

// RMTC-Config-r13-setup-measDuration-r13 ::= ENUMERATED
type RmtcConfigR13SetupMeasdurationR13 struct {
	Value utils.ENUMERATED
}

const (
	RmtcConfigR13SetupMeasdurationR13EnumeratedNothing = iota
	RmtcConfigR13SetupMeasdurationR13EnumeratedSym1
	RmtcConfigR13SetupMeasdurationR13EnumeratedSym14
	RmtcConfigR13SetupMeasdurationR13EnumeratedSym28
	RmtcConfigR13SetupMeasdurationR13EnumeratedSym42
	RmtcConfigR13SetupMeasdurationR13EnumeratedSym70
)
