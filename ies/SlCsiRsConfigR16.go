package ies

import "rrc/utils"

// SL-CSI-RS-Config-r16 ::= SEQUENCE
// Extensible
type SlCsiRsConfigR16 struct {
	SlCsiRsFreqallocationR16 *SlCsiRsConfigR16SlCsiRsFreqallocationR16
	SlCsiRsFirstsymbolR16    *utils.INTEGER `lb:0,ub:12`
}
