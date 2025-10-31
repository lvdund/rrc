package ies

// CSI-ResourceConfig ::= SEQUENCE
// Extensible
type CsiResourceconfig struct {
	CsiResourceconfigid         CsiResourceconfigid
	CsiRsResourcesetlist        *CsiResourceconfigCsiRsResourcesetlist
	BwpId                       BwpId
	Resourcetype                CsiResourceconfigResourcetype
	CsiSsbResourcesetlistextR17 *CsiSsbResourcesetid `ext`
}
