package ies

// FeatureSetDownlink-v1540 ::= SEQUENCE
type FeaturesetdownlinkV1540 struct {
	OneflDmrsTwoadditionaldmrsDl           *FeaturesetdownlinkV1540OneflDmrsTwoadditionaldmrsDl
	AdditionaldmrsDlAlt                    *FeaturesetdownlinkV1540AdditionaldmrsDlAlt
	TwoflDmrsTwoadditionaldmrsDl           *FeaturesetdownlinkV1540TwoflDmrsTwoadditionaldmrsDl
	OneflDmrsThreeadditionaldmrsDl         *FeaturesetdownlinkV1540OneflDmrsThreeadditionaldmrsDl
	PdcchMonitoringanyoccasionswithspangap *FeaturesetdownlinkV1540PdcchMonitoringanyoccasionswithspangap
	PdschSeparationwithgap                 *FeaturesetdownlinkV1540PdschSeparationwithgap
	PdschProcessingtype2                   *FeaturesetdownlinkV1540PdschProcessingtype2
	PdschProcessingtype2Limited            *FeaturesetdownlinkV1540PdschProcessingtype2Limited
	DlMcsTablealtDynamicindication         *FeaturesetdownlinkV1540DlMcsTablealtDynamicindication
}
