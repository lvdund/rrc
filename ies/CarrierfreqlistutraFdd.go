package ies

import "rrc/utils"

// CarrierFreqListUTRA-FDD ::= SEQUENCE OF CarrierFreqUTRA-FDD
// SIZE (1..maxUTRA-FDD-Carrier)
type CarrierfreqlistutraFdd struct {
	Value utils.Sequence[CarrierfrequtraFdd]
}
