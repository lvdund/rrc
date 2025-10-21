package ies

import "rrc/utils"

// MeasResultList2CDMA2000-r9 ::= SEQUENCE OF MeasResult2CDMA2000-r9
// SIZE (1..maxFreq)
type Measresultlist2cdma2000R9 struct {
	Value utils.Sequence[Measresult2cdma2000R9]
}
