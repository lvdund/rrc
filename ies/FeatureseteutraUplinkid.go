package ies

import "rrc/utils"

// FeatureSetEUTRA-UplinkId ::= utils.INTEGER (0..maxEUTRA-UL-FeatureSets)
type FeatureseteutraUplinkid struct {
	Value utils.INTEGER `lb:0,ub:maxEUTRAUlFeaturesets`
}
