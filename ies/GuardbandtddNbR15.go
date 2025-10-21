package ies

import "rrc/utils"

// GuardbandTDD-NB-r15 ::= SEQUENCE
type GuardbandtddNbR15 struct {
	RasteroffsetR15     ChannelrasteroffsetNbR13
	SibGuardbandinfoR15 interface{}
	EutraBandwitdhR15   utils.ENUMERATED
}
