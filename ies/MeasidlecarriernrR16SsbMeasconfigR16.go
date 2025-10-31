package ies

import "rrc/utils"

// MeasIdleCarrierNR-r16-ssb-MeasConfig-r16 ::= SEQUENCE
type MeasidlecarriernrR16SsbMeasconfigR16 struct {
	NrofssBlockstoaverageR16          *utils.INTEGER `lb:0,ub:maxNrofSSBlockstoaverage`
	AbsthreshssBlocksconsolidationR16 *Thresholdnr
	SmtcR16                           *SsbMtc
	SsbTomeasureR16                   *SsbTomeasure
	DerivessbIndexfromcellR16         utils.BOOLEAN
	SsRssiMeasurementR16              *SsRssiMeasurement
}
