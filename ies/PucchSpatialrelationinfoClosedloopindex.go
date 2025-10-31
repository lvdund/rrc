package ies

import "rrc/utils"

// PUCCH-SpatialRelationInfo-closedLoopIndex ::= ENUMERATED
type PucchSpatialrelationinfoClosedloopindex struct {
	Value utils.ENUMERATED
}

const (
	PucchSpatialrelationinfoClosedloopindexEnumeratedNothing = iota
	PucchSpatialrelationinfoClosedloopindexEnumeratedI0
	PucchSpatialrelationinfoClosedloopindexEnumeratedI1
)
