package ies

import "rrc/utils"

// Phy-ParametersCommon-restricted-IAB-DU-BeamReception-r17 ::= ENUMERATED
type PhyParameterscommonRestrictedIabDuBeamreceptionR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonRestrictedIabDuBeamreceptionR17EnumeratedNothing = iota
	PhyParameterscommonRestrictedIabDuBeamreceptionR17EnumeratedSupported
)
