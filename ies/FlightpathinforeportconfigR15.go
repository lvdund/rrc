package ies

import "rrc/utils"

// FlightPathInfoReportConfig-r15 ::= SEQUENCE
type FlightpathinforeportconfigR15 struct {
	MaxwaypointnumberR15 utils.INTEGER `lb:0,ub:maxWayPointR15`
	IncludetimestampR15  *bool
}
