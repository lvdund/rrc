package ies

import "rrc/utils"

// RAN-VisibleMeasurements-r17 ::= SEQUENCE
// Extensible
type RanVisiblemeasurementsR17 struct {
	ApplayerbufferlevellistR17     *[]ApplayerbufferlevelR17 `lb:1,ub:8`
	PlayoutdelayformediastartupR17 *utils.INTEGER            `lb:0,ub:30000`
	PduSessionidlistR17            *[]PduSessionid           `lb:1,ub:maxNrofPDUSessionsR17`
}
