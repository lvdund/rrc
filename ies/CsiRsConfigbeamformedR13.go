package ies

import "rrc/utils"

// CSI-RS-ConfigBeamformed-r13 ::= SEQUENCE
type CsiRsConfigbeamformedR13 struct {
	CsiRsConfignzpidlistextR13              *[]CsiRsConfignzpidR13 `lb:1,ub:7`
	CsiImConfigidlistR13                    *[]CsiImConfigidR13    `lb:1,ub:8`
	PCAndcbsrPerresourceconfiglistR13       *[]PCAndcbsrPairR13    `lb:1,ub:8`
	AceFor4txPerresourceconfiglistR13       *[]utils.BOOLEAN       `lb:1,ub:7`
	AlternativecodebookenabledbeamformedR13 *bool
	ChannelmeasrestrictionR13               *CsiRsConfigbeamformedR13ChannelmeasrestrictionR13
}
