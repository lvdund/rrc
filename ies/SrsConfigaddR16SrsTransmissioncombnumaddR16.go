package ies

import "rrc/utils"

// SRS-ConfigAdd-r16-srs-TransmissionCombNumAdd-r16 ::= ENUMERATED
type SrsConfigaddR16SrsTransmissioncombnumaddR16 struct {
	Value utils.ENUMERATED
}

const (
	SrsConfigaddR16SrsTransmissioncombnumaddR16EnumeratedNothing = iota
	SrsConfigaddR16SrsTransmissioncombnumaddR16EnumeratedN2
	SrsConfigaddR16SrsTransmissioncombnumaddR16EnumeratedN4
)
