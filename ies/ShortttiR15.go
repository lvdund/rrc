package ies

import "rrc/utils"

// ShortTTI-r15 ::= SEQUENCE
type ShortttiR15 struct {
	DlSttiLengthR15 *ShortttiLengthR15
	UlSttiLengthR15 *ShortttiLengthR15
}
