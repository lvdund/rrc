package ies

import "rrc/utils"

// PTRS-UplinkConfig-transformPrecoderEnabled ::= SEQUENCE
type PtrsUplinkconfigTransformprecoderenabled struct {
	Sampledensity                 []utils.INTEGER `lb:5,ub:5`
	Timedensitytransformprecoding *PtrsUplinkconfigTransformprecoderenabledTimedensitytransformprecoding
}
