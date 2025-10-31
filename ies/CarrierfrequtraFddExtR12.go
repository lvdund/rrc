package ies

import "rrc/utils"

// CarrierFreqUTRA-FDD-Ext-r12 ::= SEQUENCE
// Extensible
type CarrierfrequtraFddExtR12 struct {
	CarrierfreqR12             ArfcnValueutra
	CellreselectionpriorityR12 *Cellreselectionpriority
	ThreshxHighR12             Reselectionthreshold
	ThreshxLowR12              Reselectionthreshold
	QRxlevminR12               utils.INTEGER `lb:0,ub:-13`
	PMaxutraR12                utils.INTEGER `lb:0,ub:33`
	QQualminR12                utils.INTEGER `lb:0,ub:0`
	ThreshxQR12                *CarrierfrequtraFddExtR12ThreshxQR12
	MultibandinfolistR12       *[]FreqbandindicatorUtraFdd `lb:1,ub:maxMultiBands`
	ReducedmeasperformanceR12  *bool
}
