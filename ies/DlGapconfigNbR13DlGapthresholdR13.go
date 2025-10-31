package ies

import "rrc/utils"

// DL-GapConfig-NB-r13-dl-GapThreshold-r13 ::= ENUMERATED
type DlGapconfigNbR13DlGapthresholdR13 struct {
	Value utils.ENUMERATED
}

const (
	DlGapconfigNbR13DlGapthresholdR13EnumeratedNothing = iota
	DlGapconfigNbR13DlGapthresholdR13EnumeratedN32
	DlGapconfigNbR13DlGapthresholdR13EnumeratedN64
	DlGapconfigNbR13DlGapthresholdR13EnumeratedN128
	DlGapconfigNbR13DlGapthresholdR13EnumeratedN256
)
