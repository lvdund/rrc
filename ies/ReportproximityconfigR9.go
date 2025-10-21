package ies

import "rrc/utils"

// ReportProximityConfig-r9 ::= SEQUENCE
type ReportproximityconfigR9 struct {
	ProximityindicationeutraR9 *utils.ENUMERATED
	ProximityindicationutraR9  *utils.ENUMERATED
}
