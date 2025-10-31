package ies

import "rrc/utils"

// SpatialRelationsSRS-Pos-r16-spatialRelation-SRS-PosBasedOnSSB-Serving-r16 ::= ENUMERATED
type SpatialrelationssrsPosR16SpatialrelationSrsPosbasedonssbServingR16 struct {
	Value utils.ENUMERATED
}

const (
	SpatialrelationssrsPosR16SpatialrelationSrsPosbasedonssbServingR16EnumeratedNothing = iota
	SpatialrelationssrsPosR16SpatialrelationSrsPosbasedonssbServingR16EnumeratedSupported
)
