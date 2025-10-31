package ies

import "rrc/utils"

// NPRACH-ParametersFmt2-NB-r15-nprach-Parameters-r15 ::= SEQUENCE
// Extensible
type NprachParametersfmt2NbR15NprachParametersR15 struct {
	NprachPeriodicityR15              *NprachParametersfmt2NbR15NprachParametersR15NprachPeriodicityR15
	NprachStarttimeR15                *NprachParametersfmt2NbR15NprachParametersR15NprachStarttimeR15
	NprachSubcarrieroffsetR15         *NprachParametersfmt2NbR15NprachParametersR15NprachSubcarrieroffsetR15
	NprachNumsubcarriersR15           *NprachParametersfmt2NbR15NprachParametersR15NprachNumsubcarriersR15
	NprachSubcarriermsg3RangestartR15 *NprachParametersfmt2NbR15NprachParametersR15NprachSubcarriermsg3RangestartR15
	NpdcchNumrepetitionsRaR15         *NprachParametersfmt2NbR15NprachParametersR15NpdcchNumrepetitionsRaR15
	NpdcchStartsfCssRaR15             *NprachParametersfmt2NbR15NprachParametersR15NpdcchStartsfCssRaR15
	NpdcchOffsetRaR15                 *NprachParametersfmt2NbR15NprachParametersR15NpdcchOffsetRaR15
	NprachNumcbraStartsubcarriersR15  *NprachParametersfmt2NbR15NprachParametersR15NprachNumcbraStartsubcarriersR15
	NpdcchCarrierindexR15             *utils.INTEGER `lb:0,ub:maxNonAnchorCarriersNbR14`
}
