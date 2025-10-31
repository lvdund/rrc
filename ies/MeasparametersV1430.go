package ies

// MeasParameters-v1430 ::= SEQUENCE
type MeasparametersV1430 struct {
	CemeasurementsR14               *MeasparametersV1430CemeasurementsR14
	NcsgR14                         *MeasparametersV1430NcsgR14
	ShortmeasurementgapR14          *MeasparametersV1430ShortmeasurementgapR14
	PerservingcellmeasurementgapR14 *MeasparametersV1430PerservingcellmeasurementgapR14
	NonuniformgapR14                *MeasparametersV1430NonuniformgapR14
}
