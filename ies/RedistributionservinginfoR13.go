package ies

import "rrc/utils"

// RedistributionServingInfo-r13 ::= SEQUENCE
type RedistributionservinginfoR13 struct {
	RedistributionfactorservingR13 utils.INTEGER `lb:0,ub:10`
	RedistributionfactorcellR13    *bool
	T360R13                        RedistributionservinginfoR13T360R13
	RedistronpagingonlyR13         *bool
}
