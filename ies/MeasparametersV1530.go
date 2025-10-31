package ies

// MeasParameters-v1530 ::= SEQUENCE
type MeasparametersV1530 struct {
	QoeMeasreportR15              *MeasparametersV1530QoeMeasreportR15
	QoeMtsiMeasreportR15          *MeasparametersV1530QoeMtsiMeasreportR15
	CaIdlemodemeasurementsR15     *MeasparametersV1530CaIdlemodemeasurementsR15
	CaIdlemodevalidityareaR15     *MeasparametersV1530CaIdlemodevalidityareaR15
	HeightmeasR15                 *MeasparametersV1530HeightmeasR15
	MultiplecellsmeasextensionR15 *MeasparametersV1530MultiplecellsmeasextensionR15
}
