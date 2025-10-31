package ies

import "rrc/utils"

// SI-SchedulingInfo ::= SEQUENCE
// Extensible
type SiSchedulinginfo struct {
	Schedulinginfolist      []Schedulinginfo `lb:1,ub:maxSIMessage`
	SiWindowlength          SiSchedulinginfoSiWindowlength
	SiRequestconfig         *SiRequestconfig
	SiRequestconfigsul      *SiRequestconfig
	Systeminformationareaid *utils.BITSTRING `lb:24,ub:24`
}
