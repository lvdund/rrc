package ies

// IRAT-ParametersUTRA-v9c0 ::= SEQUENCE
type IratParametersutraV9c0 struct {
	VoiceoverpsHsUtraFddR9            *IratParametersutraV9c0VoiceoverpsHsUtraFddR9
	VoiceoverpsHsUtraTdd128R9         *IratParametersutraV9c0VoiceoverpsHsUtraTdd128R9
	SrvccFromutraFddToutraFddR9       *IratParametersutraV9c0SrvccFromutraFddToutraFddR9
	SrvccFromutraFddTogeranR9         *IratParametersutraV9c0SrvccFromutraFddTogeranR9
	SrvccFromutraTdd128ToutraTdd128R9 *IratParametersutraV9c0SrvccFromutraTdd128ToutraTdd128R9
	SrvccFromutraTdd128TogeranR9      *IratParametersutraV9c0SrvccFromutraTdd128TogeranR9
}
