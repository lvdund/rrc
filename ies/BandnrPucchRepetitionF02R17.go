package ies

import "rrc/utils"

// BandNR-pucch-Repetition-F0-2-r17 ::= ENUMERATED
type BandnrPucchRepetitionF02R17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPucchRepetitionF02R17EnumeratedNothing = iota
	BandnrPucchRepetitionF02R17EnumeratedSupported
)
