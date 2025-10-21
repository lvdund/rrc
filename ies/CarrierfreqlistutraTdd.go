package ies

import "rrc/utils"

// CarrierFreqListUTRA-TDD ::= SEQUENCE OF CarrierFreqUTRA-TDD
// SIZE (1..maxUTRA-TDD-Carrier)
type CarrierfreqlistutraTdd struct {
	Value utils.Sequence[CarrierfrequtraTdd]
}
