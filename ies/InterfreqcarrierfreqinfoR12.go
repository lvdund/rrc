package ies

import "rrc/utils"

// InterFreqCarrierFreqInfo-r12 ::= SEQUENCE
// Extensible
type InterfreqcarrierfreqinfoR12 struct {
	DlCarrierfreqR12            ArfcnValueeutraR9
	QRxlevminR12                QRxlevmin
	PMaxR12                     *PMax
	TReselectioneutraR12        TReselection
	TReselectioneutraSfR12      *Speedstatescalefactors
	ThreshxHighR12              Reselectionthreshold
	ThreshxLowR12               Reselectionthreshold
	AllowedmeasbandwidthR12     Allowedmeasbandwidth
	Presenceantennaport1R12     Presenceantennaport1
	CellreselectionpriorityR12  *Cellreselectionpriority
	NeighcellconfigR12          Neighcellconfig
	QOffsetfreqR12              QOffsetrange
	InterfreqneighcelllistR12   *Interfreqneighcelllist
	InterfreqblackcelllistR12   *Interfreqblackcelllist
	QQualminR12                 *QQualminR9
	ThreshxQR12                 *interface{}
	QQualminwbR12               *QQualminR9
	MultibandinfolistR12        *MultibandinfolistR11
	ReducedmeasperformanceR12   *utils.ENUMERATED
	QQualminrsrqOnallsymbolsR12 *QQualminR9
}
