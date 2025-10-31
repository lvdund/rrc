package ies

import "rrc/utils"

// MeasResultUTRA-measResult ::= SEQUENCE
// Extensible
type MeasresultutraMeasresult struct {
	UtraRscp *utils.INTEGER `lb:0,ub:91`
	UtraEcn0 *utils.INTEGER `lb:0,ub:49`
}
