package ies

import "rrc/utils"

// MeasParameters-v1430 ::= SEQUENCE
type MeasparametersV1430 struct {
	CemeasurementsR14               *utils.ENUMERATED
	NcsgR14                         *utils.ENUMERATED
	ShortmeasurementgapR14          *utils.ENUMERATED
	PerservingcellmeasurementgapR14 *utils.ENUMERATED
	NonuniformgapR14                *utils.ENUMERATED
}
