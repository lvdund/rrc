package ies

import "rrc/utils"

// RedistributionServingInfo-r13 ::= SEQUENCE
type RedistributionservinginfoR13 struct {
	RedistributionfactorservingR13 utils.INTEGER
	RedistributionfactorcellR13    *utils.ENUMERATED
	T360R13                        utils.ENUMERATED
	RedistronpagingonlyR13         *utils.ENUMERATED
}
