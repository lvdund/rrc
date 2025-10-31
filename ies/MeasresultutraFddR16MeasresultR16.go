package ies

import "rrc/utils"

// MeasResultUTRA-FDD-r16-measResult-r16 ::= SEQUENCE
type MeasresultutraFddR16MeasresultR16 struct {
	UtraFddRscpR16 *utils.INTEGER `lb:0,ub:91`
	UtraFddEcn0R16 *utils.INTEGER `lb:0,ub:49`
}
