package ies

import "rrc/utils"

// InterFreqCarrierFreqInfo ::= SEQUENCE
// Extensible
type Interfreqcarrierfreqinfo struct {
	DlCarrierfreq                  ArfcnValuenr
	Frequencybandlist              *MultifrequencybandlistnrSib
	Frequencybandlistsul           *MultifrequencybandlistnrSib
	NrofssBlockstoaverage          *utils.INTEGER `lb:0,ub:maxNrofSSBlockstoaverage`
	AbsthreshssBlocksconsolidation *Thresholdnr
	Smtc                           *SsbMtc
	Ssbsubcarrierspacing           Subcarrierspacing
	SsbTomeasure                   *SsbTomeasure
	DerivessbIndexfromcell         utils.BOOLEAN
	SsRssiMeasurement              *SsRssiMeasurement
	QRxlevmin                      QRxlevmin
	QRxlevminsul                   *QRxlevmin
	QQualmin                       *QQualmin
	PMax                           *PMax
	TReselectionnr                 TReselection
	TReselectionnrSf               *Speedstatescalefactors
	ThreshxHighp                   Reselectionthreshold
	ThreshxLowp                    Reselectionthreshold
	ThreshxQ                       *InterfreqcarrierfreqinfoThreshxQ
	Cellreselectionpriority        *Cellreselectionpriority
	Cellreselectionsubpriority     *Cellreselectionsubpriority
	QOffsetfreq                    QOffsetrange
	Interfreqneighcelllist         *Interfreqneighcelllist
	Interfreqexcludedcelllist      *Interfreqexcludedcelllist
}
