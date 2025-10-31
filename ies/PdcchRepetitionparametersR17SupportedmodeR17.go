package ies

import "rrc/utils"

// PDCCH-RepetitionParameters-r17-supportedMode-r17 ::= ENUMERATED
type PdcchRepetitionparametersR17SupportedmodeR17 struct {
	Value utils.ENUMERATED
}

const (
	PdcchRepetitionparametersR17SupportedmodeR17EnumeratedNothing = iota
	PdcchRepetitionparametersR17SupportedmodeR17EnumeratedIntra_Span
	PdcchRepetitionparametersR17SupportedmodeR17EnumeratedInter_Span
	PdcchRepetitionparametersR17SupportedmodeR17EnumeratedBoth
)
