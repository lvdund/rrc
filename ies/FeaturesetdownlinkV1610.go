package ies

// FeatureSetDownlink-v1610 ::= SEQUENCE
type FeaturesetdownlinkV1610 struct {
	CbgpdschProcessingtype1DifferenttbPerslotR16 *FeaturesetdownlinkV1610CbgpdschProcessingtype1DifferenttbPerslotR16
	CbgpdschProcessingtype2DifferenttbPerslotR16 *FeaturesetdownlinkV1610CbgpdschProcessingtype2DifferenttbPerslotR16
	IntrafreqdapsR16                             *FeaturesetdownlinkV1610IntrafreqdapsR16
	IntrabandfreqseparationdlV1620               *FreqseparationclassdlV1620
	IntrabandfreqseparationdlOnlyR16             *FreqseparationclassdlOnlyR16
	PdcchMonitoringR16                           *FeaturesetdownlinkV1610PdcchMonitoringR16
	PdcchMonitoringmixedR16                      *FeaturesetdownlinkV1610PdcchMonitoringmixedR16
	CrosscarrierschedulingprocessingDiffscsR16   *FeaturesetdownlinkV1610CrosscarrierschedulingprocessingDiffscsR16
	SingledciSdmSchemeR16                        *FeaturesetdownlinkV1610SingledciSdmSchemeR16
}
