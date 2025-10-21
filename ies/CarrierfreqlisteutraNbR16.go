package ies

import "rrc/utils"

// CarrierFreqListEUTRA-NB-r16 ::= SEQUENCE OF CarrierFreqEUTRA-NB-r16
// SIZE (1..maxFreqEUTRA-NB-r16)
type CarrierfreqlisteutraNbR16 struct {
	Value []CarrierfreqeutraNbR16
}
