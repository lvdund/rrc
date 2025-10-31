package ies

import "rrc/utils"

// SRS-Resource-partialFreqSounding-r17-enableStartRBHopping-r17 ::= ENUMERATED
type SrsResourcePartialfreqsoundingR17EnablestartrbhoppingR17 struct {
	Value utils.ENUMERATED
}

const (
	SrsResourcePartialfreqsoundingR17EnablestartrbhoppingR17EnumeratedNothing = iota
	SrsResourcePartialfreqsoundingR17EnablestartrbhoppingR17EnumeratedEnable
)
