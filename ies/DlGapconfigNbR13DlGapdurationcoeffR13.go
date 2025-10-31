package ies

import "rrc/utils"

// DL-GapConfig-NB-r13-dl-GapDurationCoeff-r13 ::= ENUMERATED
type DlGapconfigNbR13DlGapdurationcoeffR13 struct {
	Value utils.ENUMERATED
}

const (
	DlGapconfigNbR13DlGapdurationcoeffR13EnumeratedNothing = iota
	DlGapconfigNbR13DlGapdurationcoeffR13EnumeratedOneeighth
	DlGapconfigNbR13DlGapdurationcoeffR13EnumeratedOnefourth
	DlGapconfigNbR13DlGapdurationcoeffR13EnumeratedThreeeighth
	DlGapconfigNbR13DlGapdurationcoeffR13EnumeratedOnehalf
)
