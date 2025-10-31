package ies

// BWP-Uplink ::= SEQUENCE
// Extensible
type BwpUplink struct {
	BwpId        BwpId
	BwpCommon    *BwpUplinkcommon
	BwpDedicated *BwpUplinkdedicated
}
