package ies

// WLAN-OffloadInfoPerPLMN-r12 ::= SEQUENCE
// Extensible
type WlanOffloadinfoperplmnR12 struct {
	WlanOffloadconfigcommonR12 *WlanOffloadconfigR12
	WlanIdListR12              *WlanIdListR12
}
