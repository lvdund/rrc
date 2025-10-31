package ies

import "rrc/utils"

// BandNR-pucch-SpatialRelInfoMAC-CE ::= ENUMERATED
type BandnrPucchSpatialrelinfomacCe struct {
	Value utils.ENUMERATED
}

const (
	BandnrPucchSpatialrelinfomacCeEnumeratedNothing = iota
	BandnrPucchSpatialrelinfomacCeEnumeratedSupported
)
