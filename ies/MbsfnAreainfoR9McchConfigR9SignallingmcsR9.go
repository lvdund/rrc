package ies

import "rrc/utils"

// MBSFN-AreaInfo-r9-mcch-Config-r9-signallingMCS-r9 ::= ENUMERATED
type MbsfnAreainfoR9McchConfigR9SignallingmcsR9 struct {
	Value utils.ENUMERATED
}

const (
	MbsfnAreainfoR9McchConfigR9SignallingmcsR9EnumeratedNothing = iota
	MbsfnAreainfoR9McchConfigR9SignallingmcsR9EnumeratedN2
	MbsfnAreainfoR9McchConfigR9SignallingmcsR9EnumeratedN7
	MbsfnAreainfoR9McchConfigR9SignallingmcsR9EnumeratedN13
	MbsfnAreainfoR9McchConfigR9SignallingmcsR9EnumeratedN19
)
