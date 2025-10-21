package ies

import "rrc/utils"

// PUR-PUSCH-Config-r16 ::= SEQUENCE
type PurPuschConfigR16 struct {
	PurGrantinfoR16        *interface{}
	PurPuschFreqhoppingR16 bool
	P0UePuschR16           utils.INTEGER
	AlphaR16               AlphaR12
	PuschCyclicshiftR16    utils.ENUMERATED
	PuschNbMaxtbsR16       bool
	LocationceModebR16     *utils.INTEGER
}
