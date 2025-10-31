package ies

import "rrc/utils"

// PTRS-DensityRecommendationUL ::= SEQUENCE
type PtrsDensityrecommendationul struct {
	Frequencydensity1 utils.INTEGER `lb:0,ub:276`
	Frequencydensity2 utils.INTEGER `lb:0,ub:276`
	Timedensity1      utils.INTEGER `lb:0,ub:29`
	Timedensity2      utils.INTEGER `lb:0,ub:29`
	Timedensity3      utils.INTEGER `lb:0,ub:29`
	Sampledensity1    utils.INTEGER `lb:0,ub:276`
	Sampledensity2    utils.INTEGER `lb:0,ub:276`
	Sampledensity3    utils.INTEGER `lb:0,ub:276`
	Sampledensity4    utils.INTEGER `lb:0,ub:276`
	Sampledensity5    utils.INTEGER `lb:0,ub:276`
}
