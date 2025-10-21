package ies

import "rrc/utils"

// WLANConnectionStatusReport-r13-IEs ::= SEQUENCE
type WlanconnectionstatusreportR13Ies struct {
	WlanStatusR13            WlanStatusR13
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *WlanconnectionstatusreportV1430Ies
}
