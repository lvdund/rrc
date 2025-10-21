package ies

import "rrc/utils"

// MeasResultNR-r15 ::= SEQUENCE
// Extensible
type MeasresultnrR15 struct {
	RsrpresultR15   *RsrpRangenrR15
	RsrqresultR15   *RsrqRangenrR15
	RsSinrResultR15 *RsSinrRangenrR15
}
