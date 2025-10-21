package ies

import "rrc/utils"

// MeasParameters-v1530 ::= SEQUENCE
type MeasparametersV1530 struct {
	QoeMeasreportR15              *utils.ENUMERATED
	QoeMtsiMeasreportR15          *utils.ENUMERATED
	CaIdlemodemeasurementsR15     *utils.ENUMERATED
	CaIdlemodevalidityareaR15     *utils.ENUMERATED
	HeightmeasR15                 *utils.ENUMERATED
	MultiplecellsmeasextensionR15 *utils.ENUMERATED
}
