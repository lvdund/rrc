package ies

import "rrc/utils"

// DL-GapConfig-NB-v1530-dl-GapPeriodicity-v1530 ::= ENUMERATED
type DlGapconfigNbV1530DlGapperiodicityV1530 struct {
	Value utils.ENUMERATED
}

const (
	DlGapconfigNbV1530DlGapperiodicityV1530EnumeratedNothing = iota
	DlGapconfigNbV1530DlGapperiodicityV1530EnumeratedSf1024
)
