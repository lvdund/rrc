package ies

import "rrc/utils"

// CellsToAddModExt-v1710-ntn-PolarizationUL-r17 ::= ENUMERATED
type CellstoaddmodextV1710NtnPolarizationulR17 struct {
	Value utils.ENUMERATED
}

const (
	CellstoaddmodextV1710NtnPolarizationulR17EnumeratedNothing = iota
	CellstoaddmodextV1710NtnPolarizationulR17EnumeratedRhcp
	CellstoaddmodextV1710NtnPolarizationulR17EnumeratedLhcp
	CellstoaddmodextV1710NtnPolarizationulR17EnumeratedLinear
)
