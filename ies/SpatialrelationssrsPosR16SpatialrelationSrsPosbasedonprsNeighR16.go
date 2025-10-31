package ies

import "rrc/utils"

// SpatialRelationsSRS-Pos-r16-spatialRelation-SRS-PosBasedOnPRS-Neigh-r16 ::= ENUMERATED
type SpatialrelationssrsPosR16SpatialrelationSrsPosbasedonprsNeighR16 struct {
	Value utils.ENUMERATED
}

const (
	SpatialrelationssrsPosR16SpatialrelationSrsPosbasedonprsNeighR16EnumeratedNothing = iota
	SpatialrelationssrsPosR16SpatialrelationSrsPosbasedonprsNeighR16EnumeratedSupported
)
