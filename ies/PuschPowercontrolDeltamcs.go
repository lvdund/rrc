package ies

import "rrc/utils"

// PUSCH-PowerControl-deltaMCS ::= ENUMERATED
type PuschPowercontrolDeltamcs struct {
	Value utils.ENUMERATED
}

const (
	PuschPowercontrolDeltamcsEnumeratedNothing = iota
	PuschPowercontrolDeltamcsEnumeratedEnabled
)
