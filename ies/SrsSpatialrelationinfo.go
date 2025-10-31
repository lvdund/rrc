package ies

// SRS-SpatialRelationInfo ::= SEQUENCE
type SrsSpatialrelationinfo struct {
	Servingcellid   *Servcellindex
	Referencesignal SrsSpatialrelationinfoReferencesignal
}
