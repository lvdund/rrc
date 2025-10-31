package ies

import "rrc/utils"

// MIMO-ParametersPerBand-simul-SpatialRelationUpdatePUCCHResGroup-r16 ::= ENUMERATED
type MimoParametersperbandSimulSpatialrelationupdatepucchresgroupR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSimulSpatialrelationupdatepucchresgroupR16EnumeratedNothing = iota
	MimoParametersperbandSimulSpatialrelationupdatepucchresgroupR16EnumeratedSupported
)
