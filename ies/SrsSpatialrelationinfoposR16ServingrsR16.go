package ies

// SRS-SpatialRelationInfoPos-r16-servingRS-r16 ::= SEQUENCE
type SrsSpatialrelationinfoposR16ServingrsR16 struct {
	Servingcellid      *Servcellindex
	ReferencesignalR16 SrsSpatialrelationinfoposR16ServingrsR16ReferencesignalR16
}
