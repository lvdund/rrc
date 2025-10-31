package ies

import "rrc/utils"

// CarrierFreqUTRA-TDD-r12 ::= SEQUENCE
// Extensible
type CarrierfrequtraTddR12 struct {
	CarrierfreqR12             ArfcnValueutra
	CellreselectionpriorityR12 *Cellreselectionpriority
	ThreshxHighR12             Reselectionthreshold
	ThreshxLowR12              Reselectionthreshold
	QRxlevminR12               utils.INTEGER `lb:0,ub:-13`
	PMaxutraR12                utils.INTEGER `lb:0,ub:33`
	ReducedmeasperformanceR12  *bool
}
