package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PUSCH-cyclicMapping-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPuschCyclicmappingR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPuschCyclicmappingR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPuschCyclicmappingR17EnumeratedTypea
	MimoParametersperbandMtrpPuschCyclicmappingR17EnumeratedTypeb
	MimoParametersperbandMtrpPuschCyclicmappingR17EnumeratedBoth
)
