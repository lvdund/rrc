package ies

import "rrc/utils"

// RRCRelease-v1610-IEs ::= SEQUENCE
type RrcreleaseV1610 struct {
	VoicefallbackindicationR16 *utils.BOOLEAN
	MeasidleconfigR16          *utils.Setuprelease[MeasidleconfigdedicatedR16]
	Noncriticalextension       *RrcreleaseV1650
}
