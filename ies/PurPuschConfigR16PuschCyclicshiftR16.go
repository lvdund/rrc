package ies

import "rrc/utils"

// PUR-PUSCH-Config-r16-pusch-CyclicShift-r16 ::= ENUMERATED
type PurPuschConfigR16PuschCyclicshiftR16 struct {
	Value utils.ENUMERATED
}

const (
	PurPuschConfigR16PuschCyclicshiftR16EnumeratedNothing = iota
	PurPuschConfigR16PuschCyclicshiftR16EnumeratedN0
	PurPuschConfigR16PuschCyclicshiftR16EnumeratedN6
)
