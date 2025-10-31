package ies

import "rrc/utils"

// SR-NPRACH-Resource-NB-r15 ::= SEQUENCE
type SrNprachResourceNbR15 struct {
	NprachCarrierindexR15    utils.INTEGER `lb:0,ub:maxNonAnchorCarriersNbR14`
	NprachResourceindexR15   utils.INTEGER `lb:0,ub:maxNPRACHResourcesNbR13`
	NprachSubcarrierindexR15 SrNprachResourceNbR15NprachSubcarrierindexR15
	P0SrR15                  utils.INTEGER `lb:0,ub:24`
	AlphaR15                 SrNprachResourceNbR15AlphaR15
}
