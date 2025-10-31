package ies

import "rrc/utils"

// Phy-ParametersCommon-fdm-SoftResourceAvailability-DynamicIndication-r17 ::= ENUMERATED
type PhyParameterscommonFdmSoftresourceavailabilityDynamicindicationR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonFdmSoftresourceavailabilityDynamicindicationR17EnumeratedNothing = iota
	PhyParameterscommonFdmSoftresourceavailabilityDynamicindicationR17EnumeratedSupported
)
