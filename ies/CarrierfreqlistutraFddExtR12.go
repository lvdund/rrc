package ies

import "rrc/utils"

// CarrierFreqListUTRA-FDD-Ext-r12 ::= SEQUENCE OF CarrierFreqUTRA-FDD-Ext-r12
// SIZE (1..maxUTRA-FDD-Carrier)
type CarrierfreqlistutraFddExtR12 struct {
	Value []CarrierfrequtraFddExtR12
}
