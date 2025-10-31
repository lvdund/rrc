package ies

import "rrc/utils"

// PTRS-DensityRecommendationDL ::= SEQUENCE
type PtrsDensityrecommendationdl struct {
	Frequencydensity1 utils.INTEGER `lb:0,ub:276`
	Frequencydensity2 utils.INTEGER `lb:0,ub:276`
	Timedensity1      utils.INTEGER `lb:0,ub:29`
	Timedensity2      utils.INTEGER `lb:0,ub:29`
	Timedensity3      utils.INTEGER `lb:0,ub:29`
}
