package ies

import "rrc/utils"

// CarrierFreqListUTRA-TDD-Ext-r12 ::= SEQUENCE OF CarrierFreqUTRA-TDD-r12
// SIZE (1..maxUTRA-TDD-Carrier)
type CarrierfreqlistutraTddExtR12 struct {
	Value []CarrierfrequtraTddR12
}
