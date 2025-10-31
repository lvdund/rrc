package ies

import "rrc/utils"

// SidelinkPreconfigNR-r16-sl-CSI-Acquisition-r16 ::= ENUMERATED
type SidelinkpreconfignrR16SlCsiAcquisitionR16 struct {
	Value utils.ENUMERATED
}

const (
	SidelinkpreconfignrR16SlCsiAcquisitionR16EnumeratedNothing = iota
	SidelinkpreconfignrR16SlCsiAcquisitionR16EnumeratedEnabled
)
