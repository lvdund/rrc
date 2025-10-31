package ies

import "rrc/utils"

// FeatureSetUplink-v1720-extendedDC-LocationReport-r17 ::= ENUMERATED
type FeaturesetuplinkV1720ExtendeddcLocationreportR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1720ExtendeddcLocationreportR17EnumeratedNothing = iota
	FeaturesetuplinkV1720ExtendeddcLocationreportR17EnumeratedSupported
)
