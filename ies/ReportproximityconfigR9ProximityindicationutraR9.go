package ies

import "rrc/utils"

// ReportProximityConfig-r9-proximityIndicationUTRA-r9 ::= ENUMERATED
type ReportproximityconfigR9ProximityindicationutraR9 struct {
	Value utils.ENUMERATED
}

const (
	ReportproximityconfigR9ProximityindicationutraR9EnumeratedNothing = iota
	ReportproximityconfigR9ProximityindicationutraR9EnumeratedEnabled
)
