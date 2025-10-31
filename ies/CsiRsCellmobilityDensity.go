package ies

import "rrc/utils"

// CSI-RS-CellMobility-density ::= ENUMERATED
type CsiRsCellmobilityDensity struct {
	Value utils.ENUMERATED
}

const (
	CsiRsCellmobilityDensityEnumeratedNothing = iota
	CsiRsCellmobilityDensityEnumeratedD1
	CsiRsCellmobilityDensityEnumeratedD3
)
