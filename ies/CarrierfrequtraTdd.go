package ies

import "rrc/utils"

// CarrierFreqUTRA-TDD ::= SEQUENCE
// Extensible
type CarrierfrequtraTdd struct {
	Carrierfreq             ArfcnValueutra
	Cellreselectionpriority *Cellreselectionpriority
	ThreshxHigh             Reselectionthreshold
	ThreshxLow              Reselectionthreshold
	QRxlevmin               utils.INTEGER `lb:0,ub:-13`
	PMaxutra                utils.INTEGER `lb:0,ub:33`
}
