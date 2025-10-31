package ies

import "rrc/utils"

// SpatialRelationsSRS-Pos-r16-spatialRelation-SRS-PosBasedOnCSI-RS-Serving-r16 ::= ENUMERATED
type SpatialrelationssrsPosR16SpatialrelationSrsPosbasedoncsiRsServingR16 struct {
	Value utils.ENUMERATED
}

const (
	SpatialrelationssrsPosR16SpatialrelationSrsPosbasedoncsiRsServingR16EnumeratedNothing = iota
	SpatialrelationssrsPosR16SpatialrelationSrsPosbasedoncsiRsServingR16EnumeratedSupported
)
