package ies

import "rrc/utils"

// ReportProximityConfig-r9-proximityIndicationEUTRA-r9 ::= ENUMERATED
type ReportproximityconfigR9ProximityindicationeutraR9 struct {
	Value utils.ENUMERATED
}

const (
	ReportproximityconfigR9ProximityindicationeutraR9EnumeratedNothing = iota
	ReportproximityconfigR9ProximityindicationeutraR9EnumeratedEnabled
)
