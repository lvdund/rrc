package ies

import "rrc/utils"

// ReferenceSignalConfig ::= SEQUENCE
type Referencesignalconfig struct {
	SsbConfigmobility           *SsbConfigmobility
	CsiRsResourceconfigmobility *utils.Setuprelease[CsiRsResourceconfigmobility]
}
