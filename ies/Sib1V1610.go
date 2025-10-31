package ies

import "rrc/utils"

// SIB1-v1610-IEs ::= SEQUENCE
type Sib1V1610 struct {
	IdlemodemeasurementseutraR16 *utils.BOOLEAN
	IdlemodemeasurementsnrR16    *utils.BOOLEAN
	PossiSchedulinginfoR16       *PossiSchedulinginfoR16
	Noncriticalextension         *Sib1V1630
}
