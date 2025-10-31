package ies

// MIMO-ParametersPerBand-ptrs-DensityRecommendationSetDL ::= SEQUENCE
type MimoParametersperbandPtrsDensityrecommendationsetdl struct {
	Scs15khz  *PtrsDensityrecommendationdl
	Scs30khz  *PtrsDensityrecommendationdl
	Scs60khz  *PtrsDensityrecommendationdl
	Scs120khz *PtrsDensityrecommendationdl
}
