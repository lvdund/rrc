package ies

import "rrc/utils"

// MeasIdleCarrierNR-r16-subcarrierSpacingSSB-r16 ::= ENUMERATED
type MeasidlecarriernrR16SubcarrierspacingssbR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasidlecarriernrR16SubcarrierspacingssbR16EnumeratedNothing = iota
	MeasidlecarriernrR16SubcarrierspacingssbR16EnumeratedKhz15
	MeasidlecarriernrR16SubcarrierspacingssbR16EnumeratedKhz30
	MeasidlecarriernrR16SubcarrierspacingssbR16EnumeratedKhz120
	MeasidlecarriernrR16SubcarrierspacingssbR16EnumeratedKhz240
)
