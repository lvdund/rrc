package ies

// FlightPathInfoReport-r15 ::= SEQUENCE
type FlightpathinforeportR15 struct {
	FlightpathR15 *[]WaypointlocationR15 `lb:1,ub:maxWayPointR15`
	Dummy         *FlightpathinforeportR15Dummy
}
