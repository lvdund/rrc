package ies

// WayPointLocation-r15 ::= SEQUENCE
type WaypointlocationR15 struct {
	WaypointlocationR15 LocationinfoR10
	TimestampR15        *AbsolutetimeinfoR10
}
