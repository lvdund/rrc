package ies

import "rrc/utils"

// DL-GapConfig-NB-r13-dl-GapPeriodicity-r13 ::= ENUMERATED
type DlGapconfigNbR13DlGapperiodicityR13 struct {
	Value utils.ENUMERATED
}

const (
	DlGapconfigNbR13DlGapperiodicityR13EnumeratedNothing = iota
	DlGapconfigNbR13DlGapperiodicityR13EnumeratedSf64
	DlGapconfigNbR13DlGapperiodicityR13EnumeratedSf128
	DlGapconfigNbR13DlGapperiodicityR13EnumeratedSf256
	DlGapconfigNbR13DlGapperiodicityR13EnumeratedSf512
)
