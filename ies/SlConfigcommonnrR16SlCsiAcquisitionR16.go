package ies

import "rrc/utils"

// SL-ConfigCommonNR-r16-sl-CSI-Acquisition-r16 ::= ENUMERATED
type SlConfigcommonnrR16SlCsiAcquisitionR16 struct {
	Value utils.ENUMERATED
}

const (
	SlConfigcommonnrR16SlCsiAcquisitionR16EnumeratedNothing = iota
	SlConfigcommonnrR16SlCsiAcquisitionR16EnumeratedEnabled
)
