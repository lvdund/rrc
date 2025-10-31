package ies

// NPRACH-Parameters-NB-r13 ::= SEQUENCE
type NprachParametersNbR13 struct {
	NprachPeriodicityR13                NprachParametersNbR13NprachPeriodicityR13
	NprachStarttimeR13                  NprachParametersNbR13NprachStarttimeR13
	NprachSubcarrieroffsetR13           NprachParametersNbR13NprachSubcarrieroffsetR13
	NprachNumsubcarriersR13             NprachParametersNbR13NprachNumsubcarriersR13
	NprachSubcarriermsg3RangestartR13   NprachParametersNbR13NprachSubcarriermsg3RangestartR13
	MaxnumpreambleattemptceR13          NprachParametersNbR13MaxnumpreambleattemptceR13
	NumrepetitionsperpreambleattemptR13 NprachParametersNbR13NumrepetitionsperpreambleattemptR13
	NpdcchNumrepetitionsRaR13           NprachParametersNbR13NpdcchNumrepetitionsRaR13
	NpdcchStartsfCssRaR13               NprachParametersNbR13NpdcchStartsfCssRaR13
	NpdcchOffsetRaR13                   NprachParametersNbR13NpdcchOffsetRaR13
}
