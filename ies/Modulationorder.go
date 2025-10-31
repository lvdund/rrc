package ies

import "rrc/utils"

// ModulationOrder ::= ENUMERATED
type Modulationorder struct {
	Value utils.ENUMERATED
}

const (
	ModulationorderEnumeratedNothing = iota
	ModulationorderEnumeratedBpsk_Halfpi
	ModulationorderEnumeratedBpsk
	ModulationorderEnumeratedQpsk
	ModulationorderEnumeratedQam16
	ModulationorderEnumeratedQam64
	ModulationorderEnumeratedQam256
)
