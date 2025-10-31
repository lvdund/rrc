package ies

import "rrc/utils"

// BeamLinkMonitoringRS-Id-r17 ::= utils.INTEGER (0..maxNrofFailureDetectionResources-1-r17)
type BeamlinkmonitoringrsIdR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofFailureDetectionResources1R17`
}
