package ies

import "rrc/utils"

// IRAT-ParametersUTRA-v9c0 ::= SEQUENCE
type IratParametersutraV9c0 struct {
	VoiceoverpsHsUtraFddR9            *utils.ENUMERATED
	VoiceoverpsHsUtraTdd128R9         *utils.ENUMERATED
	SrvccFromutraFddToutraFddR9       *utils.ENUMERATED
	SrvccFromutraFddTogeranR9         *utils.ENUMERATED
	SrvccFromutraTdd128ToutraTdd128R9 *utils.ENUMERATED
	SrvccFromutraTdd128TogeranR9      *utils.ENUMERATED
}
