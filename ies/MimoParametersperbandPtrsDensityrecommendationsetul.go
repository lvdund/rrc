package ies

// MIMO-ParametersPerBand-ptrs-DensityRecommendationSetUL ::= SEQUENCE
type MimoParametersperbandPtrsDensityrecommendationsetul struct {
	Scs15khz  *PtrsDensityrecommendationul
	Scs30khz  *PtrsDensityrecommendationul
	Scs60khz  *PtrsDensityrecommendationul
	Scs120khz *PtrsDensityrecommendationul
}
