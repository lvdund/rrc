package ies

// SystemInformationBlockType3-NB-r13-intraFreqCellReselectionInfo-r13 ::= SEQUENCE
type Systeminformationblocktype3NbR13IntrafreqcellreselectioninfoR13 struct {
	QRxlevminR13     QRxlevmin
	QQualminR13      *QQualminR9
	PMaxR13          *PMax
	SIntrasearchpR13 Reselectionthreshold
	TReselectionR13  TReselectionNbR13
}
