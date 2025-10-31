package ies

import "rrc/utils"

// CarrierFreqEUTRA ::= SEQUENCE
type Carrierfreqeutra struct {
	Carrierfreq                ArfcnValueeutra
	EutraMultibandinfolist     *EutraMultibandinfolist
	EutraFreqneighcelllist     *EutraFreqneighcelllist
	EutraExcludedcelllist      *EutraFreqexcludedcelllist
	Allowedmeasbandwidth       EutraAllowedmeasbandwidth
	Presenceantennaport1       EutraPresenceantennaport1
	Cellreselectionpriority    *Cellreselectionpriority
	Cellreselectionsubpriority *Cellreselectionsubpriority
	ThreshxHigh                Reselectionthreshold
	ThreshxLow                 Reselectionthreshold
	QRxlevmin                  utils.INTEGER `lb:0,ub:-22`
	QQualmin                   utils.INTEGER `lb:0,ub:-3`
	PMaxeutra                  utils.INTEGER `lb:0,ub:33`
	ThreshxQ                   *CarrierfreqeutraThreshxQ
}
