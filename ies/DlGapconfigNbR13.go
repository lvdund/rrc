package ies

import "rrc/utils"

// DL-GapConfig-NB-r13 ::= SEQUENCE
type DlGapconfigNbR13 struct {
	DlGapthresholdR13     utils.ENUMERATED
	DlGapperiodicityR13   utils.ENUMERATED
	DlGapdurationcoeffR13 utils.ENUMERATED
}
