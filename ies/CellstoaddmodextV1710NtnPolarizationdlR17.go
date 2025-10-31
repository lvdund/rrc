package ies

import "rrc/utils"

// CellsToAddModExt-v1710-ntn-PolarizationDL-r17 ::= ENUMERATED
type CellstoaddmodextV1710NtnPolarizationdlR17 struct {
	Value utils.ENUMERATED
}

const (
	CellstoaddmodextV1710NtnPolarizationdlR17EnumeratedNothing = iota
	CellstoaddmodextV1710NtnPolarizationdlR17EnumeratedRhcp
	CellstoaddmodextV1710NtnPolarizationdlR17EnumeratedLhcp
	CellstoaddmodextV1710NtnPolarizationdlR17EnumeratedLinear
)
