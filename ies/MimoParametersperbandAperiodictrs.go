package ies

import "rrc/utils"

// MIMO-ParametersPerBand-aperiodicTRS ::= ENUMERATED
type MimoParametersperbandAperiodictrs struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandAperiodictrsEnumeratedNothing = iota
	MimoParametersperbandAperiodictrsEnumeratedSupported
)
