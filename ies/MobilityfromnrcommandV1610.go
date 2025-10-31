package ies

import "rrc/utils"

// MobilityFromNRCommand-v1610-IEs ::= SEQUENCE
type MobilityfromnrcommandV1610 struct {
	VoicefallbackindicationR16 *utils.BOOLEAN
	Noncriticalextension       *MobilityfromnrcommandV1610IesNoncriticalextension
}
