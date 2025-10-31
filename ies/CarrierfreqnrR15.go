package ies

import "rrc/utils"

// CarrierFreqNR-r15 ::= SEQUENCE
// Extensible
type CarrierfreqnrR15 struct {
	CarrierfreqR15                ArfcnValuenrR15
	MultibandinfolistR15          *MultifrequencybandlistnrR15
	MultibandinfolistsulR15       *MultifrequencybandlistnrR15
	MeastimingconfigR15           *MtcSsbNrR15
	SubcarrierspacingssbR15       CarrierfreqnrR15SubcarrierspacingssbR15
	SsRssiMeasurementR15          *SsRssiMeasurementR15
	CellreselectionpriorityR15    *Cellreselectionpriority
	CellreselectionsubpriorityR15 *CellreselectionsubpriorityR13
	ThreshxHighR15                Reselectionthreshold
	ThreshxLowR15                 Reselectionthreshold
	ThreshxQR15                   *CarrierfreqnrR15ThreshxQR15
	QRxlevminR15                  utils.INTEGER  `lb:0,ub:-22`
	QRxlevminsulR15               *utils.INTEGER `lb:0,ub:-22`
	PMaxnrR15                     PMaxnrR15
	NsPmaxlistnrR15               *NsPmaxlistnrR15
	QQualminR15                   *utils.INTEGER `lb:0,ub:-12`
	DerivessbIndexfromcellR15     utils.BOOLEAN
	MaxrsIndexcellqualR15         *MaxrsIndexcellqualnrR15
	ThreshrsIndexR15              *ThresholdlistnrR15
}
