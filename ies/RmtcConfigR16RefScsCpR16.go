package ies

import "rrc/utils"

// RMTC-Config-r16-ref-SCS-CP-r16 ::= ENUMERATED
type RmtcConfigR16RefScsCpR16 struct {
	Value utils.ENUMERATED
}

const (
	RmtcConfigR16RefScsCpR16EnumeratedNothing = iota
	RmtcConfigR16RefScsCpR16EnumeratedKhz15
	RmtcConfigR16RefScsCpR16EnumeratedKhz30
	RmtcConfigR16RefScsCpR16EnumeratedKhz60_Ncp
	RmtcConfigR16RefScsCpR16EnumeratedKhz60_Ecp
)
