package ies

import "rrc/utils"

// CG-ConfigInfo-IEs-scgFailureInfo ::= SEQUENCE
type CgConfiginfoIesScgfailureinfo struct {
	Failuretype   CgConfiginfoIesScgfailureinfoFailuretype
	Measresultscg utils.OCTETSTRING
}
