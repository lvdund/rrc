package ies

import "rrc/utils"

// SpatialRelationsSRS-Pos-r16-spatialRelation-SRS-PosBasedOnSSB-Neigh-r16 ::= ENUMERATED
type SpatialrelationssrsPosR16SpatialrelationSrsPosbasedonssbNeighR16 struct {
	Value utils.ENUMERATED
}

const (
	SpatialrelationssrsPosR16SpatialrelationSrsPosbasedonssbNeighR16EnumeratedNothing = iota
	SpatialrelationssrsPosR16SpatialrelationSrsPosbasedonssbNeighR16EnumeratedSupported
)
