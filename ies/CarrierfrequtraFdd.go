package ies

import "rrc/utils"

// CarrierFreqUTRA-FDD ::= SEQUENCE
// Extensible
type CarrierfrequtraFdd struct {
	Carrierfreq             ArfcnValueutra
	Cellreselectionpriority *Cellreselectionpriority
	ThreshxHigh             Reselectionthreshold
	ThreshxLow              Reselectionthreshold
	QRxlevmin               utils.INTEGER `lb:0,ub:-13`
	PMaxutra                utils.INTEGER `lb:0,ub:33`
	QQualmin                utils.INTEGER `lb:0,ub:0`
}
