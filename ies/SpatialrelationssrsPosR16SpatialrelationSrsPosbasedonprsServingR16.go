package ies

import "rrc/utils"

// SpatialRelationsSRS-Pos-r16-spatialRelation-SRS-PosBasedOnPRS-Serving-r16 ::= ENUMERATED
type SpatialrelationssrsPosR16SpatialrelationSrsPosbasedonprsServingR16 struct {
	Value utils.ENUMERATED
}

const (
	SpatialrelationssrsPosR16SpatialrelationSrsPosbasedonprsServingR16EnumeratedNothing = iota
	SpatialrelationssrsPosR16SpatialrelationSrsPosbasedonprsServingR16EnumeratedSupported
)
