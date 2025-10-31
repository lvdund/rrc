package ies

// QCL-Info ::= SEQUENCE
// Extensible
type QclInfo struct {
	Cell            *Servcellindex
	BwpId           *BwpId
	Referencesignal QclInfoReferencesignal
	QclType         QclInfoQclType
}
