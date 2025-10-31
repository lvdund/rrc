package ies

// InterFreqCarrierFreqInfo-NB-r13 ::= SEQUENCE
// Extensible
type InterfreqcarrierfreqinfoNbR13 struct {
	DlCarrierfreqR13          CarrierfreqNbR13
	QRxlevminR13              QRxlevmin
	QQualminR13               *QQualminR9
	PMaxR13                   *PMax
	QOffsetfreqR13            QOffsetrange
	InterfreqneighcelllistR13 *InterfreqneighcelllistNbR13
	InterfreqblackcelllistR13 *InterfreqblackcelllistNbR13
	MultibandinfolistR13      *MultibandinfolistNbR13
}
