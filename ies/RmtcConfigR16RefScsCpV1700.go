package ies

import "rrc/utils"

// RMTC-Config-r16-ref-SCS-CP-v1700 ::= ENUMERATED
type RmtcConfigR16RefScsCpV1700 struct {
	Value utils.ENUMERATED
}

const (
	RmtcConfigR16RefScsCpV1700EnumeratedNothing = iota
	RmtcConfigR16RefScsCpV1700EnumeratedKhz120
	RmtcConfigR16RefScsCpV1700EnumeratedKhz480
	RmtcConfigR16RefScsCpV1700EnumeratedKhz960
)
