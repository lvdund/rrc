package ies

import "rrc/utils"

// FeatureSetUplink-v1720-pucch-Repetition-F0-1-2-3-4-DynamicIndication-r17 ::= ENUMERATED
type FeaturesetuplinkV1720PucchRepetitionF01234DynamicindicationR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1720PucchRepetitionF01234DynamicindicationR17EnumeratedNothing = iota
	FeaturesetuplinkV1720PucchRepetitionF01234DynamicindicationR17EnumeratedSupported
)
