package ies

import "rrc/utils"

// RMTC-Config-r16-rmtc-Periodicity-r16 ::= ENUMERATED
type RmtcConfigR16RmtcPeriodicityR16 struct {
	Value utils.ENUMERATED
}

const (
	RmtcConfigR16RmtcPeriodicityR16EnumeratedNothing = iota
	RmtcConfigR16RmtcPeriodicityR16EnumeratedMs40
	RmtcConfigR16RmtcPeriodicityR16EnumeratedMs80
	RmtcConfigR16RmtcPeriodicityR16EnumeratedMs160
	RmtcConfigR16RmtcPeriodicityR16EnumeratedMs320
	RmtcConfigR16RmtcPeriodicityR16EnumeratedMs640
)
