package ies

import "rrc/utils"

// TDD-UL-DL-Pattern-dl-UL-TransmissionPeriodicity-v1530 ::= ENUMERATED
type TddUlDlPatternDlUlTransmissionperiodicityV1530 struct {
	Value utils.ENUMERATED
}

const (
	TddUlDlPatternDlUlTransmissionperiodicityV1530EnumeratedNothing = iota
	TddUlDlPatternDlUlTransmissionperiodicityV1530EnumeratedMs3
	TddUlDlPatternDlUlTransmissionperiodicityV1530EnumeratedMs4
)
