package ies

import "rrc/utils"

// CQI-NPDCCH-Short-NB-r14 ::= ENUMERATED
type CqiNpdcchShortNbR14 struct {
	Value utils.ENUMERATED
}

const (
	CqiNpdcchShortNbR14Nomeasurements = 0
	CqiNpdcchShortNbR14Candidaterep1  = 1
	CqiNpdcchShortNbR14Candidaterep2  = 2
	CqiNpdcchShortNbR14Candidaterep3  = 3
)
