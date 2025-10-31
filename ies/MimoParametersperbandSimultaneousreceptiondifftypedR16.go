package ies

import "rrc/utils"

// MIMO-ParametersPerBand-simultaneousReceptionDiffTypeD-r16 ::= ENUMERATED
type MimoParametersperbandSimultaneousreceptiondifftypedR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSimultaneousreceptiondifftypedR16EnumeratedNothing = iota
	MimoParametersperbandSimultaneousreceptiondifftypedR16EnumeratedSupported
)
