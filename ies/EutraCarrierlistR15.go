package ies

import "rrc/utils"

// EUTRA-CarrierList-r15 ::= SEQUENCE OF MeasIdleCarrierEUTRA-r15
// SIZE (1..maxFreqIdle-r15)
type EutraCarrierlistR15 struct {
	Value []MeasidlecarriereutraR15
}
