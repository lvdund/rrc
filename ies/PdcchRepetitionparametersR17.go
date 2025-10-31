package ies

// PDCCH-RepetitionParameters-r17 ::= SEQUENCE
type PdcchRepetitionparametersR17 struct {
	SupportedmodeR17  PdcchRepetitionparametersR17SupportedmodeR17
	LimitxPerccR17    *PdcchRepetitionparametersR17LimitxPerccR17
	LimitxAcrossccR17 *PdcchRepetitionparametersR17LimitxAcrossccR17
}
