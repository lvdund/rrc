package ies

// InterFreqCarrierFreqInfo ::= SEQUENCE
// Extensible
type Interfreqcarrierfreqinfo struct {
	DlCarrierfreq           ArfcnValueeutra
	QRxlevmin               QRxlevmin
	PMax                    *PMax
	TReselectioneutra       TReselection
	TReselectioneutraSf     *Speedstatescalefactors
	ThreshxHigh             Reselectionthreshold
	ThreshxLow              Reselectionthreshold
	Allowedmeasbandwidth    Allowedmeasbandwidth
	Presenceantennaport1    Presenceantennaport1
	Cellreselectionpriority *Cellreselectionpriority
	Neighcellconfig         Neighcellconfig
	QOffsetfreq             QOffsetrange
	Interfreqneighcelllist  *Interfreqneighcelllist
	Interfreqblackcelllist  *Interfreqblackcelllist
}
