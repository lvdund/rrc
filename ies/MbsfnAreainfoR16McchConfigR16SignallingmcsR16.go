package ies

import "rrc/utils"

// MBSFN-AreaInfo-r16-mcch-Config-r16-signallingMCS-r16 ::= ENUMERATED
type MbsfnAreainfoR16McchConfigR16SignallingmcsR16 struct {
	Value utils.ENUMERATED
}

const (
	MbsfnAreainfoR16McchConfigR16SignallingmcsR16EnumeratedNothing = iota
	MbsfnAreainfoR16McchConfigR16SignallingmcsR16EnumeratedN2
	MbsfnAreainfoR16McchConfigR16SignallingmcsR16EnumeratedN7
	MbsfnAreainfoR16McchConfigR16SignallingmcsR16EnumeratedN13
	MbsfnAreainfoR16McchConfigR16SignallingmcsR16EnumeratedN19
)
