package ies

import "rrc/utils"

// MeasConfigAppLayer-r17 ::= SEQUENCE
// Extensible
type MeasconfigapplayerR17 struct {
	MeasconfigapplayeridR17           MeasconfigapplayeridR17
	MeasconfigapplayercontainerR17    *utils.OCTETSTRING `lb:1,ub:8000`
	ServicetypeR17                    *MeasconfigapplayerR17ServicetypeR17
	PausereportingR17                 *utils.BOOLEAN
	TransmissionofsessionstartstopR17 *utils.BOOLEAN
	RanVisibleparametersR17           *utils.Setuprelease[RanVisibleparametersR17]
}
