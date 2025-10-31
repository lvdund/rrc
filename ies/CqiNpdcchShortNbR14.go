package ies

import "rrc/utils"

// CQI-NPDCCH-Short-NB-r14 ::= ENUMERATED
type CqiNpdcchShortNbR14 struct {
	Value utils.ENUMERATED
}

const (
	CqiNpdcchShortNbR14EnumeratedNothing = iota
	CqiNpdcchShortNbR14EnumeratedNomeasurements
	CqiNpdcchShortNbR14EnumeratedCandidaterep_1
	CqiNpdcchShortNbR14EnumeratedCandidaterep_2
	CqiNpdcchShortNbR14EnumeratedCandidaterep_3
)
