package ies

import "rrc/utils"

// PUR-PUSCH-Config-r16 ::= SEQUENCE
type PurPuschConfigR16 struct {
	PurGrantinfoR16        *PurPuschConfigR16PurGrantinfoR16
	PurPuschFreqhoppingR16 utils.BOOLEAN
	P0UePuschR16           utils.INTEGER `lb:0,ub:7`
	AlphaR16               AlphaR12
	PuschCyclicshiftR16    PurPuschConfigR16PuschCyclicshiftR16
	PuschNbMaxtbsR16       utils.BOOLEAN
	LocationceModebR16     *utils.INTEGER `lb:0,ub:5`
}
