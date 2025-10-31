package ies

import "rrc/utils"

// CellChangeOrder-targetRAT-Type-geran ::= SEQUENCE
type CellchangeorderTargetratTypeGeran struct {
	Physcellid          Physcellidgeran
	Carrierfreq         Carrierfreqgeran
	Networkcontrolorder *utils.BITSTRING `lb:2,ub:2`
	Systeminformation   *SiOrpsiGeran
}
