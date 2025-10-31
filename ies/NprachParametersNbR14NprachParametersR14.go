package ies

import "rrc/utils"

// NPRACH-Parameters-NB-r14-nprach-Parameters-r14 ::= SEQUENCE
// Extensible
type NprachParametersNbR14NprachParametersR14 struct {
	NprachPeriodicityR14              *NprachParametersNbR14NprachParametersR14NprachPeriodicityR14
	NprachStarttimeR14                *NprachParametersNbR14NprachParametersR14NprachStarttimeR14
	NprachSubcarrieroffsetR14         *NprachParametersNbR14NprachParametersR14NprachSubcarrieroffsetR14
	NprachNumsubcarriersR14           *NprachParametersNbR14NprachParametersR14NprachNumsubcarriersR14
	NprachSubcarriermsg3RangestartR14 *NprachParametersNbR14NprachParametersR14NprachSubcarriermsg3RangestartR14
	NpdcchNumrepetitionsRaR14         *NprachParametersNbR14NprachParametersR14NpdcchNumrepetitionsRaR14
	NpdcchStartsfCssRaR14             *NprachParametersNbR14NprachParametersR14NpdcchStartsfCssRaR14
	NpdcchOffsetRaR14                 *NprachParametersNbR14NprachParametersR14NpdcchOffsetRaR14
	NprachNumcbraStartsubcarriersR14  *NprachParametersNbR14NprachParametersR14NprachNumcbraStartsubcarriersR14
	NpdcchCarrierindexR14             *utils.INTEGER `lb:0,ub:maxNonAnchorCarriersNbR14`
}
