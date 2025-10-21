package ies

import "rrc/utils"

// FlightPathInfoReportConfig-r15 ::= SEQUENCE
type FlightpathinforeportconfigR15 struct {
	MaxwaypointnumberR15 utils.INTEGER
	IncludetimestampR15  *utils.ENUMERATED
}
