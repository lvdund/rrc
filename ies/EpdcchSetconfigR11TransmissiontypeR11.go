package ies

import "rrc/utils"

// EPDCCH-SetConfig-r11-transmissionType-r11 ::= ENUMERATED
type EpdcchSetconfigR11TransmissiontypeR11 struct {
	Value utils.ENUMERATED
}

const (
	EpdcchSetconfigR11TransmissiontypeR11EnumeratedNothing = iota
	EpdcchSetconfigR11TransmissiontypeR11EnumeratedLocalised
	EpdcchSetconfigR11TransmissiontypeR11EnumeratedDistributed
)
