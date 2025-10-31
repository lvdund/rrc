package ies

import "rrc/utils"

// RadioLinkMonitoringRS-Id ::= utils.INTEGER (0..maxNrofFailureDetectionResources-1)
type RadiolinkmonitoringrsId struct {
	Value utils.INTEGER `lb:0,ub:maxNrofFailureDetectionResources1`
}
