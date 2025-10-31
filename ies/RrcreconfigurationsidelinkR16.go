package ies

import "rrc/utils"

// RRCReconfigurationSidelink-r16-IEs ::= SEQUENCE
type RrcreconfigurationsidelinkR16 struct {
	SlrbConfigtoaddmodlistR16  *[]SlrbConfigR16         `lb:1,ub:maxNrofSLRBR16`
	SlrbConfigtoreleaselistR16 *[]SlrbPc5ConfigindexR16 `lb:1,ub:maxNrofSLRBR16`
	SlMeasconfigR16            *utils.Setuprelease[SlMeasconfigR16]
	SlCsiRsConfigR16           *utils.Setuprelease[SlCsiRsConfigR16]
	SlResetconfigR16           *utils.BOOLEAN
	SlLatencyboundcsiReportR16 *utils.INTEGER `lb:0,ub:160`
	Latenoncriticalextension   *utils.OCTETSTRING
	Noncriticalextension       *RrcreconfigurationsidelinkV1700
}
