package ies

// WLANConnectionStatusReport-v1430-IEs ::= SEQUENCE
type WlanconnectionstatusreportV1430 struct {
	WlanStatusV1430      WlanStatusV1430
	Noncriticalextension *WlanconnectionstatusreportV1430IesNoncriticalextension
}
