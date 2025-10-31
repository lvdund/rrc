package ies

import "rrc/utils"

// CG-ConfigInfo-v1560-IEs-scgFailureInfoEUTRA ::= SEQUENCE
type CgConfiginfoV1560IesScgfailureinfoeutra struct {
	Failuretypeeutra   CgConfiginfoV1560IesScgfailureinfoeutraFailuretypeeutra
	MeasresultscgEutra utils.OCTETSTRING
}
