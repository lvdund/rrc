package ies

// BWP-Downlink ::= SEQUENCE
// Extensible
type BwpDownlink struct {
	BwpId        BwpId
	BwpCommon    *BwpDownlinkcommon
	BwpDedicated *BwpDownlinkdedicated
}
