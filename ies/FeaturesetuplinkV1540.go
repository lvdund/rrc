package ies

// FeatureSetUplink-v1540 ::= SEQUENCE
type FeaturesetuplinkV1540 struct {
	Zeroslotoffsetaperiodicsrs     *FeaturesetuplinkV1540Zeroslotoffsetaperiodicsrs
	PaPhasediscontinuityimpacts    *FeaturesetuplinkV1540PaPhasediscontinuityimpacts
	PuschSeparationwithgap         *FeaturesetuplinkV1540PuschSeparationwithgap
	PuschProcessingtype2           *FeaturesetuplinkV1540PuschProcessingtype2
	UlMcsTablealtDynamicindication *FeaturesetuplinkV1540UlMcsTablealtDynamicindication
}
