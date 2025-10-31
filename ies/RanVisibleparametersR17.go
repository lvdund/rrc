package ies

import "rrc/utils"

// RAN-VisibleParameters-r17 ::= SEQUENCE
// Extensible
type RanVisibleparametersR17 struct {
	RanVisibleperiodicityR17             *RanVisibleparametersR17RanVisibleperiodicityR17
	NumberofbufferlevelentriesR17        *utils.INTEGER `lb:0,ub:8`
	ReportplayoutdelayformediastartupR17 *utils.BOOLEAN
}
