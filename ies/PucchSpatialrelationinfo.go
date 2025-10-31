package ies

// PUCCH-SpatialRelationInfo ::= SEQUENCE
type PucchSpatialrelationinfo struct {
	PucchSpatialrelationinfoid PucchSpatialrelationinfoid
	Servingcellid              *Servcellindex
	Referencesignal            PucchSpatialrelationinfoReferencesignal
	PucchPathlossreferencersId PucchPathlossreferencersId
	P0PucchId                  P0PucchId
	Closedloopindex            PucchSpatialrelationinfoClosedloopindex
}
