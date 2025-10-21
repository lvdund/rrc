package ies

import "rrc/utils"

// CSI-RS-ConfigBeamformed-r14 ::= SEQUENCE
type CsiRsConfigbeamformedR14 struct {
	CsiRsConfignzpidlistextR14              *interface{}
	CsiImConfigidlistR14                    *interface{}
	PCAndcbsrPerresourceconfiglistR14       *interface{}
	AceFor4txPerresourceconfiglistR14       *interface{}
	AlternativecodebookenabledbeamformedR14 *utils.ENUMERATED
	ChannelmeasrestrictionR14               *utils.ENUMERATED
	CsiRsConfignzpAplistR14                 *interface{}
	NzpResourceconfigoriginalV1430          *CsiRsConfigNzpV1430
	CsiRsNzpActivationR14                   *CsiRsConfignzpActivationR14
}
