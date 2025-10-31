package ies

import "rrc/utils"

// FeatureSetEUTRA-DownlinkId ::= utils.INTEGER (0..maxEUTRA-DL-FeatureSets)
type FeatureseteutraDownlinkid struct {
	Value utils.INTEGER `lb:0,ub:maxEUTRADlFeaturesets`
}
