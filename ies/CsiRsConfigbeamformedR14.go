package ies

import "rrc/utils"

// CSI-RS-ConfigBeamformed-r14 ::= SEQUENCE
type CsiRsConfigbeamformedR14 struct {
	CsiRsConfignzpidlistextR14              *[]CsiRsConfignzpidR13 `lb:1,ub:7`
	CsiImConfigidlistR14                    *[]CsiImConfigidR13    `lb:1,ub:8`
	PCAndcbsrPerresourceconfiglistR14       *[]PCAndcbsrPairR13    `lb:1,ub:8`
	AceFor4txPerresourceconfiglistR14       *[]utils.BOOLEAN       `lb:1,ub:7`
	AlternativecodebookenabledbeamformedR14 *bool
	ChannelmeasrestrictionR14               *CsiRsConfigbeamformedR14ChannelmeasrestrictionR14
	CsiRsConfignzpAplistR14                 *[]CsiRsConfignzpR11 `lb:1,ub:8`
	NzpResourceconfigoriginalV1430          *CsiRsConfigNzpV1430
	CsiRsNzpActivationR14                   *CsiRsConfignzpActivationR14
}
