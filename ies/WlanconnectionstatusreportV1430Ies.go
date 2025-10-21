package ies

import "rrc/utils"

// WLANConnectionStatusReport-v1430-IEs ::= SEQUENCE
type WlanconnectionstatusreportV1430Ies struct {
	WlanStatusV1430      WlanStatusV1430
	Noncriticalextension *interface{}
}
