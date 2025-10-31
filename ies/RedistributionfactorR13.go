package ies

import "rrc/utils"

// RedistributionFactor-r13 ::= utils.INTEGER (1..10)
type RedistributionfactorR13 struct {
	Value utils.INTEGER `lb:0,ub:10`
}
