package ies

import "rrc/utils"

// SRI-PUSCH-PowerControl-sri-PUSCH-ClosedLoopIndex ::= ENUMERATED
type SriPuschPowercontrolSriPuschClosedloopindex struct {
	Value utils.ENUMERATED
}

const (
	SriPuschPowercontrolSriPuschClosedloopindexEnumeratedNothing = iota
	SriPuschPowercontrolSriPuschClosedloopindexEnumeratedI0
	SriPuschPowercontrolSriPuschClosedloopindexEnumeratedI1
)
