package ies

import "rrc/utils"

// MeasIdleCarrierNR-r16-ssb-MeasConfig-r16 ::= SEQUENCE
type MeasidlecarriernrR16SsbMeasconfigR16 struct {
	MaxrsIndexcellqualR16     *MaxrsIndexcellqualnrR15
	ThreshrsIndexR16          *ThresholdlistnrR15
	MeastimingconfigR16       *MtcSsbNrR15
	SsbTomeasureR16           *SsbTomeasureR15
	DerivessbIndexfromcellR16 utils.BOOLEAN
	SsRssiMeasurementR16      *SsRssiMeasurementR15
}
