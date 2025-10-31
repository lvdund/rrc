package ies

import "rrc/utils"

// SL-PHY-MAC-RLC-Config-r16-sl-CSI-Acquisition-r16 ::= ENUMERATED
type SlPhyMacRlcConfigR16SlCsiAcquisitionR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPhyMacRlcConfigR16SlCsiAcquisitionR16EnumeratedNothing = iota
	SlPhyMacRlcConfigR16SlCsiAcquisitionR16EnumeratedEnabled
)
