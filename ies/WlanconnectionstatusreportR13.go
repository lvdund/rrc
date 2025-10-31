package ies

import "rrc/utils"

// WLANConnectionStatusReport-r13-IEs ::= SEQUENCE
type WlanconnectionstatusreportR13 struct {
	WlanStatusR13            WlanStatusR13
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *WlanconnectionstatusreportV1430
}
