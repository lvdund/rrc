package ies

import "rrc/utils"

// SIB2-intraFreqCellReselectionInfo ::= SEQUENCE
// Extensible
type Sib2Intrafreqcellreselectioninfo struct {
	QRxlevmin               QRxlevmin
	QRxlevminsul            *QRxlevmin
	QQualmin                *QQualmin
	SIntrasearchp           Reselectionthreshold
	SIntrasearchq           *Reselectionthresholdq
	TReselectionnr          TReselection
	Frequencybandlist       *MultifrequencybandlistnrSib
	Frequencybandlistsul    *MultifrequencybandlistnrSib
	PMax                    *PMax
	Smtc                    *SsbMtc
	SsRssiMeasurement       *SsRssiMeasurement
	SsbTomeasure            *SsbTomeasure
	DerivessbIndexfromcell  utils.BOOLEAN
	TReselectionnrSf        *Speedstatescalefactors
	Smtc2LpR16              *SsbMtc2LpR16
	SsbPositionqclCommonR16 *SsbPositionqclRelationR16
	SsbPositionqclCommonR17 *SsbPositionqclRelationR17
	Smtc4listR17            *SsbMtc4listR17
}
