package ies

import "rrc/utils"

// CarrierFreqsInfoGERAN-commonInfo ::= SEQUENCE
type CarrierfreqsinfogeranCommoninfo struct {
	Cellreselectionpriority *Cellreselectionpriority
	NccPermitted            utils.BITSTRING `lb:8,ub:8`
	QRxlevmin               utils.INTEGER   `lb:0,ub:45`
	PMaxgeran               *utils.INTEGER  `lb:0,ub:39`
	ThreshxHigh             Reselectionthreshold
	ThreshxLow              Reselectionthreshold
}
